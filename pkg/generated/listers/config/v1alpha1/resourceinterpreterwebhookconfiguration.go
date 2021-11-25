// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceInterpreterWebhookConfigurationLister helps list ResourceInterpreterWebhookConfigurations.
// All objects returned here must be treated as read-only.
type ResourceInterpreterWebhookConfigurationLister interface {
	// List lists all ResourceInterpreterWebhookConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceInterpreterWebhookConfiguration, err error)
	// Get retrieves the ResourceInterpreterWebhookConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceInterpreterWebhookConfiguration, error)
	ResourceInterpreterWebhookConfigurationListerExpansion
}

// resourceInterpreterWebhookConfigurationLister implements the ResourceInterpreterWebhookConfigurationLister interface.
type resourceInterpreterWebhookConfigurationLister struct {
	indexer cache.Indexer
}

// NewResourceInterpreterWebhookConfigurationLister returns a new ResourceInterpreterWebhookConfigurationLister.
func NewResourceInterpreterWebhookConfigurationLister(indexer cache.Indexer) ResourceInterpreterWebhookConfigurationLister {
	return &resourceInterpreterWebhookConfigurationLister{indexer: indexer}
}

// List lists all ResourceInterpreterWebhookConfigurations in the indexer.
func (s *resourceInterpreterWebhookConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceInterpreterWebhookConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceInterpreterWebhookConfiguration))
	})
	return ret, err
}

// Get retrieves the ResourceInterpreterWebhookConfiguration from the index for a given name.
func (s *resourceInterpreterWebhookConfigurationLister) Get(name string) (*v1alpha1.ResourceInterpreterWebhookConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourceinterpreterwebhookconfiguration"), name)
	}
	return obj.(*v1alpha1.ResourceInterpreterWebhookConfiguration), nil
}