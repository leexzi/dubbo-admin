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

package fake

import (
	"context"
	v1beta1 "github.com/apache/dubbo-admin/pkg/rule/apis/dubbo.apache.org/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceNameMappings implements ServiceNameMappingInterface
type FakeServiceNameMappings struct {
	Fake *FakeDubboV1beta1
	ns   string
}

var servicenamemappingsResource = schema.GroupVersionResource{Group: "dubbo.apache.org", Version: "v1beta1", Resource: "servicenamemappings"}

var servicenamemappingsKind = schema.GroupVersionKind{Group: "dubbo.apache.org", Version: "v1beta1", Kind: "ServiceNameMapping"}

// Get takes name of the serviceNameMapping, and returns the corresponding serviceNameMapping object, and an error if there is any.
func (c *FakeServiceNameMappings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ServiceNameMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicenamemappingsResource, c.ns, name), &v1beta1.ServiceNameMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceNameMapping), err
}

// List takes label and field selectors, and returns the list of ServiceNameMappings that match those selectors.
func (c *FakeServiceNameMappings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ServiceNameMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicenamemappingsResource, servicenamemappingsKind, c.ns, opts), &v1beta1.ServiceNameMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceNameMappingList{ListMeta: obj.(*v1beta1.ServiceNameMappingList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceNameMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceNameMappings.
func (c *FakeServiceNameMappings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicenamemappingsResource, c.ns, opts))

}

// Create takes the representation of a serviceNameMapping and creates it.  Returns the server's representation of the serviceNameMapping, and an error, if there is any.
func (c *FakeServiceNameMappings) Create(ctx context.Context, serviceNameMapping *v1beta1.ServiceNameMapping, opts v1.CreateOptions) (result *v1beta1.ServiceNameMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicenamemappingsResource, c.ns, serviceNameMapping), &v1beta1.ServiceNameMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceNameMapping), err
}

// Update takes the representation of a serviceNameMapping and updates it. Returns the server's representation of the serviceNameMapping, and an error, if there is any.
func (c *FakeServiceNameMappings) Update(ctx context.Context, serviceNameMapping *v1beta1.ServiceNameMapping, opts v1.UpdateOptions) (result *v1beta1.ServiceNameMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicenamemappingsResource, c.ns, serviceNameMapping), &v1beta1.ServiceNameMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceNameMapping), err
}

// Delete takes name of the serviceNameMapping and deletes it. Returns an error if one occurs.
func (c *FakeServiceNameMappings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(servicenamemappingsResource, c.ns, name, opts), &v1beta1.ServiceNameMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceNameMappings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicenamemappingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ServiceNameMappingList{})
	return err
}

// Patch applies the patch and returns the patched serviceNameMapping.
func (c *FakeServiceNameMappings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ServiceNameMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicenamemappingsResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceNameMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceNameMapping), err
}