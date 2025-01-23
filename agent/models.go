/* auto-generated */

package agent

import "gitlab.com/controlplane/controlplane/go-libs/schema/base"

type BootstrapConfig struct {
	RegistrationToken string `json:"registrationToken,omitempty"`
	AgentId           string `json:"agentId,omitempty"`
	AgentLink         string `json:"agentLink,omitempty"`
	HubEndpoint       string `json:"hubEndpoint,omitempty"`
}

type AgentStatus struct {
	BootstrapConfig BootstrapConfig `json:"bootstrapConfig,omitempty"`
}

type Agent struct {
	Id           string      `json:"id,omitempty"`
	Name         base.Name   `json:"name,omitempty"`
	Kind         base.Kind   `json:"kind,omitempty"`
	Version      float32     `json:"version"`
	Description  string      `json:"description,omitempty"`
	Tags         base.Tags   `json:"tags,omitempty"`
	Created      string      `json:"created,omitempty"`
	LastModified string      `json:"lastModified,omitempty"`
	Links        base.Links  `json:"links,omitempty"`
	Status       AgentStatus `json:"status,omitempty"`
}

type AgentInfoEnv map[string]any

type AgentInfo struct {
	LastActive   string       `json:"lastActive,omitempty"`
	Env          AgentInfoEnv `json:"env,omitempty"`
	PeerCount    float32      `json:"peerCount"`
	ServiceCount float32      `json:"serviceCount"`
}
