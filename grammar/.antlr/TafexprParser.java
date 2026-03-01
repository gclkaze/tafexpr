// Generated from c:/Users/gclka/Eva Lang/cmd/evalang/tafexpr/grammar/Tafexpr.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class TafexprParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, MUL=28, DIV=29, MOD=30, ADD=31, SUB=32, 
		INTEGER=33, DOUBLE=34, WHITESPACE=35, LBR=36, RBR=37, CON=38, NULL_TOKEN=39, 
		LESSER_THAN=40, LESSER_THAN_EQUAL=41, EQUAL=42, UNEQUAL=43, GREATER_THAN=44, 
		GREATER_THAN_EQUAL=45, LOGICAL_AND=46, LOGICAL_OR=47, LOGICAL_NOT=48, 
		DOLLAR=49, STRING=50, UnterminatedStringLiteral=51, BOOLEAN=52, NUMBER=53, 
		VARIABLE_NAME=54, PROP=55, EXCEPT_QUOTE=56, JSON_NUMBER=57, WS=58, UNKNOWN=59;
	public static final int
		RULE_taf_expression = 0, RULE_libfunc = 1, RULE_expression = 2, RULE_var_expression = 3, 
		RULE_indx_expr = 4, RULE_var_path = 5, RULE_jsonpath_expr = 6, RULE_identifierWithQualifier = 7, 
		RULE_index_expression = 8, RULE_parenthesisExpression = 9, RULE_any = 10, 
		RULE_json = 11, RULE_obj = 12, RULE_pair = 13, RULE_arr = 14, RULE_value = 15;
	private static String[] makeRuleNames() {
		return new String[] {
			"taf_expression", "libfunc", "expression", "var_expression", "indx_expr", 
			"var_path", "jsonpath_expr", "identifierWithQualifier", "index_expression", 
			"parenthesisExpression", "any", "json", "obj", "pair", "arr", "value"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'randomDoubleInRange'", "'('", "','", "')'", "'length'", "'findOneByXPATH'", 
			"'findOneStringByXPATH'", "'findOneDoubleByXPATH'", "'findOneIntegerByXPATH'", 
			"'findOneBooleanByXPATH'", "'findByXPATH'", "'extractOneByREGEX'", "'replaceAllStringOccurrences'", 
			"'toString'", "'toBoolean'", "'toInteger'", "'toDouble'", "'containsString'", 
			"'startsWith'", "'endsWith'", "'trimLeft'", "'trimRight'", "'trim'", 
			"':'", "'{'", "'}'", "'\"'", "'*'", "'/'", "'%'", "'+'", "'-'", null, 
			null, null, "'['", "']'", "'.'", "'null'", "'<'", "'<='", "'=='", "'!='", 
			"'>'", "'>='", "'&&'", "'||'", "'!'", "'$'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "MUL", "DIV", "MOD", "ADD", "SUB", "INTEGER", 
			"DOUBLE", "WHITESPACE", "LBR", "RBR", "CON", "NULL_TOKEN", "LESSER_THAN", 
			"LESSER_THAN_EQUAL", "EQUAL", "UNEQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL", 
			"LOGICAL_AND", "LOGICAL_OR", "LOGICAL_NOT", "DOLLAR", "STRING", "UnterminatedStringLiteral", 
			"BOOLEAN", "NUMBER", "VARIABLE_NAME", "PROP", "EXCEPT_QUOTE", "JSON_NUMBER", 
			"WS", "UNKNOWN"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Tafexpr.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public TafexprParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Taf_expressionContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(TafexprParser.EOF, 0); }
		public LibfuncContext libfunc() {
			return getRuleContext(LibfuncContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Taf_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_taf_expression; }
	}

	public final Taf_expressionContext taf_expression() throws RecognitionException {
		Taf_expressionContext _localctx = new Taf_expressionContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_taf_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(34);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(32);
				libfunc();
				}
				break;
			case 2:
				{
				setState(33);
				expression(0);
				}
				break;
			}
			setState(36);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LibfuncContext extends ParserRuleContext {
		public LibfuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_libfunc; }
	 
		public LibfuncContext() { }
		public void copyFrom(LibfuncContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleRandomDoubleInRangeContext extends LibfuncContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public HandleRandomDoubleInRangeContext(LibfuncContext ctx) { copyFrom(ctx); }
	}

	public final LibfuncContext libfunc() throws RecognitionException {
		LibfuncContext _localctx = new LibfuncContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_libfunc);
		try {
			_localctx = new HandleRandomDoubleInRangeContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			match(T__0);
			setState(39);
			match(T__1);
			setState(40);
			expression(0);
			setState(41);
			match(T__2);
			setState(42);
			expression(0);
			setState(43);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleFindOneByXPATHContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleFindOneByXPATHContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleTrimContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleTrimContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleVarExpressionContext extends ExpressionContext {
		public Var_expressionContext var_expression() {
			return getRuleContext(Var_expressionContext.class,0);
		}
		public HandleVarExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleFindOneDoubleByXPATHContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleFindOneDoubleByXPATHContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MulDivContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MUL() { return getToken(TafexprParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(TafexprParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(TafexprParser.MOD, 0); }
		public MulDivContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleFindOneStringByXPATHContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleFindOneStringByXPATHContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleToStringContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleToStringContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleLibfuncContext extends ExpressionContext {
		public LibfuncContext libfunc() {
			return getRuleContext(LibfuncContext.class,0);
		}
		public HandleLibfuncContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleFindByXPATHContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleFindByXPATHContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleStringContext extends ExpressionContext {
		public TerminalNode STRING() { return getToken(TafexprParser.STRING, 0); }
		public HandleStringContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleExtractOneByREGEXContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleExtractOneByREGEXContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleBoolContext extends ExpressionContext {
		public TerminalNode BOOLEAN() { return getToken(TafexprParser.BOOLEAN, 0); }
		public HandleBoolContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NumberContext extends ExpressionContext {
		public TerminalNode INTEGER() { return getToken(TafexprParser.INTEGER, 0); }
		public NumberContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleJsonContext extends ExpressionContext {
		public JsonContext json() {
			return getRuleContext(JsonContext.class,0);
		}
		public HandleJsonContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleLogicalContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode LOGICAL_AND() { return getToken(TafexprParser.LOGICAL_AND, 0); }
		public TerminalNode LOGICAL_OR() { return getToken(TafexprParser.LOGICAL_OR, 0); }
		public HandleLogicalContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleToBooleanContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleToBooleanContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleTrimLeftContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleTrimLeftContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleFindOneBooleanByXPATHContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleFindOneBooleanByXPATHContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleLogicalNegationContext extends ExpressionContext {
		public TerminalNode LOGICAL_NOT() { return getToken(TafexprParser.LOGICAL_NOT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public HandleLogicalNegationContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleLengthContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleLengthContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleNegationContext extends ExpressionContext {
		public TerminalNode SUB() { return getToken(TafexprParser.SUB, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public HandleNegationContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode ADD() { return getToken(TafexprParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(TafexprParser.SUB, 0); }
		public AddSubContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleFindOneIntegerByXPATHContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleFindOneIntegerByXPATHContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleNullContext extends ExpressionContext {
		public TerminalNode NULL_TOKEN() { return getToken(TafexprParser.NULL_TOKEN, 0); }
		public HandleNullContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleToDoubleContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleToDoubleContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleEndsWithContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleEndsWithContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleContainsStringContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleContainsStringContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OrderedEvaluationContext extends ExpressionContext {
		public ParenthesisExpressionContext parenthesisExpression() {
			return getRuleContext(ParenthesisExpressionContext.class,0);
		}
		public OrderedEvaluationContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LogicalOperationContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode LESSER_THAN() { return getToken(TafexprParser.LESSER_THAN, 0); }
		public TerminalNode LESSER_THAN_EQUAL() { return getToken(TafexprParser.LESSER_THAN_EQUAL, 0); }
		public TerminalNode EQUAL() { return getToken(TafexprParser.EQUAL, 0); }
		public TerminalNode GREATER_THAN() { return getToken(TafexprParser.GREATER_THAN, 0); }
		public TerminalNode GREATER_THAN_EQUAL() { return getToken(TafexprParser.GREATER_THAN_EQUAL, 0); }
		public TerminalNode UNEQUAL() { return getToken(TafexprParser.UNEQUAL, 0); }
		public LogicalOperationContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleToIntegerContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleToIntegerContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleTrimRightContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleTrimRightContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleStartsWithContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleStartsWithContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DoubleValueContext extends ExpressionContext {
		public TerminalNode DOUBLE() { return getToken(TafexprParser.DOUBLE, 0); }
		public DoubleValueContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleReplaceAllStringOccurrencesContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public HandleReplaceAllStringOccurrencesContext(ExpressionContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				_localctx = new HandleNegationContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(46);
				match(SUB);
				setState(47);
				expression(34);
				}
				break;
			case LOGICAL_NOT:
				{
				_localctx = new HandleLogicalNegationContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(48);
				match(LOGICAL_NOT);
				setState(49);
				expression(33);
				}
				break;
			case T__0:
				{
				_localctx = new HandleLibfuncContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(50);
				libfunc();
				}
				break;
			case INTEGER:
				{
				_localctx = new NumberContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(51);
				match(INTEGER);
				}
				break;
			case DOUBLE:
				{
				_localctx = new DoubleValueContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(52);
				match(DOUBLE);
				}
				break;
			case T__1:
				{
				_localctx = new OrderedEvaluationContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(53);
				parenthesisExpression();
				}
				break;
			case VARIABLE_NAME:
				{
				_localctx = new HandleVarExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(54);
				var_expression();
				}
				break;
			case BOOLEAN:
				{
				_localctx = new HandleBoolContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(55);
				match(BOOLEAN);
				}
				break;
			case NULL_TOKEN:
				{
				_localctx = new HandleNullContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(56);
				match(NULL_TOKEN);
				}
				break;
			case STRING:
				{
				_localctx = new HandleStringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(57);
				match(STRING);
				}
				break;
			case T__24:
			case LBR:
				{
				_localctx = new HandleJsonContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(58);
				json();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(178);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(176);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
					case 1:
						{
						_localctx = new HandleLogicalContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(61);
						if (!(precpred(_ctx, 32))) throw new FailedPredicateException(this, "precpred(_ctx, 32)");
						setState(62);
						((HandleLogicalContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==LOGICAL_AND || _la==LOGICAL_OR) ) {
							((HandleLogicalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(63);
						expression(33);
						}
						break;
					case 2:
						{
						_localctx = new MulDivContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(64);
						if (!(precpred(_ctx, 31))) throw new FailedPredicateException(this, "precpred(_ctx, 31)");
						setState(65);
						((MulDivContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1879048192L) != 0)) ) {
							((MulDivContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(66);
						expression(32);
						}
						break;
					case 3:
						{
						_localctx = new AddSubContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(67);
						if (!(precpred(_ctx, 30))) throw new FailedPredicateException(this, "precpred(_ctx, 30)");
						setState(68);
						((AddSubContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((AddSubContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(69);
						expression(31);
						}
						break;
					case 4:
						{
						_localctx = new LogicalOperationContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(70);
						if (!(precpred(_ctx, 29))) throw new FailedPredicateException(this, "precpred(_ctx, 29)");
						setState(71);
						((LogicalOperationContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 69269232549888L) != 0)) ) {
							((LogicalOperationContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(72);
						expression(30);
						}
						break;
					case 5:
						{
						_localctx = new HandleLengthContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(73);
						if (!(precpred(_ctx, 28))) throw new FailedPredicateException(this, "precpred(_ctx, 28)");
						setState(74);
						match(CON);
						setState(75);
						match(T__4);
						}
						break;
					case 6:
						{
						_localctx = new HandleFindOneByXPATHContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(76);
						if (!(precpred(_ctx, 27))) throw new FailedPredicateException(this, "precpred(_ctx, 27)");
						setState(77);
						match(CON);
						setState(78);
						match(T__5);
						setState(79);
						match(T__1);
						setState(80);
						expression(0);
						setState(81);
						match(T__3);
						}
						break;
					case 7:
						{
						_localctx = new HandleFindOneStringByXPATHContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(83);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(84);
						match(CON);
						setState(85);
						match(T__6);
						setState(86);
						match(T__1);
						setState(87);
						expression(0);
						setState(88);
						match(T__3);
						}
						break;
					case 8:
						{
						_localctx = new HandleFindOneDoubleByXPATHContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(90);
						if (!(precpred(_ctx, 25))) throw new FailedPredicateException(this, "precpred(_ctx, 25)");
						setState(91);
						match(CON);
						setState(92);
						match(T__7);
						setState(93);
						match(T__1);
						setState(94);
						expression(0);
						setState(95);
						match(T__3);
						}
						break;
					case 9:
						{
						_localctx = new HandleFindOneIntegerByXPATHContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(97);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(98);
						match(CON);
						setState(99);
						match(T__8);
						setState(100);
						match(T__1);
						setState(101);
						expression(0);
						setState(102);
						match(T__3);
						}
						break;
					case 10:
						{
						_localctx = new HandleFindOneBooleanByXPATHContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(104);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(105);
						match(CON);
						setState(106);
						match(T__9);
						setState(107);
						match(T__1);
						setState(108);
						expression(0);
						setState(109);
						match(T__3);
						}
						break;
					case 11:
						{
						_localctx = new HandleFindByXPATHContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(111);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(112);
						match(CON);
						setState(113);
						match(T__10);
						setState(114);
						match(T__1);
						setState(115);
						expression(0);
						setState(116);
						match(T__3);
						}
						break;
					case 12:
						{
						_localctx = new HandleExtractOneByREGEXContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(118);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(119);
						match(CON);
						setState(120);
						match(T__11);
						setState(121);
						match(T__1);
						setState(122);
						expression(0);
						setState(123);
						match(T__3);
						}
						break;
					case 13:
						{
						_localctx = new HandleReplaceAllStringOccurrencesContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(125);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(126);
						match(CON);
						setState(127);
						match(T__12);
						setState(128);
						match(T__1);
						setState(129);
						expression(0);
						setState(130);
						match(T__2);
						setState(131);
						expression(0);
						setState(132);
						match(T__3);
						}
						break;
					case 14:
						{
						_localctx = new HandleToStringContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(134);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(135);
						match(CON);
						setState(136);
						match(T__13);
						}
						break;
					case 15:
						{
						_localctx = new HandleToBooleanContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(137);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(138);
						match(CON);
						setState(139);
						match(T__14);
						}
						break;
					case 16:
						{
						_localctx = new HandleToIntegerContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(140);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(141);
						match(CON);
						setState(142);
						match(T__15);
						}
						break;
					case 17:
						{
						_localctx = new HandleToDoubleContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(143);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(144);
						match(CON);
						setState(145);
						match(T__16);
						}
						break;
					case 18:
						{
						_localctx = new HandleContainsStringContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(146);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(147);
						match(CON);
						setState(148);
						match(T__17);
						setState(149);
						match(T__1);
						setState(150);
						expression(0);
						setState(151);
						match(T__3);
						}
						break;
					case 19:
						{
						_localctx = new HandleStartsWithContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(153);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(154);
						match(CON);
						setState(155);
						match(T__18);
						setState(156);
						match(T__1);
						setState(157);
						expression(0);
						setState(158);
						match(T__3);
						}
						break;
					case 20:
						{
						_localctx = new HandleEndsWithContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(160);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(161);
						match(CON);
						setState(162);
						match(T__19);
						setState(163);
						match(T__1);
						setState(164);
						expression(0);
						setState(165);
						match(T__3);
						}
						break;
					case 21:
						{
						_localctx = new HandleTrimLeftContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(167);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(168);
						match(CON);
						setState(169);
						match(T__20);
						}
						break;
					case 22:
						{
						_localctx = new HandleTrimRightContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(170);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(171);
						match(CON);
						setState(172);
						match(T__21);
						}
						break;
					case 23:
						{
						_localctx = new HandleTrimContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(173);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(174);
						match(CON);
						setState(175);
						match(T__22);
						}
						break;
					}
					} 
				}
				setState(180);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Var_expressionContext extends ParserRuleContext {
		public TerminalNode VARIABLE_NAME() { return getToken(TafexprParser.VARIABLE_NAME, 0); }
		public Indx_exprContext indx_expr() {
			return getRuleContext(Indx_exprContext.class,0);
		}
		public Var_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_expression; }
	}

	public final Var_expressionContext var_expression() throws RecognitionException {
		Var_expressionContext _localctx = new Var_expressionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_var_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(181);
			match(VARIABLE_NAME);
			setState(182);
			indx_expr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Indx_exprContext extends ParserRuleContext {
		public List<TerminalNode> LBR() { return getTokens(TafexprParser.LBR); }
		public TerminalNode LBR(int i) {
			return getToken(TafexprParser.LBR, i);
		}
		public List<Index_expressionContext> index_expression() {
			return getRuleContexts(Index_expressionContext.class);
		}
		public Index_expressionContext index_expression(int i) {
			return getRuleContext(Index_expressionContext.class,i);
		}
		public List<TerminalNode> RBR() { return getTokens(TafexprParser.RBR); }
		public TerminalNode RBR(int i) {
			return getToken(TafexprParser.RBR, i);
		}
		public TerminalNode CON() { return getToken(TafexprParser.CON, 0); }
		public Var_pathContext var_path() {
			return getRuleContext(Var_pathContext.class,0);
		}
		public Indx_exprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_indx_expr; }
	}

	public final Indx_exprContext indx_expr() throws RecognitionException {
		Indx_exprContext _localctx = new Indx_exprContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_indx_expr);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(184);
				match(LBR);
				setState(185);
				index_expression();
				setState(186);
				match(RBR);
				setState(193);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(187);
						match(LBR);
						setState(188);
						index_expression();
						setState(189);
						match(RBR);
						}
						} 
					}
					setState(195);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
				}
				}
				break;
			}
			setState(200);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(198);
				match(CON);
				setState(199);
				var_path();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Var_pathContext extends ParserRuleContext {
		public List<Jsonpath_exprContext> jsonpath_expr() {
			return getRuleContexts(Jsonpath_exprContext.class);
		}
		public Jsonpath_exprContext jsonpath_expr(int i) {
			return getRuleContext(Jsonpath_exprContext.class,i);
		}
		public List<TerminalNode> CON() { return getTokens(TafexprParser.CON); }
		public TerminalNode CON(int i) {
			return getToken(TafexprParser.CON, i);
		}
		public Var_pathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_path; }
	}

	public final Var_pathContext var_path() throws RecognitionException {
		Var_pathContext _localctx = new Var_pathContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_var_path);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			jsonpath_expr();
			setState(207);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(203);
					match(CON);
					setState(204);
					jsonpath_expr();
					}
					} 
				}
				setState(209);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Jsonpath_exprContext extends ParserRuleContext {
		public IdentifierWithQualifierContext identifierWithQualifier() {
			return getRuleContext(IdentifierWithQualifierContext.class,0);
		}
		public TerminalNode PROP() { return getToken(TafexprParser.PROP, 0); }
		public Jsonpath_exprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jsonpath_expr; }
	}

	public final Jsonpath_exprContext jsonpath_expr() throws RecognitionException {
		Jsonpath_exprContext _localctx = new Jsonpath_exprContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_jsonpath_expr);
		try {
			setState(212);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(210);
				identifierWithQualifier();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(211);
				match(PROP);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierWithQualifierContext extends ParserRuleContext {
		public TerminalNode PROP() { return getToken(TafexprParser.PROP, 0); }
		public List<TerminalNode> LBR() { return getTokens(TafexprParser.LBR); }
		public TerminalNode LBR(int i) {
			return getToken(TafexprParser.LBR, i);
		}
		public List<Index_expressionContext> index_expression() {
			return getRuleContexts(Index_expressionContext.class);
		}
		public Index_expressionContext index_expression(int i) {
			return getRuleContext(Index_expressionContext.class,i);
		}
		public List<TerminalNode> RBR() { return getTokens(TafexprParser.RBR); }
		public TerminalNode RBR(int i) {
			return getToken(TafexprParser.RBR, i);
		}
		public IdentifierWithQualifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierWithQualifier; }
	}

	public final IdentifierWithQualifierContext identifierWithQualifier() throws RecognitionException {
		IdentifierWithQualifierContext _localctx = new IdentifierWithQualifierContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_identifierWithQualifier);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(PROP);
			setState(215);
			match(LBR);
			setState(216);
			index_expression();
			setState(217);
			match(RBR);
			setState(224);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(218);
					match(LBR);
					setState(219);
					index_expression();
					setState(220);
					match(RBR);
					}
					} 
				}
				setState(226);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Index_expressionContext extends ParserRuleContext {
		public Index_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_index_expression; }
	 
		public Index_expressionContext() { }
		public void copyFrom(Index_expressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexExpressionContext extends Index_expressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public IndexExpressionContext(Index_expressionContext ctx) { copyFrom(ctx); }
	}

	public final Index_expressionContext index_expression() throws RecognitionException {
		Index_expressionContext _localctx = new Index_expressionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_index_expression);
		try {
			_localctx = new IndexExpressionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(227);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParenthesisExpressionContext extends ParserRuleContext {
		public ParenthesisExpressionContext parenthesisExpression() {
			return getRuleContext(ParenthesisExpressionContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ParenthesisExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parenthesisExpression; }
	}

	public final ParenthesisExpressionContext parenthesisExpression() throws RecognitionException {
		ParenthesisExpressionContext _localctx = new ParenthesisExpressionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_parenthesisExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(229);
			match(T__1);
			setState(232);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(230);
				parenthesisExpression();
				}
				break;
			case 2:
				{
				setState(231);
				expression(0);
				}
				break;
			}
			setState(234);
			match(T__3);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AnyContext extends ParserRuleContext {
		public TerminalNode LBR() { return getToken(TafexprParser.LBR, 0); }
		public TerminalNode RBR() { return getToken(TafexprParser.RBR, 0); }
		public AnyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_any; }
	}

	public final AnyContext any() throws RecognitionException {
		AnyContext _localctx = new AnyContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_any);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 206410088448L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class JsonContext extends ParserRuleContext {
		public JsonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_json; }
	 
		public JsonContext() { }
		public void copyFrom(JsonContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleObjectContext extends JsonContext {
		public ObjContext obj() {
			return getRuleContext(ObjContext.class,0);
		}
		public HandleObjectContext(JsonContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleArrayContext extends JsonContext {
		public ArrContext arr() {
			return getRuleContext(ArrContext.class,0);
		}
		public HandleArrayContext(JsonContext ctx) { copyFrom(ctx); }
	}

	public final JsonContext json() throws RecognitionException {
		JsonContext _localctx = new JsonContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_json);
		try {
			setState(240);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__24:
				_localctx = new HandleObjectContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(238);
				obj();
				}
				break;
			case LBR:
				_localctx = new HandleArrayContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(239);
				arr();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ObjContext extends ParserRuleContext {
		public ObjContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_obj; }
	 
		public ObjContext() { }
		public void copyFrom(ObjContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleObjectDataContext extends ObjContext {
		public List<PairContext> pair() {
			return getRuleContexts(PairContext.class);
		}
		public PairContext pair(int i) {
			return getRuleContext(PairContext.class,i);
		}
		public HandleObjectDataContext(ObjContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleEmptyObjectDataContext extends ObjContext {
		public HandleEmptyObjectDataContext(ObjContext ctx) { copyFrom(ctx); }
	}

	public final ObjContext obj() throws RecognitionException {
		ObjContext _localctx = new ObjContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_obj);
		int _la;
		try {
			setState(255);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				_localctx = new HandleObjectDataContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(242);
				match(T__24);
				setState(243);
				pair();
				setState(248);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__2) {
					{
					{
					setState(244);
					match(T__2);
					setState(245);
					pair();
					}
					}
					setState(250);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(251);
				match(T__25);
				}
				break;
			case 2:
				_localctx = new HandleEmptyObjectDataContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(253);
				match(T__24);
				setState(254);
				match(T__25);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PairContext extends ParserRuleContext {
		public PairContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pair; }
	 
		public PairContext() { }
		public void copyFrom(PairContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleObjectPairContext extends PairContext {
		public TerminalNode STRING() { return getToken(TafexprParser.STRING, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public HandleObjectPairContext(PairContext ctx) { copyFrom(ctx); }
	}

	public final PairContext pair() throws RecognitionException {
		PairContext _localctx = new PairContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_pair);
		try {
			_localctx = new HandleObjectPairContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(257);
			match(STRING);
			setState(258);
			match(T__23);
			setState(259);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrContext extends ParserRuleContext {
		public TerminalNode LBR() { return getToken(TafexprParser.LBR, 0); }
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public TerminalNode RBR() { return getToken(TafexprParser.RBR, 0); }
		public ArrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arr; }
	}

	public final ArrContext arr() throws RecognitionException {
		ArrContext _localctx = new ArrContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_arr);
		int _la;
		try {
			setState(274);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(261);
				match(LBR);
				setState(262);
				value();
				setState(267);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__2) {
					{
					{
					setState(263);
					match(T__2);
					setState(264);
					value();
					}
					}
					setState(269);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(270);
				match(RBR);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(272);
				match(LBR);
				setState(273);
				match(RBR);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValueContext extends ParserRuleContext {
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	 
		public ValueContext() { }
		public void copyFrom(ValueContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleJJContext extends ValueContext {
		public JsonContext json() {
			return getRuleContext(JsonContext.class,0);
		}
		public HandleJJContext(ValueContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class HandleFooContext extends ValueContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public HandleFooContext(ValueContext ctx) { copyFrom(ctx); }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_value);
		try {
			setState(278);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new HandleJJContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(276);
				json();
				}
				break;
			case 2:
				_localctx = new HandleFooContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(277);
				expression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 2:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 32);
		case 1:
			return precpred(_ctx, 31);
		case 2:
			return precpred(_ctx, 30);
		case 3:
			return precpred(_ctx, 29);
		case 4:
			return precpred(_ctx, 28);
		case 5:
			return precpred(_ctx, 27);
		case 6:
			return precpred(_ctx, 26);
		case 7:
			return precpred(_ctx, 25);
		case 8:
			return precpred(_ctx, 24);
		case 9:
			return precpred(_ctx, 23);
		case 10:
			return precpred(_ctx, 22);
		case 11:
			return precpred(_ctx, 21);
		case 12:
			return precpred(_ctx, 20);
		case 13:
			return precpred(_ctx, 19);
		case 14:
			return precpred(_ctx, 18);
		case 15:
			return precpred(_ctx, 17);
		case 16:
			return precpred(_ctx, 16);
		case 17:
			return precpred(_ctx, 15);
		case 18:
			return precpred(_ctx, 14);
		case 19:
			return precpred(_ctx, 13);
		case 20:
			return precpred(_ctx, 12);
		case 21:
			return precpred(_ctx, 10);
		case 22:
			return precpred(_ctx, 9);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001;\u0119\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0001\u0000\u0001\u0000\u0003\u0000#\b\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002<\b\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0005\u0002\u00b1\b\u0002\n\u0002\f\u0002\u00b4\t\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0005\u0004\u00c0\b\u0004\n\u0004\f\u0004"+
		"\u00c3\t\u0004\u0003\u0004\u00c5\b\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004\u00c9\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005\u00ce"+
		"\b\u0005\n\u0005\f\u0005\u00d1\t\u0005\u0001\u0006\u0001\u0006\u0003\u0006"+
		"\u00d5\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u00df\b\u0007\n\u0007"+
		"\f\u0007\u00e2\t\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0003\t"+
		"\u00e9\b\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0003"+
		"\u000b\u00f1\b\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0005\f\u00f7\b\f"+
		"\n\f\f\f\u00fa\t\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u0100\b\f\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0005\u000e\u010a\b\u000e\n\u000e\f\u000e\u010d\t\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u0113\b\u000e\u0001\u000f\u0001"+
		"\u000f\u0003\u000f\u0117\b\u000f\u0001\u000f\u0000\u0001\u0004\u0010\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e\u0000\u0005\u0001\u0000./\u0001\u0000\u001c\u001e\u0001\u0000\u001f"+
		" \u0001\u0000(-\u0002\u0000\u0018\u001b$%\u0137\u0000\"\u0001\u0000\u0000"+
		"\u0000\u0002&\u0001\u0000\u0000\u0000\u0004;\u0001\u0000\u0000\u0000\u0006"+
		"\u00b5\u0001\u0000\u0000\u0000\b\u00c4\u0001\u0000\u0000\u0000\n\u00ca"+
		"\u0001\u0000\u0000\u0000\f\u00d4\u0001\u0000\u0000\u0000\u000e\u00d6\u0001"+
		"\u0000\u0000\u0000\u0010\u00e3\u0001\u0000\u0000\u0000\u0012\u00e5\u0001"+
		"\u0000\u0000\u0000\u0014\u00ec\u0001\u0000\u0000\u0000\u0016\u00f0\u0001"+
		"\u0000\u0000\u0000\u0018\u00ff\u0001\u0000\u0000\u0000\u001a\u0101\u0001"+
		"\u0000\u0000\u0000\u001c\u0112\u0001\u0000\u0000\u0000\u001e\u0116\u0001"+
		"\u0000\u0000\u0000 #\u0003\u0002\u0001\u0000!#\u0003\u0004\u0002\u0000"+
		"\" \u0001\u0000\u0000\u0000\"!\u0001\u0000\u0000\u0000#$\u0001\u0000\u0000"+
		"\u0000$%\u0005\u0000\u0000\u0001%\u0001\u0001\u0000\u0000\u0000&\'\u0005"+
		"\u0001\u0000\u0000\'(\u0005\u0002\u0000\u0000()\u0003\u0004\u0002\u0000"+
		")*\u0005\u0003\u0000\u0000*+\u0003\u0004\u0002\u0000+,\u0005\u0004\u0000"+
		"\u0000,\u0003\u0001\u0000\u0000\u0000-.\u0006\u0002\uffff\uffff\u0000"+
		"./\u0005 \u0000\u0000/<\u0003\u0004\u0002\"01\u00050\u0000\u00001<\u0003"+
		"\u0004\u0002!2<\u0003\u0002\u0001\u00003<\u0005!\u0000\u00004<\u0005\""+
		"\u0000\u00005<\u0003\u0012\t\u00006<\u0003\u0006\u0003\u00007<\u00054"+
		"\u0000\u00008<\u0005\'\u0000\u00009<\u00052\u0000\u0000:<\u0003\u0016"+
		"\u000b\u0000;-\u0001\u0000\u0000\u0000;0\u0001\u0000\u0000\u0000;2\u0001"+
		"\u0000\u0000\u0000;3\u0001\u0000\u0000\u0000;4\u0001\u0000\u0000\u0000"+
		";5\u0001\u0000\u0000\u0000;6\u0001\u0000\u0000\u0000;7\u0001\u0000\u0000"+
		"\u0000;8\u0001\u0000\u0000\u0000;9\u0001\u0000\u0000\u0000;:\u0001\u0000"+
		"\u0000\u0000<\u00b2\u0001\u0000\u0000\u0000=>\n \u0000\u0000>?\u0007\u0000"+
		"\u0000\u0000?\u00b1\u0003\u0004\u0002!@A\n\u001f\u0000\u0000AB\u0007\u0001"+
		"\u0000\u0000B\u00b1\u0003\u0004\u0002 CD\n\u001e\u0000\u0000DE\u0007\u0002"+
		"\u0000\u0000E\u00b1\u0003\u0004\u0002\u001fFG\n\u001d\u0000\u0000GH\u0007"+
		"\u0003\u0000\u0000H\u00b1\u0003\u0004\u0002\u001eIJ\n\u001c\u0000\u0000"+
		"JK\u0005&\u0000\u0000K\u00b1\u0005\u0005\u0000\u0000LM\n\u001b\u0000\u0000"+
		"MN\u0005&\u0000\u0000NO\u0005\u0006\u0000\u0000OP\u0005\u0002\u0000\u0000"+
		"PQ\u0003\u0004\u0002\u0000QR\u0005\u0004\u0000\u0000R\u00b1\u0001\u0000"+
		"\u0000\u0000ST\n\u001a\u0000\u0000TU\u0005&\u0000\u0000UV\u0005\u0007"+
		"\u0000\u0000VW\u0005\u0002\u0000\u0000WX\u0003\u0004\u0002\u0000XY\u0005"+
		"\u0004\u0000\u0000Y\u00b1\u0001\u0000\u0000\u0000Z[\n\u0019\u0000\u0000"+
		"[\\\u0005&\u0000\u0000\\]\u0005\b\u0000\u0000]^\u0005\u0002\u0000\u0000"+
		"^_\u0003\u0004\u0002\u0000_`\u0005\u0004\u0000\u0000`\u00b1\u0001\u0000"+
		"\u0000\u0000ab\n\u0018\u0000\u0000bc\u0005&\u0000\u0000cd\u0005\t\u0000"+
		"\u0000de\u0005\u0002\u0000\u0000ef\u0003\u0004\u0002\u0000fg\u0005\u0004"+
		"\u0000\u0000g\u00b1\u0001\u0000\u0000\u0000hi\n\u0017\u0000\u0000ij\u0005"+
		"&\u0000\u0000jk\u0005\n\u0000\u0000kl\u0005\u0002\u0000\u0000lm\u0003"+
		"\u0004\u0002\u0000mn\u0005\u0004\u0000\u0000n\u00b1\u0001\u0000\u0000"+
		"\u0000op\n\u0016\u0000\u0000pq\u0005&\u0000\u0000qr\u0005\u000b\u0000"+
		"\u0000rs\u0005\u0002\u0000\u0000st\u0003\u0004\u0002\u0000tu\u0005\u0004"+
		"\u0000\u0000u\u00b1\u0001\u0000\u0000\u0000vw\n\u0015\u0000\u0000wx\u0005"+
		"&\u0000\u0000xy\u0005\f\u0000\u0000yz\u0005\u0002\u0000\u0000z{\u0003"+
		"\u0004\u0002\u0000{|\u0005\u0004\u0000\u0000|\u00b1\u0001\u0000\u0000"+
		"\u0000}~\n\u0014\u0000\u0000~\u007f\u0005&\u0000\u0000\u007f\u0080\u0005"+
		"\r\u0000\u0000\u0080\u0081\u0005\u0002\u0000\u0000\u0081\u0082\u0003\u0004"+
		"\u0002\u0000\u0082\u0083\u0005\u0003\u0000\u0000\u0083\u0084\u0003\u0004"+
		"\u0002\u0000\u0084\u0085\u0005\u0004\u0000\u0000\u0085\u00b1\u0001\u0000"+
		"\u0000\u0000\u0086\u0087\n\u0013\u0000\u0000\u0087\u0088\u0005&\u0000"+
		"\u0000\u0088\u00b1\u0005\u000e\u0000\u0000\u0089\u008a\n\u0012\u0000\u0000"+
		"\u008a\u008b\u0005&\u0000\u0000\u008b\u00b1\u0005\u000f\u0000\u0000\u008c"+
		"\u008d\n\u0011\u0000\u0000\u008d\u008e\u0005&\u0000\u0000\u008e\u00b1"+
		"\u0005\u0010\u0000\u0000\u008f\u0090\n\u0010\u0000\u0000\u0090\u0091\u0005"+
		"&\u0000\u0000\u0091\u00b1\u0005\u0011\u0000\u0000\u0092\u0093\n\u000f"+
		"\u0000\u0000\u0093\u0094\u0005&\u0000\u0000\u0094\u0095\u0005\u0012\u0000"+
		"\u0000\u0095\u0096\u0005\u0002\u0000\u0000\u0096\u0097\u0003\u0004\u0002"+
		"\u0000\u0097\u0098\u0005\u0004\u0000\u0000\u0098\u00b1\u0001\u0000\u0000"+
		"\u0000\u0099\u009a\n\u000e\u0000\u0000\u009a\u009b\u0005&\u0000\u0000"+
		"\u009b\u009c\u0005\u0013\u0000\u0000\u009c\u009d\u0005\u0002\u0000\u0000"+
		"\u009d\u009e\u0003\u0004\u0002\u0000\u009e\u009f\u0005\u0004\u0000\u0000"+
		"\u009f\u00b1\u0001\u0000\u0000\u0000\u00a0\u00a1\n\r\u0000\u0000\u00a1"+
		"\u00a2\u0005&\u0000\u0000\u00a2\u00a3\u0005\u0014\u0000\u0000\u00a3\u00a4"+
		"\u0005\u0002\u0000\u0000\u00a4\u00a5\u0003\u0004\u0002\u0000\u00a5\u00a6"+
		"\u0005\u0004\u0000\u0000\u00a6\u00b1\u0001\u0000\u0000\u0000\u00a7\u00a8"+
		"\n\f\u0000\u0000\u00a8\u00a9\u0005&\u0000\u0000\u00a9\u00b1\u0005\u0015"+
		"\u0000\u0000\u00aa\u00ab\n\n\u0000\u0000\u00ab\u00ac\u0005&\u0000\u0000"+
		"\u00ac\u00b1\u0005\u0016\u0000\u0000\u00ad\u00ae\n\t\u0000\u0000\u00ae"+
		"\u00af\u0005&\u0000\u0000\u00af\u00b1\u0005\u0017\u0000\u0000\u00b0=\u0001"+
		"\u0000\u0000\u0000\u00b0@\u0001\u0000\u0000\u0000\u00b0C\u0001\u0000\u0000"+
		"\u0000\u00b0F\u0001\u0000\u0000\u0000\u00b0I\u0001\u0000\u0000\u0000\u00b0"+
		"L\u0001\u0000\u0000\u0000\u00b0S\u0001\u0000\u0000\u0000\u00b0Z\u0001"+
		"\u0000\u0000\u0000\u00b0a\u0001\u0000\u0000\u0000\u00b0h\u0001\u0000\u0000"+
		"\u0000\u00b0o\u0001\u0000\u0000\u0000\u00b0v\u0001\u0000\u0000\u0000\u00b0"+
		"}\u0001\u0000\u0000\u0000\u00b0\u0086\u0001\u0000\u0000\u0000\u00b0\u0089"+
		"\u0001\u0000\u0000\u0000\u00b0\u008c\u0001\u0000\u0000\u0000\u00b0\u008f"+
		"\u0001\u0000\u0000\u0000\u00b0\u0092\u0001\u0000\u0000\u0000\u00b0\u0099"+
		"\u0001\u0000\u0000\u0000\u00b0\u00a0\u0001\u0000\u0000\u0000\u00b0\u00a7"+
		"\u0001\u0000\u0000\u0000\u00b0\u00aa\u0001\u0000\u0000\u0000\u00b0\u00ad"+
		"\u0001\u0000\u0000\u0000\u00b1\u00b4\u0001\u0000\u0000\u0000\u00b2\u00b0"+
		"\u0001\u0000\u0000\u0000\u00b2\u00b3\u0001\u0000\u0000\u0000\u00b3\u0005"+
		"\u0001\u0000\u0000\u0000\u00b4\u00b2\u0001\u0000\u0000\u0000\u00b5\u00b6"+
		"\u00056\u0000\u0000\u00b6\u00b7\u0003\b\u0004\u0000\u00b7\u0007\u0001"+
		"\u0000\u0000\u0000\u00b8\u00b9\u0005$\u0000\u0000\u00b9\u00ba\u0003\u0010"+
		"\b\u0000\u00ba\u00c1\u0005%\u0000\u0000\u00bb\u00bc\u0005$\u0000\u0000"+
		"\u00bc\u00bd\u0003\u0010\b\u0000\u00bd\u00be\u0005%\u0000\u0000\u00be"+
		"\u00c0\u0001\u0000\u0000\u0000\u00bf\u00bb\u0001\u0000\u0000\u0000\u00c0"+
		"\u00c3\u0001\u0000\u0000\u0000\u00c1\u00bf\u0001\u0000\u0000\u0000\u00c1"+
		"\u00c2\u0001\u0000\u0000\u0000\u00c2\u00c5\u0001\u0000\u0000\u0000\u00c3"+
		"\u00c1\u0001\u0000\u0000\u0000\u00c4\u00b8\u0001\u0000\u0000\u0000\u00c4"+
		"\u00c5\u0001\u0000\u0000\u0000\u00c5\u00c8\u0001\u0000\u0000\u0000\u00c6"+
		"\u00c7\u0005&\u0000\u0000\u00c7\u00c9\u0003\n\u0005\u0000\u00c8\u00c6"+
		"\u0001\u0000\u0000\u0000\u00c8\u00c9\u0001\u0000\u0000\u0000\u00c9\t\u0001"+
		"\u0000\u0000\u0000\u00ca\u00cf\u0003\f\u0006\u0000\u00cb\u00cc\u0005&"+
		"\u0000\u0000\u00cc\u00ce\u0003\f\u0006\u0000\u00cd\u00cb\u0001\u0000\u0000"+
		"\u0000\u00ce\u00d1\u0001\u0000\u0000\u0000\u00cf\u00cd\u0001\u0000\u0000"+
		"\u0000\u00cf\u00d0\u0001\u0000\u0000\u0000\u00d0\u000b\u0001\u0000\u0000"+
		"\u0000\u00d1\u00cf\u0001\u0000\u0000\u0000\u00d2\u00d5\u0003\u000e\u0007"+
		"\u0000\u00d3\u00d5\u00057\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000"+
		"\u00d4\u00d3\u0001\u0000\u0000\u0000\u00d5\r\u0001\u0000\u0000\u0000\u00d6"+
		"\u00d7\u00057\u0000\u0000\u00d7\u00d8\u0005$\u0000\u0000\u00d8\u00d9\u0003"+
		"\u0010\b\u0000\u00d9\u00e0\u0005%\u0000\u0000\u00da\u00db\u0005$\u0000"+
		"\u0000\u00db\u00dc\u0003\u0010\b\u0000\u00dc\u00dd\u0005%\u0000\u0000"+
		"\u00dd\u00df\u0001\u0000\u0000\u0000\u00de\u00da\u0001\u0000\u0000\u0000"+
		"\u00df\u00e2\u0001\u0000\u0000\u0000\u00e0\u00de\u0001\u0000\u0000\u0000"+
		"\u00e0\u00e1\u0001\u0000\u0000\u0000\u00e1\u000f\u0001\u0000\u0000\u0000"+
		"\u00e2\u00e0\u0001\u0000\u0000\u0000\u00e3\u00e4\u0003\u0004\u0002\u0000"+
		"\u00e4\u0011\u0001\u0000\u0000\u0000\u00e5\u00e8\u0005\u0002\u0000\u0000"+
		"\u00e6\u00e9\u0003\u0012\t\u0000\u00e7\u00e9\u0003\u0004\u0002\u0000\u00e8"+
		"\u00e6\u0001\u0000\u0000\u0000\u00e8\u00e7\u0001\u0000\u0000\u0000\u00e9"+
		"\u00ea\u0001\u0000\u0000\u0000\u00ea\u00eb\u0005\u0004\u0000\u0000\u00eb"+
		"\u0013\u0001\u0000\u0000\u0000\u00ec\u00ed\u0007\u0004\u0000\u0000\u00ed"+
		"\u0015\u0001\u0000\u0000\u0000\u00ee\u00f1\u0003\u0018\f\u0000\u00ef\u00f1"+
		"\u0003\u001c\u000e\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f0\u00ef"+
		"\u0001\u0000\u0000\u0000\u00f1\u0017\u0001\u0000\u0000\u0000\u00f2\u00f3"+
		"\u0005\u0019\u0000\u0000\u00f3\u00f8\u0003\u001a\r\u0000\u00f4\u00f5\u0005"+
		"\u0003\u0000\u0000\u00f5\u00f7\u0003\u001a\r\u0000\u00f6\u00f4\u0001\u0000"+
		"\u0000\u0000\u00f7\u00fa\u0001\u0000\u0000\u0000\u00f8\u00f6\u0001\u0000"+
		"\u0000\u0000\u00f8\u00f9\u0001\u0000\u0000\u0000\u00f9\u00fb\u0001\u0000"+
		"\u0000\u0000\u00fa\u00f8\u0001\u0000\u0000\u0000\u00fb\u00fc\u0005\u001a"+
		"\u0000\u0000\u00fc\u0100\u0001\u0000\u0000\u0000\u00fd\u00fe\u0005\u0019"+
		"\u0000\u0000\u00fe\u0100\u0005\u001a\u0000\u0000\u00ff\u00f2\u0001\u0000"+
		"\u0000\u0000\u00ff\u00fd\u0001\u0000\u0000\u0000\u0100\u0019\u0001\u0000"+
		"\u0000\u0000\u0101\u0102\u00052\u0000\u0000\u0102\u0103\u0005\u0018\u0000"+
		"\u0000\u0103\u0104\u0003\u001e\u000f\u0000\u0104\u001b\u0001\u0000\u0000"+
		"\u0000\u0105\u0106\u0005$\u0000\u0000\u0106\u010b\u0003\u001e\u000f\u0000"+
		"\u0107\u0108\u0005\u0003\u0000\u0000\u0108\u010a\u0003\u001e\u000f\u0000"+
		"\u0109\u0107\u0001\u0000\u0000\u0000\u010a\u010d\u0001\u0000\u0000\u0000"+
		"\u010b\u0109\u0001\u0000\u0000\u0000\u010b\u010c\u0001\u0000\u0000\u0000"+
		"\u010c\u010e\u0001\u0000\u0000\u0000\u010d\u010b\u0001\u0000\u0000\u0000"+
		"\u010e\u010f\u0005%\u0000\u0000\u010f\u0113\u0001\u0000\u0000\u0000\u0110"+
		"\u0111\u0005$\u0000\u0000\u0111\u0113\u0005%\u0000\u0000\u0112\u0105\u0001"+
		"\u0000\u0000\u0000\u0112\u0110\u0001\u0000\u0000\u0000\u0113\u001d\u0001"+
		"\u0000\u0000\u0000\u0114\u0117\u0003\u0016\u000b\u0000\u0115\u0117\u0003"+
		"\u0004\u0002\u0000\u0116\u0114\u0001\u0000\u0000\u0000\u0116\u0115\u0001"+
		"\u0000\u0000\u0000\u0117\u001f\u0001\u0000\u0000\u0000\u0011\";\u00b0"+
		"\u00b2\u00c1\u00c4\u00c8\u00cf\u00d4\u00e0\u00e8\u00f0\u00f8\u00ff\u010b"+
		"\u0112\u0116";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}