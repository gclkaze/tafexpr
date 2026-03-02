package variablecontext

import (
	"github.com/gclkaze/tafexpr/stackvalue"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/parameters"
)

type StringVariableContext struct {
}

//var TheVariableContext *VariableContext

func (vc *StringVariableContext) Init(isVerbose bool) {
}

func (vc *StringVariableContext) GetVariable(s string) *parameters.IVariableParameterValue {
	return nil
}
func (vc *StringVariableContext) FreeVariable(s string) {

}
func (vc *StringVariableContext) SetVariable(s string, v *parameters.IVariableParameterValue) {

}

func (vc *StringVariableContext) SetParameter(s string, v globals.ParameterValue) {

}
func (vc *StringVariableContext) GetLength() int {
	return 0
}

func (vc *StringVariableContext) GetVariableIntValue(s string) (res float64, ok error) {
	//log.Println(s, "=>", 1)
	return 1, nil
}
func (vc *StringVariableContext) GetVariableValue(s string, secretAware bool) (v stackvalue.StackValue, err error) {
	return stackvalue.NewStringStackValue("value"), nil
}

func (vc *StringVariableContext) EvaluateJSONVariableIntValue(s string, path string) (res float64, ok error) {
	//log.Println(s, path, "=>", 2)
	return 2, nil
}

func (vc *StringVariableContext) EvaluateJSONVariable(s string, path string, secretAware bool) (stackvalue.StackValue, error) {
	return stackvalue.NewStringStackValue("value"), nil
}

func (vc *StringVariableContext) ClearVariableContext() {

}
