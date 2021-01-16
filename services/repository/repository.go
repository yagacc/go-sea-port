package main

import (
	"context"
	repo "github.com/yagacc/go-sea-port/domain/repository/v1"
)

type Repository struct {
	storage PortStorage
}

func (r *Repository) Get(ctx context.Context, req *repo.GetRequest) (*repo.GetResponse, error) {
	p := r.storage.Get(req.PortId)
	return &repo.GetResponse{Port: p}, nil
}

func (r *Repository) Save(ctx context.Context, req *repo.SaveRequest) (*repo.SaveResponse, error) {
	p, err := r.storage.Update(req.Port)
	if err != nil {
		return nil, err
	}
	return &repo.SaveResponse{Port: p}, nil
}
