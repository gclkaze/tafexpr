package variablecontext

import (
	"fmt"
	"log"
	"strings"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"

	mine "github.com/gclkaze/tafexpr/stackvalue"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/parameters"
)

type JSONObjectMockVariableContext struct {
	variableMap map[string]globals.JSONObjectGen
	Verbose     bool
}

func (vc *JSONObjectMockVariableContext) SetValue(p string, v globals.JSONObjectGen) {
	vc.variableMap[p] = v
}

func (vc *JSONObjectMockVariableContext) FreeVariable(s string) {
	_, ok := vc.variableMap[s]
	if !ok {
		return
	}
	delete(vc.variableMap, s)
}
func (vc *JSONObjectMockVariableContext) SetParameter(s string, v globals.ParameterValue) {

}
func (vc *JSONObjectMockVariableContext) GetLength() int {
	return len(vc.variableMap)
}

func (vc *JSONObjectMockVariableContext) GetVariable(s string) parameters.IVariableParameterValue {
	return nil
}
func (vc *JSONObjectMockVariableContext) SetVariable(s string, v parameters.IVariableParameterValue) {

}

func (vc *JSONObjectMockVariableContext) GetVariableValue(s string, secretAware bool) (v stackvalue.StackValue, err error) {
	val, ok := vc.variableMap[s]
	if !ok {
		return nil, fmt.Errorf("didn't find elem with key %s", s)
	}
	return mine.NewJSONStackValue(val), nil
}

func (vc *JSONObjectMockVariableContext) Init(isVerbose bool) {
	vc.Verbose = isVerbose
	vc.variableMap = map[string]globals.JSONObjectGen{}
}

func (vc *JSONObjectMockVariableContext) GetVariableIntValue(s string) (res float64, err error) {
	val, ok := vc.variableMap[s]
	if !ok {
		log.Println("Didn't find element at " + s)
		panic("Didn't find element at " + s)
	}
	if isEmptyJSONValue(val) {
		return 0, nil
	}
	return 1, nil
}

// isEmptyJSONValue reports whether a JSONObjectGen value should be
// treated as "empty" for truthy/falsy purposes: nil, or an empty string.
func isEmptyJSONValue(val globals.JSONObjectGen) bool {
	if val == nil {
		return true
	}
	if s, ok := val.(string); ok && s == "" {
		return true
	}
	return false
}

func (vc *JSONObjectMockVariableContext) EvaluateJSONVariableIntValue(s string) (res globals.JSONObjectGen, err error) {
	return vc.GetVariableIntValue(s)
}

func (vc *JSONObjectMockVariableContext) EvaluateJSONVariable(s string, path string, secretAware bool) (stackvalue.StackValue, error) {
	if strings.HasPrefix(path, s) {
		return vc.GetVariableValue(path, secretAware)
	}
	return vc.GetVariableValue(s+"."+path, secretAware)
}
func (vc *JSONObjectMockVariableContext) ClearVariableContext() {

}
