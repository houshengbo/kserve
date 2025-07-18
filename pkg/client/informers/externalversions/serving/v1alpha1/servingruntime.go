/*
Copyright 2023 The KServe Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apisservingv1alpha1 "github.com/kserve/kserve/pkg/apis/serving/v1alpha1"
	versioned "github.com/kserve/kserve/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kserve/kserve/pkg/client/informers/externalversions/internalinterfaces"
	servingv1alpha1 "github.com/kserve/kserve/pkg/client/listers/serving/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServingRuntimeInformer provides access to a shared informer and lister for
// ServingRuntimes.
type ServingRuntimeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() servingv1alpha1.ServingRuntimeLister
}

type servingRuntimeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServingRuntimeInformer constructs a new informer for ServingRuntime type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServingRuntimeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServingRuntimeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServingRuntimeInformer constructs a new informer for ServingRuntime type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServingRuntimeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().ServingRuntimes(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().ServingRuntimes(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().ServingRuntimes(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha1().ServingRuntimes(namespace).Watch(ctx, options)
			},
		},
		&apisservingv1alpha1.ServingRuntime{},
		resyncPeriod,
		indexers,
	)
}

func (f *servingRuntimeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServingRuntimeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *servingRuntimeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisservingv1alpha1.ServingRuntime{}, f.defaultInformer)
}

func (f *servingRuntimeInformer) Lister() servingv1alpha1.ServingRuntimeLister {
	return servingv1alpha1.NewServingRuntimeLister(f.Informer().GetIndexer())
}
