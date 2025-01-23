/* auto-generated */

package repositorySync

type RepositorySyncSourceDirectory struct {
	Recurse bool `json:"recurse,omitempty"`
}

type RepositorySyncSource struct {
	RepoURL           string                        `json:"repoURL,omitempty"`
	TargetRevision    string                        `json:"targetRevision,omitempty"`
	Path              string                        `json:"path,omitempty"`
	AuthSecretLink    string                        `json:"authSecretLink,omitempty"`
	WebhookSecretLink string                        `json:"webhookSecretLink,omitempty"`
	Directory         RepositorySyncSourceDirectory `json:"directory,omitempty"`
}

type RepositorySync struct {
	Source RepositorySyncSource `json:"source,omitempty"`
}
