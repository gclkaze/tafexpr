package stackvalue

import (
	"fmt"
	"strconv"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	tafargumenetlistenererrortypes "github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	log "github.com/sirupsen/logrus"
)

type BoolStackValue struct {
	value bool
}

func NewBoolStackValue(i bool) *BoolStackValue {
	return &BoolStackValue{value: i}
}

func (s BoolStackValue) ToInteger() (result int64, err error) {
	if s.value {
		return 1, nil
	}
	return 0, nil
}

func (s BoolStackValue) ToUintPtr() uintptr {
	ptr := uintptr(0)
	if s.value {
		ptr = 1
	}
	return ptr
}

func (s BoolStackValue) ToDouble() (result float64, err error) {
	if s.value {
		return float64(1), nil
	}
	return 0, nil
}

func (s BoolStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s BoolStackValue) GetType() stackvalue.StackValueType {
	return stackvalue.BOOL
}

func (s BoolStackValue) GetValue() bool {
	return s.value
}

func (s BoolStackValue) IsTruthy() bool {
	return s.value
}

func (s BoolStackValue) Length() (result int, err error) {
	return -1, fmt.Errorf("Length does not apply to bool value")
}

func (s BoolStackValue) ToString() string {
	return strconv.FormatBool(s.value)
}

func (s BoolStackValue) IsScalar() bool {
	return true
}

func (s BoolStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}

func (s BoolStackValue) Equals(other stackvalue.StackValue) bool {
	if other.GetType() != stackvalue.BOOL {
		if other.GetType() == stackvalue.JSON_OBJECT {
			b, e, x := other.Equal(&s)
			if e == tafargumentlistenererrortypes.NONE && x == nil {
				bb := b.(*BoolStackValue)
				return bb.GetValue()
			} else {
				log.Println(x.Error())
				return false
			}
		}
		return false
	}
	dV, ok := other.(*BoolStackValue)
	if !ok {
		return false
	}
	return s.value == dV.value
}

func (s BoolStackValue) Copy() stackvalue.StackValue {
	return NewBoolStackValue(s.value)
}

func (b BoolStackValue) Add(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := 0
	if b.value {
		intValue = 1
	}
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
				return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Add(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewIntegerStackValue(intV.GetValue() + intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewIntegerStackValue(int(doubleV.GetValue()) + intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if boolV.GetValue() {
			intValue++
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible add operation between a boolean and a %s", s.GetType().String())
}
func (b BoolStackValue) Sub(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := 0
	if b.value {
		intValue = 1
	}
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
				return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Sub(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewIntegerStackValue(-intV.GetValue() + intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewIntegerStackValue(int(-doubleV.GetValue()) + intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if boolV.GetValue() {
			intValue--
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible sub operation between a boolean and a %s", s.GetType().String())

}
func (b BoolStackValue) Mul(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := 0
	if b.value {
		intValue = 1
	}
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
				return NewIntegerStackValue(0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Mul(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewIntegerStackValue(intV.GetValue() * intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewIntegerStackValue(int(doubleV.GetValue()) * intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			intValue = 0
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewIntegerStackValue(0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mul operation between a boolean and a %s", s.GetType().String())
}

func (b BoolStackValue) Div(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := 0
	if b.value {
		intValue = 1
	}
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
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" inner value is empty")
			}
			return b.Div(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		rhsIntValue := intV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue / rhsIntValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		rhsIntValue := doubleV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(int(intValue) / int(rhsIntValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible div operation between a boolean and a %s", s.GetType().String())

}
func (b BoolStackValue) Mod(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	intValue := 0
	if b.value {
		intValue = 1
	}
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
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" inner value is empty")
			}
			return b.Mod(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		rhsIntValue := intV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue % rhsIntValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		rhsIntValue := doubleV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(int(intValue) % int(rhsIntValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mod operation between a boolean and a %s", s.GetType().String())
}
