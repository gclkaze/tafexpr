package tafexpr

import (
	"io"
	"os"
	"testing"

	"github.com/gclkaze/tafexpr/variablecontext"
	"github.com/stretchr/testify/assert"
)

// captureStdout gives what f writes to stdout.
func captureStdout(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	out := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		out <- string(b)
	}()
	old := os.Stdout
	os.Stdout = w
	func() {
		defer func() {
			os.Stdout = old
			w.Close()
		}()
		f()
	}()
	s := <-out
	r.Close()
	return s
}

// BL-270: parsing never writes to stdout, where a script's output goes.
func TestParsingNeverWritesToStdout(t *testing.T) {
	expressions := []string{
		`"plain"`, `""`, `"say \"BL270\" twice"`, `"\"BL270\""`, `"{\"token\": \"BL270\"}"`,
		`"a \"b\"" + "c \"d\""`, `"value {{$i}} and \"BL270\""`,
		`{"token": "say \"BL270\""}`, `["BL270", "\"x\"", 1]`,
		`1 + 2 * 3`, `1.5`, `true`, `null`, `$i`, `$i + 1`, `$i == 60`,
		`"unterminated`, `1 +`, `$undeclared`, ``, // errors stay off stdout too
	}
	for _, e := range expressions {
		t.Run(e, func(t *testing.T) {
			p := SetupParser()
			p.VariableContext = SetupMockVariableContext()
			p.VariableContext.(*variablecontext.MockVariableContext).SetValue("$i", 60)
			assert.Empty(t, captureStdout(t, func() { p.Parse(e) }), "Parse")
			assert.Empty(t, captureStdout(t, func() { p.Interpolate(e) }), "Interpolate")
			assert.Empty(t, captureStdout(t, func() { Validate(e) }), "Validate")
		})
	}
}
