package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

type UserDefinedStackValue struct {
	ptr uintptr
}

func NewUserDefinedStackValue(ptr uintptr) *UserDefinedStackValue {
	//	bptr := (*globals.ParameterValue)(unsafe.Pointer(ptr))
	return &UserDefinedStackValue{ptr: ptr}
}

func (s UserDefinedStackValue) ToInteger() (result int64, err error) {
	if s.ptr == 0 {
		return 0, nil
	}
	return 1, nil
}

func (s UserDefinedStackValue) ToUintPtr() uintptr {
	return s.ptr
}

func (s UserDefinedStackValue) ToDouble() (result float64, err error) {
	if s.ptr == 0 {
		return 0, nil
	}
	return 1, nil
}

func (s UserDefinedStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s UserDefinedStackValue) GetType() StackValueType {
	return USER_DEFINED
}

func (s UserDefinedStackValue) Copy() StackValue {
	return NewUserDefinedStackValue(s.ptr)
}

func (s UserDefinedStackValue) IsTruthy() bool {

	if s.ptr == 0 {
		return false
	}
	return true
	/*	p := s.ptr
		if p == 0 || *p == nil {
			return false
		}
		return !(*p).IsEmpty()*/
}

func (s UserDefinedStackValue) ToString() string {
	return "UserDefined"
}

func (s UserDefinedStackValue) Length() (result int, err error) {
	return 0, fmt.Errorf("Length does not apply to UserDefined value")
}

func (s UserDefinedStackValue) IsScalar() bool {
	return false
}

func (s UserDefinedStackValue) Equals(other StackValue) bool {
	b, _, _ := s.Equal(other)
	if b == nil {
		return false
	}
	bb, _ := b.(*BoolStackValue)
	return bb.GetValue()
}

func (s UserDefinedStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}

func (s *UserDefinedStackValue) GetValue() uintptr {
	return s.ptr
}
func (s UserDefinedStackValue) Add(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" add operation does not apply to UserDefined and %s", t.String())
}
func (s UserDefinedStackValue) Sub(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" sub operation does not apply to UserDefined and %s", t.String())
}
func (s UserDefinedStackValue) Mul(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" mul operation does not apply to UserDefined and %s", t.String())
}
func (s UserDefinedStackValue) Div(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" div operation does not apply to UserDefined and %s", t.String())
}
func (s UserDefinedStackValue) Mod(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" mod operation does not apply to UserDefined and %s", t.String())
}
