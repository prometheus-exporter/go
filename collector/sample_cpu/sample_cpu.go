package sample_cpu

import (
	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/prometheus-exporter/go/collector/sample_cpu/metric"
	"github.com/prometheus/client_golang/prometheus"
)

func New() prometheus.Collector {
	return exporter.NewCollector(
		[]exporter.Metric{
			&metric.ModelInfo{},
			&metric.TotalSeconds{},
		})
}
