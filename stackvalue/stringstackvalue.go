package stackvalue

import (
	"fmt"
	"strconv"
	"unsafe"

	log "github.com/sirupsen/logrus"

	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

type StringStackValue struct {
	value string
}

func NewStringStackValue(v string) *StringStackValue {
	return &StringStackValue{value: v}
}

func (s StringStackValue) GetType() stackvalue.StackValueType {
	return stackvalue.STRING
}

func (s StringStackValue) Length() (result int, err error) {
	return len(s.value), nil
}

func (s StringStackValue) IsScalar() bool {
	return false
}

func (s StringStackValue) GetValue() string {
	return s.value
}
func (s StringStackValue) Copy() stackvalue.StackValue {
	return NewStringStackValue(s.value)
}

func (s StringStackValue) ToUintPtr() uintptr {
	if s.value == "" {
		return uintptr(0)
	}
	ptr := uintptr(unsafe.Pointer(&[]byte(s.value)[0]))
	return ptr
}

func (s StringStackValue) ToString() string {
	return s.value
}
func (s StringStackValue) ToInteger() (result int64, err error) {
	v, err := strconv.Atoi(s.value)
	return int64(v), err
}

func (s StringStackValue) ToDouble() (result float64, err error) {
	v, err := strconv.ParseFloat(s.value, 64)
	return v, err
}

func (s StringStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s StringStackValue) IsTruthy() bool {
	return s.value != ""
}
func (s StringStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}
func (s StringStackValue) Equals(other stackvalue.StackValue) bool {
	if other.GetType() != stackvalue.STRING {
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
	strV, ok := other.(*StringStackValue)
	if !ok {
		return false
	}
	return s.value == strV.value
}

func (s StringStackValue) Add(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible add operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Sub(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible sub operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Mul(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mul operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Div(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible div operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Mod(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mod operation between a string and a %s", s.GetType().String())
}
