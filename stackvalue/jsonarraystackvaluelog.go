package stackvalue

import (
	"fmt"

	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

func (me JSONArrayStackValue) LesserThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" lesser than operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) LesserThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" lesser than equal operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Equal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case stackvalue.NULL:
		return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_ARRAY:
		return NewBoolStackValue(me.Equals(s)), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation == between a JSON Array and a %s", s.GetType().String())
}

func (me JSONArrayStackValue) Unequal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case stackvalue.NULL:
		return NewBoolStackValue(me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		return NewBoolStackValue(!me.Equals(s)), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation != between a JSON Array and a %s", s.GetType().String())
}

func (me JSONArrayStackValue) GreaterThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" greater than operation does not apply to JSON Array")
}
func (me JSONArrayStackValue) GreaterThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" greater than equal operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) And(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case stackvalue.NULL:
		return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	default:
		l := me.IsTruthy()
		r := s.IsTruthy()
		return NewBoolStackValue(l && r), tafargumenetlistenererrortypes.NONE, nil
	}
}

func (me JSONArrayStackValue) Or(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if me.IsTruthy() {
		return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
	}
	return NewBoolStackValue(s.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}

func (me JSONArrayStackValue) Not() (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}

func (me JSONArrayStackValue) Add(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" add operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Sub(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" sub operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Mul(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" mul operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Div(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" div operation does not apply to JSON Array")
}

func (me JSONArrayStackValue) Mod(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumentlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf(" mod operation does not apply to JSON Array")
}
