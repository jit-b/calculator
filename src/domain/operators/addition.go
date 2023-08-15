package operators

import "Calculator/src/domain"

type Addition struct {
	expression1 domain.Expression
	expression2 domain.Expression
}

func (s Addition) CreateExpression(expression1 domain.Expression, expression2 domain.Expression) domain.Expression {
	return &Addition{
		expression1: expression1,
		expression2: expression2,
	}
}

func (s Addition) Evaluate() float32 {
	return s.expression1.Evaluate() + s.expression2.Evaluate()
}

func (s Addition) AsString() string {
	return "+"
}
