package tafexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegCalc(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("-50.2 * ( 10 + 100 )"))
	assert.Equal(t, -5522, p.IntValue)
}

func TestNegCalcWithNegParenthesis(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("-(5522 - 10.5 + 10.5)"))
	assert.Equal(t, -5522, p.IntValue)
}

func TestNegEqualsDoubleDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("-50.2 * ( 10 + 100 ) == -(5522 - 10.5 + 10.5)"))
	assert.Equal(t, true, p.BoolValue)
}

func TestNegEqualsIntIntSimple(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("-(2) == -2"))
	assert.Equal(t, true, p.BoolValue)
}

func TestNegEqualsIntInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("-(50 * ( 10 + 100 )) == -5500 - 10 + 10"))
	assert.Equal(t, true, p.BoolValue)
}

func TestNegEqualsDoubleInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("-2 * 3.5 == -7"))
	assert.Equal(t, true, p.BoolValue)
}

func TestNegEqualsIntDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse(" -8 == -3 * 3.2 + (0.8 * 2)"))
	assert.Equal(t, true, p.BoolValue)
}
