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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	cisv1 "github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngressLinks implements IngressLinkInterface
type FakeIngressLinks struct {
	Fake *FakeK8sV1
	ns   string
}

var ingresslinksResource = schema.GroupVersionResource{Group: "k8s.nginx.org", Version: "v1", Resource: "ingresslinks"}

var ingresslinksKind = schema.GroupVersionKind{Group: "k8s.nginx.org", Version: "v1", Kind: "IngressLink"}

// Get takes name of the ingressLink, and returns the corresponding ingressLink object, and an error if there is any.
func (c *FakeIngressLinks) Get(name string, options v1.GetOptions) (result *cisv1.IngressLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingresslinksResource, c.ns, name), &cisv1.IngressLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cisv1.IngressLink), err
}

// List takes label and field selectors, and returns the list of IngressLinks that match those selectors.
func (c *FakeIngressLinks) List(opts v1.ListOptions) (result *cisv1.IngressLinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingresslinksResource, ingresslinksKind, c.ns, opts), &cisv1.IngressLinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cisv1.IngressLinkList{ListMeta: obj.(*cisv1.IngressLinkList).ListMeta}
	for _, item := range obj.(*cisv1.IngressLinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressLinks.
func (c *FakeIngressLinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingresslinksResource, c.ns, opts))

}

// Create takes the representation of a ingressLink and creates it.  Returns the server's representation of the ingressLink, and an error, if there is any.
func (c *FakeIngressLinks) Create(ingressLink *cisv1.IngressLink) (result *cisv1.IngressLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingresslinksResource, c.ns, ingressLink), &cisv1.IngressLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cisv1.IngressLink), err
}

// Update takes the representation of a ingressLink and updates it. Returns the server's representation of the ingressLink, and an error, if there is any.
func (c *FakeIngressLinks) Update(ingressLink *cisv1.IngressLink) (result *cisv1.IngressLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingresslinksResource, c.ns, ingressLink), &cisv1.IngressLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cisv1.IngressLink), err
}

// Delete takes name of the ingressLink and deletes it. Returns an error if one occurs.
func (c *FakeIngressLinks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ingresslinksResource, c.ns, name), &cisv1.IngressLink{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressLinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingresslinksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cisv1.IngressLinkList{})
	return err
}

// Patch applies the patch and returns the patched ingressLink.
func (c *FakeIngressLinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cisv1.IngressLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingresslinksResource, c.ns, name, pt, data, subresources...), &cisv1.IngressLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cisv1.IngressLink), err
}