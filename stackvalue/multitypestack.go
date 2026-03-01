package stackvalue

import (
	"github.com/gclkaze/evalang/evalangparser/globals"
	log "github.com/sirupsen/logrus"
)

type MultiTypeStack struct {
	stack     []StackValue
	onError   bool
	hierarchy []*[]StackValue
	current   *[]StackValue
}

func (s MultiTypeStack) OnError() bool {
	return s.onError
}

func NewMultiTypeStack() *MultiTypeStack {
	inst := &MultiTypeStack{}
	inst.current = &inst.stack
	inst.hierarchy = append(inst.hierarchy, inst.current)
	return inst
}

func (s MultiTypeStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *MultiTypeStack) Push(st StackValue) {
	*s.current = append(*s.current, st)
}

func (s *MultiTypeStack) CreateStack() {

	//var st StackValue = NewEmptyJSONArrayStackValue()
	newStack := &[]StackValue{}
	s.current = newStack

	//*s.current = append(*s.current, NewEmptyJSONArrayStackValue())
	s.hierarchy = append(s.hierarchy, s.current)
}

func (s *MultiTypeStack) ExitStack() *globals.JSONArrayGen {
	ptr := s.hierarchy[len(s.hierarchy)-1]

	//pop hierarchy
	s.hierarchy = s.hierarchy[:len(s.hierarchy)-1]
	s.current = s.hierarchy[len(s.hierarchy)-1]

	//fmt.Printf("ptr: %v\n", ptr)
	return NewJSONArrayStackValueFromStackvalue(ptr)
}

func (s *MultiTypeStack) PushInteger(i int) {
	*s.current = append(*s.current, NewIntegerStackValue(i))
}

func (s *MultiTypeStack) PushDouble(i float64) {
	*s.current = append(*s.current, NewDoubleStackValue(i))
}

func (s *MultiTypeStack) PushNull() {
	*s.current = append(*s.current, NewNullStackValue())
}

func (s *MultiTypeStack) PushString(st string) {
	*s.current = append(*s.current, NewStringStackValue(st))
}

func (s *MultiTypeStack) PushBool(b bool) {
	*s.current = append(*s.current, NewBoolStackValue(b))
}

func (s *MultiTypeStack) Pop() StackValue {
	if len(*s.current) < 1 {
		s.onError = true
		return nil
		//		panic("stack is empty unable to pop")
	}
	// Get the last value from the stack.
	result := (*s.current)[len(*s.current)-1]

	// Remove the last element from the stack.
	*s.current = (*s.current)[:len(*s.current)-1]
	log.Debug("popped ", s.stack)
	return result
}
