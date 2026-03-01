// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tafexpr

import "github.com/antlr4-go/antlr/v4"

// TafexprListener is a complete listener for a parse tree produced by TafexprParser.
type TafexprListener interface {
	antlr.ParseTreeListener

	// EnterTaf_expression is called when entering the taf_expression production.
	EnterTaf_expression(c *Taf_expressionContext)

	// EnterHandleRandomDoubleInRange is called when entering the HandleRandomDoubleInRange production.
	EnterHandleRandomDoubleInRange(c *HandleRandomDoubleInRangeContext)

	// EnterHandleFindOneByXPATH is called when entering the HandleFindOneByXPATH production.
	EnterHandleFindOneByXPATH(c *HandleFindOneByXPATHContext)

	// EnterHandleTrim is called when entering the HandleTrim production.
	EnterHandleTrim(c *HandleTrimContext)

	// EnterHandleVarExpression is called when entering the HandleVarExpression production.
	EnterHandleVarExpression(c *HandleVarExpressionContext)

	// EnterHandleFindOneDoubleByXPATH is called when entering the HandleFindOneDoubleByXPATH production.
	EnterHandleFindOneDoubleByXPATH(c *HandleFindOneDoubleByXPATHContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterHandleFindOneStringByXPATH is called when entering the HandleFindOneStringByXPATH production.
	EnterHandleFindOneStringByXPATH(c *HandleFindOneStringByXPATHContext)

	// EnterHandleToString is called when entering the HandleToString production.
	EnterHandleToString(c *HandleToStringContext)

	// EnterHandleLibfunc is called when entering the HandleLibfunc production.
	EnterHandleLibfunc(c *HandleLibfuncContext)

	// EnterHandleFindByXPATH is called when entering the HandleFindByXPATH production.
	EnterHandleFindByXPATH(c *HandleFindByXPATHContext)

	// EnterHandleString is called when entering the HandleString production.
	EnterHandleString(c *HandleStringContext)

	// EnterHandleExtractOneByREGEX is called when entering the HandleExtractOneByREGEX production.
	EnterHandleExtractOneByREGEX(c *HandleExtractOneByREGEXContext)

	// EnterHandleBool is called when entering the HandleBool production.
	EnterHandleBool(c *HandleBoolContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterHandleJson is called when entering the HandleJson production.
	EnterHandleJson(c *HandleJsonContext)

	// EnterHandleLogical is called when entering the HandleLogical production.
	EnterHandleLogical(c *HandleLogicalContext)

	// EnterHandleToBoolean is called when entering the HandleToBoolean production.
	EnterHandleToBoolean(c *HandleToBooleanContext)

	// EnterHandleTrimLeft is called when entering the HandleTrimLeft production.
	EnterHandleTrimLeft(c *HandleTrimLeftContext)

	// EnterHandleFindOneBooleanByXPATH is called when entering the HandleFindOneBooleanByXPATH production.
	EnterHandleFindOneBooleanByXPATH(c *HandleFindOneBooleanByXPATHContext)

	// EnterHandleLogicalNegation is called when entering the HandleLogicalNegation production.
	EnterHandleLogicalNegation(c *HandleLogicalNegationContext)

	// EnterHandleLength is called when entering the HandleLength production.
	EnterHandleLength(c *HandleLengthContext)

	// EnterHandleNegation is called when entering the HandleNegation production.
	EnterHandleNegation(c *HandleNegationContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterHandleFindOneIntegerByXPATH is called when entering the HandleFindOneIntegerByXPATH production.
	EnterHandleFindOneIntegerByXPATH(c *HandleFindOneIntegerByXPATHContext)

	// EnterHandleNull is called when entering the HandleNull production.
	EnterHandleNull(c *HandleNullContext)

	// EnterHandleToDouble is called when entering the HandleToDouble production.
	EnterHandleToDouble(c *HandleToDoubleContext)

	// EnterHandleEndsWith is called when entering the HandleEndsWith production.
	EnterHandleEndsWith(c *HandleEndsWithContext)

	// EnterHandleContainsString is called when entering the HandleContainsString production.
	EnterHandleContainsString(c *HandleContainsStringContext)

	// EnterOrderedEvaluation is called when entering the OrderedEvaluation production.
	EnterOrderedEvaluation(c *OrderedEvaluationContext)

	// EnterLogicalOperation is called when entering the LogicalOperation production.
	EnterLogicalOperation(c *LogicalOperationContext)

	// EnterHandleToInteger is called when entering the HandleToInteger production.
	EnterHandleToInteger(c *HandleToIntegerContext)

	// EnterHandleTrimRight is called when entering the HandleTrimRight production.
	EnterHandleTrimRight(c *HandleTrimRightContext)

	// EnterHandleStartsWith is called when entering the HandleStartsWith production.
	EnterHandleStartsWith(c *HandleStartsWithContext)

	// EnterDoubleValue is called when entering the DoubleValue production.
	EnterDoubleValue(c *DoubleValueContext)

	// EnterHandleReplaceAllStringOccurrences is called when entering the HandleReplaceAllStringOccurrences production.
	EnterHandleReplaceAllStringOccurrences(c *HandleReplaceAllStringOccurrencesContext)

	// EnterVar_expression is called when entering the var_expression production.
	EnterVar_expression(c *Var_expressionContext)

	// EnterIndx_expr is called when entering the indx_expr production.
	EnterIndx_expr(c *Indx_exprContext)

	// EnterVar_path is called when entering the var_path production.
	EnterVar_path(c *Var_pathContext)

	// EnterJsonpath_expr is called when entering the jsonpath_expr production.
	EnterJsonpath_expr(c *Jsonpath_exprContext)

	// EnterIdentifierWithQualifier is called when entering the identifierWithQualifier production.
	EnterIdentifierWithQualifier(c *IdentifierWithQualifierContext)

	// EnterIndexExpression is called when entering the IndexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterParenthesisExpression is called when entering the parenthesisExpression production.
	EnterParenthesisExpression(c *ParenthesisExpressionContext)

	// EnterAny is called when entering the any production.
	EnterAny(c *AnyContext)

	// EnterHandleObject is called when entering the HandleObject production.
	EnterHandleObject(c *HandleObjectContext)

	// EnterHandleArray is called when entering the HandleArray production.
	EnterHandleArray(c *HandleArrayContext)

	// EnterHandleObjectData is called when entering the HandleObjectData production.
	EnterHandleObjectData(c *HandleObjectDataContext)

	// EnterHandleEmptyObjectData is called when entering the HandleEmptyObjectData production.
	EnterHandleEmptyObjectData(c *HandleEmptyObjectDataContext)

	// EnterHandleObjectPair is called when entering the HandleObjectPair production.
	EnterHandleObjectPair(c *HandleObjectPairContext)

	// EnterArr is called when entering the arr production.
	EnterArr(c *ArrContext)

	// EnterHandleJJ is called when entering the HandleJJ production.
	EnterHandleJJ(c *HandleJJContext)

	// EnterHandleFoo is called when entering the HandleFoo production.
	EnterHandleFoo(c *HandleFooContext)

	// ExitTaf_expression is called when exiting the taf_expression production.
	ExitTaf_expression(c *Taf_expressionContext)

	// ExitHandleRandomDoubleInRange is called when exiting the HandleRandomDoubleInRange production.
	ExitHandleRandomDoubleInRange(c *HandleRandomDoubleInRangeContext)

	// ExitHandleFindOneByXPATH is called when exiting the HandleFindOneByXPATH production.
	ExitHandleFindOneByXPATH(c *HandleFindOneByXPATHContext)

	// ExitHandleTrim is called when exiting the HandleTrim production.
	ExitHandleTrim(c *HandleTrimContext)

	// ExitHandleVarExpression is called when exiting the HandleVarExpression production.
	ExitHandleVarExpression(c *HandleVarExpressionContext)

	// ExitHandleFindOneDoubleByXPATH is called when exiting the HandleFindOneDoubleByXPATH production.
	ExitHandleFindOneDoubleByXPATH(c *HandleFindOneDoubleByXPATHContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitHandleFindOneStringByXPATH is called when exiting the HandleFindOneStringByXPATH production.
	ExitHandleFindOneStringByXPATH(c *HandleFindOneStringByXPATHContext)

	// ExitHandleToString is called when exiting the HandleToString production.
	ExitHandleToString(c *HandleToStringContext)

	// ExitHandleLibfunc is called when exiting the HandleLibfunc production.
	ExitHandleLibfunc(c *HandleLibfuncContext)

	// ExitHandleFindByXPATH is called when exiting the HandleFindByXPATH production.
	ExitHandleFindByXPATH(c *HandleFindByXPATHContext)

	// ExitHandleString is called when exiting the HandleString production.
	ExitHandleString(c *HandleStringContext)

	// ExitHandleExtractOneByREGEX is called when exiting the HandleExtractOneByREGEX production.
	ExitHandleExtractOneByREGEX(c *HandleExtractOneByREGEXContext)

	// ExitHandleBool is called when exiting the HandleBool production.
	ExitHandleBool(c *HandleBoolContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitHandleJson is called when exiting the HandleJson production.
	ExitHandleJson(c *HandleJsonContext)

	// ExitHandleLogical is called when exiting the HandleLogical production.
	ExitHandleLogical(c *HandleLogicalContext)

	// ExitHandleToBoolean is called when exiting the HandleToBoolean production.
	ExitHandleToBoolean(c *HandleToBooleanContext)

	// ExitHandleTrimLeft is called when exiting the HandleTrimLeft production.
	ExitHandleTrimLeft(c *HandleTrimLeftContext)

	// ExitHandleFindOneBooleanByXPATH is called when exiting the HandleFindOneBooleanByXPATH production.
	ExitHandleFindOneBooleanByXPATH(c *HandleFindOneBooleanByXPATHContext)

	// ExitHandleLogicalNegation is called when exiting the HandleLogicalNegation production.
	ExitHandleLogicalNegation(c *HandleLogicalNegationContext)

	// ExitHandleLength is called when exiting the HandleLength production.
	ExitHandleLength(c *HandleLengthContext)

	// ExitHandleNegation is called when exiting the HandleNegation production.
	ExitHandleNegation(c *HandleNegationContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitHandleFindOneIntegerByXPATH is called when exiting the HandleFindOneIntegerByXPATH production.
	ExitHandleFindOneIntegerByXPATH(c *HandleFindOneIntegerByXPATHContext)

	// ExitHandleNull is called when exiting the HandleNull production.
	ExitHandleNull(c *HandleNullContext)

	// ExitHandleToDouble is called when exiting the HandleToDouble production.
	ExitHandleToDouble(c *HandleToDoubleContext)

	// ExitHandleEndsWith is called when exiting the HandleEndsWith production.
	ExitHandleEndsWith(c *HandleEndsWithContext)

	// ExitHandleContainsString is called when exiting the HandleContainsString production.
	ExitHandleContainsString(c *HandleContainsStringContext)

	// ExitOrderedEvaluation is called when exiting the OrderedEvaluation production.
	ExitOrderedEvaluation(c *OrderedEvaluationContext)

	// ExitLogicalOperation is called when exiting the LogicalOperation production.
	ExitLogicalOperation(c *LogicalOperationContext)

	// ExitHandleToInteger is called when exiting the HandleToInteger production.
	ExitHandleToInteger(c *HandleToIntegerContext)

	// ExitHandleTrimRight is called when exiting the HandleTrimRight production.
	ExitHandleTrimRight(c *HandleTrimRightContext)

	// ExitHandleStartsWith is called when exiting the HandleStartsWith production.
	ExitHandleStartsWith(c *HandleStartsWithContext)

	// ExitDoubleValue is called when exiting the DoubleValue production.
	ExitDoubleValue(c *DoubleValueContext)

	// ExitHandleReplaceAllStringOccurrences is called when exiting the HandleReplaceAllStringOccurrences production.
	ExitHandleReplaceAllStringOccurrences(c *HandleReplaceAllStringOccurrencesContext)

	// ExitVar_expression is called when exiting the var_expression production.
	ExitVar_expression(c *Var_expressionContext)

	// ExitIndx_expr is called when exiting the indx_expr production.
	ExitIndx_expr(c *Indx_exprContext)

	// ExitVar_path is called when exiting the var_path production.
	ExitVar_path(c *Var_pathContext)

	// ExitJsonpath_expr is called when exiting the jsonpath_expr production.
	ExitJsonpath_expr(c *Jsonpath_exprContext)

	// ExitIdentifierWithQualifier is called when exiting the identifierWithQualifier production.
	ExitIdentifierWithQualifier(c *IdentifierWithQualifierContext)

	// ExitIndexExpression is called when exiting the IndexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitParenthesisExpression is called when exiting the parenthesisExpression production.
	ExitParenthesisExpression(c *ParenthesisExpressionContext)

	// ExitAny is called when exiting the any production.
	ExitAny(c *AnyContext)

	// ExitHandleObject is called when exiting the HandleObject production.
	ExitHandleObject(c *HandleObjectContext)

	// ExitHandleArray is called when exiting the HandleArray production.
	ExitHandleArray(c *HandleArrayContext)

	// ExitHandleObjectData is called when exiting the HandleObjectData production.
	ExitHandleObjectData(c *HandleObjectDataContext)

	// ExitHandleEmptyObjectData is called when exiting the HandleEmptyObjectData production.
	ExitHandleEmptyObjectData(c *HandleEmptyObjectDataContext)

	// ExitHandleObjectPair is called when exiting the HandleObjectPair production.
	ExitHandleObjectPair(c *HandleObjectPairContext)

	// ExitArr is called when exiting the arr production.
	ExitArr(c *ArrContext)

	// ExitHandleJJ is called when exiting the HandleJJ production.
	ExitHandleJJ(c *HandleJJContext)

	// ExitHandleFoo is called when exiting the HandleFoo production.
	ExitHandleFoo(c *HandleFooContext)
}
