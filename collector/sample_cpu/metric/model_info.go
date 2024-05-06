package metric

import (
	"fmt"

	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/common-library/go/log/klog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/cpu"
)

type ModelInfo struct {
}

func (this *ModelInfo) GetDesc() *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "model_info"),
		"cpu model info",
		[]string{"cpu", "model"},
		nil)
}

func (this *ModelInfo) GetValueType() prometheus.ValueType {
	return prometheus.GaugeValue
}

func (this *ModelInfo) GetValues() []exporter.Value {
	values := []exporter.Value{}

	if infoStats, err := cpu.Info(); err != nil {
		klog.ErrorS(err, "")
	} else {
		for _, infoStat := range infoStats {
			values = append(values, exporter.Value{Value: 1, LabelValues: []string{fmt.Sprintf("%d", infoStat.CPU), infoStat.ModelName}})
		}
	}

	return values
}
