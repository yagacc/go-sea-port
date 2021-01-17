package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/yagacc/go-sea-port/domain/domain"
	"testing"
)

func TestInMemPortStorageErrorsWhenUpdatingToNilDb(t *testing.T) {
	//setup
	inMemStorage := &InMemPortStorage{}
	p, err := inMemStorage.Update(&domain.Port{
		Id:      "AB1",
		Name:    "AB1-Name",
		City:    "AB1-City",
		Country: "AB1-Country",
	})
	assert.Nil(t, p)
	assert.NotNil(t, err)
}

func TestInMemPortStorage(t *testing.T) {
	//setup
	inMemStorage := NewInMemPortStorage()
	inMemStorage.Update(&domain.Port{
		Id:      "AB1",
		Name:    "AB1-Name",
		City:    "AB1-City",
		Country: "AB1-Country",
	})
	//test
	tests := []struct {
		name        string
		storage     PortStorage
		id          string
		expectedNil bool
	}{
		{"get when not initialised", &InMemPortStorage{}, "AB1", true},
		{"get when present", inMemStorage, "AB1", false},
		{"get when missing", inMemStorage, "AB2", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.storage.Get(tt.id)
			if tt.expectedNil {
				assert.Nil(t, p)
			} else {
				assert.NotNil(t, p)
			}
		})
	}
}
