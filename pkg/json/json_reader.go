package json

import (
	"encoding/json"
	"github.com/yagacc/go-sea-port/domain/domain"
	"io/ioutil"
)

type NumberRecords int
type RecordProcessor func(string, *domain.Port) error
type Ports map[string]*domain.Port

type JsonReader interface {
	ReadFile(string, RecordProcessor) (NumberRecords, error)
}

//phase 1 get it working....inefficient impl...
type InMemJsonReader struct{}

func (s *InMemJsonReader) ReadFile(filename string, processor RecordProcessor) (NumberRecords, error) {
	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	var data Ports
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return 0, err
	}

	for key, port := range data {
		err = processor(key, port)
		if err != nil {
			return 0, err
		}
	}

	return NumberRecords(len(data)), nil
}
