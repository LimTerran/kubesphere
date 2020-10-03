/*
Copyright 2020 The KubeSphere Authors.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
)

// FederatedResourceQuotaLister helps list FederatedResourceQuotas.
type FederatedResourceQuotaLister interface {
	// List lists all FederatedResourceQuotas in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.FederatedResourceQuota, err error)
	// FederatedResourceQuotas returns an object that can list and get FederatedResourceQuotas.
	FederatedResourceQuotas(namespace string) FederatedResourceQuotaNamespaceLister
	FederatedResourceQuotaListerExpansion
}

// federatedResourceQuotaLister implements the FederatedResourceQuotaLister interface.
type federatedResourceQuotaLister struct {
	indexer cache.Indexer
}

// NewFederatedResourceQuotaLister returns a new FederatedResourceQuotaLister.
func NewFederatedResourceQuotaLister(indexer cache.Indexer) FederatedResourceQuotaLister {
	return &federatedResourceQuotaLister{indexer: indexer}
}

// List lists all FederatedResourceQuotas in the indexer.
func (s *federatedResourceQuotaLister) List(selector labels.Selector) (ret []*v1beta1.FederatedResourceQuota, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedResourceQuota))
	})
	return ret, err
}

// FederatedResourceQuotas returns an object that can list and get FederatedResourceQuotas.
func (s *federatedResourceQuotaLister) FederatedResourceQuotas(namespace string) FederatedResourceQuotaNamespaceLister {
	return federatedResourceQuotaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedResourceQuotaNamespaceLister helps list and get FederatedResourceQuotas.
type FederatedResourceQuotaNamespaceLister interface {
	// List lists all FederatedResourceQuotas in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.FederatedResourceQuota, err error)
	// Get retrieves the FederatedResourceQuota from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.FederatedResourceQuota, error)
	FederatedResourceQuotaNamespaceListerExpansion
}

// federatedResourceQuotaNamespaceLister implements the FederatedResourceQuotaNamespaceLister
// interface.
type federatedResourceQuotaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedResourceQuotas in the indexer for a given namespace.
func (s federatedResourceQuotaNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedResourceQuota, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedResourceQuota))
	})
	return ret, err
}

// Get retrieves the FederatedResourceQuota from the indexer for a given namespace and name.
func (s federatedResourceQuotaNamespaceLister) Get(name string) (*v1beta1.FederatedResourceQuota, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedresourcequota"), name)
	}
	return obj.(*v1beta1.FederatedResourceQuota), nil
}
