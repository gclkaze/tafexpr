package tafexpr

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gclkaze/tafexpr/parser"
	"github.com/gclkaze/tafexpr/stackvalue"
	"github.com/gclkaze/tafexpr/tafargumentlistenererrortypes"
	"github.com/gclkaze/tafexpr/variablecontext"

	"github.com/gclkaze/evalang-globals/globals"

	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

type TAFArgumentParser struct {
	tokens          []antlr.Token
	symbolicNames   []string
	IntValue        int
	DoubleValue     float64
	BoolValue       bool
	JSONValue       globals.JSONObjectGen
	JSONArray       globals.JSONArrayGen
	RefValue        *globals.ParameterValue
	VariableStorage string

	StringValue string

	OnError         bool
	VariableContext variablecontext.IVariableContext

	lexerErrors                *TAFArgumentErrorListener
	parserErrors               *TAFArgumentErrorListener
	ErrorMsgs                  []TAFParserArgumentError
	Result                     stackvalue.StackValue
	VariableStorageDestination VarExpression
	VariableExpressions        []VarExpression

	Debug                bool
	IsFloat              bool
	interpolationApplied bool

	interpol *regexp.Regexp

	secretAware bool
}

func (ap *TAFArgumentParser) SetSecretAwareness(secretAware bool) {
	ap.secretAware = secretAware
}

func (ap *TAFArgumentParser) clearErrors() {
	if ap.lexerErrors != nil {
		ap.lexerErrors = nil
	}

	if ap.parserErrors != nil {
		ap.parserErrors = nil
	}
}

func (ap TAFArgumentParser) GetErrors() []string {
	var errors []string
	if ap.parserErrors != nil && len(ap.parserErrors.Errors) != 0 {
		for i := 0; i < len(ap.parserErrors.Errors); i++ {
			cur := ap.parserErrors.Errors[i]
			er := fmt.Sprintf("%v", cur.Error())
			errors = append(errors, er)
		}
	}

	if ap.parserErrors != nil && len(ap.lexerErrors.Errors) != 0 {
		for i := 0; i < len(ap.lexerErrors.Errors); i++ {
			cur := ap.lexerErrors.Errors[i]
			er := fmt.Sprintf("%v", cur.Error())
			errors = append(errors, er)
		}
	}

	if ap.ErrorMsgs != nil && len(ap.ErrorMsgs) != 0 {
		for i := 0; i < len(ap.ErrorMsgs); i++ {
			cur := ap.ErrorMsgs[i]
			er := fmt.Sprintf("%v", cur.Msg.Error())
			errors = append(errors, er)
		}
	}

	return errors
}

func (ap *TAFArgumentParser) CanParse(arg string) bool {

	ap.clearErrors()

	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.OnError = true

	ap.VariableContext = &variablecontext.VariableContext{}

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	if l.IsFloat {
		ap.DoubleValue = l.popFloat()

	} else {
		last := l.popStack()
		i, ok := last.(*stackvalue.IntegerStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		ap.IntValue = i.GetValue()
	}
	if l.OnError {
		ap.OnError = true
		return false
	}

	ap.OnError = false
	return true
}
func (ap *TAFArgumentParser) ParseVariableStorageDestination(arg string) bool {
	ap.clearErrors()

	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.JSONArray = nil
	ap.JSONValue = nil
	ap.OnError = true

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug, onVariableStorageDeclarationMode: true}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	if l.OnError {
		ap.OnError = true
		ap.ErrorMsgs = l.ErrorMsgs
		return false
	}

	//	_ = l.popStack()
	ap.VariableStorageDestination = l.GetLastVariableDestinationExpresison()
	return true
}

func (ap *TAFArgumentParser) Analyze(arg string) bool {
	ap.clearErrors()

	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.JSONArray = nil
	ap.JSONValue = nil
	ap.OnError = true

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug, onVariableStorageDeclarationMode: false}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	ap.VariableExpressions = l.GetReferencedVariables()
	return true
}

func (ap *TAFArgumentParser) ParsePureVariableStorageDestination(arg string) bool {
	ap.clearErrors()

	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.JSONArray = nil
	ap.JSONValue = nil
	ap.OnError = true

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug, onVariableStorageDeclarationMode: true}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	errs := ap.GetErrors()
	if len(errs) > 0 {
		ap.OnError = true
		return false
	}

	if l.LastExit != "VarExpression" {
		ap.OnError = true
		return false
	}

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	if l.OnError {
		ap.OnError = true
		ap.ErrorMsgs = l.ErrorMsgs
		return false
	}

	//	_ = l.popStack()
	ap.VariableStorageDestination = l.GetLastVariableDestinationExpresison()
	return true
}

