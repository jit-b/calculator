package web

import (
	"net/http"

	"Calculator/src/application/web"
)

type server struct {
	controller web.Controller
}

func NewServer() *server {
	return &server{}
}

func (s server) Run(port string, controller web.Controller) {
	s.controller = controller
	http.ListenAndServe(":"+port, s)
}

func (s server) ServeHTTP(originalResponse http.ResponseWriter, originalRequest *http.Request) {
	request := request{
		originalRequest: originalRequest,
	}
	response := response{originalResponse: originalResponse}
	s.controller.Handle(response, request)
}
