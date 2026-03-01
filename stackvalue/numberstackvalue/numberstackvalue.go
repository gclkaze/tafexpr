package numberstackvalue

import (
	"github.com/gclkaze/tafexpr/stackvalue"
	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
)

type NumberStackValue interface {
	Add(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Sub(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Mul(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Div(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Mod(s stackvalue.StackValue) (result stackvalue.StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
}
