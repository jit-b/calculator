package domain

type Value interface {
	CreateExpression(s string) (Expression, error)
}
