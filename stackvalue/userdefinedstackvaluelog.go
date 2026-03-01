package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

func (c UserDefinedStackValue) LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" lesser than operation does not apply to UserDefined and %s", t.String())
}

func (c UserDefinedStackValue) LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" lesser than equal operation does not apply to UserDefined and %s", t.String())
}

func (c UserDefinedStackValue) Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	ct := c.GetType()
	if t != ct {
		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}

	//both referemces
	me, ok := s.(*UserDefinedStackValue)
	if !ok {
		return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert other stack value to UserDefined object")
	}
	other, ok := s.(*UserDefinedStackValue)
	if !ok {
		return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert current stack value to UserDefined object")
	}

	return NewBoolStackValue(me.ptr == other.ptr), tafargumenetlistenererrortypes.NONE, nil

}
func (c UserDefinedStackValue) Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	b, t, e := c.Equal(s)
	if b == nil {
		return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, e
	}
	bb, _ := b.(*BoolStackValue)
	return NewBoolStackValue(bb.GetValue()), t, nil
}

func (c UserDefinedStackValue) GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than operation does not apply to UserDefined and %s", t.String())
}

func (c UserDefinedStackValue) GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than equal operation does not apply to UserDefined and %s", t.String())
}

func (c UserDefinedStackValue) And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(c.IsTruthy() && s.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}
func (c UserDefinedStackValue) Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(c.IsTruthy() || s.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}
func (c UserDefinedStackValue) Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue((!c.IsTruthy())), tafargumenetlistenererrortypes.NONE, nil
}
