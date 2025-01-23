/* auto-generated */

package envoy

import "github.com/controlplane-com/types-go/envoyAccessLog"
import "github.com/controlplane-com/types-go/envoyCluster"
import "github.com/controlplane-com/types-go/envoyExcExtAuth"
import "github.com/controlplane-com/types-go/envoyHttp"
import "github.com/controlplane-com/types-go/volumeSpec"

type EnvoyFilters struct {
	AccessLog            []envoyAccessLog.AccessLog          `json:"accessLog,omitempty"`
	Clusters             []envoyCluster.Cluster              `json:"clusters,omitempty"`
	ExcludedExternalAuth []envoyExcExtAuth.ExcExtAuth        `json:"excludedExternalAuth,omitempty"`
	ExcludedRateLimit    []envoyExcExtAuth.ExcludedRateLimit `json:"excludedRateLimit,omitempty"`
	Http                 []envoyHttp.HttpFilter              `json:"http,omitempty"`
	Network              []any                               `json:"network,omitempty"`
	Volumes              []volumeSpec.VolumeSpec             `json:"volumes,omitempty"`
}
