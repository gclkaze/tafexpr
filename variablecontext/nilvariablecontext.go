package variablecontext

import (
	"errors"

	"github.com/gclkaze/tafexpr/stackvalue"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/parameters"
)

type NilVariableContext struct {
	variableMap map[string]float64
	Verbose     bool
}

//var TheVariableContext *VariableContext

func (vc *NilVariableContext) SetValue(p string, v float64) {
	vc.variableMap[p] = v
}

func (vc *NilVariableContext) Init(isVerbose bool) {
	vc.Verbose = isVerbose
	vc.variableMap = map[string]float64{}
}
func (vc *NilVariableContext) GetVariable(s string) *parameters.VariableParameterValue {
	return nil
}
func (vc *NilVariableContext) SetVariable(s string, v *parameters.VariableParameterValue) {

}
func (vc *NilVariableContext) SetParameter(s string, v globals.ParameterValue) {

}
func (vc *NilVariableContext) FreeVariable(s string) {
	_, ok := vc.variableMap[s]
	if !ok {
		return
	}
	delete(vc.variableMap, s)

}

func (vc *NilVariableContext) GetLength() int {
	return len(vc.variableMap)
}

func (vc *NilVariableContext) GetVariableValue(s string, secretAware bool) (v stackvalue.StackValue, err error) {
	return nil, nil
}

func (vc *NilVariableContext) GetVariableIntValue(s string) (res float64, err error) {
	return 0, errors.New("Undeclared variable \"" + s + "\"")
}

func (vc *NilVariableContext) EvaluateJSONVariableIntValue(s string, path string) (res float64, err error) {
	return vc.GetVariableIntValue(s + "." + path)
}
func (vc *NilVariableContext) EvaluateJSONVariable(s string, path string, secretAware bool) (stackvalue.StackValue, error) {
	return stackvalue.NewNullStackValue(), nil
}
func (vc *NilVariableContext) ClearVariableContext() {

}
