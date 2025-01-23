/* auto-generated */

package volumeSpec

type VolumeSpecRecoveryPolicy string

const (
	VolumeSpecRecoveryPolicyRetain  VolumeSpecRecoveryPolicy = "retain"
	VolumeSpecRecoveryPolicyRecycle VolumeSpecRecoveryPolicy = "recycle"
)

type VolumeSpec struct {
	Uri            string                   `json:"uri,omitempty"`
	RecoveryPolicy VolumeSpecRecoveryPolicy `json:"recoveryPolicy,omitempty"`
	Path           string                   `json:"path,omitempty"`
}
