// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Another failing example. A list of SSIS object metadata.
// API Version: 2018-06-01.
func GetIntegrationRuntimeObjectMetadatum(ctx *pulumi.Context, args *GetIntegrationRuntimeObjectMetadatumArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeObjectMetadatumResult, error) {
	var rv GetIntegrationRuntimeObjectMetadatumResult
	err := ctx.Invoke("mypkg::getIntegrationRuntimeObjectMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeObjectMetadatumArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath *string `pulumi:"metadataPath"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A list of SSIS object metadata.
type GetIntegrationRuntimeObjectMetadatumResult struct {
	// The link to the next page of results, if any remaining results exist.
	NextLink *string `pulumi:"nextLink"`
	// List of SSIS object metadata.
	Value []interface{} `pulumi:"value"`
}

func GetIntegrationRuntimeObjectMetadatumOutput(ctx *pulumi.Context, args GetIntegrationRuntimeObjectMetadatumOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeObjectMetadatumResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIntegrationRuntimeObjectMetadatumResult, error) {
			args := v.(GetIntegrationRuntimeObjectMetadatumArgs)
			r, err := GetIntegrationRuntimeObjectMetadatum(ctx, &args, opts...)
			var s GetIntegrationRuntimeObjectMetadatumResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIntegrationRuntimeObjectMetadatumResultOutput)
}

type GetIntegrationRuntimeObjectMetadatumOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	// Metadata path.
	MetadataPath pulumi.StringPtrInput `pulumi:"metadataPath"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetIntegrationRuntimeObjectMetadatumOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumArgs)(nil)).Elem()
}

// A list of SSIS object metadata.
type GetIntegrationRuntimeObjectMetadatumResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeObjectMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeObjectMetadatumResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToGetIntegrationRuntimeObjectMetadatumResultOutput() GetIntegrationRuntimeObjectMetadatumResultOutput {
	return o
}

func (o GetIntegrationRuntimeObjectMetadatumResultOutput) ToGetIntegrationRuntimeObjectMetadatumResultOutputWithContext(ctx context.Context) GetIntegrationRuntimeObjectMetadatumResultOutput {
	return o
}

// The link to the next page of results, if any remaining results exist.
func (o GetIntegrationRuntimeObjectMetadatumResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeObjectMetadatumResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of SSIS object metadata.
func (o GetIntegrationRuntimeObjectMetadatumResultOutput) Value() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeObjectMetadatumResult) []interface{} { return v.Value }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIntegrationRuntimeObjectMetadatumResultOutput{})
}
