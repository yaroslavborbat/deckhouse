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
	v1alpha1 "bashible-apiserver/pkg/apis/bashible/v1alpha1"
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubernetesBundles implements KubernetesBundleInterface
type FakeKubernetesBundles struct {
	Fake *FakeBashibleV1alpha1
}

var kubernetesbundlesResource = schema.GroupVersionResource{Group: "bashible.deckhouse.io", Version: "v1alpha1", Resource: "kubernetesbundles"}

var kubernetesbundlesKind = schema.GroupVersionKind{Group: "bashible.deckhouse.io", Version: "v1alpha1", Kind: "KubernetesBundle"}

// Get takes name of the kubernetesBundle, and returns the corresponding kubernetesBundle object, and an error if there is any.
func (c *FakeKubernetesBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubernetesBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubernetesbundlesResource, name), &v1alpha1.KubernetesBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesBundle), err
}

// List takes label and field selectors, and returns the list of KubernetesBundles that match those selectors.
func (c *FakeKubernetesBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubernetesBundleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubernetesbundlesResource, kubernetesbundlesKind, opts), &v1alpha1.KubernetesBundleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubernetesBundleList{ListMeta: obj.(*v1alpha1.KubernetesBundleList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubernetesBundleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubernetesBundles.
func (c *FakeKubernetesBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubernetesbundlesResource, opts))
}

// Create takes the representation of a kubernetesBundle and creates it.  Returns the server's representation of the kubernetesBundle, and an error, if there is any.
func (c *FakeKubernetesBundles) Create(ctx context.Context, kubernetesBundle *v1alpha1.KubernetesBundle, opts v1.CreateOptions) (result *v1alpha1.KubernetesBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubernetesbundlesResource, kubernetesBundle), &v1alpha1.KubernetesBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesBundle), err
}

// Update takes the representation of a kubernetesBundle and updates it. Returns the server's representation of the kubernetesBundle, and an error, if there is any.
func (c *FakeKubernetesBundles) Update(ctx context.Context, kubernetesBundle *v1alpha1.KubernetesBundle, opts v1.UpdateOptions) (result *v1alpha1.KubernetesBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubernetesbundlesResource, kubernetesBundle), &v1alpha1.KubernetesBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesBundle), err
}

// Delete takes name of the kubernetesBundle and deletes it. Returns an error if one occurs.
func (c *FakeKubernetesBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kubernetesbundlesResource, name), &v1alpha1.KubernetesBundle{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubernetesBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubernetesbundlesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubernetesBundleList{})
	return err
}

// Patch applies the patch and returns the patched kubernetesBundle.
func (c *FakeKubernetesBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubernetesBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubernetesbundlesResource, name, pt, data, subresources...), &v1alpha1.KubernetesBundle{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesBundle), err
}
