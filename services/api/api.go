package main

import (
	"context"
	"fmt"
	api "github.com/yagacc/go-sea-port/domain/api/v1"
	pb "github.com/yagacc/go-sea-port/domain/repository/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	unimplErrorMsg   = "not yet implemented"
	notFoundErrorFmt = "%s not found"
)

type Api struct {
	*RepositoryClient
}

func (a *Api) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	p, err := a.RepositoryClient.Get(ctx, &pb.GetRequest{
		PortId: req.PortId,
	})
	if err != nil {
		return nil, err
	}
	if p.Port == nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf(notFoundErrorFmt, req.PortId))
	}
	return &api.GetResponse{Port: p.Port}, nil
}

func (a *Api) List(context.Context, *api.ListRequest) (*api.ListResponse, error) {
	return nil, status.Error(codes.Unimplemented, unimplErrorMsg)
}

func (a *Api) Save(context.Context, *api.SaveRequest) (*api.SaveResponse, error) {
	return nil, status.Error(codes.Unimplemented, unimplErrorMsg)
}

func (a *Api) Create(context.Context, *api.CreateRequest) (*api.CreateResponse, error) {
	return nil, status.Error(codes.Unimplemented, unimplErrorMsg)
}

func (a *Api) Delete(context.Context, *api.DeleteRequest) (*api.DeleteResponse, error) {
	return nil, status.Error(codes.Unimplemented, unimplErrorMsg)
}
