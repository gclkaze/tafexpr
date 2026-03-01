package tafexpr

import (
	"testing"

	"github.com/gclkaze/tafexpr/variablecontext"

	"github.com/stretchr/testify/assert"
)

func TestSimpleVar(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)

	assert.Equal(t, true, p.ParsePureVariableStorageDestination("$i"))
	v := p.VariableStorageDestination
	assert.Equal(t, "$i", v.VarExpression)
	assert.Equal(t, "$i", v.VarName)
}

func TestSimpleVarIncr(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)

	assert.Equal(t, false, p.ParsePureVariableStorageDestination("$i + 1"))
}

func TestSimpleVarIncrAssign(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)

	assert.Equal(t, false, p.ParseVariableStorageDestination("$i = 1"))
}

func TestLogicalExpressionNotVarExpression(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)

	assert.Equal(t, false, p.ParsePureVariableStorageDestination("$i == 5"))

}

func TestComplexVarExpressionIsNotDest(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$d[60][10]", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$json.a.b.c[10].l.m", 60.05)

	assert.Equal(t, false, p.ParsePureVariableStorageDestination("$json           .        a     .  b. c[   $d[ $i  ][  $j  ]]. l.m + 200"))

}

func TestComplexVarExpressionIsDest(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$j", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$d[70][10]", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$json.a.b.c[10].l.m", 60.05)

	assert.Equal(t, true, p.ParsePureVariableStorageDestination("$json           .        a     .  b. c[   $d[ $i + 10  ][  $j  ]]. l.m "))

}
