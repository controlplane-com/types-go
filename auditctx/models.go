/* auto-generated */

package auditctx

import "gitlab.com/controlplane/controlplane/go-libs/schema/base"

type AuditContextStatus map[string]any

type AuditContextOrigin string

const (
	AuditContextOriginDefault AuditContextOrigin = "default"
	AuditContextOriginBuiltin AuditContextOrigin = "builtin"
)

type AuditContext struct {
	Id           string             `json:"id,omitempty"`
	Name         base.Name          `json:"name,omitempty"`
	Kind         base.Kind          `json:"kind,omitempty"`
	Version      float32            `json:"version"`
	Description  string             `json:"description,omitempty"`
	Tags         base.Tags          `json:"tags,omitempty"`
	Created      string             `json:"created,omitempty"`
	LastModified string             `json:"lastModified,omitempty"`
	Links        base.Links         `json:"links,omitempty"`
	Status       AuditContextStatus `json:"status,omitempty"`
	Origin       AuditContextOrigin `json:"origin,omitempty"`
}
