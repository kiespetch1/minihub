package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type BaseMetrics struct {
	namespace string

	HTTPRequestsTotal     *prometheus.CounterVec
	HTTPRequestDuration   *prometheus.HistogramVec
	HTTPRequestSize       *prometheus.HistogramVec
	HTTPResponseSize      *prometheus.HistogramVec
	HTTPActiveConnections prometheus.Gauge

	GRPCRequestsTotal   *prometheus.CounterVec
	GRPCRequestDuration *prometheus.HistogramVec

	AppInfo      *prometheus.GaugeVec
	AppUptime    prometheus.Gauge
	AppStartTime prometheus.Gauge
}

func NewBaseMetrics(namespace, serviceName string) *BaseMetrics {
	return &BaseMetrics{
		namespace: namespace,

		HTTPRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "http_requests_total",
				Help:      "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),

		HTTPRequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "http_request_duration_seconds",
				Help:      "HTTP request latency in seconds",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method", "path", "status"},
		),

		HTTPRequestSize: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "http_request_size_bytes",
				Help:      "HTTP request size in bytes",
				Buckets:   prometheus.ExponentialBuckets(100, 10, 8),
			},
			[]string{"method", "path"},
		),

		HTTPResponseSize: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "http_response_size_bytes",
				Help:      "HTTP response size in bytes",
				Buckets:   prometheus.ExponentialBuckets(100, 10, 8),
			},
			[]string{"method", "path"},
		),

		HTTPActiveConnections: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "http_active_connections",
				Help:      "Number of active HTTP connections",
			},
		),

		GRPCRequestsTotal: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "grpc_requests_total",
				Help:      "Total number of gRPC requests",
			},
			[]string{"method", "status"},
		),

		GRPCRequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "grpc_request_duration_seconds",
				Help:      "gRPC request latency in seconds",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method", "status"},
		),

		AppInfo: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "app_info",
				Help:      "Application information",
			},
			[]string{"version", "env", "service"},
		),

		AppUptime: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "app_uptime_seconds",
				Help:      "Application uptime in seconds",
			},
		),

		AppStartTime: promauto.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: serviceName,
				Name:      "app_start_time_seconds",
				Help:      "Application start time in unix timestamp",
			},
		),
	}
}

func (m *BaseMetrics) SetAppInfo(version, env, service string) {
	m.AppInfo.WithLabelValues(version, env, service).Set(1)
}
