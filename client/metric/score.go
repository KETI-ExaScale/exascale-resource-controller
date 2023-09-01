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

type ScoreGetter interface {
	Scores() GPUInterface
}

type ScoreInterface interface {
	Create(*metricv1.Score) (*metricv1.Score, error)
	Update(*metricv1.Score) (*metricv1.Score, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*metricv1.Score, error)
	List(opts metav1.ListOptions) (*metricv1.ScoreList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *metricv1.Score, err error)
}

type scores struct {
	client rest.Interface
}

func newScores(c *MetricV1Client) *scores {
	return &scores{
		client: c.RESTClient(),
	}
}

func (c *scores) Get(name string, options metav1.GetOptions) (result *metricv1.Score, err error) {
	result = &metricv1.Score{}
	err = c.client.Get().
		Resource("scores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *scores) List(opts metav1.ListOptions) (result *metricv1.ScoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &metricv1.ScoreList{}
	err = c.client.Get().
		Resource("scores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *scores) Create(score *metricv1.Score) (result *metricv1.Score, err error) {
	result = &metricv1.Score{}
	err = c.client.Post().
		Resource("scores").
		Body(score).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *scores) Update(score *metricv1.Score) (result *metricv1.Score, err error) {
	result = &metricv1.Score{}
	err = c.client.Put().
		Resource("scores").
		Name(score.Name).
		Body(score).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *scores) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("scores").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *scores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *metricv1.Score, err error) {
	result = &metricv1.Score{}
	err = c.client.Patch(pt).
		Resource("scores").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}
