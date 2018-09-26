/*
Copyright 2018 The Kube Vault Authors.

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
	v1alpha1 "github.com/kubevault/operator/apis/kubevault/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VaultPolicyLister helps list VaultPolicies.
type VaultPolicyLister interface {
	// List lists all VaultPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VaultPolicy, err error)
	// VaultPolicies returns an object that can list and get VaultPolicies.
	VaultPolicies(namespace string) VaultPolicyNamespaceLister
	VaultPolicyListerExpansion
}

// vaultPolicyLister implements the VaultPolicyLister interface.
type vaultPolicyLister struct {
	indexer cache.Indexer
}

// NewVaultPolicyLister returns a new VaultPolicyLister.
func NewVaultPolicyLister(indexer cache.Indexer) VaultPolicyLister {
	return &vaultPolicyLister{indexer: indexer}
}

// List lists all VaultPolicies in the indexer.
func (s *vaultPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.VaultPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultPolicy))
	})
	return ret, err
}

// VaultPolicies returns an object that can list and get VaultPolicies.
func (s *vaultPolicyLister) VaultPolicies(namespace string) VaultPolicyNamespaceLister {
	return vaultPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VaultPolicyNamespaceLister helps list and get VaultPolicies.
type VaultPolicyNamespaceLister interface {
	// List lists all VaultPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VaultPolicy, err error)
	// Get retrieves the VaultPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VaultPolicy, error)
	VaultPolicyNamespaceListerExpansion
}

// vaultPolicyNamespaceLister implements the VaultPolicyNamespaceLister
// interface.
type vaultPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VaultPolicies in the indexer for a given namespace.
func (s vaultPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VaultPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultPolicy))
	})
	return ret, err
}

// Get retrieves the VaultPolicy from the indexer for a given namespace and name.
func (s vaultPolicyNamespaceLister) Get(name string) (*v1alpha1.VaultPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vaultpolicy"), name)
	}
	return obj.(*v1alpha1.VaultPolicy), nil
}
