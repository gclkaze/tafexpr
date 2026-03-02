package stackvalue

import (
	"fmt"

	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"

	"github.com/gclkaze/evalang-globals/globals"
	"github.com/gclkaze/evalang-globals/globals/stackvalue"
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

func (s ReferenceStackValue) GetType() stackvalue.StackValueType {
	return stackvalue.REFERENCE
	/*	p := s.value
		if p == nil || *p == nil {
			return REFERENCE
		}

		t := (*p).GetType()
		switch t {
		case stackvalue.globals.STRING:
			return STRING
		case stackvalue.globals.INTEGER:
			return INTEGER
		case stackvalue.globals.DOUBLE:
			return DOUBLE
		case stackvalue.globals.JSON:
			return JSON_OBJECT
		case stackvalue.globals.JSON_ARRAY:
			return JSON_ARRAY
		case stackvalue.globals.BOOLEAN:
			return BOOL
		case stackvalue.globals.DB:
			return DB
		case stackvalue.globals.BROWSER:
			return BROWSER
		case stackvalue.globals.NULL:
			return NULL
		case stackvalue.globals.REFERENCE:
			return REFERENCE
		case stackvalue.globals.HTML_ELEM:
		case stackvalue.globals.DATE:
		case stackvalue.globals.VARIABLE:
		case stackvalue.globals.VARIABLE_STORAGE:
		case stackvalue.globals.NUMBER:
		case stackvalue.globals.LIST:
		case stackvalue.globals.CALCULATED:
		case stackvalue.globals.EXPRESSION:
		}
		return REFERENCE*/
}

func (s ReferenceStackValue) Copy() stackvalue.StackValue {
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

func (s ReferenceStackValue) Equals(other stackvalue.StackValue) bool {
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
func (s ReferenceStackValue) Add(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" add operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Sub(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" sub operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Mul(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" mul operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Div(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" div operation does not apply to REFERENCE and %s", t.String())
}
func (s ReferenceStackValue) Mod(other stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	t := s.GetType()
	return nil, tafargumentlistenererrortypes.CONVERSION_ERROR, fmt.Errorf(" mod operation does not apply to REFERENCE and %s", t.String())
}
