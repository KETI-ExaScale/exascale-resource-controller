package metric

import (
	metricv1 "github.com/KETI-ExaScale/exascale-resource-controller/apis/metric/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type MetricV1Interface interface {
	RESTClient() rest.Interface
	GPUGetter
	ScoreGetter
	CollectorGetter
}

type MetricV1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*MetricV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: metricv1.GroupVersion.Group, Version: metricv1.GroupVersion.Version}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &MetricV1Client{restClient: client}, nil
}

func (c *MetricV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *MetricV1Client) GPUs() GPUInterface {
	return newGPUs(c)
}

func (c *MetricV1Client) Scores() ScoreInterface {
	return newScores(c)
}

func (c *MetricV1Client) Collectors() CollectorInterface {
	return newCollectors(c)
}
