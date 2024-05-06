# Prometheus Exporter using Go

## run
 - `go build -o prometheus-exporter ./main.go`
 - `./prometheus-exporter -address=:10000 -timeout=10s -url-path=/metric`
 - run `127.0.0.1:10000/metric` in your browser
