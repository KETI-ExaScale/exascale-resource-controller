package metric

import (
	"context"
	"time"

	metricv1 "github.com/KETI-ExaScale/exascale-resource-controller/apis/metric/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type CollectorGetter interface {
	Collectors() CollectorInterface
}

type CollectorInterface interface {
	Create(*metricv1.Collector) (*metricv1.Collector, error)
	Update(*metricv1.Collector) (*metricv1.Collector, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*metricv1.Collector, error)
	List(opts metav1.ListOptions) (*metricv1.CollectorList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *metricv1.Collector, err error)
}

type collectors struct {
	client rest.Interface
}

func newCollectors(c *MetricV1Client) *collectors {
	return &collectors{
		client: c.RESTClient(),
	}
}

func (c *collectors) Get(name string, options metav1.GetOptions) (result *metricv1.Collector, err error) {
	result = &metricv1.Collector{}
	err = c.client.Get().
		Resource("collectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *collectors) List(opts metav1.ListOptions) (result *metricv1.CollectorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &metricv1.CollectorList{}
	err = c.client.Get().
		Resource("collectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *collectors) Create(collector *metricv1.Collector) (result *metricv1.Collector, err error) {
	result = &metricv1.Collector{}
	err = c.client.Post().
		Resource("collectors").
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *collectors) Update(collector *metricv1.Collector) (result *metricv1.Collector, err error) {
	result = &metricv1.Collector{}
	err = c.client.Put().
		Resource("collectors").
		Name(collector.Name).
		Body(collector).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *collectors) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("collectors").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *collectors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *metricv1.Collector, err error) {
	result = &metricv1.Collector{}
	err = c.client.Patch(pt).
		Resource("collectors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}
