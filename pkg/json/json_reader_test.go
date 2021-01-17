package json_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yagacc/go-sea-port/domain/domain"
	"github.com/yagacc/go-sea-port/pkg/json"
	"io/ioutil"
	"testing"
)

const (
	filename = "/tmp/file.json"
)

const (
	jsonContent = `{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }}`
)

func init() {
	ioutil.WriteFile(filename, []byte(jsonContent), 0666)
}

func TestInMemJsonReaderReadsSuccessfully(t *testing.T) {
	inMemJsonReader := &json.InMemJsonReader{}
	doNothingFunc := func(string, *domain.Port) error { return nil }
	count, err := inMemJsonReader.ReadFile(filename, doNothingFunc)
	assert.Nil(t, err)
	assert.Equal(t, 2, int(count), "number of records should be 2")
}

func TestJsonUnmarshalsToDomainCorrectly(t *testing.T) {
	inMemJsonReader := &json.InMemJsonReader{}
	serialiseIntoPorts := func(key string, port *domain.Port) error {
		if key == "AEAJM" {
			assert.Equal(t, "Ajman", port.Name, "name should be Ajman")
		} else if key == "AEAUH" {
			assert.Equal(t, "Abu Dhabi", port.Name, "name should be Abu Dhabi")
		}
		return nil
	}

	count, err := inMemJsonReader.ReadFile(filename, serialiseIntoPorts)
	assert.Nil(t, err)
	assert.Equal(t, 2, int(count), "number of records should be 2")
}

func TestBufferedJsonReaderReadsSuccessfully(t *testing.T) {
	bufferedJsonReader := &json.BufferedJsonReader{BufferSize: 128}
	doNothingFunc := func(string, *domain.Port) error { return nil }
	count, err := bufferedJsonReader.ReadFile(filename, doNothingFunc)
	assert.Nil(t, err)
	assert.Equal(t, 2, int(count), "number of records should be 2")
}
