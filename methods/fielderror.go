package methods

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

// FieldReadError explains why  receiver.field  (written WITHOUT parentheses, so it is a property
// read) could not be read. It says what the receiver actually holds. When the property name
// is also a method name the message depends on the receiver's type:
//
//   - the method works on this type:      Did you mean the method "trim"? Call it with parentheses: $s.trim()
//   - it works on this type but needs
//     arguments:                          Did you mean the method "contains"? It takes 1 argument: $s.contains(...)
//   - it exists, but only for other types: "trim" is a method, but it only works on String values
//
// The usual cause is a script still written in the old style ($list.length instead of
// $list.length()). parent is the value of the receiver, or nil when it could not be determined.
func FieldReadError(receiver, field string, parent stackvalue.StackValue) error {
	var msg string
	switch {
	case parent == nil:
		msg = fmt.Sprintf("cannot read property %q of %s", field, receiver)
	case parent.GetType() == stackvalue.JSON_OBJECT:
		msg = fmt.Sprintf("%s has no property %q", receiver, field)
	case parent.GetType() == stackvalue.NULL:
		msg = fmt.Sprintf("%s is null, so it has no property %q", receiver, field)
	default:
		msg = fmt.Sprintf("%s is of type %s, not an object, so it has no property %q", receiver, parent.GetType(), field)
	}
	return errors.New(msg + methodHint(receiver, field, parent))
}

func methodHint(receiver, field string, parent stackvalue.StackValue) string {
	name := field
	if i := strings.Index(name, "["); i >= 0 { // length[0] -> length
		name = name[:i]
	}

	if parent == nil { // type unknown: go by whether any receiver has such a method
		ms := methodsNamed(name)
		if len(ms) == 0 {
			return ""
		}
		if CheckMethod(name, 0) == nil {
			return callHint(receiver, name, true, 0, 0)
		}
		min, max := ms[0].bounds()
		return callHint(receiver, name, false, min, max)
	}

	t := parent.GetType()
	if t == stackvalue.NULL {
		return "" // a method call on null is an error anyway; do not suggest it
	}
	if m, ok := lookupFor(t, name); ok {
		min, max := m.bounds()
		return callHint(receiver, name, min == 0, min, max)
	}
	if types, every := ReceiverTypes(name); len(types) > 0 && !every {
		names := make([]string, len(types))
		for i, rt := range types {
			names[i] = rt.String()
		}
		return fmt.Sprintf(". %q is a method, but it only works on %s values", name, strings.Join(names, " or "))
	}
	return ""
}

func callHint(receiver, name string, noArgs bool, min, max int) string {
	if noArgs {
		return fmt.Sprintf(". Did you mean the method %q? Call it with parentheses: %s.%s()", name, receiver, name)
	}
	return fmt.Sprintf(". Did you mean the method %q? It takes %s: %s.%s(...)", name, argText(min, max), receiver, name)
}

func argText(min, max int) string {
	switch {
	case min == max && min == 1:
		return "1 argument"
	case min == max:
		return fmt.Sprintf("%d arguments", min)
	default:
		return fmt.Sprintf("%d to %d arguments", min, max)
	}
}
