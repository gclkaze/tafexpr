package variablecontext

import (
	"github.com/gclkaze/tafexpr/stackvalue"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/parameters"
)

type VariableContext struct {
}

//var TheVariableContext *VariableContext

func (vc *VariableContext) Init(isVerbose bool) {
}

func (vc *VariableContext) GetVariable(s string) parameters.IVariableParameterValue {
	return nil
}
func (vc *VariableContext) FreeVariable(s string) {

}
func (vc *VariableContext) SetVariable(s string, v parameters.IVariableParameterValue) {

}

func (vc *VariableContext) SetParameter(s string, v globals.ParameterValue) {

}

func (vc *VariableContext) GetLength() int {
	return 0
}

func (vc *VariableContext) GetVariableIntValue(s string) (res float64, ok error) {
	//log.Println(s, "=>", 1)
	return 1, nil
}
func (vc *VariableContext) GetVariableValue(s string, secretAware bool) (v stackvalue.StackValue, err error) {
	return stackvalue.NewIntegerStackValue(1), nil
}

func (vc *VariableContext) EvaluateJSONVariableIntValue(s string, path string) (res float64, ok error) {
	//log.Println(s, path, "=>", 2)
	return 2, nil
}

func (vc *VariableContext) EvaluateJSONVariable(s string, path string, secretAware bool) (stackvalue.StackValue, error) {
	return stackvalue.NewIntegerStackValue(2), nil
}

func (vc *VariableContext) ClearVariableContext() {

}
