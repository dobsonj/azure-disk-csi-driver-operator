/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ManagedFieldsEntryApplyConfiguration represents an declarative configuration of the ManagedFieldsEntry type for use
// with apply.
type ManagedFieldsEntryApplyConfiguration struct {
	Manager    *string                        `json:"manager,omitempty"`
	Operation  *v1.ManagedFieldsOperationType `json:"operation,omitempty"`
	APIVersion *string                        `json:"apiVersion,omitempty"`
	Time       *v1.Time                       `json:"time,omitempty"`
	FieldsType *string                        `json:"fieldsType,omitempty"`
	FieldsV1   *v1.FieldsV1                   `json:"fieldsV1,omitempty"`
}

// ManagedFieldsEntryApplyConfiguration constructs an declarative configuration of the ManagedFieldsEntry type for use with
// apply.
func ManagedFieldsEntry() *ManagedFieldsEntryApplyConfiguration {
	return &ManagedFieldsEntryApplyConfiguration{}
}

// WithManager sets the Manager field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Manager field is set to the value of the last call.
func (b *ManagedFieldsEntryApplyConfiguration) WithManager(value string) *ManagedFieldsEntryApplyConfiguration {
	b.Manager = &value
	return b
}

// WithOperation sets the Operation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Operation field is set to the value of the last call.
func (b *ManagedFieldsEntryApplyConfiguration) WithOperation(value v1.ManagedFieldsOperationType) *ManagedFieldsEntryApplyConfiguration {
	b.Operation = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ManagedFieldsEntryApplyConfiguration) WithAPIVersion(value string) *ManagedFieldsEntryApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithTime sets the Time field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Time field is set to the value of the last call.
func (b *ManagedFieldsEntryApplyConfiguration) WithTime(value v1.Time) *ManagedFieldsEntryApplyConfiguration {
	b.Time = &value
	return b
}

// WithFieldsType sets the FieldsType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FieldsType field is set to the value of the last call.
func (b *ManagedFieldsEntryApplyConfiguration) WithFieldsType(value string) *ManagedFieldsEntryApplyConfiguration {
	b.FieldsType = &value
	return b
}

// WithFieldsV1 sets the FieldsV1 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FieldsV1 field is set to the value of the last call.
func (b *ManagedFieldsEntryApplyConfiguration) WithFieldsV1(value v1.FieldsV1) *ManagedFieldsEntryApplyConfiguration {
	b.FieldsV1 = &value
	return b
}
