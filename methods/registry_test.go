package methods

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

// val satisfies the full stackvalue.StackValue interface by embedding it; only the two
// methods the registry actually calls are implemented.
type val struct {
	stackvalue.StackValue
	t stackvalue.StackValueType
	s string
}

func (v val) GetType() stackvalue.StackValueType { return v.t }
func (v val) ToString() string                   { return v.s }

func mk(t stackvalue.StackValueType, s string) stackvalue.StackValue { return val{t: t, s: s} }

func str(s string) stackvalue.StackValue { return mk(stackvalue.STRING, s) }
func num(s string) stackvalue.StackValue { return mk(stackvalue.INTEGER, s) }
func null() stackvalue.StackValue        { return mk(stackvalue.NULL, "") }

// registerFixtures adds the test-only methods to whatever tables are active.
func registerFixtures() {
	// trim([cutset])  -- optional parameter with a default
	Register("trim", Method{
		Group: "string", Returns: "STRING",
		Params: []Param{{Name: "cutset", Types: []stackvalue.StackValueType{stackvalue.STRING}, Optional: true, Default: str(" \t\r\n")}},
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			return str(strings.Trim(c.Recv.ToString(), c.Args[0].ToString())), nil
		},
	}, stackvalue.STRING)

	// get(key[, default]) on objects -- "was a default supplied?" matters
	Register("get", Method{
		Group: "object", Returns: "ANY",
		Params: []Param{
			{Name: "key", Types: []stackvalue.StackValueType{stackvalue.STRING}},
			{Name: "default", Optional: true},
		},
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			if c.Args[0].ToString() == "present" {
				return str("found"), nil
			}
			if c.Has(1) {
				return c.Args[1], nil // an explicit default, even an explicit null
			}
			return nil, errMissing(c.Args[0].ToString())
		},
	}, stackvalue.JSON_OBJECT)

	// get(index) on lists -- same name, different receiver type
	Register("get", Method{
		Group: "list", Returns: "ANY",
		Params: []Param{{Name: "index", Types: []stackvalue.StackValueType{stackvalue.INTEGER}}},
		Fn:     func(c *Call) (stackvalue.StackValue, error) { return str("item" + c.Args[0].ToString()), nil },
	}, stackvalue.JSON_ARRAY)

	// contains(x) on strings, with the old name kept as an alias
	Register("contains", Method{
		Group: "string", Returns: "BOOLEAN",
		Params: []Param{{Name: "needle", Types: []stackvalue.StackValueType{stackvalue.STRING}}},
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			return mk(stackvalue.BOOL, boolStr(strings.Contains(c.Recv.ToString(), c.Args[0].ToString()))), nil
		},
	}, stackvalue.STRING)
	Alias("containsString", "contains")

	// type() valid on every receiver
	Register("type", Method{
		Group: "introspection", Returns: "STRING",
		Fn: func(c *Call) (stackvalue.StackValue, error) { return str(c.Recv.GetType().String()), nil },
	})

	// a free function with two optional parameters
	RegisterFunc("randomInRange", Method{
		Group: "random", Returns: "DOUBLE",
		Params: []Param{{Name: "min", Optional: true, Default: num("0")}, {Name: "max", Optional: true, Default: num("1")}},
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			return str(c.Args[0].ToString() + ".." + c.Args[1].ToString()), nil
		},
	})
}

// withFixtures gives the calling test private tables holding only the fixtures, and restores the
// real registrations afterwards. Tests about the REAL registrations (aliases, reserved words,
// legacy names) deliberately do not call it.
func withFixtures(t *testing.T) {
	t.Helper()
	sb, sa, sf, sal := byRecv, anyRecv, funcs, aliases
	byRecv, anyRecv, funcs, aliases = map[key]Method{}, map[string]Method{}, map[string]Method{}, map[string]string{}
	t.Cleanup(func() { byRecv, anyRecv, funcs, aliases = sb, sa, sf, sal })
	registerFixtures()
}

type errMissing string

