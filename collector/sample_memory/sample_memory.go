package sample_memory

import (
	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/prometheus-exporter/go/collector/sample_memory/metric"
	"github.com/prometheus/client_golang/prometheus"
)

func New() prometheus.Collector {
	return exporter.NewCollector(
		[]exporter.Metric{
			&metric.FreeBytes{},
			&metric.TotalBytes{},
			&metric.UsedBytes{},
			&metric.UsedPercent{},
		})
}
