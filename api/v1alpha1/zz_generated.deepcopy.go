// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSecurityConfig) DeepCopyInto(out *SmbSecurityConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSecurityConfig.
func (in *SmbSecurityConfig) DeepCopy() *SmbSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(SmbSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmbSecurityConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSecurityConfigList) DeepCopyInto(out *SmbSecurityConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SmbSecurityConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSecurityConfigList.
func (in *SmbSecurityConfigList) DeepCopy() *SmbSecurityConfigList {
	if in == nil {
		return nil
	}
	out := new(SmbSecurityConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmbSecurityConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSecurityConfigSpec) DeepCopyInto(out *SmbSecurityConfigSpec) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = new(SmbSecurityUsersSpec)
		**out = **in
	}
	if in.JoinSources != nil {
		in, out := &in.JoinSources, &out.JoinSources
		*out = make([]SmbSecurityJoinSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSecurityConfigSpec.
func (in *SmbSecurityConfigSpec) DeepCopy() *SmbSecurityConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SmbSecurityConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSecurityConfigStatus) DeepCopyInto(out *SmbSecurityConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSecurityConfigStatus.
func (in *SmbSecurityConfigStatus) DeepCopy() *SmbSecurityConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SmbSecurityConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSecurityJoinSpec) DeepCopyInto(out *SmbSecurityJoinSpec) {
	*out = *in
	if in.UserJoin != nil {
		in, out := &in.UserJoin, &out.UserJoin
		*out = new(SmbSecurityUserJoinSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSecurityJoinSpec.
func (in *SmbSecurityJoinSpec) DeepCopy() *SmbSecurityJoinSpec {
	if in == nil {
		return nil
	}
	out := new(SmbSecurityJoinSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSecurityUserJoinSpec) DeepCopyInto(out *SmbSecurityUserJoinSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSecurityUserJoinSpec.
func (in *SmbSecurityUserJoinSpec) DeepCopy() *SmbSecurityUserJoinSpec {
	if in == nil {
		return nil
	}
	out := new(SmbSecurityUserJoinSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSecurityUsersSpec) DeepCopyInto(out *SmbSecurityUsersSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSecurityUsersSpec.
func (in *SmbSecurityUsersSpec) DeepCopy() *SmbSecurityUsersSpec {
	if in == nil {
		return nil
	}
	out := new(SmbSecurityUsersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbShare) DeepCopyInto(out *SmbShare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbShare.
func (in *SmbShare) DeepCopy() *SmbShare {
	if in == nil {
		return nil
	}
	out := new(SmbShare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmbShare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbShareList) DeepCopyInto(out *SmbShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SmbShare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbShareList.
func (in *SmbShareList) DeepCopy() *SmbShareList {
	if in == nil {
		return nil
	}
	out := new(SmbShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmbShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbSharePvcSpec) DeepCopyInto(out *SmbSharePvcSpec) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbSharePvcSpec.
func (in *SmbSharePvcSpec) DeepCopy() *SmbSharePvcSpec {
	if in == nil {
		return nil
	}
	out := new(SmbSharePvcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbShareSpec) DeepCopyInto(out *SmbShareSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbShareSpec.
func (in *SmbShareSpec) DeepCopy() *SmbShareSpec {
	if in == nil {
		return nil
	}
	out := new(SmbShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbShareStatus) DeepCopyInto(out *SmbShareStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbShareStatus.
func (in *SmbShareStatus) DeepCopy() *SmbShareStatus {
	if in == nil {
		return nil
	}
	out := new(SmbShareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmbShareStorageSpec) DeepCopyInto(out *SmbShareStorageSpec) {
	*out = *in
	if in.Pvc != nil {
		in, out := &in.Pvc, &out.Pvc
		*out = new(SmbSharePvcSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmbShareStorageSpec.
func (in *SmbShareStorageSpec) DeepCopy() *SmbShareStorageSpec {
	if in == nil {
		return nil
	}
	out := new(SmbShareStorageSpec)
	in.DeepCopyInto(out)
	return out
}
