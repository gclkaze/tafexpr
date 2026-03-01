package stackvalue

import (
	"fmt"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"

	"github.com/gclkaze/evalang/evalangparser/globals"
)

type ReferenceStackValue struct {
	value *globals.ParameterValue
}

func NewReferenceStackValue(parameter *globals.ParameterValue) *ReferenceStackValue {
	return &ReferenceStackValue{value: parameter}
}

func (s ReferenceStackValue) ToInteger() (result int64, err error) {
	if s.value == nil {
		return 0, nil
	}
	return 1, nil
}

func (s ReferenceStackValue) ToUintPtr() uintptr {
	return uintptr(0)
}

func (s ReferenceStackValue) ToDouble() (result float64, err error) {
	if s.value == nil {
		return 0, nil
	}
	return 1, nil
}

func (s ReferenceStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s ReferenceStackValue) GetType() StackValueType {
	return REFERENCE
	/*	p := s.value
		if p == nil || *p == nil {
			return REFERENCE
		}

		t := (*p).GetType()
		switch t {
		case globals.STRING:
			return STRING
		case globals.INTEGER:
			return INTEGER
		case globals.DOUBLE:
			return DOUBLE
		case globals.JSON:
			return JSON_OBJECT
		case globals.JSON_ARRAY:
			return JSON_ARRAY
		case globals.BOOLEAN:
			return BOOL
		case globals.DB:
			return DB
		case globals.BROWSER:
			return BROWSER
		case globals.NULL:
			return NULL
		case globals.REFERENCE:
			return REFERENCE
		case globals.HTML_ELEM:
		case globals.DATE:
		case globals.VARIABLE:
		case globals.VARIABLE_STORAGE:
		case globals.NUMBER:
		case globals.LIST:
		case globals.CALCULATED:
		case globals.EXPRESSION:
		}
		return REFERENCE*/
}

func (s ReferenceStackValue) Copy() StackValue {
	return NewReferenceStackValue(s.value)
}

func (s ReferenceStackValue) IsTruthy() bool {

	if s.value == nil {
		return false
	}
	p := s.value
	if p == nil || *p == nil {
		return false
	}
	return !(*p).IsEmpty()
}

func (s ReferenceStackValue) ToString() string {
	return "reference"
}

func (s ReferenceStackValue) Length() (result int, err error) {
	return 0, fmt.Errorf("Length does not apply to null value")
}

func (s ReferenceStackValue) IsScalar() bool {
	return false
}

func (s ReferenceStackValue) Equals(other StackValue) bool {
	b, _, _ := s.Equal(other)
	if b == nil {
		return false
	}
	bb, _ := b.(*BoolStackValue)
	return bb.GetValue()
}

func (s ReferenceStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}

func (s *ReferenceStackValue) GetValue() *globals.ParameterValue {
	return s.value
}
func (s ReferenceStackValue) Add(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" add operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Sub(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" sub operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Mul(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" mul operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Div(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" div operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Mod(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" mod operation does not apply to REFERENCE and %s", t.String())
}
