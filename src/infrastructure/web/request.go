package web

import "net/http"

type request struct {
	originalRequest *http.Request
}

func (r request) GetBody() []byte {
	ioWriter := r.originalRequest.Body
	result := []byte{}
	ioWriter.Read(result)

	return result
}
