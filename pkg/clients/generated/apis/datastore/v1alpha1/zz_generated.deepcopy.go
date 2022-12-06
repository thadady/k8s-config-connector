//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreIndex) DeepCopyInto(out *DatastoreIndex) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreIndex.
func (in *DatastoreIndex) DeepCopy() *DatastoreIndex {
	if in == nil {
		return nil
	}
	out := new(DatastoreIndex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreIndex) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreIndexList) DeepCopyInto(out *DatastoreIndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatastoreIndex, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreIndexList.
func (in *DatastoreIndexList) DeepCopy() *DatastoreIndexList {
	if in == nil {
		return nil
	}
	out := new(DatastoreIndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatastoreIndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreIndexSpec) DeepCopyInto(out *DatastoreIndexSpec) {
	*out = *in
	if in.Ancestor != nil {
		in, out := &in.Ancestor, &out.Ancestor
		*out = new(string)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]IndexProperties, len(*in))
		copy(*out, *in)
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreIndexSpec.
func (in *DatastoreIndexSpec) DeepCopy() *DatastoreIndexSpec {
	if in == nil {
		return nil
	}
	out := new(DatastoreIndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatastoreIndexStatus) DeepCopyInto(out *DatastoreIndexStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.IndexId != nil {
		in, out := &in.IndexId, &out.IndexId
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatastoreIndexStatus.
func (in *DatastoreIndexStatus) DeepCopy() *DatastoreIndexStatus {
	if in == nil {
		return nil
	}
	out := new(DatastoreIndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexProperties) DeepCopyInto(out *IndexProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexProperties.
func (in *IndexProperties) DeepCopy() *IndexProperties {
	if in == nil {
		return nil
	}
	out := new(IndexProperties)
	in.DeepCopyInto(out)
	return out
}