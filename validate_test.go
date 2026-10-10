package tafexpr

import (
	"strings"
	"testing"
)

func TestValidateAcceptsValidExpressions(t *testing.T) {
	for _, e := range []string{
		`$s.trim()`, `[1,2,3].length()`, `{"a":1}.length()`,
		`$a == 1 && $b == 2`, `$a==1&&$b==2`,
		`randomDoubleInRange(1, 5)`, `"abc".contains("b")`,
		`$o.tags[0]`, `5-1`, `-$a`, `$i -1`,
		`$s.trim().contains("x")`, `$s.replaceAll("a", "b")`,
		`$s.containsString("x")`,           // a legacy alias still resolves
		`[$n, $d, $s]`, `"a b"`, `$s.trim`, // a property read is valid syntax
	} {
		if errs := Validate(e); len(errs) != 0 {
			t.Errorf("%s: unexpected errors %v", e, errs)
		}
	}
}

func TestValidateFindsMistakesBeforeTheScriptRuns(t *testing.T) {
	for _, c := range []struct{ expr, want string }{
		{`$s.trimm()`, `unknown method "trimm"`},
		{`$s.trimm()`, `did you mean`},
		{`$s.contains()`, `takes 1 argument`},
		{`$s.trim(1)`, `takes 0 argument`},
		{`$s.replaceAll("a")`, `takes 2 argument`},
		{`foo(1)`, `unknown function "foo"`},
		{`randomDoubleInRange(1)`, `takes 2 argument`},
		{`$s.contains($t.trimm())`, `unknown method "trimm"`}, // inside an argument
		{`$s.trim(`, `syntax error`},
		{`$a ==`, `syntax error`},
		{`[1,2,3].length`, `syntax error`}, // a literal cannot be followed by a bare .length
	} {
		errs := Validate(c.expr)
		if len(errs) == 0 {
			t.Errorf("%s: want an error containing %q, got none", c.expr, c.want)
			continue
		}
		var all []string
		for _, e := range errs {
			all = append(all, e.Error())
		}
		if !strings.Contains(strings.Join(all, " | "), c.want) {
			t.Errorf("%s: want an error containing %q, got %v", c.expr, c.want, all)
		}
	}
}

func TestValidateReportsEveryProblem(t *testing.T) {
	if errs := Validate(`$s.trimm().foo()`); len(errs) != 2 {
		t.Errorf("want 2 errors (trimm and foo), got %v", errs)
	}
}
