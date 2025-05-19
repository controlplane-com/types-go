/* auto-generated */

package locationDns

type LocationDnsMessageOperation string

const (
	LocationDnsMessageOperationUpdate LocationDnsMessageOperation = "update"
	LocationDnsMessageOperationDelete LocationDnsMessageOperation = "delete"
)

type LocationDnsMessage struct {
	Org              string                      `json:"org,omitempty"`
	Gvc              string                      `json:"gvc,omitempty"`
	GvcAlias         string                      `json:"gvcAlias,omitempty"`
	Workload         string                      `json:"workload,omitempty"`
	Operation        LocationDnsMessageOperation `json:"operation,omitempty"`
	DeleteReason     string                      `json:"deleteReason,omitempty"`
	Name             string                      `json:"name,omitempty"`
	Target           string                      `json:"target,omitempty"`
	AvailabilityZone string                      `json:"availabilityZone,omitempty"`
}
