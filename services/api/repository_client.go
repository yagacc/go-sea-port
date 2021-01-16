package main

import (
	"context"
	pb "github.com/yagacc/go-sea-port/domain/repository/v1"
	"github.com/yagacc/go-sea-port/pkg/grpc"
)

type RepositoryClient struct {
	pb.PortRepositoryClient
	grpcConnections []*grpc.Connection
}

func (c *RepositoryClient) connect(ctx context.Context, repositoryAddr string) {
	conn := &grpc.Connection{Add: repositoryAddr}
	conn.Connect(ctx)
	c.PortRepositoryClient = pb.NewPortRepositoryClient(conn.Connection)
	c.grpcConnections = append(c.grpcConnections, conn)
}

func (g *RepositoryClient) stop() {
	for _, f := range g.grpcConnections {
		f.Stop()
	}
}
