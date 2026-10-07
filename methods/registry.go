// Package methods is the single table of every method and free function the
// expression language can call. The grammar knows nothing about individual
// methods; it only produces "receiver . name ( args )" and "name ( args )".
package methods

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/gclkaze/evalang-globals/globals/stackvalue"
)

// Param describes one parameter. Optional parameters must come last.
type Param struct {
	Name     string
	Types    []stackvalue.StackValueType // accepted argument types; nil = any
	Optional bool
	Default  stackvalue.StackValue // used when an optional argument is omitted; may be nil
}

// Call is what a method body receives. Args always has len(Params) entries:
// omitted optional arguments are already replaced by their Default.
type Call struct {
	Recv  stackvalue.StackValue // nil for free functions
	Args  []stackvalue.StackValue
	given int
}

// Has reports whether the script actually supplied argument i. It lets a
// method tell "no default given" apart from "default is null".
func (c *Call) Has(i int) bool { return i < c.given }

type Method struct {
	Group   string // for docs and tooling: "string", "list", "object", ...
	Doc     string
	Returns string
	Params  []Param
	Fn      func(*Call) (stackvalue.StackValue, error)
}

type key struct {
	recv stackvalue.StackValueType
	name string
}

var (
	byRecv  = map[key]Method{}    // receiver-specific methods
	anyRecv = map[string]Method{} // methods valid on every receiver (type, isNull, ...)
	funcs   = map[string]Method{} // free functions: randomDoubleInRange(...), now(), ...
	aliases = map[string]string{} // old name -> canonical name
)

// reserved lists the words the SCRIPT lexer (Evalang.g4) turns into keyword
// tokens. Tafexpr would accept them as method names, but a script could never
// write them: "$x.in(1)" fails with "expecting PROP" before tafexpr sees it.
// Keep in sync with Evalang.g4; TestReservedWordsMatchGrammar cross-checks it.
var reserved = map[string]bool{
	"if": true, "else": true, "for": true, "in": true, "as": true, "import": true,
	"call": true, "next": true, "break": true, "continue": true, "return": true,
	"true": true, "false": true, "TRUE": true, "FALSE": true, "null": true,
	"Meta": true, "Component": true, "Licence": true,
	"ModuleDeclaration": true, "ModuleFunctionExport": true,
}

// validName mirrors the PROP token: a letter, then letters, digits or underscores.
var validName = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)

func checkName(name string) {
	if !validName.MatchString(name) {
		panic(fmt.Sprintf("methods: %q is not a valid method name (must match PROP: letter, then letters/digits/_)", name))
	}
	if reserved[name] {
		panic(fmt.Sprintf("methods: %q is a script keyword and could never be written in a script; pick another name", name))
	}
}

func validate(name string, m Method) {
	checkName(name)
	seenOptional := false
	for _, p := range m.Params {
		if p.Optional {
			seenOptional = true
		} else if seenOptional {
			panic(fmt.Sprintf("methods: %s: required parameter %q follows an optional one", name, p.Name))
		}
	}
	if m.Fn == nil {
		panic(fmt.Sprintf("methods: %s has no Fn", name))
	}
}

// Register adds a method for the given receiver types. With no receiver
// types the method is valid on every receiver. Registering the same
// receiver and name twice panics at startup, so clashes are found
// immediately instead of one silently winning.
func Register(name string, m Method, recv ...stackvalue.StackValueType) {
	validate(name, m)
	if len(recv) == 0 {
		if _, dup := anyRecv[name]; dup {
			panic(fmt.Sprintf("methods: %s registered twice for any receiver", name))
		}
		anyRecv[name] = m
		return
	}
	for _, t := range recv {
		k := key{t, name}
		if _, dup := byRecv[k]; dup {
			panic(fmt.Sprintf("methods: %s.%s registered twice", t, name))
		}
		byRecv[k] = m
	}
}

func RegisterFunc(name string, m Method) {
	validate(name, m)
	if _, dup := funcs[name]; dup {
		panic(fmt.Sprintf("methods: function %s registered twice", name))
	}
	funcs[name] = m
}

// Alias keeps an old name working (e.g. containsString -> contains).
func Alias(old, canonical string) {
	checkName(old)
	aliases[old] = canonical
}

