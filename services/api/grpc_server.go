package main

import (
	api "github.com/yagacc/go-sea-port/domain/api/v1"
	"github.com/yagacc/go-sea-port/pkg/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func (s *GrpcServer) listenAndServe(c *RepositoryClient) {
	api.RegisterPortApiServer(s.server.Server, &Api{RepositoryClient: c})
	s.server.Start()
}

func (s *GrpcServer) stop() {
	s.server.Stop()
}
