package stackvalue

import (
	"fmt"
	"strings"

	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

func (c NullStackValue) LesserThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) < 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 < other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.DOUBLE:
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
func (c NullStackValue) LesserThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) <= 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 <= other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.DOUBLE:
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

func (c NullStackValue) Equal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) == 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(other.GetValue() == 0), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(other.GetValue() == 0), tafargumenetlistenererrortypes.NONE, nil

		}

	case stackvalue.BOOL:
		{
			other, ok := s.(*BoolStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
			}
			return NewBoolStackValue(!other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(!arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(!oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil

	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than operation does not apply to NULL and %s", t.String())
}

func (c NullStackValue) Unequal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) != 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(other.GetValue() != 0), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.DOUBLE:
		{
			other, ok := s.(*DoubleStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
			}
			return NewBoolStackValue(other.GetValue() != 0), tafargumenetlistenererrortypes.NONE, nil

		}

	case stackvalue.BOOL:
		{
			other, ok := s.(*BoolStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
			}
			return NewBoolStackValue(other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" greater than operation does not apply to NULL and %s", t.String())
}

func (c NullStackValue) GreaterThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) > 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 > other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.DOUBLE:
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

func (c NullStackValue) GreaterThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {

	t := s.GetType()
	switch t {
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare("", otherS.value) >= 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.INTEGER:
		{
			other, ok := s.(*IntegerStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to int")
			}
			return NewBoolStackValue(0 >= other.GetValue()), tafargumenetlistenererrortypes.NONE, nil

		}
	case stackvalue.DOUBLE:
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

func (b NullStackValue) And(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	isTruthy := false
	switch s.GetType() {
	case stackvalue.STRING:
		strValue, ok := s.(*StringStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
		}
		return NewBoolStackValue(isTruthy && strValue.GetValue() != ""), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(isTruthy && intV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(isTruthy && doubleV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}

		return NewBoolStackValue(isTruthy && boolV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}
		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(false && !arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(false && !oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" and operation does not apply to NULL and %s", s.GetType().String())
}

func (b NullStackValue) Or(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	switch s.GetType() {
	case stackvalue.STRING:
		strValue, ok := s.(*StringStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
		}
		return NewBoolStackValue(strValue.GetValue() != ""), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(doubleV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}

		return NewBoolStackValue(boolV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
		}
		return NewBoolStackValue(oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf("or operation does not apply to NULL and %s", s.GetType().String())
}

func (b NullStackValue) Not() (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
}
