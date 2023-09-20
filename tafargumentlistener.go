package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"tafexpr/parser"
	"tafexpr/variablecontext"
)

type TAFArgumentListenerError int

const (
	DIVISION_BY_ZERO TAFArgumentListenerError = iota
)

type TAFArgumentListener struct {
	*parser.BaseTafexprListener
	stack           []float64
	IsFloat         bool
	OnError         bool
	Error           TAFArgumentListenerError
	VariableContext variablecontext.IVariableContext
	CurrentPath     string
	Scope           int
	Index           IndexStack
}

func (l *TAFArgumentListener) push(i float64) {
	l.stack = append(l.stack, i)
	//log.Println("pushed ", l.stack)

	l.IsFloat = !isIntegral(i)
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

func (l *TAFArgumentListener) pop() any {

	if l.OnError {
		return -1
	}
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}
	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]
	//log.Println("popped ", l.stack)
	if l.IsFloat {
		return result
	}
	return int(result)
}

func (l *TAFArgumentListener) popFloat() float64 {
	if l.OnError {
		return -1
	}

	if !l.IsFloat {
		panic("popFloat can be called only if the result is a FLOAT64 number.")
	}
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}
	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	const prec = 200
	a := new(big.Float).SetPrec(prec).SetFloat64(result)
	//fmt.Printf("a.String(): %v\n", a.String())
	result, _ = strconv.ParseFloat(a.String(), 64)
	return result
}

func (l *TAFArgumentListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	if l.IsFloat {
		rightFloat := right.(float64)
		if rightFloat == 0 {
			l.OnError = true
			l.Error = DIVISION_BY_ZERO
			return
		}
		leftFloat := left.(float64)

		switch c.GetOp().GetTokenType() {
		case parser.TafexprParserMUL:
			v := leftFloat * rightFloat
			l.push(v)
		case parser.TafexprParserDIV:
			v := leftFloat / rightFloat
			l.push(v)
		default:
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}

	} else {
		rightInt := right.(int)
		if rightInt == 0 {
			l.OnError = true
			l.Error = DIVISION_BY_ZERO
			return
		}

		leftInt := left.(int)

		switch c.GetOp().GetTokenType() {
		case parser.TafexprParserMUL:
			v := leftInt * rightInt
			l.push(float64(v))
			break
		case parser.TafexprParserDIV:
			v := leftInt / rightInt
			l.push(float64(v))
			break
		default:
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}
	}

}

func (l *TAFArgumentListener) ExitAddSub(c *parser.AddSubContext) {
	right := l.pop()
	left := l.pop()

	if l.IsFloat {
		rightFloat := right.(float64)
		leftFloat := left.(float64)
		switch c.GetOp().GetTokenType() {
		case parser.TafexprParserADD:
			v := leftFloat + rightFloat
			/*			fmt.Printf("leftFloat: %v\n", leftFloat)
						fmt.Printf("rightFloat: %v\n", rightFloat)
						fmt.Printf("v: %v\n", v)*/
			l.push(v)
			break
		case parser.TafexprParserSUB:
			v := leftFloat - rightFloat
			l.push(v)
			break
		default:
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
			break
		}

	} else {
		rightInt := right.(int)
		leftInt := left.(int)

		switch c.GetOp().GetTokenType() {
		case parser.TafexprParserADD:
			v := leftInt + rightInt
			l.push(float64(v))
			break
		case parser.TafexprParserSUB:
			v := leftInt - rightInt
			l.push(float64(v))
			break
		default:
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}
	}
}

func (l *TAFArgumentListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(float64(i))
}

func (l *TAFArgumentListener) ExitDoubleValue(c *parser.DoubleValueContext) {
	//fmt.Printf("c.GetText(): %v\n", c.GetText())
	i, err := strconv.ParseFloat(c.GetText(), -1)
	if err != nil {
		panic(err.Error())
	}

	l.IsFloat = true
	//log.Println(i)

	l.push(i)
}

func (l *TAFArgumentListener) EnterTaf_expression(c *parser.Taf_expressionContext) {

}

func (l *TAFArgumentListener) ExitTaf_expression(c *parser.Taf_expressionContext) {
	log.Println("ExitTaf_expression" + c.GetText())
	l.Index.Clear()
}

func (l *TAFArgumentListener) ExitNumberValue(c *parser.NumberContext) {
	log.Println("ExitNumberValue " + c.GetText())
}

