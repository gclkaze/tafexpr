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

// EnterHandleRandomDoubleInRange is called when production HandleRandomDoubleInRange is entered.
func (s *BaseTafexprListener) EnterHandleRandomDoubleInRange(ctx *HandleRandomDoubleInRangeContext) {}

// ExitHandleRandomDoubleInRange is called when production HandleRandomDoubleInRange is exited.
func (s *BaseTafexprListener) ExitHandleRandomDoubleInRange(ctx *HandleRandomDoubleInRangeContext) {}

// EnterHandleFindOneByXPATH is called when production HandleFindOneByXPATH is entered.
func (s *BaseTafexprListener) EnterHandleFindOneByXPATH(ctx *HandleFindOneByXPATHContext) {}

// ExitHandleFindOneByXPATH is called when production HandleFindOneByXPATH is exited.
func (s *BaseTafexprListener) ExitHandleFindOneByXPATH(ctx *HandleFindOneByXPATHContext) {}

// EnterHandleTrim is called when production HandleTrim is entered.
func (s *BaseTafexprListener) EnterHandleTrim(ctx *HandleTrimContext) {}

// ExitHandleTrim is called when production HandleTrim is exited.
func (s *BaseTafexprListener) ExitHandleTrim(ctx *HandleTrimContext) {}

// EnterHandleVarExpression is called when production HandleVarExpression is entered.
func (s *BaseTafexprListener) EnterHandleVarExpression(ctx *HandleVarExpressionContext) {}

// ExitHandleVarExpression is called when production HandleVarExpression is exited.
func (s *BaseTafexprListener) ExitHandleVarExpression(ctx *HandleVarExpressionContext) {}

// EnterHandleFindOneDoubleByXPATH is called when production HandleFindOneDoubleByXPATH is entered.
func (s *BaseTafexprListener) EnterHandleFindOneDoubleByXPATH(ctx *HandleFindOneDoubleByXPATHContext) {
}

// ExitHandleFindOneDoubleByXPATH is called when production HandleFindOneDoubleByXPATH is exited.
func (s *BaseTafexprListener) ExitHandleFindOneDoubleByXPATH(ctx *HandleFindOneDoubleByXPATHContext) {
}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseTafexprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseTafexprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterHandleFindOneStringByXPATH is called when production HandleFindOneStringByXPATH is entered.
func (s *BaseTafexprListener) EnterHandleFindOneStringByXPATH(ctx *HandleFindOneStringByXPATHContext) {
}

// ExitHandleFindOneStringByXPATH is called when production HandleFindOneStringByXPATH is exited.
func (s *BaseTafexprListener) ExitHandleFindOneStringByXPATH(ctx *HandleFindOneStringByXPATHContext) {
}

// EnterHandleToString is called when production HandleToString is entered.
func (s *BaseTafexprListener) EnterHandleToString(ctx *HandleToStringContext) {}

// ExitHandleToString is called when production HandleToString is exited.
func (s *BaseTafexprListener) ExitHandleToString(ctx *HandleToStringContext) {}

// EnterHandleLibfunc is called when production HandleLibfunc is entered.
func (s *BaseTafexprListener) EnterHandleLibfunc(ctx *HandleLibfuncContext) {}

// ExitHandleLibfunc is called when production HandleLibfunc is exited.
func (s *BaseTafexprListener) ExitHandleLibfunc(ctx *HandleLibfuncContext) {}

// EnterHandleFindByXPATH is called when production HandleFindByXPATH is entered.
func (s *BaseTafexprListener) EnterHandleFindByXPATH(ctx *HandleFindByXPATHContext) {}

// ExitHandleFindByXPATH is called when production HandleFindByXPATH is exited.
func (s *BaseTafexprListener) ExitHandleFindByXPATH(ctx *HandleFindByXPATHContext) {}

// EnterHandleString is called when production HandleString is entered.
func (s *BaseTafexprListener) EnterHandleString(ctx *HandleStringContext) {}

