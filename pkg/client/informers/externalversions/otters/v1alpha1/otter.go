// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	otters_v1alpha1 "github.com/zachpuck/otter-controller/pkg/apis/otters/v1alpha1"
	versioned "github.com/zachpuck/otter-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/zachpuck/otter-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/zachpuck/otter-controller/pkg/client/listers/otters/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OtterInformer provides access to a shared informer and lister for
// Otters.
type OtterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OtterLister
}

type otterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOtterInformer constructs a new informer for Otter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOtterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOtterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOtterInformer constructs a new informer for Otter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOtterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OttersV1alpha1().Otters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OttersV1alpha1().Otters(namespace).Watch(options)
			},
		},
		&otters_v1alpha1.Otter{},
		resyncPeriod,
		indexers,
	)
}

func (f *otterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOtterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *otterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&otters_v1alpha1.Otter{}, f.defaultInformer)
}

func (f *otterInformer) Lister() v1alpha1.OtterLister {
	return v1alpha1.NewOtterLister(f.Informer().GetIndexer())
}
