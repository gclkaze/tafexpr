package stackvalue

import (
	"fmt"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
)

func (b BoolStackValue) LesserThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	//true < X
	//false < x
	boolValue := b.value
	switch s.GetType() {
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
				intBoolValue := 0
				if boolValue {
					intBoolValue = 1
				}
				return NewBoolStackValue(intBoolValue < 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.LesserThan(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		intValue := intV.GetValue()
		intBoolValue := 0
		if boolValue {
			intBoolValue = 1
		}
		return NewBoolStackValue(intBoolValue < intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		doubleValue := doubleV.GetValue()
		doubleBoolValue := 0
		if boolValue {
			doubleBoolValue = 1
		}

		return NewBoolStackValue(doubleBoolValue < int(doubleValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}

		lhsValue := 0
		if boolValue {
			lhsValue = 1
		}
		return NewBoolStackValue(lhsValue < rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than operation between a boolean and a %s", s.GetType().String())
}
func (b BoolStackValue) LesserThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	boolValue := b.value
	switch s.GetType() {
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
				intBoolValue := 0
				if boolValue {
					intBoolValue = 1
				}
				return NewBoolStackValue(intBoolValue <= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.LesserThanEqual(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		intValue := intV.GetValue()
		intBoolValue := 0
		if boolValue {
			intBoolValue = 1
		}
		return NewBoolStackValue(intBoolValue <= intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		doubleValue := doubleV.GetValue()
		doubleBoolValue := 0
		if boolValue {
			doubleBoolValue = 1
		}

		return NewBoolStackValue(doubleBoolValue <= int(doubleValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}

		lhsValue := 0
		if boolValue {
			lhsValue = 1
		}
		return NewBoolStackValue(lhsValue <= rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible lesser than equal operation between a boolean and a %s", s.GetType().String())
}
func (b BoolStackValue) Equal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	boolValue := b.value
	switch s.GetType() {
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
				intBoolValue := 0
				if boolValue {
					intBoolValue = 1
				}
				return NewBoolStackValue(intBoolValue == 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Equal(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		intValue := intV.GetValue()
		/*		intBoolValue := 0
				if boolValue {
					intBoolValue = 1
				}

		*/
		res := false
		if boolValue {
			if intValue <= 0 {
				res = false
			} else {
				res = true
			}
		} else {
			if intValue <= 0 {
				res = true
			} else {
				res = false
			}
		}

		return NewBoolStackValue(res), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		doubleValue := doubleV.GetValue()
		res := false
		if boolValue {
			if doubleValue <= 0 {
				res = false
			} else {
				res = true
			}
		} else {
			if doubleValue <= 0 {
				res = true
			} else {
				res = false
			}
		}

		return NewBoolStackValue(res), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		return NewBoolStackValue(boolV.GetValue() == boolValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(!boolValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible equality operation between a boolean and a %s", s.GetType().String())
}
func (b BoolStackValue) Unequal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	boolValue := b.value
	switch s.GetType() {
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
				intBoolValue := 0
				if boolValue {
					intBoolValue = 1
				}
				return NewBoolStackValue(intBoolValue != 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Unequal(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		intValue := intV.GetValue()
		res := false
		if boolValue {
			if intValue <= 0 {
				res = false
			} else {
				res = true
			}
		} else {
			if intValue <= 0 {
				res = true
			} else {
				res = false
			}
		}

		return NewBoolStackValue(!res), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		doubleValue := doubleV.GetValue()
		res := false
		if boolValue {
			if doubleValue <= 0 {
				res = false
			} else {
				res = true
			}
		} else {
			if doubleValue <= 0 {
				res = true
			} else {
				res = false
			}
		}

		return NewBoolStackValue(!res), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		return NewBoolStackValue(boolV.GetValue() != boolValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(boolValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible unequality operation between a boolean and a %s", s.GetType().String())
}
func (b BoolStackValue) GreaterThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	boolValue := b.value
	switch s.GetType() {
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
				intBoolValue := 0
				if boolValue {
					intBoolValue = 1
				}
				return NewBoolStackValue(intBoolValue > 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.GreaterThan(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		intValue := intV.GetValue()
		intBoolValue := 0
		if boolValue {
			intBoolValue = 1
		}
		return NewBoolStackValue(intBoolValue > intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		doubleValue := doubleV.GetValue()
		doubleBoolValue := 0
		if boolValue {
			doubleBoolValue = 1
		}

		return NewBoolStackValue(doubleBoolValue > int(doubleValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}

		lhsValue := 0
		if boolValue {
			lhsValue = 1
		}
		return NewBoolStackValue(lhsValue > rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than operation between a boolean and a %s", s.GetType().String())
}
func (b BoolStackValue) GreaterThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	boolValue := b.value
	switch s.GetType() {
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
				intBoolValue := 0
				if boolValue {
					intBoolValue = 1
				}
				return NewBoolStackValue(intBoolValue >= 0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.GreaterThanEqual(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		intValue := intV.GetValue()
		intBoolValue := 0
		if boolValue {
			intBoolValue = 1
		}
		return NewBoolStackValue(intBoolValue >= intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		doubleValue := doubleV.GetValue()
		doubleBoolValue := 0
		if boolValue {
			doubleBoolValue = 1
		}

		return NewBoolStackValue(doubleBoolValue >= int(doubleValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		rhsValue := 0
		if boolV.GetValue() {
			rhsValue = 1
		}

		lhsValue := 0
		if boolValue {
			lhsValue = 1
		}
		return NewBoolStackValue(lhsValue >= rhsValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible greater than equal operation between a boolean and a %s", s.GetType().String())
}

func (b BoolStackValue) And(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	if !b.value {
		return NewBoolStackValue(false), tafargumenetlistenererrortypes.NONE, nil
	}
	isTruthy := true
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
		return NewBoolStackValue(isTruthy && arrV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.JSON_OBJECT:
		oV, ok := s.(*JSONStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to json")
		}
		return NewBoolStackValue(isTruthy && oV.IsTruthy()), tafargumenetlistenererrortypes.NONE, nil

	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible and operation between a boolean and a %s", s.GetType().String())
}

func (b BoolStackValue) Or(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	isTruthy := b.value
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
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible or operation between a boolean and a %s", s.GetType().String())
}

func (b BoolStackValue) Not() (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return NewBoolStackValue(!b.value), tafargumenetlistenererrortypes.NONE, nil
}