// ExitHandleString is called when production HandleString is exited.
func (s *BaseTafexprListener) ExitHandleString(ctx *HandleStringContext) {}

// EnterHandleExtractOneByREGEX is called when production HandleExtractOneByREGEX is entered.
func (s *BaseTafexprListener) EnterHandleExtractOneByREGEX(ctx *HandleExtractOneByREGEXContext) {}

// ExitHandleExtractOneByREGEX is called when production HandleExtractOneByREGEX is exited.
func (s *BaseTafexprListener) ExitHandleExtractOneByREGEX(ctx *HandleExtractOneByREGEXContext) {}

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

// EnterHandleLogical is called when production HandleLogical is entered.
func (s *BaseTafexprListener) EnterHandleLogical(ctx *HandleLogicalContext) {}

// ExitHandleLogical is called when production HandleLogical is exited.
func (s *BaseTafexprListener) ExitHandleLogical(ctx *HandleLogicalContext) {}

// EnterHandleToBoolean is called when production HandleToBoolean is entered.
func (s *BaseTafexprListener) EnterHandleToBoolean(ctx *HandleToBooleanContext) {}

// ExitHandleToBoolean is called when production HandleToBoolean is exited.
func (s *BaseTafexprListener) ExitHandleToBoolean(ctx *HandleToBooleanContext) {}

// EnterHandleTrimLeft is called when production HandleTrimLeft is entered.
func (s *BaseTafexprListener) EnterHandleTrimLeft(ctx *HandleTrimLeftContext) {}

// ExitHandleTrimLeft is called when production HandleTrimLeft is exited.
func (s *BaseTafexprListener) ExitHandleTrimLeft(ctx *HandleTrimLeftContext) {}

// EnterHandleFindOneBooleanByXPATH is called when production HandleFindOneBooleanByXPATH is entered.
func (s *BaseTafexprListener) EnterHandleFindOneBooleanByXPATH(ctx *HandleFindOneBooleanByXPATHContext) {
}

// ExitHandleFindOneBooleanByXPATH is called when production HandleFindOneBooleanByXPATH is exited.
func (s *BaseTafexprListener) ExitHandleFindOneBooleanByXPATH(ctx *HandleFindOneBooleanByXPATHContext) {
}

// EnterHandleLogicalNegation is called when production HandleLogicalNegation is entered.
func (s *BaseTafexprListener) EnterHandleLogicalNegation(ctx *HandleLogicalNegationContext) {}

// ExitHandleLogicalNegation is called when production HandleLogicalNegation is exited.
func (s *BaseTafexprListener) ExitHandleLogicalNegation(ctx *HandleLogicalNegationContext) {}

// EnterHandleLength is called when production HandleLength is entered.
func (s *BaseTafexprListener) EnterHandleLength(ctx *HandleLengthContext) {}

// ExitHandleLength is called when production HandleLength is exited.
func (s *BaseTafexprListener) ExitHandleLength(ctx *HandleLengthContext) {}

// EnterHandleNegation is called when production HandleNegation is entered.
func (s *BaseTafexprListener) EnterHandleNegation(ctx *HandleNegationContext) {}

