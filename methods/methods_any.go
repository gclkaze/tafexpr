package methods

import (
	"strings"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	mine "github.com/gclkaze/tafexpr/stackvalue"
)

// Methods valid on every receiver. Ported unchanged from the legacy handlers
// ExitHandleToString / ToInteger / ToBoolean / ToDouble / Length.
func init() {
	Register("toString", Method{
		Group: "convert", Returns: "String", Doc: "The value as text.",
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			return mine.NewStringStackValue(c.Recv.ToString()), nil
		},
	})

	Register("toInteger", Method{
		Group: "convert", Returns: "Integer", Doc: "The value as an integer; fails when it has none.",
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			i, err := c.Recv.ToInteger()
			if err != nil {
				return nil, err
			}
			return mine.NewIntegerStackValue(int(i)), nil
		},
	})

	Register("toBoolean", Method{
		Group: "convert", Returns: "Bool", Doc: "The value as a boolean.",
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			b, err := c.Recv.ToBoolean()
			if err != nil {
				return nil, err
			}
			return mine.NewBoolStackValue(b), nil
		},
	})

	Register("toDouble", Method{
		Group: "convert", Returns: "Double", Doc: "The value as a double; fails when it has none.",
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			d, err := c.Recv.ToDouble()
			if err != nil {
				return nil, err
			}
			return mine.NewDoubleStackValue(d), nil
		},
	})

	Register("length", Method{
		Group: "convert", Returns: "Integer", Doc: "Length of a string, list or object; fails for types without one.",
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			n, err := c.Recv.Length()
			if err != nil {
				return nil, err
			}
			return mine.NewIntegerStackValue(n), nil
		},
	})

	// The data type of the value: lowercase, one word (string, integer, double, bool, jsonobject,
	// jsonarray, userdefined, ...). The names are part of the language, because scripts compare
	// against them. They are StackValueType.String() in lowercase. Null has no type: like every
	// method, type() on null is an error ($x == null is how a script tests for null).
	Register("type", Method{
		Group: "introspection", Returns: "String",
		Doc: "The data type of the value, lowercase and one word: string, integer, double, bool, jsonobject, jsonarray, userdefined. Fails on null.",
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			return mine.NewStringStackValue(strings.ToLower(c.Recv.GetType().String())), nil
		},
	})
}
