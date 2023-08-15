package web

import "net/http"

type response struct {
	originalResponse http.ResponseWriter
}

func (r response) Write(body []byte) (int, error) {
	return r.originalResponse.Write(body)
}
