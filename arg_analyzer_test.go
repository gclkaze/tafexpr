package tafexpr

import (
	"testing"

	"github.com/gclkaze/tafexpr/variablecontext"
	"github.com/stretchr/testify/assert"
)

func SetupTruthyVariableContext() variablecontext.IVariableContext {
	v := &variablecontext.TruthyVariablecontext{}
	v.Init(true)

	var context variablecontext.IVariableContext = v
	return context
}

func TestAnalyzeScalarVariableStatement(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupTruthyVariableContext()

	assert.Equal(t, true, p.Analyze("$x + $y + $z + $l"))
	exprs := p.VariableExpressions
	assert.Equal(t, 4, len(exprs))
	assert.Equal(t, "$x", exprs[0].VarName)
	assert.Equal(t, "$y", exprs[1].VarName)
	assert.Equal(t, "$z", exprs[2].VarName)
	assert.Equal(t, "$l", exprs[3].VarName)
}

func TestAnalyzeScalarVariableStatementWithConstants(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupTruthyVariableContext()

	assert.Equal(t, true, p.Analyze("1 - 0 + $x * 2 + ($y-3) + $z + $l + 100 -2 "))
	exprs := p.VariableExpressions
	assert.Equal(t, 4, len(exprs))
	assert.Equal(t, "$x", exprs[0].VarName)
	assert.Equal(t, "$y", exprs[1].VarName)
	assert.Equal(t, "$z", exprs[2].VarName)
	assert.Equal(t, "$l", exprs[3].VarName)
}

func TestAnalyzeObjectVariables(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupTruthyVariableContext()

	assert.Equal(t, true, p.Analyze("$x.y + $y.z + $z.l + $l.m.a.b[$nn + 100 - $u.a.x.z[10]]"))
	exprs := p.VariableExpressions
	assert.Equal(t, 6, len(exprs))
	assert.Equal(t, "$x", exprs[0].VarName)
	assert.Equal(t, "$y", exprs[1].VarName)
	assert.Equal(t, "$z", exprs[2].VarName)
	assert.Equal(t, "$nn", exprs[3].VarName)
	assert.Equal(t, "$u", exprs[4].VarName)
	assert.Equal(t, "$l", exprs[5].VarName)
}

func TestAnalyzeErroneousExpression(t *testing.T) {
	p := SetupParser()
	p.VariableContext = SetupTruthyVariableContext()

	assert.Equal(t, false, p.Analyze("$x.y + asdsadasda + $y"))
}
