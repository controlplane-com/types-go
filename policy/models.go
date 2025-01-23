/* auto-generated */

package policy

import "gitlab.com/controlplane/controlplane/go-libs/schema/base"
import "gitlab.com/controlplane/controlplane/go-libs/schema/query"

type Binding struct {
	Permissions    []string `json:"permissions,omitempty"`
	PrincipalLinks []string `json:"principalLinks,omitempty"`
}

type PolicyTarget string

const (
	PolicyTargetAll PolicyTarget = "all"
)

type PolicyOrigin string

const (
	PolicyOriginDefault PolicyOrigin = "default"
	PolicyOriginBuiltin PolicyOrigin = "builtin"
)

type Policy struct {
	Id           string       `json:"id,omitempty"`
	Name         base.Name    `json:"name,omitempty"`
	Kind         base.Kind    `json:"kind,omitempty"`
	Version      float32      `json:"version"`
	Description  string       `json:"description,omitempty"`
	Tags         base.Tags    `json:"tags,omitempty"`
	Created      string       `json:"created,omitempty"`
	LastModified string       `json:"lastModified,omitempty"`
	Links        base.Links   `json:"links,omitempty"`
	TargetKind   base.Kind    `json:"targetKind,omitempty"`
	TargetLinks  []string     `json:"targetLinks,omitempty"`
	TargetQuery  query.Query  `json:"targetQuery,omitempty"`
	Target       PolicyTarget `json:"target,omitempty"`
	Origin       PolicyOrigin `json:"origin,omitempty"`
	Bindings     []Binding    `json:"bindings,omitempty"`
}
