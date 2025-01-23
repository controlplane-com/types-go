/* auto-generated */

package ipSet

import "github.com/controlplane-com/types-go/base"

type IpAddressState string

const (
	IpAddressStateBound   IpAddressState = "bound"
	IpAddressStateUnbound IpAddressState = "unbound"
)

type IpAddress struct {
	Name    string         `json:"name,omitempty"`
	Ip      string         `json:"ip,omitempty"`
	Id      string         `json:"id,omitempty"`
	State   IpAddressState `json:"state,omitempty"`
	Created string         `json:"created,omitempty"`
}

type IpSetStatus struct {
	IpAddresses []IpAddress `json:"ipAddresses,omitempty"`
	Error       string      `json:"error,omitempty"`
}

type IpSetLocationRetentionPolicy string

const (
	IpSetLocationRetentionPolicyKeep IpSetLocationRetentionPolicy = "keep"
	IpSetLocationRetentionPolicyFree IpSetLocationRetentionPolicy = "free"
)

type IpSetLocation struct {
	Name            string                       `json:"name,omitempty"`
	RetentionPolicy IpSetLocationRetentionPolicy `json:"retentionPolicy,omitempty"`
}

type IpSetSpec struct {
	Link      string          `json:"link,omitempty"`
	Locations []IpSetLocation `json:"locations,omitempty"`
}

type IpSet struct {
	Id           string      `json:"id,omitempty"`
	Name         base.Name   `json:"name,omitempty"`
	Kind         base.Kind   `json:"kind,omitempty"`
	Version      float32     `json:"version"`
	Description  string      `json:"description,omitempty"`
	Tags         base.Tags   `json:"tags,omitempty"`
	Created      string      `json:"created,omitempty"`
	LastModified string      `json:"lastModified,omitempty"`
	Links        base.Links  `json:"links,omitempty"`
	Spec         IpSetSpec   `json:"spec,omitempty"`
	Status       IpSetStatus `json:"status,omitempty"`
}
