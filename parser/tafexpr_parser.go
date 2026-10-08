// Code generated from grammar/Tafexpr.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tafexpr
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TafexprParser struct {
	*antlr.BaseParser
}

var TafexprParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tafexprParserInit() {
	staticData := &TafexprParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "','", "')'", "'{'", "'}'", "':'", "'*'", "'/'", "'%'", "'+'",
		"'-'", "", "", "", "'['", "']'", "'.'", "'null'", "'<'", "'<='", "'=='",
		"'!='", "'>'", "'>='", "'&&'", "'||'", "'!'", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "MUL", "DIV", "MOD", "ADD", "SUB", "DOUBLE",
		"INTEGER", "WHITESPACE", "LBR", "RBR", "CON", "NULL_TOKEN", "LESSER_THAN",
		"LESSER_THAN_EQUAL", "EQUAL", "UNEQUAL", "GREATER_THAN", "GREATER_THAN_EQUAL",
		"LOGICAL_AND", "LOGICAL_OR", "LOGICAL_NOT", "DOLLAR", "STRING", "BOOLEAN",
		"NUMBER", "VARIABLE_NAME", "PROP", "JSON_NUMBER", "WS", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"taf_expression", "expression", "var_expression", "indx_expr", "var_path",
		"jsonpath_expr", "identifierWithQualifier", "index_expression", "parenthesisExpression",
		"json", "obj", "pair", "arr", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 195, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 38, 8, 1, 10, 1, 12, 1, 41, 9, 1, 3, 1,
		43, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 58, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 85, 8, 1, 10, 1, 12, 1, 88, 9,
		1, 3, 1, 90, 8, 1, 1, 1, 5, 1, 93, 8, 1, 10, 1, 12, 1, 96, 9, 1, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 108, 8, 3,
		10, 3, 12, 3, 111, 9, 3, 3, 3, 113, 8, 3, 1, 3, 1, 3, 3, 3, 117, 8, 3,
		1, 4, 1, 4, 1, 4, 5, 4, 122, 8, 4, 10, 4, 12, 4, 125, 9, 4, 1, 5, 1, 5,
		3, 5, 129, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6,
		139, 8, 6, 10, 6, 12, 6, 142, 9, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8,
		149, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 155, 8, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 161, 8, 10, 10, 10, 12, 10, 164, 9, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 170, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 5, 12, 180, 8, 12, 10, 12, 12, 12, 183, 9, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 189, 8, 12, 1, 13, 1, 13, 3, 13, 193, 8, 13, 1,
		13, 0, 1, 2, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0,
		4, 1, 0, 7, 9, 1, 0, 10, 11, 2, 0, 19, 20, 23, 24, 1, 0, 21, 22, 214, 0,
		28, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4, 97, 1, 0, 0, 0, 6, 112, 1, 0, 0,
		0, 8, 118, 1, 0, 0, 0, 10, 128, 1, 0, 0, 0, 12, 130, 1, 0, 0, 0, 14, 143,
		1, 0, 0, 0, 16, 145, 1, 0, 0, 0, 18, 154, 1, 0, 0, 0, 20, 169, 1, 0, 0,
		0, 22, 171, 1, 0, 0, 0, 24, 188, 1, 0, 0, 0, 26, 192, 1, 0, 0, 0, 28, 29,
		3, 2, 1, 0, 29, 30, 5, 0, 0, 1, 30, 1, 1, 0, 0, 0, 31, 32, 6, 1, -1, 0,
		32, 33, 5, 33, 0, 0, 33, 42, 5, 1, 0, 0, 34, 39, 3, 2, 1, 0, 35, 36, 5,
		2, 0, 0, 36, 38, 3, 2, 1, 0, 37, 35, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39,
		37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0,
		0, 42, 34, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 58,
		5, 3, 0, 0, 45, 46, 5, 11, 0, 0, 46, 58, 3, 2, 1, 16, 47, 48, 5, 27, 0,
		0, 48, 58, 3, 2, 1, 15, 49, 58, 5, 13, 0, 0, 50, 58, 5, 12, 0, 0, 51, 58,
		3, 16, 8, 0, 52, 58, 3, 4, 2, 0, 53, 58, 5, 30, 0, 0, 54, 58, 5, 18, 0,
		0, 55, 58, 5, 29, 0, 0, 56, 58, 3, 18, 9, 0, 57, 31, 1, 0, 0, 0, 57, 45,
		1, 0, 0, 0, 57, 47, 1, 0, 0, 0, 57, 49, 1, 0, 0, 0, 57, 50, 1, 0, 0, 0,
		57, 51, 1, 0, 0, 0, 57, 52, 1, 0, 0, 0, 57, 53, 1, 0, 0, 0, 57, 54, 1,
		0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 94, 1, 0, 0, 0, 59,
		60, 10, 14, 0, 0, 60, 61, 7, 0, 0, 0, 61, 93, 3, 2, 1, 15, 62, 63, 10,
		13, 0, 0, 63, 64, 7, 1, 0, 0, 64, 93, 3, 2, 1, 14, 65, 66, 10, 12, 0, 0,
		66, 67, 7, 2, 0, 0, 67, 93, 3, 2, 1, 13, 68, 69, 10, 11, 0, 0, 69, 70,
		7, 3, 0, 0, 70, 93, 3, 2, 1, 12, 71, 72, 10, 10, 0, 0, 72, 73, 5, 25, 0,
		0, 73, 93, 3, 2, 1, 11, 74, 75, 10, 9, 0, 0, 75, 76, 5, 26, 0, 0, 76, 93,
		3, 2, 1, 10, 77, 78, 10, 18, 0, 0, 78, 79, 5, 17, 0, 0, 79, 80, 5, 33,
		0, 0, 80, 89, 5, 1, 0, 0, 81, 86, 3, 2, 1, 0, 82, 83, 5, 2, 0, 0, 83, 85,
		3, 2, 1, 0, 84, 82, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 81, 1,
		0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 5, 3, 0, 0, 92,
		59, 1, 0, 0, 0, 92, 62, 1, 0, 0, 0, 92, 65, 1, 0, 0, 0, 92, 68, 1, 0, 0,
		0, 92, 71, 1, 0, 0, 0, 92, 74, 1, 0, 0, 0, 92, 77, 1, 0, 0, 0, 93, 96,
		1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 3, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 97, 98, 5, 32, 0, 0, 98, 99, 3, 6, 3, 0, 99, 5, 1,
		0, 0, 0, 100, 101, 5, 15, 0, 0, 101, 102, 3, 14, 7, 0, 102, 109, 5, 16,
		0, 0, 103, 104, 5, 15, 0, 0, 104, 105, 3, 14, 7, 0, 105, 106, 5, 16, 0,
		0, 106, 108, 1, 0, 0, 0, 107, 103, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109,
		107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 112, 100, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 116, 1, 0,
		0, 0, 114, 115, 5, 17, 0, 0, 115, 117, 3, 8, 4, 0, 116, 114, 1, 0, 0, 0,
		116, 117, 1, 0, 0, 0, 117, 7, 1, 0, 0, 0, 118, 123, 3, 10, 5, 0, 119, 120,
		5, 17, 0, 0, 120, 122, 3, 10, 5, 0, 121, 119, 1, 0, 0, 0, 122, 125, 1,
		0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 9, 1, 0, 0, 0,
		125, 123, 1, 0, 0, 0, 126, 129, 3, 12, 6, 0, 127, 129, 5, 33, 0, 0, 128,
		126, 1, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 11, 1, 0, 0, 0, 130, 131, 5,
		33, 0, 0, 131, 132, 5, 15, 0, 0, 132, 133, 3, 14, 7, 0, 133, 140, 5, 16,
		0, 0, 134, 135, 5, 15, 0, 0, 135, 136, 3, 14, 7, 0, 136, 137, 5, 16, 0,
		0, 137, 139, 1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140,
		138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 13, 1, 0, 0, 0, 142, 140, 1,
		0, 0, 0, 143, 144, 3, 2, 1, 0, 144, 15, 1, 0, 0, 0, 145, 148, 5, 1, 0,
		0, 146, 149, 3, 16, 8, 0, 147, 149, 3, 2, 1, 0, 148, 146, 1, 0, 0, 0, 148,
		147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 5, 3, 0, 0, 151, 17, 1,
		0, 0, 0, 152, 155, 3, 20, 10, 0, 153, 155, 3, 24, 12, 0, 154, 152, 1, 0,
		0, 0, 154, 153, 1, 0, 0, 0, 155, 19, 1, 0, 0, 0, 156, 157, 5, 4, 0, 0,
		157, 162, 3, 22, 11, 0, 158, 159, 5, 2, 0, 0, 159, 161, 3, 22, 11, 0, 160,
		158, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163,
		1, 0, 0, 0, 163, 165, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 166, 5, 5,
		0, 0, 166, 170, 1, 0, 0, 0, 167, 168, 5, 4, 0, 0, 168, 170, 5, 5, 0, 0,
		169, 156, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 21, 1, 0, 0, 0, 171, 172,
		5, 29, 0, 0, 172, 173, 5, 6, 0, 0, 173, 174, 3, 26, 13, 0, 174, 23, 1,
		0, 0, 0, 175, 176, 5, 15, 0, 0, 176, 181, 3, 26, 13, 0, 177, 178, 5, 2,
		0, 0, 178, 180, 3, 26, 13, 0, 179, 177, 1, 0, 0, 0, 180, 183, 1, 0, 0,
		0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184, 1, 0, 0, 0, 183,
		181, 1, 0, 0, 0, 184, 185, 5, 16, 0, 0, 185, 189, 1, 0, 0, 0, 186, 187,
		5, 15, 0, 0, 187, 189, 5, 16, 0, 0, 188, 175, 1, 0, 0, 0, 188, 186, 1,
		0, 0, 0, 189, 25, 1, 0, 0, 0, 190, 193, 3, 18, 9, 0, 191, 193, 3, 2, 1,
		0, 192, 190, 1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 27, 1, 0, 0, 0, 20,
		39, 42, 57, 86, 89, 92, 94, 109, 112, 116, 123, 128, 140, 148, 154, 162,
		169, 181, 188, 192,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TafexprParserInit initializes any static state used to implement TafexprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTafexprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TafexprParserInit() {
	staticData := &TafexprParserStaticData
	staticData.once.Do(tafexprParserInit)
}

// NewTafexprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTafexprParser(input antlr.TokenStream) *TafexprParser {
	TafexprParserInit()
	this := new(TafexprParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TafexprParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tafexpr.g4"

	return this
}

// TafexprParser tokens.
const (
	TafexprParserEOF                = antlr.TokenEOF
	TafexprParserT__0               = 1
	TafexprParserT__1               = 2
	TafexprParserT__2               = 3
	TafexprParserT__3               = 4
	TafexprParserT__4               = 5
	TafexprParserT__5               = 6
	TafexprParserMUL                = 7
	TafexprParserDIV                = 8
	TafexprParserMOD                = 9
	TafexprParserADD                = 10
	TafexprParserSUB                = 11
	TafexprParserDOUBLE             = 12
	TafexprParserINTEGER            = 13
	TafexprParserWHITESPACE         = 14
	TafexprParserLBR                = 15
	TafexprParserRBR                = 16
	TafexprParserCON                = 17
	TafexprParserNULL_TOKEN         = 18
	TafexprParserLESSER_THAN        = 19
	TafexprParserLESSER_THAN_EQUAL  = 20
	TafexprParserEQUAL              = 21
	TafexprParserUNEQUAL            = 22
	TafexprParserGREATER_THAN       = 23
	TafexprParserGREATER_THAN_EQUAL = 24
	TafexprParserLOGICAL_AND        = 25
	TafexprParserLOGICAL_OR         = 26
	TafexprParserLOGICAL_NOT        = 27
	TafexprParserDOLLAR             = 28
	TafexprParserSTRING             = 29
	TafexprParserBOOLEAN            = 30
	TafexprParserNUMBER             = 31
	TafexprParserVARIABLE_NAME      = 32
	TafexprParserPROP               = 33
	TafexprParserJSON_NUMBER        = 34
	TafexprParserWS                 = 35
	TafexprParserUNKNOWN            = 36
)

// TafexprParser rules.
const (
	TafexprParserRULE_taf_expression          = 0
	TafexprParserRULE_expression              = 1
	TafexprParserRULE_var_expression          = 2
	TafexprParserRULE_indx_expr               = 3
	TafexprParserRULE_var_path                = 4
	TafexprParserRULE_jsonpath_expr           = 5
	TafexprParserRULE_identifierWithQualifier = 6
	TafexprParserRULE_index_expression        = 7
	TafexprParserRULE_parenthesisExpression   = 8
	TafexprParserRULE_json                    = 9
	TafexprParserRULE_obj                     = 10
	TafexprParserRULE_pair                    = 11
	TafexprParserRULE_arr                     = 12
	TafexprParserRULE_value                   = 13
)

// ITaf_expressionContext is an interface to support dynamic dispatch.
type ITaf_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsTaf_expressionContext differentiates from other interfaces.
	IsTaf_expressionContext()
}

