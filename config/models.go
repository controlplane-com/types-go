/* auto-generated */

package config

import "gitlab.com/controlplane/controlplane/go-libs/schema/base"

type ConfigKind string

const (
	ConfigKindConfig ConfigKind = "config"
)

type Config[D any] struct {
	Kind         ConfigKind `json:"kind,omitempty"`
	Id           string     `json:"id,omitempty"`
	LastModified string     `json:"lastModified,omitempty"`
	Version      float32    `json:"version,omitempty"`
	Data         D          `json:"data,omitempty"`
	Links        base.Links `json:"links,omitempty"`
}