// UnresolvedAliases lists aliases whose target is not registered; a test
// asserts it is empty once every file's init() has run.
func UnresolvedAliases() []string {
	var bad []string
	for old, target := range aliases {
		if len(methodsNamed(target)) == 0 {
			if _, ok := funcs[target]; !ok {
				bad = append(bad, old+" -> "+target)
			}
		}
	}
	sort.Strings(bad)
	return bad
}

func canonical(name string) string {
	if c, ok := aliases[name]; ok {
		return c
	}
	return name
}

// Invoke runs recv.name(args...).
func Invoke(recv stackvalue.StackValue, name string, args []stackvalue.StackValue) (stackvalue.StackValue, error) {
	name = canonical(name)
	t := recv.GetType()
	m, ok := byRecv[key{t, name}]
	if !ok {
		m, ok = anyRecv[name]
	}
	if !ok {
		return nil, fmt.Errorf("%s has no method %q; available: %s", t, name, strings.Join(availableFor(t), ", "))
	}
	return run(fmt.Sprintf("%s.%s", t, name), recv, m, args)
}

// InvokeFunc runs a free function: name(args...).
func InvokeFunc(name string, args []stackvalue.StackValue) (stackvalue.StackValue, error) {
	name = canonical(name)
	m, ok := funcs[name]
	if !ok {
		return nil, fmt.Errorf("unknown function %q", name)
	}
	return run(name, nil, m, args)
}

func run(label string, recv stackvalue.StackValue, m Method, args []stackvalue.StackValue) (stackvalue.StackValue, error) {
	min, max := 0, len(m.Params)
	for _, p := range m.Params {
		if !p.Optional {
			min++
		}
	}
	if len(args) < min || len(args) > max {
		want := fmt.Sprintf("%d", max)
		if min != max {
			want = fmt.Sprintf("%d to %d", min, max)
		}
		return nil, fmt.Errorf("%s expects %s argument(s), got %d", label, want, len(args))
	}
	for i, a := range args {
		p := m.Params[i]
		if len(p.Types) == 0 {
			continue
		}
		okType := false
		for _, want := range p.Types {
			if a.GetType() == want {
				okType = true
				break
			}
		}
		if !okType {
			return nil, fmt.Errorf("%s: argument %d (%s) must be %s, got %s", label, i+1, p.Name, typeList(p.Types), a.GetType())
		}
	}
	full := make([]stackvalue.StackValue, len(m.Params))
	copy(full, args)
	for i := len(args); i < len(m.Params); i++ {
		full[i] = m.Params[i].Default
	}
	return safeCall(label, m, &Call{Recv: recv, Args: full, given: len(args)})
}

// safeCall runs a method body and turns a panic inside it into an error. A method
// can never crash the evaluator, whatever the code it delegates to does (the
// no-match XPath case panicked with "index out of range" in the legacy handler).
func safeCall(label string, m Method, c *Call) (res stackvalue.StackValue, err error) {
	defer func() {
		if r := recover(); r != nil {
			res, err = nil, fmt.Errorf("%s: internal error: %v", label, r)
		}
	}()
	return m.Fn(c)
}

func typeList(ts []stackvalue.StackValueType) string {
	s := make([]string, len(ts))
	for i, t := range ts {
		s[i] = t.String()
	}
	return strings.Join(s, " or ")
}

func availableFor(t stackvalue.StackValueType) []string {
	set := map[string]bool{}
	for k := range byRecv {
		if k.recv == t {
			set[k.name] = true
		}
	}
	for n := range anyRecv {
		set[n] = true
	}
	out := make([]string, 0, len(set))
	for n := range set {
		out = append(out, n)
	}
	sort.Strings(out)
	return out
}

// Catalog lists every registered method, for tooling (VSIX completions,
// wizard method picker, docs).
type Entry struct {
	Receiver string
	Name     string
	Group    string
	Doc      string
	Returns  string
	Params   []Param
}

