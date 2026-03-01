// Generated from c:/Users/gclka/Eva Lang/tafexpr/grammar/Tafexpr.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link TafexprParser}.
 */
public interface TafexprListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link TafexprParser#taf_expression}.
	 * @param ctx the parse tree
	 */
	void enterTaf_expression(TafexprParser.Taf_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#taf_expression}.
	 * @param ctx the parse tree
	 */
	void exitTaf_expression(TafexprParser.Taf_expressionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleRandomDoubleInRange}
	 * labeled alternative in {@link TafexprParser#libfunc}.
	 * @param ctx the parse tree
	 */
	void enterHandleRandomDoubleInRange(TafexprParser.HandleRandomDoubleInRangeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleRandomDoubleInRange}
	 * labeled alternative in {@link TafexprParser#libfunc}.
	 * @param ctx the parse tree
	 */
	void exitHandleRandomDoubleInRange(TafexprParser.HandleRandomDoubleInRangeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleFindOneByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleFindOneByXPATH(TafexprParser.HandleFindOneByXPATHContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleFindOneByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleFindOneByXPATH(TafexprParser.HandleFindOneByXPATHContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleTrim}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleTrim(TafexprParser.HandleTrimContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleTrim}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleTrim(TafexprParser.HandleTrimContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleVarExpression}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleVarExpression(TafexprParser.HandleVarExpressionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleVarExpression}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleVarExpression(TafexprParser.HandleVarExpressionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleFindOneDoubleByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleFindOneDoubleByXPATH(TafexprParser.HandleFindOneDoubleByXPATHContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleFindOneDoubleByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleFindOneDoubleByXPATH(TafexprParser.HandleFindOneDoubleByXPATHContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(TafexprParser.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(TafexprParser.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleFindOneStringByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleFindOneStringByXPATH(TafexprParser.HandleFindOneStringByXPATHContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleFindOneStringByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleFindOneStringByXPATH(TafexprParser.HandleFindOneStringByXPATHContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleToString}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleToString(TafexprParser.HandleToStringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleToString}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleToString(TafexprParser.HandleToStringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleFindByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleFindByXPATH(TafexprParser.HandleFindByXPATHContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleFindByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleFindByXPATH(TafexprParser.HandleFindByXPATHContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleString}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleString(TafexprParser.HandleStringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleString}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleString(TafexprParser.HandleStringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleExtractOneByREGEX}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleExtractOneByREGEX(TafexprParser.HandleExtractOneByREGEXContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleExtractOneByREGEX}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleExtractOneByREGEX(TafexprParser.HandleExtractOneByREGEXContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleBool}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleBool(TafexprParser.HandleBoolContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleBool}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleBool(TafexprParser.HandleBoolContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Number}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterNumber(TafexprParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Number}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitNumber(TafexprParser.NumberContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleJson}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleJson(TafexprParser.HandleJsonContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleJson}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleJson(TafexprParser.HandleJsonContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleLogical}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleLogical(TafexprParser.HandleLogicalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleLogical}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleLogical(TafexprParser.HandleLogicalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleToBoolean}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleToBoolean(TafexprParser.HandleToBooleanContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleToBoolean}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleToBoolean(TafexprParser.HandleToBooleanContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleTrimLeft}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleTrimLeft(TafexprParser.HandleTrimLeftContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleTrimLeft}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleTrimLeft(TafexprParser.HandleTrimLeftContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleFindOneBooleanByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleFindOneBooleanByXPATH(TafexprParser.HandleFindOneBooleanByXPATHContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleFindOneBooleanByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleFindOneBooleanByXPATH(TafexprParser.HandleFindOneBooleanByXPATHContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleLogicalNegation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleLogicalNegation(TafexprParser.HandleLogicalNegationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleLogicalNegation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleLogicalNegation(TafexprParser.HandleLogicalNegationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleLength}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleLength(TafexprParser.HandleLengthContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleLength}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleLength(TafexprParser.HandleLengthContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleNegation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleNegation(TafexprParser.HandleNegationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleNegation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleNegation(TafexprParser.HandleNegationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(TafexprParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(TafexprParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleFindOneIntegerByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleFindOneIntegerByXPATH(TafexprParser.HandleFindOneIntegerByXPATHContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleFindOneIntegerByXPATH}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleFindOneIntegerByXPATH(TafexprParser.HandleFindOneIntegerByXPATHContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleNull}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleNull(TafexprParser.HandleNullContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleNull}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleNull(TafexprParser.HandleNullContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleToDouble}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleToDouble(TafexprParser.HandleToDoubleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleToDouble}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleToDouble(TafexprParser.HandleToDoubleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleEndsWith}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleEndsWith(TafexprParser.HandleEndsWithContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleEndsWith}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleEndsWith(TafexprParser.HandleEndsWithContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleContainsString}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleContainsString(TafexprParser.HandleContainsStringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleContainsString}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleContainsString(TafexprParser.HandleContainsStringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code OrderedEvaluation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterOrderedEvaluation(TafexprParser.OrderedEvaluationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code OrderedEvaluation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitOrderedEvaluation(TafexprParser.OrderedEvaluationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LogicalOperation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOperation(TafexprParser.LogicalOperationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LogicalOperation}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOperation(TafexprParser.LogicalOperationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleToInteger}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleToInteger(TafexprParser.HandleToIntegerContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleToInteger}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleToInteger(TafexprParser.HandleToIntegerContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleTrimRight}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleTrimRight(TafexprParser.HandleTrimRightContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleTrimRight}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleTrimRight(TafexprParser.HandleTrimRightContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleStartsWith}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleStartsWith(TafexprParser.HandleStartsWithContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleStartsWith}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleStartsWith(TafexprParser.HandleStartsWithContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DoubleValue}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterDoubleValue(TafexprParser.DoubleValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DoubleValue}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitDoubleValue(TafexprParser.DoubleValueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleReplaceAllStringOccurrences}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterHandleReplaceAllStringOccurrences(TafexprParser.HandleReplaceAllStringOccurrencesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleReplaceAllStringOccurrences}
	 * labeled alternative in {@link TafexprParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitHandleReplaceAllStringOccurrences(TafexprParser.HandleReplaceAllStringOccurrencesContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#var_expression}.
	 * @param ctx the parse tree
	 */
	void enterVar_expression(TafexprParser.Var_expressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#var_expression}.
	 * @param ctx the parse tree
	 */
	void exitVar_expression(TafexprParser.Var_expressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#indx_expr}.
	 * @param ctx the parse tree
	 */
	void enterIndx_expr(TafexprParser.Indx_exprContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#indx_expr}.
	 * @param ctx the parse tree
	 */
	void exitIndx_expr(TafexprParser.Indx_exprContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#var_path}.
	 * @param ctx the parse tree
	 */
	void enterVar_path(TafexprParser.Var_pathContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#var_path}.
	 * @param ctx the parse tree
	 */
	void exitVar_path(TafexprParser.Var_pathContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#jsonpath_expr}.
	 * @param ctx the parse tree
	 */
	void enterJsonpath_expr(TafexprParser.Jsonpath_exprContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#jsonpath_expr}.
	 * @param ctx the parse tree
	 */
	void exitJsonpath_expr(TafexprParser.Jsonpath_exprContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#identifierWithQualifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifierWithQualifier(TafexprParser.IdentifierWithQualifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#identifierWithQualifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifierWithQualifier(TafexprParser.IdentifierWithQualifierContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IndexExpression}
	 * labeled alternative in {@link TafexprParser#index_expression}.
	 * @param ctx the parse tree
	 */
	void enterIndexExpression(TafexprParser.IndexExpressionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IndexExpression}
	 * labeled alternative in {@link TafexprParser#index_expression}.
	 * @param ctx the parse tree
	 */
	void exitIndexExpression(TafexprParser.IndexExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#parenthesisExpression}.
	 * @param ctx the parse tree
	 */
	void enterParenthesisExpression(TafexprParser.ParenthesisExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#parenthesisExpression}.
	 * @param ctx the parse tree
	 */
	void exitParenthesisExpression(TafexprParser.ParenthesisExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#any}.
	 * @param ctx the parse tree
	 */
	void enterAny(TafexprParser.AnyContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#any}.
	 * @param ctx the parse tree
	 */
	void exitAny(TafexprParser.AnyContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleObject}
	 * labeled alternative in {@link TafexprParser#json}.
	 * @param ctx the parse tree
	 */
	void enterHandleObject(TafexprParser.HandleObjectContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleObject}
	 * labeled alternative in {@link TafexprParser#json}.
	 * @param ctx the parse tree
	 */
	void exitHandleObject(TafexprParser.HandleObjectContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleArray}
	 * labeled alternative in {@link TafexprParser#json}.
	 * @param ctx the parse tree
	 */
	void enterHandleArray(TafexprParser.HandleArrayContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleArray}
	 * labeled alternative in {@link TafexprParser#json}.
	 * @param ctx the parse tree
	 */
	void exitHandleArray(TafexprParser.HandleArrayContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleObjectData}
	 * labeled alternative in {@link TafexprParser#obj}.
	 * @param ctx the parse tree
	 */
	void enterHandleObjectData(TafexprParser.HandleObjectDataContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleObjectData}
	 * labeled alternative in {@link TafexprParser#obj}.
	 * @param ctx the parse tree
	 */
	void exitHandleObjectData(TafexprParser.HandleObjectDataContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleEmptyObjectData}
	 * labeled alternative in {@link TafexprParser#obj}.
	 * @param ctx the parse tree
	 */
	void enterHandleEmptyObjectData(TafexprParser.HandleEmptyObjectDataContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleEmptyObjectData}
	 * labeled alternative in {@link TafexprParser#obj}.
	 * @param ctx the parse tree
	 */
	void exitHandleEmptyObjectData(TafexprParser.HandleEmptyObjectDataContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleObjectPair}
	 * labeled alternative in {@link TafexprParser#pair}.
	 * @param ctx the parse tree
	 */
	void enterHandleObjectPair(TafexprParser.HandleObjectPairContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleObjectPair}
	 * labeled alternative in {@link TafexprParser#pair}.
	 * @param ctx the parse tree
	 */
	void exitHandleObjectPair(TafexprParser.HandleObjectPairContext ctx);
	/**
	 * Enter a parse tree produced by {@link TafexprParser#arr}.
	 * @param ctx the parse tree
	 */
	void enterArr(TafexprParser.ArrContext ctx);
	/**
	 * Exit a parse tree produced by {@link TafexprParser#arr}.
	 * @param ctx the parse tree
	 */
	void exitArr(TafexprParser.ArrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleJJ}
	 * labeled alternative in {@link TafexprParser#value}.
	 * @param ctx the parse tree
	 */
	void enterHandleJJ(TafexprParser.HandleJJContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleJJ}
	 * labeled alternative in {@link TafexprParser#value}.
	 * @param ctx the parse tree
	 */
	void exitHandleJJ(TafexprParser.HandleJJContext ctx);
	/**
	 * Enter a parse tree produced by the {@code HandleFoo}
	 * labeled alternative in {@link TafexprParser#value}.
	 * @param ctx the parse tree
	 */
	void enterHandleFoo(TafexprParser.HandleFooContext ctx);
	/**
	 * Exit a parse tree produced by the {@code HandleFoo}
	 * labeled alternative in {@link TafexprParser#value}.
	 * @param ctx the parse tree
	 */
	void exitHandleFoo(TafexprParser.HandleFooContext ctx);
}