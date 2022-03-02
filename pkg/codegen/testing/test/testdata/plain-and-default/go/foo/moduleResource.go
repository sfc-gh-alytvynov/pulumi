// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package foo

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ModuleResource struct {
	pulumi.CustomResourceState
}

// NewModuleResource registers a new resource with the given unique name, arguments, and options.
func NewModuleResource(ctx *pulumi.Context,
	name string, args *ModuleResourceArgs, opts ...pulumi.ResourceOption) (*ModuleResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if isZero(args.Optional_bool) {
		args.Optional_bool = pulumi.BoolPtr(true)
	}
	args.Optional_const = pulumi.StringPtr("val")
	if isZero(args.Optional_enum) {
		args.Optional_enum = EnumThing(8)
	}
	if isZero(args.Optional_number) {
		args.Optional_number = pulumi.Float64Ptr(42.0)
	}
	if isZero(args.Optional_string) {
		args.Optional_string = pulumi.StringPtr("buzzer")
	}
	if isZero(args.Plain_optional_bool) {
		plain_optional_bool_ := true
		args.Plain_optional_bool = &plain_optional_bool_
	}
	plain_optional_const_ := "val"
	args.Plain_optional_const = &plain_optional_const_
	if isZero(args.Plain_optional_number) {
		plain_optional_number_ := 42.0
		args.Plain_optional_number = &plain_optional_number_
	}
	if isZero(args.Plain_optional_string) {
		plain_optional_string_ := "buzzer"
		args.Plain_optional_string = &plain_optional_string_
	}
	if isZero(args.Plain_required_bool) {
		args.Plain_required_bool = true
	}
	args.Plain_required_const = "val"
	if isZero(args.Plain_required_number) {
		args.Plain_required_number = 42.0
	}
	if isZero(args.Plain_required_string) {
		args.Plain_required_string = "buzzer"
	}
	if isZero(args.Required_bool) {
		args.Required_bool = pulumi.Bool(true)
	}
	if isZero(args.Required_enum) {
		args.Required_enum = EnumThing(4)
	}
	if isZero(args.Required_number) {
		args.Required_number = pulumi.Float64(42.0)
	}
	if isZero(args.Required_string) {
		args.Required_string = pulumi.String("buzzer")
	}
	var resource ModuleResource
	err := ctx.RegisterResource("foobar::ModuleResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModuleResource gets an existing ModuleResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModuleResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModuleResourceState, opts ...pulumi.ResourceOption) (*ModuleResource, error) {
	var resource ModuleResource
	err := ctx.ReadResource("foobar::ModuleResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModuleResource resources.
type moduleResourceState struct {
}

type ModuleResourceState struct {
}

func (ModuleResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleResourceState)(nil)).Elem()
}

type moduleResourceArgs struct {
	Optional_bool         *bool      `pulumi:"optional_bool"`
	Optional_const        *string    `pulumi:"optional_const"`
	Optional_enum         *EnumThing `pulumi:"optional_enum"`
	Optional_number       *float64   `pulumi:"optional_number"`
	Optional_string       *string    `pulumi:"optional_string"`
	Plain_optional_bool   *bool      `pulumi:"plain_optional_bool"`
	Plain_optional_const  *string    `pulumi:"plain_optional_const"`
	Plain_optional_number *float64   `pulumi:"plain_optional_number"`
	Plain_optional_string *string    `pulumi:"plain_optional_string"`
	Plain_required_bool   bool       `pulumi:"plain_required_bool"`
	Plain_required_const  string     `pulumi:"plain_required_const"`
	Plain_required_number float64    `pulumi:"plain_required_number"`
	Plain_required_string string     `pulumi:"plain_required_string"`
	Required_bool         bool       `pulumi:"required_bool"`
	Required_enum         EnumThing  `pulumi:"required_enum"`
	Required_number       float64    `pulumi:"required_number"`
	Required_string       string     `pulumi:"required_string"`
}

// The set of arguments for constructing a ModuleResource resource.
type ModuleResourceArgs struct {
	Optional_bool         pulumi.BoolPtrInput
	Optional_const        pulumi.StringPtrInput
	Optional_enum         EnumThingPtrInput
	Optional_number       pulumi.Float64PtrInput
	Optional_string       pulumi.StringPtrInput
	Plain_optional_bool   *bool
	Plain_optional_const  *string
	Plain_optional_number *float64
	Plain_optional_string *string
	Plain_required_bool   bool
	Plain_required_const  string
	Plain_required_number float64
	Plain_required_string string
	Required_bool         pulumi.BoolInput
	Required_enum         EnumThingInput
	Required_number       pulumi.Float64Input
	Required_string       pulumi.StringInput
}

func (ModuleResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleResourceArgs)(nil)).Elem()
}

type ModuleResourceInput interface {
	pulumi.Input

	ToModuleResourceOutput() ModuleResourceOutput
	ToModuleResourceOutputWithContext(ctx context.Context) ModuleResourceOutput
}

func (*ModuleResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleResource)(nil)).Elem()
}

func (i *ModuleResource) ToModuleResourceOutput() ModuleResourceOutput {
	return i.ToModuleResourceOutputWithContext(context.Background())
}

func (i *ModuleResource) ToModuleResourceOutputWithContext(ctx context.Context) ModuleResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleResourceOutput)
}

type ModuleResourceOutput struct{ *pulumi.OutputState }

func (ModuleResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleResource)(nil)).Elem()
}

func (o ModuleResourceOutput) ToModuleResourceOutput() ModuleResourceOutput {
	return o
}

func (o ModuleResourceOutput) ToModuleResourceOutputWithContext(ctx context.Context) ModuleResourceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleResourceInput)(nil)).Elem(), &ModuleResource{})
	pulumi.RegisterOutputType(ModuleResourceOutput{})
}