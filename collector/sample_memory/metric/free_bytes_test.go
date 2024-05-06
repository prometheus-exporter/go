package metric_test

import (
	"testing"

	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/prometheus-exporter/go/collector/sample_memory/metric"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestFreeBytes(t *testing.T) {
	collector := exporter.NewCollector([]exporter.Metric{&metric.FreeBytes{}})

	if _, err := testutil.CollectAndLint(collector); err != nil {
		t.Fatal(err)
	}

	if count := testutil.CollectAndCount(collector); count != 1 {
		t.Fatal("invalid -", count)
	}
}
