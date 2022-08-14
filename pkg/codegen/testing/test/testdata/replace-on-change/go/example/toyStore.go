// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ToyStore struct {
	pulumi.CustomResourceState

	Chew   ChewPtrOutput  `pulumi:"chew"`
	Laser  LaserPtrOutput `pulumi:"laser"`
	Stuff  ToyArrayOutput `pulumi:"stuff"`
	Wanted ToyArrayOutput `pulumi:"wanted"`
}

// NewToyStore registers a new resource with the given unique name, arguments, and options.
func NewToyStore(ctx *pulumi.Context,
	name string, args *ToyStoreArgs, opts ...pulumi.ResourceOption) (*ToyStore, error) {
	if args == nil {
		args = &ToyStoreArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"chew.owner",
		"laser.batteries",
		"stuff[*].associated.color",
		"stuff[*].color",
		"wanted[*]",
	})
	opts = append(opts, replaceOnChanges)
	var resource ToyStore
	err := ctx.RegisterResource("example::ToyStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetToyStore gets an existing ToyStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetToyStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ToyStoreState, opts ...pulumi.ResourceOption) (*ToyStore, error) {
	var resource ToyStore
	err := ctx.ReadResource("example::ToyStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ToyStore resources.
type toyStoreState struct {
}

type ToyStoreState struct {
}

func (ToyStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*toyStoreState)(nil)).Elem()
}

type toyStoreArgs struct {
}

// The set of arguments for constructing a ToyStore resource.
type ToyStoreArgs struct {
}

func (ToyStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*toyStoreArgs)(nil)).Elem()
}

type ToyStoreInput interface {
	pulumi.Input

	ToToyStoreOutput() ToyStoreOutput
	ToToyStoreOutputWithContext(ctx context.Context) ToyStoreOutput
}

func (*ToyStore) ElementType() reflect.Type {
	return reflect.TypeOf((**ToyStore)(nil)).Elem()
}

func (i *ToyStore) ToToyStoreOutput() ToyStoreOutput {
	return i.ToToyStoreOutputWithContext(context.Background())
}

func (i *ToyStore) ToToyStoreOutputWithContext(ctx context.Context) ToyStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyStoreOutput)
}

// ToyStoreArrayInput is an input type that accepts ToyStoreArray and ToyStoreArrayOutput values.
// You can construct a concrete instance of `ToyStoreArrayInput` via:
//
//          ToyStoreArray{ ToyStoreArgs{...} }
type ToyStoreArrayInput interface {
	pulumi.Input

	ToToyStoreArrayOutput() ToyStoreArrayOutput
	ToToyStoreArrayOutputWithContext(context.Context) ToyStoreArrayOutput
}

type ToyStoreArray []ToyStoreInput

func (ToyStoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ToyStore)(nil)).Elem()
}

func (i ToyStoreArray) ToToyStoreArrayOutput() ToyStoreArrayOutput {
	return i.ToToyStoreArrayOutputWithContext(context.Background())
}

func (i ToyStoreArray) ToToyStoreArrayOutputWithContext(ctx context.Context) ToyStoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyStoreArrayOutput)
}

// ToyStoreMapInput is an input type that accepts ToyStoreMap and ToyStoreMapOutput values.
// You can construct a concrete instance of `ToyStoreMapInput` via:
//
//          ToyStoreMap{ "key": ToyStoreArgs{...} }
type ToyStoreMapInput interface {
	pulumi.Input

	ToToyStoreMapOutput() ToyStoreMapOutput
	ToToyStoreMapOutputWithContext(context.Context) ToyStoreMapOutput
}

type ToyStoreMap map[string]ToyStoreInput

func (ToyStoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ToyStore)(nil)).Elem()
}

func (i ToyStoreMap) ToToyStoreMapOutput() ToyStoreMapOutput {
	return i.ToToyStoreMapOutputWithContext(context.Background())
}

func (i ToyStoreMap) ToToyStoreMapOutputWithContext(ctx context.Context) ToyStoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ToyStoreMapOutput)
}

type ToyStoreOutput struct{ *pulumi.OutputState }

func (ToyStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ToyStore)(nil)).Elem()
}

func (o ToyStoreOutput) ToToyStoreOutput() ToyStoreOutput {
	return o
}

func (o ToyStoreOutput) ToToyStoreOutputWithContext(ctx context.Context) ToyStoreOutput {
	return o
}

func (o ToyStoreOutput) Chew() ChewPtrOutput {
	return o.ApplyT(func(v *ToyStore) ChewPtrOutput { return v.Chew }).(ChewPtrOutput)
}

func (o ToyStoreOutput) Laser() LaserPtrOutput {
	return o.ApplyT(func(v *ToyStore) LaserPtrOutput { return v.Laser }).(LaserPtrOutput)
}

func (o ToyStoreOutput) Stuff() ToyArrayOutput {
	return o.ApplyT(func(v *ToyStore) ToyArrayOutput { return v.Stuff }).(ToyArrayOutput)
}

func (o ToyStoreOutput) Wanted() ToyArrayOutput {
	return o.ApplyT(func(v *ToyStore) ToyArrayOutput { return v.Wanted }).(ToyArrayOutput)
}

type ToyStoreArrayOutput struct{ *pulumi.OutputState }

func (ToyStoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ToyStore)(nil)).Elem()
}

func (o ToyStoreArrayOutput) ToToyStoreArrayOutput() ToyStoreArrayOutput {
	return o
}

func (o ToyStoreArrayOutput) ToToyStoreArrayOutputWithContext(ctx context.Context) ToyStoreArrayOutput {
	return o
}

func (o ToyStoreArrayOutput) Index(i pulumi.IntInput) ToyStoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ToyStore {
		return vs[0].([]*ToyStore)[vs[1].(int)]
	}).(ToyStoreOutput)
}

type ToyStoreMapOutput struct{ *pulumi.OutputState }

func (ToyStoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ToyStore)(nil)).Elem()
}

func (o ToyStoreMapOutput) ToToyStoreMapOutput() ToyStoreMapOutput {
	return o
}

func (o ToyStoreMapOutput) ToToyStoreMapOutputWithContext(ctx context.Context) ToyStoreMapOutput {
	return o
}

func (o ToyStoreMapOutput) MapIndex(k pulumi.StringInput) ToyStoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ToyStore {
		return vs[0].(map[string]*ToyStore)[vs[1].(string)]
	}).(ToyStoreOutput)
}

func init() {
	pulumi.RegisterOutputType(ToyStoreOutput{})
	pulumi.RegisterOutputType(ToyStoreArrayOutput{})
	pulumi.RegisterOutputType(ToyStoreMapOutput{})
}
