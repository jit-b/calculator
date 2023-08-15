package web

type Request interface {
	GetBody() []byte
}