type Taf_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaf_expressionContext() *Taf_expressionContext {
	var p = new(Taf_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_taf_expression
	return p
}

func InitEmptyTaf_expressionContext(p *Taf_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_taf_expression
}

func (*Taf_expressionContext) IsTaf_expressionContext() {}

func NewTaf_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Taf_expressionContext {
	var p = new(Taf_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_taf_expression

	return p
}

func (s *Taf_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Taf_expressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Taf_expressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TafexprParserEOF, 0)
}

func (s *Taf_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Taf_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Taf_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterTaf_expression(s)
	}
}

func (s *Taf_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitTaf_expression(s)
	}
}

func (p *TafexprParser) Taf_expression() (localctx ITaf_expressionContext) {
	localctx = NewTaf_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TafexprParserRULE_taf_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.expression(0)
	}
	{
		p.SetState(29)
		p.Match(TafexprParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HandleLogicalNegationContext struct {
	ExpressionContext
}

func NewHandleLogicalNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleLogicalNegationContext {
	var p = new(HandleLogicalNegationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleLogicalNegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleLogicalNegationContext) LOGICAL_NOT() antlr.TerminalNode {
	return s.GetToken(TafexprParserLOGICAL_NOT, 0)
}

func (s *HandleLogicalNegationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HandleLogicalNegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleLogicalNegation(s)
	}
}

func (s *HandleLogicalNegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleLogicalNegation(s)
	}
}

type HandleNegationContext struct {
	ExpressionContext
}

func NewHandleNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleNegationContext {
	var p = new(HandleNegationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleNegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleNegationContext) SUB() antlr.TerminalNode {
	return s.GetToken(TafexprParserSUB, 0)
}

func (s *HandleNegationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HandleNegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleNegation(s)
	}
}

func (s *HandleNegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleNegation(s)
	}
}

