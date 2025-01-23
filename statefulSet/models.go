/* auto-generated */

package statefulSet

import "github.com/controlplane-com/types-go/volumeSet"

type StatefulSetStatus struct {
	ReplicaCount float32                   `json:"replicaCount"`
	VolumeSet    volumeSet.VolumeSetStatus `json:"volumeSet,omitempty"`
}
