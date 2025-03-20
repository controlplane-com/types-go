/* auto-generated */

package mk8sAddons

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
