/* auto-generated */

package mk8sAddons

import "github.com/controlplane-com/types-go/pkg/workload"

type RegularExpression string

type FlexibleAddonConfig map[string]string

type NonCustomizableAddonConfig struct {
}

type AzureAddonConfig struct {
	TenantId string `json:"tenantId,omitempty"`
}

type AwsEFSAddonConfig struct {
	RoleArn string `json:"roleArn,omitempty"`
}

type MetricsAddonConfigScrapeAnnotated struct {
	IntervalSeconds   float32           `json:"intervalSeconds"`
	IncludeNamespaces RegularExpression `json:"includeNamespaces,omitempty"`
	ExcludeNamespaces RegularExpression `json:"excludeNamespaces,omitempty"`
	RetainLabels      RegularExpression `json:"retainLabels,omitempty"`
}

type MetricsAddonConfig struct {
	KubeState       bool                              `json:"kubeState,omitempty"`
	CoreDns         bool                              `json:"coreDns,omitempty"`
	Kubelet         bool                              `json:"kubelet,omitempty"`
	Apiserver       bool                              `json:"apiserver,omitempty"`
	NodeExporter    bool                              `json:"nodeExporter,omitempty"`
	Cadvisor        bool                              `json:"cadvisor,omitempty"`
	ScrapeAnnotated MetricsAddonConfigScrapeAnnotated `json:"scrapeAnnotated,omitempty"`
}

type RegistryMirrorConfigMirrors struct {
	Registry any/* TODO: [object Object]*/ `json:"registry,omitempty"`
	Mirrors  []string `json:"mirrors,omitempty"`
}

type RegistryMirrorConfig struct {
	Mirrors []RegistryMirrorConfigMirrors `json:"mirrors,omitempty"`
}

type MetricsAddonStatusRemoteWriteConfig map[string]any

type MetricsAddonStatus struct {
	PrometheusEndpoint string                              `json:"prometheusEndpoint,omitempty"`
	RemoteWriteConfig  MetricsAddonStatusRemoteWriteConfig `json:"remoteWriteConfig,omitempty"`
}

type LogsAddonConfig struct {
	AuditEnabled      bool              `json:"auditEnabled,omitempty"`
	IncludeNamespaces RegularExpression `json:"includeNamespaces,omitempty"`
	ExcludeNamespaces RegularExpression `json:"excludeNamespaces,omitempty"`
	Docker            bool              `json:"docker,omitempty"`
	Kubelet           bool              `json:"kubelet,omitempty"`
	Kernel            bool              `json:"kernel,omitempty"`
	Events            bool              `json:"events,omitempty"`
}

type LogsAddonStatus struct {
	LokiAddress string `json:"lokiAddress,omitempty"`
}

type DashboardAddonStatus struct {
	Url string `json:"url,omitempty"`
}

type AwsWorkloadIdentityAddonStatusOidcProviderConfig struct {
	ProviderUrl string `json:"providerUrl,omitempty"`
	Audience    string `json:"audience,omitempty"`
}

type AwsWorkloadIdentityAddonStatusTrustPolicy map[string]any

type AwsWorkloadIdentityAddonStatus struct {
	OidcProviderConfig AwsWorkloadIdentityAddonStatusOidcProviderConfig `json:"oidcProviderConfig,omitempty"`
	TrustPolicy        AwsWorkloadIdentityAddonStatusTrustPolicy        `json:"trustPolicy,omitempty"`
}

type AwsTrustPolicyConfigTrustPolicy map[string]any

type AwsTrustPolicyConfig struct {
	TrustPolicy AwsTrustPolicyConfigTrustPolicy `json:"trustPolicy,omitempty"`
}

type NvidiaAddonConfig struct {
	TaintGPUNodes bool `json:"taintGPUNodes,omitempty"`
}

type AwsECRAddonConfig struct {
	RoleArn string `json:"roleArn,omitempty"`
}

type AwsELBAddonConfig struct {
	RoleArn string `json:"roleArn,omitempty"`
}

type AzureACRAddonConfig struct {
	ClientId string `json:"clientId,omitempty"`
}

type ByokAddonConfigConfigActuatorLogLevel string

