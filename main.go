package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-library/go/command-line/flags"
	"github.com/common-library/go/database/prometheus/exporter"
	"github.com/common-library/go/log/klog"
	"github.com/prometheus-exporter/go/collector/sample_cpu"
	"github.com/prometheus-exporter/go/collector/sample_memory"
	"github.com/prometheus/client_golang/prometheus"
)

func parseFlags() error {
	flagInfos := []flags.FlagInfo{
		{FlagName: "address", Usage: ":10000", DefaultValue: string(":10000")},
		{FlagName: "timeout", Usage: "60", DefaultValue: time.Duration(60)},
		{FlagName: "url-path", Usage: "/metrics", DefaultValue: string("/metrics")},
	}

	if err := flags.Parse(flagInfos); err != nil {
		flag.Usage()
		return err
	}

	flag.Visit(func(f *flag.Flag) { klog.InfoS("flags", "name", f.Name, "value", f.Value) })

	return nil
}

func run() error {
	if err := parseFlags(); err != nil {
		return err
	}

	address := flags.Get[string]("address")
	timeout := flags.Get[time.Duration]("timeout")
	urlPath := flags.Get[string]("url-path")

	collectors := []prometheus.Collector{
		sample_cpu.New(),
		sample_memory.New(),
	}

	if err := exporter.RegisterCollector(collectors...); err != nil {
		return err
	}
	defer exporter.UnRegisterCollector(collectors...)

	listenAndServeFailureFunc := func(err error) { klog.ErrorS(err, "") }
	if err := exporter.Start(address, urlPath, listenAndServeFailureFunc); err != nil {
		return err
	}
	defer exporter.Stop(timeout)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	klog.InfoS("signal", "kind", <-signals)

	return nil
}

func main() {
	defer klog.Flush()

	klog.InfoS("main start")
	defer klog.InfoS("main end")

	if err := run(); err != nil {
		klog.ErrorS(err, "")
	}
}
