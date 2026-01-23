/* auto-generated */

package command

import "github.com/controlplane-com/types-go/pkg/base"
import "github.com/controlplane-com/types-go/pkg/env"
import "github.com/controlplane-com/types-go/pkg/workload"
import "github.com/controlplane-com/types-go/pkg/volumeSet"

type CommandLifecycleStage string

const (
	CommandLifecycleStagePending   CommandLifecycleStage = "pending"
	CommandLifecycleStageRunning   CommandLifecycleStage = "running"
	CommandLifecycleStageCancelled CommandLifecycleStage = "cancelled"
	CommandLifecycleStageCompleted CommandLifecycleStage = "completed"
	CommandLifecycleStageFailed    CommandLifecycleStage = "failed"
)

type CommandSpec map[string]any

type CommandStatus map[string]any

type Command struct {
	Id             string                `json:"id,omitempty"`
	OwnerId        string                `json:"ownerId,omitempty"`
	Kind           base.Kind             `json:"kind,omitempty"`
	Version        float32               `json:"version"`
	Created        string                `json:"created,omitempty"`
	LastModified   string                `json:"lastModified,omitempty"`
	Links          base.Links            `json:"links,omitempty"`
	Tags           base.Tags             `json:"tags,omitempty"`
	Type           string                `json:"type,omitempty"`
	LifecycleStage CommandLifecycleStage `json:"lifecycleStage,omitempty"`
	Spec           CommandSpec           `json:"spec,omitempty"`
	Status         CommandStatus         `json:"status,omitempty"`
}

type CronWorkloadContainerOverrides struct {
	Name    string          `json:"name,omitempty"`
	Env     []env.EnvVar    `json:"env,omitempty"`
	Command string          `json:"command,omitempty"`
	Args    []string        `json:"args,omitempty"`
	Memory  workload.Memory `json:"memory,omitempty"`
	Cpu     workload.Cpu    `json:"cpu,omitempty"`
	Image   base.ImageLink  `json:"image,omitempty"`
}

type RunCronWorkloadSpec struct {
	Location           string                           `json:"location,omitempty"`
	ContainerOverrides []CronWorkloadContainerOverrides `json:"containerOverrides,omitempty"`
}

type RunCronWorkloadStatusClusterIdByLocation map[string]string

type RunCronWorkloadStatus struct {
	Replica                string                                   `json:"replica,omitempty"`
	ClusterIdByLocation    RunCronWorkloadStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	MinimumWorkloadVersion float32                                  `json:"minimumWorkloadVersion"`
}

type StopReplicaSpec struct {
	Location string `json:"location,omitempty"`
	Replica  string `json:"replica,omitempty"`
}

type StopReplicaStatusClusterIdByLocation map[string]string

type StopReplicaStatus struct {
	ClusterId               string                               `json:"clusterId,omitempty"`
	ClusterIdByLocation     StopReplicaStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	PodCreatedByCommandLink string                               `json:"podCreatedByCommandLink,omitempty"`
	PodId                   string                               `json:"podId,omitempty"`
}

type ReplaceVolumeSpec struct {
	Location    string  `json:"location,omitempty"`
	VolumeIndex float32 `json:"volumeIndex"`
}

type RestoreVolumeSpec struct {
	VolumeIndex  float32 `json:"volumeIndex"`
	Location     string  `json:"location,omitempty"`
	SnapshotName string  `json:"snapshotName,omitempty"`
	Zone         string  `json:"zone,omitempty"`
}

type ReplaceVolumeStatusStage string

const (
	ReplaceVolumeStatusStageCreateVolume               ReplaceVolumeStatusStage = "create-volume"
	ReplaceVolumeStatusStageCleanupAfterVolumeCreation ReplaceVolumeStatusStage = "cleanup-after-volume-creation"
	ReplaceVolumeStatusStageUpdateVolumeSet            ReplaceVolumeStatusStage = "update-volume-set"
	ReplaceVolumeStatusStageConfigureStorageResources  ReplaceVolumeStatusStage = "configure-storage-resources"
	ReplaceVolumeStatusStageRemoveFinalizer            ReplaceVolumeStatusStage = "remove-finalizer"
	ReplaceVolumeStatusStageShutdownReplica            ReplaceVolumeStatusStage = "shutdown-replica"
	ReplaceVolumeStatusStageAwaitReplicaTermination    ReplaceVolumeStatusStage = "await-replica-termination"
	ReplaceVolumeStatusStageCleanupK8S                 ReplaceVolumeStatusStage = "cleanup-k8s"
	ReplaceVolumeStatusStageFail                       ReplaceVolumeStatusStage = "fail"
	ReplaceVolumeStatusStageRevert                     ReplaceVolumeStatusStage = "revert"
	ReplaceVolumeStatusStageCleanupOldStorageDevice    ReplaceVolumeStatusStage = "cleanup-old-storage-device"
	ReplaceVolumeStatusStageRestartReplica             ReplaceVolumeStatusStage = "restart-replica"
)

