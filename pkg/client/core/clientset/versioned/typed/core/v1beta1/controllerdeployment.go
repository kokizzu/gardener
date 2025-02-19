// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	corev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ControllerDeploymentsGetter has a method to return a ControllerDeploymentInterface.
// A group's client should implement this interface.
type ControllerDeploymentsGetter interface {
	ControllerDeployments() ControllerDeploymentInterface
}

// ControllerDeploymentInterface has methods to work with ControllerDeployment resources.
type ControllerDeploymentInterface interface {
	Create(ctx context.Context, controllerDeployment *corev1beta1.ControllerDeployment, opts v1.CreateOptions) (*corev1beta1.ControllerDeployment, error)
	Update(ctx context.Context, controllerDeployment *corev1beta1.ControllerDeployment, opts v1.UpdateOptions) (*corev1beta1.ControllerDeployment, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*corev1beta1.ControllerDeployment, error)
	List(ctx context.Context, opts v1.ListOptions) (*corev1beta1.ControllerDeploymentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1beta1.ControllerDeployment, err error)
	ControllerDeploymentExpansion
}

// controllerDeployments implements ControllerDeploymentInterface
type controllerDeployments struct {
	*gentype.ClientWithList[*corev1beta1.ControllerDeployment, *corev1beta1.ControllerDeploymentList]
}

// newControllerDeployments returns a ControllerDeployments
func newControllerDeployments(c *CoreV1beta1Client) *controllerDeployments {
	return &controllerDeployments{
		gentype.NewClientWithList[*corev1beta1.ControllerDeployment, *corev1beta1.ControllerDeploymentList](
			"controllerdeployments",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *corev1beta1.ControllerDeployment { return &corev1beta1.ControllerDeployment{} },
			func() *corev1beta1.ControllerDeploymentList { return &corev1beta1.ControllerDeploymentList{} },
		),
	}
}
