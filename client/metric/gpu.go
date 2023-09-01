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

type GPUGetter interface {
	GPUs() GPUInterface
}

type GPUInterface interface {
	Create(*metricv1.GPU) (*metricv1.GPU, error)
	Update(*metricv1.GPU) (*metricv1.GPU, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*metricv1.GPU, error)
	List(opts metav1.ListOptions) (*metricv1.GPUList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *metricv1.GPU, err error)
}

type gpus struct {
	client rest.Interface
}

func newGPUs(c *MetricV1Client) *gpus {
	return &gpus{
		client: c.RESTClient(),
	}
}

func (c *gpus) Get(name string, options metav1.GetOptions) (result *metricv1.GPU, err error) {
	result = &metricv1.GPU{}
	err = c.client.Get().
		Resource("gpus").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *gpus) List(opts metav1.ListOptions) (result *metricv1.GPUList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &metricv1.GPUList{}
	err = c.client.Get().
		Resource("gpus").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *gpus) Create(gpu *metricv1.GPU) (result *metricv1.GPU, err error) {
	result = &metricv1.GPU{}
	err = c.client.Post().
		Resource("gpus").
		Body(gpu).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *gpus) Update(gpu *metricv1.GPU) (result *metricv1.GPU, err error) {
	result = &metricv1.GPU{}
	err = c.client.Put().
		Resource("gpus").
		Name(gpu.Name).
		Body(gpu).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *gpus) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("gpus").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *gpus) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *metricv1.GPU, err error) {
	result = &metricv1.GPU{}
	err = c.client.Patch(pt).
		Resource("gpus").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}
