package variablecontext

import (
	"log"
	"strings"

	"github.com/gclkaze/tafexpr/stackvalue"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/parameters"
)

type TruthyVariablecontext struct {
	variableMap map[string]float64
	Verbose     bool
}

//var TheVariableContext *VariableContext

func (vc *TruthyVariablecontext) SetValue(p string, v float64) {
	vc.variableMap[p] = v
}
func (vc *TruthyVariablecontext) SetParameter(s string, v globals.ParameterValue) {

}

func (vc *TruthyVariablecontext) FreeVariable(s string) {
	_, ok := vc.variableMap[s]
	if !ok {
		return
	}
	delete(vc.variableMap, s)
}

func (vc *TruthyVariablecontext) GetLength() int {
	return len(vc.variableMap)
}

func (vc *TruthyVariablecontext) GetVariable(s string) parameters.IVariableParameterValue {
	return nil
}
func (vc *TruthyVariablecontext) SetVariable(s string, v parameters.IVariableParameterValue) {

}

func (vc *TruthyVariablecontext) GetVariableValue(s string, secretAware bool) (v stackvalue.StackValue, err error) {
	val, ok := vc.variableMap[s]
	if !ok {
		vc.SetValue(s, 1.0)
		return stackvalue.NewDoubleStackValue(val), nil
	}
	return stackvalue.NewDoubleStackValue(val), nil
}

func (vc *TruthyVariablecontext) Init(isVerbose bool) {
	vc.Verbose = isVerbose
	vc.variableMap = map[string]float64{}
}

func (vc *TruthyVariablecontext) GetVariableIntValue(s string) (res float64, err error) {
	val, ok := vc.variableMap[s]
	if !ok {
		log.Println("Didn't find element at " + s)
		panic("Didn't find element at " + s)
	}
	return val, nil
}

func (vc *TruthyVariablecontext) EvaluateJSONVariableIntValue(s string, path string) (res float64, err error) {
	return vc.GetVariableIntValue(s + "." + path)
}

func (vc *TruthyVariablecontext) EvaluateJSONVariable(s string, path string, secretAware bool) (stackvalue.StackValue, error) {
	//return stackvalue.NewIntegerStackValue(2), nil
	if strings.HasPrefix(path, s) {
		return vc.GetVariableValue(path, secretAware)
	}
	return vc.GetVariableValue(s+"."+path, secretAware)
}
func (vc *TruthyVariablecontext) ClearVariableContext() {

}
