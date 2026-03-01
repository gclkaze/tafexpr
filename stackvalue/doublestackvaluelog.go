package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

func (b DoubleStackValue) LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewBoolStackValue(doubleValue < 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.LesserThan(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(doubleValue < float64(intV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(doubleValue < doubleV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(int(doubleValue) < rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(doubleValue < 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewBoolStackValue(doubleValue <= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.LesserThanEqual(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(int(doubleValue) <= intV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(doubleValue <= (doubleV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(int(doubleValue) <= rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(doubleValue <= 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than equal operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewBoolStackValue(doubleValue == 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Equal(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(float64(intV.GetValue()) == doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue((doubleV.GetValue()) == doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(int(doubleValue) == rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(doubleValue == 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible equality operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewBoolStackValue(doubleValue != 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Unequal(inner)
		}
	case STRING:
		return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf("incompatible operation != between a double and a string")
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(float64(intV.GetValue()) != doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue((doubleV.GetValue()) != doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(int(doubleValue) != rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(doubleValue != 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible unequality operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewBoolStackValue(doubleValue > 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.GreaterThan(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(doubleValue > float64(intV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(doubleValue > doubleV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(int(doubleValue) > rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(doubleValue > 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewBoolStackValue(doubleValue >= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.GreaterThanEqual(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewBoolStackValue(doubleValue >= float64(intV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewBoolStackValue(doubleValue >= doubleV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}
		return NewBoolStackValue(int(doubleValue) >= rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(doubleValue >= 0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than equal operation between a double and a %s", s.GetType().String())
}
func (b DoubleStackValue) And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
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
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible add operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
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
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible or operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if b.value <= 0 {
		return NewBoolStackValue(true), tafargumenetlistenererrortypes.NONE, nil
	}
	return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
}
