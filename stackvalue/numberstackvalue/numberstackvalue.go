package numberstackvalue

import (
	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
)

type NumberStackValue interface {
	Add(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Sub(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Mul(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Div(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Mod(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
}
