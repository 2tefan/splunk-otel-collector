// Code generated by monitor-code-gen. DO NOT EDIT.

package exec

import (
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "telegraf/exec"

var groupSet = map[string]bool{}

var metricSet = map[string]monitors.MetricInfo{}

var defaultMetrics = map[string]bool{}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "telegraf/exec",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         true,
}