type ReplaceVolumeStatusClusterIdByLocation map[string]string

type ReplaceVolumeStatusNewVolumeAttributes map[string]string

type ReplaceVolumeStatus struct {
	Stage                   ReplaceVolumeStatusStage               `json:"stage,omitempty"`
	Messages                []string                               `json:"messages,omitempty"`
	ClusterId               string                                 `json:"clusterId,omitempty"`
	ClusterIdByLocation     ReplaceVolumeStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	InUseByWorkloadId       string                                 `json:"inUseByWorkloadId,omitempty"`
	StorageDeviceIdToRemove string                                 `json:"storageDeviceIdToRemove,omitempty"`
	NewStorageDeviceId      string                                 `json:"newStorageDeviceId,omitempty"`
	NewVolumeAttributes     ReplaceVolumeStatusNewVolumeAttributes `json:"newVolumeAttributes,omitempty"`
	NewResourceName         string                                 `json:"newResourceName,omitempty"`
	NextVolumeSize          float32                                `json:"nextVolumeSize"`
}

type CreateVolumeSnapshotSpecSnapshotTags map[string]string

type CreateVolumeSnapshotSpec struct {
	Location               string                                 `json:"location,omitempty"`
	VolumeIndex            float32                                `json:"volumeIndex"`
	SnapshotName           string                                 `json:"snapshotName,omitempty"`
	SnapshotExpirationDate string                                 `json:"snapshotExpirationDate,omitempty"`
	SnapshotTags           []CreateVolumeSnapshotSpecSnapshotTags `json:"snapshotTags,omitempty"`
}

type CreateVolumeSnapshotStatusStage string

const (
	CreateVolumeSnapshotStatusStageCreateSnapshot  CreateVolumeSnapshotStatusStage = "create-snapshot"
	CreateVolumeSnapshotStatusStageUpdateVolumeSet CreateVolumeSnapshotStatusStage = "update-volume-set"
	CreateVolumeSnapshotStatusStageCleanupK8S      CreateVolumeSnapshotStatusStage = "cleanup-k8s"
	CreateVolumeSnapshotStatusStageRevert          CreateVolumeSnapshotStatusStage = "revert"
)

type CreateVolumeSnapshotStatusClusterIdByLocation map[string]string

type CreateVolumeSnapshotStatus struct {
	Stage               CreateVolumeSnapshotStatusStage               `json:"stage,omitempty"`
	Messages            []string                                      `json:"messages,omitempty"`
	ClusterId           string                                        `json:"clusterId,omitempty"`
	ClusterIdByLocation CreateVolumeSnapshotStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	NewSnapshotId       string                                        `json:"newSnapshotId,omitempty"`
	NewSnapshotSize     float32                                       `json:"newSnapshotSize"`
	CreationStartTime   string                                        `json:"creationStartTime,omitempty"`
}

type ExpandVolumeSpec struct {
	Location           string  `json:"location,omitempty"`
	VolumeIndex        float32 `json:"volumeIndex"`
	NewStorageCapacity float32 `json:"newStorageCapacity"`
	TimeoutSeconds     float32 `json:"timeoutSeconds"`
}

type ExpandVolumeStatusClusterIdByLocation map[string]string

type ExpandVolumeStatusStage string

const (
	ExpandVolumeStatusStageExpandVolume            ExpandVolumeStatusStage = "expand-volume"
	ExpandVolumeStatusStageDeleteStatefulSet       ExpandVolumeStatusStage = "delete-stateful-set"
	ExpandVolumeStatusStageAwaitReplicaTermination ExpandVolumeStatusStage = "await-replica-termination"
	ExpandVolumeStatusStageAwaitExpansionCompleted ExpandVolumeStatusStage = "await-expansion-completed"
	ExpandVolumeStatusStageUpdateVolumeSet         ExpandVolumeStatusStage = "update-volume-set"
	ExpandVolumeStatusStageRecreateReplica         ExpandVolumeStatusStage = "recreate-replica"
	ExpandVolumeStatusStageCleanupK8S              ExpandVolumeStatusStage = "cleanup-k8s"
	ExpandVolumeStatusStageRevert                  ExpandVolumeStatusStage = "revert"
)

