/* auto-generated */

package workload

import "github.com/controlplane-com/types-go/port"
import "github.com/controlplane-com/types-go/env"
import "github.com/controlplane-com/types-go/volumeSpec"
import "github.com/controlplane-com/types-go/workloadOptions"
import "github.com/controlplane-com/types-go/envoy"
import "github.com/controlplane-com/types-go/base"

type Memory string

type Cpu string

type HealthCheckSpecExec struct {
	Command []string `json:"command,omitempty"`
}

type HealthCheckSpecGrpc struct {
	Port port.Port `json:"port,omitempty"`
}

type HealthCheckSpecTcpSocket struct {
	Port float32 `json:"port"`
}

type HealthCheckSpecHttpGetHttpHeaders struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type HealthCheckSpecHttpGetScheme string

const (
	HealthCheckSpecHttpGetSchemeHttp  HealthCheckSpecHttpGetScheme = "HTTP"
	HealthCheckSpecHttpGetSchemeHttps HealthCheckSpecHttpGetScheme = "HTTPS"
)

type HealthCheckSpecHttpGet struct {
	Path        string                              `json:"path,omitempty"`
	Port        float32                             `json:"port"`
	HttpHeaders []HealthCheckSpecHttpGetHttpHeaders `json:"httpHeaders,omitempty"`
	Scheme      HealthCheckSpecHttpGetScheme        `json:"scheme,omitempty"`
}

type HealthCheckSpec struct {
	Exec                HealthCheckSpecExec      `json:"exec,omitempty"`
	Grpc                HealthCheckSpecGrpc      `json:"grpc,omitempty"`
	TcpSocket           HealthCheckSpecTcpSocket `json:"tcpSocket,omitempty"`
	HttpGet             HealthCheckSpecHttpGet   `json:"httpGet,omitempty"`
	InitialDelaySeconds float32                  `json:"initialDelaySeconds"`
	PeriodSeconds       float32                  `json:"periodSeconds"`
	TimeoutSeconds      float32                  `json:"timeoutSeconds"`
	SuccessThreshold    float32                  `json:"successThreshold"`
	FailureThreshold    float32                  `json:"failureThreshold"`
}

type RolloutOptionsScalingPolicy string

const (
	RolloutOptionsScalingPolicyOrderedReady RolloutOptionsScalingPolicy = "OrderedReady"
	RolloutOptionsScalingPolicyParallel     RolloutOptionsScalingPolicy = "Parallel"
)

type RolloutOptions struct {
	MinReadySeconds        float32                     `json:"minReadySeconds"`
	MaxUnavailableReplicas string                      `json:"maxUnavailableReplicas,omitempty"`
	MaxSurgeReplicas       string                      `json:"maxSurgeReplicas,omitempty"`
	ScalingPolicy          RolloutOptionsScalingPolicy `json:"scalingPolicy,omitempty"`
}

type RolloutOptionsStatefulScalingPolicy string

const (
	RolloutOptionsStatefulScalingPolicyOrderedReady RolloutOptionsStatefulScalingPolicy = "OrderedReady"
	RolloutOptionsStatefulScalingPolicyParallel     RolloutOptionsStatefulScalingPolicy = "Parallel"
)

type RolloutOptionsStateful struct {
	MinReadySeconds        float32                             `json:"minReadySeconds"`
	MaxSurgeReplicas       string                              `json:"maxSurgeReplicas,omitempty"`
	ScalingPolicy          RolloutOptionsStatefulScalingPolicy `json:"scalingPolicy,omitempty"`
	MaxUnavailableReplicas string                              `json:"maxUnavailableReplicas,omitempty"`
}

type SecurityOptions struct {
	FilesystemGroupId float32 `json:"filesystemGroupId"`
}

type GpuResourceNvidia struct {
	Model    any     `json:"model,omitempty"`
	Quantity float32 `json:"quantity"`
}

type GpuResource struct {
	Nvidia GpuResourceNvidia `json:"nvidia,omitempty"`
}

type ContainerSpecMetrics struct {
	Port float32 `json:"port"`
	Path string  `json:"path,omitempty"`
}

type ContainerSpecPortsProtocol string

const (
	ContainerSpecPortsProtocolHttp  ContainerSpecPortsProtocol = "http"
	ContainerSpecPortsProtocolHttp2 ContainerSpecPortsProtocol = "http2"
	ContainerSpecPortsProtocolGrpc  ContainerSpecPortsProtocol = "grpc"
	ContainerSpecPortsProtocolTcp   ContainerSpecPortsProtocol = "tcp"
)