type HandleVarExpressionContext struct {
	ExpressionContext
}

func NewHandleVarExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleVarExpressionContext {
	var p = new(HandleVarExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleVarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleVarExpressionContext) Var_expression() IVar_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_expressionContext)
}

func (s *HandleVarExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleVarExpression(s)
	}
}

func (s *HandleVarExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleVarExpression(s)
	}
}

type HandleMethodCallContext struct {
	ExpressionContext
}

func NewHandleMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleMethodCallContext {
	var p = new(HandleMethodCallContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleMethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleMethodCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *HandleMethodCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HandleMethodCallContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleMethodCallContext) PROP() antlr.TerminalNode {
	return s.GetToken(TafexprParserPROP, 0)
}

func (s *HandleMethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleMethodCall(s)
	}
}

func (s *HandleMethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleMethodCall(s)
	}
}

type MulDivContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(TafexprParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(TafexprParserDIV, 0)
}

func (s *MulDivContext) MOD() antlr.TerminalNode {
	return s.GetToken(TafexprParserMOD, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(TafexprParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(TafexprParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type HandleFunctionCallContext struct {
	ExpressionContext
}

func NewHandleFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFunctionCallContext {
	var p = new(HandleFunctionCallContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFunctionCallContext) PROP() antlr.TerminalNode {
	return s.GetToken(TafexprParserPROP, 0)
}

func (s *HandleFunctionCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *HandleFunctionCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HandleFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFunctionCall(s)
	}
}

func (s *HandleFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFunctionCall(s)
	}
}

type HandleNullContext struct {
	ExpressionContext
}

func NewHandleNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleNullContext {
	var p = new(HandleNullContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleNullContext) NULL_TOKEN() antlr.TerminalNode {
	return s.GetToken(TafexprParserNULL_TOKEN, 0)
}

func (s *HandleNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleNull(s)
	}
}

func (s *HandleNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleNull(s)
	}
}

type HandleStringContext struct {
	ExpressionContext
}

func NewHandleStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleStringContext {
	var p = new(HandleStringContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(TafexprParserSTRING, 0)
}

func (s *HandleStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleString(s)
	}
}

func (s *HandleStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleString(s)
	}
}

type OrderedEvaluationContext struct {
	ExpressionContext
}

func NewOrderedEvaluationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrderedEvaluationContext {
	var p = new(OrderedEvaluationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrderedEvaluationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderedEvaluationContext) ParenthesisExpression() IParenthesisExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesisExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesisExpressionContext)
}