const (
	ByokAddonConfigConfigActuatorLogLevelTrace ByokAddonConfigConfigActuatorLogLevel = "trace"
	ByokAddonConfigConfigActuatorLogLevelInfo  ByokAddonConfigConfigActuatorLogLevel = "info"
	ByokAddonConfigConfigActuatorLogLevelError ByokAddonConfigConfigActuatorLogLevel = "error"
)

type ByokAddonConfigConfigActuatorEnv map[string]string

type ByokAddonConfigConfigActuator struct {
	MinCpu    workload.Cpu                          `json:"minCpu,omitempty"`
	MaxCpu    workload.Cpu                          `json:"maxCpu,omitempty"`
	MinMemory workload.Memory                       `json:"minMemory,omitempty"`
	MaxMemory workload.Memory                       `json:"maxMemory,omitempty"`
	LogLevel  ByokAddonConfigConfigActuatorLogLevel `json:"logLevel,omitempty"`
	Env       ByokAddonConfigConfigActuatorEnv      `json:"env,omitempty"`
}

type ByokAddonConfigConfigMiddlebox struct {
	Enabled            bool    `json:"enabled,omitempty"`
	BandwidthAlertMbps float32 `json:"bandwidthAlertMbps"`
}

type ByokAddonConfigConfigCommonPdb struct {
	MaxUnavailable float32 `json:"maxUnavailable"`
}

type ByokAddonConfigConfigCommon struct {
	DeploymentReplicas float32                        `json:"deploymentReplicas"`
	Pdb                ByokAddonConfigConfigCommonPdb `json:"pdb,omitempty"`
}

type ByokAddonConfigConfigLonghorn struct {
	Replicas float32 `json:"replicas"`
}

type ByokAddonConfigConfigIngress struct {
	Cpu           workload.Cpu    `json:"cpu,omitempty"`
	Memory        workload.Memory `json:"memory,omitempty"`
	TargetPercent float32         `json:"targetPercent"`
}

type ByokAddonConfigConfigIstioIstiod struct {
	Replicas  float32         `json:"replicas"`
	MinCpu    workload.Cpu    `json:"minCpu,omitempty"`
	MaxCpu    workload.Cpu    `json:"maxCpu,omitempty"`
	MinMemory workload.Memory `json:"minMemory,omitempty"`
	MaxMemory workload.Memory `json:"maxMemory,omitempty"`
	Pdb       float32         `json:"pdb"`
}

type ByokAddonConfigConfigIstioIngressgateway struct {
	Replicas  float32         `json:"replicas"`
	MaxCpu    workload.Cpu    `json:"maxCpu,omitempty"`
	MaxMemory workload.Memory `json:"maxMemory,omitempty"`
}

type ByokAddonConfigConfigIstioSidecar struct {
	MinCpu    workload.Cpu    `json:"minCpu,omitempty"`
	MinMemory workload.Memory `json:"minMemory,omitempty"`
}

type ByokAddonConfigConfigIstio struct {
	Istiod         ByokAddonConfigConfigIstioIstiod         `json:"istiod,omitempty"`
	Ingressgateway ByokAddonConfigConfigIstioIngressgateway `json:"ingressgateway,omitempty"`
	Sidecar        ByokAddonConfigConfigIstioSidecar        `json:"sidecar,omitempty"`
}

type ByokAddonConfigConfigLogSplitter struct {
	MinCpu        workload.Cpu    `json:"minCpu,omitempty"`
	MaxCpu        workload.Cpu    `json:"maxCpu,omitempty"`
	MinMemory     workload.Memory `json:"minMemory,omitempty"`
	MaxMemory     workload.Memory `json:"maxMemory,omitempty"`
	MemBufferSize string          `json:"memBufferSize,omitempty"`
	PerPodRate    float32         `json:"perPodRate"`
}

type ByokAddonConfigConfigMonitoringKubeStateMetrics struct {
	MinMemory workload.Memory `json:"minMemory,omitempty"`
}

type ByokAddonConfigConfigMonitoringPrometheusMain struct {
	Storage workload.Memory `json:"storage,omitempty"`
}

