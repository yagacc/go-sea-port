package main

import (
	"fmt"
	"github.com/yagacc/go-sea-port/domain/domain"
)

type InMemPortStorage struct {
	db map[string]*domain.Port
}

func NewInMemPortStorage() *InMemPortStorage {
	return &InMemPortStorage{db: make(map[string]*domain.Port)}
}

func (r *InMemPortStorage) Update(p *domain.Port) (*domain.Port, error) {
	if r.db == nil {
		return nil, fmt.Errorf("inmem storage is not initialised")
	}
	r.db[p.Id] = p
	return r.db[p.Id], nil
}

func (r *InMemPortStorage) Get(id string) *domain.Port {
	if r.db == nil {
		return nil
	}
	return r.db[id]
}

func (r *InMemPortStorage) Ping() error {
	return nil
}

func (r *InMemPortStorage) Close() error {
	return nil
}
