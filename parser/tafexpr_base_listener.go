// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tafexpr
import "github.com/antlr4-go/antlr/v4"

// BaseTafexprListener is a complete listener for a parse tree produced by TafexprParser.
type BaseTafexprListener struct{}

var _ TafexprListener = &BaseTafexprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTafexprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTafexprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTafexprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTafexprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTaf_expression is called when production taf_expression is entered.
func (s *BaseTafexprListener) EnterTaf_expression(ctx *Taf_expressionContext) {}

// ExitTaf_expression is called when production taf_expression is exited.
func (s *BaseTafexprListener) ExitTaf_expression(ctx *Taf_expressionContext) {}

// EnterHandleLogicalNegation is called when production HandleLogicalNegation is entered.
func (s *BaseTafexprListener) EnterHandleLogicalNegation(ctx *HandleLogicalNegationContext) {}

// ExitHandleLogicalNegation is called when production HandleLogicalNegation is exited.
func (s *BaseTafexprListener) ExitHandleLogicalNegation(ctx *HandleLogicalNegationContext) {}

// EnterHandleNegation is called when production HandleNegation is entered.
func (s *BaseTafexprListener) EnterHandleNegation(ctx *HandleNegationContext) {}

// ExitHandleNegation is called when production HandleNegation is exited.
func (s *BaseTafexprListener) ExitHandleNegation(ctx *HandleNegationContext) {}

// EnterHandleVarExpression is called when production HandleVarExpression is entered.
func (s *BaseTafexprListener) EnterHandleVarExpression(ctx *HandleVarExpressionContext) {}

// ExitHandleVarExpression is called when production HandleVarExpression is exited.
func (s *BaseTafexprListener) ExitHandleVarExpression(ctx *HandleVarExpressionContext) {}

// EnterHandleMethodCall is called when production HandleMethodCall is entered.
func (s *BaseTafexprListener) EnterHandleMethodCall(ctx *HandleMethodCallContext) {}

