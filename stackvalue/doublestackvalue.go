package stackvalue

import (
	"fmt"
	"math"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
)

type DoubleStackValue struct {
	value float64
}

func NewDoubleStackValue(i float64) *DoubleStackValue {
	return &DoubleStackValue{value: i}
}

func (s DoubleStackValue) ToInteger() (result int64, err error) {
	return int64(s.value), nil
}

func (s DoubleStackValue) ToDouble() (result float64, err error) {
	return float64(s.value), nil
}

func (s DoubleStackValue) ToUintPtr() uintptr {
	bits := math.Float64bits(s.value)
	ptr := uintptr(bits)
	return ptr
}

func (s DoubleStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s DoubleStackValue) Length() (result int, err error) {
	return -1, fmt.Errorf("Length does not apply to double value")
}

func (s DoubleStackValue) IsScalar() bool {
	return true
}

func (s DoubleStackValue) ToString() string {
	return strconv.FormatFloat(s.value, 'f', -1, 64)
}

func (s DoubleStackValue) GetType() stackvalue.StackValueType {
	return stackvalue.DOUBLE
}

func (s DoubleStackValue) GetValue() float64 {
	return s.value
}

func (s DoubleStackValue) Copy() stackvalue.StackValue {
	return NewDoubleStackValue(s.value)
}
func (s DoubleStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}
func (s DoubleStackValue) Equals(other stackvalue.StackValue) bool {
	if other.GetType() != stackvalue.DOUBLE {
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
	dV, ok := other.(*DoubleStackValue)
	if !ok {
		return false
	}
	return s.value == dV.value
}

func (s DoubleStackValue) IsTruthy() bool {
	return s.value != 0
}

func (b DoubleStackValue) Add(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Add(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewDoubleStackValue(float64(intV.GetValue()) + doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewDoubleStackValue(doubleV.GetValue() + doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if boolV.GetValue() {
			doubleValue++
		}
		return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible add operation between a double and a %s", s.GetType().String())
}
func (b DoubleStackValue) Sub(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Sub(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewDoubleStackValue(doubleValue - float64(intV.GetValue())), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewDoubleStackValue(doubleValue - doubleV.GetValue()), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if boolV.GetValue() {
			doubleValue--
		}
		return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible sub operation between a double and a %s", s.GetType().String())
}

func (b DoubleStackValue) Mul(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
				return NewDoubleStackValue(0), tafargumenetlistenererrortypes.NONE, nil
			}
			return b.Mul(inner)
		}
	case stackvalue.INTEGER:
		intV, ok := s.(*IntegerStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to integer")
		}
		return NewDoubleStackValue(float64(intV.GetValue()) * doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		return NewDoubleStackValue(doubleV.GetValue() * doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			doubleValue = 0
		}
		return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return NewDoubleStackValue(0), tafargumenetlistenererrortypes.NONE, nil
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mul operation between a double and a %s", s.GetType().String())

}
func (b DoubleStackValue) Div(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := b.value
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
		return NewDoubleStackValue(doubleValue / float64(rhsIntValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		rhsDoubleValue := doubleV.GetValue()
		if rhsDoubleValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewDoubleStackValue(doubleValue / rhsDoubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewDoubleStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible div operation between a double and a %s", s.GetType().String())
}
func (b DoubleStackValue) Mod(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumenetlistenererrortypes.TAFArgumentListenerError, err error) {
	doubleValue := int(b.value)
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
		return NewIntegerStackValue(doubleValue % rhsIntValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.DOUBLE:
		doubleV, ok := s.(*DoubleStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to double")
		}
		rhsIntValue := doubleV.GetValue()
		if rhsIntValue == 0 {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(int(doubleValue) % int(rhsIntValue)), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.BOOL:
		boolV, ok := s.(*BoolStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to bool")
		}
		if !boolV.GetValue() {
			return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
		}
		return NewIntegerStackValue(doubleValue), tafargumenetlistenererrortypes.NONE, nil
	case stackvalue.NULL:
		_, ok := s.(*NullStackValue)
		if !ok {
			return nil, tafargumenetlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" couldn't convert stack value to null")
		}

		return nil, tafargumenetlistenererrortypes.DIVISION_BY_ZERO, fmt.Errorf(" division by zero ")
	}
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mod operation between a double and a %s", s.GetType().String())
}
