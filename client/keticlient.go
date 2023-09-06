package client

import (
	"fmt"

	"github.com/KETI-ExaScale/exascale-resource-controller/client/cloud"
	"github.com/KETI-ExaScale/exascale-resource-controller/client/metric"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	MetricV1() metric.MetricV1Interface
	CloudV1() cloud.CloudV1Interface
}

type ClientSet struct {
	metricV1 *metric.MetricV1Client
	cloudV1  *cloud.CloudV1Client
}

func NewForConfig(c *rest.Config) (*ClientSet, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("Burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs ClientSet
	var err error
	cs.metricV1, err = metric.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cloudV1, err = cloud.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}

func (c *ClientSet) MetricV1() metric.MetricV1Interface {
	return c.metricV1
}

func (c *ClientSet) CloudV1() cloud.CloudV1Interface {
	return c.cloudV1
}