func (s *OrderedEvaluationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterOrderedEvaluation(s)
	}
}

func (s *OrderedEvaluationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitOrderedEvaluation(s)
	}
}

type LogicalOperationContext struct {
	ExpressionContext
	op antlr.Token
}

func NewLogicalOperationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOperationContext {
	var p = new(LogicalOperationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalOperationContext) GetOp() antlr.Token { return s.op }

func (s *LogicalOperationContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperationContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOperationContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalOperationContext) LESSER_THAN() antlr.TerminalNode {
	return s.GetToken(TafexprParserLESSER_THAN, 0)
}

func (s *LogicalOperationContext) LESSER_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(TafexprParserLESSER_THAN_EQUAL, 0)
}

func (s *LogicalOperationContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(TafexprParserGREATER_THAN, 0)
}

func (s *LogicalOperationContext) GREATER_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(TafexprParserGREATER_THAN_EQUAL, 0)
}

func (s *LogicalOperationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TafexprParserEQUAL, 0)
}

func (s *LogicalOperationContext) UNEQUAL() antlr.TerminalNode {
	return s.GetToken(TafexprParserUNEQUAL, 0)
}

func (s *LogicalOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterLogicalOperation(s)
	}
}

func (s *LogicalOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitLogicalOperation(s)
	}
}

type HandleBoolContext struct {
	ExpressionContext
}

func NewHandleBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleBoolContext {
	var p = new(HandleBoolContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleBoolContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(TafexprParserBOOLEAN, 0)
}

func (s *HandleBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleBool(s)
	}
}

func (s *HandleBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleBool(s)
	}
}

type NumberContext struct {
	ExpressionContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(TafexprParserINTEGER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitNumber(s)
	}
}

type HandleJsonContext struct {
	ExpressionContext
}

func NewHandleJsonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleJsonContext {
	var p = new(HandleJsonContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleJsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleJsonContext) Json() IJsonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *HandleJsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleJson(s)
	}
}

func (s *HandleJsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleJson(s)
	}
}

type DoubleValueContext struct {
	ExpressionContext
}

func NewDoubleValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleValueContext {
	var p = new(DoubleValueContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *DoubleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleValueContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(TafexprParserDOUBLE, 0)
}

func (s *DoubleValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterDoubleValue(s)
	}
}

func (s *DoubleValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitDoubleValue(s)
	}
}

type HandleLogicalContext struct {
	ExpressionContext
	op antlr.Token
}

func NewHandleLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleLogicalContext {
	var p = new(HandleLogicalContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleLogicalContext) GetOp() antlr.Token { return s.op }

func (s *HandleLogicalContext) SetOp(v antlr.Token) { s.op = v }

func (s *HandleLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleLogicalContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *HandleLogicalContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HandleLogicalContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(TafexprParserLOGICAL_AND, 0)
}

func (s *HandleLogicalContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(TafexprParserLOGICAL_OR, 0)
}

func (s *HandleLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleLogical(s)
	}
}

