/* auto-generated */

package serviceaccount

import "github.com/controlplane-com/types-go/pkg/base"

type ServiceAccountKey struct {
	Name        string `json:"name,omitempty"`
	Created     string `json:"created,omitempty"`
	Key         string `json:"key,omitempty"`
	Description string `json:"description,omitempty"`
}

type ServiceAccountOrigin string

const (
	ServiceAccountOriginDefault ServiceAccountOrigin = "default"
	ServiceAccountOriginBuiltin ServiceAccountOrigin = "builtin"
)

type ServiceAccount struct {
	Id           string               `json:"id,omitempty"`
	Name         base.Name            `json:"name,omitempty"`
	Kind         base.Kind            `json:"kind,omitempty"`
	Version      float32              `json:"version"`
	Description  string               `json:"description,omitempty"`
	Tags         base.Tags            `json:"tags,omitempty"`
	Created      string               `json:"created,omitempty"`
	LastModified string               `json:"lastModified,omitempty"`
	Links        base.Links           `json:"links,omitempty"`
	Keys         []ServiceAccountKey  `json:"keys,omitempty"`
	Origin       ServiceAccountOrigin `json:"origin,omitempty"`
}
