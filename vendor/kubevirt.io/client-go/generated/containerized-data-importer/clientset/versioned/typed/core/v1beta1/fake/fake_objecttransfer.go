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
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// FakeObjectTransfers implements ObjectTransferInterface
type FakeObjectTransfers struct {
	Fake *FakeCdiV1beta1
}

var objecttransfersResource = schema.GroupVersionResource{Group: "cdi.kubevirt.io", Version: "v1beta1", Resource: "objecttransfers"}

var objecttransfersKind = schema.GroupVersionKind{Group: "cdi.kubevirt.io", Version: "v1beta1", Kind: "ObjectTransfer"}

// Get takes name of the objectTransfer, and returns the corresponding objectTransfer object, and an error if there is any.
func (c *FakeObjectTransfers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ObjectTransfer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(objecttransfersResource, name), &v1beta1.ObjectTransfer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectTransfer), err
}

// List takes label and field selectors, and returns the list of ObjectTransfers that match those selectors.
func (c *FakeObjectTransfers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ObjectTransferList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(objecttransfersResource, objecttransfersKind, opts), &v1beta1.ObjectTransferList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ObjectTransferList{ListMeta: obj.(*v1beta1.ObjectTransferList).ListMeta}
	for _, item := range obj.(*v1beta1.ObjectTransferList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested objectTransfers.
func (c *FakeObjectTransfers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(objecttransfersResource, opts))
}

// Create takes the representation of a objectTransfer and creates it.  Returns the server's representation of the objectTransfer, and an error, if there is any.
func (c *FakeObjectTransfers) Create(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.CreateOptions) (result *v1beta1.ObjectTransfer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(objecttransfersResource, objectTransfer), &v1beta1.ObjectTransfer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectTransfer), err
}

// Update takes the representation of a objectTransfer and updates it. Returns the server's representation of the objectTransfer, and an error, if there is any.
func (c *FakeObjectTransfers) Update(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.UpdateOptions) (result *v1beta1.ObjectTransfer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(objecttransfersResource, objectTransfer), &v1beta1.ObjectTransfer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectTransfer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeObjectTransfers) UpdateStatus(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.UpdateOptions) (*v1beta1.ObjectTransfer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(objecttransfersResource, "status", objectTransfer), &v1beta1.ObjectTransfer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectTransfer), err
}

// Delete takes name of the objectTransfer and deletes it. Returns an error if one occurs.
func (c *FakeObjectTransfers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(objecttransfersResource, name), &v1beta1.ObjectTransfer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeObjectTransfers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(objecttransfersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ObjectTransferList{})
	return err
}

// Patch applies the patch and returns the patched objectTransfer.
func (c *FakeObjectTransfers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ObjectTransfer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(objecttransfersResource, name, pt, data, subresources...), &v1beta1.ObjectTransfer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectTransfer), err
}