func (e errMissing) Error() string { return "no such key: " + string(e) }
func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func TestOptionalParamUsesDefault(t *testing.T) {
	withFixtures(t)
	r, err := Invoke(str("  hi  "), "trim", nil)
	if err != nil || r.ToString() != "hi" {
		t.Fatalf("got %v, %v", r, err)
	}
	r, err = Invoke(str("xxhixx"), "trim", []stackvalue.StackValue{str("x")})
	if err != nil || r.ToString() != "hi" {
		t.Fatalf("got %v, %v", r, err)
	}
}

func TestHasDistinguishesOmittedFromNull(t *testing.T) {
	withFixtures(t)
	obj := mk(stackvalue.JSON_OBJECT, "")
	if _, err := Invoke(obj, "get", []stackvalue.StackValue{str("nope")}); err == nil {
		t.Fatal("omitted default should fail on a missing key")
	}
	r, err := Invoke(obj, "get", []stackvalue.StackValue{str("nope"), null()})
	if err != nil || r.GetType() != stackvalue.NULL {
		t.Fatalf("explicit null default must be returned as null, got %v, %v", r, err)
	}
	r, _ = Invoke(obj, "get", []stackvalue.StackValue{str("nope"), str("fallback")})
	if r.ToString() != "fallback" {
		t.Fatalf("got %v", r)
	}
}

func TestArityMessages(t *testing.T) {
	withFixtures(t)
	cases := []struct {
		recv stackvalue.StackValue
		name string
		args []stackvalue.StackValue
		want string
	}{
		{str("a"), "trim", []stackvalue.StackValue{str("x"), str("y")}, "String.trim expects 0 to 1 argument(s), got 2"},
		{str("a"), "contains", nil, "String.contains expects 1 argument(s), got 0"},
		{mk(stackvalue.JSON_OBJECT, ""), "get", nil, "JSONObject.get expects 1 to 2 argument(s), got 0"},
	}
	for _, c := range cases {
		_, err := Invoke(c.recv, c.name, c.args)
		if err == nil || err.Error() != c.want {
			t.Fatalf("want %q, got %v", c.want, err)
		}
	}
}

func TestArgumentTypeCheck(t *testing.T) {
	withFixtures(t)
	_, err := Invoke(str("abc"), "contains", []stackvalue.StackValue{num("1")})
	want := "String.contains: argument 1 (needle) must be String, got Integer"
	if err == nil || err.Error() != want {
		t.Fatalf("want %q, got %v", want, err)
	}
}

func TestSameNameDispatchesOnReceiverType(t *testing.T) {
	withFixtures(t)
	a, _ := Invoke(mk(stackvalue.JSON_ARRAY, ""), "get", []stackvalue.StackValue{num("3")})
	o, _ := Invoke(mk(stackvalue.JSON_OBJECT, ""), "get", []stackvalue.StackValue{str("present")})
	if a.ToString() != "item3" || o.ToString() != "found" {
		t.Fatalf("got %v / %v", a, o)
	}
}

func TestAliasAndAnyReceiver(t *testing.T) {
	withFixtures(t)
	r, err := Invoke(str("hello"), "containsString", []stackvalue.StackValue{str("ell")})
	if err != nil || r.ToString() != "true" {
		t.Fatalf("alias failed: %v %v", r, err)
	}
	for _, v := range []stackvalue.StackValue{str("x"), num("1"), null()} {
		r, err := Invoke(v, "type", nil)
		if err != nil || r.ToString() != v.GetType().String() {
			t.Fatalf("type() failed on %v: %v %v", v.GetType(), r, err)
		}
	}
}

func TestUnknownMethodListsWhatExists(t *testing.T) {
	withFixtures(t)
	_, err := Invoke(str("a"), "trimm", nil)
	if err == nil || !strings.Contains(err.Error(), `String has no method "trimm"`) || !strings.Contains(err.Error(), "trim") {
		t.Fatalf("got %v", err)
	}
	// methods of other receiver types must not be offered
	if strings.Contains(err.Error(), "keys") {
		t.Fatal("offered a method that is not valid on STRING")
	}
}

