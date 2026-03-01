package tafexpr

import (
	"testing"

	"github.com/gclkaze/evalang/evalangparser/utils"

	"github.com/stretchr/testify/assert"
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
