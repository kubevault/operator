/*
Copyright The KubeVault Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	scheme "kubevault.dev/operator/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSRolesGetter has a method to return a AWSRoleInterface.
// A group's client should implement this interface.
type AWSRolesGetter interface {
	AWSRoles(namespace string) AWSRoleInterface
}

// AWSRoleInterface has methods to work with AWSRole resources.
type AWSRoleInterface interface {
	Create(ctx context.Context, aWSRole *v1alpha1.AWSRole, opts v1.CreateOptions) (*v1alpha1.AWSRole, error)
	Update(ctx context.Context, aWSRole *v1alpha1.AWSRole, opts v1.UpdateOptions) (*v1alpha1.AWSRole, error)
	UpdateStatus(ctx context.Context, aWSRole *v1alpha1.AWSRole, opts v1.UpdateOptions) (*v1alpha1.AWSRole, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AWSRole, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AWSRoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSRole, err error)
	AWSRoleExpansion
}

// aWSRoles implements AWSRoleInterface
type aWSRoles struct {
	client rest.Interface
	ns     string
}

// newAWSRoles returns a AWSRoles
func newAWSRoles(c *EngineV1alpha1Client, namespace string) *aWSRoles {
	return &aWSRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSRole, and returns the corresponding aWSRole object, and an error if there is any.
func (c *aWSRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSRole, err error) {
	result = &v1alpha1.AWSRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSRoles that match those selectors.
func (c *aWSRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AWSRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSRoles.
func (c *aWSRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aWSRole and creates it.  Returns the server's representation of the aWSRole, and an error, if there is any.
func (c *aWSRoles) Create(ctx context.Context, aWSRole *v1alpha1.AWSRole, opts v1.CreateOptions) (result *v1alpha1.AWSRole, err error) {
	result = &v1alpha1.AWSRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSRole).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aWSRole and updates it. Returns the server's representation of the aWSRole, and an error, if there is any.
func (c *aWSRoles) Update(ctx context.Context, aWSRole *v1alpha1.AWSRole, opts v1.UpdateOptions) (result *v1alpha1.AWSRole, err error) {
	result = &v1alpha1.AWSRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsroles").
		Name(aWSRole.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSRole).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aWSRoles) UpdateStatus(ctx context.Context, aWSRole *v1alpha1.AWSRole, opts v1.UpdateOptions) (result *v1alpha1.AWSRole, err error) {
	result = &v1alpha1.AWSRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsroles").
		Name(aWSRole.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSRole).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aWSRole and deletes it. Returns an error if one occurs.
func (c *aWSRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsroles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsroles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aWSRole.
func (c *aWSRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSRole, err error) {
	result = &v1alpha1.AWSRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsroles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
