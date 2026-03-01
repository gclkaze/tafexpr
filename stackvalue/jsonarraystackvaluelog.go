package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

func (me JSONArrayStackValue) LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" lesser than operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" lesser than equal operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case NULL:
		return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_ARRAY:
		return NewBoolStackValue(me.Equals(s)), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation == between a JSON Array and a %s", s.GetType().String())
}

func (me JSONArrayStackValue) Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case NULL:
		return NewBoolStackValue(me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		return NewBoolStackValue(!me.Equals(s)), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation != between a JSON Array and a %s", s.GetType().String())
}

func (me JSONArrayStackValue) GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" greater than operation does not apply to JSON Array")
}
func (me JSONArrayStackValue) GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" greater than equal operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case NULL:
		return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	default:
		l := me.IsTruthy()
		r := s.IsTruthy()
		return NewBoolStackValue(l && r), tafargumenetlistenererrortypes.NONE, nil
	}
}

func (me JSONArrayStackValue) Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if me.IsTruthy() {
		return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
	}
	return NewBoolStackValue(s.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}

func (me JSONArrayStackValue) Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}

func (me JSONArrayStackValue) Add(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" add operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Sub(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" sub operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Mul(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" mul operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Div(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" div operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Mod(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" mod operation does not apply to JSON Array")
}
