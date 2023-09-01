package cloud

import (
	"context"
	"time"

	cloudv1 "github.com/KETI-ExaScale/exascale-resource-controller/apis/cloud/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type ClusterGetter interface {
	Clusters() ClusterInterface
}

type ClusterInterface interface {
	Create(*cloudv1.Cluster) (*cloudv1.Cluster, error)
	Update(*cloudv1.Cluster) (*cloudv1.Cluster, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*cloudv1.Cluster, error)
	List(opts metav1.ListOptions) (*cloudv1.ClusterList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.Cluster, err error)
}

type clusters struct {
	client rest.Interface
}

func newClusters(c *CloudV1Client) *clusters {
	return &clusters{
		client: c.RESTClient(),
	}
}

func (c *clusters) Get(name string, options metav1.GetOptions) (result *cloudv1.Cluster, err error) {
	result = &cloudv1.Cluster{}
	err = c.client.Get().
		Resource("clusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *clusters) List(opts metav1.ListOptions) (result *cloudv1.ClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &cloudv1.ClusterList{}
	err = c.client.Get().
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *clusters) Create(cluster *cloudv1.Cluster) (result *cloudv1.Cluster, err error) {
	result = &cloudv1.Cluster{}
	err = c.client.Post().
		Resource("clusters").
		Body(cluster).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *clusters) Update(cluster *cloudv1.Cluster) (result *cloudv1.Cluster, err error) {
	result = &cloudv1.Cluster{}
	err = c.client.Put().
		Resource("clusters").
		Name(cluster.Name).
		Body(cluster).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *clusters) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusters").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *clusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.Cluster, err error) {
	result = &cloudv1.Cluster{}
	err = c.client.Patch(pt).
		Resource("clusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}
