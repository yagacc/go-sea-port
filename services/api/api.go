package main

import (
	"context"
	api "github.com/yagacc/go-sea-port/domain/api/v1"
	"github.com/yagacc/go-sea-port/domain/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	unimplErrorMsg = "not yet implemented"
)

type Api struct {}

func (a *Api) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	return &api.GetResponse{Port: &domain.Port{
		Id: req.PortId,
	}}, nil
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