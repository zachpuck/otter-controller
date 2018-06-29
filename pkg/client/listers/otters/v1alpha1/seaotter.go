// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zachpuck/otter-controller/pkg/apis/otters/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SeaOtterLister helps list SeaOtters.
type SeaOtterLister interface {
	// List lists all SeaOtters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SeaOtter, err error)
	// SeaOtters returns an object that can list and get SeaOtters.
	SeaOtters(namespace string) SeaOtterNamespaceLister
	SeaOtterListerExpansion
}

// seaOtterLister implements the SeaOtterLister interface.
type seaOtterLister struct {
	indexer cache.Indexer
}

// NewSeaOtterLister returns a new SeaOtterLister.
func NewSeaOtterLister(indexer cache.Indexer) SeaOtterLister {
	return &seaOtterLister{indexer: indexer}
}

// List lists all SeaOtters in the indexer.
func (s *seaOtterLister) List(selector labels.Selector) (ret []*v1alpha1.SeaOtter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SeaOtter))
	})
	return ret, err
}

// SeaOtters returns an object that can list and get SeaOtters.
func (s *seaOtterLister) SeaOtters(namespace string) SeaOtterNamespaceLister {
	return seaOtterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SeaOtterNamespaceLister helps list and get SeaOtters.
type SeaOtterNamespaceLister interface {
	// List lists all SeaOtters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SeaOtter, err error)
	// Get retrieves the SeaOtter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SeaOtter, error)
	SeaOtterNamespaceListerExpansion
}

// seaOtterNamespaceLister implements the SeaOtterNamespaceLister
// interface.
type seaOtterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SeaOtters in the indexer for a given namespace.
func (s seaOtterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SeaOtter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SeaOtter))
	})
	return ret, err
}

// Get retrieves the SeaOtter from the indexer for a given namespace and name.
func (s seaOtterNamespaceLister) Get(name string) (*v1alpha1.SeaOtter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("seaotter"), name)
	}
	return obj.(*v1alpha1.SeaOtter), nil
}