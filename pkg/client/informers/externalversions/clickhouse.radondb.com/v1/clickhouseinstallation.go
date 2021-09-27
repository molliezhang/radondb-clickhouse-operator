// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	clickhouseradondbcomv1 "github.com/TCeason/clickhouse-operator/pkg/apis/clickhouse.radondb.com/v1"
	versioned "github.com/TCeason/clickhouse-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/TCeason/clickhouse-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/TCeason/clickhouse-operator/pkg/client/listers/clickhouse.radondb.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClickHouseInstallationInformer provides access to a shared informer and lister for
// ClickHouseInstallations.
type ClickHouseInstallationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClickHouseInstallationLister
}

type clickHouseInstallationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClickHouseInstallationInformer constructs a new informer for ClickHouseInstallation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClickHouseInstallationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClickHouseInstallationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClickHouseInstallationInformer constructs a new informer for ClickHouseInstallation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClickHouseInstallationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClickhouseV1().ClickHouseInstallations(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClickhouseV1().ClickHouseInstallations(namespace).Watch(options)
			},
		},
		&clickhouseradondbcomv1.ClickHouseInstallation{},
		resyncPeriod,
		indexers,
	)
}

func (f *clickHouseInstallationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClickHouseInstallationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clickHouseInstallationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clickhouseradondbcomv1.ClickHouseInstallation{}, f.defaultInformer)
}

func (f *clickHouseInstallationInformer) Lister() v1.ClickHouseInstallationLister {
	return v1.NewClickHouseInstallationLister(f.Informer().GetIndexer())
}