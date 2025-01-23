/* auto-generated */

package gvc

import (
	"github.com/controlplane-com/types-go/pkg/query"
)
import "github.com/controlplane-com/types-go/pkg/tracing"
import "github.com/controlplane-com/types-go/pkg/env"
import "github.com/controlplane-com/types-go/pkg/base"

type GvcSpecLoadBalancer struct {
	Dedicated bool `json:"dedicated,omitempty"`
}

type GvcConfigCluster struct {
	ClusterId string `json:"clusterId,omitempty"`
	Since     string `json:"since,omitempty"`
}

type GvcConfigClusters map[string]GvcConfigCluster

type GvcStatus map[string]any

type StaticPlacement struct {
	LocationLinks []string    `json:"locationLinks,omitempty"`
	LocationQuery query.Query `json:"locationQuery,omitempty"`
}

type GvcSpec struct {
	StaticPlacement StaticPlacement     `json:"staticPlacement,omitempty"`
	PullSecretLinks []string            `json:"pullSecretLinks,omitempty"`
	Domain          string              `json:"domain,omitempty"`
	Tracing         tracing.Tracing     `json:"tracing,omitempty"`
	Env             []env.EnvVar        `json:"env,omitempty"`
	LoadBalancer    GvcSpecLoadBalancer `json:"loadBalancer,omitempty"`
}

type Gvc struct {
	Id           string     `json:"id,omitempty"`
	Name         base.Name  `json:"name,omitempty"`
	Kind         base.Kind  `json:"kind,omitempty"`
	Version      float32    `json:"version,omitempty"`
	Description  string     `json:"description,omitempty"`
	Tags         base.Tags  `json:"tags,omitempty"`
	Created      string     `json:"created,omitempty"`
	LastModified string     `json:"lastModified,omitempty"`
	Links        base.Links `json:"links,omitempty"`
	Alias        string     `json:"alias,omitempty"`
	Spec         GvcSpec    `json:"spec,omitempty"`
	Status       GvcStatus  `json:"status,omitempty"`
}

type GvcConfig struct {
	Clusters GvcConfigClusters `json:"clusters,omitempty"`
}
