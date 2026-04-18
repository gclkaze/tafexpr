package tafexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gclkaze/evalang-globals/utils"
)

func TestJsonObject(t *testing.T) {
	p := SetupParser()
	s := utils.ReadFile("jsonexprs\\simple.json")
	res := p.Parse(*s)
	assert.Equal(t, true, res)
	assert.NotNil(t, p.JSONValue)
}

func TestJsonArray(t *testing.T) {
	p := SetupParser()
	s := utils.ReadFile("jsonexprs\\array.json")
	res := p.Parse(*s)
	assert.Equal(t, true, res)
	assert.NotNil(t, p.JSONArray)
}

func TestJsonArrayScalar(t *testing.T) {
	p := SetupParser()
	expr := "[1,2,3,4]"
	res := p.Parse(expr)
	assert.Equal(t, true, res)
	assert.NotNil(t, p.JSONArray)
}

func TestJsonComplex(t *testing.T) {
	p := SetupParser()
	s := utils.ReadFile("jsonexprs\\complex.json")
	res := p.Parse(*s)
	assert.Equal(t, true, res)
	assert.NotNil(t, p.JSONValue)
}

func TestJsonToString(t *testing.T) {
	p := SetupParser()
	//	s := utils.ReadFile("jsonexprs\\complex.json")
	expr := "{\"2\":3}.toString"
	res := p.Parse(expr)
	assert.Equal(t, true, res)
	assert.NotNil(t, p.StringValue)
}

func TestJJsonToString(t *testing.T) {
	p := SetupParser()
	//	s := utils.ReadFile("jsonexprs\\complex.json")
	expr := "{\"a\":{\"b\":{\"c\":666}}}.toString"
	res := p.Parse(expr)
	assert.Equal(t, true, res)
	assert.NotNil(t, p.StringValue)
}

func TestJSONArrayLength(t *testing.T) {
	p := SetupParser()
	//	s := utils.ReadFile("jsonexprs\\complex.json")
	expr := "[1,2,3].length == [4,5,6].length"
	res := p.Parse(expr)
	assert.Equal(t, true, res)
	assert.NotNil(t, p.BoolValue)
}
