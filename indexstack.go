package tafexpr

import (
	"strconv"
	"strings"

	"github.com/gclkaze/evalang/evalangparser/utils"
	regexp "github.com/s-kozlov/goback/regexp"

	log "github.com/sirupsen/logrus"
)

type IndexValue struct {
	Expr  string
	Value int
}

type IndexStack struct {
	value             int
	symbolStartPos    int
	varExpression     string
	evalVarExpression string
	change            bool
	init              bool
	stack             []IndexValue
}

func (ind *IndexStack) SetExpression(v string) {
	if ind.init == true {
		return
	}
	ind.varExpression = v
	ind.evalVarExpression = v
	ind.init = true
}

func (ind IndexStack) IsEmpty() bool {
	return ind.varExpression == "" && ind.evalVarExpression == "" && ind.value == 0 && ind.symbolStartPos == 0
}

func (ind IndexStack) IsActuallyEmpty() bool {
	return len(ind.stack) == 0
}

func (ind *IndexStack) Push(expr string, val int) {
	ind.stack = append(ind.stack, IndexValue{Expr: expr, Value: val})
	//ind.evalVarExpression = strings.Replace(ind.evalVarExpression, expr, strconv.Itoa(val), 1)
}

func (ind *IndexStack) Clear() {
	ind.varExpression = ""
	ind.evalVarExpression = ""

	ind.value = 0
	ind.symbolStartPos = 0
	ind.change = false
	ind.stack = nil
	ind.init = false
}

func (ind IndexStack) Eval() string {
	/*	log.Print("\n\n===========================================>")
		log.Println("Initially ", ind.varExpression)
		log.Println("Rep ", ind.evalVarExpression)
		log.Print("===========================================>")*/

	s := ind.EvalExpr(ind.varExpression)
	//log.Println("s: ", s)
	return s
}

func (ind IndexStack) EvalExpr(expr string) string {
	if len(ind.stack) == 0 {
		log.Debug("stack is empty")
		return expr
	}
	rep := expr
	for i := len(ind.stack) - 1; i >= 0; i-- {
		rep = ind.ReplaceValue(rep, ind.stack[i])
	}
	return rep
}

func (ind IndexStack) OldReplace(rep string, stored IndexValue) string {
	rep = strings.Replace(rep, stored.Expr, strconv.Itoa(stored.Value), 1)
	return rep
}

func (ind IndexStack) ReplaceValue(rep string, stored IndexValue) string {
	regx := utils.GetBracketedString(stored.Expr) + `(?=[^a-zA-Z0-9])`
	r := regexp.MustCompile(regx)
	/*	fmt.Printf("**Before : %v\n", rep)
		b := r.MatchString(rep)
		fmt.Printf("**b: %v\n", b)
		fmt.Println("**Replacing " + stored.Expr + " => " + strconv.Itoa(stored.Value))*/
	rep = r.ReplaceAllString(rep, strconv.Itoa(stored.Value))
	/*	fmt.Println("**After " + rep)*/

	return rep
}
