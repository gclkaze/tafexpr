package tafexpr

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// /LESSER EQUALS
func TestLesserEqualsThanIntInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1 <= 100 + 2"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserEqualsThanDoubleInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1.5 <= 100 + 2"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserEqualsThanIntDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1 <= 100 + 2.72"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserEqualsThanDoubleDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 2 * 5.6 <= 100 + 2.72 * 9.99"))
	assert.Equal(t, true, p.BoolValue)
}

// //////////////////////////////////////////////////////////////////////////////////////
// LESSER THAN
func TestLesserThanDoubleDoubleInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1.11 < 100 + 2"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserThanDoubleDoubleDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1.11 < 100 + 2 * 8.31 - 2"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserThanDoubleIntDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1 * 33 < 100 + 2 * 1091.23"))
	assert.Equal(t, true, p.BoolValue)
}

func TestLesserThanIntInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 40 * 32 < 100 + 2 * 99 * 22"))
	assert.Equal(t, true, p.BoolValue)
}

// //////////////////////////////////////////////////////////////////////////////////////
// EQUALS
func TestEqualsDoubleDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("50.2 * ( 10 + 100 ) == 5522 - 10.5 + 10.5"))
	assert.Equal(t, true, p.BoolValue)
}

func TestEqualsIntInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("50 * ( 10 + 100 ) == 5500 - 10 + 10"))
	assert.Equal(t, true, p.BoolValue)
}

func TestEqualsDoubleInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("2 * 3.5 == 7"))
	assert.Equal(t, true, p.BoolValue)
}

func TestEqualsIntDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse(" 8 == 3 * 3.2 - (0.8 * 2)"))
	assert.Equal(t, true, p.BoolValue)
}

// //////////////////////////////////////////////////////////////////////////////////////
// UNEQUALS

func TestUnEqualsDoubleDouble(t *testing.T) {
	p := SetupParser()
	p.Debug = true
	res := p.Parse("50.2 * ( 10 + 100 ) + 0.2 != 5522 - 10.5 + 10.5 + 0.1")
	for i := 0; i < len(p.ErrorMsgs); i++ {
		log.Println(p.ErrorMsgs[i])
	}
	assert.Equal(t, true, res)
	assert.Equal(t, true, p.BoolValue)
}

func TestUnEqualsDoubleInt(t *testing.T) {
	p := SetupParser()
	p.Debug = true
	res := p.Parse("50.2 * ( 10 + 100 ) + 0.2 != 5522 - 10 + 10 + 2")
	for i := 0; i < len(p.ErrorMsgs); i++ {
		log.Println(p.ErrorMsgs[i])
	}
	assert.Equal(t, true, res)
	assert.Equal(t, true, p.BoolValue)
}

func TestUnEqualsIntDouble(t *testing.T) {
	p := SetupParser()
	p.Debug = true
	res := p.Parse("50 * ( 10 + 100 )  != 5522 - 10.5 + 10.5 + 0.1")
	for i := 0; i < len(p.ErrorMsgs); i++ {
		log.Println(p.ErrorMsgs[i])
	}
	assert.Equal(t, true, res)
	assert.Equal(t, true, p.BoolValue)
}

func TestUnEqualsIntInt(t *testing.T) {
	p := SetupParser()
	p.Debug = true
	res := p.Parse("50 * ( 10 + 100 )  != 5522 - 10 + 10 + 10")
	for i := 0; i < len(p.ErrorMsgs); i++ {
		log.Println(p.ErrorMsgs[i])
	}
	assert.Equal(t, true, res)
	assert.Equal(t, true, p.BoolValue)
}

/////////////////////////////
//Greater EQUALS THAN

func TestGreaterEqualsThanIntInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1 >= 100 + 2"))
	assert.Equal(t, false, p.BoolValue)
}
func TestGreaterEqualsThanDoubleDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1.5 * 3.1 >= 100 + 2.22 + 222.33"))
	assert.Equal(t, false, p.BoolValue)
}
func TestGreaterEqualsThanDoubleInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + (1.5 * 3.1) >= 100 + 2"))
	assert.Equal(t, true, p.BoolValue)
}
func TestGreaterEqualsThanIntDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1 >= 100 + 1.5 * 3.1"))
	assert.Equal(t, false, p.BoolValue)
}

// ////////////////////////////////////////////////
// Greater THAN
func TestGreaterDoubleInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1.11 > 100 + 2"))
	assert.Equal(t, false, p.BoolValue)
}

func TestGreaterDoubleDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1.11 > 100 + 2.22 * 33.3"))
	assert.Equal(t, false, p.BoolValue)
}

func TestGreaterIntInt(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 300 > 50 + 100 + 2"))
	assert.Equal(t, true, p.BoolValue)
}

func TestGreaterIntDouble(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 1 > 100 + 2.323 * 1.31"))
	assert.Equal(t, false, p.BoolValue)
}
