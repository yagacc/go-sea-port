package json

import (
	"bufio"
	"encoding/json"
	"github.com/yagacc/go-sea-port/domain/domain"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	open          = '{'
	close         = '}'
	jsonPortRegex = regexp.MustCompile("(?s)\"([A-Z]{5})\": \\{(.*?)\\}")
)

type NumberRecords int
type RecordProcessor func(string, *domain.Port) error
type Ports map[string]*domain.Port

type JsonReader interface {
	ReadFile(string, RecordProcessor) (NumberRecords, error)
}

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

type BufferedJsonReader struct {
	BufferSize int
}

func (r *BufferedJsonReader) ReadFile(filename string, processor RecordProcessor) (NumberRecords, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	var totalRecords NumberRecords
	reader := bufio.NewReader(jsonFile)

	var rawJson *RawJson
	var currentRawJsonObject string
	var tokens []rune
	for {
		buf := make([]byte, r.BufferSize)
		n, err := reader.Read(buf)
		buf = buf[:n]

		chunk := string(buf)

		rawJson, currentRawJsonObject, tokens = getJsonObject(chunk, currentRawJsonObject, tokens)
		if rawJson != nil {
			var port domain.Port
			err := json.Unmarshal([]byte("{"+rawJson.object+"}"), &port)
			if err != nil {
				return 0, err
			}
			err = processor(rawJson.key, &port)
			if err != nil {
				return 0, err
			}
			totalRecords++
		}

		if n == 0 {
			if err == io.EOF {
				break
			}
			//this should never happen - err!=EOF but no bytes read!
			return 0, err
		}
	}

	return totalRecords, jsonFile.Close()
}

type RawJson struct {
	key, object string
}

func getJsonObject(chunk, current string, tokens []rune) (*RawJson, string, []rune) {
	var rawJson *RawJson
	concat := current
	for _, c := range chunk {
		concat += string(c)
		if c == open {
			tokens = append(tokens, c)
		} else if c == close {
			tokens = tokens[:len(tokens)-1]
			if len(tokens) == 1 { //only opening { left, we have a full Json object (0=EOF)
				match := jsonPortRegex.FindStringSubmatch(concat)
				if match != nil {
					idx := jsonPortRegex.FindStringIndex(concat)
					rawJson = &RawJson{
						key:    match[1],
						object: match[2],
					}
					concat = concat[idx[1]:]
				}
			}
		}
	}
	return rawJson, concat, tokens
}
