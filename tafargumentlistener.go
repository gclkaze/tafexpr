package tafexpr

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/gclkaze/tafexpr/parser"
	"github.com/gclkaze/tafexpr/stackvalue"
	"github.com/gclkaze/tafexpr/stackvalue/logicalstackvalue"
	"github.com/gclkaze/tafexpr/stackvalue/numberstackvalue"
	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"

	"github.com/gclkaze/tafexpr/variablecontext"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/utils"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

type InterpolationToken struct {
	token    string
	startPos int
	endPos   int
	value    stackvalue.StackValue
}

func (i InterpolationToken) GetToken() string {
	return i.token
}

func (i InterpolationToken) GetStart() int {
	return i.startPos
}

func (i InterpolationToken) GetEnd() int {
	return i.endPos
}

func (i InterpolationToken) GetValue() stackvalue.StackValue {
	return i.value
}

type VarExpression struct {
	VarExpression string
	VarName       string
}

func (v VarExpression) IsPath() bool {
	return v.VarExpression != v.VarName
}

type TAFArgumentListener struct {
	*parser.BaseTafexprListener
	//	stack           []float64
	Escaped                          bool
	stack                            *stackvalue.MultiTypeStack
	IsFloat                          bool
	OnError                          bool
	ErrorMsgs                        []TAFParserArgumentError
	VariableContext                  variablecontext.IVariableContext
	CurrentPath                      string
	Scope                            int
	Index                            IndexStack
	Debug                            bool
	onVariableStorageDeclarationMode bool
	lastVarExpression                VarExpression

	detectedVariableExpressions []VarExpression

	arrayScope           int
	jsonStack            []*globals.JSONStruct
	itokens              []*InterpolationToken
	iexpr                string
	interpolationApplied bool
	onInterpolationMode  bool
	LastExit             string

	secretAware bool
}

type TAFParserArgumentError struct {
	Msg  error
	Type tafargumentlistenererrortypes.TAFArgumentListenerError
}

func (l TAFParserArgumentError) ToString() string {
	return fmt.Sprintf("%v", l.Msg)
}

func (l TAFArgumentListener) StackIsEmpty() bool {
	return l.stack == nil || l.stack.IsEmpty()
}
func (s *TAFArgumentListener) EnterTaf_expression(ctx *parser.Taf_expressionContext) {
	s.stack = stackvalue.NewMultiTypeStack()
}

func (s *TAFArgumentListener) GetLastVariableDestinationExpresison() VarExpression {
	return s.lastVarExpression
}

func (s *TAFArgumentListener) GetReferencedVariables() []VarExpression {
	return s.detectedVariableExpressions
}

func NewInterpolationToken(start int, end int, token string, v stackvalue.StackValue) *InterpolationToken {
	return &InterpolationToken{startPos: start, endPos: end, token: token, value: v}
}

func (i *InterpolationToken) SetStackValue(v stackvalue.StackValue) {
	i.value = v
}

func (l TAFArgumentListener) InterpolationApplied() bool {
	return l.interpolationApplied
}

func (l *TAFArgumentListener) SetSecretAware(secretAware bool) {
	l.secretAware = secretAware
}

func (l *TAFArgumentListener) push(i float64) {
	//l.stack = append(l.stack, i)
	if l.OnError {
		return
	}

	if isIntegral(i) {
		l.stack.PushInteger(int(i))
	} else {
		l.stack.PushDouble(i)

	}
	log.Debug("pushed ", l.stack)

	if !l.IsFloat {
		l.IsFloat = !isIntegral(i)
	}
}

