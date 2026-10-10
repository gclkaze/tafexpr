package methods

import (
	"strings"
	"testing"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

// The names type() returns are part of the language: scripts compare against them. They are
// StackValueType.String() in lowercase, one word, no whitespace.
func TestTypeReturnsLowercaseOneWordNames(t *testing.T) {
	for typ, want := range map[stackvalue.StackValueType]string{
		stackvalue.STRING: "string", stackvalue.INTEGER: "integer", stackvalue.DOUBLE: "double",
		stackvalue.BOOL: "bool", stackvalue.JSON_OBJECT: "jsonobject", stackvalue.JSON_ARRAY: "jsonarray",
		stackvalue.USER_DEFINED: "userdefined",
	} {
		r, err := Invoke(mk(typ, ""), "type", nil)
		if err != nil || r.ToString() != want {
			t.Errorf("%v: want %q, got %v err=%v", typ, want, r, err)
			continue
		}
		if want != strings.ToLower(want) || strings.ContainsAny(want, " \t\n") {
			t.Errorf("%q must be lowercase with no whitespace", want)
		}
	}
}

// Null is not a data type: it has no type() either, like every other method.
func TestTypeOnNullIsAnError(t *testing.T) {
	r, err := Invoke(mk(stackvalue.NULL, ""), "type", nil)
	if err == nil || r != nil || err.Error() != `cannot call method "type" on null` {
		t.Errorf("got result=%v err=%v", r, err)
	}
}

func TestTypeTakesNoArguments(t *testing.T) {
	if _, err := Invoke(str("x"), "type", []stackvalue.StackValue{str("y")}); err == nil {
		t.Error("type() with an argument must be an error")
	}
	if err := CheckMethod("type", 0); err != nil {
		t.Errorf("type() must pass the compile-time check: %v", err)
	}
	if err := CheckMethod("type", 1); err == nil {
		t.Error("type(1) must fail the compile-time check")
	}
}