type ContainerSpecPorts struct {
	Protocol ContainerSpecPortsProtocol `json:"protocol,omitempty"`
	Number   float32                    `json:"number"`
}

type ContainerSpecReadinessProbeExec struct {
	Command []string `json:"command,omitempty"`
}

type ContainerSpecReadinessProbeGrpc struct {
	Port port.Port `json:"port,omitempty"`
}

type ContainerSpecReadinessProbeTcpSocket struct {
	Port float32 `json:"port"`
}

type ContainerSpecReadinessProbeHttpGetHttpHeaders struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ContainerSpecReadinessProbeHttpGetScheme string

const (
	ContainerSpecReadinessProbeHttpGetSchemeHttp  ContainerSpecReadinessProbeHttpGetScheme = "HTTP"
	ContainerSpecReadinessProbeHttpGetSchemeHttps ContainerSpecReadinessProbeHttpGetScheme = "HTTPS"
)

type ContainerSpecReadinessProbeHttpGet struct {
	Path        string                                          `json:"path,omitempty"`
	Port        float32                                         `json:"port"`
	HttpHeaders []ContainerSpecReadinessProbeHttpGetHttpHeaders `json:"httpHeaders,omitempty"`
	Scheme      ContainerSpecReadinessProbeHttpGetScheme        `json:"scheme,omitempty"`
}

type ContainerSpecReadinessProbe struct {
	Exec                ContainerSpecReadinessProbeExec      `json:"exec,omitempty"`
	Grpc                ContainerSpecReadinessProbeGrpc      `json:"grpc,omitempty"`
	TcpSocket           ContainerSpecReadinessProbeTcpSocket `json:"tcpSocket,omitempty"`
	HttpGet             ContainerSpecReadinessProbeHttpGet   `json:"httpGet,omitempty"`
	InitialDelaySeconds float32                              `json:"initialDelaySeconds"`
	PeriodSeconds       float32                              `json:"periodSeconds"`
	TimeoutSeconds      float32                              `json:"timeoutSeconds"`
	SuccessThreshold    float32                              `json:"successThreshold"`
	FailureThreshold    float32                              `json:"failureThreshold"`
}

type ContainerSpecLivenessProbeExec struct {
	Command []string `json:"command,omitempty"`
}

type ContainerSpecLivenessProbeGrpc struct {
	Port port.Port `json:"port,omitempty"`
}

type ContainerSpecLivenessProbeTcpSocket struct {
	Port float32 `json:"port"`
}

type ContainerSpecLivenessProbeHttpGetHttpHeaders struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ContainerSpecLivenessProbeHttpGetScheme string

const (
	ContainerSpecLivenessProbeHttpGetSchemeHttp  ContainerSpecLivenessProbeHttpGetScheme = "HTTP"
	ContainerSpecLivenessProbeHttpGetSchemeHttps ContainerSpecLivenessProbeHttpGetScheme = "HTTPS"
)

type ContainerSpecLivenessProbeHttpGet struct {
	Path        string                                         `json:"path,omitempty"`
	Port        float32                                        `json:"port"`
	HttpHeaders []ContainerSpecLivenessProbeHttpGetHttpHeaders `json:"httpHeaders,omitempty"`
	Scheme      ContainerSpecLivenessProbeHttpGetScheme        `json:"scheme,omitempty"`
}

type ContainerSpecLivenessProbe struct {
	Exec                ContainerSpecLivenessProbeExec      `json:"exec,omitempty"`
	Grpc                ContainerSpecLivenessProbeGrpc      `json:"grpc,omitempty"`
	TcpSocket           ContainerSpecLivenessProbeTcpSocket `json:"tcpSocket,omitempty"`
	HttpGet             ContainerSpecLivenessProbeHttpGet   `json:"httpGet,omitempty"`
	InitialDelaySeconds float32                             `json:"initialDelaySeconds"`
	PeriodSeconds       float32                             `json:"periodSeconds"`
	TimeoutSeconds      float32                             `json:"timeoutSeconds"`
	SuccessThreshold    float32                             `json:"successThreshold"`
	FailureThreshold    float32                             `json:"failureThreshold"`
}

type ContainerSpecGpuNvidia struct {
	Model    any     `json:"model,omitempty"`
	Quantity float32 `json:"quantity"`
}

