package tafexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLesserEqualsThanString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"abc\" <= \"efg\""))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserThanString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"abc\" < \"efg\""))
	assert.Equal(t, true, p.BoolValue)
}

func TestEqualsString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"abc\" == \"abc\""))
	assert.Equal(t, true, p.BoolValue)
}

func TestUnEqualsString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"abc\" != \"abc\""))
	assert.Equal(t, false, p.BoolValue)
}

func TestGreaterEqualsThanString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"efg\" >= \"abc\""))
	assert.Equal(t, true, p.BoolValue)
}

func TestGreaterThanString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"efg\" > \"abc\""))
	assert.Equal(t, true, p.BoolValue)
}

func TestEqualsStringInteger(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, false, p.Parse("\"123\" == 123"))
	assert.Equal(t, false, p.BoolValue)
}

func TestEqualsStringNull(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"\" == null"))
	assert.Equal(t, true, p.BoolValue)
}

func TestUnEqualsStringNull(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"string\" != null"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLargerAndEqualStringNull(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"string\" >= null"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLargerStringNull(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("\"string\" > null"))
	assert.Equal(t, true, p.BoolValue)
}

func TestUnEqualsNullString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse(" null != \"string\""))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserAndEqualNullString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("null  <= \"string\""))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserNullString(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("null < \"string\" "))
	assert.Equal(t, true, p.BoolValue)
}
