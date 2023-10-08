package grpc

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_xrequest_id "github.com/higebu/go-grpc-interceptor/xrequestid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/vladjong/go_project_template/configs"
	"github.com/vladjong/go_project_template/internal/controller/grpc/users"
	"github.com/vladjong/go_project_template/internal/services"
	users_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/users"
)

type GRPC struct {
	grpc *grpc.Server
}

func New() *GRPC {
	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_recovery.StreamServerInterceptor(),
				grpc_xrequest_id.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_xrequest_id.UnaryServerInterceptor(),
			),
		),
	)
	reflection.Register(server)
	return nil
}

func (s *GRPC) Start(cfg configs.Config) error {
	address := ""

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	return s.grpc.Serve(l)
}

func (s *GRPC) Stop() error {
	s.grpc.GracefulStop()
	return nil
}

func (s *GRPC) InitServices(service services.Services) {
	users_grpc.RegisterUsersServer(s.grpc, users.New(service))
}
