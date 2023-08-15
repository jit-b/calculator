package operators

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Calculator/src/domain/operators"
	"Calculator/test/domain/mocks"
)

func TestAdditionCreateExpressionAndEvaluate(t *testing.T) {
	mockedExpression1 := mocks.NewExpression(t)
	mockedExpression1.EXPECT().Evaluate().Return(5).Once()

	mockedExpression2 := mocks.NewExpression(t)
	mockedExpression2.EXPECT().Evaluate().Return(7).Once()

	operator := operators.Addition{}
	resultExpression := operator.CreateExpression(mockedExpression1, mockedExpression2)

	assert.Equal(t, float32(12), resultExpression.Evaluate())
}

func TestAdditionAsString(t *testing.T) {
	operator := operators.Addition{}
	assert.Equal(t, string("+"), operator.AsString())
}
