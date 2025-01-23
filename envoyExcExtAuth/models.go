/* auto-generated */

package envoyExcExtAuth

import "gitlab.com/controlplane/controlplane/go-libs/schema/envoyCommon"
import "gitlab.com/controlplane/controlplane/go-libs/schema/port"

type ExcExtAuth struct {
	Match   envoyCommon.RouteMatch `json:"match,omitempty"`
	Port    port.Port              `json:"port,omitempty"`
	SvcPort port.Port              `json:"svcPort,omitempty"`
}

type ExcludedRateLimit struct {
	Match   envoyCommon.RouteMatch `json:"match,omitempty"`
	Port    port.Port              `json:"port,omitempty"`
	SvcPort port.Port              `json:"svcPort,omitempty"`
}
