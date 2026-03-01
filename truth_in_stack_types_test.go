package tafexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerTrue(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("25"))
	assert.Equal(t, true, p.Result.IsTruthy())
}

func TestIntegerFalse(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("0"))
	assert.Equal(t, false, p.Result.IsTruthy())
}

func TestDoubleTrue(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("25.32 + 501.32"))
	assert.Equal(t, true, p.Result.IsTruthy())
}

func TestDoubleFalse(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("(25.32 * 1000.22 ) * 0"))
	assert.Equal(t, false, p.Result.IsTruthy())
}

func TestBoolTrue(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("25.32 + 501.32 > 0"))
	assert.Equal(t, true, p.Result.IsTruthy())
}

func TestBoolFalse(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("(25.32 * 1000.22 ) * 0 > 0"))
	assert.Equal(t, false, p.Result.IsTruthy())
}

func TestBoolPrimTrue(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("true"))
	assert.Equal(t, true, p.Result.IsTruthy())
}

func TestBoolPrimFalse(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("false"))
	assert.Equal(t, false, p.Result.IsTruthy())
}

func TestStringTrue(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse(" \"string\""))
	assert.Equal(t, true, p.Result.IsTruthy())
}

func TestStringFalse(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"\""))
	assert.Equal(t, false, p.Result.IsTruthy())
}

func TestNull(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("null"))
	assert.Equal(t, false, p.Result.IsTruthy())
}
