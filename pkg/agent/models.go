/* auto-generated */

package agent

import "github.com/controlplane-com/types-go/pkg/base"

type BootstrapConfig struct {
	RegistrationToken string `json:"registrationToken"`
	AgentId           string `json:"agentId"`
	AgentLink         string `json:"agentLink"`
	HubEndpoint       string `json:"hubEndpoint"`
}

type AgentStatus struct {
	BootstrapConfig BootstrapConfig `json:"bootstrapConfig,omitempty"`
}

type AgentTags map[string]any

type Agent struct {
	Id           string      `json:"id,omitempty"`
	Name         base.Name   `json:"name,omitempty"`
	Kind         base.Kind   `json:"kind,omitempty"`
	Version      float32     `json:"version"`
	Description  string      `json:"description,omitempty"`
	Tags         AgentTags   `json:"tags,omitempty"`
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
