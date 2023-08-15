package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Calculator/src/domain/operators"
	"Calculator/test/domain/mocks"
)

func TestSubtractionCreateExpressionAndEvaluate(t *testing.T) {
	mockedExpression1 := mocks.NewExpression(t)
	mockedExpression1.EXPECT().Evaluate().Return(7).Once()

	mockedExpression2 := mocks.NewExpression(t)
	mockedExpression2.EXPECT().Evaluate().Return(4).Once()

	operator := operators.Subtraction{}
	resultExpression := operator.CreateExpression(mockedExpression1, mockedExpression2)

	assert.Equal(t, float32(3), resultExpression.Evaluate())
}

func TestSubtractionAsString(t *testing.T) {
	operator := operators.Subtraction{}
	assert.Equal(t, string("-"), operator.AsString())
}
