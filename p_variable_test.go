package tafexpr

import (
	"testing"

	"github.com/gclkaze/tafexpr/variablecontext"
	"github.com/stretchr/testify/assert"
)

func TestVarArg(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$x", 13)

	assert.Equal(t, true, p.Parse("$x"))
	assert.Equal(t, 13, p.IntValue)
}

func TestListVarArgIndexLiteral(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$x[1][2][3].a[4].b.c.d[15]", 13)

	assert.Equal(t, true, p.Parse("$x[1][2][3].a[4].b.c.d[$i + 5]"))
	assert.Equal(t, 13, p.IntValue)
}

func TestListVarArgIndexLiteralArrayAccess(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 666)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$first[666]", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$sum", 10)

	assert.Equal(t, true, p.Parse(" $sum + $first[$i]"))
	assert.Equal(t, 20, p.IntValue)
}

func TestListVarArgIndexLiteralArrayAccessLiteral(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupMockVariableContext()
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$first[666].rating", 10)
	p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$sum", 10)

	assert.Equal(t, true, p.Parse(" $sum + $first[666].rating"))
	assert.Equal(t, 20, p.IntValue)
}

/*
func TestListVarArgIndexExpression(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 200 +3090"))
	assert.Equal(t, 100+200+3090, p.IntValue)
}

func TestListVarArgComposite(t *testing.T) {
	p := SetupParser()
	assert.Equal(t, true, p.Parse("100 + 200 +3090"))
	assert.Equal(t, 100+200+3090, p.IntValue)
}*/
