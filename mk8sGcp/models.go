/* auto-generated */

package mk8sGcp

import "gitlab.com/controlplane/controlplane/go-libs/schema/mk8sCommon"

type ImageRecommended string

const (
	ImageRecommendedUbuntu2004Lts ImageRecommended = "ubuntu-2004-lts"
	ImageRecommendedUbuntu2204Lts ImageRecommended = "ubuntu-2204-lts"
	ImageRecommendedUbuntu2404Lts ImageRecommended = "ubuntu-2404-lts"
)

type Image struct {
	Recommended ImageRecommended `json:"recommended,omitempty"`
	Exact       string           `json:"exact,omitempty"`
}

type GcpPoolTargetShape string

const (
	GcpPoolTargetShapeAny              GcpPoolTargetShape = "ANY"
	GcpPoolTargetShapeBalanced         GcpPoolTargetShape = "BALANCED"
	GcpPoolTargetShapeClosestToHealthy GcpPoolTargetShape = "CLOSEST_TO_HEALTHY"
)

type GcpPool struct {
	Name                  string             `json:"name,omitempty"`
	Labels                mk8sCommon.Labels  `json:"labels,omitempty"`
	Taints                mk8sCommon.Taints  `json:"taints,omitempty"`
	InstanceType          string             `json:"instanceType,omitempty"`
	OverrideImage         Image              `json:"overrideImage,omitempty"`
	BootDiskSize          float32            `json:"bootDiskSize"`
	MinSize               float32            `json:"minSize"`
	MaxSize               float32            `json:"maxSize"`
	SubnetId              string             `json:"subnetId,omitempty"`
	Zones                 []string           `json:"zones,omitempty"`
	TargetShape           GcpPoolTargetShape `json:"targetShape,omitempty"`
	ExtraSecurityGroupIds []string           `json:"extraSecurityGroupIds,omitempty"`
}

type GcpProviderRegion string

const (
	GcpProviderRegionAfricaSouth1           GcpProviderRegion = "africa-south1"
	GcpProviderRegionAsiaEast1              GcpProviderRegion = "asia-east1"
	GcpProviderRegionAsiaEast2              GcpProviderRegion = "asia-east2"
	GcpProviderRegionAsiaNortheast1         GcpProviderRegion = "asia-northeast1"
	GcpProviderRegionAsiaNortheast2         GcpProviderRegion = "asia-northeast2"
	GcpProviderRegionAsiaNortheast3         GcpProviderRegion = "asia-northeast3"
	GcpProviderRegionAsiaSouth1             GcpProviderRegion = "asia-south1"
	GcpProviderRegionAsiaSouth2             GcpProviderRegion = "asia-south2"
	GcpProviderRegionAsiaSoutheast1         GcpProviderRegion = "asia-southeast1"
	GcpProviderRegionAsiaSoutheast2         GcpProviderRegion = "asia-southeast2"
	GcpProviderRegionAustraliaSoutheast1    GcpProviderRegion = "australia-southeast1"
	GcpProviderRegionAustraliaSoutheast2    GcpProviderRegion = "australia-southeast2"
	GcpProviderRegionEuropeCentral2         GcpProviderRegion = "europe-central2"
	GcpProviderRegionEuropeNorth1           GcpProviderRegion = "europe-north1"
	GcpProviderRegionEuropeSouthwest1       GcpProviderRegion = "europe-southwest1"
	GcpProviderRegionEuropeWest1            GcpProviderRegion = "europe-west1"
	GcpProviderRegionEuropeWest10           GcpProviderRegion = "europe-west10"
	GcpProviderRegionEuropeWest12           GcpProviderRegion = "europe-west12"
	GcpProviderRegionEuropeWest2            GcpProviderRegion = "europe-west2"
	GcpProviderRegionEuropeWest3            GcpProviderRegion = "europe-west3"
	GcpProviderRegionEuropeWest4            GcpProviderRegion = "europe-west4"
	GcpProviderRegionEuropeWest6            GcpProviderRegion = "europe-west6"
	GcpProviderRegionEuropeWest8            GcpProviderRegion = "europe-west8"
	GcpProviderRegionEuropeWest9            GcpProviderRegion = "europe-west9"
	GcpProviderRegionMeCentral1             GcpProviderRegion = "me-central1"
	GcpProviderRegionMeCentral2             GcpProviderRegion = "me-central2"
	GcpProviderRegionMeWest1                GcpProviderRegion = "me-west1"
	GcpProviderRegionNorthamericaNortheast1 GcpProviderRegion = "northamerica-northeast1"
	GcpProviderRegionNorthamericaNortheast2 GcpProviderRegion = "northamerica-northeast2"
	GcpProviderRegionSouthamericaEast1      GcpProviderRegion = "southamerica-east1"
	GcpProviderRegionSouthamericaWest1      GcpProviderRegion = "southamerica-west1"
	GcpProviderRegionUsCentral1             GcpProviderRegion = "us-central1"
	GcpProviderRegionUsEast1                GcpProviderRegion = "us-east1"
	GcpProviderRegionUsEast4                GcpProviderRegion = "us-east4"
	GcpProviderRegionUsEast5                GcpProviderRegion = "us-east5"
	GcpProviderRegionUsSouth1               GcpProviderRegion = "us-south1"
	GcpProviderRegionUsWest1                GcpProviderRegion = "us-west1"
	GcpProviderRegionUsWest2                GcpProviderRegion = "us-west2"
	GcpProviderRegionUsWest3                GcpProviderRegion = "us-west3"
	GcpProviderRegionUsWest4                GcpProviderRegion = "us-west4"
)

