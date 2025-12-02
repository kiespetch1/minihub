package metrics

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor(base *BaseMetrics) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start).Seconds()
		code := status.Code(err).String()

		base.GRPCRequestsTotal.WithLabelValues(info.FullMethod, code).Inc()
		base.GRPCRequestDuration.WithLabelValues(info.FullMethod, code).Observe(duration)

		return resp, err
	}
}

func StreamServerInterceptor(base *BaseMetrics) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		start := time.Now()

		err := handler(srv, ss)

		duration := time.Since(start).Seconds()
		code := status.Code(err).String()

		base.GRPCRequestsTotal.WithLabelValues(info.FullMethod, code).Inc()
		base.GRPCRequestDuration.WithLabelValues(info.FullMethod, code).Observe(duration)

		return err
	}
}