type ContainerSpecGpu struct {
	Nvidia ContainerSpecGpuNvidia `json:"nvidia,omitempty"`
}

type ContainerSpecLifecyclePostStartExec struct {
	Command []string `json:"command,omitempty"`
}

type ContainerSpecLifecyclePostStart struct {
	Exec ContainerSpecLifecyclePostStartExec `json:"exec,omitempty"`
}

type ContainerSpecLifecyclePreStopExec struct {
	Command []string `json:"command,omitempty"`
}

type ContainerSpecLifecyclePreStop struct {
	Exec ContainerSpecLifecyclePreStopExec `json:"exec,omitempty"`
}

type ContainerSpecLifecycle struct {
	PostStart ContainerSpecLifecyclePostStart `json:"postStart,omitempty"`
	PreStop   ContainerSpecLifecyclePreStop   `json:"preStop,omitempty"`
}

type ContainerSpec struct {
	Name           string                      `json:"name,omitempty"`
	Image          string                      `json:"image,omitempty"`
	WorkingDir     string                      `json:"workingDir,omitempty"`
	Metrics        ContainerSpecMetrics        `json:"metrics,omitempty"`
	Port           float32                     `json:"port"`
	Ports          []ContainerSpecPorts        `json:"ports,omitempty"`
	Memory         string                      `json:"memory,omitempty"`
	ReadinessProbe ContainerSpecReadinessProbe `json:"readinessProbe,omitempty"`
	LivenessProbe  ContainerSpecLivenessProbe  `json:"livenessProbe,omitempty"`
	Cpu            string                      `json:"cpu,omitempty"`
	MinCpu         string                      `json:"minCpu,omitempty"`
	MinMemory      string                      `json:"minMemory,omitempty"`
	Env            []env.EnvVar                `json:"env,omitempty"`
	Gpu            ContainerSpecGpu            `json:"gpu,omitempty"`
	InheritEnv     bool                        `json:"inheritEnv,omitempty"`
	Command        string                      `json:"command,omitempty"`
	Args           []string                    `json:"args,omitempty"`
	Lifecycle      ContainerSpecLifecycle      `json:"lifecycle,omitempty"`
	Volumes        []volumeSpec.VolumeSpec     `json:"volumes,omitempty"`
}

type HealthCheckStatus struct {
	Active      bool    `json:"active,omitempty"`
	Success     bool    `json:"success,omitempty"`
	Code        float32 `json:"code"`
	Message     string  `json:"message,omitempty"`
	Failures    float32 `json:"failures"`
	Successes   float32 `json:"successes"`
	LastChecked string  `json:"lastChecked,omitempty"`
}

type LoadBalancerStatus struct {
	Origin string `json:"origin,omitempty"`
	Url    string `json:"url,omitempty"`
}

type ResolvedImageManifestsPlatform map[string]string

type ResolvedImageManifests struct {
	Image     string                         `json:"image,omitempty"`
	MediaType string                         `json:"mediaType,omitempty"`
	Digest    string                         `json:"digest,omitempty"`
	Platform  ResolvedImageManifestsPlatform `json:"platform,omitempty"`
}

type ResolvedImage struct {
	Digest    string                   `json:"digest,omitempty"`
	Manifests []ResolvedImageManifests `json:"manifests,omitempty"`
}

type ResolvedImages struct {
	ResolvedForVersion float32         `json:"resolvedForVersion"`
	ResolvedAt         string          `json:"resolvedAt,omitempty"`
	ErrorMessages      []string        `json:"errorMessages,omitempty"`
	Images             []ResolvedImage `json:"images,omitempty"`
}

type WorkloadStatus struct {
	ParentId            string               `json:"parentId,omitempty"`
	CanonicalEndpoint   string               `json:"canonicalEndpoint,omitempty"`
	Endpoint            string               `json:"endpoint,omitempty"`
	InternalName        string               `json:"internalName,omitempty"`
	HealthCheck         HealthCheckStatus    `json:"healthCheck,omitempty"`
	CurrentReplicaCount float32              `json:"currentReplicaCount"`
	ResolvedImages      ResolvedImages       `json:"resolvedImages,omitempty"`
	LoadBalancer        []LoadBalancerStatus `json:"loadBalancer,omitempty"`

	/* WARNING!! Arbitrary properties are being ignored! */
}

type FirewallSpecExternalOutboundAllowPortProtocol string

