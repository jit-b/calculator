package web

type Response interface {
	Write([]byte) (int, error)
}