type ExpandVolumeStatus struct {
	ClusterId           string                                `json:"clusterId,omitempty"`
	ClusterIdByLocation ExpandVolumeStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	Messages            []string                              `json:"messages,omitempty"`
	Stage               ExpandVolumeStatusStage               `json:"stage,omitempty"`
	ReplicaRestartedAt  string                                `json:"replicaRestartedAt,omitempty"`
	LockNames           []string                              `json:"lockNames,omitempty"`
}

type DeleteVolumeSpec struct {
	Location    string  `json:"location,omitempty"`
	VolumeIndex float32 `json:"volumeIndex"`
}

type DeleteOrphanedVolumeSpec struct {
	Location                     string  `json:"location,omitempty"`
	StorageDeviceId              string  `json:"storageDeviceId,omitempty"`
	VolumeIndex                  float32 `json:"volumeIndex"`
	NewlyObservedStorageDeviceId string  `json:"newlyObservedStorageDeviceId,omitempty"`
}

type DeleteVolumeStatusStage string

const (
	DeleteVolumeStatusStageUpdateVolumeSet         DeleteVolumeStatusStage = "update-volume-set"
	DeleteVolumeStatusStageDeleteStorageResources  DeleteVolumeStatusStage = "delete-storage-resources"
	DeleteVolumeStatusStageShutdownReplica         DeleteVolumeStatusStage = "shutdown-replica"
	DeleteVolumeStatusStageAwaitReplicaTermination DeleteVolumeStatusStage = "await-replica-termination"
	DeleteVolumeStatusStageFail                    DeleteVolumeStatusStage = "fail"
	DeleteVolumeStatusStageCleanupK8S              DeleteVolumeStatusStage = "cleanup-k8s"
)

type DeleteVolumeStatusClusterIdByLocation map[string]string

type DeleteVolumeStatus struct {
	Stage                   DeleteVolumeStatusStage               `json:"stage,omitempty"`
	ClusterId               string                                `json:"clusterId,omitempty"`
	ClusterIdByLocation     DeleteVolumeStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	Messages                []string                              `json:"messages,omitempty"`
	InUseByWorkloadId       string                                `json:"inUseByWorkloadId,omitempty"`
	StorageDeviceIdToRemove string                                `json:"storageDeviceIdToRemove,omitempty"`
}

type DeleteOrphanedVolumeStatusStage string

const (
	DeleteOrphanedVolumeStatusStageDeleteStorageResources DeleteOrphanedVolumeStatusStage = "delete-storage-resources"
	DeleteOrphanedVolumeStatusStageCleanupK8S             DeleteOrphanedVolumeStatusStage = "cleanup-k8s"
	DeleteOrphanedVolumeStatusStageUpdateVolumeSet        DeleteOrphanedVolumeStatusStage = "update-volume-set"
	DeleteOrphanedVolumeStatusStageFail                   DeleteOrphanedVolumeStatusStage = "fail"
)

type DeleteOrphanedVolumeStatusClusterIdByLocation map[string]string

type DeleteOrphanedVolumeStatus struct {
	Stage               DeleteOrphanedVolumeStatusStage               `json:"stage,omitempty"`
	ClusterId           string                                        `json:"clusterId,omitempty"`
	ClusterIdByLocation DeleteOrphanedVolumeStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	Messages            []string                                      `json:"messages,omitempty"`
}

type DeleteOrphanedVolumeSnapshotSpec struct {
	Location    string  `json:"location,omitempty"`
	SnapshotId  string  `json:"snapshotId,omitempty"`
	VolumeIndex float32 `json:"volumeIndex"`
}

type DeleteOrphanedVolumeSnapshotStatusStage string

const (
	DeleteOrphanedVolumeSnapshotStatusStageDeleteSnapshot  DeleteOrphanedVolumeSnapshotStatusStage = "delete-snapshot"
	DeleteOrphanedVolumeSnapshotStatusStageUpdateVolumeSet DeleteOrphanedVolumeSnapshotStatusStage = "update-volume-set"
	DeleteOrphanedVolumeSnapshotStatusStageFail            DeleteOrphanedVolumeSnapshotStatusStage = "fail"
)

