package cloud

import (
	"context"
	"time"

	cloudv1 "github.com/KETI-ExaScale/exascale-resource-controller/apis/cloud/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type NodeGetter interface {
	Nodes() NodeInterface
}

type NodeInterface interface {
	Create(*cloudv1.Node) (*cloudv1.Node, error)
	Update(*cloudv1.Node) (*cloudv1.Node, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*cloudv1.Node, error)
	List(opts metav1.ListOptions) (*cloudv1.NodeList, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.Node, err error)
}

type nodes struct {
	client rest.Interface
}

func newNodes(c *CloudV1Client) *nodes {
	return &nodes{
		client: c.RESTClient(),
	}

}

func (c *nodes) Get(name string, options metav1.GetOptions) (result *cloudv1.Node, err error) {
	result = &cloudv1.Node{}
	err = c.client.Get().
		Resource("nodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of clusters that match those selectors.
func (c *nodes) List(opts metav1.ListOptions) (result *cloudv1.NodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &cloudv1.NodeList{}
	err = c.client.Get().
		Resource("nodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *nodes) Create(node *cloudv1.Node) (result *cloudv1.Node, err error) {
	result = &cloudv1.Node{}
	err = c.client.Post().
		Resource("nodes").
		Body(node).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *nodes) Update(node *cloudv1.Node) (result *cloudv1.Node, err error) {
	result = &cloudv1.Node{}
	err = c.client.Put().
		Resource("nodes").
		Name(node.Name).
		Body(node).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *nodes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodes").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// Patch applies the patch and returns the patched pod.
func (c *nodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.Node, err error) {
	result = &cloudv1.Node{}
	err = c.client.Patch(pt).
		Resource("nodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.Background()).
		Into(result)
	return
}
