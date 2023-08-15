package values

import (
	"Calculator/src/domain"
	"errors"
	"strconv"
)

type Value float32

func (v Value) Evaluate() float32 {
	return float32(v)
}

func (v Value) CreateExpression(valueAsStr string) (domain.Expression, error) {
	valueAsFloat, err := strconv.ParseFloat(valueAsStr, 64)
	if err != nil {
		return nil, errors.New("value isn't float")
	}

	return Value(valueAsFloat), nil
}
