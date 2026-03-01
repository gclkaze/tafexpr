package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

func (b IntegerStackValue) LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := b.value
	switch s.GetType() {
	case JSON_OBJECT:
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
				return NewBoolStackValue(intValue < 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.LesserThan(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intValue < intV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(intValue < int(doubleV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(intValue < rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(intValue < 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := b.value
	switch s.GetType() {
	case JSON_OBJECT:
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
				return NewBoolStackValue(intValue <= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.LesserThanEqual(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intValue <= intV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(intValue <= int(doubleV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(intValue <= rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(intValue <= 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than equal operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := b.value
	switch s.GetType() {
	case JSON_OBJECT:
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
				return NewBoolStackValue(intValue == 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Equal(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intV.GetValue() == intValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(int(doubleV.GetValue()) == intValue), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(intValue == rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(intValue == 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible equality operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := b.value
	switch s.GetType() {
	case JSON_OBJECT:
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
				return NewBoolStackValue(intValue != 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Unequal(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intV.GetValue() != intValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(int(doubleV.GetValue()) != intValue), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(intValue != rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(intValue != 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible unequality operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := b.value
	switch s.GetType() {
	case JSON_OBJECT:
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
				return NewBoolStackValue(intValue > 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.GreaterThan(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intValue > intV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(intValue > int(doubleV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(intValue > rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(intValue > 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := b.value
	switch s.GetType() {
	case JSON_OBJECT:
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
				return NewBoolStackValue(intValue >= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.GreaterThanEqual(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(intValue >= intV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(intValue >= int(doubleV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(intValue >= rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(intValue >= 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than equal operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if b.value <= 0 {
		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}
	isTruthy := true
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
		return NewBoolStackValue(isTruthy && arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json")
		}
		return NewBoolStackValue(isTruthy && oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible and operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := b.value
	isTruthy := false
	if intValue > 0 {
		isTruthy = true
	}
	switch s.GetType() {
	case STRING:
		strValue, ok := s.(*StringStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to string")
		}
		return NewBoolStackValue(isTruthy || strValue.GetValue() != ""), tafargumenetlistenererrortypes.NONE, nil
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(isTruthy || intV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(isTruthy || doubleV.GetValue() > 0), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}

		return NewBoolStackValue(isTruthy || boolV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(intValue >= 0), tafargumenetlistenererrortypes.NONE, nil
	case JSON_ARRAY:
		arrV, ok := s.(*JSONArrayStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json array")
		}
		return NewBoolStackValue(isTruthy || arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json")
		}
		return NewBoolStackValue(isTruthy || oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible or operation between an integer and a %s", s.GetType().String())
}

func (b IntegerStackValue) Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if b.value <= 0 {
		return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
	}
	return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
}
