/* auto-generated */

package workloadOptions

import "github.com/controlplane-com/types-go/pkg/base"

type OptionsAutoscalingMetric string

const (
	OptionsAutoscalingMetricConcurrency OptionsAutoscalingMetric = "concurrency"
	OptionsAutoscalingMetricCpu         OptionsAutoscalingMetric = "cpu"
	OptionsAutoscalingMetricMemory      OptionsAutoscalingMetric = "memory"
	OptionsAutoscalingMetricRps         OptionsAutoscalingMetric = "rps"
	OptionsAutoscalingMetricLatency     OptionsAutoscalingMetric = "latency"
	OptionsAutoscalingMetricDisabled    OptionsAutoscalingMetric = "disabled"
)

type OptionsAutoscaling struct {
	Metric           OptionsAutoscalingMetric `json:"metric,omitempty"`
	Multi            any                      `json:"multi,omitempty"`
	MetricPercentile any                      `json:"metricPercentile,omitempty"`
	Target           float32                  `json:"target"`
	MaxScale         float32                  `json:"maxScale"`
	MinScale         float32                  `json:"minScale"`
	ScaleToZeroDelay float32                  `json:"scaleToZeroDelay"`
	MaxConcurrency   float32                  `json:"maxConcurrency"`
}

type Options struct {
	Autoscaling    OptionsAutoscaling    `json:"autoscaling,omitempty"`
	TimeoutSeconds float32               `json:"timeoutSeconds"`
	CapacityAI     bool                  `json:"capacityAI,omitempty"`
	Spot           bool                  `json:"spot,omitempty"`
	Debug          bool                  `json:"debug,omitempty"`
	Suspend        bool                  `json:"suspend,omitempty"`
	MultiZone      base.MultiZoneOptions `json:"multiZone,omitempty"`
	Location       string                `json:"location,omitempty"`
}

type DefaultOptions Options

type LocalOptions []Options
