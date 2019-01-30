// Copyright 2018 The prometheus-operator Authors
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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	monitoring_v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceMonitors implements ServiceMonitorInterface
type FakeServiceMonitors struct {
	Fake *FakeMonitoringV1
	ns   string
}

var servicemonitorsResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "servicemonitors"}

var servicemonitorsKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "ServiceMonitor"}

// Get takes name of the serviceMonitor, and returns the corresponding serviceMonitor object, and an error if there is any.
func (c *FakeServiceMonitors) Get(name string, options v1.GetOptions) (result *monitoring_v1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemonitorsResource, c.ns, name), &monitoring_v1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}

// List takes label and field selectors, and returns the list of ServiceMonitors that match those selectors.
func (c *FakeServiceMonitors) List(opts v1.ListOptions) (result *monitoring_v1.ServiceMonitorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemonitorsResource, servicemonitorsKind, c.ns, opts), &monitoring_v1.ServiceMonitorList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitorList), err
}

// Watch returns a watch.Interface that watches the requested serviceMonitors.
func (c *FakeServiceMonitors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemonitorsResource, c.ns, opts))

}

// Create takes the representation of a serviceMonitor and creates it.  Returns the server's representation of the serviceMonitor, and an error, if there is any.
func (c *FakeServiceMonitors) Create(serviceMonitor *monitoring_v1.ServiceMonitor) (result *monitoring_v1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemonitorsResource, c.ns, serviceMonitor), &monitoring_v1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}

// Update takes the representation of a serviceMonitor and updates it. Returns the server's representation of the serviceMonitor, and an error, if there is any.
func (c *FakeServiceMonitors) Update(serviceMonitor *monitoring_v1.ServiceMonitor) (result *monitoring_v1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemonitorsResource, c.ns, serviceMonitor), &monitoring_v1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}

// Delete takes name of the serviceMonitor and deletes it. Returns an error if one occurs.
func (c *FakeServiceMonitors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicemonitorsResource, c.ns, name), &monitoring_v1.ServiceMonitor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMonitors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemonitorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &monitoring_v1.ServiceMonitorList{})
	return err
}

// Patch applies the patch and returns the patched serviceMonitor.
func (c *FakeServiceMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *monitoring_v1.ServiceMonitor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemonitorsResource, c.ns, name, data, subresources...), &monitoring_v1.ServiceMonitor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoring_v1.ServiceMonitor), err
}
