package grpcserver

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	server          *grpc.Server
	notify          chan error
	shutdownTimeout time.Duration
	address         string
}

func New(opts ...Option) *Server {
	s := &Server{
		server:          grpc.NewServer(),
		notify:          make(chan error, 1),
		shutdownTimeout: 5 * time.Second,
		address:         ":50051",
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) GetServer() *grpc.Server {
	return s.server
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}

	go func() {
		slog.Info("Starting gRPC server", "address", s.address)
		s.notify <- s.server.Serve(lis)
		close(s.notify)
	}()

	return nil
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() {
	done := make(chan struct{})

	go func() {
		s.server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		slog.Info("gRPC server stopped gracefully")
	case <-time.After(s.shutdownTimeout):
		slog.Warn("gRPC server shutdown timeout, forcing stop")
		s.server.Stop()
	}
}

func Run(registerFunc func(*grpc.Server), opts ...Option) error {
	srv := New(opts...)

	registerFunc(srv.GetServer())

	if err := srv.Start(); err != nil {
		return err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-interrupt:
		slog.Info("Received shutdown signal", "signal", sig.String())
	case err := <-srv.Notify():
		slog.Error("gRPC server error", "error", err)
	}

	srv.Shutdown()

	slog.Info("gRPC server stopped")
	return nil
}
