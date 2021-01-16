package main

import (
	pb "github.com/yagacc/go-sea-port/domain/repository/v1"
	"github.com/yagacc/go-sea-port/pkg/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func (s *GrpcServer) listenAndServe(repo PortStorage) {
	pb.RegisterPortRepositoryServer(s.server.Server, &Repository{storage: repo})
	s.server.Start()
}

func (s *GrpcServer) stop() {
	s.server.Stop()
}
