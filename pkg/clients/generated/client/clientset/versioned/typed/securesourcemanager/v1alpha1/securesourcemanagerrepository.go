// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/securesourcemanager/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecureSourceManagerRepositoriesGetter has a method to return a SecureSourceManagerRepositoryInterface.
// A group's client should implement this interface.
type SecureSourceManagerRepositoriesGetter interface {
	SecureSourceManagerRepositories(namespace string) SecureSourceManagerRepositoryInterface
}

// SecureSourceManagerRepositoryInterface has methods to work with SecureSourceManagerRepository resources.
type SecureSourceManagerRepositoryInterface interface {
	Create(ctx context.Context, secureSourceManagerRepository *v1alpha1.SecureSourceManagerRepository, opts v1.CreateOptions) (*v1alpha1.SecureSourceManagerRepository, error)
	Update(ctx context.Context, secureSourceManagerRepository *v1alpha1.SecureSourceManagerRepository, opts v1.UpdateOptions) (*v1alpha1.SecureSourceManagerRepository, error)
	UpdateStatus(ctx context.Context, secureSourceManagerRepository *v1alpha1.SecureSourceManagerRepository, opts v1.UpdateOptions) (*v1alpha1.SecureSourceManagerRepository, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SecureSourceManagerRepository, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SecureSourceManagerRepositoryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecureSourceManagerRepository, err error)
	SecureSourceManagerRepositoryExpansion
}

// secureSourceManagerRepositories implements SecureSourceManagerRepositoryInterface
type secureSourceManagerRepositories struct {
	client rest.Interface
	ns     string
}

// newSecureSourceManagerRepositories returns a SecureSourceManagerRepositories
func newSecureSourceManagerRepositories(c *SecuresourcemanagerV1alpha1Client, namespace string) *secureSourceManagerRepositories {
	return &secureSourceManagerRepositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secureSourceManagerRepository, and returns the corresponding secureSourceManagerRepository object, and an error if there is any.
func (c *secureSourceManagerRepositories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecureSourceManagerRepository, err error) {
	result = &v1alpha1.SecureSourceManagerRepository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecureSourceManagerRepositories that match those selectors.
func (c *secureSourceManagerRepositories) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecureSourceManagerRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecureSourceManagerRepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secureSourceManagerRepositories.
func (c *secureSourceManagerRepositories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a secureSourceManagerRepository and creates it.  Returns the server's representation of the secureSourceManagerRepository, and an error, if there is any.
func (c *secureSourceManagerRepositories) Create(ctx context.Context, secureSourceManagerRepository *v1alpha1.SecureSourceManagerRepository, opts v1.CreateOptions) (result *v1alpha1.SecureSourceManagerRepository, err error) {
	result = &v1alpha1.SecureSourceManagerRepository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secureSourceManagerRepository).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a secureSourceManagerRepository and updates it. Returns the server's representation of the secureSourceManagerRepository, and an error, if there is any.
func (c *secureSourceManagerRepositories) Update(ctx context.Context, secureSourceManagerRepository *v1alpha1.SecureSourceManagerRepository, opts v1.UpdateOptions) (result *v1alpha1.SecureSourceManagerRepository, err error) {
	result = &v1alpha1.SecureSourceManagerRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		Name(secureSourceManagerRepository.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secureSourceManagerRepository).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *secureSourceManagerRepositories) UpdateStatus(ctx context.Context, secureSourceManagerRepository *v1alpha1.SecureSourceManagerRepository, opts v1.UpdateOptions) (result *v1alpha1.SecureSourceManagerRepository, err error) {
	result = &v1alpha1.SecureSourceManagerRepository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		Name(secureSourceManagerRepository.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(secureSourceManagerRepository).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the secureSourceManagerRepository and deletes it. Returns an error if one occurs.
func (c *secureSourceManagerRepositories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secureSourceManagerRepositories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched secureSourceManagerRepository.
func (c *secureSourceManagerRepositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecureSourceManagerRepository, err error) {
	result = &v1alpha1.SecureSourceManagerRepository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("securesourcemanagerrepositories").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}