func (ap *TAFArgumentParser) Interpolate(arg string) bool {
	r := regexp.MustCompile(`\{\{(.*?)\}\}`)
	submatches := r.FindAllString(arg, -1)
	indices := r.FindAllStringIndex(arg, -1)
	inter := []InterpolationToken{}
	for i := 0; i < len(submatches); i++ {
		/*fmt.Printf("submatches[i]: %v\n", submatches[i])
		fmt.Printf("indices: %v\n", indices[i])*/
		trimmed := strings.Trim(submatches[i], "{")
		trimmed = strings.Trim(trimmed, "}")

		parsed := ap.Parse(trimmed)
		if parsed {
			//sValue := ap.Result.ToString()
			//fmt.Printf("sValue: %v\n", sValue)
			inter = append(inter, *NewInterpolationToken(indices[i][0], indices[i][1], submatches[i], ap.Result))
		}
	}
	if len(inter) == 0 {
		ap.interpolationApplied = false
		ap.StringValue = arg
		return true
	}
	current := 0
	substring := ""
	tok := inter[0]
	for i := 0; i < len(arg); i++ {
		if i == tok.GetStart() {
			substring += tok.GetValue().ToString()
			i = tok.GetEnd() - 1
			current++
			if current == len(inter) {
				if i+1 >= len(arg) {
					break
				}
				lastPart := arg[i+1:]
				substring += lastPart
				substring = strings.TrimSuffix(substring, "<EOF>")
				break
			}
			tok = inter[current]
			continue
		}
		substring += string(arg[i])
	}
	//l.interpolationApplied = true
	//fmt.Printf("substring: %v\n", substring)
	ap.StringValue = substring
	ap.interpolationApplied = arg != substring
	return true
}