func TestFreeFunctionDefaults(t *testing.T) {
	withFixtures(t)
	r, _ := InvokeFunc("randomInRange", nil)
	if r.ToString() != "0..1" {
		t.Fatalf("got %v", r)
	}
	r, _ = InvokeFunc("randomInRange", []stackvalue.StackValue{num("5"), num("9")})
	if r.ToString() != "5..9" {
		t.Fatalf("got %v", r)
	}
}

func TestRegistrationGuards(t *testing.T) {
	withFixtures(t)
	mustPanic := func(name string, f func()) {
		defer func() {
			if recover() == nil {
				t.Fatalf("%s: expected a panic", name)
			}
		}()
		f()
	}
	mustPanic("duplicate", func() {
		Register("trim", Method{Fn: func(*Call) (stackvalue.StackValue, error) { return nil, nil }}, stackvalue.STRING)
	})
	mustPanic("required after optional", func() {
		Register("bad", Method{
			Params: []Param{{Name: "a", Optional: true}, {Name: "b"}},
			Fn:     func(*Call) (stackvalue.StackValue, error) { return nil, nil },
		}, stackvalue.STRING)
	})
}

func TestCatalogForTooling(t *testing.T) {
	withFixtures(t)
	var found bool
	for _, e := range Catalog() {
		if e.Receiver == "String" && e.Name == "trim" && e.Group == "string" && len(e.Params) == 1 && e.Params[0].Optional {
			found = true
		}
	}
	if !found {
		t.Fatal("catalog is missing trim")
	}
}

func TestCompileTimeCheckMethod(t *testing.T) {
	withFixtures(t)
	cases := []struct {
		name string
		argc int
		want string // "" = valid, otherwise the exact error text
	}{
		{"trim", 0, ""},
		{"trim", 1, ""},
		{"trim", 2, `method "trim" takes 0 or 1 argument(s), got 2`},
		{"get", 1, ""}, // list.get(index) and object.get(key)
		{"get", 2, ""}, // object.get(key, default)
		{"get", 3, `method "get" takes 1 or 2 argument(s), got 3`},  // no receiver accepts 3
		{"containsString", 1, ""},                                   // alias resolves
		{"type", 0, ""},                                             // any-receiver method
		{"trimm", 0, `unknown method "trimm"; did you mean: trim?`}, // typo suggestion
	}
	for _, c := range cases {
		err := CheckMethod(c.name, c.argc)
		switch {
		case c.want == "" && err != nil:
			t.Fatalf("%s/%d: unexpected error %v", c.name, c.argc, err)
		case c.want != "" && (err == nil || err.Error() != c.want):
			t.Fatalf("%s/%d: want %q, got %v", c.name, c.argc, c.want, err)
		}
	}
}

func TestMethodNamesAreCaseSensitive(t *testing.T) {
	withFixtures(t)
	err := CheckMethod("TRIM", 0)
	if err == nil || !strings.Contains(err.Error(), "did you mean: trim") {
		t.Fatalf("TRIM is unknown (names are case-sensitive) but should suggest trim, got %v", err)
	}
}

func TestCompileTimeCheckFunc(t *testing.T) {
	withFixtures(t)
	if err := CheckFunc("randomInRange", 2); err != nil {
		t.Fatal(err)
	}
	if err := CheckFunc("randomInRange", 0); err != nil { // both params optional
		t.Fatal(err)
	}
	if err := CheckFunc("randomInRange", 3); err == nil || err.Error() != `function "randomInRange" takes 0, 1 or 2 argument(s), got 3` {
		t.Fatalf("got %v", err)
	}
	if err := CheckFunc("randomInRang", 1); err == nil || !strings.Contains(err.Error(), "did you mean: randomInRange") {
		t.Fatalf("got %v", err)
	}
	if err := CheckFunc("zzzzzz", 0); err == nil || strings.Contains(err.Error(), "did you mean") {
		t.Fatalf("a far-off name must not get a suggestion, got %v", err)
	}
}