type GcpProviderGcpTags map[string]string

type GcpProviderNetworkingServiceNetwork string

const (
	GcpProviderNetworkingServiceNetwork10430016   GcpProviderNetworkingServiceNetwork = "10.43.0.0/16"
	GcpProviderNetworkingServiceNetwork1921680016 GcpProviderNetworkingServiceNetwork = "192.168.0.0/16"
)

type GcpProviderNetworkingPodNetwork string

const (
	GcpProviderNetworkingPodNetworkVpc       GcpProviderNetworkingPodNetwork = "vpc"
	GcpProviderNetworkingPodNetwork10420016  GcpProviderNetworkingPodNetwork = "10.42.0.0/16"
	GcpProviderNetworkingPodNetwork172160015 GcpProviderNetworkingPodNetwork = "172.16.0.0/15"
	GcpProviderNetworkingPodNetwork172180015 GcpProviderNetworkingPodNetwork = "172.18.0.0/15"
	GcpProviderNetworkingPodNetwork172200015 GcpProviderNetworkingPodNetwork = "172.20.0.0/15"
	GcpProviderNetworkingPodNetwork172220015 GcpProviderNetworkingPodNetwork = "172.22.0.0/15"
	GcpProviderNetworkingPodNetwork172240015 GcpProviderNetworkingPodNetwork = "172.24.0.0/15"
	GcpProviderNetworkingPodNetwork172260015 GcpProviderNetworkingPodNetwork = "172.26.0.0/15"
	GcpProviderNetworkingPodNetwork172280015 GcpProviderNetworkingPodNetwork = "172.28.0.0/15"
	GcpProviderNetworkingPodNetwork172300015 GcpProviderNetworkingPodNetwork = "172.30.0.0/15"
)

type GcpProviderNetworking struct {
	ServiceNetwork GcpProviderNetworkingServiceNetwork `json:"serviceNetwork,omitempty"`
	PodNetwork     GcpProviderNetworkingPodNetwork     `json:"podNetwork,omitempty"`
}

type GcpProviderImageRecommended string

const (
	GcpProviderImageRecommendedUbuntu2004Lts GcpProviderImageRecommended = "ubuntu-2004-lts"
	GcpProviderImageRecommendedUbuntu2204Lts GcpProviderImageRecommended = "ubuntu-2204-lts"
	GcpProviderImageRecommendedUbuntu2404Lts GcpProviderImageRecommended = "ubuntu-2404-lts"
)

type GcpProviderImage struct {
	Recommended GcpProviderImageRecommended `json:"recommended,omitempty"`
	Exact       string                      `json:"exact,omitempty"`
}

type GcpProvider struct {
	Region               GcpProviderRegion           `json:"region,omitempty"`
	GcpTags              GcpProviderGcpTags          `json:"gcpTags,omitempty"`
	SkipCreateRoles      bool                        `json:"skipCreateRoles,omitempty"`
	Networking           GcpProviderNetworking       `json:"networking,omitempty"`
	PreInstallScript     string                      `json:"preInstallScript,omitempty"`
	Image                GcpProviderImage            `json:"image,omitempty"`
	DeployServiceAccount string                      `json:"deployServiceAccount,omitempty"`
	VpcId                string                      `json:"vpcId,omitempty"`
	ProjectId            string                      `json:"projectId,omitempty"`
	KeyPair              string                      `json:"keyPair,omitempty"`
	DiskEncryptionKey    string                      `json:"diskEncryptionKey,omitempty"`
	NodePools            []GcpPool                   `json:"nodePools,omitempty"`
	Autoscaler           mk8sCommon.AutoscalerConfig `json:"autoscaler,omitempty"`
}

type GcpProviderStatus map[string]any