type ByokAddonConfigConfigMonitoringPrometheus struct {
	Main ByokAddonConfigConfigMonitoringPrometheusMain `json:"main,omitempty"`
}

type ByokAddonConfigConfigMonitoring struct {
	MinMemory        workload.Memory                                 `json:"minMemory,omitempty"`
	MaxMemory        workload.Memory                                 `json:"maxMemory,omitempty"`
	KubeStateMetrics ByokAddonConfigConfigMonitoringKubeStateMetrics `json:"kubeStateMetrics,omitempty"`
	Prometheus       ByokAddonConfigConfigMonitoringPrometheus       `json:"prometheus,omitempty"`
}

type ByokAddonConfigConfigRedis struct {
	MinCpu    workload.Cpu    `json:"minCpu,omitempty"`
	MaxCpu    workload.Cpu    `json:"maxCpu,omitempty"`
	MinMemory workload.Memory `json:"minMemory,omitempty"`
	MaxMemory workload.Memory `json:"maxMemory,omitempty"`
	Storage   workload.Memory `json:"storage,omitempty"`
}

type ByokAddonConfigConfigRedisHa struct {
	MinCpu    workload.Cpu    `json:"minCpu,omitempty"`
	MaxCpu    workload.Cpu    `json:"maxCpu,omitempty"`
	MinMemory workload.Memory `json:"minMemory,omitempty"`
	MaxMemory workload.Memory `json:"maxMemory,omitempty"`
	Storage   float32         `json:"storage"`
}

type ByokAddonConfigConfigRedisSentinel struct {
	MinCpu    workload.Cpu    `json:"minCpu,omitempty"`
	MaxCpu    workload.Cpu    `json:"maxCpu,omitempty"`
	MinMemory workload.Memory `json:"minMemory,omitempty"`
	MaxMemory workload.Memory `json:"maxMemory,omitempty"`
	Storage   float32         `json:"storage"`
}

type ByokAddonConfigConfigTempoAgent struct {
	MinCpu    workload.Cpu    `json:"minCpu,omitempty"`
	MinMemory workload.Memory `json:"minMemory,omitempty"`
}

type ByokAddonConfigConfigInternalDns struct {
	MinCpu    workload.Cpu    `json:"minCpu,omitempty"`
	MaxCpu    workload.Cpu    `json:"maxCpu,omitempty"`
	MinMemory workload.Memory `json:"minMemory,omitempty"`
	MaxMemory workload.Memory `json:"maxMemory,omitempty"`
}

type ByokAddonConfigConfig struct {
	Actuator      ByokAddonConfigConfigActuator      `json:"actuator,omitempty"`
	Middlebox     ByokAddonConfigConfigMiddlebox     `json:"middlebox,omitempty"`
	Common        ByokAddonConfigConfigCommon        `json:"common,omitempty"`
	Longhorn      ByokAddonConfigConfigLonghorn      `json:"longhorn,omitempty"`
	Ingress       ByokAddonConfigConfigIngress       `json:"ingress,omitempty"`
	Istio         ByokAddonConfigConfigIstio         `json:"istio,omitempty"`
	LogSplitter   ByokAddonConfigConfigLogSplitter   `json:"logSplitter,omitempty"`
	Monitoring    ByokAddonConfigConfigMonitoring    `json:"monitoring,omitempty"`
	Redis         ByokAddonConfigConfigRedis         `json:"redis,omitempty"`
	RedisHa       ByokAddonConfigConfigRedisHa       `json:"redisHa,omitempty"`
	RedisSentinel ByokAddonConfigConfigRedisSentinel `json:"redisSentinel,omitempty"`
	TempoAgent    ByokAddonConfigConfigTempoAgent    `json:"tempoAgent,omitempty"`
	InternalDns   ByokAddonConfigConfigInternalDns   `json:"internalDns,omitempty"`

	/* WARNING!! Arbitrary properties are being ignored! */
}

type ByokAddonConfig struct {
	IgnoreUpdates bool                  `json:"ignoreUpdates,omitempty"`
	Location      string                `json:"location,omitempty"`
	Config        ByokAddonConfigConfig `json:"config,omitempty"`
}