func (ap *TAFArgumentParser) Parse(arg string) bool {

	ap.clearErrors()

	toBeEscaped := false
	/*	if strings.Contains(arg, "\"") && !strings.Contains(arg, "\\\"") {
		toBeEscaped = true
		fmt.Println(arg)
		arg = strings.ReplaceAll(arg, "\"", "\\\"")
		fmt.Println(arg)
	}*/
	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.JSONArray = nil
	ap.JSONValue = nil
	ap.OnError = true
	ap.interpolationApplied = false
	ap.StringValue = ""
	ap.BoolValue = false

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug, Escaped: toBeEscaped}
	l.SetSecretAware(ap.secretAware)

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}
	if l.OnError {
		ap.OnError = true
		ap.ErrorMsgs = l.ErrorMsgs
		return false
	}

	if l.StackIsEmpty() {
		ap.OnError = true
		return false
	}
	last := l.popStack()
	ap.Result = last
	switch last.GetType() {
	case stackvalue.REFERENCE:
		{
			i, ok := last.(*stackvalue.ReferenceStackValue)
			if !ok {
				panic("cannot cast to reference in the parser")
			}
			ap.RefValue = i.GetValue()
		}
	case stackvalue.JSON_OBJECT:
		i, ok := last.(*stackvalue.JSONStackValue)
		if !ok {
			panic("cannot cast to json in the parser")
		}
		ap.JSONValue = i.GetValue()
	case stackvalue.JSON_ARRAY:
		i, ok := last.(*stackvalue.JSONArrayStackValue)
		if !ok {
			panic("cannot cast to json array in the parser")
		}
		ap.JSONArray = i.GetValue()
	case stackvalue.INTEGER:
		i, ok := last.(*stackvalue.IntegerStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		ap.IntValue = i.GetValue()
	case stackvalue.BOOL:
		i, ok := last.(*stackvalue.BoolStackValue)
		if !ok {
			panic("cannot cast to bool in the parser")
		}
		ap.BoolValue = i.GetValue()
	case stackvalue.DOUBLE:
		i, ok := last.(*stackvalue.DoubleStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		ap.DoubleValue = i.GetValue()
		if isIntegral(ap.DoubleValue) {
			ap.IntValue = int(ap.DoubleValue)
		}
	case stackvalue.NULL:
		_, ok := last.(*stackvalue.NullStackValue)
		if !ok {
			panic("cannot cast to null in the parser")
		}
		ap.IntValue = 0
	case stackvalue.STRING:
		i, ok := last.(*stackvalue.StringStackValue)
		if !ok {
			panic("cannot cast to string in the parser")
		}
		ap.StringValue = i.GetValue()
		if toBeEscaped {
			fmt.Println(ap.StringValue)
			ap.StringValue = strings.ReplaceAll(ap.StringValue, "\\\"", "\"")
			fmt.Println(ap.StringValue)
			ap.StringValue = "\"" + ap.StringValue + "\""
		}
	}

	if l.OnError {
		ap.OnError = true
		return false
	}

	ap.OnError = false
	return true
}
func (ap *TAFArgumentParser) ParseExpressionAndCheckAgainstContract(arg string, contract string) bool {

	if contract == "variable_storage" {
		res := ap.ParsePureVariableStorageDestination(arg)
		return res
	}
	ap.clearErrors()

	toBeEscaped := false
	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.JSONArray = nil
	ap.JSONValue = nil
	ap.OnError = true
	ap.interpolationApplied = false
	ap.StringValue = ""
	ap.BoolValue = false

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug, Escaped: toBeEscaped}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	/*	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}*/

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}
	if l.OnError {
		ap.OnError = true
		ap.ErrorMsgs = l.ErrorMsgs
		return false
	}

	arg = strings.ReplaceAll(arg, `"`, ``)

	if l.StackIsEmpty() {
		ap.OnError = true
		ap.ErrorMsgs = append(ap.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("%s is not a %s Expression", arg, contract), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return false
	}
	last := l.popStack()
	ap.Result = last
	if l.LastExit == "VarExpression" && contract == "variable_storage" {
		ap.OnError = false
		return true
	}
	switch last.GetType() {
	case stackvalue.INTEGER:
		_, ok := last.(*stackvalue.IntegerStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		if contract == "double" || contract == "number" || contract == "integer" || contract == "variable" {
			return true
		}
		ap.ErrorMsgs = append(ap.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("%s is not a %s Expression", arg, contract), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
	case stackvalue.STRING:
		_, ok := last.(*stackvalue.StringStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		if contract == "string" || contract == "variable" {
			return true
		}
		ap.ErrorMsgs = append(ap.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("%s is not a %s Expression", arg, contract), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
	case stackvalue.DOUBLE:
		_, ok := last.(*stackvalue.DoubleStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		if contract == "double" || contract == "number" || contract == "integer" || contract == "variable" {
			return true
		}
		ap.ErrorMsgs = append(ap.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("%s is not a %s Expression", arg, contract), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
	case stackvalue.BOOL:
		_, ok := last.(*stackvalue.BoolStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		if contract == "bool" || contract == "variable" {
			return true
		}
		ap.ErrorMsgs = append(ap.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("%s is not a %s Expression", arg, contract), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
	default:
		ap.OnError = true
		ap.ErrorMsgs = append(ap.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("%s is not a %s Expression", arg, contract), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return false
	}

	ap.OnError = true
	return false
}

func (ap *TAFArgumentParser) ParsePotentialIntegerExpression(arg string) bool {

	ap.clearErrors()

	toBeEscaped := false
	ap.IsFloat = false
	ap.DoubleValue = -1
	ap.IntValue = -1
	ap.JSONArray = nil
	ap.JSONValue = nil
	ap.OnError = true
	ap.interpolationApplied = false
	ap.StringValue = ""
	ap.BoolValue = false

	ap.lexerErrors = &TAFArgumentErrorListener{}
	a := antlr.NewInputStream(arg)

	//log.Println(arg)
	// Create the Lexer
	lexer := parser.NewTafexprLexer(a)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ap.lexerErrors)

	ap.symbolicNames = lexer.SymbolicNames
	if len(ap.lexerErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewTafexprParser(stream)
	l := &TAFArgumentListener{VariableContext: ap.VariableContext, Debug: ap.Debug, Escaped: toBeEscaped}

	ap.parserErrors = &TAFArgumentErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(ap.parserErrors)

	antlr.ParseTreeWalkerDefault.Walk(l, p.Taf_expression())

	/*	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}*/

	if len(ap.parserErrors.Errors) != 0 {
		ap.OnError = true
		return false
	}
	if l.OnError {
		ap.OnError = true
		ap.ErrorMsgs = l.ErrorMsgs
		return false
	}

	if l.StackIsEmpty() {
		ap.OnError = true
		return false
	}
	last := l.popStack()
	ap.Result = last
	switch last.GetType() {
	case stackvalue.REFERENCE:
		{
			i, ok := last.(*stackvalue.ReferenceStackValue)
			if !ok {
				panic("cannot cast to reference in the parser")
			}
			ap.RefValue = i.GetValue()
		}
	case stackvalue.INTEGER:
		i, ok := last.(*stackvalue.IntegerStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		ap.IntValue = i.GetValue()
	case stackvalue.DOUBLE:
		i, ok := last.(*stackvalue.DoubleStackValue)
		if !ok {
			panic("cannot cast to int in the parser")
		}
		ap.DoubleValue = i.GetValue()
		if isIntegral(ap.DoubleValue) {
			ap.IntValue = int(ap.DoubleValue)
		}
	default:
		ap.OnError = true
		ap.ErrorMsgs = append(ap.ErrorMsgs, TAFParserArgumentError{Msg: fmt.Errorf("%s is not an Integer Expression", arg), Type: tafargumentlistenererrortypes.RUNTIME_ERROR})
		return false
	}

	if l.OnError {
		ap.OnError = true
		return false
	}

	ap.OnError = false
	return true
}

func (ap TAFArgumentParser) InterpolationApplied() bool {
	return ap.interpolationApplied
}

func (ap *TAFArgumentParser) PrintTokens() {
	for i := 0; i < len(ap.tokens); i++ {
		t := ap.tokens[i]
		log.Printf("%s (%q)\n", ap.symbolicNames[t.GetTokenType()], t.GetText())
	}
}

func (ap *TAFArgumentParser) GetTokensLength() int {
	if ap == nil {
		return 0
	}
	return len(ap.tokens)
}

func (ap *TAFArgumentParser) TokenExists(t string, txt string) bool {
	if ap == nil {
		return false
	}

	for i := 0; i < len(ap.tokens); i++ {
		tt := ap.tokens[i]
		if ap.symbolicNames[tt.GetTokenType()] == t && txt == tt.GetText() {
			return true
		}
	}
	return false
}