const (
	FirewallSpecExternalOutboundAllowPortProtocolHttp  FirewallSpecExternalOutboundAllowPortProtocol = "http"
	FirewallSpecExternalOutboundAllowPortProtocolHttps FirewallSpecExternalOutboundAllowPortProtocol = "https"
	FirewallSpecExternalOutboundAllowPortProtocolTcp   FirewallSpecExternalOutboundAllowPortProtocol = "tcp"
)

type FirewallSpecExternalOutboundAllowPort struct {
	Protocol FirewallSpecExternalOutboundAllowPortProtocol `json:"protocol,omitempty"`
	Number   float32                                       `json:"number"`
}

type FirewallSpecExternal struct {
	InboundAllowCIDR      []string                                `json:"inboundAllowCIDR,omitempty"`
	InboundBlockedCIDR    []string                                `json:"inboundBlockedCIDR,omitempty"`
	OutboundAllowHostname []string                                `json:"outboundAllowHostname,omitempty"`
	OutboundAllowPort     []FirewallSpecExternalOutboundAllowPort `json:"outboundAllowPort,omitempty"`
	OutboundAllowCIDR     []string                                `json:"outboundAllowCIDR,omitempty"`
	OutboundBlockedCIDR   []string                                `json:"outboundBlockedCIDR,omitempty"`
}

type FirewallSpecInternalInboundAllowType string

const (
	FirewallSpecInternalInboundAllowTypeNone         FirewallSpecInternalInboundAllowType = "none"
	FirewallSpecInternalInboundAllowTypeSameGvc      FirewallSpecInternalInboundAllowType = "same-gvc"
	FirewallSpecInternalInboundAllowTypeSameOrg      FirewallSpecInternalInboundAllowType = "same-org"
	FirewallSpecInternalInboundAllowTypeWorkloadList FirewallSpecInternalInboundAllowType = "workload-list"
)

type FirewallSpecInternal struct {
	InboundAllowType     FirewallSpecInternalInboundAllowType `json:"inboundAllowType,omitempty"`
	InboundAllowWorkload []string                             `json:"inboundAllowWorkload,omitempty"`
}

type FirewallSpec struct {
	External FirewallSpecExternal `json:"external,omitempty"`
	Internal FirewallSpecInternal `json:"internal,omitempty"`
}

type ScheduleType string

type JobSpecConcurrencyPolicy string

const (
	JobSpecConcurrencyPolicyForbid  JobSpecConcurrencyPolicy = "Forbid"
	JobSpecConcurrencyPolicyReplace JobSpecConcurrencyPolicy = "Replace"
)

type JobSpecRestartPolicy string

const (
	JobSpecRestartPolicyOnFailure JobSpecRestartPolicy = "OnFailure"
	JobSpecRestartPolicyNever     JobSpecRestartPolicy = "Never"
)

type JobSpec struct {
	Schedule              ScheduleType             `json:"schedule,omitempty"`
	ConcurrencyPolicy     JobSpecConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`
	HistoryLimit          float32                  `json:"historyLimit"`
	RestartPolicy         JobSpecRestartPolicy     `json:"restartPolicy,omitempty"`
	ActiveDeadlineSeconds float32                  `json:"activeDeadlineSeconds"`
}

type LoadBalancerPortProtocol string

const (
	LoadBalancerPortProtocolTcp LoadBalancerPortProtocol = "TCP"
	LoadBalancerPortProtocolUdp LoadBalancerPortProtocol = "UDP"
)

type LoadBalancerPortScheme string

const (
	LoadBalancerPortSchemeHttp  LoadBalancerPortScheme = "http"
	LoadBalancerPortSchemeTcp   LoadBalancerPortScheme = "tcp"
	LoadBalancerPortSchemeHttps LoadBalancerPortScheme = "https"
	LoadBalancerPortSchemeWs    LoadBalancerPortScheme = "ws"
	LoadBalancerPortSchemeWss   LoadBalancerPortScheme = "wss"
)

type LoadBalancerPort struct {
	ExternalPort  float32                  `json:"externalPort"`
	Protocol      LoadBalancerPortProtocol `json:"protocol,omitempty"`
	Scheme        LoadBalancerPortScheme   `json:"scheme,omitempty"`
	ContainerPort float32                  `json:"containerPort"`
}

type LoadBalancerSpecDirect struct {
	Enabled bool               `json:"enabled,omitempty"`
	Ports   []LoadBalancerPort `json:"ports,omitempty"`
	IpSet   string             `json:"ipSet,omitempty"`
}

type LoadBalancerSpecGeoLocationHeaders struct {
	Asn     string `json:"asn,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
}

