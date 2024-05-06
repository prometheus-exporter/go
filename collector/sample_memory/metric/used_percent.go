package metric

import (
	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/common-library/go/log/klog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/mem"
)

type UsedPercent struct {
}

func (this *UsedPercent) GetDesc() *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "used_percent"),
		"memory total bytes",
		nil,
		nil)
}

func (this *UsedPercent) GetValueType() prometheus.ValueType {
	return prometheus.GaugeValue
}

func (this *UsedPercent) GetValues() []exporter.Value {
	if virtualMemoryStat, err := mem.VirtualMemory(); err != nil {
		klog.ErrorS(err, "")
		return nil
	} else {
		return []exporter.Value{exporter.Value{Value: virtualMemoryStat.UsedPercent}}
	}
}
