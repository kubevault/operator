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
	v1alpha1 "github.com/soter/vault-operator/apis/extensions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	testing "k8s.io/client-go/testing"
)

// FakeSecrets implements SecretInterface
type FakeSecrets struct {
	Fake *FakeExtensionsV1alpha1
	ns   string
}

var secretsResource = schema.GroupVersionResource{Group: "extensions.vault.soter.ac", Version: "v1alpha1", Resource: "secrets"}

var secretsKind = schema.GroupVersionKind{Group: "extensions.vault.soter.ac", Version: "v1alpha1", Kind: "Secret"}

// Get takes name of the secret, and returns the corresponding secret object, and an error if there is any.
func (c *FakeSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretsResource, c.ns, name), &v1alpha1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Secret), err
}

// List takes label and field selectors, and returns the list of Secrets that match those selectors.
func (c *FakeSecrets) List(opts v1.ListOptions) (result *v1alpha1.SecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretsResource, secretsKind, c.ns, opts), &v1alpha1.SecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecretList{}
	for _, item := range obj.(*v1alpha1.SecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create takes the representation of a secret and creates it.  Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Create(secret *v1alpha1.Secret) (result *v1alpha1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretsResource, c.ns, secret), &v1alpha1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Secret), err
}

// Update takes the representation of a secret and updates it. Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Update(secret *v1alpha1.Secret) (result *v1alpha1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretsResource, c.ns, secret), &v1alpha1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Secret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecrets) UpdateStatus(secret *v1alpha1.Secret) (*v1alpha1.Secret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretsResource, "status", c.ns, secret), &v1alpha1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Secret), err
}

// Delete takes name of the secret and deletes it. Returns an error if one occurs.
func (c *FakeSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretsResource, c.ns, name), &v1alpha1.Secret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecretList{})
	return err
}

// Patch applies the patch and returns the patched secret.
func (c *FakeSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretsResource, c.ns, name, data, subresources...), &v1alpha1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Secret), err
}