type LoadBalancerSpecGeoLocation struct {
	Enabled bool                               `json:"enabled,omitempty"`
	Headers LoadBalancerSpecGeoLocationHeaders `json:"headers,omitempty"`
}

type LoadBalancerSpec struct {
	Direct      LoadBalancerSpecDirect      `json:"direct,omitempty"`
	GeoLocation LoadBalancerSpecGeoLocation `json:"geoLocation,omitempty"`
}

type Extras struct {
	Affinity                  any   `json:"affinity,omitempty"`
	Tolerations               []any `json:"tolerations,omitempty"`
	TopologySpreadConstraints []any `json:"topologySpreadConstraints,omitempty"`
}

type RequestRetryPolicy struct {
	Attempts float32 `json:"attempts"`
	RetryOn  string  `json:"retryOn,omitempty"`
}

type WorkloadSpecType string

const (
	WorkloadSpecTypeServerless WorkloadSpecType = "serverless"
	WorkloadSpecTypeStandard   WorkloadSpecType = "standard"
	WorkloadSpecTypeCron       WorkloadSpecType = "cron"
	WorkloadSpecTypeStateful   WorkloadSpecType = "stateful"
)

type WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocol string

const (
	WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocolHttp  WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocol = "http"
	WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocolHttps WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocol = "https"
	WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocolTcp   WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocol = "tcp"
)

type WorkloadSpecFirewallConfigExternalOutboundAllowPort struct {
	Protocol WorkloadSpecFirewallConfigExternalOutboundAllowPortProtocol `json:"protocol,omitempty"`
	Number   float32                                                     `json:"number"`
}

type WorkloadSpecFirewallConfigExternal struct {
	InboundAllowCIDR      []string                                              `json:"inboundAllowCIDR,omitempty"`
	InboundBlockedCIDR    []string                                              `json:"inboundBlockedCIDR,omitempty"`
	OutboundAllowHostname []string                                              `json:"outboundAllowHostname,omitempty"`
	OutboundAllowPort     []WorkloadSpecFirewallConfigExternalOutboundAllowPort `json:"outboundAllowPort,omitempty"`
	OutboundAllowCIDR     []string                                              `json:"outboundAllowCIDR,omitempty"`
	OutboundBlockedCIDR   []string                                              `json:"outboundBlockedCIDR,omitempty"`
}

type WorkloadSpecFirewallConfigInternalInboundAllowType string

const (
	WorkloadSpecFirewallConfigInternalInboundAllowTypeNone         WorkloadSpecFirewallConfigInternalInboundAllowType = "none"
	WorkloadSpecFirewallConfigInternalInboundAllowTypeSameGvc      WorkloadSpecFirewallConfigInternalInboundAllowType = "same-gvc"
	WorkloadSpecFirewallConfigInternalInboundAllowTypeSameOrg      WorkloadSpecFirewallConfigInternalInboundAllowType = "same-org"
	WorkloadSpecFirewallConfigInternalInboundAllowTypeWorkloadList WorkloadSpecFirewallConfigInternalInboundAllowType = "workload-list"
)

type WorkloadSpecFirewallConfigInternal struct {
	InboundAllowType     WorkloadSpecFirewallConfigInternalInboundAllowType `json:"inboundAllowType,omitempty"`
	InboundAllowWorkload []string                                           `json:"inboundAllowWorkload,omitempty"`
}

type WorkloadSpecFirewallConfig struct {
	External WorkloadSpecFirewallConfigExternal `json:"external,omitempty"`
	Internal WorkloadSpecFirewallConfigInternal `json:"internal,omitempty"`
}

type WorkloadSpecJobConcurrencyPolicy string

const (
	WorkloadSpecJobConcurrencyPolicyForbid  WorkloadSpecJobConcurrencyPolicy = "Forbid"
	WorkloadSpecJobConcurrencyPolicyReplace WorkloadSpecJobConcurrencyPolicy = "Replace"
)

type WorkloadSpecJobRestartPolicy string

const (
	WorkloadSpecJobRestartPolicyOnFailure WorkloadSpecJobRestartPolicy = "OnFailure"
	WorkloadSpecJobRestartPolicyNever     WorkloadSpecJobRestartPolicy = "Never"
)

