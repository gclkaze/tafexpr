package stackvalue

import (
	"fmt"
	"strings"

	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

func (st StringStackValue) LesserThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case stackvalue.JSON_OBJECT:
		{
			v, ok := s.(*JSONStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
			}
			inner, err := v.GetInnerValue()
			if err != nil {
				return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
			}
			if inner == nil {
				return NewBoolStackValue(strings.Compare(st.value, "") < 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return st.LesserThan(inner)
		}
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare(st.value, otherS.value) < 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(strings.Compare(st.value, "") < 0), tafargumenetlistenererrortypes.NONE, nil

		}

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than operation between a string and a %s", s.GetType().String())
}
func (st StringStackValue) LesserThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case stackvalue.JSON_OBJECT:
		{
			v, ok := s.(*JSONStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
			}
			inner, err := v.GetInnerValue()
			if err != nil {
				return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
			}
			if inner == nil {
				return NewBoolStackValue(strings.Compare(st.value, "") <= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return st.LesserThanEqual(inner)
		}

	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare(st.value, otherS.value) <= 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(strings.Compare(st.value, "") <= 0), tafargumenetlistenererrortypes.NONE, nil

		}

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than equal operation between a string and a %s", s.GetType().String())
}
func (st StringStackValue) Equal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case stackvalue.JSON_OBJECT:
		{
			v, ok := s.(*JSONStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
			}
			inner, err := v.GetInnerValue()
			if err != nil {
				return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
			}
			if inner == nil {
				return NewBoolStackValue(strings.Compare(st.value, "") == 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return st.Equal(inner)
		}
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare(st.value, otherS.value) == 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(strings.Compare(st.value, "") == 0), tafargumenetlistenererrortypes.NONE, nil

		}
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible equality operation between a string and a %s", s.GetType().String())

}
func (st StringStackValue) Unequal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case stackvalue.JSON_OBJECT:
		{
			v, ok := s.(*JSONStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
			}
			inner, err := v.GetInnerValue()
			if err != nil {
				return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
			}
			if inner == nil {
				return NewBoolStackValue(strings.Compare(st.value, "") != 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return st.Unequal(inner)
		}

	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare(st.value, otherS.value) != 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(strings.Compare(st.value, "") != 0), tafargumenetlistenererrortypes.NONE, nil

		}

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible unequality operation between a string and a %s", s.GetType().String())

}
func (st StringStackValue) GreaterThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case stackvalue.JSON_OBJECT:
		{
			v, ok := s.(*JSONStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
			}
			inner, err := v.GetInnerValue()
			if err != nil {
				return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
			}
			if inner == nil {
				return NewBoolStackValue(strings.Compare(st.value, "") > 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return st.GreaterThan(inner)
		}
	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare(st.value, otherS.value) > 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(strings.Compare(st.value, "") > 0), tafargumenetlistenererrortypes.NONE, nil

		}

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than operation between a string and a %s", s.GetType().String())

}
func (st StringStackValue) GreaterThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	switch t {
	case stackvalue.JSON_OBJECT:
		{
			v, ok := s.(*JSONStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json object")
			}
			inner, err := v.GetInnerValue()
			if err != nil {
				return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, err
			}
			if inner == nil {
				return NewBoolStackValue(strings.Compare(st.value, "") >= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return st.GreaterThanEqual(inner)
		}

	case stackvalue.STRING:
		{
			otherS, ok := s.(*StringStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
			}
			return NewBoolStackValue(strings.Compare(st.value, otherS.value) >= 0), tafargumenetlistenererrortypes.NONE, nil
		}
	case stackvalue.NULL:
		{
			_, ok := s.(*NullStackValue)
			if !ok {
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
			}
			return NewBoolStackValue(strings.Compare(st.value, "") >= 0), tafargumenetlistenererrortypes.NONE, nil

		}

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than equal operation between a string and a %s", s.GetType().String())

}
func (b StringStackValue) And(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if b.value == "" {
		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}
	strValue := b.value
	isTruthy := false
	if strValue != "" {
		isTruthy = true
	}
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

		return NewBoolStackValue(!isTruthy), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(isTruthy && arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json")
		}
		return NewBoolStackValue(isTruthy && oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible and operation between a string and a %s", s.GetType().String())

}

func (b StringStackValue) Or(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	strValue := b.value
	isTruthy := false
	if strValue != "" {
		isTruthy = true
	}
	switch s.GetType() {
	case stackvalue.STRING:
		strValue, ok := s.(*StringStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
		}
		return NewBoolStackValue(isTruthy || strValue.GetValue() != ""), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(isTruthy || intV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(isTruthy || doubleV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		return NewBoolStackValue(isTruthy || boolV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}
		return NewBoolStackValue(isTruthy), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(isTruthy || arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json")
		}
		return NewBoolStackValue(isTruthy || oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible or operation between a string and a %s", s.GetType().String())
}

func (b StringStackValue) Not() (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(b.value == ""), tafargumenetlistenererrortypes.NONE, nil
}
