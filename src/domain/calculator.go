package domain

import (
	"errors"
	"strings"
)

type Calculator struct {
	operators []Operator
	value     Value
}

func NewCalculator(operators []Operator, value Value) *Calculator {
	return &Calculator{operators: operators, value: value}
}

func (c Calculator) Calculate(expression string) (float32, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	evaluatableExpression, err := c.makeExpression(expression)

	if err != nil {
		return 0, err
	}

	return evaluatableExpression.Evaluate(), nil
}

func (c Calculator) makeExpression(expression string) (Expression, error) {
	for i := range c.operators {
		operator := c.operators[i]

		if string(expression[0]) == operator.AsString() {
			return nil, errors.New("operator has zero position")
		}

		indexOperator := strings.Index(expression, operator.AsString())

		if indexOperator == -1 {
			continue
		}

		if indexOperator == len(expression)-1 {
			return nil, errors.New("operator has last position")
		}

		leftOperand, err := c.makeExpression(expression[:indexOperator]) //
		if err != nil {
			return nil, err
		}

		rightOperand, err := c.makeExpression(expression[indexOperator+1:])
		if err != nil {
			return nil, err
		}

		return operator.CreateExpression(leftOperand, rightOperand), nil
	}

	return c.value.CreateExpression(expression)
}
