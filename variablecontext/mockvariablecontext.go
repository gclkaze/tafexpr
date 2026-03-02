package variablecontext

import (
	"fmt"
	"log"
	"strings"

	"github.com/gclkaze/tafexpr/stackvalue"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/parameters"
)

type MockVariableContext struct {
	variableMap map[string]float64
	Verbose     bool
}

//var TheVariableContext *VariableContext

func (vc *MockVariableContext) SetValue(p string, v float64) {
	vc.variableMap[p] = v
}

func (vc *MockVariableContext) FreeVariable(s string) {
	_, ok := vc.variableMap[s]
	if !ok {
		return
	}
	delete(vc.variableMap, s)
}
func (vc *MockVariableContext) SetParameter(s string, v globals.ParameterValue) {

}
func (vc *MockVariableContext) GetLength() int {
	return len(vc.variableMap)
}

func (vc *MockVariableContext) GetVariable(s string) parameters.IVariableParameterValue {
	return nil
}
func (vc *MockVariableContext) SetVariable(s string, v parameters.IVariableParameterValue) {

}

func (vc *MockVariableContext) GetVariableValue(s string, secretAware bool) (v stackvalue.StackValue, err error) {
	val, ok := vc.variableMap[s]
	if !ok {
		return nil, fmt.Errorf("didn't find elem with key %s", s)
	}
	return stackvalue.NewDoubleStackValue(val), nil
}

func (vc *MockVariableContext) Init(isVerbose bool) {
	vc.Verbose = isVerbose
	vc.variableMap = map[string]float64{}
}

func (vc *MockVariableContext) GetVariableIntValue(s string) (res float64, err error) {
	val, ok := vc.variableMap[s]
	if !ok {
		log.Println("Didn't find element at " + s)
		panic("Didn't find element at " + s)
	}
	return val, nil
}

func (vc *MockVariableContext) EvaluateJSONVariableIntValue(s string, path string) (res float64, err error) {
	return vc.GetVariableIntValue(s + "." + path)
}

func (vc *MockVariableContext) EvaluateJSONVariable(s string, path string, secretAware bool) (stackvalue.StackValue, error) {
	//return stackvalue.NewIntegerStackValue(2), nil
	if strings.HasPrefix(path, s) {
		return vc.GetVariableValue(path, secretAware)
	}
	return vc.GetVariableValue(s+"."+path, secretAware)
}
func (vc *MockVariableContext) ClearVariableContext() {

}
