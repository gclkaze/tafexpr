package tafexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerContract(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupVariableContext()
	assert.Equal(t, true, p.ParseExpressionAndCheckAgainstContract("123", "integer"))
	assert.Equal(t, true, p.ParseExpressionAndCheckAgainstContract("   123  ", "integer"))
	assert.Equal(t, false, p.ParseExpressionAndCheckAgainstContract("   123  2222", "integer"))
	assert.Equal(t, false, p.ParseExpressionAndCheckAgainstContract("\"123\"", "integer"))

	//variables
	assert.Equal(t, true, p.ParseExpressionAndCheckAgainstContract("$myVar + 1", "integer"))
	assert.Equal(t, true, p.ParseExpressionAndCheckAgainstContract("$myVar", "integer"))
	assert.Equal(t, true, p.ParseExpressionAndCheckAgainstContract("$myVar * $myVar2", "integer"))
}
