// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example
{
    [ExampleResourceType("example::NoRecursive")]
    public partial class NoRecursive : global::Pulumi.CustomResource
    {
        [Output("rec")]
        public Output<Outputs.Rec?> Rec { get; private set; } = null!;

        [Output("replace")]
        public Output<string?> Replace { get; private set; } = null!;


        /// <summary>
        /// Create a NoRecursive resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NoRecursive(string name, NoRecursiveArgs? args = null, CustomResourceOptions? options = null)
            : base("example::NoRecursive", name, args ?? new NoRecursiveArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NoRecursive(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("example::NoRecursive", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "replace",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NoRecursive resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NoRecursive Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NoRecursive(name, id, options);
        }
    }

    public sealed class NoRecursiveArgs : global::Pulumi.ResourceArgs
    {
        public NoRecursiveArgs()
        {
        }
        public static new NoRecursiveArgs Empty => new NoRecursiveArgs();
    }
}
