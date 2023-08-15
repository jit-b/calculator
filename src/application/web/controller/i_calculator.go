package controller

type calculator interface {
	Calculate(expression string) (float32, error)
}
