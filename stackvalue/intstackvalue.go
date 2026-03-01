package stackvalue

import (
	"fmt"
	"strconv"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	log "github.com/sirupsen/logrus"
)

type IntegerStackValue struct {
	value int
}

func NewIntegerStackValue(i int) *IntegerStackValue {
	return &IntegerStackValue{value: i}
}

func (s IntegerStackValue) ToInteger() (result int64, err error) {
	return int64(s.value), nil
}

func (s IntegerStackValue) ToDouble() (result float64, err error) {
	return float64(s.value), nil
}

func (s IntegerStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s IntegerStackValue) ToUintPtr() uintptr {
	return uintptr(s.value)
}

func (s IntegerStackValue) GetType() StackValueType {
	return INTEGER
}

func (s IntegerStackValue) ToString() string {
	return strconv.Itoa(s.value)
}

func (s IntegerStackValue) GetValue() int {
	return s.value
}

func (s IntegerStackValue) Length() (result int, err error) {
	return -1, fmt.Errorf("Length does not apply to integer value")
}

func (s IntegerStackValue) IsScalar() bool {
	return true
}

func (s IntegerStackValue) IsTruthy() bool {
	return s.value != 0
}

func (s IntegerStackValue) ToJson() *JSONStackValue {
	j := NewJSONStackValue(s.GetValue())
	return j
}

func (s IntegerStackValue) Equals(other StackValue) bool {
	if other.GetType() != INTEGER {
		if other.GetType() == JSON_OBJECT {
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
	intV, ok := other.(*IntegerStackValue)
	if !ok {
		return false
	}
	return s.value == intV.value
}

func (s IntegerStackValue) Copy() StackValue {
	return NewIntegerStackValue(s.value)
}

func (b IntegerStackValue) Add(s StackValue) (result StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
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
				return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Add(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewIntegerStackValue(intV.GetValue() + intValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewDoubleStackValue((doubleV.GetValue()) + float64(intValue)), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if boolV.GetValue() {
			intValue++
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible add operation between an integer and a %s", s.GetType().String())
}
func (b IntegerStackValue) Sub(s StackValue) (result StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
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
				return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Sub(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewIntegerStackValue(intValue - intV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewDoubleStackValue(float64(intValue) - doubleV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if boolV.GetValue() {
			intValue--
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible sub operation between an integer and a %s", s.GetType().String())
}
func (b IntegerStackValue) Mul(s StackValue) (result StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
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
				return NewIntegerStackValue(0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Mul(inner)
		}

	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewIntegerStackValue(intV.GetValue() * intValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewDoubleStackValue(doubleV.GetValue() * float64(intValue)), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			intValue = 0
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewIntegerStackValue(0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mul operation between an integer and a %s", s.GetType().String())
}
func (b IntegerStackValue) Div(s StackValue) (result StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
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
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" inner value is empty")
			}
			return b.Div(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		rhsIntValue := intV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue / rhsIntValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		rhsIntValue := doubleV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewDoubleStackValue(float64(intValue) / rhsIntValue), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible div operation between an integer and a %s", s.GetType().String())

}
func (b IntegerStackValue) Mod(s StackValue) (result StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
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
				return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" inner value is empty")
			}
			return b.Mod(inner)
		}
	case INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		rhsIntValue := intV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue % rhsIntValue), tafargumenetlistenererrortypes.NONE, nil
	case DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		rhsIntValue := doubleV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(int(intValue) % int(rhsIntValue)), tafargumenetlistenererrortypes.NONE, nil
	case BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(intValue), tafargumenetlistenererrortypes.NONE, nil
	case NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mod operation between an integer and a %s", s.GetType().String())
}
