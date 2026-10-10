package methods

import (
	"testing"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

// These use the REAL registrations: length and toString are valid on every receiver, trim and
// contains only on a String (trim takes no arguments, contains takes one).
func TestFieldReadErrorMessages(t *testing.T) {
	cases := []struct {
		name     string
		receiver string
		field    string
		parent   stackvalue.StackValue
		want     string
	}{
		// the method works on this type, and takes no arguments
		{"list, old-style length", "$arr", "length", mk(stackvalue.JSON_ARRAY, ""),
			`$arr is of type JSONArray, not an object, so it has no property "length". Did you mean the method "length"? Call it with parentheses: $arr.length()`},
		{"string, old-style trim", "$s", "trim", mk(stackvalue.STRING, "x"),
			`$s is of type String, not an object, so it has no property "trim". Did you mean the method "trim"? Call it with parentheses: $s.trim()`},
		{"integer, old-style toString", "$n", "toString", mk(stackvalue.INTEGER, "5"),
			`$n is of type Integer, not an object, so it has no property "toString". Did you mean the method "toString"? Call it with parentheses: $n.toString()`},
		{"object lacking a property that is a method", "$o", "length", mk(stackvalue.JSON_OBJECT, ""),
			`$o has no property "length". Did you mean the method "length"? Call it with parentheses: $o.length()`},
		{"nested receiver", "$a.b", "trim", mk(stackvalue.STRING, "x"),
			`$a.b is of type String, not an object, so it has no property "trim". Did you mean the method "trim"? Call it with parentheses: $a.b.trim()`},
		{"qualified property name", "$a", "length[0]", mk(stackvalue.JSON_ARRAY, ""),
			`$a is of type JSONArray, not an object, so it has no property "length[0]". Did you mean the method "length"? Call it with parentheses: $a.length()`},

		// the method works on this type but needs arguments: say how many
		{"string, contains needs 1 argument", "$s", "contains", mk(stackvalue.STRING, "x"),
			`$s is of type String, not an object, so it has no property "contains". Did you mean the method "contains"? It takes 1 argument: $s.contains(...)`},
		{"string, replaceAll needs 2 arguments", "$s", "replaceAll", mk(stackvalue.STRING, "x"),
			`$s is of type String, not an object, so it has no property "replaceAll". Did you mean the method "replaceAll"? It takes 2 arguments: $s.replaceAll(...)`},
		{"string, legacy alias keeps the name the user wrote", "$s", "containsString", mk(stackvalue.STRING, "x"),
			`$s is of type String, not an object, so it has no property "containsString". Did you mean the method "containsString"? It takes 1 argument: $s.containsString(...)`},

		// the method exists, but only for other types
		{"integer, trim is a String method", "$n", "trim", mk(stackvalue.INTEGER, "5"),
			`$n is of type Integer, not an object, so it has no property "trim". "trim" is a method, but it only works on String values`},
		{"object, trim is a String method", "$o", "trim", mk(stackvalue.JSON_OBJECT, ""),
			`$o has no property "trim". "trim" is a method, but it only works on String values`},
		{"list, contains is a String method", "$a", "contains", mk(stackvalue.JSON_ARRAY, ""),
			`$a is of type JSONArray, not an object, so it has no property "contains". "contains" is a method, but it only works on String values`},

		// no method of that name: no hint at all
		{"integer, a property that is no method", "$n", "foo", mk(stackvalue.INTEGER, "5"),
			`$n is of type Integer, not an object, so it has no property "foo"`},
		{"a free function is not a method", "$n", "randomDoubleInRange", mk(stackvalue.INTEGER, "5"),
			`$n is of type Integer, not an object, so it has no property "randomDoubleInRange"`},
		{"null: no hint, a call on null fails too", "$x", "length", mk(stackvalue.NULL, ""),
			`$x is null, so it has no property "length"`},

		// receiver unknown: go by whether any type has such a method
		{"receiver unknown, no-argument method", "$a", "length", nil,
			`cannot read property "length" of $a. Did you mean the method "length"? Call it with parentheses: $a.length()`},
		{"receiver unknown, method with arguments", "$a", "contains", nil,
			`cannot read property "contains" of $a. Did you mean the method "contains"? It takes 1 argument: $a.contains(...)`},
		{"receiver unknown, not a method", "$a", "zzz", nil,
			`cannot read property "zzz" of $a`},
	}
	for _, c := range cases {
		got := FieldReadError(c.receiver, c.field, c.parent)
		if got == nil || got.Error() != c.want {
			t.Errorf("%s:\n  got  %v\n  want %s", c.name, got, c.want)
		}
	}
}

func TestHasMethodFor(t *testing.T) {
	if !HasMethodFor(stackvalue.JSON_ARRAY, "length", 0) || !HasMethodFor(stackvalue.STRING, "trim", 0) {
		t.Error("length (any receiver) and trim (String) must be found")
	}
	if HasMethodFor(stackvalue.JSON_OBJECT, "trim", 0) {
		t.Error("trim must not be offered for an object")
	}
	if HasMethodFor(stackvalue.STRING, "trim", 2) {
		t.Error("the ported trim takes no arguments, so two must not be offered")
	}
	if !HasMethodFor(stackvalue.STRING, "containsString", 1) {
		t.Error("an alias must resolve")
	}
}

func TestReceiverTypes(t *testing.T) {
	types, every := ReceiverTypes("trim")
	if every || len(types) != 1 || types[0] != stackvalue.STRING {
		t.Errorf("trim: want [String] and not every, got %v every=%v", types, every)
	}
	types, every = ReceiverTypes("length")
	if !every || len(types) != 0 {
		t.Errorf("length: want every receiver and no specific types, got %v every=%v", types, every)
	}
	types, every = ReceiverTypes("containsString") // alias
	if every || len(types) != 1 || types[0] != stackvalue.STRING {
		t.Errorf("alias containsString must resolve to contains: got %v every=%v", types, every)
	}
	if types, every = ReceiverTypes("zzz"); every || len(types) != 0 {
		t.Errorf("unknown name must have no receivers, got %v every=%v", types, every)
	}
}
