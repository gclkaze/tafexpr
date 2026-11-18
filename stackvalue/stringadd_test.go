package stackvalue

import (
	"testing"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

// BL-197: "text" + x is not text concatenation. The error names the types of
// both operands and says to build text with {{...}}.
func TestAddingToAStringNamesBothTypesAndTheWayToBuildText(t *testing.T) {
	const hint = `: + does not join text; use {{...}} to build text, e.g. "{{$a}}{{$b}}"`
	for _, c := range []struct {
		other stackvalue.StackValue
		want  string
	}{
		{NewIntegerStackValue(1), "incompatible add operation between a String and an Integer" + hint},
		{NewStringStackValue("b"), "incompatible add operation between a String and a String" + hint},
		{NewDoubleStackValue(1.5), "incompatible add operation between a String and a Double" + hint},
		{NewBoolStackValue(true), "incompatible add operation between a String and a Bool" + hint},
		{NewEmptyJSONStackValue(), "incompatible add operation between a String and a JSONObject" + hint},
		{NewEmptyJSONArrayStackValue(), "incompatible add operation between a String and a JSONArray" + hint},
		{NewNullStackValue(), "incompatible add operation between a String and Null" + hint},
	} {
		result, _, err := NewStringStackValue("a").Add(c.other)
		if result != nil || err == nil || err.Error() != c.want {
			t.Errorf("%s: got %v, %v; want %q", c.other.GetType(), result, err, c.want)
		}
	}
}