func (l *TAFArgumentListener) pushValue(s stackvalue.StackValue) {
	if l.OnError {
		return
	}

	l.stack.Push(s)
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func (l *TAFArgumentListener) _pop() any {
	if l.OnError {
		return -1
	}

	last := l.stack.Pop()
	if last.GetType() == stackvalue.INTEGER {
		i, ok := last.(*stackvalue.IntegerStackValue)
		if !ok {
			panic(666)
		}
		return i.GetValue()
	}
	if last.GetType() == stackvalue.DOUBLE {
		i, ok := last.(*stackvalue.DoubleStackValue)
		if !ok {
			panic(777)
		}
		return i.GetValue()
	}
	panic(888)
	return -1
}

func (l *TAFArgumentListener) popStack() stackvalue.StackValue {
	if l.OnError {
		return nil
	}

	return l.stack.Pop()
}

func (l *TAFArgumentListener) popFloat() float64 {
	if l.OnError {
		return -1
	}

	last := l.stack.Pop()
	if last.GetType() == stackvalue.DOUBLE {
		i, ok := last.(*stackvalue.DoubleStackValue)
		if !ok {
			panic(433)
		}
		f := i.GetValue()
		const prec = 200
		a := new(big.Float).SetPrec(prec).SetFloat64(f)
		result, _ := strconv.ParseFloat(a.String(), 64)
		return result
	} else {
		i, ok := last.(*stackvalue.IntegerStackValue)
		if !ok {
			panic(4344)
		}
		return float64(i.GetValue())
	}
	panic(444)
	return -1
}

func (l TAFArgumentListener) castToFloat(unkn any) (res float64, err error) {
	switch i := unkn.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int:
		return float64(i), nil
	}
	return -1, errors.New("couldnt cast it")
}

