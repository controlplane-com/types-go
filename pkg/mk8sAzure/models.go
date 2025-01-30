/* auto-generated */

package mk8sAzure

import "github.com/controlplane-com/types-go/pkg/mk8sCommon"

type ImageRecommended string

const (
	ImageRecommendedUbuntuNoble2404 ImageRecommended = "ubuntu/noble-24.04"
	ImageRecommendedUbuntuJammy2204 ImageRecommended = "ubuntu/jammy-22.04"
	ImageRecommendedUbuntuFocal2004 ImageRecommended = "ubuntu/focal-20.04"
)

type ImageReference struct {
	Publisher string `json:"publisher,omitempty"`
	Offer     string `json:"offer,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Version   string `json:"version,omitempty"`
}

type Image struct {
	Recommended ImageRecommended `json:"recommended,omitempty"`
	Reference   ImageReference   `json:"reference,omitempty"`
}

type AzurePool struct {
	Name          string            `json:"name,omitempty"`
	Labels        mk8sCommon.Labels `json:"labels,omitempty"`
	Taints        mk8sCommon.Taints `json:"taints,omitempty"`
	Size          string            `json:"size,omitempty"`
	SubnetId      string            `json:"subnetId,omitempty"`
	Zones         []float32         `json:"zones,omitempty"`
	OverrideImage Image             `json:"overrideImage,omitempty"`
	BootDiskSize  float32           `json:"bootDiskSize"`
	MinSize       float32           `json:"minSize"`
	MaxSize       float32           `json:"maxSize"`
}

type AzureProviderLocation string

const (
	AzureProviderLocationCentralus      AzureProviderLocation = "centralus"
	AzureProviderLocationEastus2        AzureProviderLocation = "eastus2"
	AzureProviderLocationSouthcentralus AzureProviderLocation = "southcentralus"
)

type AzureProviderAzureTags map[string]string

type AzureProviderNetworkingServiceNetwork string

const (
	AzureProviderNetworkingServiceNetwork10430016   AzureProviderNetworkingServiceNetwork = "10.43.0.0/16"
	AzureProviderNetworkingServiceNetwork1921680016 AzureProviderNetworkingServiceNetwork = "192.168.0.0/16"
)

type AzureProviderNetworkingPodNetwork string

const (
	AzureProviderNetworkingPodNetwork10420016  AzureProviderNetworkingPodNetwork = "10.42.0.0/16"
	AzureProviderNetworkingPodNetwork172160015 AzureProviderNetworkingPodNetwork = "172.16.0.0/15"
	AzureProviderNetworkingPodNetwork172180015 AzureProviderNetworkingPodNetwork = "172.18.0.0/15"
	AzureProviderNetworkingPodNetwork172200015 AzureProviderNetworkingPodNetwork = "172.20.0.0/15"
	AzureProviderNetworkingPodNetwork172220015 AzureProviderNetworkingPodNetwork = "172.22.0.0/15"
	AzureProviderNetworkingPodNetwork172240015 AzureProviderNetworkingPodNetwork = "172.24.0.0/15"
	AzureProviderNetworkingPodNetwork172260015 AzureProviderNetworkingPodNetwork = "172.26.0.0/15"
	AzureProviderNetworkingPodNetwork172280015 AzureProviderNetworkingPodNetwork = "172.28.0.0/15"
	AzureProviderNetworkingPodNetwork172300015 AzureProviderNetworkingPodNetwork = "172.30.0.0/15"
)

type AzureProviderNetworking struct {
	ServiceNetwork AzureProviderNetworkingServiceNetwork `json:"serviceNetwork,omitempty"`
	PodNetwork     AzureProviderNetworkingPodNetwork     `json:"podNetwork,omitempty"`
}

type AzureProviderImageRecommended string

const (
	AzureProviderImageRecommendedUbuntuNoble2404 AzureProviderImageRecommended = "ubuntu/noble-24.04"
	AzureProviderImageRecommendedUbuntuJammy2204 AzureProviderImageRecommended = "ubuntu/jammy-22.04"
	AzureProviderImageRecommendedUbuntuFocal2004 AzureProviderImageRecommended = "ubuntu/focal-20.04"
)

type AzureProviderImageReference struct {
	Publisher string `json:"publisher,omitempty"`
	Offer     string `json:"offer,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Version   string `json:"version,omitempty"`
}

type AzureProviderImage struct {
	Recommended AzureProviderImageRecommended `json:"recommended,omitempty"`
	Reference   AzureProviderImageReference   `json:"reference,omitempty"`
}

type AzureProviderTags map[string]string

type AzureProvider struct {
	Location         AzureProviderLocation       `json:"location,omitempty"`
	AzureTags        AzureProviderAzureTags      `json:"azureTags,omitempty"`
	SubscriptionId   string                      `json:"subscriptionId,omitempty"`
	SdkSecretLink    string                      `json:"sdkSecretLink,omitempty"`
	ResourceGroup    string                      `json:"resourceGroup,omitempty"`
	Networking       AzureProviderNetworking     `json:"networking,omitempty"`
	PreInstallScript string                      `json:"preInstallScript,omitempty"`
	Image            AzureProviderImage          `json:"image,omitempty"`
	NetworkId        string                      `json:"networkId,omitempty"`
	Tags             AzureProviderTags           `json:"tags,omitempty"`
	NodePools        []AzurePool                 `json:"nodePools,omitempty"`
	Autoscaler       mk8sCommon.AutoscalerConfig `json:"autoscaler,omitempty"`
}

type AzureProviderStatus map[string]any
