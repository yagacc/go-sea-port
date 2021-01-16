package main

import "github.com/yagacc/go-sea-port/domain/domain"

type PortStorage interface {
	Ping() error
	Update(*domain.Port) (*domain.Port, error)
	Get(string) *domain.Port
	Close() error
}
