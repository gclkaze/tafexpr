package stackvalue

/*json value*/
import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	jsoniter "github.com/json-iterator/go"
	"github.com/nsf/jsondiff"
	log "github.com/sirupsen/logrus"

	//	"github.com/gclkaze/evalang/evalangparser/globals"

	"github.com/gclkaze/evalang-globals/globals"
)

type JSONStackValue struct {
	value globals.JSONObjectGen
}

func NewJSONStackValue(v globals.JSONObjectGen) *JSONStackValue {
	return &JSONStackValue{value: v}
}

func NewEmptyJSONStackValue() *JSONStackValue {
	return &JSONStackValue{}
}

func (s JSONStackValue) ToInteger() (result int64, err error) {
	return 0, fmt.Errorf("cannot convert JSON to integer")
}

func (s JSONStackValue) ToDouble() (result float64, err error) {
	return 0, fmt.Errorf("cannot convert JSON to double")
}

func (s JSONStackValue) ToUintPtr() uintptr {
	p, ok := s.value.(uintptr)
	if !ok {
		return uintptr(0)
	}
	return p
}

func (s JSONStackValue) Flatten() (map[string]string, error) {
	if s.value == nil {
		return nil, nil
	}
	converted := make(map[string]string)
	cd, ok := s.value.(map[string]interface{})
	if !ok {
		return nil, nil
	}

	for k, v := range cd {
		if str, ok := v.(string); ok {
			converted[k] = str
			continue
		}
		if istr, ok := v.(int); ok {
			converted[k] = strconv.Itoa(istr)
			continue
		}
		if bstr, ok := v.(bool); ok {
			converted[k] = strconv.FormatBool(bstr)
			continue
		}
		if fstr, ok := v.(float64); ok {
			converted[k] = fmt.Sprintf("%.2f", fstr)
			continue
		}

	}

	return converted, nil
}

func (s JSONStackValue) ToBoolean() (result bool, err error) {
	return s.IsTruthy(), nil
}

func (s JSONStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}
func NewJSONStackValueFromStackValue(v globals.JSONObjectGen) *JSONStackValue {
	return NewJSONStackValue(v)
}

func GetGenericValue(v StackValue) globals.JSONObjectGen {
	t := v.GetType()
	switch t {
	case USER_DEFINED:
		i, ok := v.(*UserDefinedStackValue)
		if !ok {
			return nil
		}
		return i.GetValue()

	case INTEGER:
		i, ok := v.(*IntegerStackValue)
		if !ok {
			return nil
		}
		return i.GetValue()

	case BOOL:
		i, ok := v.(*BoolStackValue)
		if !ok {
			return nil
		}
		return i.GetValue()
	case DOUBLE:
		i, ok := v.(*DoubleStackValue)
		if !ok {
			return nil
		}
		return i.GetValue()

	case JSON_ARRAY:
		i, ok := v.(*JSONArrayStackValue)
		if !ok {
			return nil
		}
		return i.GetValue()
	case STRING:
		i, ok := v.(*StringStackValue)
		if !ok {
			return nil
		}
		return i.GetValue()
	case NULL:
		return nil
	case JSON_OBJECT:
		i, ok := v.(*JSONStackValue)
		if !ok {
			return nil
		}
		return i.GetValue()

	}
	return nil
}

func (s JSONStackValue) ToString() string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	v, err := json.Marshal(s.value)
	if err != nil {
		return ""
	}
	st := string(v)
	st = strings.ReplaceAll(st, "\"", "\\\"")
	return st
}

func (s JSONStackValue) Length() (result int, err error) {
	actualValue := s.value

	if actualValue == nil {
		return 0, nil
	}
	switch vt := actualValue.(type) {
	case int:
		return NewIntegerStackValue(vt).Length()
	case float64:
		return NewDoubleStackValue(vt).Length()
	case bool:
		return NewBoolStackValue(vt).Length()
	case string:
		return NewStringStackValue(vt).Length()
	case globals.JSONArrayGen:
		return NewJSONArrayStackValue(vt).Length()
	case globals.JSONObjectGen:
		castToJSON, ok := actualValue.(globals.JSONStruct)
		if !ok {
			return 0, fmt.Errorf("cannot cast object")
		}
		return len(castToJSON), nil
	default:
		return 0, fmt.Errorf("unknown type in stack value %v", vt)
	}
}

func (s JSONStackValue) IsScalar() bool {
	return false
}

func (s JSONStackValue) GetInnerValue() (st StackValue, err error) {
	actualValue := s.value

	if actualValue == nil {
		return nil, fmt.Errorf("json stack value is nil")
	}
	switch vt := actualValue.(type) {
	case int:
		return NewIntegerStackValue(vt), nil
	case float64:
		return NewDoubleStackValue(vt), nil
	case bool:
		return NewBoolStackValue(vt), nil
	case string:
		return NewStringStackValue(vt), nil
	case globals.JSONArrayGen:
		return NewJSONArrayStackValue(vt), nil
	case globals.JSONObjectGen:
		return NewJSONStackValue(vt), nil
	default:
		return nil, fmt.Errorf("unknown type in stack value %v", vt)
	}

}

func (s JSONStackValue) GetType() StackValueType {
	return JSON_OBJECT
}

func (s JSONStackValue) GetValue() globals.JSONObjectGen {
	return s.value
}

func (s *JSONStackValue) GetStorageValue() *globals.JSONObjectGen {
	return &s.value
}

func (s JSONStackValue) Copy() StackValue {
	return NewJSONStackValue(s.value)
}

func (s JSONStackValue) IsTruthy() bool {
	inner, err := s.GetInnerValue()
	return s.value != nil && err == nil && inner != nil
}

func (s JSONStackValue) Equals(other StackValue) bool {
	if other.GetType() != JSON_OBJECT {
		i, err := s.GetInnerValue()
		if err != nil {
			//check if other is null stack value
			_, ok := other.(*NullStackValue)
			return ok
		}
		b, e, x := i.Equal(other)
		if e == tafargumentlistenererrortypes.NONE && x == nil {
			bb := b.(*BoolStackValue)
			return bb.GetValue()
		} else {
			log.Println(x.Error())
			return false
		}
	}
	otherJ, ok := other.(*JSONStackValue)
	if !ok {
		otherJInner, ok := other.(JSONStackValue)
		if !ok {
			return false
		}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		currentString, err := json.Marshal(s.value)
		if err != nil {
			log.Errorf("%s", err.Error())
			return false
		}
		otherString, err := json.Marshal(otherJInner.value)
		if err != nil {
			log.Errorf("%s", err.Error())
			return false
		}
		opt := jsondiff.DefaultJSONOptions()

		diff, diffString := jsondiff.Compare([]byte(currentString), []byte(otherString), &opt)
		if diff == jsondiff.FullMatch {
			return true
		}
		log.Debug("JSON DIFF " + diffString)
		return false //s.value == strV.value

	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	currentString, err := json.Marshal(s.value)
	if err != nil {
		log.Errorf("%s", err.Error())
		return false
	}
	otherString, err := json.Marshal(otherJ.value)
	if err != nil {
		log.Errorf("%s", err.Error())
		return false
	}
	opt := jsondiff.DefaultJSONOptions()

	diff, diffString := jsondiff.Compare([]byte(currentString), []byte(otherString), &opt)
	if diff == jsondiff.FullMatch {
		return true
	}
	log.Debug("JSON DIFF " + diffString)
	return false //s.value == strV.value
}
