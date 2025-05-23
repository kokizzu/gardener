// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// RollingUpdateStrategyApplyConfiguration represents an declarative configuration of the RollingUpdateStrategy type for use
// with apply.
type RollingUpdateStrategyApplyConfiguration struct {
	Partition *int32 `json:"partition,omitempty"`
}

// RollingUpdateStrategyApplyConfiguration constructs an declarative configuration of the RollingUpdateStrategy type for use with
// apply.
func RollingUpdateStrategy() *RollingUpdateStrategyApplyConfiguration {
	return &RollingUpdateStrategyApplyConfiguration{}
}

// WithPartition sets the Partition field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Partition field is set to the value of the last call.
func (b *RollingUpdateStrategyApplyConfiguration) WithPartition(value int32) *RollingUpdateStrategyApplyConfiguration {
	b.Partition = &value
	return b
}