func TestRegistrationRejectsKeywordsAndBadNames(t *testing.T) {
	noop := func(*Call) (stackvalue.StackValue, error) { return nil, nil }
	for _, name := range []string{"in", "as", "next", "call", "for", "if", "return", "null", "1abc", "_x", "a-b", "", "has space"} {
		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("Register(%q) should panic", name)
				}
			}()
			Register(name, Method{Fn: noop}, stackvalue.STRING)
		}()
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Error("Alias with a keyword name should panic")
			}
		}()
		Alias("in", "contains")
	}()
}

func TestAliasesResolve(t *testing.T) {
	if bad := UnresolvedAliases(); len(bad) != 0 {
		t.Fatalf("aliases without a target: %v", bad)
	}
}

// Set EVALANG_G4 to the real Evalang.g4 (CI does) to prove the reserved list
// still covers every keyword the script lexer defines, and that no registered
// name collides with one.
func TestReservedWordsMatchGrammar(t *testing.T) {
	path := os.Getenv("EVALANG_G4")
	if path == "" {
		t.Skip("EVALANG_G4 not set")
	}
	src, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	words := map[string]bool{}
	for _, m := range regexp.MustCompile(`'([A-Za-z][A-Za-z0-9_]*)'`).FindAllStringSubmatch(string(src), -1) {
		words[m[1]] = true
	}
	for w := range words {
		for k := range byRecv {
			if k.name == w {
				t.Errorf("method %s.%s collides with grammar keyword %q", k.recv, k.name, w)
			}
		}
		for n := range anyRecv {
			if n == w {
				t.Errorf("method %s collides with grammar keyword %q", n, w)
			}
		}
		for n := range funcs {
			if n == w {
				t.Errorf("function %s collides with grammar keyword %q", n, w)
			}
		}
		if !reserved[w] && w != "Main" && !strings.HasPrefix(w, "On") && w != "randomDoubleInRange" {
			t.Logf("grammar keyword %q is not in the reserved list (add it if a method could be given that name)", w)
		}
	}
	for r := range reserved {
		if !words[r] {
			t.Logf("reserved word %q no longer appears in the grammar", r)
		}
	}
}

func TestMethodPanicBecomesError(t *testing.T) {
	withFixtures(t)
	Register("boom", Method{
		Group: "test",
		Fn: func(c *Call) (stackvalue.StackValue, error) {
			var empty []string
			return str(empty[0]), nil // index out of range, like the legacy no-match XPath
		},
	}, stackvalue.STRING)
	RegisterFunc("boomFn", Method{Group: "test", Fn: func(*Call) (stackvalue.StackValue, error) { panic("kaboom") }})

	r, err := Invoke(str("x"), "boom", nil)
	if err == nil || r != nil || !strings.Contains(err.Error(), "String.boom: internal error") || !strings.Contains(err.Error(), "index out of range") {
		t.Fatalf("a panicking method must return a descriptive error, got %v, %v", r, err)
	}
	if _, err := InvokeFunc("boomFn", nil); err == nil || !strings.Contains(err.Error(), "boomFn: internal error: kaboom") {
		t.Fatalf("got %v", err)
	}
	// and a well-behaved method is unaffected
	if r, err := Invoke(str("  hi  "), "trim", nil); err != nil || r.ToString() != "hi" {
		t.Fatalf("got %v, %v", r, err)
	}
}

// The names below are what type() returns and what error messages print: they are the
// language's vocabulary, so a change to StackValueType.String() must be deliberate.
func TestReceiverNamesInMessages(t *testing.T) {
	for typ, want := range map[stackvalue.StackValueType]string{
		stackvalue.STRING: "String", stackvalue.INTEGER: "Integer", stackvalue.DOUBLE: "Double",
		stackvalue.BOOL: "Bool", stackvalue.NULL: "Null", stackvalue.JSON_OBJECT: "JSONObject",
		stackvalue.JSON_ARRAY: "JSONArray", stackvalue.USER_DEFINED: "UserDefined",
	} {
		_, err := Invoke(mk(typ, ""), "noSuchMethod", nil)
		if err == nil || !strings.HasPrefix(err.Error(), want+` has no method "noSuchMethod"`) {
			t.Errorf("%v: want a message starting with %q, got %v", typ, want, err)
		}
	}
}
