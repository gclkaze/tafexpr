package stackvalue

type BrowserStackValue struct {
}

func NewBrowserStackValue(v string) *BrowserStackValue {
	return &BrowserStackValue{}
}

/*
func (s BrowserStackValue) GetType() StackValueType {
	return STRING
}

func (s BrowserStackValue) Length() (result int, err error) {
	return len(s.value), nil
}

func (s BrowserStackValue) IsScalar() bool {
	return false
}

func (s BrowserStackValue) GetValue() string {
	return s.value
}
func (s BrowserStackValue) Copy() StackValue {
	return NewStringStackValue(s.value)
}

func (s BrowserStackValue) ToString() string {
	return s.value
}

func (s BrowserStackValue) IsTruthy() bool {
	return s.value != ""
}
func (s BrowserStackValue) ToJson() *JSONStackValue {
	return NewJSONStackValue(s.GetValue())
}
func (s BrowserStackValue) Equals(other StackValue) bool {
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

func (s BrowserStackValue) Add(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible add operation between a browser value and a %s", s.GetType().String())
}
func (s BrowserStackValue) Sub(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible sub operation between a browser value and a %s", s.GetType().String())
}
func (s BrowserStackValue) Mul(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mul operation between a browser value and a %s", s.GetType().String())
}
func (s BrowserStackValue) Div(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible div operation between a browser value and a %s", s.GetType().String())
}
func (s BrowserStackValue) Mod(other StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error) {
	return nil, tafargumenetlistenererrortypes.INVALID_OPERATION_ERROR, fmt.Errorf("incompatible mod operation between a browser value and a %s", s.GetType().String())
}
*/
