// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedGroovyScript) DeepCopyInto(out *AppliedGroovyScript) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedGroovyScript.
func (in *AppliedGroovyScript) DeepCopy() *AppliedGroovyScript {
	if in == nil {
		return nil
	}
	out := new(AppliedGroovyScript)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Casc) DeepCopyInto(out *Casc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Casc.
func (in *Casc) DeepCopy() *Casc {
	if in == nil {
		return nil
	}
	out := new(Casc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Casc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CascList) DeepCopyInto(out *CascList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Casc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CascList.
func (in *CascList) DeepCopy() *CascList {
	if in == nil {
		return nil
	}
	out := new(CascList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CascList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CascSpec) DeepCopyInto(out *CascSpec) {
	*out = *in
	in.GroovyScripts.DeepCopyInto(&out.GroovyScripts)
	in.ConfigurationAsCode.DeepCopyInto(&out.ConfigurationAsCode)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CascSpec.
func (in *CascSpec) DeepCopy() *CascSpec {
	if in == nil {
		return nil
	}
	out := new(CascSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CascStatus) DeepCopyInto(out *CascStatus) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.AppliedGroovyScripts != nil {
		in, out := &in.AppliedGroovyScripts, &out.AppliedGroovyScripts
		*out = make([]AppliedGroovyScript, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CascStatus.
func (in *CascStatus) DeepCopy() *CascStatus {
	if in == nil {
		return nil
	}
	out := new(CascStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapRef) DeepCopyInto(out *ConfigMapRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapRef.
func (in *ConfigMapRef) DeepCopy() *ConfigMapRef {
	if in == nil {
		return nil
	}
	out := new(ConfigMapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Customization) DeepCopyInto(out *Customization) {
	*out = *in
	out.Secret = in.Secret
	if in.Configurations != nil {
		in, out := &in.Configurations, &out.Configurations
		*out = make([]ConfigMapRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Customization.
func (in *Customization) DeepCopy() *Customization {
	if in == nil {
		return nil
	}
	out := new(Customization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}