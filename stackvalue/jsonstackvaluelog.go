package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

func (me JSONStackValue) LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	/*	i, err := me.GetInnerValue()
		if err != nil {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
		}
		if s.GetType() == JSON_OBJECT {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation < between two JSON Objects")
		}
		if i == nil {
			v := NewIntegerStackValue(0)
			return v.LesserThan(s)
		}
		return i.LesserThan(s)*/
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation < between a JSON Object and a %s", s.GetType().String())
}

func (me JSONStackValue) LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	/*	i, err := me.GetInnerValue()
		if err != nil {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
		}
		if s.GetType() == JSON_OBJECT {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation <= between two JSON Objects")
		}

		if i == nil {
			v := NewIntegerStackValue(0)
			return v.LesserThanEqual(s)
		}
		return i.LesserThanEqual(s)*/
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation <= between a JSON Object and a %s", s.GetType().String())
}

func (me JSONStackValue) Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case NULL:
		return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		return NewBoolStackValue(me.Equals(s)), tafargumenetlistenererrortypes.NONE, nil
		/*	case INTEGER:
			return NewBoolStackValue((me.Equals(s))), tafargumenetlistenererrortypes.NONE, nil*/
	}
	i, err := me.GetInnerValue()
	if err != nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
	}
	if i == nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation == between a JSON Object and a %s", s.GetType().String())
	}
	if i.GetType() != JSON_OBJECT {
		return i.Equal(s)
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation == between a JSON Object and a %s", s.GetType().String()) //i.Equal(s)
}

func (me JSONStackValue) Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case NULL:
		return NewBoolStackValue(me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		return NewBoolStackValue(!me.Equals(s)), tafargumenetlistenererrortypes.NONE, nil
	}
	i, err := me.GetInnerValue()
	if err != nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
	}
	if i == nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation != between a JSON Object and a %s", s.GetType().String())
	}
	if i.GetType() != JSON_OBJECT {
		return i.Equal(s)
	}

	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation != between a JSON Object and a %s", s.GetType().String()) //i.Unequal(s)
}

func (me JSONStackValue) GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	/*	i, err := me.GetInnerValue()
		if err != nil {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
		}
		if s.GetType() == JSON_OBJECT {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation > between two JSON Objects")
		}
		if i == nil {
			v := NewIntegerStackValue(0)
			return v.GreaterThan(s)
		}
		return i.GreaterThan(s)*/
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation > between a JSON Object and a %s", s.GetType().String())

}

func (me JSONStackValue) GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	/*	i, err := me.GetInnerValue()
		if err != nil {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
		}
		if s.GetType() == JSON_OBJECT {
			return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation >= between two JSON Objects")
		}

		if i == nil {
			v := NewIntegerStackValue(0)
			return v.GreaterThanEqual(s)
		}
		return i.GreaterThanEqual(s)*/
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible operation >= between a JSON Object and a %s", s.GetType().String())

}

func (me JSONStackValue) And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case NULL:
		return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	default:
		l := me.IsTruthy()
		r := s.IsTruthy()
		return NewBoolStackValue(l && r), tafargumenetlistenererrortypes.NONE, nil
	}
}

func (me JSONStackValue) Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if me.IsTruthy() {
		return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
	}
	return NewBoolStackValue(s.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
}

func (me JSONStackValue) Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(!me.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil

}

func (me JSONStackValue) Add(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	i, err := me.GetInnerValue()
	if err != nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
	}
	if i == nil {
		v := NewIntegerStackValue(0)
		return v.Add(s)
	}

	return i.Add(s)
}

func (me JSONStackValue) Sub(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	i, err := me.GetInnerValue()
	if err != nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
	}
	if i == nil {
		v := NewIntegerStackValue(0)
		return v.Sub(s)
	}

	return i.Sub(s)
}

func (me JSONStackValue) Mul(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	i, err := me.GetInnerValue()
	if err != nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
	}
	if i == nil {
		return NewIntegerStackValue(0), tafargumenetlistenererrortypes.NONE, nil
	}

	return i.Mul(s)
}

func (me JSONStackValue) Div(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	i, err := me.GetInnerValue()
	if err != nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
	}
	if i == nil {
		return NewIntegerStackValue(0).Div(s)
	}

	return i.Div(s)
}

func (me JSONStackValue) Mod(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	i, err := me.GetInnerValue()
	if err != nil {
		return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
	}
	if i == nil {
		return NewIntegerStackValue(0).Mod(s)
	}

	return i.Mod(s)
}
