/* auto-generated */

package mk8sCommon

type GoDuration string

type Labels map[string]string

type TaintEffect string

const (
	TaintEffectNoSchedule       TaintEffect = "NoSchedule"
	TaintEffectPreferNoSchedule TaintEffect = "PreferNoSchedule"
	TaintEffectNoExecute        TaintEffect = "NoExecute"
)

type Taint struct {
	Key    string      `json:"key,omitempty"`
	Value  string      `json:"value,omitempty"`
	Effect TaintEffect `json:"effect,omitempty"`
}

type Taints []Taint

type NodePoolName string

type SshPublicKey string

type AutoscalerConfigExpander string

const (
	AutoscalerConfigExpanderRandom     AutoscalerConfigExpander = "random"
	AutoscalerConfigExpanderMostPods   AutoscalerConfigExpander = "most-pods"
	AutoscalerConfigExpanderLeastWaste AutoscalerConfigExpander = "least-waste"
	AutoscalerConfigExpanderPrice      AutoscalerConfigExpander = "price"
)

type AutoscalerConfig struct {
	Expander             []AutoscalerConfigExpander `json:"expander,omitempty"`
	UnneededTime         string                     `json:"unneededTime,omitempty"`
	UnreadyTime          string                     `json:"unreadyTime,omitempty"`
	UtilizationThreshold float32                    `json:"utilizationThreshold"`
}

type UnmanagedPool struct {
	Name   string `json:"name,omitempty"`
	Labels Labels `json:"labels,omitempty"`
	Taints Taints `json:"taints,omitempty"`
}

type PreInstallScript string

type CacertsRes struct {
	Cacerts string `json:"cacerts,omitempty"`
}

type ReadyRes struct {
	Ready   bool   `json:"ready,omitempty"`
	Message string `json:"message,omitempty"`
}

type JoinBody map[string]any

type JoinRes struct {
	Script string `json:"script,omitempty"`
}

type InitBody map[string]any

type InitRes struct {
	Script string `json:"script,omitempty"`
}

type KubeConfigBody map[string]any

type KubeConfigRes struct {
	Kubeconfig string `json:"kubeconfig,omitempty"`
	FileName   string `json:"fileName,omitempty"`
}
