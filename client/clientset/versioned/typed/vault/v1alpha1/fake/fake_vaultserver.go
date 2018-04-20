/*
Copyright 2018 The Vault Operator Authors.

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

package fake

import (
	v1alpha1 "github.com/soter/vault-operator/apis/vault/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVaultServers implements VaultServerInterface
type FakeVaultServers struct {
	Fake *FakeVaultV1alpha1
	ns   string
}

var vaultserversResource = schema.GroupVersionResource{Group: "vault.soter.ac", Version: "v1alpha1", Resource: "vaultservers"}

var vaultserversKind = schema.GroupVersionKind{Group: "vault.soter.ac", Version: "v1alpha1", Kind: "VaultServer"}

// Get takes name of the vaultServer, and returns the corresponding vaultServer object, and an error if there is any.
func (c *FakeVaultServers) Get(name string, options v1.GetOptions) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vaultserversResource, c.ns, name), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}

// List takes label and field selectors, and returns the list of VaultServers that match those selectors.
func (c *FakeVaultServers) List(opts v1.ListOptions) (result *v1alpha1.VaultServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vaultserversResource, vaultserversKind, c.ns, opts), &v1alpha1.VaultServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VaultServerList{}
	for _, item := range obj.(*v1alpha1.VaultServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vaultServers.
func (c *FakeVaultServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vaultserversResource, c.ns, opts))

}

// Create takes the representation of a vaultServer and creates it.  Returns the server's representation of the vaultServer, and an error, if there is any.
func (c *FakeVaultServers) Create(vaultServer *v1alpha1.VaultServer) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vaultserversResource, c.ns, vaultServer), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}

// Update takes the representation of a vaultServer and updates it. Returns the server's representation of the vaultServer, and an error, if there is any.
func (c *FakeVaultServers) Update(vaultServer *v1alpha1.VaultServer) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vaultserversResource, c.ns, vaultServer), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}

// Delete takes name of the vaultServer and deletes it. Returns an error if one occurs.
func (c *FakeVaultServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vaultserversResource, c.ns, name), &v1alpha1.VaultServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVaultServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vaultserversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VaultServerList{})
	return err
}

// Patch applies the patch and returns the patched vaultServer.
func (c *FakeVaultServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vaultserversResource, c.ns, name, data, subresources...), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}
