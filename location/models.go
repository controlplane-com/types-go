/* auto-generated */

package location

import "github.com/controlplane-com/types-go/base"

type LocationSpec struct {
	Enabled bool `json:"enabled,omitempty"`
}

type LocationStatusGeo struct {
	Lat       float32 `json:"lat"`
	Lon       float32 `json:"lon"`
	Country   string  `json:"country,omitempty"`
	State     string  `json:"state,omitempty"`
	City      string  `json:"city,omitempty"`
	Continent string  `json:"continent,omitempty"`
}

type LocationStatus struct {
	Geo      LocationStatusGeo `json:"geo,omitempty"`
	IpRanges []string          `json:"ipRanges,omitempty"`
}

type LocationOrigin string

const (
	LocationOriginBuiltin LocationOrigin = "builtin"
	LocationOriginDefault LocationOrigin = "default"
	LocationOriginCustom  LocationOrigin = "custom"
)

type LocationProvider string

const (
	LocationProviderAws     LocationProvider = "aws"
	LocationProviderGcp     LocationProvider = "gcp"
	LocationProviderAzure   LocationProvider = "azure"
	LocationProviderByok    LocationProvider = "byok"
	LocationProviderLinode  LocationProvider = "linode"
	LocationProviderVultr   LocationProvider = "vultr"
	LocationProviderEquinix LocationProvider = "equinix"
)

type Location struct {
	Id           string           `json:"id,omitempty"`
	Name         base.Name        `json:"name,omitempty"`
	Kind         base.Kind        `json:"kind,omitempty"`
	Version      float32          `json:"version"`
	Description  string           `json:"description,omitempty"`
	Tags         base.Tags        `json:"tags,omitempty"`
	Created      string           `json:"created,omitempty"`
	LastModified string           `json:"lastModified,omitempty"`
	Links        base.Links       `json:"links,omitempty"`
	Origin       LocationOrigin   `json:"origin,omitempty"`
	Provider     LocationProvider `json:"provider,omitempty"`
	Region       string           `json:"region,omitempty"`
	Spec         LocationSpec     `json:"spec,omitempty"`
	Status       LocationStatus   `json:"status,omitempty"`
}
