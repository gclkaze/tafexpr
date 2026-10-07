package methods

import (
	"strconv"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	"github.com/gclkaze/evalang-globals/utils"
	mine "github.com/gclkaze/tafexpr/stackvalue"
)

// findOne is the shared lookup behind every findOne*ByXPATH method: exactly one match, as text.
// A no-match used to PANIC inside the helper (index out of range); the registry's safeCall turns
// that into an error, and fixing the helper itself makes the message clearer.
func findOne(c *Call) (string, error) {
	return utils.ParseHTMLByXPATHAndGetOne(c.Recv.ToString(), c.Args[0].ToString())
}

var xpathParam = []Param{{Name: "xpath", Types: stringOnly}}

// HTML / XPath methods, ported from the legacy handlers.
func init() {
	Register("findByXPATH", Method{Group: "html", Returns: "JSONArray", Doc: "Every element matching the XPath.",
		Params: xpathParam,
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			res, err := utils.ParseHTMLByXPATH(c.Recv.ToString(), c.Args[0].ToString())
			if err != nil {
				return nil, err
			}
			var payload []interface{} // stays nil when nothing matched, exactly as the legacy handler built it
			for _, r := range res {
				payload = append(payload, r)
			}
			return mine.NewJSONArrayStackValue(payload), nil
		}}, stackvalue.STRING)

	Register("findOneByXPATH", Method{Group: "html", Returns: "String", Doc: "The single element matching the XPath; fails on none or several.",
		Params: xpathParam,
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			s, err := findOne(c)
			if err != nil {
				return nil, err
			}
			return mine.NewStringStackValue(s), nil
		}}, stackvalue.STRING)

	Register("findOneStringByXPATH", Method{Group: "html", Returns: "String", Doc: "Like findOneByXPATH, as a string.",
		Params: xpathParam,
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			s, err := findOne(c)
			if err != nil {
				return nil, err
			}
			return mine.NewStringStackValue(s), nil
		}}, stackvalue.STRING)

	Register("findOneIntegerByXPATH", Method{Group: "html", Returns: "Integer", Doc: "The single match, converted to an integer.",
		Params: xpathParam,
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			s, err := findOne(c)
			if err != nil {
				return nil, err
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			return mine.NewIntegerStackValue(i), nil
		}}, stackvalue.STRING)

	Register("findOneDoubleByXPATH", Method{Group: "html", Returns: "Double", Doc: "The single match, converted to a double.",
		Params: xpathParam,
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			s, err := findOne(c)
			if err != nil {
				return nil, err
			}
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, err
			}
			return mine.NewDoubleStackValue(f), nil
		}}, stackvalue.STRING)

	Register("findOneBooleanByXPATH", Method{Group: "html", Returns: "Bool", Doc: "The single match, converted to a boolean.",
		Params: xpathParam,
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			s, err := findOne(c)
			if err != nil {
				return nil, err
			}
			b, err := strconv.ParseBool(s)
			if err != nil {
				return nil, err
			}
			return mine.NewBoolStackValue(b), nil
		}}, stackvalue.STRING)
}