func (s *HandleLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleLogical(s)
	}
}

func (p *TafexprParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TafexprParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, TafexprParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TafexprParserPROP:
		localctx = NewHandleFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(32)
			p.Match(TafexprParserPROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.Match(TafexprParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14630041618) != 0 {
			{
				p.SetState(34)
				p.expression(0)
			}
			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == TafexprParserT__1 {
				{
					p.SetState(35)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(36)
					p.expression(0)
				}

				p.SetState(41)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(44)
			p.Match(TafexprParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserSUB:
		localctx = NewHandleNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(TafexprParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.expression(16)
		}

	case TafexprParserLOGICAL_NOT:
		localctx = NewHandleLogicalNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.Match(TafexprParserLOGICAL_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.expression(15)
		}

	case TafexprParserINTEGER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(TafexprParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserDOUBLE:
		localctx = NewDoubleValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(TafexprParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserT__0:
		localctx = NewOrderedEvaluationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(51)
			p.ParenthesisExpression()
		}

	case TafexprParserVARIABLE_NAME:
		localctx = NewHandleVarExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Var_expression()
		}

	case TafexprParserBOOLEAN:
		localctx = NewHandleBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Match(TafexprParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserNULL_TOKEN:
		localctx = NewHandleNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(TafexprParserNULL_TOKEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserSTRING:
		localctx = NewHandleStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
			p.Match(TafexprParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserT__3, TafexprParserLBR:
		localctx = NewHandleJsonContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(56)
			p.Json()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(60)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)
					p.expression(15)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(63)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TafexprParserADD || _la == TafexprParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(64)
					p.expression(14)
				}

			case 3:
				localctx = NewLogicalOperationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(66)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalOperationContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26738688) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalOperationContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(67)
					p.expression(13)
				}

			case 4:
				localctx = NewLogicalOperationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(69)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalOperationContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TafexprParserEQUAL || _la == TafexprParserUNEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalOperationContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(70)
					p.expression(12)
				}

			case 5:
				localctx = NewHandleLogicalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(72)

					var _m = p.Match(TafexprParserLOGICAL_AND)

					localctx.(*HandleLogicalContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(73)
					p.expression(11)
				}

			case 6:
				localctx = NewHandleLogicalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(75)

					var _m = p.Match(TafexprParserLOGICAL_OR)

					localctx.(*HandleLogicalContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(76)
					p.expression(10)
				}

			case 7:
				localctx = NewHandleMethodCallContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(78)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(79)
					p.Match(TafexprParserPROP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.Match(TafexprParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(89)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14630041618) != 0 {
					{
						p.SetState(81)
						p.expression(0)
					}
					p.SetState(86)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					for _la == TafexprParserT__1 {
						{
							p.SetState(82)
							p.Match(TafexprParserT__1)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(83)
							p.expression(0)
						}

						p.SetState(88)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(91)
					p.Match(TafexprParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_expressionContext is an interface to support dynamic dispatch.
type IVar_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARIABLE_NAME() antlr.TerminalNode
	Indx_expr() IIndx_exprContext

	// IsVar_expressionContext differentiates from other interfaces.
	IsVar_expressionContext()
}

type Var_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_expressionContext() *Var_expressionContext {
	var p = new(Var_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_expression
	return p
}

func InitEmptyVar_expressionContext(p *Var_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_expression
}

func (*Var_expressionContext) IsVar_expressionContext() {}

func NewVar_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_expressionContext {
	var p = new(Var_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_var_expression

	return p
}

func (s *Var_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_expressionContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(TafexprParserVARIABLE_NAME, 0)
}

func (s *Var_expressionContext) Indx_expr() IIndx_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndx_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndx_exprContext)
}

func (s *Var_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterVar_expression(s)
	}
}

func (s *Var_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitVar_expression(s)
	}
}

func (p *TafexprParser) Var_expression() (localctx IVar_expressionContext) {
	localctx = NewVar_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TafexprParserRULE_var_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(TafexprParserVARIABLE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Indx_expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndx_exprContext is an interface to support dynamic dispatch.
type IIndx_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLBR() []antlr.TerminalNode
	LBR(i int) antlr.TerminalNode
	AllIndex_expression() []IIndex_expressionContext
	Index_expression(i int) IIndex_expressionContext
	AllRBR() []antlr.TerminalNode
	RBR(i int) antlr.TerminalNode
	CON() antlr.TerminalNode
	Var_path() IVar_pathContext

	// IsIndx_exprContext differentiates from other interfaces.
	IsIndx_exprContext()
}

type Indx_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndx_exprContext() *Indx_exprContext {
	var p = new(Indx_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_indx_expr
	return p
}

func InitEmptyIndx_exprContext(p *Indx_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_indx_expr
}

func (*Indx_exprContext) IsIndx_exprContext() {}

func NewIndx_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Indx_exprContext {
	var p = new(Indx_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_indx_expr

	return p
}

func (s *Indx_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Indx_exprContext) AllLBR() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserLBR)
}

func (s *Indx_exprContext) LBR(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserLBR, i)
}

func (s *Indx_exprContext) AllIndex_expression() []IIndex_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndex_expressionContext); ok {
			len++
		}
	}

	tst := make([]IIndex_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndex_expressionContext); ok {
			tst[i] = t.(IIndex_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Indx_exprContext) Index_expression(i int) IIndex_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_expressionContext)
}

func (s *Indx_exprContext) AllRBR() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserRBR)
}

