package operators

import "Calculator/src/domain"

type Subtraction struct {
	expression1 domain.Expression
	expression2 domain.Expression
}

func (s Subtraction) CreateExpression(expression1 domain.Expression, expression2 domain.Expression) domain.Expression {
	return &Subtraction{
		expression1: expression1,
		expression2: expression2,
	}
}

func (s Subtraction) Evaluate() float32 {
	return s.expression1.Evaluate() - s.expression2.Evaluate()
}

func (s Subtraction) AsString() string {
	return "-"
}
