package tafexpr

import (
	"testing"

	"github.com/gclkaze/tafexpr/variablecontext"
	"github.com/stretchr/testify/assert"
)

func TestStringInterpolation(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 10)

	assert.Equal(t, true, p.Interpolate("mike{{$i}}and{{$j}}and{{$i}}"))
	assert.Equal(t, "mike60and10and60", p.StringValue)
	assert.Equal(t, true, p.InterpolationApplied())
}

func TestStringInterpolationMulVars(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$k", 12)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$l", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$m", 110)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$n", 10)

	assert.Equal(t, true, p.Interpolate("mike{{$i}}and{{$j}}and{{$k}}and{{$l}}and{{$m}}and{{$n}}"))
	assert.Equal(t, "mike60and10and12and10and110and10", p.StringValue)
	assert.Equal(t, true, p.InterpolationApplied())
}

func TestStringInterpolationDidnApply(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 10)

	assert.Equal(t, true, p.Interpolate("mike{{$ij}}and{{$ji}}"))
	assert.Equal(t, false, p.InterpolationApplied())
	assert.Equal(t, "mike{{$ij}}and{{$ji}}", p.StringValue)
}

func TestStringInterpolationDidnApplyInStringJSONObject(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()

	assert.Equal(t, true, p.Interpolate("\"{\"a\":{\"b\":{\"c\":666}}}\""))
	assert.Equal(t, false, p.InterpolationApplied())
	assert.Equal(t, "\"{\"a\":{\"b\":{\"c\":666}}}\"", p.StringValue)
}

func TestStringInterpolationApplyInStringJSONObject(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
	assert.Equal(t, true, p.Interpolate("{\"a\":{\"b\":{\"c\":{{$i}}}}}"))
	assert.Equal(t, true, p.InterpolationApplied())
	assert.Equal(t, "{\"a\":{\"b\":{\"c\":60}}}", p.StringValue)
}

func TestStringInterpolationPartial(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 10)

	assert.Equal(t, true, p.Interpolate("mike{{$ij}}and{{$j}}"))
	assert.Equal(t, true, p.InterpolationApplied())
	assert.Equal(t, "mike{{$ij}}and10", p.StringValue)
}

func TestStringInterpolationJSONKey(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 1)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 2)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$k", 3)

	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$value", 10)

	assert.Equal(t, true, p.Interpolate("{   \"{{$i + 1}}\" : (value * 2) ,   \"{{$j + 2}}\" : (value * 3),   \"{{$k + 3}}\" : (value * 4)               }"))
	assert.Equal(t, true, p.InterpolationApplied())
	assert.Equal(t, "{   \"2\" : (value * 2) ,   \"4\" : (value * 3),   \"6\" : (value * 4)               }", p.StringValue)
}

//
