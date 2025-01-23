/* auto-generated */

package cronjob

import "github.com/controlplane-com/types-go/containerstatus"

type JobExecutionConditionStatus string

const (
	JobExecutionConditionStatusTrue    JobExecutionConditionStatus = "True"
	JobExecutionConditionStatusFalse   JobExecutionConditionStatus = "False"
	JobExecutionConditionStatusUnknown JobExecutionConditionStatus = "Unknown"
)

type JobExecutionConditionType string

const (
	JobExecutionConditionTypeComplete  JobExecutionConditionType = "Complete"
	JobExecutionConditionTypeFailed    JobExecutionConditionType = "Failed"
	JobExecutionConditionTypeSuspended JobExecutionConditionType = "Suspended"
)

type JobExecutionCondition struct {
	Status             JobExecutionConditionStatus `json:"status,omitempty"`
	Type               JobExecutionConditionType   `json:"type,omitempty"`
	LastDetectionTime  string                      `json:"lastDetectionTime,omitempty"`
	LastTransitionTime string                      `json:"lastTransitionTime,omitempty"`
	Message            string                      `json:"message,omitempty"`
	Reason             string                      `json:"reason,omitempty"`
}

type JobExecutionStatusStatus string

const (
	JobExecutionStatusStatusSuccessful JobExecutionStatusStatus = "successful"
	JobExecutionStatusStatusFailed     JobExecutionStatusStatus = "failed"
	JobExecutionStatusStatusActive     JobExecutionStatusStatus = "active"
	JobExecutionStatusStatusPending    JobExecutionStatusStatus = "pending"
	JobExecutionStatusStatusInvalid    JobExecutionStatusStatus = "invalid"
	JobExecutionStatusStatusRemoved    JobExecutionStatusStatus = "removed"
	JobExecutionStatusStatusEmpty      JobExecutionStatusStatus = ""
)

type JobExecutionStatusContainers map[string]containerstatus.ContainerStatus

type JobExecutionStatus struct {
	WorkloadVersion float32                      `json:"workloadVersion"`
	Status          JobExecutionStatusStatus     `json:"status,omitempty"`
	StartTime       string                       `json:"startTime,omitempty"`
	CompletionTime  string                       `json:"completionTime,omitempty"`
	Conditions      []JobExecutionCondition      `json:"conditions,omitempty"`
	Name            string                       `json:"name,omitempty"`
	Replica         string                       `json:"replica,omitempty"`
	Containers      JobExecutionStatusContainers `json:"containers,omitempty"`
}