func (l *TAFArgumentListener) ExitHandleFindByXPATH(c *parser.HandleFindByXPATHContext) {
	l.LastExit = "FindByXPATH"
	if l.OnError {
		return
	}
	v := l.popStack()
	xpath := v.ToString()
	if v.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("XPATH needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of findByXPATH should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	varContent := theVar.ToString()
	result, err := utils.ParseHTMLByXPATH(varContent, xpath)

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	var payload []interface{}
	for i := 0; i < len(result); i++ {
		payload = append(payload, result[i])
	}

	l.stack.Push(stackvalue.NewJSONArrayStackValue(payload))
}

func (l *TAFArgumentListener) ExitHandleFindOneByXPATH(c *parser.HandleFindOneByXPATHContext) {
	l.LastExit = "FindOneByXPATH"
	if l.OnError {
		return
	}
	v := l.popStack()
	xpath := v.ToString()
	if v.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("XPATH needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of findByXPATH should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	varContent := theVar.ToString()
	result, err := utils.ParseHTMLByXPATHAndGetOne(varContent, xpath)

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushString(result)
}

func (l *TAFArgumentListener) findOneXPATHString(expr string) (res string, err error) {
	v := l.popStack()
	xpath := v.ToString()
	if v.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("XPATH needs to be a STRING TYPE: %s", expr), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of findOne<VALUE TYPE>ByXPATH should be of STRING TYPE: %s", expr), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	varContent := theVar.ToString()
	result, err := utils.ParseHTMLByXPATHAndGetOne(varContent, xpath)

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	return result, err
}

func (l *TAFArgumentListener) ExitHandleFindOneStringByXPATH(c *parser.HandleFindOneStringByXPATHContext) {
	l.LastExit = "FindOneStringByXPATH"
	if l.OnError {
		return
	}
	s, err := l.findOneXPATHString(c.GetText())
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushString(s)
}
func (l *TAFArgumentListener) ExitHandleFindOneIntegerByXPATH(c *parser.HandleFindOneIntegerByXPATHContext) {
	l.LastExit = "FindOneIntegerByXPATH"
	if l.OnError {
		return
	}
	s, err := l.findOneXPATHString(c.GetText())
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushInteger(i)

}
func (l *TAFArgumentListener) ExitHandleFindOneDoubleByXPATH(c *parser.HandleFindOneDoubleByXPATHContext) {
	l.LastExit = "FindOneDoubleByXPATH"
	if l.OnError {
		return
	}
	s, err := l.findOneXPATHString(c.GetText())
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushDouble(i)

}
func (l *TAFArgumentListener) ExitHandleFindOneBooleanByXPATH(c *parser.HandleFindOneBooleanByXPATHContext) {
	l.LastExit = "FindOneBooleanByXPATH"
	if l.OnError {
		return
	}
	s, err := l.findOneXPATHString(c.GetText())
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	i, err := strconv.ParseBool(s)
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushBool(i)
}

func (l *TAFArgumentListener) ExitHandleExtractOneByREGEX(c *parser.HandleExtractOneByREGEXContext) {
	l.LastExit = "ExtractOneByREGEX"
	if l.OnError {
		return
	}
	v := l.popStack()
	xpath := v.ToString()
	if v.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("REGEX needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of extractOneByREGEX should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	varContent := theVar.ToString()
	result, err := utils.ExtractOneFromRegex(varContent, xpath)

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushString(result)
}

func (l *TAFArgumentListener) ExitHandleRandomDoubleInRange(c *parser.HandleRandomDoubleInRangeContext) {
	l.LastExit = "RandomDoubleInRange"
	vmax := l.popStack()
	if vmax.GetType() != stackvalue.INTEGER && vmax.GetType() != stackvalue.DOUBLE {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("MAX needs to be an INTEGER or DOUBLE TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	vmaxValue, err := vmax.ToDouble()
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("Couldn't cast MAX to DOUBLE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	vmin := l.popStack()
	if vmin.GetType() != stackvalue.INTEGER && vmin.GetType() != stackvalue.DOUBLE {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("MIN needs to be an INTEGER or DOUBLE TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	vminValue, err := vmin.ToDouble()
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("Couldn't cast MIN to DOUBLE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	v := utils.GetRandomDoubleInRange(vminValue, vmaxValue)
	l.stack.PushDouble(v)

}

func (l *TAFArgumentListener) ExitHandleContainsString(c *parser.HandleContainsStringContext) {
	l.LastExit = "ContainsString"
	if l.OnError {
		return
	}
	param := l.popStack()
	if param.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("containsString parameter needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	searchString := param.ToString()
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of containsString should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	result, err := utils.ContainsString(theVar.ToString(), searchString)

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushBool(result)
}

func (l *TAFArgumentListener) ExitHandleStartsWith(c *parser.HandleStartsWithContext) {
	l.LastExit = "StartsWith"
	if l.OnError {
		return
	}
	param := l.popStack()
	if param.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("startsWith parameter needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	searchString := param.ToString()
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of startsWith should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	result, err := utils.StarsWith(theVar.ToString(), searchString)

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushBool(result)
}

func (l *TAFArgumentListener) ExitHandleEndsWith(c *parser.HandleEndsWithContext) {
	l.LastExit = "EndsWith"
	if l.OnError {
		return
	}
	param := l.popStack()
	if param.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("endsWith parameter needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	searchString := param.ToString()
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of endsWith should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	result, err := utils.EndsWith(theVar.ToString(), searchString)

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushBool(result)
}

func (l *TAFArgumentListener) ExitHandleTrim(c *parser.HandleTrimContext) {
	l.LastExit = "Trim"
	if l.OnError {
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of endsWith should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	result, err := utils.Trim(theVar.ToString())

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushString(result)
}

func (l *TAFArgumentListener) ExitHandleTrimLeft(c *parser.HandleTrimLeftContext) {
	l.LastExit = "TrimLeft"
	if l.OnError {
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of trimLeft should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	result, err := utils.TrimLeft(theVar.ToString())

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushString(result)
}

func (l *TAFArgumentListener) ExitHandleTrimRight(c *parser.HandleTrimRightContext) {
	l.LastExit = "TrimRight"
	if l.OnError {
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of trimRight should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	result, err := utils.TrimRight(theVar.ToString())

	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushString(result)
}

/*
| expression '.' 'trimLeft'#HandleTrimLeft
| expression '.' 'trimRight'#HandleTrimRight
| expression '.' 'trim'#HandleTrim
*/

func (l *TAFArgumentListener) ExitHandleToString(c *parser.HandleToStringContext) {
	l.LastExit = "ToString"
	if l.OnError {
		return
	}
	v := l.popStack()
	s := v.ToString()
	l.stack.PushString(s)
}

func (l *TAFArgumentListener) ExitHandleToInteger(c *parser.HandleToIntegerContext) {
	l.LastExit = "ToInteger"
	if l.OnError {
		return
	}
	v := l.popStack()
	s, err := v.ToInteger()
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushInteger(int(s))
}
func (l *TAFArgumentListener) ExitHandleToBoolean(c *parser.HandleToBooleanContext) {
	l.LastExit = "ToBoolean"
	if l.OnError {
		return
	}
	v := l.popStack()
	s, err := v.ToBoolean()
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushBool(s)

}
func (l *TAFArgumentListener) ExitHandleToDouble(c *parser.HandleToDoubleContext) {
	l.LastExit = "ToDouble"
	if l.OnError {
		return
	}
	v := l.popStack()
	s, err := v.ToDouble()
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}

	l.stack.PushDouble(s)
}

func (l *TAFArgumentListener) ExitHandleReplaceAllStringOccurrences(c *parser.HandleReplaceAllStringOccurrencesContext) {
	l.LastExit = "ExtractOneByREGEX"
	if l.OnError {
		return
	}
	v := l.popStack()
	newValue := v.ToString()
	if v.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("new replacement needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	v = l.popStack()
	oldValue := v.ToString()
	if v.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("old value needs to be a STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	theVar := l.popStack()
	if theVar.GetType() != stackvalue.STRING {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("LHS of replaceAllStringOccurrences should be of STRING TYPE: %s", c.GetText()), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	varContent := theVar.ToString()
	result := utils.ReplaceAllString(varContent, oldValue, newValue)
	l.stack.PushString(result)
}

func (l *TAFArgumentListener) ExitHandleLength(c *parser.HandleLengthContext) {
	l.LastExit = "Length"
	if l.OnError {
		return
	}
	v := l.popStack()
	length, err := v.Length()
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return
	}
	l.stack.PushInteger(length)
	//fmt.Printf("v: %v\n", v)
}

// LESSER_THAN|LESSER_THAN_EQUAL|EQUAL|GREATER_THAN|GREATER_THAN_EQUAL|UNEQUAL
func (l *TAFArgumentListener) ExitLogicalOperation(c *parser.LogicalOperationContext) {
	l.LastExit = "Logical"

	if l.OnError {
		return
	}
	right, left := l.popStack(), l.popStack() //l.pop(), l.pop()
	leftNum, ok := left.(logicalstackvalue.LogicalStackValue)
	if !ok {
		if !l.stack.OnError() {
			panic(666)
		}
		l.OnError = true
		return
	}

	switch c.GetOp().GetTokenType() {
	case parser.TafexprParserLESSER_THAN:
		res, errType, err := leftNum.LesserThan(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
		//l.push(v)
	case parser.TafexprParserLESSER_THAN_EQUAL:
		res, errType, err := leftNum.LesserThanEqual(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	case parser.TafexprParserEQUAL:
		res, errType, err := leftNum.Equal(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	case parser.TafexprParserUNEQUAL:
		res, errType, err := leftNum.Unequal(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	case parser.TafexprParserGREATER_THAN:
		res, errType, err := leftNum.GreaterThan(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	case parser.TafexprParserGREATER_THAN_EQUAL:
		res, errType, err := leftNum.GreaterThanEqual(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	}
}

func (l *TAFArgumentListener) ExitHandleLogicalNegation(c *parser.HandleLogicalNegationContext) {
	l.LastExit = "Logical Negation"

	if l.OnError {
		return
	}
	current := l.popStack()
	logValue, ok := current.(logicalstackvalue.LogicalStackValue)
	if !ok {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("not a logical expression: %s ", c.GetText()), Type: tafargumentlistenererrortypes.CONVERSION_ERROR})
		return
	}
	negated, st, err := logValue.Not()
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: st})
		return
	}
	l.stack.Push(negated)
}

func (l *TAFArgumentListener) ExitHandleNegation(c *parser.HandleNegationContext) {
	l.LastExit = "Negation"
	if l.OnError {
		return
	}
	current := l.popStack()
	numValue, ok := current.(numberstackvalue.NumberStackValue)
	if !ok {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("not a numerical expression: %s ", c.GetText()), Type: tafargumentlistenererrortypes.CONVERSION_ERROR})
		return
	}
	neg := stackvalue.NewIntegerStackValue(-1)
	negated, st, err := numValue.Mul(neg)
	if err != nil {
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: st})
		return
	}

	l.stack.Push(negated)
}

func (l *TAFArgumentListener) ExitHandleLogical(c *parser.HandleLogicalContext) {
	l.LastExit = "Logical"

	if l.OnError {
		return
	}
	right, left := l.popStack(), l.popStack() //l.pop(), l.pop()
	leftNum, ok := left.(logicalstackvalue.LogicalStackValue)
	if !ok {
		panic(666)
	}
	switch c.GetOp().GetTokenType() {

	case parser.TafexprParserLOGICAL_AND:
		res, errType, err := leftNum.And(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	case parser.TafexprParserLOGICAL_OR:
		res, errType, err := leftNum.Or(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	}
}

func (l *TAFArgumentListener) ExitMulDiv(c *parser.MulDivContext) {
	l.LastExit = "Mul Div"

	if l.OnError {
		return
	}
	right, left := l.popStack(), l.popStack() //l.pop(), l.pop()
	leftNum, ok := left.(numberstackvalue.NumberStackValue)
	if !ok {
		panic(666)
	}

	switch c.GetOp().GetTokenType() {
	case parser.TafexprParserMUL:
		res, errType, err := leftNum.Mul(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
		//l.push(v)
	case parser.TafexprParserDIV:
		res, errType, err := leftNum.Div(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	case parser.TafexprParserMOD:
		res, errType, err := leftNum.Mod(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return

	}
}

func (l *TAFArgumentListener) ExitAddSub(c *parser.AddSubContext) {
	l.LastExit = "Addsub"

	if l.OnError {
		return
	}
	right, left := l.popStack(), l.popStack()
	leftNum, ok := left.(numberstackvalue.NumberStackValue)
	if !ok {
		//panic(666)
		if !l.stack.OnError() {
			panic(666)
		}
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("left value is undeclared"), Type: tafargumentlistenererrortypes.UNDECLARED_VARIABLE})
		l.OnError = true
		return
	}

	switch c.GetOp().GetTokenType() {
	case parser.TafexprParserADD:
		res, errType, err := leftNum.Add(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
		//l.push(v)
	case parser.TafexprParserSUB:
		res, errType, err := leftNum.Sub(right)
		if errType == tafargumentlistenererrortypes.NONE {
			l.pushValue(res)
			return
		}
		l.OnError = true
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: err, Type: errType})
		return
	}

}

func (l *TAFArgumentListener) ExitNumber(c *parser.NumberContext) {
	l.LastExit = "Number"

	if l.OnError {
		return
	}
	i, err := strconv.Atoi(c.GetText())
	if err != nil {

		if l.Debug {
			log.Error(err.Error())
			panic(err.Error())
		} else {
			log.Error(err.Error())
			l.OnError = true
			return
		}
	}

	l.push(float64(i))
}

func (l *TAFArgumentListener) ExitDoubleValue(c *parser.DoubleValueContext) {
	l.LastExit = "Double"

	if l.OnError {
		return
	}
	i, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {

		if l.Debug {
			log.Error(err.Error())
			panic(err.Error())
		} else {
			log.Error(err.Error())
			l.OnError = true
			return
		}
	}

	l.IsFloat = true

	l.push(i)
}

func (l *TAFArgumentListener) ExitTaf_expression(c *parser.Taf_expressionContext) {
	//log.Debug("ExitTaf_expression" + c.GetText())
	l.Index.Clear()
}

func (l *TAFArgumentListener) EnterVar_expression(c *parser.Var_expressionContext) {
	varExpression := c.GetText()
	//log.Debug("EnterVar_expression " + varExpression)

	l.Index.SetExpression(varExpression)
}

func (l *TAFArgumentListener) EnterHandleVarExpression(c *parser.HandleVarExpressionContext) {
	//log.Debug("EnterHandleVarExpression")
	if l.Scope == 0 {
		l.CurrentPath = c.GetText()
	}
	l.Scope++
}

func (l *TAFArgumentListener) EnterHandleJJ(ctx *parser.HandleJJContext) {
	//strValue := ctx.GetText()
	//fmt.Println("WHAT??" + strValue)
}

func (l *TAFArgumentListener) ExitHandleString(ctx *parser.HandleStringContext) {
	l.LastExit = "String"

	strValue := ctx.GetText()
	/*	strValue = strings.TrimSuffix(strValue, "\"")
		strValue = strings.TrimPrefix(strValue, "\"")*/
	if l.Escaped {
		strValue = strings.TrimSuffix(strValue, "\\\"")
		strValue = strings.TrimPrefix(strValue, "\"")
	} else {
		strValue = strings.TrimSuffix(strValue, "\"")
		strValue = strings.TrimPrefix(strValue, "\"")
	}

	l.pushValue(stackvalue.NewStringStackValue(strValue))
}

func (l *TAFArgumentListener) ExitHandleNull(ctx *parser.HandleNullContext) {
	l.LastExit = "Null"

	l.pushValue(stackvalue.NewNullStackValue())
}

func (l *TAFArgumentListener) ExitHandleBool(c *parser.HandleBoolContext) {
	l.LastExit = "Bool"

	b := c.GetText()
	b = strings.ToLower(b)
	if strings.Compare(b, "true") == 0 {
		l.pushValue(stackvalue.NewBoolStackValue(true))
	} else if strings.Compare(b, "false") == 0 {
		l.pushValue(stackvalue.NewBoolStackValue(false))
	} else {
		panic("Unhandled boolean expression")
	}
}

func (l *TAFArgumentListener) ExitHandleVarExpression(c *parser.HandleVarExpressionContext) {
	l.LastExit = "VarExpression"

	//log.Debug("ExitHandleVarExpression")
	l.Scope--
	if l.Scope == 0 && !l.Index.IsEmpty() {
		log.Debug("l.Index:", l.Index)
		//we need to store index value to the stack
		log.Debug("==============================================OUT OF SCOPE for ==>" + c.GetText())
		log.Debug("l.stack: ", l.stack)
		l.Index.Clear()

		return
	}

}

/*func (s *TAFArgumentListener) ExitMultidim_var_expression(ctx *parser.Multidim_var_expressionContext) {
	name := ctx.GetText()
	fmt.Printf("name: %v\n", name)
}*/

func (l *TAFArgumentListener) ExtractPath(t string, myVar string) (path string) {
	return strings.Replace(t, myVar+".", "", 1)
}

func (l *TAFArgumentListener) ExitIndx_expr(c *parser.Indx_exprContext) {
	//t := c.GetText()
	//fmt.Printf("t: %v\n", t)
}

// ExitVar_expression is called when exiting the var_expression production.
func (l *TAFArgumentListener) ExitVar_expression(c *parser.Var_expressionContext) {
	l.LastExit = "Var_expression"

	if l.OnError {
		return
	}
	//log.Debug("Scope ", l.Scope, "ExitVar_expression: "+c.GetText())
	//currExpr := c.GetText()
	//fmt.Printf("currExpr: %v\n", currExpr)

	if l.VariableContext == nil {
		//no variable context, runtime eror
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: errors.New("the Variable Context is not allocated"), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		l.OnError = true
		return
	}

	if c == nil {
		//runtime error here??
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: errors.New("the Variable Context Expression is not allocated"), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		l.OnError = true
		return
	}
	if c.Indx_expr().Var_path() == nil {
		if c.VARIABLE_NAME() == nil {
			//runtime error here??
			l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: errors.New("the Variable name is empty"), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
			l.OnError = true
			return
		}
		indExprTok := c.Indx_expr().GetText()
		//fmt.Printf("indExprTok: %v\n", indExprTok)
		if indExprTok != "" {
			txt := c.VARIABLE_NAME().GetText() + indExprTok
			evalExpr := l.Index.EvalExpr(txt)
			//fmt.Printf("evalExpr: %v\n", evalExpr)

			l.lastVarExpression = VarExpression{VarName: c.VARIABLE_NAME().GetText(), VarExpression: evalExpr}
			l.detectedVariableExpressions = append(l.detectedVariableExpressions, l.lastVarExpression)
			v, ok := l.VariableContext.GetVariableValue(l.lastVarExpression.VarExpression, l.secretAware)
			if ok != nil {
				if l.onVariableStorageDeclarationMode {
					return
				}
				l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: ok, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
				l.OnError = true
				return
			}
			l.stack.Push(v)
			return
		}
		l.lastVarExpression = VarExpression{VarName: c.VARIABLE_NAME().GetText(), VarExpression: c.VARIABLE_NAME().GetText()}
		l.detectedVariableExpressions = append(l.detectedVariableExpressions, l.lastVarExpression)

		v, ok := l.VariableContext.GetVariableValue(l.lastVarExpression.VarExpression, l.secretAware)
		if ok != nil {
			if l.onVariableStorageDeclarationMode {
				return
			}
			if l.onInterpolationMode {
				//	l.stack.Push(stackvalue.NewStringStackValue("(" + l.lastVarExpression.VarExpression + ")"))
				return
			}
			l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: ok, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
			l.OnError = true
			return
		}
		l.stack.Push(v)
		return
	}
	varName := c.VARIABLE_NAME().GetText()
	if l.Scope > 1 {
		ex := c.GetText()
		//we have a subexpression and its value, we need to push that
		log.Debug("l.stack:", l.stack)
		if ex == varName {
			l.lastVarExpression = VarExpression{VarName: varName, VarExpression: varName}
			l.detectedVariableExpressions = append(l.detectedVariableExpressions, l.lastVarExpression)

			v, ok := l.VariableContext.GetVariableIntValue(varName)
			if ok != nil {
				if l.onVariableStorageDeclarationMode {
					return
				}
				l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: ok, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
				l.OnError = true
				return
			}

			l.Index.Push(ex, int(v))
			l.push(v)
			return
		}

		evalExpr := l.Index.EvalExpr(ex)
		path := l.ExtractPath(evalExpr, varName)
		if ex == varName {
			l.lastVarExpression = VarExpression{VarName: varName, VarExpression: varName}
			l.detectedVariableExpressions = append(l.detectedVariableExpressions, l.lastVarExpression)

			v, ok := l.VariableContext.GetVariableIntValue(varName)
			if ok != nil {
				if l.onVariableStorageDeclarationMode {
					return
				}

				l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: ok, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
				l.OnError = true
				return
			}

			l.Index.Push(ex, int(v))
			l.push(v)
			return
		}
		l.lastVarExpression = VarExpression{VarName: varName, VarExpression: varName + "." + path}
		l.detectedVariableExpressions = append(l.detectedVariableExpressions, l.lastVarExpression)

		v, ok := l.VariableContext.EvaluateJSONVariable(varName, path, l.secretAware)
		if ok != nil {
			if l.onVariableStorageDeclarationMode {
				return
			}

			l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: ok, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
			l.OnError = true
			return
		}

		vIdx, err := l.GetIndex(v)
		if err != nil {
			panic(err)
		}
		l.Index.Push(ex, vIdx)
		l.stack.Push(v)
		return
	}

	expr := l.Index.Eval()

	if expr == varName {
		l.lastVarExpression = VarExpression{VarName: varName, VarExpression: varName}
		l.detectedVariableExpressions = append(l.detectedVariableExpressions, l.lastVarExpression)

		v, ok := l.VariableContext.GetVariableValue(varName, l.secretAware)
		if l.onVariableStorageDeclarationMode {
			return
		}

		if ok != nil {
			l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: ok, Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
			l.OnError = true
			return
		}
		l.stack.Push(v)
		return
	}

	path := l.ExtractPath(expr, varName)
	l.lastVarExpression = VarExpression{VarName: varName, VarExpression: expr}
	l.detectedVariableExpressions = append(l.detectedVariableExpressions, l.lastVarExpression)

	v, ok := l.VariableContext.EvaluateJSONVariable(varName, path, l.secretAware)
	if ok != nil {
		if l.onVariableStorageDeclarationMode {
			return
		}
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("couldnt' obtain path %s from variable %s", path, varName), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		l.OnError = true
		return
	}

	l.stack.Push(v)
}

func (l *TAFArgumentListener) EnterIndexExpression(c *parser.IndexExpressionContext) {
	//log.Debug("Enter Index Expression")
	//l.CurrentPath = c.GetText()
}

func (l *TAFArgumentListener) GetIndex(st stackvalue.StackValue) (index int, err error) {
	if l.OnError {
		return
	}
	indexdV, ok := st.(*stackvalue.DoubleStackValue)
	if !ok {
		indexiV, ok := st.(*stackvalue.IntegerStackValue)
		if !ok {
			return -1, fmt.Errorf("index needs to be a number")
		}
		return indexiV.GetValue(), nil
	}

	return int(indexdV.GetValue()), nil
}

func (l *TAFArgumentListener) ExitIndexExpression(c *parser.IndexExpressionContext) {
	l.LastExit = "Index Expression"

	if l.OnError {
		return
	}
	log.Debug(l.Scope, " Exit Index Expression ")
	p := c.GetText()

	index := l.popStack()
	i, err := l.GetIndex(index)
	if err != nil {
		panic(err)
	}
	l.Index.Push(p, i)

	if l.Scope == 1 {
		//l.Index.Clear()
	}
	return
}

/*func (l *TAFArgumentListener) EnterVarPath(c *parser.VarPathContext) {
}

func (l *TAFArgumentListener) ExitVarPath(c *parser.VarPathContext) {
	log.Debug("ExitVarPath " + c.GetText())
}*/

// ExitParenthesisExpression is called when exiting the parenthesisExpression production.
/*func (l *TAFArgumentListener) EnterParenthesisExpression(c *parser.HandleParenthesisExpressionContext) {
	log.Debug("EnterParenthesisExpression" + c.GetText())
}

func (l *TAFArgumentListener) ExitParenthesisExpression(c *parser.HandleParenthesisExpressionContext) {
	log.Debug("ExitParenthesisExpression" + c.GetText())
}*/

func (l *TAFArgumentListener) ExitHandleJson(ctx *parser.HandleJsonContext) {
	l.LastExit = "JSON"

	//	x := ctx.GetText()
	//	log.Debug("ExitHandleJson" + x)
	/*
	   var json = jsoniter.ConfigCompatibleWithStandardLibrary
	   var payload globals.JSONStruct
	   err := json.Unmarshal([]byte(x), &payload)

	   	if err == nil {
	   		fmt.Printf("payload: %v\n", payload)
	   	} else {

	   		panic("Failed!")
	   	}
	*/
}

func (l *TAFArgumentListener) EnterHandleObjectData(ctx *parser.HandleObjectDataContext) {
	//fmt.Printf("ctx.GetText(): %v\n", ctx.GetText())

}

func (l *TAFArgumentListener) EnterHandleObjectPair(ctx *parser.HandleObjectPairContext) {
	//fmt.Printf("Pair: %v\n", ctx.GetText())
}

func (l *TAFArgumentListener) ExitHandleObjectPair(ctx *parser.HandleObjectPairContext) {
	l.LastExit = "Object Pair"

	current := l.jsonStack[len(l.jsonStack)-1]
	key := ctx.STRING().GetText()
	val := l.popStack()
	tj := stackvalue.GetGenericValue(val)
	//fmt.Printf("tj: %v\n", tj)
	key = strings.TrimLeft(key, "\"")
	key = strings.TrimRight(key, "\"")

	(*current)[key] = tj
}

func (l *TAFArgumentListener) ExitHandleObjectData(ctx *parser.HandleObjectDataContext) {
	l.LastExit = "Object Data"
	//fmt.Printf("ExitHandleObjectData: %v\n", ctx.GetText())
}
func (l *TAFArgumentListener) EnterHandleObject(ctx *parser.HandleObjectContext) {
	st := &globals.JSONStruct{}
	//fmt.Printf("ctx.GetText(): %v\n", ctx.GetText())
	l.jsonStack = append(l.jsonStack, st)
}

func (l *TAFArgumentListener) ExitHandleObject(ctx *parser.HandleObjectContext) {
	l.LastExit = "Object"
	if l.OnError {
		return
	}

	last := l.jsonStack[len(l.jsonStack)-1]
	l.jsonStack = l.jsonStack[:len(l.jsonStack)-1]
	if len(l.jsonStack) == 0 {
		l.stack.Push(stackvalue.NewJSONStackValue(*last))
	} else {
		l.stack.Push(stackvalue.NewJSONStackValue(*last))
	}
}

func (l *TAFArgumentListener) EnterHandleArray(ctx *parser.HandleArrayContext) {
	l.arrayScope++
	l.stack.CreateStack()
}

func (l *TAFArgumentListener) ExitHandleArray(ctx *parser.HandleArrayContext) {
	l.LastExit = "Array"
	if l.OnError {
		return
	}
	st := l.stack.ExitStack()

	l.arrayScope--

	x := ctx.GetText()
	//log.Debug("ExitHandleArrayt" + x)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var payload globals.JSONArrayGen

	if st != nil {
		l.stack.Push(stackvalue.NewJSONArrayStackValue(*st))
		return
	}
	err := json.Unmarshal([]byte(x), &payload)
	if err == nil {
		l.stack.Push(stackvalue.NewJSONArrayStackValue(payload))
	} else {
		l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf(" couldn't import JSON Array"), Type: tafargumentlistenererrortypes.JSON_ARRAY_IMPORT_ERROR})
		l.OnError = true
		return
	}
}
