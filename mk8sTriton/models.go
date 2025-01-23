/* auto-generated */

package mk8sTriton

import "gitlab.com/controlplane/controlplane/go-libs/schema/mk8sCommon"

type TritonPoolTritonTags map[string]string

type TritonPool struct {
	Name              string               `json:"name,omitempty"`
	Labels            mk8sCommon.Labels    `json:"labels,omitempty"`
	Taints            mk8sCommon.Taints    `json:"taints,omitempty"`
	PackageId         string               `json:"packageId,omitempty"`
	OverrideImageId   string               `json:"overrideImageId,omitempty"`
	PublicNetworkId   string               `json:"publicNetworkId,omitempty"`
	PrivateNetworkIds []string             `json:"privateNetworkIds,omitempty"`
	TritonTags        TritonPoolTritonTags `json:"tritonTags,omitempty"`
	MinSize           float32              `json:"minSize"`
	MaxSize           float32              `json:"maxSize"`
}

type ManualMetadata map[string]string

type ManualTags map[string]string

type Manual struct {
	PackageId         string         `json:"packageId,omitempty"`
	ImageId           string         `json:"imageId,omitempty"`
	PublicNetworkId   string         `json:"publicNetworkId,omitempty"`
	PrivateNetworkIds []string       `json:"privateNetworkIds,omitempty"`
	Metadata          ManualMetadata `json:"metadata,omitempty"`
	Tags              ManualTags     `json:"tags,omitempty"`
	Count             float32        `json:"count"`
	CnsInternalDomain string         `json:"cnsInternalDomain,omitempty"`
	CnsPublicDomain   string         `json:"cnsPublicDomain,omitempty"`
}

type LoadBalancerConfigGateway struct {
}

type LoadBalancerConfig struct {
	Manual  Manual                    `json:"manual,omitempty"`
	Gateway LoadBalancerConfigGateway `json:"gateway,omitempty"`
}

type SdcConnection struct {
	Url                  string `json:"url,omitempty"`
	Account              string `json:"account,omitempty"`
	User                 string `json:"user,omitempty"`
	PrivateKeySecretLink string `json:"privateKeySecretLink,omitempty"`
}

type TritonProviderConnection struct {
	Url                  string `json:"url,omitempty"`
	Account              string `json:"account,omitempty"`
	User                 string `json:"user,omitempty"`
	PrivateKeySecretLink string `json:"privateKeySecretLink,omitempty"`
}

type TritonProviderNetworkingServiceNetwork string

const (
	TritonProviderNetworkingServiceNetwork10430016   TritonProviderNetworkingServiceNetwork = "10.43.0.0/16"
	TritonProviderNetworkingServiceNetwork1921680016 TritonProviderNetworkingServiceNetwork = "192.168.0.0/16"
)

type TritonProviderNetworkingPodNetwork string

const (
	TritonProviderNetworkingPodNetwork10420016  TritonProviderNetworkingPodNetwork = "10.42.0.0/16"
	TritonProviderNetworkingPodNetwork172160015 TritonProviderNetworkingPodNetwork = "172.16.0.0/15"
	TritonProviderNetworkingPodNetwork172180015 TritonProviderNetworkingPodNetwork = "172.18.0.0/15"
	TritonProviderNetworkingPodNetwork172200015 TritonProviderNetworkingPodNetwork = "172.20.0.0/15"
	TritonProviderNetworkingPodNetwork172220015 TritonProviderNetworkingPodNetwork = "172.22.0.0/15"
	TritonProviderNetworkingPodNetwork172240015 TritonProviderNetworkingPodNetwork = "172.24.0.0/15"
	TritonProviderNetworkingPodNetwork172260015 TritonProviderNetworkingPodNetwork = "172.26.0.0/15"
	TritonProviderNetworkingPodNetwork172280015 TritonProviderNetworkingPodNetwork = "172.28.0.0/15"
	TritonProviderNetworkingPodNetwork172300015 TritonProviderNetworkingPodNetwork = "172.30.0.0/15"
)

type TritonProviderNetworking struct {
	ServiceNetwork TritonProviderNetworkingServiceNetwork `json:"serviceNetwork,omitempty"`
	PodNetwork     TritonProviderNetworkingPodNetwork     `json:"podNetwork,omitempty"`
}

type TritonProviderLocation string

const (
	TritonProviderLocationAwsEuCentral1 TritonProviderLocation = "aws-eu-central-1"
	TritonProviderLocationAwsUsEast2    TritonProviderLocation = "aws-us-east-2"
	TritonProviderLocationAwsUsWest2    TritonProviderLocation = "aws-us-west-2"
	TritonProviderLocationGcpMeWest1    TritonProviderLocation = "gcp-me-west1"
	TritonProviderLocationGcpUsEast1    TritonProviderLocation = "gcp-us-east1"
)

type TritonProvider struct {
	Connection       TritonProviderConnection    `json:"connection,omitempty"`
	Networking       TritonProviderNetworking    `json:"networking,omitempty"`
	PreInstallScript mk8sCommon.PreInstallScript `json:"preInstallScript,omitempty"`
	Location         TritonProviderLocation      `json:"location,omitempty"`
	LoadBalancer     LoadBalancerConfig          `json:"loadBalancer,omitempty"`
	PrivateNetworkId string                      `json:"privateNetworkId,omitempty"`
	FirewallEnabled  bool                        `json:"firewallEnabled,omitempty"`
	NodePools        []TritonPool                `json:"nodePools,omitempty"`
	ImageId          string                      `json:"imageId,omitempty"`
	SshKeys          []mk8sCommon.SshPublicKey   `json:"sshKeys,omitempty"`
	Autoscaler       mk8sCommon.AutoscalerConfig `json:"autoscaler,omitempty"`
}

type TritonProviderStatus map[string]any

type TritonJoinParams struct {
	NodePoolName string `json:"nodePoolName,omitempty"`

	/* WARNING!! Arbitrary properties are being ignored! */
}
