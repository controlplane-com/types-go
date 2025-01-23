/* auto-generated */

package envoy

import "gitlab.com/controlplane/controlplane/go-libs/schema/envoyAccessLog"
import "gitlab.com/controlplane/controlplane/go-libs/schema/envoyCluster"
import "gitlab.com/controlplane/controlplane/go-libs/schema/envoyExcExtAuth"
import "gitlab.com/controlplane/controlplane/go-libs/schema/envoyHttp"
import "gitlab.com/controlplane/controlplane/go-libs/schema/volumeSpec"

type EnvoyFilters struct {
	AccessLog            []envoyAccessLog.AccessLog          `json:"accessLog,omitempty"`
	Clusters             []envoyCluster.Cluster              `json:"clusters,omitempty"`
	ExcludedExternalAuth []envoyExcExtAuth.ExcExtAuth        `json:"excludedExternalAuth,omitempty"`
	ExcludedRateLimit    []envoyExcExtAuth.ExcludedRateLimit `json:"excludedRateLimit,omitempty"`
	Http                 []envoyHttp.HttpFilter              `json:"http,omitempty"`
	Network              []any                               `json:"network,omitempty"`
	Volumes              []volumeSpec.VolumeSpec             `json:"volumes,omitempty"`
}
