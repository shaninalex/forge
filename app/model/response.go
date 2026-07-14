package model

import (
	"net/http"
)

type Response struct {
	Body       []byte
	StatusCode int
	Header     http.Header
	Duration   int64
}

type DataBank map[string]*Response

func (s DataBank) GetDataMap() map[string]string {
	output := make(map[string]string)

	for k, d := range s {
		output[k] = string(d.Body)
	}

	return output
}
