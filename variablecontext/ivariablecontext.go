package variablecontext

import (
	"github.com/gclkaze/tafexpr/stackvalue"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/parameters"
)

type IVariableContext interface {
	Init(isVerbose bool)

	GetVariable(s string) *parameters.VariableParameterValue
	SetVariable(s string, v *parameters.VariableParameterValue)
	SetParameter(s string, v globals.ParameterValue)

	GetVariableValue(s string, secretAware bool) (stackvalue.StackValue, error)
	GetVariableIntValue(s string) (float64, error)
	EvaluateJSONVariable(s string, path string, secretAware bool) (stackvalue.StackValue, error)
	GetLength() int
	FreeVariable(s string)

	ClearVariableContext()
}
