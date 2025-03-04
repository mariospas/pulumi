﻿// Copyright 2016-2021, Pulumi Corporation

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Threading.Tasks;
using Google.Protobuf.WellKnownTypes;

namespace Pulumi
{
    public partial class Deployment
    {
        private async Task<PrepareResult> PrepareResourceAsync(
            string label, Resource res, bool custom, bool remote,
            ResourceArgs args, ResourceOptions options)
        {
            // Before we can proceed, all our dependencies must be finished.
            var type = res.GetResourceType();
            var name = res.GetResourceName();

            LogExcessive($"Gathering explicit dependencies: t={type}, name={name}, custom={custom}, remote={remote}");
            var explicitDirectDependencies = new HashSet<Resource>(
                await GatherExplicitDependenciesAsync(options.DependsOn).ConfigureAwait(false));
            LogExcessive($"Gathered explicit dependencies: t={type}, name={name}, custom={custom}, remote={remote}");

            // Serialize out all our props to their final values.  In doing so, we'll also collect all
            // the Resources pointed to by any Dependency objects we encounter, adding them to 'propertyDependencies'.
            LogExcessive($"Serializing properties: t={type}, name={name}, custom={custom}, remote={remote}");
            var dictionary = await args.ToDictionaryAsync().ConfigureAwait(false);
            var (serializedProps, propertyToDirectDependencies) =
                await SerializeResourcePropertiesAsync(
                        label,
                        dictionary,
                        await this.MonitorSupportsResourceReferences().ConfigureAwait(false),
                        keepOutputValues: remote && await MonitorSupportsOutputValues().ConfigureAwait(false)).ConfigureAwait(false);
            LogExcessive($"Serialized properties: t={type}, name={name}, custom={custom}, remote={remote}");

            // Wait for the parent to complete.
            // If no parent was provided, parent to the root resource.
            LogExcessive($"Getting parent urn: t={type}, name={name}, custom={custom}, remote={remote}");
            var parentUrn = options.Parent != null
                ? await options.Parent.Urn.GetValueAsync(whenUnknown: default!).ConfigureAwait(false)
                : await GetRootResourceAsync(type).ConfigureAwait(false);
            LogExcessive($"Got parent urn: t={type}, name={name}, custom={custom}, remote={remote}");

            string? providerRef = null;
            if (custom)
            {
                var customOpts = options as CustomResourceOptions;
                providerRef = await ProviderResource.RegisterAsync(customOpts?.Provider).ConfigureAwait(false);

                // Note: because of the hard distinction between custom and
                // component resources, we don't fully mirror the behavior found
                // in other SDKs. Because custom resources are passed with
                // CustomResourceOptions, there is no Providers list. This makes
                // it nonsensical to find a candidate for Provider in Providers.

                if (providerRef == null)
                {
                    var t = res.GetResourceType();
                    var parentRef = customOpts?.Parent?.GetProvider(t);
                    providerRef = await ProviderResource.RegisterAsync(parentRef).ConfigureAwait(false);
                }
            }

            var providerRefs = new Dictionary<string, string>();
            if (remote && options is ComponentResourceOptions componentOpts)
            {
                if (componentOpts.Provider != null)
                {
                    var duplicate = false;
                    foreach (var p in componentOpts.Providers)
                    {
                        if (p.Package == componentOpts.Provider.Package)
                        {
                            duplicate = true;
                            await _logger.WarnAsync(
                                $"Conflict between provider and providers field for package '{p.Package}'. "+
                                "This behavior is depreciated, and will turn into an error July 2022. "+
                                "For more information, see https://github.com/mariospas/pulumi/issues/8799.", res)
                                .ConfigureAwait(false);
                        }
                    }
                    if (!duplicate)
                    {
                        componentOpts.Providers.Add(componentOpts.Provider);
                    }
                }

                foreach (var provider in componentOpts.Providers)
                {
                    var pref = await ProviderResource.RegisterAsync(provider).ConfigureAwait(false);
                    if (pref != null)
                    {
                        providerRefs.Add(provider.Package, pref);
                    }
                }
            }

            // Collect the URNs for explicit/implicit dependencies for the engine so that it can understand
            // the dependency graph and optimize operations accordingly.

            // The list of all dependencies (implicit or explicit).
            var allDirectDependencies = new HashSet<Resource>(explicitDirectDependencies);

            var allDirectDependencyUrns = await GetAllTransitivelyReferencedResourceUrnsAsync(explicitDirectDependencies).ConfigureAwait(false);
            var propertyToDirectDependencyUrns = new Dictionary<string, HashSet<string>>();

            foreach (var (propertyName, directDependencies) in propertyToDirectDependencies)
            {
                allDirectDependencies.AddRange(directDependencies);

                var urns = await GetAllTransitivelyReferencedResourceUrnsAsync(directDependencies).ConfigureAwait(false);
                allDirectDependencyUrns.AddRange(urns);
                propertyToDirectDependencyUrns[propertyName] = urns;
            }

            // Wait for all aliases. Note that we use 'res._aliases' instead of 'options.aliases' as
            // the former has been processed in the Resource constructor prior to calling
            // 'registerResource' - both adding new inherited aliases and simplifying aliases down
            // to URNs.
            var aliases = new List<string>();
            var uniqueAliases = new HashSet<string>();
            foreach (var alias in res._aliases)
            {
                var aliasVal = await alias.ToOutput().GetValueAsync(whenUnknown: "").ConfigureAwait(false);
                if (aliasVal != "" && uniqueAliases.Add(aliasVal))
                {
                    aliases.Add(aliasVal);
                }
            }

            return new PrepareResult(
                serializedProps,
                parentUrn ?? "",
                providerRef ?? "",
                providerRefs,
                allDirectDependencyUrns,
                propertyToDirectDependencyUrns,
                aliases);

            void LogExcessive(string message)
            {
                if (_excessiveDebugOutput)
                    Log.Debug(message);
            }
        }

