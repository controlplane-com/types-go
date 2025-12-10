/* auto-generated */

package volumeSet

import "github.com/controlplane-com/types-go/pkg/base"

type FileSystem struct {
	Name              string   `json:"name,omitempty"`
	AccessMode        string   `json:"accessMode,omitempty"`
	CommandsSupported []string `json:"commandsSupported,omitempty"`
}

type PerformanceClassFeaturesSupported string

const (
	PerformanceClassFeaturesSupportedAutomaticExpansion PerformanceClassFeaturesSupported = "automatic-expansion"
	PerformanceClassFeaturesSupportedSnapshots          PerformanceClassFeaturesSupported = "snapshots"
)

type PerformanceClass struct {
	Name              string                              `json:"name,omitempty"`
	MinCapacity       float32                             `json:"minCapacity"`
	MaxCapacity       float32                             `json:"maxCapacity"`
	FeaturesSupported []PerformanceClassFeaturesSupported `json:"featuresSupported,omitempty"`
}

type VolumeSnapshotTags map[string]string

type VolumeSnapshot struct {
	Name    string               `json:"name,omitempty"`
	Id      string               `json:"id,omitempty"`
	Created string               `json:"created,omitempty"`
	Expires string               `json:"expires,omitempty"`
	Size    float32              `json:"size"`
	Tags    []VolumeSnapshotTags `json:"tags,omitempty"`
}

type PersistentVolumeStatusLifecycle string

const (
	PersistentVolumeStatusLifecycleCreating  PersistentVolumeStatusLifecycle = "creating"
	PersistentVolumeStatusLifecycleUnused    PersistentVolumeStatusLifecycle = "unused"
	PersistentVolumeStatusLifecycleUnbound   PersistentVolumeStatusLifecycle = "unbound"
	PersistentVolumeStatusLifecycleBound     PersistentVolumeStatusLifecycle = "bound"
	PersistentVolumeStatusLifecycleDeleted   PersistentVolumeStatusLifecycle = "deleted"
	PersistentVolumeStatusLifecycleRepairing PersistentVolumeStatusLifecycle = "repairing"
)

type PersistentVolumeStatusAttributes map[string]string

type PersistentVolumeStatus struct {
	Lifecycle           PersistentVolumeStatusLifecycle  `json:"lifecycle,omitempty"`
	StorageDeviceId     string                           `json:"storageDeviceId,omitempty"`
	OldStorageDeviceIds []string                         `json:"oldStorageDeviceIds,omitempty"`
	ResourceName        string                           `json:"resourceName,omitempty"`
	Index               float32                          `json:"index"`
	CurrentSize         float32                          `json:"currentSize"`
	CurrentBytesUsed    float32                          `json:"currentBytesUsed"`
	Iops                float32                          `json:"iops"`
	Throughput          float32                          `json:"throughput"`
	Driver              string                           `json:"driver,omitempty"`
	VolumeSnapshots     []VolumeSnapshot                 `json:"volumeSnapshots,omitempty"`
	Attributes          PersistentVolumeStatusAttributes `json:"attributes,omitempty"`
	Zone                string                           `json:"zone,omitempty"`
}

type VolumeSetStatusLocation struct {
	Name               string                   `json:"name,omitempty"`
	Volumes            []PersistentVolumeStatus `json:"volumes,omitempty"`
	DesiredVolumeCount float32                  `json:"desiredVolumeCount"`
	ClusterId          string                   `json:"clusterId,omitempty"`
}

type VolumeSetStatus struct {
	ParentId       string                    `json:"parentId,omitempty"`
	UsedByWorkload string                    `json:"usedByWorkload,omitempty"`
	WorkloadLinks  []string                  `json:"workloadLinks,omitempty"`
	BindingId      string                    `json:"bindingId,omitempty"`
	Locations      []VolumeSetStatusLocation `json:"locations,omitempty"`
}

type SnapshotSpec struct {
	CreateFinalSnapshot bool   `json:"createFinalSnapshot,omitempty"`
	RetentionDuration   string `json:"retentionDuration,omitempty"`
	Schedule            string `json:"schedule,omitempty"`
}

type PerformanceClassName string

const (
	PerformanceClassNameGeneralPurposeSsd PerformanceClassName = "general-purpose-ssd"
	PerformanceClassNameHighThroughputSsd PerformanceClassName = "high-throughput-ssd"
	PerformanceClassNameShared            PerformanceClassName = "shared"
)

type FileSystemType string

const (
	FileSystemTypeExt4   FileSystemType = "ext4"
	FileSystemTypeXfs    FileSystemType = "xfs"
	FileSystemTypeShared FileSystemType = "shared"
)

type MountResources struct {
	MaxCpu    string `json:"maxCpu,omitempty"`
	MinCpu    string `json:"minCpu,omitempty"`
	MinMemory string `json:"minMemory,omitempty"`
	MaxMemory string `json:"maxMemory,omitempty"`
}

type CustomEncryptionRegion struct {
	KeyId string `json:"keyId,omitempty"`
}

type VolumeSetSpecCustomEncryptionRegions map[string]CustomEncryptionRegion

type VolumeSetSpecCustomEncryption struct {
	Regions VolumeSetSpecCustomEncryptionRegions `json:"regions,omitempty"`
}

type VolumeSetSpecAutoscaling struct {
	MaxCapacity       float32 `json:"maxCapacity"`
	MinFreePercentage float32 `json:"minFreePercentage"`
	ScalingFactor     float32 `json:"scalingFactor"`
}

type VolumeSetSpecMountOptionsResources struct {
	MaxCpu    string `json:"maxCpu,omitempty"`
	MinCpu    string `json:"minCpu,omitempty"`
	MinMemory string `json:"minMemory,omitempty"`
	MaxMemory string `json:"maxMemory,omitempty"`
}

type VolumeSetSpecMountOptions struct {
	Resources VolumeSetSpecMountOptionsResources `json:"resources,omitempty"`
}

type VolumeSetSpec struct {
	InitialCapacity    float32                       `json:"initialCapacity"`
	PerformanceClass   PerformanceClassName          `json:"performanceClass,omitempty"`
	StorageClassSuffix string                        `json:"storageClassSuffix,omitempty"`
	FileSystemType     FileSystemType                `json:"fileSystemType,omitempty"`
	CustomEncryption   VolumeSetSpecCustomEncryption `json:"customEncryption,omitempty"`
	Snapshots          SnapshotSpec                  `json:"snapshots,omitempty"`
	Autoscaling        VolumeSetSpecAutoscaling      `json:"autoscaling,omitempty"`
	MountOptions       VolumeSetSpecMountOptions     `json:"mountOptions,omitempty"`
}

type VolumeSet struct {
	Id           string          `json:"id,omitempty"`
	Name         base.Name       `json:"name,omitempty"`
	Kind         base.Kind       `json:"kind,omitempty"`
	Version      float32         `json:"version"`
	Description  string          `json:"description,omitempty"`
	Tags         base.Tags       `json:"tags,omitempty"`
	Created      string          `json:"created,omitempty"`
	LastModified string          `json:"lastModified,omitempty"`
	Links        base.Links      `json:"links,omitempty"`
	Spec         VolumeSetSpec   `json:"spec,omitempty"`
	Status       VolumeSetStatus `json:"status,omitempty"`
	Gvc          any             `json:"gvc,omitempty"`
}
