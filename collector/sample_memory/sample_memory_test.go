package sample_memory_test

import (
	"testing"

	"github.com/prometheus-exporter/go/collector/sample_memory"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestNew(t *testing.T) {
	collector := sample_memory.New()

	if _, err := testutil.CollectAndLint(collector); err != nil {
		t.Fatal(err)
	}
}
