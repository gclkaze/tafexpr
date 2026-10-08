// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tafexpr
import "github.com/antlr4-go/antlr/v4"

// TafexprListener is a complete listener for a parse tree produced by TafexprParser.
type TafexprListener interface {
	antlr.ParseTreeListener

	// EnterTaf_expression is called when entering the taf_expression production.
	EnterTaf_expression(c *Taf_expressionContext)

	// EnterHandleLogicalNegation is called when entering the HandleLogicalNegation production.
	EnterHandleLogicalNegation(c *HandleLogicalNegationContext)

	// EnterHandleNegation is called when entering the HandleNegation production.
	EnterHandleNegation(c *HandleNegationContext)

	// EnterHandleVarExpression is called when entering the HandleVarExpression production.
	EnterHandleVarExpression(c *HandleVarExpressionContext)

	// EnterHandleMethodCall is called when entering the HandleMethodCall production.
	EnterHandleMethodCall(c *HandleMethodCallContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterHandleFunctionCall is called when entering the HandleFunctionCall production.
	EnterHandleFunctionCall(c *HandleFunctionCallContext)

	// EnterHandleNull is called when entering the HandleNull production.
	EnterHandleNull(c *HandleNullContext)

	// EnterHandleString is called when entering the HandleString production.
	EnterHandleString(c *HandleStringContext)

	// EnterOrderedEvaluation is called when entering the OrderedEvaluation production.
	EnterOrderedEvaluation(c *OrderedEvaluationContext)

	// EnterLogicalOperation is called when entering the LogicalOperation production.
	EnterLogicalOperation(c *LogicalOperationContext)

	// EnterHandleBool is called when entering the HandleBool production.
	EnterHandleBool(c *HandleBoolContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterHandleJson is called when entering the HandleJson production.
	EnterHandleJson(c *HandleJsonContext)

	// EnterDoubleValue is called when entering the DoubleValue production.
	EnterDoubleValue(c *DoubleValueContext)

	// EnterHandleLogical is called when entering the HandleLogical production.
	EnterHandleLogical(c *HandleLogicalContext)

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

	// ExitHandleLogicalNegation is called when exiting the HandleLogicalNegation production.
	ExitHandleLogicalNegation(c *HandleLogicalNegationContext)

	// ExitHandleNegation is called when exiting the HandleNegation production.
	ExitHandleNegation(c *HandleNegationContext)

	// ExitHandleVarExpression is called when exiting the HandleVarExpression production.
	ExitHandleVarExpression(c *HandleVarExpressionContext)

	// ExitHandleMethodCall is called when exiting the HandleMethodCall production.
	ExitHandleMethodCall(c *HandleMethodCallContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitHandleFunctionCall is called when exiting the HandleFunctionCall production.
	ExitHandleFunctionCall(c *HandleFunctionCallContext)

	// ExitHandleNull is called when exiting the HandleNull production.
	ExitHandleNull(c *HandleNullContext)

	// ExitHandleString is called when exiting the HandleString production.
	ExitHandleString(c *HandleStringContext)

	// ExitOrderedEvaluation is called when exiting the OrderedEvaluation production.
	ExitOrderedEvaluation(c *OrderedEvaluationContext)

	// ExitLogicalOperation is called when exiting the LogicalOperation production.
	ExitLogicalOperation(c *LogicalOperationContext)

	// ExitHandleBool is called when exiting the HandleBool production.
	ExitHandleBool(c *HandleBoolContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitHandleJson is called when exiting the HandleJson production.
	ExitHandleJson(c *HandleJsonContext)

	// ExitDoubleValue is called when exiting the DoubleValue production.
	ExitDoubleValue(c *DoubleValueContext)

	// ExitHandleLogical is called when exiting the HandleLogical production.
	ExitHandleLogical(c *HandleLogicalContext)

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
