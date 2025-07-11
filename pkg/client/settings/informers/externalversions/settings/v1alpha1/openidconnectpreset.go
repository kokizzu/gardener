// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apissettingsv1alpha1 "github.com/gardener/gardener/pkg/apis/settings/v1alpha1"
	versioned "github.com/gardener/gardener/pkg/client/settings/clientset/versioned"
	internalinterfaces "github.com/gardener/gardener/pkg/client/settings/informers/externalversions/internalinterfaces"
	settingsv1alpha1 "github.com/gardener/gardener/pkg/client/settings/listers/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OpenIDConnectPresetInformer provides access to a shared informer and lister for
// OpenIDConnectPresets.
type OpenIDConnectPresetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() settingsv1alpha1.OpenIDConnectPresetLister
}

type openIDConnectPresetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOpenIDConnectPresetInformer constructs a new informer for OpenIDConnectPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOpenIDConnectPresetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOpenIDConnectPresetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOpenIDConnectPresetInformer constructs a new informer for OpenIDConnectPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOpenIDConnectPresetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().OpenIDConnectPresets(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().OpenIDConnectPresets(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().OpenIDConnectPresets(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().OpenIDConnectPresets(namespace).Watch(ctx, options)
			},
		},
		&apissettingsv1alpha1.OpenIDConnectPreset{},
		resyncPeriod,
		indexers,
	)
}

func (f *openIDConnectPresetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOpenIDConnectPresetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *openIDConnectPresetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apissettingsv1alpha1.OpenIDConnectPreset{}, f.defaultInformer)
}

func (f *openIDConnectPresetInformer) Lister() settingsv1alpha1.OpenIDConnectPresetLister {
	return settingsv1alpha1.NewOpenIDConnectPresetLister(f.Informer().GetIndexer())
}
