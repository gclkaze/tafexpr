package methods

import (
	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/utils"
	mine "github.com/gclkaze/tafexpr/stackvalue"
)

var stringOnly = []stackvalue.StackValueType{stackvalue.STRING}

// textFn adapts a func(string) (string, error) helper to a method body on a String receiver.
func textFn(f func(string) (string, error)) func(*Call) (stackvalue.StackValue, error) {
	return func(c *Call) (stackvalue.StackValue, error) {
		s, err := f(c.Recv.ToString())
		if err != nil {
			return nil, err
		}
		return mine.NewStringStackValue(s), nil
	}
}

// testFn adapts a func(string, string) (bool, error) helper taking the receiver and argument 1.
func testFn(f func(string, string) (bool, error)) func(*Call) (stackvalue.StackValue, error) {
	return func(c *Call) (stackvalue.StackValue, error) {
		ok, err := f(c.Recv.ToString(), c.Args[0].ToString())
		if err != nil {
			return nil, err
		}
		return mine.NewBoolStackValue(ok), nil
	}
}

// String methods. Ported from the legacy handlers; each keeps its utils call.
// Parity only: new behaviour (e.g. trim's optional cutset) is a later, separate step.
func init() {
	Register("trim", Method{Group: "string", Returns: "String", Doc: "Removes leading and trailing whitespace.",
		Fn: textFn(utils.Trim)}, stackvalue.STRING)
	Register("trimLeft", Method{Group: "string", Returns: "String", Doc: "Removes leading whitespace.",
		Fn: textFn(utils.TrimLeft)}, stackvalue.STRING)
	Register("trimRight", Method{Group: "string", Returns: "String", Doc: "Removes trailing whitespace.",
		Fn: textFn(utils.TrimRight)}, stackvalue.STRING)

	Register("contains", Method{Group: "string", Returns: "Bool", Doc: "True when the text contains the argument.",
		Params: []Param{{Name: "text", Types: stringOnly}},
		Fn:     testFn(utils.ContainsString)}, stackvalue.STRING)
	Register("startsWith", Method{Group: "string", Returns: "Bool", Doc: "True when the text starts with the argument.",
		Params: []Param{{Name: "prefix", Types: stringOnly}},
		Fn:     testFn(utils.StarsWith)}, stackvalue.STRING) // sic: the helper is spelled StarsWith
	Register("endsWith", Method{Group: "string", Returns: "Bool", Doc: "True when the text ends with the argument.",
		Params: []Param{{Name: "suffix", Types: stringOnly}},
		Fn:     testFn(utils.EndsWith)}, stackvalue.STRING)

	Register("replaceAll", Method{Group: "string", Returns: "String", Doc: "Replaces every occurrence of old with new.",
		Params: []Param{{Name: "old", Types: stringOnly}, {Name: "new", Types: stringOnly}},
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			return mine.NewStringStackValue(utils.ReplaceAllString(c.Recv.ToString(), c.Args[0].ToString(), c.Args[1].ToString())), nil
		}}, stackvalue.STRING)

	Register("extract", Method{Group: "string", Returns: "String", Doc: "The first match of a regular expression.",
		Params: []Param{{Name: "regex", Types: stringOnly}},
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			s, err := utils.ExtractOneFromRegex(c.Recv.ToString(), c.Args[0].ToString())
			if err != nil {
				return nil, err
			}
			return mine.NewStringStackValue(s), nil
		}}, stackvalue.STRING)
}
