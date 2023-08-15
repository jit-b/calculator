package domain

import (
	"Calculator/src/domain"
	mocks2 "Calculator/test/domain/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidExpReturnNoErr(t *testing.T) {
	mockedOperator := mocks2.NewOperator(t)
	mockedOperator.EXPECT().AsString().Return("+").Times(6)

	mockedLeftOperand := mocks2.NewExpression(t)
	mockedValue := mocks2.NewValue(t)
	mockedValue.EXPECT().CreateExpression("1").Return(mockedLeftOperand, nil).Once()

	mockedRightOperand := mocks2.NewExpression(t)
	mockedValue.EXPECT().CreateExpression("2").Return(mockedRightOperand, nil).Once()

	resultMockedExpression := mocks2.NewExpression(t)

	mockedOperator.EXPECT().CreateExpression(mockedLeftOperand, mockedRightOperand).Return(resultMockedExpression).Once()

	resultMockedExpression.EXPECT().Evaluate().Return(float32(3)).Once()

	calculator := domain.NewCalculator([]domain.Operator{mockedOperator}, mockedValue)
	result, err := calculator.Calculate("1+2")
	assert.NoError(t, err)
	assert.Equal(t, float32(3), result)

}

func TestInvalidExpReturnError(t *testing.T) {
	mockedOperator := mocks2.NewOperator(t)
	mockedOperator.EXPECT().AsString().Return("+").Once()

	calculator := domain.NewCalculator([]domain.Operator{mockedOperator}, nil)
	result, err := calculator.Calculate("+1")
	assert.Error(t, err)
	assert.Equal(t, float32(0), result)

}

// в отдельном файле тест для операторов для + и -
// и в калькуляторе допокрыть то что не протестировали
