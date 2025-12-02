package grpcserver

import (
	"time"

	"google.golang.org/grpc"
)

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		s.address = ":" + port
	}
}

func Address(addr string) Option {
	return func(s *Server) {
		s.address = addr
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

func WithGRPCOptions(opts ...grpc.ServerOption) Option {
	return func(s *Server) {
		s.server = grpc.NewServer(opts...)
	}
}
