package metric

import (
	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/common-library/go/log/klog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
)

type TotalSeconds struct {
}

func (this *TotalSeconds) GetDesc() *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "total_seconds"),
		"cpu total seconds",
		[]string{"cpu", "mode"},
		nil)
}

func (this *TotalSeconds) GetValueType() prometheus.ValueType {
	return prometheus.CounterValue
}

func (this *TotalSeconds) GetValues() []exporter.Value {
	values := []exporter.Value{}

	if timesStats, err := cpu.Times(true); err != nil {
		klog.ErrorS(err, "")
	} else {
		for _, timesStat := range timesStats {
			values = append(values, exporter.Value{Value: timesStat.User, LabelValues: []string{timesStat.CPU, "user"}})
			values = append(values, exporter.Value{Value: timesStat.System, LabelValues: []string{timesStat.CPU, "system"}})
			values = append(values, exporter.Value{Value: timesStat.Idle, LabelValues: []string{timesStat.CPU, "idle"}})
		}
	}

	return values
}
