package model

import "net/http"

type Response struct {
	Body       []byte
	StatusCode int
	Header     http.Header
	Duration   int64
}
