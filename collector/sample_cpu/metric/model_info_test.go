package metric_test

import (
	"testing"

	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/prometheus-exporter/go/collector/sample_cpu/metric"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestModelInfo(t *testing.T) {
	collector := exporter.NewCollector([]exporter.Metric{&metric.ModelInfo{}})

	if _, err := testutil.CollectAndLint(collector); err != nil {
		t.Fatal(err)
	}
}
