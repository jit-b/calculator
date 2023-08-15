package web

type Controller interface {
	Handle(response Response, request Request)
}
