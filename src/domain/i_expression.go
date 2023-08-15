package domain

type Expression interface {
	Evaluate() float32
}
