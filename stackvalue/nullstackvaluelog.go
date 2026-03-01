package stackvalue

import (
	"fmt"
	"strings"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"

	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

func (c NullStackValue) LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) < 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 < other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(0 < other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" lesser than operation does not apply to NULL and %s", t.String())
}
func (c NullStackValue) LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) <= 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 <= other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(0 <= other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" lesser than equals operation does not apply to NULL and %s", t.String())
}

func (c NullStackValue) Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) == 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(other.GetValue() == 0), tafargumenetlistenererrortypes.NONE, nil

		}
	case DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(other.GetValue() == 0), tafargumenetlistenererrortypes.NONE, nil

		}

	case BOOL:
		{
			other, ok := s.(*BoolStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
			}
			return NewBoolStackValue(!other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(!arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(!oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil

	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than operation does not apply to NULL and %s", t.String())
}

func (c NullStackValue) Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) != 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
		}
	case INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(other.GetValue() != 0), tafargumenetlistenererrortypes.NONE, nil

		}
	case DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(other.GetValue() != 0), tafargumenetlistenererrortypes.NONE, nil

		}

	case BOOL:
		{
			other, ok := s.(*BoolStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
			}
			return NewBoolStackValue(other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than operation does not apply to NULL and %s", t.String())
}

func (c NullStackValue) GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) > 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 > other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(0 > other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than operation does not apply to NULL and %s", t.String())
}

func (c NullStackValue) GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) >= 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 >= other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(0 >= other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than equals operation does not apply to NULL and %s", t.String())
}

func (b NullStackValue) And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	isTruthy := false
	switch s.GetType() {
	case STRING:
		strValue, ok := s.(*StringStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
		}
		return NewBoolStackValue(isTruthy && strValue.GetValue() != ""), tafargumenetlistenererrortypes.NONE, nil
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(isTruthy && intV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(isTruthy && doubleV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}

		return NewBoolStackValue(isTruthy && boolV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}
		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	case JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(false && !arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(false && !oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" and operation does not apply to NULL and %s", s.GetType().String())
}

func (b NullStackValue) Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case STRING:
		strValue, ok := s.(*StringStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
		}
		return NewBoolStackValue(strValue.GetValue() != ""), tafargumenetlistenererrortypes.NONE, nil
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(doubleV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}

		return NewBoolStackValue(boolV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	case JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf("or operation does not apply to NULL and %s", s.GetType().String())
}

func (b NullStackValue) Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
}
