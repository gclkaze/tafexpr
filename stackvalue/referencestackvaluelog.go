package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

func (c ReferenceStackValue) LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" lesser than operation does not apply to REFERENCE and %s", t.String())
}

func (c ReferenceStackValue) LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" lesser than equal operation does not apply to REFERENCE and %s", t.String())
}

func (c ReferenceStackValue) Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	ct := c.GetType()
	if t != ct {
		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}

	//both referemces
	me, ok := s.(*ReferenceStackValue)
	if !ok {
		return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert other stack value to reference object")
	}
	other, ok := s.(*ReferenceStackValue)
	if !ok {
		return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert current stack value to reference object")
	}

	return NewBoolStackValue(me.value == other.value), tafargumenetlistenererrortypes.NONE, nil

}
func (c ReferenceStackValue) Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	b, t, e := c.Equal(s)
	if b == nil {
		return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, e
	}
	bb, _ := b.(*BoolStackValue)
	return NewBoolStackValue(bb.GetValue()), t, nil
}

func (c ReferenceStackValue) GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than operation does not apply to REFERENCE and %s", t.String())
}

func (c ReferenceStackValue) GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than equal operation does not apply to REFERENCE and %s", t.String())
}

func (c ReferenceStackValue) And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(c.IsTruthy() && s.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}
func (c ReferenceStackValue) Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(c.IsTruthy() || s.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}
func (c ReferenceStackValue) Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue((!c.IsTruthy())), tafargumenetlistenererrortypes.NONE, nil
}
