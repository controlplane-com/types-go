/* auto-generated */

package group

import "github.com/controlplane-com/types-go/base"
import "github.com/controlplane-com/types-go/query"

type GroupIdentityMatcherLanguage string

const (
	GroupIdentityMatcherLanguageJmespath   GroupIdentityMatcherLanguage = "jmespath"
	GroupIdentityMatcherLanguageJavascript GroupIdentityMatcherLanguage = "javascript"
)

type GroupIdentityMatcher struct {
	Expression string                       `json:"expression,omitempty"`
	Language   GroupIdentityMatcherLanguage `json:"language,omitempty"`
}

type GroupOrigin string

const (
	GroupOriginSynthetic GroupOrigin = "synthetic"
	GroupOriginDefault   GroupOrigin = "default"
	GroupOriginBuiltin   GroupOrigin = "builtin"
)

type Group struct {
	Id              string               `json:"id,omitempty"`
	Name            base.Name            `json:"name,omitempty"`
	Kind            base.Kind            `json:"kind,omitempty"`
	Version         float32              `json:"version"`
	Description     string               `json:"description,omitempty"`
	Tags            base.Tags            `json:"tags,omitempty"`
	Created         string               `json:"created,omitempty"`
	LastModified    string               `json:"lastModified,omitempty"`
	Links           base.Links           `json:"links,omitempty"`
	MemberLinks     []string             `json:"memberLinks,omitempty"`
	MemberQuery     query.Query          `json:"memberQuery,omitempty"`
	IdentityMatcher GroupIdentityMatcher `json:"identityMatcher,omitempty"`
	Origin          GroupOrigin          `json:"origin,omitempty"`
}