func Catalog() []Entry {
	var out []Entry
	for k, m := range byRecv {
		out = append(out, Entry{k.recv.String(), k.name, m.Group, m.Doc, m.Returns, m.Params})
	}
	for n, m := range anyRecv {
		out = append(out, Entry{"ANY", n, m.Group, m.Doc, m.Returns, m.Params})
	}
	for n, m := range funcs {
		out = append(out, Entry{"(function)", n, m.Group, m.Doc, m.Returns, m.Params})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Receiver != out[j].Receiver {
			return out[i].Receiver < out[j].Receiver
		}
		return out[i].Name < out[j].Name
	})
	return out
}

// ---- compile-time checks ---------------------------------------------------
//
// At compile time the receiver's type is not known (variables are dynamic),
// so these checks are deliberately loose: a call passes if SOME receiver type
// could accept it. They exist to catch typos and wrong argument counts before
// a long script starts running, not to replace the precise runtime checks.

func (m Method) bounds() (min, max int) {
	for _, p := range m.Params {
		if !p.Optional {
			min++
		}
	}
	return min, len(m.Params)
}

func methodsNamed(name string) []Method {
	name = canonical(name)
	var out []Method
	for k, m := range byRecv {
		if k.name == name {
			out = append(out, m)
		}
	}
	if m, ok := anyRecv[name]; ok {
		out = append(out, m)
	}
	return out
}

// CheckMethod validates recv.name(argc args) without knowing recv's type.
func CheckMethod(name string, argc int) error {
	ms := methodsNamed(name)
	if len(ms) == 0 {
		return unknownName("method", name, allMethodNames())
	}
	return checkArity("method", name, ms, argc)
}

// CheckFunc validates a free function call name(argc args).
func CheckFunc(name string, argc int) error {
	name = canonical(name)
	m, ok := funcs[name]
	if !ok {
		return unknownName("function", name, allFuncNames())
	}
	return checkArity("function", name, []Method{m}, argc)
}

func checkArity(kind, name string, ms []Method, argc int) error {
	accepted := map[int]bool{}
	for _, m := range ms {
		min, max := m.bounds()
		for n := min; n <= max; n++ {
			accepted[n] = true
		}
	}
	if accepted[argc] {
		return nil
	}
	counts := make([]int, 0, len(accepted))
	for n := range accepted {
		counts = append(counts, n)
	}
	sort.Ints(counts)
	parts := make([]string, len(counts))
	for i, n := range counts {
		parts[i] = fmt.Sprintf("%d", n)
	}
	return fmt.Errorf("%s %q takes %s argument(s), got %d", kind, name, orList(parts), argc)
}

func orList(parts []string) string {
	if len(parts) <= 1 {
		return strings.Join(parts, "")
	}
	return strings.Join(parts[:len(parts)-1], ", ") + " or " + parts[len(parts)-1]
}

func unknownName(kind, name string, candidates []string) error {
	if s := suggest(name, candidates); len(s) > 0 {
		return fmt.Errorf("unknown %s %q; did you mean: %s?", kind, name, strings.Join(s, ", "))
	}
	return fmt.Errorf("unknown %s %q", kind, name)
}

func allMethodNames() []string {
	set := map[string]bool{}
	for k := range byRecv {
		set[k.name] = true
	}
	for n := range anyRecv {
		set[n] = true
	}
	return sortedKeys(set)
}

func allFuncNames() []string {
	set := map[string]bool{}
	for n := range funcs {
		set[n] = true
	}
	return sortedKeys(set)
}

func sortedKeys(set map[string]bool) []string {
	out := make([]string, 0, len(set))
	for n := range set {
		out = append(out, n)
	}
	sort.Strings(out)
	return out
}

// suggest returns up to three candidates within edit distance 2, closest first.
func suggest(name string, candidates []string) []string {
	type scored struct {
		name string
		d    int
	}
	var hits []scored
	for _, c := range candidates {
		if d := editDistance(strings.ToLower(name), strings.ToLower(c)); d <= 2 {
			hits = append(hits, scored{c, d})
		}
	}
	sort.Slice(hits, func(i, j int) bool {
		if hits[i].d != hits[j].d {
			return hits[i].d < hits[j].d
		}
		return hits[i].name < hits[j].name
	})
	var out []string
	for i := 0; i < len(hits) && i < 3; i++ {
		out = append(out, hits[i].name)
	}
	return out
}

func editDistance(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur := make([]int, len(rb)+1)
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min3(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev = cur
	}
	return prev[len(rb)]
}

func min3(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}
