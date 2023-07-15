// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/*
Copyright The Kubernetes Authors.

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

// Code generated by clientgen-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	v1beta1 "github.com/apache/dubbo-admin/pkg/rule/apis/dubbo.apache.org/v1beta1"
	scheme "github.com/apache/dubbo-admin/pkg/rule/clientgen/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuthorizationPoliciesGetter has a method to return a AuthorizationPolicyInterface.
// A group's clientgen should implement this interface.
type AuthorizationPoliciesGetter interface {
	AuthorizationPolicies(namespace string) AuthorizationPolicyInterface
}

// AuthorizationPolicyInterface has methods to work with AuthorizationPolicy resources.
type AuthorizationPolicyInterface interface {
	Create(ctx context.Context, authorizationPolicy *v1beta1.AuthorizationPolicy, opts v1.CreateOptions) (*v1beta1.AuthorizationPolicy, error)
	Update(ctx context.Context, authorizationPolicy *v1beta1.AuthorizationPolicy, opts v1.UpdateOptions) (*v1beta1.AuthorizationPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.AuthorizationPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.AuthorizationPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AuthorizationPolicy, err error)
	AuthorizationPolicyExpansion
}

// authorizationPolicies implements AuthorizationPolicyInterface
type authorizationPolicies struct {
	client rest.Interface
	ns     string
}

// newAuthorizationPolicies returns a AuthorizationPolicies
func newAuthorizationPolicies(c *DubboV1beta1Client, namespace string) *authorizationPolicies {
	return &authorizationPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the authorizationPolicy, and returns the corresponding authorizationPolicy object, and an error if there is any.
func (c *authorizationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AuthorizationPolicies that match those selectors.
func (c *authorizationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AuthorizationPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.AuthorizationPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authorizationPolicies.
func (c *authorizationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a authorizationPolicy and creates it.  Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *authorizationPolicies) Create(ctx context.Context, authorizationPolicy *v1beta1.AuthorizationPolicy, opts v1.CreateOptions) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authorizationPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a authorizationPolicy and updates it. Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *authorizationPolicies) Update(ctx context.Context, authorizationPolicy *v1beta1.AuthorizationPolicy, opts v1.UpdateOptions) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(authorizationPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authorizationPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the authorizationPolicy and deletes it. Returns an error if one occurs.
func (c *authorizationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authorizationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authorizationpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched authorizationPolicy.
func (c *authorizationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AuthorizationPolicy, err error) {
	result = &v1beta1.AuthorizationPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("authorizationpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
