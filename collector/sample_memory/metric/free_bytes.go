package metric

import (
	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/common-library/go/log/klog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/mem"
)

type FreeBytes struct {
}

func (this *FreeBytes) GetDesc() *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "free_bytes"),
		"memory total bytes",
		nil,
		nil)
}

func (this *FreeBytes) GetValueType() prometheus.ValueType {
	return prometheus.GaugeValue
}

func (this *FreeBytes) GetValues() []exporter.Value {
	if virtualMemoryStat, err := mem.VirtualMemory(); err != nil {
		klog.ErrorS(err, "")
		return nil
	} else {
		return []exporter.Value{exporter.Value{Value: float64(virtualMemoryStat.Free)}}
	}
}
