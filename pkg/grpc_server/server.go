package grpc_server

import (
	"log/slog"
	"net"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	_defaultAddr            = ":96"
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	server          *grpc.Server
	notify          chan error
	shutdownTimeout time.Duration
	addr            string
}

func New(opts ...Option) *Server {
	logger := slog.Default()

	optsGrpc := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), optsGrpc...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(logger), optsGrpc...),
		),
	)

	server := &Server{
		server:          s,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		addr:            _defaultAddr,
	}
	reflection.Register(s)

	for _, opt := range opts {
		opt(server)
	}

	server.start()
	slog.Info("Listening to gRPC", "port", server.addr)

	return server
}

func (s *Server) start() {
	go func() {
		l, err := net.Listen("tcp", s.addr)
		if err != nil {
			s.notify <- err
			return
		}
		s.notify <- s.server.Serve(l)
		close(s.notify)
	}()
}

func (s *Server) Server() *grpc.Server {
	return s.server
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
