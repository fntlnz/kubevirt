/*
Copyright The KubeVirt Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "kubevirt.io/api/instancetype/v1alpha2"
)

// FakeVirtualMachineClusterInstancetypes implements VirtualMachineClusterInstancetypeInterface
type FakeVirtualMachineClusterInstancetypes struct {
	Fake *FakeInstancetypeV1alpha2
}

var virtualmachineclusterinstancetypesResource = schema.GroupVersionResource{Group: "instancetype.kubevirt.io", Version: "v1alpha2", Resource: "virtualmachineclusterinstancetypes"}

var virtualmachineclusterinstancetypesKind = schema.GroupVersionKind{Group: "instancetype.kubevirt.io", Version: "v1alpha2", Kind: "VirtualMachineClusterInstancetype"}

// Get takes name of the virtualMachineClusterInstancetype, and returns the corresponding virtualMachineClusterInstancetype object, and an error if there is any.
func (c *FakeVirtualMachineClusterInstancetypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.VirtualMachineClusterInstancetype, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(virtualmachineclusterinstancetypesResource, name), &v1alpha2.VirtualMachineClusterInstancetype{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.VirtualMachineClusterInstancetype), err
}

// List takes label and field selectors, and returns the list of VirtualMachineClusterInstancetypes that match those selectors.
func (c *FakeVirtualMachineClusterInstancetypes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.VirtualMachineClusterInstancetypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(virtualmachineclusterinstancetypesResource, virtualmachineclusterinstancetypesKind, opts), &v1alpha2.VirtualMachineClusterInstancetypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.VirtualMachineClusterInstancetypeList{ListMeta: obj.(*v1alpha2.VirtualMachineClusterInstancetypeList).ListMeta}
	for _, item := range obj.(*v1alpha2.VirtualMachineClusterInstancetypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineClusterInstancetypes.
func (c *FakeVirtualMachineClusterInstancetypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(virtualmachineclusterinstancetypesResource, opts))
}

// Create takes the representation of a virtualMachineClusterInstancetype and creates it.  Returns the server's representation of the virtualMachineClusterInstancetype, and an error, if there is any.
func (c *FakeVirtualMachineClusterInstancetypes) Create(ctx context.Context, virtualMachineClusterInstancetype *v1alpha2.VirtualMachineClusterInstancetype, opts v1.CreateOptions) (result *v1alpha2.VirtualMachineClusterInstancetype, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(virtualmachineclusterinstancetypesResource, virtualMachineClusterInstancetype), &v1alpha2.VirtualMachineClusterInstancetype{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.VirtualMachineClusterInstancetype), err
}

// Update takes the representation of a virtualMachineClusterInstancetype and updates it. Returns the server's representation of the virtualMachineClusterInstancetype, and an error, if there is any.
func (c *FakeVirtualMachineClusterInstancetypes) Update(ctx context.Context, virtualMachineClusterInstancetype *v1alpha2.VirtualMachineClusterInstancetype, opts v1.UpdateOptions) (result *v1alpha2.VirtualMachineClusterInstancetype, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(virtualmachineclusterinstancetypesResource, virtualMachineClusterInstancetype), &v1alpha2.VirtualMachineClusterInstancetype{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.VirtualMachineClusterInstancetype), err
}

// Delete takes name of the virtualMachineClusterInstancetype and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineClusterInstancetypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(virtualmachineclusterinstancetypesResource, name), &v1alpha2.VirtualMachineClusterInstancetype{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineClusterInstancetypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(virtualmachineclusterinstancetypesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.VirtualMachineClusterInstancetypeList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineClusterInstancetype.
func (c *FakeVirtualMachineClusterInstancetypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.VirtualMachineClusterInstancetype, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(virtualmachineclusterinstancetypesResource, name, pt, data, subresources...), &v1alpha2.VirtualMachineClusterInstancetype{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.VirtualMachineClusterInstancetype), err
}
