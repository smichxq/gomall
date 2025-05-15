package mtl

import "github.com/kitex-contrib/obs-opentelemetry/provider"

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("192.168.3.6:4317"),
		provider.WithInsecure(),
		// 改变opentelemetry自带的Metric防止与Prometheus冲突
		provider.WithEnableMetrics(false),
	)

	return p
}
