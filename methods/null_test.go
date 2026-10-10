package methods

import (
	"fmt"
	"testing"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

// Calling any method on null is an error, with one message, whatever the method is.
func TestMethodOnNullIsAnError(t *testing.T) {
	names := []string{"toString", "toInteger", "toBoolean", "toDouble", "length", // valid on every receiver
		"trim", "contains", "containsString", "findByXPATH", // String methods
		"noSuchMethod"} // not a method at all: still the null rule, not "has no method"
	for _, name := range names {
		r, err := Invoke(mk(stackvalue.NULL, ""), name, nil)
		want := fmt.Sprintf("cannot call method %q on null", name)
		if err == nil || r != nil || err.Error() != want {
			t.Errorf("null.%s(): got result=%v err=%v, want error %q", name, r, err, want)
		}
	}
}

// The rule must not leak: other receivers behave as before.
func TestNullRuleDoesNotAffectOtherReceivers(t *testing.T) {
	if r, err := Invoke(str("  x  "), "trim", nil); err != nil || r.ToString() != "x" {
		t.Errorf("string.trim() broke: %v %v", r, err)
	}
	if _, err := Invoke(mk(stackvalue.INTEGER, "5"), "toString", nil); err != nil {
		t.Errorf("integer.toString() broke: %v", err)
	}
}
