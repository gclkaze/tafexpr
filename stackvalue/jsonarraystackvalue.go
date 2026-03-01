package stackvalue

import (
	"fmt"
	"strings"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	log "github.com/sirupsen/logrus"

	"github.com/gclkaze/evalang/evalangparser/globals"
)

type JSONArrayStackValue struct {
	//	value globals.JSONArray

	value []JSONStackValue
	ptr   *globals.JSONArrayGen
}

func NewJSONArrayStackValue(v globals.JSONArrayGen) *JSONArrayStackValue {
	inst := &JSONArrayStackValue{}
	for i := 0; i < len(v); i++ {
		inst.value = append(inst.value, *NewJSONStackValue(v[i]))
	}
	inst.ptr = &v
	return inst
}

func (s JSONArrayStackValue) ToUintPtr() uintptr {
	return uintptr(0)
}

func NewJSONArrayStackValueFromStackvalue(st *[]StackValue) *globals.JSONArrayGen {
	inst := globals.JSONArrayGen{}
	for i := 0; i < len(*st); i++ {
		p := (*st)[i]
		pp := GetGenericValue(p)
		inst = append(inst, pp)
	}
	return &inst
}

func (s JSONArrayStackValue) ToVargs() ([]uintptr, error) {
	var ptrs []uintptr
	for i := range s.value {
		value := s.value[i]
		inner, err := value.GetInnerValue()
		if err != nil {
			return nil, err
		}
		ptrs = append(ptrs, inner.ToUintPtr())
	}
	return ptrs, nil
}

func (s JSONArrayStackValue) ToInteger() (result int64, err error) {
	return 0, fmt.Errorf("cannot convert JSONArray to integer")
}

func (s JSONArrayStackValue) ToDouble() (result float64, err error) {
	return 0, fmt.Errorf("cannot convert JSONArray to double")
}

func (s JSONArrayStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s JSONArrayStackValue) ToString() string {
	var ar []string
	for i := 0; i < len(s.value); i++ {
		s := s.value[i].ToString()
		ar = append(ar, s)
	}
	//	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonn := "[" + strings.Join(ar, ",") + "]"

	return jsonn
}

func (s JSONArrayStackValue) Length() (result int, err error) {
	return len(s.value), nil
}

func (s JSONArrayStackValue) IsScalar() bool {
	return false
}

func (s JSONArrayStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}

func NewEmptyJSONArrayStackValue() *JSONArrayStackValue {
	inst := &JSONArrayStackValue{}
	return inst
}

func (s *JSONArrayStackValue) Push(v globals.JSONObjectGen) {
	s.value = append(s.value, *NewJSONStackValue(v))
}

func (s JSONArrayStackValue) GetPtr() *globals.JSONArrayGen {
	return s.ptr
}

func (s JSONArrayStackValue) GetPtrGen() *globals.JSONArrayGen {
	var a globals.JSONArrayGen = globals.JSONArrayGen{}
	for i := 0; i < len(*s.ptr); i++ {
		one := (*s.ptr)[i]
		a = append(a, one)
	}
	return &a
}

func (s JSONArrayStackValue) GetType() StackValueType {
	return JSON_ARRAY
}

func (s JSONArrayStackValue) GetValue() globals.JSONArrayGen {
	return *s.ptr
}
func (s JSONArrayStackValue) Copy() StackValue {
	inst := &JSONArrayStackValue{}
	for i := 0; i < len(s.value); i++ {
		inst.value = append(inst.value, *NewJSONStackValue(s.value[i].GetValue()))
	}
	return inst
}

func (s JSONArrayStackValue) IsTruthy() bool {
	return s.value != nil && len(s.value) != 0
}

func (s JSONArrayStackValue) Equals(other StackValue) bool {
	if other.GetType() != JSON_ARRAY {
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
	otherArr, ok := other.(*JSONArrayStackValue)
	if !ok {
		return false
	}
	if len(s.value) != len(otherArr.value) {
		return false
	}
	for i := 0; i < len(s.value); i++ {
		if !s.value[i].Equals(otherArr.value[i]) {
			return false
		}
	}
	return true
}
