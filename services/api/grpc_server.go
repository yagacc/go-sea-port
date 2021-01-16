package main

import (
	api "github.com/yagacc/go-sea-port/domain/api/v1"
	"github.com/yagacc/go-sea-port/pkg/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func (s *GrpcServer) listenAndServe() {
	api.RegisterPortApiServer(s.server.Server, &Api{})
	s.server.Start()
}

func (s *GrpcServer) stop() {
	s.server.Stop()
}
