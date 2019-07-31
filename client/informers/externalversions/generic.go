/*
Copyright 2019 The Kube Vault Authors.

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

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1alpha1 "kubevault.dev/operator/apis/catalog/v1alpha1"
	enginev1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	kubevaultv1alpha1 "kubevault.dev/operator/apis/kubevault/v1alpha1"
	policyv1alpha1 "kubevault.dev/operator/apis/policy/v1alpha1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=catalog.kubevault.com, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("vaultserverversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Catalog().V1alpha1().VaultServerVersions().Informer()}, nil

		// Group=engine.kubevault.com, Version=v1alpha1
	case enginev1alpha1.SchemeGroupVersion.WithResource("awsaccesskeyrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Engine().V1alpha1().AWSAccessKeyRequests().Informer()}, nil
	case enginev1alpha1.SchemeGroupVersion.WithResource("awsroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Engine().V1alpha1().AWSRoles().Informer()}, nil
	case enginev1alpha1.SchemeGroupVersion.WithResource("azureaccesskeyrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Engine().V1alpha1().AzureAccessKeyRequests().Informer()}, nil
	case enginev1alpha1.SchemeGroupVersion.WithResource("azureroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Engine().V1alpha1().AzureRoles().Informer()}, nil
	case enginev1alpha1.SchemeGroupVersion.WithResource("gcpaccesskeyrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Engine().V1alpha1().GCPAccessKeyRequests().Informer()}, nil
	case enginev1alpha1.SchemeGroupVersion.WithResource("gcproles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Engine().V1alpha1().GCPRoles().Informer()}, nil
	case enginev1alpha1.SchemeGroupVersion.WithResource("secretengines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Engine().V1alpha1().SecretEngines().Informer()}, nil

		// Group=kubevault.com, Version=v1alpha1
	case kubevaultv1alpha1.SchemeGroupVersion.WithResource("vaultservers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubevault().V1alpha1().VaultServers().Informer()}, nil

		// Group=policy.kubevault.com, Version=v1alpha1
	case policyv1alpha1.SchemeGroupVersion.WithResource("vaultpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().V1alpha1().VaultPolicies().Informer()}, nil
	case policyv1alpha1.SchemeGroupVersion.WithResource("vaultpolicybindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().V1alpha1().VaultPolicyBindings().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