type DeleteOrphanedVolumeSnapshotStatusClusterIdByLocation map[string]string

type DeleteOrphanedVolumeSnapshotStatus struct {
	Stage               DeleteOrphanedVolumeSnapshotStatusStage               `json:"stage,omitempty"`
	ClusterId           string                                                `json:"clusterId,omitempty"`
	ClusterIdByLocation DeleteOrphanedVolumeSnapshotStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	Messages            []string                                              `json:"messages,omitempty"`
}

type SnapshotDeletionStatusStage string

const (
	SnapshotDeletionStatusStagePending             SnapshotDeletionStatusStage = "pending"
	SnapshotDeletionStatusStageK8SResourcesCreated SnapshotDeletionStatusStage = "k8s-resources-created"
	SnapshotDeletionStatusStageDeleted             SnapshotDeletionStatusStage = "deleted"
)

type SnapshotDeletionStatus struct {
	Stage    SnapshotDeletionStatusStage `json:"stage,omitempty"`
	Messages []string                    `json:"messages,omitempty"`
}

type DeleteCloudDevicesStatusClusterIdByLocation map[string]string

type DeleteCloudDevicesStatusVolumeLifecycle string

const (
	DeleteCloudDevicesStatusVolumeLifecycleCreating  DeleteCloudDevicesStatusVolumeLifecycle = "creating"
	DeleteCloudDevicesStatusVolumeLifecycleUnused    DeleteCloudDevicesStatusVolumeLifecycle = "unused"
	DeleteCloudDevicesStatusVolumeLifecycleUnbound   DeleteCloudDevicesStatusVolumeLifecycle = "unbound"
	DeleteCloudDevicesStatusVolumeLifecycleBound     DeleteCloudDevicesStatusVolumeLifecycle = "bound"
	DeleteCloudDevicesStatusVolumeLifecycleDeleted   DeleteCloudDevicesStatusVolumeLifecycle = "deleted"
	DeleteCloudDevicesStatusVolumeLifecycleRepairing DeleteCloudDevicesStatusVolumeLifecycle = "repairing"
)

type DeleteCloudDevicesStatusVolumeAttributes map[string]string

type DeleteCloudDevicesStatusVolume struct {
	Lifecycle           DeleteCloudDevicesStatusVolumeLifecycle  `json:"lifecycle,omitempty"`
	StorageDeviceId     string                                   `json:"storageDeviceId,omitempty"`
	OldStorageDeviceIds []string                                 `json:"oldStorageDeviceIds,omitempty"`
	ResourceName        string                                   `json:"resourceName,omitempty"`
	Index               float32                                  `json:"index"`
	CurrentSize         float32                                  `json:"currentSize"`
	CurrentBytesUsed    float32                                  `json:"currentBytesUsed"`
	Iops                float32                                  `json:"iops"`
	Throughput          float32                                  `json:"throughput"`
	Driver              string                                   `json:"driver,omitempty"`
	VolumeSnapshots     []volumeSet.VolumeSnapshot               `json:"volumeSnapshots,omitempty"`
	Attributes          DeleteCloudDevicesStatusVolumeAttributes `json:"attributes,omitempty"`
	Zone                string                                   `json:"zone,omitempty"`
}

type DeleteCloudDevicesStatusStage string

const (
	DeleteCloudDevicesStatusStageDeleteSnapshots        DeleteCloudDevicesStatusStage = "delete-snapshots"
	DeleteCloudDevicesStatusStageDeleteVolume           DeleteCloudDevicesStatusStage = "delete-volume"
	DeleteCloudDevicesStatusStageFinalizeVolumeDeletion DeleteCloudDevicesStatusStage = "finalize-volume-deletion"
	DeleteCloudDevicesStatusStageUpdateVolumeSet        DeleteCloudDevicesStatusStage = "update-volume-set"
)

type DeleteCloudDevicesStatusPvcRef struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}

type DeleteCloudDevicesStatusSnapshotDeletionStatus map[string]SnapshotDeletionStatus

type DeleteCloudDevicesStatus struct {
	ClusterId              string                                         `json:"clusterId,omitempty"`
	ClusterIdByLocation    DeleteCloudDevicesStatusClusterIdByLocation    `json:"clusterIdByLocation,omitempty"`
	Volume                 DeleteCloudDevicesStatusVolume                 `json:"volume,omitempty"`
	Stage                  DeleteCloudDevicesStatusStage                  `json:"stage,omitempty"`
	Messages               []string                                       `json:"messages,omitempty"`
	PvcRef                 DeleteCloudDevicesStatusPvcRef                 `json:"pvcRef,omitempty"`
	SnapshotDeletionStatus DeleteCloudDevicesStatusSnapshotDeletionStatus `json:"snapshotDeletionStatus,omitempty"`
}

