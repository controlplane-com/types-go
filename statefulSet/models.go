/* auto-generated */

package statefulSet

import "gitlab.com/controlplane/controlplane/go-libs/schema/volumeSet"

type StatefulSetStatus struct {
	ReplicaCount float32                   `json:"replicaCount"`
	VolumeSet    volumeSet.VolumeSetStatus `json:"volumeSet,omitempty"`
}