type WorkloadSpecJob struct {
	Schedule              ScheduleType                     `json:"schedule,omitempty"`
	ConcurrencyPolicy     WorkloadSpecJobConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`
	HistoryLimit          float32                          `json:"historyLimit"`
	RestartPolicy         WorkloadSpecJobRestartPolicy     `json:"restartPolicy,omitempty"`
	ActiveDeadlineSeconds float32                          `json:"activeDeadlineSeconds"`
}

type WorkloadSpecSidecar struct {
	Envoy envoy.EnvoyFilters `json:"envoy,omitempty"`
}

type WorkloadSpecSecurityOptions struct {
	FilesystemGroupId float32 `json:"filesystemGroupId"`
}

type WorkloadSpecLoadBalancerDirect struct {
	Enabled bool               `json:"enabled,omitempty"`
	Ports   []LoadBalancerPort `json:"ports,omitempty"`
	IpSet   string             `json:"ipSet,omitempty"`
}

type WorkloadSpecLoadBalancerGeoLocationHeaders struct {
	Asn     string `json:"asn,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
}

type WorkloadSpecLoadBalancerGeoLocation struct {
	Enabled bool                                       `json:"enabled,omitempty"`
	Headers WorkloadSpecLoadBalancerGeoLocationHeaders `json:"headers,omitempty"`
}

type WorkloadSpecLoadBalancer struct {
	Direct      WorkloadSpecLoadBalancerDirect      `json:"direct,omitempty"`
	GeoLocation WorkloadSpecLoadBalancerGeoLocation `json:"geoLocation,omitempty"`
}

type WorkloadSpecExtras struct {
	Affinity                  any   `json:"affinity,omitempty"`
	Tolerations               []any `json:"tolerations,omitempty"`
	TopologySpreadConstraints []any `json:"topologySpreadConstraints,omitempty"`
}

type WorkloadSpecRequestRetryPolicy struct {
	Attempts float32 `json:"attempts"`
	RetryOn  string  `json:"retryOn,omitempty"`
}

type WorkloadSpec struct {
	Type               WorkloadSpecType               `json:"type,omitempty"`
	IdentityLink       string                         `json:"identityLink,omitempty"`
	Containers         []ContainerSpec                `json:"containers,omitempty"`
	FirewallConfig     WorkloadSpecFirewallConfig     `json:"firewallConfig,omitempty"`
	DefaultOptions     workloadOptions.DefaultOptions `json:"defaultOptions,omitempty"`
	LocalOptions       workloadOptions.LocalOptions   `json:"localOptions,omitempty"`
	Job                WorkloadSpecJob                `json:"job,omitempty"`
	Sidecar            WorkloadSpecSidecar            `json:"sidecar,omitempty"`
	SupportDynamicTags bool                           `json:"supportDynamicTags,omitempty"`
	RolloutOptions     any                            `json:"rolloutOptions,omitempty"`
	SecurityOptions    WorkloadSpecSecurityOptions    `json:"securityOptions,omitempty"`
	LoadBalancer       WorkloadSpecLoadBalancer       `json:"loadBalancer,omitempty"`
	Extras             WorkloadSpecExtras             `json:"extras,omitempty"`
	RequestRetryPolicy WorkloadSpecRequestRetryPolicy `json:"requestRetryPolicy,omitempty"`
}

type Workload struct {
	Id           string         `json:"id,omitempty"`
	Kind         base.Kind      `json:"kind,omitempty"`
	Version      float32        `json:"version"`
	Description  string         `json:"description,omitempty"`
	Tags         base.Tags      `json:"tags,omitempty"`
	Created      string         `json:"created,omitempty"`
	LastModified string         `json:"lastModified,omitempty"`
	Links        base.Links     `json:"links,omitempty"`
	Name         string         `json:"name,omitempty"`
	Gvc          any            `json:"gvc,omitempty"`
	Spec         WorkloadSpec   `json:"spec,omitempty"`
	Status       WorkloadStatus `json:"status,omitempty"`
}

type WorkloadConfigScheduling struct {
	Fingerprint string  `json:"fingerprint,omitempty"`
	Version     float32 `json:"version"`
}

type WorkloadConfig struct {
	Scheduling    WorkloadConfigScheduling `json:"scheduling,omitempty"`
	ThinProvision float32                  `json:"thinProvision"`
}
