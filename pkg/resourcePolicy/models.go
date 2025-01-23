/* auto-generated */

package resourcePolicy

import "github.com/controlplane-com/types-go/pkg/base"

type ResourcePolicyBindings struct {
	Permissions    []string `json:"permissions,omitempty"`
	PrincipalLinks []string `json:"principalLinks,omitempty"`
}

type ResourcePolicy struct {
	Id           string                   `json:"id,omitempty"`
	Kind         base.Kind                `json:"kind,omitempty"`
	Version      float32                  `json:"version"`
	Created      string                   `json:"created,omitempty"`
	LastModified string                   `json:"lastModified,omitempty"`
	Links        base.Links               `json:"links,omitempty"`
	Bindings     []ResourcePolicyBindings `json:"bindings,omitempty"`
}
