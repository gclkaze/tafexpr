package tafexpr

import (
	"fmt"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
	"github.com/gclkaze/tafexpr/methods"
	"github.com/gclkaze/tafexpr/parser"
	mine "github.com/gclkaze/tafexpr/stackvalue"
)

// ExitHandleMethodCall evaluates  receiver.name(args...)  through the method registry.
// It replaces the 19 per-method handlers: the grammar no longer knows any method name.
func (l *TAFArgumentListener) ExitHandleMethodCall(c *parser.HandleMethodCallContext) {
	l.LastExit = "MethodCall"
	if l.OnError {
		return
	}
	name := c.PROP().GetText()
	args := l.popArgs(len(c.AllExpression()) - 1) // the first expression is the receiver
	recv := normalize(l.popStack())
	if recv == nil || hasNil(args) {
		l.failCall(c.GetText(), fmt.Errorf("cannot call %s: a value is missing", name))
		return
	}
	res, err := methods.Invoke(recv, name, args)
	l.finishCall(c.GetText(), res, err)
}

// ExitHandleFunctionCall evaluates  name(args...)  (no receiver), e.g. randomDoubleInRange(1, 5).
func (l *TAFArgumentListener) ExitHandleFunctionCall(c *parser.HandleFunctionCallContext) {
	l.LastExit = "FunctionCall"
	if l.OnError {
		return
	}
	name := c.PROP().GetText()
	args := l.popArgs(len(c.AllExpression())) // no receiver: every expression is an argument
	if hasNil(args) {
		l.failCall(c.GetText(), fmt.Errorf("cannot call %s: a value is missing", name))
		return
	}
	res, err := methods.InvokeFunc(name, args)
	l.finishCall(c.GetText(), res, err)
}

// popArgs pops n arguments (they were pushed left to right) and returns them in source order.
func (l *TAFArgumentListener) popArgs(n int) []stackvalue.StackValue {
	if n < 0 {
		n = 0
	}
	args := make([]stackvalue.StackValue, n)
	for i := n - 1; i >= 0; i-- {
		args[i] = normalize(l.popStack())
	}
	return args
}

func hasNil(vs []stackvalue.StackValue) bool {
	for _, v := range vs {
		if v == nil {
			return true
		}
	}
	return false
}

func (l *TAFArgumentListener) failCall(text string, err error) {
	l.OnError = true
	l.ErrorMsgs = append(l.ErrorMsgs, TAFParserArgumentError{
		Msg:  fmt.Errorf("%w: %s", err, text),
		Type: tafargumentlistenererrortypes.RUNTIME_ERROR,
	})
}

func (l *TAFArgumentListener) finishCall(text string, res stackvalue.StackValue, err error) {
	if err != nil {
		l.failCall(text, err)
		return
	}
	if res == nil {
		l.failCall(text, fmt.Errorf("the call returned no value"))
		return
	}
	l.pushValue(res)
}

// normalize unwraps a JSONStackValue that merely wraps a scalar or a list. A JSONStackValue always
// reports JSON_OBJECT as its type, so without this a string read from a JSON path would dispatch
// as an object and $row.name.trim() would not find trim. A nil inner value becomes Null.
func normalize(v stackvalue.StackValue) stackvalue.StackValue {
	var j *mine.JSONStackValue
	switch t := v.(type) {
	case *mine.JSONStackValue:
		j = t
	case mine.JSONStackValue:
		j = &t
	default:
		return v
	}
	if j.GetValue() == nil {
		return mine.NewNullStackValue()
	}
	if inner, err := j.GetInnerValue(); err == nil {
		return inner
	}
	return v
}