type DeleteVolumeSnapshotSpec struct {
	Location     string  `json:"location,omitempty"`
	VolumeIndex  float32 `json:"volumeIndex"`
	SnapshotName string  `json:"snapshotName,omitempty"`
}

type DeleteVolumeSnapshotStatusClusterIdByLocation map[string]string

type DeleteVolumeSnapshotStatusSnapshotTags map[string]string

type DeleteVolumeSnapshotStatusSnapshot struct {
	Name    string                                   `json:"name,omitempty"`
	Id      string                                   `json:"id,omitempty"`
	Created string                                   `json:"created,omitempty"`
	Expires string                                   `json:"expires,omitempty"`
	Size    float32                                  `json:"size"`
	Tags    []DeleteVolumeSnapshotStatusSnapshotTags `json:"tags,omitempty"`
}

type DeleteVolumeSnapshotStatusStage string

const (
	DeleteVolumeSnapshotStatusStageDeleteSnapshot  DeleteVolumeSnapshotStatusStage = "delete-snapshot"
	DeleteVolumeSnapshotStatusStageUpdateVolumeSet DeleteVolumeSnapshotStatusStage = "update-volume-set"
)

type DeleteVolumeSnapshotStatus struct {
	ClusterId           string                                        `json:"clusterId,omitempty"`
	ClusterIdByLocation DeleteVolumeSnapshotStatusClusterIdByLocation `json:"clusterIdByLocation,omitempty"`
	Snapshot            DeleteVolumeSnapshotStatusSnapshot            `json:"snapshot,omitempty"`
	Stage               DeleteVolumeSnapshotStatusStage               `json:"stage,omitempty"`
	Messages            []string                                      `json:"messages,omitempty"`
	SnapshotId          string                                        `json:"snapshotId,omitempty"`
}

type Cluster struct {
	ClusterId string `json:"clusterId,omitempty"`
	Since     string `json:"since,omitempty"`
}

type Clusters map[string]Cluster

type DeleteVolumeSetSpecVolumeSet struct {
	Id           string                    `json:"id,omitempty"`
	Name         base.Name                 `json:"name,omitempty"`
	Kind         base.Kind                 `json:"kind,omitempty"`
	Version      float32                   `json:"version"`
	Description  string                    `json:"description,omitempty"`
	Tags         base.Tags                 `json:"tags,omitempty"`
	Created      string                    `json:"created,omitempty"`
	LastModified string                    `json:"lastModified,omitempty"`
	Links        base.Links                `json:"links,omitempty"`
	Spec         volumeSet.VolumeSetSpec   `json:"spec,omitempty"`
	Status       volumeSet.VolumeSetStatus `json:"status,omitempty"`
	Gvc          any                       `json:"gvc,omitempty"`
}

type DeleteVolumeSetSpec struct {
	VolumeSet DeleteVolumeSetSpecVolumeSet `json:"volumeSet,omitempty"`
	Locations []string                     `json:"locations,omitempty"`
}

type DeleteVolumeSetLocationStatusStage string

const (
	DeleteVolumeSetLocationStatusStageDeleteVolumes         DeleteVolumeSetLocationStatusStage = "delete-volumes"
	DeleteVolumeSetLocationStatusStageDeleteOrphanedVolumes DeleteVolumeSetLocationStatusStage = "delete-orphaned-volumes"
	DeleteVolumeSetLocationStatusStageCleanupFilesystem     DeleteVolumeSetLocationStatusStage = "cleanup-filesystem"
	DeleteVolumeSetLocationStatusStageComplete              DeleteVolumeSetLocationStatusStage = "complete"
)

type DeleteVolumeSetLocationStatusVolumes map[string]DeleteCloudDevicesStatus

type DeleteVolumeSetLocationStatus struct {
	Stage   DeleteVolumeSetLocationStatusStage   `json:"stage,omitempty"`
	Volumes DeleteVolumeSetLocationStatusVolumes `json:"volumes,omitempty"`
}

type DeleteVolumeSetStatus map[string]DeleteVolumeSetLocationStatus