func (l *TAFArgumentListener) EnterVar_expression(c *parser.Var_expressionContext) {
	varExpression := c.GetText()
	log.Println("EnterVar_expression " + varExpression)

	l.Index.SetExpression(varExpression)
}

func (l *TAFArgumentListener) EnterHandleVarExpression(c *parser.HandleVarExpressionContext) {
	log.Println("EnterHandleVarExpression")
	if l.Scope == 0 {
		l.CurrentPath = c.GetText()
	}
	l.Scope++
}

func (l *TAFArgumentListener) ExitHandleVarExpression(c *parser.HandleVarExpressionContext) {
	log.Println("ExitHandleVarExpression")
	l.Scope--
	if l.Scope == 0 && !l.Index.IsEmpty() {
		fmt.Printf("l.Index: %v\n", l.Index)
		//we need to store index value to the stack
		log.Println("==============================================OUT OF SCOPE for ==>" + c.GetText())
		fmt.Printf("l.stack: %v\n", l.stack)
		l.Index.Clear()

		return
	}

}

func (l *TAFArgumentListener) ExtractPath(t string, myVar string) (path string) {
	return strings.Replace(t, myVar+".", "", 1)
}

// ExitVar_expression is called when exiting the var_expression production.
func (l *TAFArgumentListener) ExitVar_expression(c *parser.Var_expressionContext) {
	log.Println("Scope ", l.Scope, "ExitVar_expression: "+c.GetText())

	if l.VariableContext == nil {
		//no variable context, runtime eror
		return
	}

	if c == nil {
		//runtime error here??
		return
	}
	if c.Var_path() == nil {
		if c.VARIABLE_NAME() == nil {
			//runtime error here??
			return
		}
		v := l.VariableContext.GetVariableIntValue(c.VARIABLE_NAME().GetText())
		l.push(v)
		return
	}
	varName := c.VARIABLE_NAME().GetText()
	if l.Scope > 1 {
		ex := c.GetText()
		//we have a subexpression and its value, we need to push that
		fmt.Printf("l.stack: %v\n", l.stack)
		path := l.ExtractPath(ex, varName)
		if ex == varName {
			v := l.VariableContext.GetVariableIntValue(varName)
			l.Index.Push(ex, int(v))
			l.push(v)
			return
		}

		evalExpr := l.Index.EvalExpr(ex)
		path = l.ExtractPath(evalExpr, varName)
		if ex == varName {
			v := l.VariableContext.GetVariableIntValue(varName)
			l.Index.Push(ex, int(v))
			l.push(v)
			return
		}

		v := l.VariableContext.EvaluateJSONVariableIntValue(varName, path)
		l.Index.Push(ex, int(v))
		l.push(v)
		return
	}

	expr := l.Index.Eval()
	if expr == varName {
		v := l.VariableContext.GetVariableIntValue(varName)
		l.push(v)
		return
	}

	path := l.ExtractPath(expr, varName)
	v := l.VariableContext.EvaluateJSONVariableIntValue(varName, path)
	l.push(v)
	return
}

func (l *TAFArgumentListener) EnterIndexExpression(c *parser.IndexExpressionContext) {
	log.Println("Enter Index Expression")
	//l.CurrentPath = c.GetText()
}

func (l *TAFArgumentListener) ExitIndexExpression(c *parser.IndexExpressionContext) {
	log.Println(l.Scope, " Exit Index Expression ")
	p := c.GetText()

	index := l.pop()
	l.Index.Push(p, index.(int))

	if l.Scope == 1 {
		//l.Index.Clear()
	}
	return
}

func (l *TAFArgumentListener) EnterVarPath(c *parser.VarPathContext) {
	/*	log.Println("EnterVarPath " + c.GetText())
		l.CurrentPath = c.GetText()*/
}

func (l *TAFArgumentListener) ExitVarPath(c *parser.VarPathContext) {
	log.Println("ExitVarPath " + c.GetText())
	//l.CurrentPath = c.GetText()
	//fmt.Printf("l.CurrentPath: %v\n", l.CurrentPath)
}

// ExitParenthesisExpression is called when exiting the parenthesisExpression production.
func (l *TAFArgumentListener) ExitParenthesisExpression(c *parser.ParenthesisExpressionContext) {
	log.Println("ExitParenthesisExpression" + c.GetText())
}