// ExitHandleMethodCall is called when production HandleMethodCall is exited.
func (s *BaseTafexprListener) ExitHandleMethodCall(ctx *HandleMethodCallContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseTafexprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseTafexprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseTafexprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseTafexprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterHandleFunctionCall is called when production HandleFunctionCall is entered.
func (s *BaseTafexprListener) EnterHandleFunctionCall(ctx *HandleFunctionCallContext) {}

// ExitHandleFunctionCall is called when production HandleFunctionCall is exited.
func (s *BaseTafexprListener) ExitHandleFunctionCall(ctx *HandleFunctionCallContext) {}

// EnterHandleNull is called when production HandleNull is entered.
func (s *BaseTafexprListener) EnterHandleNull(ctx *HandleNullContext) {}

// ExitHandleNull is called when production HandleNull is exited.
func (s *BaseTafexprListener) ExitHandleNull(ctx *HandleNullContext) {}

// EnterHandleString is called when production HandleString is entered.
func (s *BaseTafexprListener) EnterHandleString(ctx *HandleStringContext) {}

// ExitHandleString is called when production HandleString is exited.
func (s *BaseTafexprListener) ExitHandleString(ctx *HandleStringContext) {}

// EnterOrderedEvaluation is called when production OrderedEvaluation is entered.
func (s *BaseTafexprListener) EnterOrderedEvaluation(ctx *OrderedEvaluationContext) {}

// ExitOrderedEvaluation is called when production OrderedEvaluation is exited.
func (s *BaseTafexprListener) ExitOrderedEvaluation(ctx *OrderedEvaluationContext) {}

// EnterLogicalOperation is called when production LogicalOperation is entered.
func (s *BaseTafexprListener) EnterLogicalOperation(ctx *LogicalOperationContext) {}

// ExitLogicalOperation is called when production LogicalOperation is exited.
func (s *BaseTafexprListener) ExitLogicalOperation(ctx *LogicalOperationContext) {}

// EnterHandleBool is called when production HandleBool is entered.
func (s *BaseTafexprListener) EnterHandleBool(ctx *HandleBoolContext) {}

// ExitHandleBool is called when production HandleBool is exited.
func (s *BaseTafexprListener) ExitHandleBool(ctx *HandleBoolContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseTafexprListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseTafexprListener) ExitNumber(ctx *NumberContext) {}

// EnterHandleJson is called when production HandleJson is entered.
func (s *BaseTafexprListener) EnterHandleJson(ctx *HandleJsonContext) {}

// ExitHandleJson is called when production HandleJson is exited.
func (s *BaseTafexprListener) ExitHandleJson(ctx *HandleJsonContext) {}

// EnterDoubleValue is called when production DoubleValue is entered.
func (s *BaseTafexprListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production DoubleValue is exited.
func (s *BaseTafexprListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterHandleLogical is called when production HandleLogical is entered.
func (s *BaseTafexprListener) EnterHandleLogical(ctx *HandleLogicalContext) {}

// ExitHandleLogical is called when production HandleLogical is exited.
func (s *BaseTafexprListener) ExitHandleLogical(ctx *HandleLogicalContext) {}

// EnterVar_expression is called when production var_expression is entered.
func (s *BaseTafexprListener) EnterVar_expression(ctx *Var_expressionContext) {}

// ExitVar_expression is called when production var_expression is exited.
func (s *BaseTafexprListener) ExitVar_expression(ctx *Var_expressionContext) {}

// EnterIndx_expr is called when production indx_expr is entered.
func (s *BaseTafexprListener) EnterIndx_expr(ctx *Indx_exprContext) {}

// ExitIndx_expr is called when production indx_expr is exited.
func (s *BaseTafexprListener) ExitIndx_expr(ctx *Indx_exprContext) {}

// EnterVar_path is called when production var_path is entered.
func (s *BaseTafexprListener) EnterVar_path(ctx *Var_pathContext) {}

// ExitVar_path is called when production var_path is exited.
func (s *BaseTafexprListener) ExitVar_path(ctx *Var_pathContext) {}

// EnterJsonpath_expr is called when production jsonpath_expr is entered.
func (s *BaseTafexprListener) EnterJsonpath_expr(ctx *Jsonpath_exprContext) {}

// ExitJsonpath_expr is called when production jsonpath_expr is exited.
func (s *BaseTafexprListener) ExitJsonpath_expr(ctx *Jsonpath_exprContext) {}

// EnterIdentifierWithQualifier is called when production identifierWithQualifier is entered.
func (s *BaseTafexprListener) EnterIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}

// ExitIdentifierWithQualifier is called when production identifierWithQualifier is exited.
func (s *BaseTafexprListener) ExitIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}

// EnterIndexExpression is called when production IndexExpression is entered.
func (s *BaseTafexprListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production IndexExpression is exited.
func (s *BaseTafexprListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterParenthesisExpression is called when production parenthesisExpression is entered.
func (s *BaseTafexprListener) EnterParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// ExitParenthesisExpression is called when production parenthesisExpression is exited.
func (s *BaseTafexprListener) ExitParenthesisExpression(ctx *ParenthesisExpressionContext) {}

// EnterHandleObject is called when production HandleObject is entered.
func (s *BaseTafexprListener) EnterHandleObject(ctx *HandleObjectContext) {}

// ExitHandleObject is called when production HandleObject is exited.
func (s *BaseTafexprListener) ExitHandleObject(ctx *HandleObjectContext) {}

// EnterHandleArray is called when production HandleArray is entered.
func (s *BaseTafexprListener) EnterHandleArray(ctx *HandleArrayContext) {}

// ExitHandleArray is called when production HandleArray is exited.
func (s *BaseTafexprListener) ExitHandleArray(ctx *HandleArrayContext) {}

// EnterHandleObjectData is called when production HandleObjectData is entered.
func (s *BaseTafexprListener) EnterHandleObjectData(ctx *HandleObjectDataContext) {}

// ExitHandleObjectData is called when production HandleObjectData is exited.
func (s *BaseTafexprListener) ExitHandleObjectData(ctx *HandleObjectDataContext) {}

// EnterHandleEmptyObjectData is called when production HandleEmptyObjectData is entered.
func (s *BaseTafexprListener) EnterHandleEmptyObjectData(ctx *HandleEmptyObjectDataContext) {}

// ExitHandleEmptyObjectData is called when production HandleEmptyObjectData is exited.
func (s *BaseTafexprListener) ExitHandleEmptyObjectData(ctx *HandleEmptyObjectDataContext) {}

// EnterHandleObjectPair is called when production HandleObjectPair is entered.
func (s *BaseTafexprListener) EnterHandleObjectPair(ctx *HandleObjectPairContext) {}

// ExitHandleObjectPair is called when production HandleObjectPair is exited.
func (s *BaseTafexprListener) ExitHandleObjectPair(ctx *HandleObjectPairContext) {}

// EnterArr is called when production arr is entered.
func (s *BaseTafexprListener) EnterArr(ctx *ArrContext) {}

// ExitArr is called when production arr is exited.
func (s *BaseTafexprListener) ExitArr(ctx *ArrContext) {}

// EnterHandleJJ is called when production HandleJJ is entered.
func (s *BaseTafexprListener) EnterHandleJJ(ctx *HandleJJContext) {}

// ExitHandleJJ is called when production HandleJJ is exited.
func (s *BaseTafexprListener) ExitHandleJJ(ctx *HandleJJContext) {}

// EnterHandleFoo is called when production HandleFoo is entered.
func (s *BaseTafexprListener) EnterHandleFoo(ctx *HandleFooContext) {}

// ExitHandleFoo is called when production HandleFoo is exited.
func (s *BaseTafexprListener) ExitHandleFoo(ctx *HandleFooContext) {}
