package methods

import (
	"strings"
	"testing"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
	mine "github.com/gclkaze/tafexpr/stackvalue"
)

// These tests exercise the REAL registrations. They assert wiring (names, arity, argument
// order, result types, error vs panic), not the semantics of the utils helpers: those are
// pinned by the evalang corpus, which is the parity oracle for the cutover.

func s(v string) stackvalue.StackValue { return mine.NewStringStackValue(v) }
func i(v int) stackvalue.StackValue    { return mine.NewIntegerStackValue(v) }

func TestLegacyMethodsStillResolve(t *testing.T) {
	legacy := map[string]int{
		"length": 0, "toString": 0, "toBoolean": 0, "toInteger": 0, "toDouble": 0,
		"trim": 0, "trimLeft": 0, "trimRight": 0,
		"startsWith": 1, "endsWith": 1, "containsString": 1, "extractOneByREGEX": 1,
		"findByXPATH": 1, "findOneByXPATH": 1, "findOneStringByXPATH": 1,
		"findOneDoubleByXPATH": 1, "findOneIntegerByXPATH": 1, "findOneBooleanByXPATH": 1,
		"replaceAllStringOccurrences": 2,
	}
	for name, argc := range legacy {
		if err := CheckMethod(name, argc); err != nil {
			t.Errorf("%s with %d argument(s): %v", name, argc, err)
		}
		// and the wrong argument count must be rejected
		if err := CheckMethod(name, argc+3); err == nil {
			t.Errorf("%s accepted %d arguments", name, argc+3)
		}
	}
	if len(legacy) != 19 {
		t.Fatalf("the cutover rewrites 19 methods, the list has %d", len(legacy))
	}
	if err := CheckFunc("randomDoubleInRange", 2); err != nil {
		t.Error(err)
	}
}

func TestRealAliasesResolve(t *testing.T) {
	if bad := UnresolvedAliases(); len(bad) != 0 {
		t.Fatalf("aliases without a target: %v", bad)
	}
	for old, canon := range map[string]string{"containsString": "contains", "replaceAllStringOccurrences": "replaceAll", "extractOneByREGEX": "extract"} {
		if canonical(old) != canon {
			t.Errorf("%s should resolve to %s", old, canon)
		}
	}
}

func TestConversionsOnAnyReceiver(t *testing.T) {
	check := func(recv stackvalue.StackValue, name string, wantType stackvalue.StackValueType) stackvalue.StackValue {
		t.Helper()
		r, err := Invoke(recv, name, nil)
		if err != nil || r.GetType() != wantType {
			t.Fatalf("%s.%s: got %v, %v; want type %v", recv.GetType(), name, r, err, wantType)
		}
		return r
	}
	if check(i(5), "toString", stackvalue.STRING).ToString() != "5" {
		t.Error("toString of 5")
	}
	if n, _ := check(s("12"), "toInteger", stackvalue.INTEGER).ToInteger(); n != 12 {
		t.Errorf("toInteger of \"12\" = %d", n)
	}
	if d, _ := check(s("2.5"), "toDouble", stackvalue.DOUBLE).ToDouble(); d != 2.5 {
		t.Errorf("toDouble of \"2.5\" = %v", d)
	}
	if b, _ := check(s("x"), "toBoolean", stackvalue.BOOL).ToBoolean(); !b {
		t.Error("toBoolean of a non-empty string")
	}
	if n, _ := check(s("abc"), "length", stackvalue.INTEGER).ToInteger(); n != 3 {
		t.Errorf("length of \"abc\" = %d", n)
	}
	if _, err := Invoke(s("abc"), "toInteger", nil); err == nil {
		t.Error("toInteger of \"abc\" must fail")
	}
	if _, err := Invoke(i(5), "length", nil); err == nil {
		t.Error("length of an integer must fail")
	}
}

