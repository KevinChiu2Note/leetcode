package design_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddOperator_Result(t *testing.T) {
	var op Operator
	op = AddOperatorFactory{}.CreateFactory()
	op.SetA(10)
	op.SetB(20)
	assert.Equal(t, op.Result(), 30)
}

func TestSubOperator_Result(t *testing.T) {
	var op Operator
	op = SubOperatorFactory{}.CreateFactory()
	op.SetA(10)
	op.SetB(20)
	assert.Equal(t, op.Result(), 10)
}