// ExitHandleNegation is called when production HandleNegation is exited.
func (s *BaseTafexprListener) ExitHandleNegation(ctx *HandleNegationContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseTafexprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseTafexprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterHandleFindOneIntegerByXPATH is called when production HandleFindOneIntegerByXPATH is entered.
func (s *BaseTafexprListener) EnterHandleFindOneIntegerByXPATH(ctx *HandleFindOneIntegerByXPATHContext) {
}

// ExitHandleFindOneIntegerByXPATH is called when production HandleFindOneIntegerByXPATH is exited.
func (s *BaseTafexprListener) ExitHandleFindOneIntegerByXPATH(ctx *HandleFindOneIntegerByXPATHContext) {
}

// EnterHandleNull is called when production HandleNull is entered.
func (s *BaseTafexprListener) EnterHandleNull(ctx *HandleNullContext) {}

// ExitHandleNull is called when production HandleNull is exited.
func (s *BaseTafexprListener) ExitHandleNull(ctx *HandleNullContext) {}

// EnterHandleToDouble is called when production HandleToDouble is entered.
func (s *BaseTafexprListener) EnterHandleToDouble(ctx *HandleToDoubleContext) {}

// ExitHandleToDouble is called when production HandleToDouble is exited.
func (s *BaseTafexprListener) ExitHandleToDouble(ctx *HandleToDoubleContext) {}

// EnterHandleEndsWith is called when production HandleEndsWith is entered.
func (s *BaseTafexprListener) EnterHandleEndsWith(ctx *HandleEndsWithContext) {}

// ExitHandleEndsWith is called when production HandleEndsWith is exited.
func (s *BaseTafexprListener) ExitHandleEndsWith(ctx *HandleEndsWithContext) {}

// EnterHandleContainsString is called when production HandleContainsString is entered.
func (s *BaseTafexprListener) EnterHandleContainsString(ctx *HandleContainsStringContext) {}

// ExitHandleContainsString is called when production HandleContainsString is exited.
func (s *BaseTafexprListener) ExitHandleContainsString(ctx *HandleContainsStringContext) {}

// EnterOrderedEvaluation is called when production OrderedEvaluation is entered.
func (s *BaseTafexprListener) EnterOrderedEvaluation(ctx *OrderedEvaluationContext) {}

// ExitOrderedEvaluation is called when production OrderedEvaluation is exited.
func (s *BaseTafexprListener) ExitOrderedEvaluation(ctx *OrderedEvaluationContext) {}

// EnterLogicalOperation is called when production LogicalOperation is entered.
func (s *BaseTafexprListener) EnterLogicalOperation(ctx *LogicalOperationContext) {}

// ExitLogicalOperation is called when production LogicalOperation is exited.
func (s *BaseTafexprListener) ExitLogicalOperation(ctx *LogicalOperationContext) {}

// EnterHandleToInteger is called when production HandleToInteger is entered.
func (s *BaseTafexprListener) EnterHandleToInteger(ctx *HandleToIntegerContext) {}

// ExitHandleToInteger is called when production HandleToInteger is exited.
func (s *BaseTafexprListener) ExitHandleToInteger(ctx *HandleToIntegerContext) {}

// EnterHandleTrimRight is called when production HandleTrimRight is entered.
func (s *BaseTafexprListener) EnterHandleTrimRight(ctx *HandleTrimRightContext) {}

// ExitHandleTrimRight is called when production HandleTrimRight is exited.
func (s *BaseTafexprListener) ExitHandleTrimRight(ctx *HandleTrimRightContext) {}

// EnterHandleStartsWith is called when production HandleStartsWith is entered.
func (s *BaseTafexprListener) EnterHandleStartsWith(ctx *HandleStartsWithContext) {}

// ExitHandleStartsWith is called when production HandleStartsWith is exited.
func (s *BaseTafexprListener) ExitHandleStartsWith(ctx *HandleStartsWithContext) {}

// EnterDoubleValue is called when production DoubleValue is entered.
func (s *BaseTafexprListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production DoubleValue is exited.
func (s *BaseTafexprListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterHandleReplaceAllStringOccurrences is called when production HandleReplaceAllStringOccurrences is entered.
func (s *BaseTafexprListener) EnterHandleReplaceAllStringOccurrences(ctx *HandleReplaceAllStringOccurrencesContext) {
}

// ExitHandleReplaceAllStringOccurrences is called when production HandleReplaceAllStringOccurrences is exited.
func (s *BaseTafexprListener) ExitHandleReplaceAllStringOccurrences(ctx *HandleReplaceAllStringOccurrencesContext) {
}

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

// EnterAny is called when production any is entered.
func (s *BaseTafexprListener) EnterAny(ctx *AnyContext) {}

// ExitAny is called when production any is exited.
func (s *BaseTafexprListener) ExitAny(ctx *AnyContext) {}

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
