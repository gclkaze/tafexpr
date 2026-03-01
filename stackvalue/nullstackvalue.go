package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

type NullStackValue struct {
}

func NewNullStackValue() *NullStackValue {
	return &NullStackValue{}
}

func (s NullStackValue) GetType() StackValueType {
	return NULL
}
func (s NullStackValue) ToUintPtr() uintptr {
	return uintptr(0)
}

func (s NullStackValue) Copy() StackValue {
	return NewNullStackValue()
}
func (s NullStackValue) ToInteger() (result int64, err error) {
	return 0, nil
}

func (s NullStackValue) ToDouble() (result float64, err error) {
	return float64(0), nil
}

func (s NullStackValue) ToBoolean() (result bool, err error) {
	return false, nil
}

func (s NullStackValue) IsTruthy() bool {
	return false
}

func (s NullStackValue) ToString() string {
	return "null"
}

func (s NullStackValue) Length() (result int, err error) {
	return 0, fmt.Errorf("Length does not apply to null value")
}

func (s NullStackValue) IsScalar() bool {
	return true
}

func (s NullStackValue) Equals(other StackValue) bool {
	if other.GetType() != NULL {
		inner, ok := other.(*JSONStackValue)
		if !ok {
			return false
		}
		isNull := !inner.IsTruthy()
		return isNull
	}
	_, ok := other.(*NullStackValue)
	return ok
}

func (s NullStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}

func (s NullStackValue) GetValue() interface{} {
	return nil
}
func (s NullStackValue) Add(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Sub(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Mul(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Div(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Mod(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
