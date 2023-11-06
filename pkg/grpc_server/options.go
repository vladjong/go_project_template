package grpc_server

import "net"

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		s.addr= net.JoinHostPort("", port)
	}
}

