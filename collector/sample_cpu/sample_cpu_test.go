package sample_cpu_test

import (
	"testing"

	"github.com/prometheus-exporter/go/collector/sample_cpu"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestNew(t *testing.T) {
	collector := sample_cpu.New()

	if _, err := testutil.CollectAndLint(collector); err != nil {
		t.Fatal(err)
	}
}
