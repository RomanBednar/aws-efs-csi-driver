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

// ListMetaApplyConfiguration represents an declarative configuration of the ListMeta type for use
// with apply.
type ListMetaApplyConfiguration struct {
	SelfLink           *string `json:"selfLink,omitempty"`
	ResourceVersion    *string `json:"resourceVersion,omitempty"`
	Continue           *string `json:"continue,omitempty"`
	RemainingItemCount *int64  `json:"remainingItemCount,omitempty"`
}

// ListMetaApplyConfiguration constructs an declarative configuration of the ListMeta type for use with
// apply.
func ListMeta() *ListMetaApplyConfiguration {
	return &ListMetaApplyConfiguration{}
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *ListMetaApplyConfiguration) WithSelfLink(value string) *ListMetaApplyConfiguration {
	b.SelfLink = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ListMetaApplyConfiguration) WithResourceVersion(value string) *ListMetaApplyConfiguration {
	b.ResourceVersion = &value
	return b
}

// WithContinue sets the Continue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Continue field is set to the value of the last call.
func (b *ListMetaApplyConfiguration) WithContinue(value string) *ListMetaApplyConfiguration {
	b.Continue = &value
	return b
}

// WithRemainingItemCount sets the RemainingItemCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemainingItemCount field is set to the value of the last call.
func (b *ListMetaApplyConfiguration) WithRemainingItemCount(value int64) *ListMetaApplyConfiguration {
	b.RemainingItemCount = &value
	return b
}
