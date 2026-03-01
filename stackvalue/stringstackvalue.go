package stackvalue

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	tafargumenetlistenererrortypes "github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	log "github.com/sirupsen/logrus"
)

type StringStackValue struct {
	value string
}

func NewStringStackValue(v string) *StringStackValue {
	return &StringStackValue{value: v}
}

func (s StringStackValue) GetType() StackValueType {
	return STRING
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
func (s StringStackValue) Copy() StackValue {
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
func (s StringStackValue) Equals(other StackValue) bool {
	if other.GetType() != STRING {
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
	strV, ok := other.(*StringStackValue)
	if !ok {
		return false
	}
	return s.value == strV.value
}

func (s StringStackValue) Add(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible add operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Sub(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible sub operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Mul(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mul operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Div(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible div operation between a string and a %s", s.GetType().String())
}
func (s StringStackValue) Mod(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mod operation between a string and a %s", s.GetType().String())
}
