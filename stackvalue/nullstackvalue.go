package stackvalue

import (
	"fmt"

	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

type NullStackValue struct {
}

func NewNullStackValue() *NullStackValue {
	return &NullStackValue{}
}

func (s NullStackValue) GetType() stackvalue.StackValueType {
	return stackvalue.NULL
}
func (s NullStackValue) ToUintPtr() uintptr {
	return uintptr(0)
}

func (s NullStackValue) Copy() stackvalue.StackValue {
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

func (s NullStackValue) Equals(other stackvalue.StackValue) bool {
	if other.GetType() != stackvalue.NULL {
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
func (s NullStackValue) Add(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Sub(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Mul(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Div(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
func (s NullStackValue) Mod(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return other.Copy(), tafargumentlistenererrortypes.NONE, nil
}
