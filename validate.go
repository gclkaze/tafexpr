package tafexpr

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gclkaze/tafexpr/methods"
	"github.com/gclkaze/tafexpr/parser"
)

// Validate checks an expression WITHOUT evaluating it: that it parses, that every method and
// function it calls exists, and that each call gets an acceptable number of arguments. It needs
// no variable values, so it can run when a script is compiled. It returns one error per problem,
// or nil when the expression is fine.
//
// What it cannot know: the type a variable will hold (so $n.trim() on an integer is only found at
// run time), and what a property read like $x.length refers to (it is valid syntax, so it is only
// found when the line runs).
func Validate(expression string) []error {
	syntax := &syntaxCollector{DefaultErrorListener: antlr.NewDefaultErrorListener()}

	lexer := parser.NewTafexprLexer(antlr.NewInputStream(expression))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(syntax)

	p := parser.NewTafexprParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	p.RemoveErrorListeners()
	p.AddErrorListener(syntax)

	tree := p.Taf_expression()
	if len(syntax.errs) > 0 {
		return syntax.errs
	}

	checker := &callChecker{BaseTafexprListener: &parser.BaseTafexprListener{}}
	antlr.ParseTreeWalkerDefault.Walk(checker, tree)
	return checker.errs
}

// syntaxCollector gathers lexer and parser errors instead of printing them.
type syntaxCollector struct {
	*antlr.DefaultErrorListener
	errs []error
}

func (c *syntaxCollector) SyntaxError(_ antlr.Recognizer, _ interface{}, _, column int, msg string, _ antlr.RecognitionException) {
	c.errs = append(c.errs, fmt.Errorf("syntax error at position %d: %s", column+1, msg))
}

// callChecker walks a parsed expression and checks every call against the method registry.
type callChecker struct {
	*parser.BaseTafexprListener
	errs []error
}

// ExitHandleMethodCall checks  receiver.name(args...) . The first expression is the receiver.
func (v *callChecker) ExitHandleMethodCall(c *parser.HandleMethodCallContext) {
	if err := methods.CheckMethod(c.PROP().GetText(), len(c.AllExpression())-1); err != nil {
		v.errs = append(v.errs, err)
	}
}

// ExitHandleFunctionCall checks  name(args...) , a call with no receiver.
func (v *callChecker) ExitHandleFunctionCall(c *parser.HandleFunctionCallContext) {
	if err := methods.CheckFunc(c.PROP().GetText(), len(c.AllExpression())); err != nil {
		v.errs = append(v.errs, err)
	}
}