func TestStringMethodsWiring(t *testing.T) {
	str := func(recv, name string, args ...stackvalue.StackValue) string {
		t.Helper()
		r, err := Invoke(s(recv), name, args)
		if err != nil || r.GetType() != stackvalue.STRING {
			t.Fatalf("%s.%s: got %v, %v", recv, name, r, err)
		}
		return r.ToString()
	}
	if got := str("  Hi  ", "trim"); got != "Hi" {
		t.Errorf("trim: %q", got)
	}
	if got := str("  Hi  ", "trimLeft"); got != "Hi  " {
		t.Errorf("trimLeft: %q", got)
	}
	if got := str("  Hi  ", "trimRight"); got != "  Hi" {
		t.Errorf("trimRight: %q", got)
	}
	// argument ORDER: replaceAllStringOccurrences(old, new)
	if got := str("foo", "replaceAllStringOccurrences", s("o"), s("0")); got != "f00" {
		t.Errorf("replaceAll(old, new) order: %q", got)
	}
	if got := str("ab123cd", "extractOneByREGEX", s("[0-9]+")); got == "" {
		t.Errorf("extract returned nothing: %q", got)
	}
	for name, want := range map[string]bool{"startsWith": true, "endsWith": false, "containsString": true} {
		arg := map[string]string{"startsWith": "He", "endsWith": "x", "containsString": "ell"}[name]
		r, err := Invoke(s("Hello"), name, []stackvalue.StackValue{s(arg)})
		if err != nil || r.GetType() != stackvalue.BOOL {
			t.Fatalf("%s: %v, %v", name, r, err)
		}
		if b, _ := r.ToBoolean(); b != want {
			t.Errorf("%s(%q) on Hello = %v, want %v", name, arg, b, want)
		}
	}
}

func TestTypeErrorsAreErrors(t *testing.T) {
	// wrong argument type (the corpus line  $s.startsWith(1)  is ERR on main)
	if _, err := Invoke(s("abc"), "startsWith", []stackvalue.StackValue{i(1)}); err == nil || !strings.Contains(err.Error(), "argument 1") {
		t.Errorf("startsWith(1): %v", err)
	}
	// wrong receiver type (the corpus line  $n.trim  is ERR on main)
	if _, err := Invoke(i(5), "trim", nil); err == nil || !strings.Contains(err.Error(), "has no method") {
		t.Errorf("5.trim: %v", err)
	}
	// wrong argument count
	if _, err := Invoke(s("abc"), "replaceAllStringOccurrences", []stackvalue.StackValue{s("a")}); err == nil {
		t.Error("replaceAll with one argument must fail")
	}
}

func TestXPathFailuresAreErrorsNotPanics(t *testing.T) {
	html := s(`<a href="x">one</a><a href="y">two</a>`)
	// two matches: the corpus line is ERR on main
	if _, err := Invoke(html, "findOneByXPATH", []stackvalue.StackValue{s("//a")}); err == nil {
		t.Error("two matches must be an error")
	}
	// no match: PANIC on main ("index out of range"), an error here, whatever the helper does
	for _, name := range []string{"findOneByXPATH", "findOneStringByXPATH", "findOneIntegerByXPATH", "findOneDoubleByXPATH", "findOneBooleanByXPATH"} {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%s with no match escaped as a panic: %v", name, r)
				}
			}()
			if _, err := Invoke(html, name, []stackvalue.StackValue{s("//zzz")}); err == nil {
				t.Errorf("%s with no match must be an error", name)
			}
		}()
	}
}

func TestFindByXPATHReturnsAList(t *testing.T) {
	r, err := Invoke(s(`<a href="x">one</a><a href="y">two</a>`), "findByXPATH", []stackvalue.StackValue{s("//a")})
	if err != nil || r.GetType() != stackvalue.JSON_ARRAY {
		t.Fatalf("got %v, %v", r, err)
	}
	if n, _ := r.Length(); n != 2 {
		t.Errorf("want 2 elements, got %d", n)
	}
	// nothing matched: if the helper reports an empty result (rather than an error), the list must
	// print as [], never as null
	r, err = Invoke(s("<p>x</p>"), "findByXPATH", []stackvalue.StackValue{s("//zzz")})
	if err == nil && r.ToString() != "[]" {
		t.Errorf("no match must print as [], got %q", r.ToString())
	}
}

func TestRandomDoubleInRangeWiring(t *testing.T) {
	r, err := InvokeFunc("randomDoubleInRange", []stackvalue.StackValue{i(1), i(1)})
	if err != nil || r.GetType() != stackvalue.DOUBLE {
		t.Fatalf("got %v, %v", r, err)
	}
	if _, err := InvokeFunc("randomDoubleInRange", []stackvalue.StackValue{s("a"), i(1)}); err == nil {
		t.Error("a string bound must be rejected")
	}
	if _, err := InvokeFunc("randomDoubleInRange", []stackvalue.StackValue{i(1)}); err == nil {
		t.Error("one argument must be rejected")
	}
}

func TestEveryRealMethodIsDocumented(t *testing.T) {
	for _, e := range Catalog() {
		if e.Doc == "" || e.Group == "" || e.Returns == "" {
			t.Errorf("%s.%s is missing Doc, Group or Returns (the VSIX and wizard are fed from this)", e.Receiver, e.Name)
		}
	}
}
