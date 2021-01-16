package grpc

import (
	g "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

type Server struct {
	Server             *g.Server
	GracefullyShutdown bool
	tcpAddress         string
}

func NewServer(tcpAddress string) *Server {
	//new server with generic health endpoint
	svr := &Server{
		Server:     g.NewServer(),
		tcpAddress: tcpAddress,
	}
	grpc_health_v1.RegisterHealthServer(svr.Server, health.NewServer()) //TODO: generic health - will not check dependencies
	return svr
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp", s.tcpAddress)
	if err != nil {
		log.Fatalf("Failed to start gRPC server : %v", err)
	}
	defer lis.Close()

	if err := s.Server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

func (s *Server) Stop() {
	s.Server.GracefulStop()
}
