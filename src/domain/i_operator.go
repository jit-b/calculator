package domain

type Operator interface {
	AsString() string
	CreateExpression(exp1 Expression, exp2 Expression) Expression
}
