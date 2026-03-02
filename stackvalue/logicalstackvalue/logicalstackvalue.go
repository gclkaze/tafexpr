package logicalstackvalue

import (
	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
)

type LogicalStackValue interface {
	LesserThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	LesserThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Equal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Unequal(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	GreaterThan(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	GreaterThanEqual(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)

	And(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Or(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Not() (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
}
