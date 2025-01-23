/* auto-generated */

package deployment

import "github.com/controlplane-com/types-go/pkg/containerstatus"
import "github.com/controlplane-com/types-go/pkg/cronjob"
import "github.com/controlplane-com/types-go/pkg/base"

type DeploymentVersionContainers map[string]containerstatus.ContainerStatus

type DeploymentVersion struct {
	Name       string                      `json:"name,omitempty"`
	Created    string                      `json:"created,omitempty"`
	Workload   float32                     `json:"workload"`
	Gvc        float32                     `json:"gvc"`
	Containers DeploymentVersionContainers `json:"containers,omitempty"`
	Ready      bool                        `json:"ready,omitempty"`
	Message    string                      `json:"message,omitempty"`
}

type DeploymentStatus struct {
	Endpoint                  string                       `json:"endpoint,omitempty"`
	Remote                    string                       `json:"remote,omitempty"`
	LastProcessedVersion      float32                      `json:"lastProcessedVersion"`
	ExpectedDeploymentVersion float32                      `json:"expectedDeploymentVersion"`
	Internal                  any                          `json:"internal,omitempty"`
	Ready                     bool                         `json:"ready,omitempty"`
	Deploying                 bool                         `json:"deploying,omitempty"`
	Versions                  []DeploymentVersion          `json:"versions,omitempty"`
	JobExecutions             []cronjob.JobExecutionStatus `json:"jobExecutions,omitempty"`
	Message                   string                       `json:"message,omitempty"`
}

type Deployment struct {
	Name         string           `json:"name,omitempty"`
	Kind         base.Kind        `json:"kind,omitempty"`
	Links        base.Links       `json:"links,omitempty"`
	LastModified string           `json:"lastModified,omitempty"`
	Status       DeploymentStatus `json:"status,omitempty"`
}
