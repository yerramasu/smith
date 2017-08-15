// +build !ignore_autogenerated

// Generated file, do not modify manually!

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Bundle, InType: reflect.TypeOf(&Bundle{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BundleCondition, InType: reflect.TypeOf(&BundleCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BundleList, InType: reflect.TypeOf(&BundleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BundleSpec, InType: reflect.TypeOf(&BundleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_BundleStatus, InType: reflect.TypeOf(&BundleStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Resource, InType: reflect.TypeOf(&Resource{})},
	)
}

// DeepCopy_v1_Bundle is an autogenerated deepcopy function.
func DeepCopy_v1_Bundle(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Bundle)
		out := out.(*Bundle)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*meta_v1.ObjectMeta)
		}
		if err := DeepCopy_v1_BundleSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_BundleStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_v1_BundleCondition is an autogenerated deepcopy function.
func DeepCopy_v1_BundleCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BundleCondition)
		out := out.(*BundleCondition)
		*out = *in
		out.LastUpdateTime = in.LastUpdateTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_v1_BundleList is an autogenerated deepcopy function.
func DeepCopy_v1_BundleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BundleList)
		out := out.(*BundleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Bundle, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Bundle(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_BundleSpec is an autogenerated deepcopy function.
func DeepCopy_v1_BundleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BundleSpec)
		out := out.(*BundleSpec)
		*out = *in
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make([]Resource, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Resource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_BundleStatus is an autogenerated deepcopy function.
func DeepCopy_v1_BundleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BundleStatus)
		out := out.(*BundleStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]BundleCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_BundleCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_v1_Resource is an autogenerated deepcopy function.
func DeepCopy_v1_Resource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Resource)
		out := out.(*Resource)
		*out = *in
		if in.DependsOn != nil {
			in, out := &in.DependsOn, &out.DependsOn
			*out = make([]ResourceName, len(*in))
			copy(*out, *in)
		}
		// in.Spec is kind 'Interface'
		if in.Spec != nil {
			if newVal, err := c.DeepCopy(&in.Spec); err != nil {
				return err
			} else {
				out.Spec = *newVal.(*runtime.Object)
			}
		}
		return nil
	}
}