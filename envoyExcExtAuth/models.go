/* auto-generated */

package envoyExcExtAuth

import "github.com/controlplane-com/types-go/envoyCommon"
import "github.com/controlplane-com/types-go/port"

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