func (s *Indx_exprContext) RBR(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserRBR, i)
}

func (s *Indx_exprContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *Indx_exprContext) Var_path() IVar_pathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_pathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_pathContext)
}

func (s *Indx_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Indx_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Indx_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterIndx_expr(s)
	}
}

func (s *Indx_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitIndx_expr(s)
	}
}

func (p *TafexprParser) Indx_expr() (localctx IIndx_exprContext) {
	localctx = NewIndx_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TafexprParserRULE_indx_expr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(100)
			p.Match(TafexprParserLBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Index_expression()
		}
		{
			p.SetState(102)
			p.Match(TafexprParserRBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(103)
					p.Match(TafexprParserLBR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(104)
					p.Index_expression()
				}
				{
					p.SetState(105)
					p.Match(TafexprParserRBR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(114)
			p.Match(TafexprParserCON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Var_path()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_pathContext is an interface to support dynamic dispatch.
type IVar_pathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJsonpath_expr() []IJsonpath_exprContext
	Jsonpath_expr(i int) IJsonpath_exprContext
	AllCON() []antlr.TerminalNode
	CON(i int) antlr.TerminalNode

	// IsVar_pathContext differentiates from other interfaces.
	IsVar_pathContext()
}

type Var_pathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_pathContext() *Var_pathContext {
	var p = new(Var_pathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_path
	return p
}

func InitEmptyVar_pathContext(p *Var_pathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_var_path
}

func (*Var_pathContext) IsVar_pathContext() {}

func NewVar_pathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_pathContext {
	var p = new(Var_pathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_var_path

	return p
}

func (s *Var_pathContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_pathContext) AllJsonpath_expr() []IJsonpath_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJsonpath_exprContext); ok {
			len++
		}
	}

	tst := make([]IJsonpath_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJsonpath_exprContext); ok {
			tst[i] = t.(IJsonpath_exprContext)
			i++
		}
	}

	return tst
}

func (s *Var_pathContext) Jsonpath_expr(i int) IJsonpath_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonpath_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonpath_exprContext)
}

func (s *Var_pathContext) AllCON() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserCON)
}

func (s *Var_pathContext) CON(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, i)
}

func (s *Var_pathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_pathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_pathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterVar_path(s)
	}
}

func (s *Var_pathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitVar_path(s)
	}
}

func (p *TafexprParser) Var_path() (localctx IVar_pathContext) {
	localctx = NewVar_pathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TafexprParserRULE_var_path)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Jsonpath_expr()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(119)
				p.Match(TafexprParserCON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(120)
				p.Jsonpath_expr()
			}

		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJsonpath_exprContext is an interface to support dynamic dispatch.
type IJsonpath_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierWithQualifier() IIdentifierWithQualifierContext
	PROP() antlr.TerminalNode

	// IsJsonpath_exprContext differentiates from other interfaces.
	IsJsonpath_exprContext()
}

type Jsonpath_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonpath_exprContext() *Jsonpath_exprContext {
	var p = new(Jsonpath_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_jsonpath_expr
	return p
}

func InitEmptyJsonpath_exprContext(p *Jsonpath_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_jsonpath_expr
}

func (*Jsonpath_exprContext) IsJsonpath_exprContext() {}

func NewJsonpath_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Jsonpath_exprContext {
	var p = new(Jsonpath_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_jsonpath_expr

	return p
}

func (s *Jsonpath_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Jsonpath_exprContext) IdentifierWithQualifier() IIdentifierWithQualifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierWithQualifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierWithQualifierContext)
}

func (s *Jsonpath_exprContext) PROP() antlr.TerminalNode {
	return s.GetToken(TafexprParserPROP, 0)
}

func (s *Jsonpath_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Jsonpath_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Jsonpath_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterJsonpath_expr(s)
	}
}

func (s *Jsonpath_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitJsonpath_expr(s)
	}
}

func (p *TafexprParser) Jsonpath_expr() (localctx IJsonpath_exprContext) {
	localctx = NewJsonpath_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TafexprParserRULE_jsonpath_expr)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.IdentifierWithQualifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(TafexprParserPROP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierWithQualifierContext is an interface to support dynamic dispatch.
type IIdentifierWithQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROP() antlr.TerminalNode
	AllLBR() []antlr.TerminalNode
	LBR(i int) antlr.TerminalNode
	AllIndex_expression() []IIndex_expressionContext
	Index_expression(i int) IIndex_expressionContext
	AllRBR() []antlr.TerminalNode
	RBR(i int) antlr.TerminalNode

	// IsIdentifierWithQualifierContext differentiates from other interfaces.
	IsIdentifierWithQualifierContext()
}

type IdentifierWithQualifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierWithQualifierContext() *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_identifierWithQualifier
	return p
}

func InitEmptyIdentifierWithQualifierContext(p *IdentifierWithQualifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_identifierWithQualifier
}

func (*IdentifierWithQualifierContext) IsIdentifierWithQualifierContext() {}

func NewIdentifierWithQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_identifierWithQualifier

	return p
}

func (s *IdentifierWithQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierWithQualifierContext) PROP() antlr.TerminalNode {
	return s.GetToken(TafexprParserPROP, 0)
}

func (s *IdentifierWithQualifierContext) AllLBR() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserLBR)
}

func (s *IdentifierWithQualifierContext) LBR(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserLBR, i)
}

func (s *IdentifierWithQualifierContext) AllIndex_expression() []IIndex_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndex_expressionContext); ok {
			len++
		}
	}

	tst := make([]IIndex_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndex_expressionContext); ok {
			tst[i] = t.(IIndex_expressionContext)
			i++
		}
	}

	return tst
}

func (s *IdentifierWithQualifierContext) Index_expression(i int) IIndex_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_expressionContext)
}

func (s *IdentifierWithQualifierContext) AllRBR() []antlr.TerminalNode {
	return s.GetTokens(TafexprParserRBR)
}

func (s *IdentifierWithQualifierContext) RBR(i int) antlr.TerminalNode {
	return s.GetToken(TafexprParserRBR, i)
}

func (s *IdentifierWithQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierWithQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierWithQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterIdentifierWithQualifier(s)
	}
}

func (s *IdentifierWithQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitIdentifierWithQualifier(s)
	}
}

func (p *TafexprParser) IdentifierWithQualifier() (localctx IIdentifierWithQualifierContext) {
	localctx = NewIdentifierWithQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TafexprParserRULE_identifierWithQualifier)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(TafexprParserPROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(TafexprParserLBR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Index_expression()
	}
	{
		p.SetState(133)
		p.Match(TafexprParserRBR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(134)
				p.Match(TafexprParserLBR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(135)
				p.Index_expression()
			}
			{
				p.SetState(136)
				p.Match(TafexprParserRBR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndex_expressionContext is an interface to support dynamic dispatch.
type IIndex_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndex_expressionContext differentiates from other interfaces.
	IsIndex_expressionContext()
}

type Index_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_expressionContext() *Index_expressionContext {
	var p = new(Index_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_index_expression
	return p
}

func InitEmptyIndex_expressionContext(p *Index_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_index_expression
}

func (*Index_expressionContext) IsIndex_expressionContext() {}

func NewIndex_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_expressionContext {
	var p = new(Index_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_index_expression

	return p
}

func (s *Index_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_expressionContext) CopyAll(ctx *Index_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Index_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexExpressionContext struct {
	Index_expressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyIndex_expressionContext(&p.Index_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Index_expressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterIndexExpression(s)
	}
}

func (s *IndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitIndexExpression(s)
	}
}

func (p *TafexprParser) Index_expression() (localctx IIndex_expressionContext) {
	localctx = NewIndex_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TafexprParserRULE_index_expression)
	localctx = NewIndexExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParenthesisExpressionContext is an interface to support dynamic dispatch.
type IParenthesisExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParenthesisExpression() IParenthesisExpressionContext
	Expression() IExpressionContext

	// IsParenthesisExpressionContext differentiates from other interfaces.
	IsParenthesisExpressionContext()
}

type ParenthesisExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenthesisExpressionContext() *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_parenthesisExpression
	return p
}

func InitEmptyParenthesisExpressionContext(p *ParenthesisExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_parenthesisExpression
}

func (*ParenthesisExpressionContext) IsParenthesisExpressionContext() {}

func NewParenthesisExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_parenthesisExpression

	return p
}

func (s *ParenthesisExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenthesisExpressionContext) ParenthesisExpression() IParenthesisExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesisExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesisExpressionContext)
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenthesisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitParenthesisExpression(s)
	}
}

func (p *TafexprParser) ParenthesisExpression() (localctx IParenthesisExpressionContext) {
	localctx = NewParenthesisExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TafexprParserRULE_parenthesisExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(TafexprParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(146)
			p.ParenthesisExpression()
		}

	case 2:
		{
			p.SetState(147)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(150)
		p.Match(TafexprParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_json
	return p
}

func InitEmptyJsonContext(p *JsonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_json
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) CopyAll(ctx *JsonContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HandleObjectContext struct {
	JsonContext
}

func NewHandleObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleObjectContext {
	var p = new(HandleObjectContext)

	InitEmptyJsonContext(&p.JsonContext)
	p.parser = parser
	p.CopyAll(ctx.(*JsonContext))

	return p
}

func (s *HandleObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleObjectContext) Obj() IObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *HandleObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleObject(s)
	}
}

func (s *HandleObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleObject(s)
	}
}

type HandleArrayContext struct {
	JsonContext
}

func NewHandleArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleArrayContext {
	var p = new(HandleArrayContext)

	InitEmptyJsonContext(&p.JsonContext)
	p.parser = parser
	p.CopyAll(ctx.(*JsonContext))

	return p
}

func (s *HandleArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleArrayContext) Arr() IArrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrContext)
}

func (s *HandleArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleArray(s)
	}
}

func (s *HandleArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleArray(s)
	}
}

func (p *TafexprParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TafexprParserRULE_json)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TafexprParserT__3:
		localctx = NewHandleObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Obj()
		}

	case TafexprParserLBR:
		localctx = NewHandleArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Arr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_obj
	return p
}

func InitEmptyObjContext(p *ObjContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_obj
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) CopyAll(ctx *ObjContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HandleObjectDataContext struct {
	ObjContext
}

func NewHandleObjectDataContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleObjectDataContext {
	var p = new(HandleObjectDataContext)

	InitEmptyObjContext(&p.ObjContext)
	p.parser = parser
	p.CopyAll(ctx.(*ObjContext))

	return p
}

func (s *HandleObjectDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleObjectDataContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *HandleObjectDataContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *HandleObjectDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleObjectData(s)
	}
}

func (s *HandleObjectDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleObjectData(s)
	}
}

type HandleEmptyObjectDataContext struct {
	ObjContext
}

func NewHandleEmptyObjectDataContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleEmptyObjectDataContext {
	var p = new(HandleEmptyObjectDataContext)

	InitEmptyObjContext(&p.ObjContext)
	p.parser = parser
	p.CopyAll(ctx.(*ObjContext))

	return p
}

func (s *HandleEmptyObjectDataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleEmptyObjectDataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleEmptyObjectData(s)
	}
}

func (s *HandleEmptyObjectDataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleEmptyObjectData(s)
	}
}

func (p *TafexprParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TafexprParserRULE_obj)
	var _la int

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewHandleObjectDataContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Match(TafexprParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Pair()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TafexprParserT__1 {
			{
				p.SetState(158)
				p.Match(TafexprParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(159)
				p.Pair()
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(165)
			p.Match(TafexprParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewHandleEmptyObjectDataContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(TafexprParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(TafexprParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_pair
	return p
}

func InitEmptyPairContext(p *PairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_pair
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) CopyAll(ctx *PairContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HandleObjectPairContext struct {
	PairContext
}

func NewHandleObjectPairContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleObjectPairContext {
	var p = new(HandleObjectPairContext)

	InitEmptyPairContext(&p.PairContext)
	p.parser = parser
	p.CopyAll(ctx.(*PairContext))

	return p
}

func (s *HandleObjectPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleObjectPairContext) STRING() antlr.TerminalNode {
	return s.GetToken(TafexprParserSTRING, 0)
}

func (s *HandleObjectPairContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *HandleObjectPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleObjectPair(s)
	}
}

func (s *HandleObjectPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleObjectPair(s)
	}
}

func (p *TafexprParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TafexprParserRULE_pair)
	localctx = NewHandleObjectPairContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(TafexprParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(TafexprParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrContext is an interface to support dynamic dispatch.
type IArrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBR() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	RBR() antlr.TerminalNode

	// IsArrContext differentiates from other interfaces.
	IsArrContext()
}

type ArrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrContext() *ArrContext {
	var p = new(ArrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_arr
	return p
}

func InitEmptyArrContext(p *ArrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_arr
}

func (*ArrContext) IsArrContext() {}

func NewArrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrContext {
	var p = new(ArrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_arr

	return p
}

func (s *ArrContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrContext) LBR() antlr.TerminalNode {
	return s.GetToken(TafexprParserLBR, 0)
}

func (s *ArrContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrContext) RBR() antlr.TerminalNode {
	return s.GetToken(TafexprParserRBR, 0)
}

func (s *ArrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterArr(s)
	}
}

func (s *ArrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitArr(s)
	}
}

func (p *TafexprParser) Arr() (localctx IArrContext) {
	localctx = NewArrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TafexprParserRULE_arr)
	var _la int

	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Match(TafexprParserLBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Value()
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TafexprParserT__1 {
			{
				p.SetState(177)
				p.Match(TafexprParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(178)
				p.Value()
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(184)
			p.Match(TafexprParserRBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(TafexprParserLBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(TafexprParserRBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HandleJJContext struct {
	ValueContext
}

func NewHandleJJContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleJJContext {
	var p = new(HandleJJContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *HandleJJContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleJJContext) Json() IJsonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *HandleJJContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleJJ(s)
	}
}

func (s *HandleJJContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleJJ(s)
	}
}

type HandleFooContext struct {
	ValueContext
}

func NewHandleFooContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFooContext {
	var p = new(HandleFooContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *HandleFooContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFooContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HandleFooContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFoo(s)
	}
}

func (s *HandleFooContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFoo(s)
	}
}

func (p *TafexprParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TafexprParserRULE_value)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewHandleJJContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Json()
		}

	case 2:
		localctx = NewHandleFooContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TafexprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TafexprParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 18)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
