// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package foo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TopLevel struct {
	Buzz *string `pulumi:"buzz"`
}

// TopLevelInput is an input type that accepts TopLevelArgs and TopLevelOutput values.
// You can construct a concrete instance of `TopLevelInput` via:
//
//          TopLevelArgs{...}
type TopLevelInput interface {
	pulumi.Input

	ToTopLevelOutput() TopLevelOutput
	ToTopLevelOutputWithContext(context.Context) TopLevelOutput
}

type TopLevelArgs struct {
	Buzz pulumi.StringPtrInput `pulumi:"buzz"`
}

func (TopLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopLevel)(nil)).Elem()
}

func (i TopLevelArgs) ToTopLevelOutput() TopLevelOutput {
	return i.ToTopLevelOutputWithContext(context.Background())
}

func (i TopLevelArgs) ToTopLevelOutputWithContext(ctx context.Context) TopLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopLevelOutput)
}

func (i TopLevelArgs) ToTopLevelPtrOutput() TopLevelPtrOutput {
	return i.ToTopLevelPtrOutputWithContext(context.Background())
}

func (i TopLevelArgs) ToTopLevelPtrOutputWithContext(ctx context.Context) TopLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopLevelOutput).ToTopLevelPtrOutputWithContext(ctx)
}

// TopLevelPtrInput is an input type that accepts TopLevelArgs, TopLevelPtr and TopLevelPtrOutput values.
// You can construct a concrete instance of `TopLevelPtrInput` via:
//
//          TopLevelArgs{...}
//
//  or:
//
//          nil
type TopLevelPtrInput interface {
	pulumi.Input

	ToTopLevelPtrOutput() TopLevelPtrOutput
	ToTopLevelPtrOutputWithContext(context.Context) TopLevelPtrOutput
}

type topLevelPtrType TopLevelArgs

func TopLevelPtr(v *TopLevelArgs) TopLevelPtrInput {
	return (*topLevelPtrType)(v)
}

func (*topLevelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TopLevel)(nil)).Elem()
}

func (i *topLevelPtrType) ToTopLevelPtrOutput() TopLevelPtrOutput {
	return i.ToTopLevelPtrOutputWithContext(context.Background())
}

func (i *topLevelPtrType) ToTopLevelPtrOutputWithContext(ctx context.Context) TopLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopLevelPtrOutput)
}

type TopLevelOutput struct{ *pulumi.OutputState }

func (TopLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopLevel)(nil)).Elem()
}

func (o TopLevelOutput) ToTopLevelOutput() TopLevelOutput {
	return o
}

func (o TopLevelOutput) ToTopLevelOutputWithContext(ctx context.Context) TopLevelOutput {
	return o
}

func (o TopLevelOutput) ToTopLevelPtrOutput() TopLevelPtrOutput {
	return o.ToTopLevelPtrOutputWithContext(context.Background())
}

func (o TopLevelOutput) ToTopLevelPtrOutputWithContext(ctx context.Context) TopLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TopLevel) *TopLevel {
		return &v
	}).(TopLevelPtrOutput)
}

func (o TopLevelOutput) Buzz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TopLevel) *string { return v.Buzz }).(pulumi.StringPtrOutput)
}

type TopLevelPtrOutput struct{ *pulumi.OutputState }

func (TopLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopLevel)(nil)).Elem()
}

func (o TopLevelPtrOutput) ToTopLevelPtrOutput() TopLevelPtrOutput {
	return o
}

func (o TopLevelPtrOutput) ToTopLevelPtrOutputWithContext(ctx context.Context) TopLevelPtrOutput {
	return o
}

func (o TopLevelPtrOutput) Elem() TopLevelOutput {
	return o.ApplyT(func(v *TopLevel) TopLevel {
		if v != nil {
			return *v
		}
		var ret TopLevel
		return ret
	}).(TopLevelOutput)
}

func (o TopLevelPtrOutput) Buzz() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopLevel) *string {
		if v == nil {
			return nil
		}
		return v.Buzz
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopLevelInput)(nil)).Elem(), TopLevelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopLevelPtrInput)(nil)).Elem(), TopLevelArgs{})
	pulumi.RegisterOutputType(TopLevelOutput{})
	pulumi.RegisterOutputType(TopLevelPtrOutput{})
}
