/* auto-generated */

package image

import "gitlab.com/controlplane/controlplane/go-libs/schema/base"

type ImageManifest map[string]any

type Image struct {
	Id           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Kind         base.Kind     `json:"kind,omitempty"`
	Version      float32       `json:"version"`
	Created      string        `json:"created,omitempty"`
	LastModified string        `json:"lastModified,omitempty"`
	Tags         base.Tags     `json:"tags,omitempty"`
	Links        base.Links    `json:"links,omitempty"`
	Repository   string        `json:"repository,omitempty"`
	Tag          string        `json:"tag,omitempty"`
	Manifest     ImageManifest `json:"manifest,omitempty"`
	Digest       string        `json:"digest,omitempty"`
}

type ImageSummaryRepositories struct {
	Name      string  `json:"name,omitempty"`
	TagsCount float32 `json:"tagsCount"`
}

type ImageSummary struct {
	Kind         base.Kind                  `json:"kind,omitempty"`
	Links        base.Links                 `json:"links,omitempty"`
	Repositories []ImageSummaryRepositories `json:"repositories,omitempty"`
}