        private static Task<ImmutableArray<Resource>> GatherExplicitDependenciesAsync(InputList<Resource> resources)
            => resources.ToOutput().GetValueAsync(whenUnknown: ImmutableArray<Resource>.Empty);

        internal static async Task<HashSet<string>> GetAllTransitivelyReferencedResourceUrnsAsync(
            HashSet<Resource> resources)
        {
            // Go through 'resources', but transitively walk through **Component** resources, collecting any
            // of their child resources.  This way, a Component acts as an aggregation really of all the
            // reachable resources it parents.  This walking will stop when it hits custom resources.
            //
            // This function also terminates at remote components, whose children are not known to the Node SDK directly.
            // Remote components will always wait on all of their children, so ensuring we return the remote component
            // itself here and waiting on it will accomplish waiting on all of it's children regardless of whether they
            // are returned explicitly here.
            //
            // In other words, if we had:
            //
            //                  Comp1
            //              /     |     \
            //          Cust1   Comp2  Remote1
            //                  /   \       \
            //              Cust2   Cust3  Comp3
            //              /                 \
            //          Cust4                Cust5
            //
            // Then the transitively reachable resources of Comp1 will be [Cust1, Cust2, Cust3, Remote1]. It
            // will *not* include:
            // * Cust4 because it is a child of a custom resource
            // * Comp2 because it is a non-remote component resoruce
            // * Comp3 and Cust5 because Comp3 is a child of a remote component resource
            var transitivelyReachableResources = GetTransitivelyReferencedChildResourcesOfComponentResources(resources);

            var transitivelyReachableCustomResources = transitivelyReachableResources.Where(res =>
            {
                switch (res)
                {
                    case CustomResource _: return true;
                    case ComponentResource component: return component.remote;
                    default: return false; // Unreachable
                }
            });
            var tasks = transitivelyReachableCustomResources.Select(r => r.Urn.GetValueAsync(whenUnknown: ""));
            var urns = await Task.WhenAll(tasks).ConfigureAwait(false);
            return new HashSet<string>(urns.Where(urn => !string.IsNullOrEmpty(urn)));
        }

        /// <summary>
        /// Recursively walk the resources passed in, returning them and all resources reachable from
        /// <see cref="Resource.ChildResources"/> through any **Component** resources we encounter.
        /// </summary>
        private static HashSet<Resource> GetTransitivelyReferencedChildResourcesOfComponentResources(HashSet<Resource> resources)
        {
            // Recursively walk the dependent resources through their children, adding them to the result set.
            var result = new HashSet<Resource>();
            AddTransitivelyReferencedChildResourcesOfComponentResources(resources, result);
            return result;
        }

        private static void AddTransitivelyReferencedChildResourcesOfComponentResources(HashSet<Resource> resources, HashSet<Resource> result)
        {
            foreach (var resource in resources)
            {
                if (result.Add(resource))
                {
                    if (resource is ComponentResource)
                    {
                        HashSet<Resource> childResources;
                        lock (resource.ChildResources)
                        {
                            childResources = new HashSet<Resource>(resource.ChildResources);
                        }
                        AddTransitivelyReferencedChildResourcesOfComponentResources(childResources, result);
                    }
                }
            }
        }

        private readonly struct PrepareResult
        {
            public readonly Struct SerializedProps;
            public readonly string ParentUrn;
            public readonly string ProviderRef;
            public readonly Dictionary<string, string> ProviderRefs;
            public readonly HashSet<string> AllDirectDependencyUrns;
            public readonly Dictionary<string, HashSet<string>> PropertyToDirectDependencyUrns;
            public readonly List<string> Aliases;

            public PrepareResult(Struct serializedProps, string parentUrn, string providerRef, Dictionary<string, string> providerRefs, HashSet<string> allDirectDependencyUrns, Dictionary<string, HashSet<string>> propertyToDirectDependencyUrns, List<string> aliases)
            {
                SerializedProps = serializedProps;
                ParentUrn = parentUrn;
                ProviderRef = providerRef;
                ProviderRefs = providerRefs;
                AllDirectDependencyUrns = allDirectDependencyUrns;
                PropertyToDirectDependencyUrns = propertyToDirectDependencyUrns;
                Aliases = aliases;
            }
        }
    }
}
