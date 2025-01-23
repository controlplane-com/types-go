/* auto-generated */

package org

import "gitlab.com/controlplane/controlplane/go-libs/schema/orgLogging"
import "gitlab.com/controlplane/controlplane/go-libs/schema/tracing"
import "gitlab.com/controlplane/controlplane/go-libs/schema/base"

type OrgStatus struct {
	AccountLink string `json:"accountLink,omitempty"`
	Active      bool   `json:"active,omitempty"`
}

type AuthConfig struct {
	DomainAutoMembers []string `json:"domainAutoMembers,omitempty"`
	SamlOnly          bool     `json:"samlOnly,omitempty"`
}

type ObservabilityConfig struct {
	LogsRetentionDays    float32  `json:"logsRetentionDays"`
	MetricsRetentionDays float32  `json:"metricsRetentionDays"`
	TracesRetentionDays  float32  `json:"tracesRetentionDays"`
	DefaultAlertEmail    []string `json:"defaultAlertEmail,omitempty"`
}

type ThreatDetectionMinimumSeverity string

const (
	ThreatDetectionMinimumSeverityWarning  ThreatDetectionMinimumSeverity = "warning"
	ThreatDetectionMinimumSeverityError    ThreatDetectionMinimumSeverity = "error"
	ThreatDetectionMinimumSeverityCritical ThreatDetectionMinimumSeverity = "critical"
)

type ThreatDetectionSyslogTransport string

const (
	ThreatDetectionSyslogTransportTcp ThreatDetectionSyslogTransport = "tcp"
	ThreatDetectionSyslogTransportUdp ThreatDetectionSyslogTransport = "udp"
)

type ThreatDetectionSyslog struct {
	Transport ThreatDetectionSyslogTransport `json:"transport,omitempty"`
	Host      string                         `json:"host,omitempty"`
	Port      float32                        `json:"port"`
}

type ThreatDetection struct {
	Enabled         bool                           `json:"enabled,omitempty"`
	MinimumSeverity ThreatDetectionMinimumSeverity `json:"minimumSeverity,omitempty"`
	Syslog          ThreatDetectionSyslog          `json:"syslog,omitempty"`
}

type OrgSpecLogging struct {
	S3          orgLogging.S3Logging          `json:"s3,omitempty"`
	Coralogix   orgLogging.CoralogixLogging   `json:"coralogix,omitempty"`
	Datadog     orgLogging.DatadogLogging     `json:"datadog,omitempty"`
	Logzio      orgLogging.LogzioLogging      `json:"logzio,omitempty"`
	Elastic     orgLogging.ElasticLogging     `json:"elastic,omitempty"`
	CloudWatch  orgLogging.CloudWatchLogging  `json:"cloudWatch,omitempty"`
	Fluentd     orgLogging.FluentdLogging     `json:"fluentd,omitempty"`
	Stackdriver orgLogging.StackdriverLogging `json:"stackdriver,omitempty"`
	Syslog      orgLogging.SyslogLogging      `json:"syslog,omitempty"`
}

type OrgSpecExtraLogging struct {
	S3          orgLogging.S3Logging          `json:"s3,omitempty"`
	Coralogix   orgLogging.CoralogixLogging   `json:"coralogix,omitempty"`
	Datadog     orgLogging.DatadogLogging     `json:"datadog,omitempty"`
	Logzio      orgLogging.LogzioLogging      `json:"logzio,omitempty"`
	Elastic     orgLogging.ElasticLogging     `json:"elastic,omitempty"`
	CloudWatch  orgLogging.CloudWatchLogging  `json:"cloudWatch,omitempty"`
	Fluentd     orgLogging.FluentdLogging     `json:"fluentd,omitempty"`
	Stackdriver orgLogging.StackdriverLogging `json:"stackdriver,omitempty"`
	Syslog      orgLogging.SyslogLogging      `json:"syslog,omitempty"`
}

type OrgSpecObservability struct {
	LogsRetentionDays    float32  `json:"logsRetentionDays"`
	MetricsRetentionDays float32  `json:"metricsRetentionDays"`
	TracesRetentionDays  float32  `json:"tracesRetentionDays"`
	DefaultAlertEmail    []string `json:"defaultAlertEmail,omitempty"`
}

type OrgSpecSecurity struct {
	ThreatDetection ThreatDetection `json:"threatDetection,omitempty"`
}

type OrgSpec struct {
	Logging               OrgSpecLogging        `json:"logging,omitempty"`
	ExtraLogging          []OrgSpecExtraLogging `json:"extraLogging,omitempty"`
	Tracing               tracing.Tracing       `json:"tracing,omitempty"`
	SessionTimeoutSeconds float32               `json:"sessionTimeoutSeconds"`
	AuthConfig            AuthConfig            `json:"authConfig,omitempty"`
	Observability         OrgSpecObservability  `json:"observability,omitempty"`
	Security              OrgSpecSecurity       `json:"security,omitempty"`
}

type Org struct {
	Id           string     `json:"id,omitempty"`
	Kind         base.Kind  `json:"kind,omitempty"`
	Version      float32    `json:"version"`
	Description  string     `json:"description,omitempty"`
	Tags         base.Tags  `json:"tags,omitempty"`
	Created      string     `json:"created,omitempty"`
	LastModified string     `json:"lastModified,omitempty"`
	Links        base.Links `json:"links,omitempty"`
	Name         string     `json:"name,omitempty"`
	Spec         OrgSpec    `json:"spec,omitempty"`
	Status       OrgStatus  `json:"status,omitempty"`
}

type QuotaOverride struct {
	Name string  `json:"name,omitempty"`
	Max  float32 `json:"max"`
}

type OrgConfig struct {
	AwsPrivateLinks    []string        `json:"awsPrivateLinks,omitempty"`
	GcpServiceConnects []string        `json:"gcpServiceConnects,omitempty"`
	QuotaOverrides     []QuotaOverride `json:"quotaOverrides,omitempty"`
}
