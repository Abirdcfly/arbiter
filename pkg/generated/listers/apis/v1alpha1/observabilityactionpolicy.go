/*
Copyright 2022 The Arbiter Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kube-arbiter/arbiter/pkg/apis/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ObservabilityActionPolicyLister helps list ObservabilityActionPolicies.
// All objects returned here must be treated as read-only.
type ObservabilityActionPolicyLister interface {
	// List lists all ObservabilityActionPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ObservabilityActionPolicy, err error)
	// ObservabilityActionPolicies returns an object that can list and get ObservabilityActionPolicies.
	ObservabilityActionPolicies(namespace string) ObservabilityActionPolicyNamespaceLister
	ObservabilityActionPolicyListerExpansion
}

// observabilityActionPolicyLister implements the ObservabilityActionPolicyLister interface.
type observabilityActionPolicyLister struct {
	indexer cache.Indexer
}

// NewObservabilityActionPolicyLister returns a new ObservabilityActionPolicyLister.
func NewObservabilityActionPolicyLister(indexer cache.Indexer) ObservabilityActionPolicyLister {
	return &observabilityActionPolicyLister{indexer: indexer}
}

// List lists all ObservabilityActionPolicies in the indexer.
func (s *observabilityActionPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ObservabilityActionPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ObservabilityActionPolicy))
	})
	return ret, err
}

// ObservabilityActionPolicies returns an object that can list and get ObservabilityActionPolicies.
func (s *observabilityActionPolicyLister) ObservabilityActionPolicies(namespace string) ObservabilityActionPolicyNamespaceLister {
	return observabilityActionPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ObservabilityActionPolicyNamespaceLister helps list and get ObservabilityActionPolicies.
// All objects returned here must be treated as read-only.
type ObservabilityActionPolicyNamespaceLister interface {
	// List lists all ObservabilityActionPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ObservabilityActionPolicy, err error)
	// Get retrieves the ObservabilityActionPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ObservabilityActionPolicy, error)
	ObservabilityActionPolicyNamespaceListerExpansion
}

// observabilityActionPolicyNamespaceLister implements the ObservabilityActionPolicyNamespaceLister
// interface.
type observabilityActionPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ObservabilityActionPolicies in the indexer for a given namespace.
func (s observabilityActionPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ObservabilityActionPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ObservabilityActionPolicy))
	})
	return ret, err
}

// Get retrieves the ObservabilityActionPolicy from the indexer for a given namespace and name.
func (s observabilityActionPolicyNamespaceLister) Get(name string) (*v1alpha1.ObservabilityActionPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("observabilityactionpolicy"), name)
	}
	return obj.(*v1alpha1.ObservabilityActionPolicy), nil
}