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
		"", "'randomDoubleInRange'", "'('", "','", "')'", "'length'", "'findOneByXPATH'",
		"'findOneStringByXPATH'", "'findOneDoubleByXPATH'", "'findOneIntegerByXPATH'",
		"'findOneBooleanByXPATH'", "'findByXPATH'", "'extractOneByREGEX'", "'replaceAllStringOccurrences'",
		"'toString'", "'toBoolean'", "'toInteger'", "'toDouble'", "'containsString'",
		"'startsWith'", "'endsWith'", "'trimLeft'", "'trimRight'", "'trim'",
		"':'", "'{'", "'}'", "'\"'", "'*'", "'/'", "'%'", "'+'", "'-'", "",
		"", "", "'['", "']'", "'.'", "'null'", "'<'", "'<='", "'=='", "'!='",
		"'>'", "'>='", "'&&'", "'||'", "'!'", "'$'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "MUL", "DIV", "MOD", "ADD",
		"SUB", "INTEGER", "DOUBLE", "WHITESPACE", "LBR", "RBR", "CON", "NULL_TOKEN",
		"LESSER_THAN", "LESSER_THAN_EQUAL", "EQUAL", "UNEQUAL", "GREATER_THAN",
		"GREATER_THAN_EQUAL", "LOGICAL_AND", "LOGICAL_OR", "LOGICAL_NOT", "DOLLAR",
		"STRING", "UnterminatedStringLiteral", "BOOLEAN", "NUMBER", "VARIABLE_NAME",
		"PROP", "EXCEPT_QUOTE", "JSON_NUMBER", "WS", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"taf_expression", "libfunc", "expression", "var_expression", "indx_expr",
		"var_path", "jsonpath_expr", "identifierWithQualifier", "index_expression",
		"parenthesisExpression", "any", "json", "obj", "pair", "arr", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 281, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 3, 0, 35, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 177, 8, 2, 10, 2, 12, 2, 180, 9, 2, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 192, 8, 4, 10, 4, 12, 4, 195,
		9, 4, 3, 4, 197, 8, 4, 1, 4, 1, 4, 3, 4, 201, 8, 4, 1, 5, 1, 5, 1, 5, 5,
		5, 206, 8, 5, 10, 5, 12, 5, 209, 9, 5, 1, 6, 1, 6, 3, 6, 213, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 223, 8, 7, 10, 7, 12,
		7, 226, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 233, 8, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 241, 8, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 5, 12, 247, 8, 12, 10, 12, 12, 12, 250, 9, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 256, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 5, 14, 266, 8, 14, 10, 14, 12, 14, 269, 9, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 3, 14, 275, 8, 14, 1, 15, 1, 15, 3, 15, 279, 8, 15, 1, 15,
		0, 1, 4, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		0, 5, 1, 0, 46, 47, 1, 0, 28, 30, 1, 0, 31, 32, 1, 0, 40, 45, 2, 0, 24,
		27, 36, 37, 311, 0, 34, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4, 59, 1, 0, 0,
		0, 6, 181, 1, 0, 0, 0, 8, 196, 1, 0, 0, 0, 10, 202, 1, 0, 0, 0, 12, 212,
		1, 0, 0, 0, 14, 214, 1, 0, 0, 0, 16, 227, 1, 0, 0, 0, 18, 229, 1, 0, 0,
		0, 20, 236, 1, 0, 0, 0, 22, 240, 1, 0, 0, 0, 24, 255, 1, 0, 0, 0, 26, 257,
		1, 0, 0, 0, 28, 274, 1, 0, 0, 0, 30, 278, 1, 0, 0, 0, 32, 35, 3, 2, 1,
		0, 33, 35, 3, 4, 2, 0, 34, 32, 1, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 37, 5, 0, 0, 1, 37, 1, 1, 0, 0, 0, 38, 39, 5, 1, 0, 0,
		39, 40, 5, 2, 0, 0, 40, 41, 3, 4, 2, 0, 41, 42, 5, 3, 0, 0, 42, 43, 3,
		4, 2, 0, 43, 44, 5, 4, 0, 0, 44, 3, 1, 0, 0, 0, 45, 46, 6, 2, -1, 0, 46,
		47, 5, 32, 0, 0, 47, 60, 3, 4, 2, 34, 48, 49, 5, 48, 0, 0, 49, 60, 3, 4,
		2, 33, 50, 60, 3, 2, 1, 0, 51, 60, 5, 33, 0, 0, 52, 60, 5, 34, 0, 0, 53,
		60, 3, 18, 9, 0, 54, 60, 3, 6, 3, 0, 55, 60, 5, 52, 0, 0, 56, 60, 5, 39,
		0, 0, 57, 60, 5, 50, 0, 0, 58, 60, 3, 22, 11, 0, 59, 45, 1, 0, 0, 0, 59,
		48, 1, 0, 0, 0, 59, 50, 1, 0, 0, 0, 59, 51, 1, 0, 0, 0, 59, 52, 1, 0, 0,
		0, 59, 53, 1, 0, 0, 0, 59, 54, 1, 0, 0, 0, 59, 55, 1, 0, 0, 0, 59, 56,
		1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 178, 1, 0, 0, 0,
		61, 62, 10, 32, 0, 0, 62, 63, 7, 0, 0, 0, 63, 177, 3, 4, 2, 33, 64, 65,
		10, 31, 0, 0, 65, 66, 7, 1, 0, 0, 66, 177, 3, 4, 2, 32, 67, 68, 10, 30,
		0, 0, 68, 69, 7, 2, 0, 0, 69, 177, 3, 4, 2, 31, 70, 71, 10, 29, 0, 0, 71,
		72, 7, 3, 0, 0, 72, 177, 3, 4, 2, 30, 73, 74, 10, 28, 0, 0, 74, 75, 5,
		38, 0, 0, 75, 177, 5, 5, 0, 0, 76, 77, 10, 27, 0, 0, 77, 78, 5, 38, 0,
		0, 78, 79, 5, 6, 0, 0, 79, 80, 5, 2, 0, 0, 80, 81, 3, 4, 2, 0, 81, 82,
		5, 4, 0, 0, 82, 177, 1, 0, 0, 0, 83, 84, 10, 26, 0, 0, 84, 85, 5, 38, 0,
		0, 85, 86, 5, 7, 0, 0, 86, 87, 5, 2, 0, 0, 87, 88, 3, 4, 2, 0, 88, 89,
		5, 4, 0, 0, 89, 177, 1, 0, 0, 0, 90, 91, 10, 25, 0, 0, 91, 92, 5, 38, 0,
		0, 92, 93, 5, 8, 0, 0, 93, 94, 5, 2, 0, 0, 94, 95, 3, 4, 2, 0, 95, 96,
		5, 4, 0, 0, 96, 177, 1, 0, 0, 0, 97, 98, 10, 24, 0, 0, 98, 99, 5, 38, 0,
		0, 99, 100, 5, 9, 0, 0, 100, 101, 5, 2, 0, 0, 101, 102, 3, 4, 2, 0, 102,
		103, 5, 4, 0, 0, 103, 177, 1, 0, 0, 0, 104, 105, 10, 23, 0, 0, 105, 106,
		5, 38, 0, 0, 106, 107, 5, 10, 0, 0, 107, 108, 5, 2, 0, 0, 108, 109, 3,
		4, 2, 0, 109, 110, 5, 4, 0, 0, 110, 177, 1, 0, 0, 0, 111, 112, 10, 22,
		0, 0, 112, 113, 5, 38, 0, 0, 113, 114, 5, 11, 0, 0, 114, 115, 5, 2, 0,
		0, 115, 116, 3, 4, 2, 0, 116, 117, 5, 4, 0, 0, 117, 177, 1, 0, 0, 0, 118,
		119, 10, 21, 0, 0, 119, 120, 5, 38, 0, 0, 120, 121, 5, 12, 0, 0, 121, 122,
		5, 2, 0, 0, 122, 123, 3, 4, 2, 0, 123, 124, 5, 4, 0, 0, 124, 177, 1, 0,
		0, 0, 125, 126, 10, 20, 0, 0, 126, 127, 5, 38, 0, 0, 127, 128, 5, 13, 0,
		0, 128, 129, 5, 2, 0, 0, 129, 130, 3, 4, 2, 0, 130, 131, 5, 3, 0, 0, 131,
		132, 3, 4, 2, 0, 132, 133, 5, 4, 0, 0, 133, 177, 1, 0, 0, 0, 134, 135,
		10, 19, 0, 0, 135, 136, 5, 38, 0, 0, 136, 177, 5, 14, 0, 0, 137, 138, 10,
		18, 0, 0, 138, 139, 5, 38, 0, 0, 139, 177, 5, 15, 0, 0, 140, 141, 10, 17,
		0, 0, 141, 142, 5, 38, 0, 0, 142, 177, 5, 16, 0, 0, 143, 144, 10, 16, 0,
		0, 144, 145, 5, 38, 0, 0, 145, 177, 5, 17, 0, 0, 146, 147, 10, 15, 0, 0,
		147, 148, 5, 38, 0, 0, 148, 149, 5, 18, 0, 0, 149, 150, 5, 2, 0, 0, 150,
		151, 3, 4, 2, 0, 151, 152, 5, 4, 0, 0, 152, 177, 1, 0, 0, 0, 153, 154,
		10, 14, 0, 0, 154, 155, 5, 38, 0, 0, 155, 156, 5, 19, 0, 0, 156, 157, 5,
		2, 0, 0, 157, 158, 3, 4, 2, 0, 158, 159, 5, 4, 0, 0, 159, 177, 1, 0, 0,
		0, 160, 161, 10, 13, 0, 0, 161, 162, 5, 38, 0, 0, 162, 163, 5, 20, 0, 0,
		163, 164, 5, 2, 0, 0, 164, 165, 3, 4, 2, 0, 165, 166, 5, 4, 0, 0, 166,
		177, 1, 0, 0, 0, 167, 168, 10, 12, 0, 0, 168, 169, 5, 38, 0, 0, 169, 177,
		5, 21, 0, 0, 170, 171, 10, 10, 0, 0, 171, 172, 5, 38, 0, 0, 172, 177, 5,
		22, 0, 0, 173, 174, 10, 9, 0, 0, 174, 175, 5, 38, 0, 0, 175, 177, 5, 23,
		0, 0, 176, 61, 1, 0, 0, 0, 176, 64, 1, 0, 0, 0, 176, 67, 1, 0, 0, 0, 176,
		70, 1, 0, 0, 0, 176, 73, 1, 0, 0, 0, 176, 76, 1, 0, 0, 0, 176, 83, 1, 0,
		0, 0, 176, 90, 1, 0, 0, 0, 176, 97, 1, 0, 0, 0, 176, 104, 1, 0, 0, 0, 176,
		111, 1, 0, 0, 0, 176, 118, 1, 0, 0, 0, 176, 125, 1, 0, 0, 0, 176, 134,
		1, 0, 0, 0, 176, 137, 1, 0, 0, 0, 176, 140, 1, 0, 0, 0, 176, 143, 1, 0,
		0, 0, 176, 146, 1, 0, 0, 0, 176, 153, 1, 0, 0, 0, 176, 160, 1, 0, 0, 0,
		176, 167, 1, 0, 0, 0, 176, 170, 1, 0, 0, 0, 176, 173, 1, 0, 0, 0, 177,
		180, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 5, 1,
		0, 0, 0, 180, 178, 1, 0, 0, 0, 181, 182, 5, 54, 0, 0, 182, 183, 3, 8, 4,
		0, 183, 7, 1, 0, 0, 0, 184, 185, 5, 36, 0, 0, 185, 186, 3, 16, 8, 0, 186,
		193, 5, 37, 0, 0, 187, 188, 5, 36, 0, 0, 188, 189, 3, 16, 8, 0, 189, 190,
		5, 37, 0, 0, 190, 192, 1, 0, 0, 0, 191, 187, 1, 0, 0, 0, 192, 195, 1, 0,
		0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0,
		195, 193, 1, 0, 0, 0, 196, 184, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		200, 1, 0, 0, 0, 198, 199, 5, 38, 0, 0, 199, 201, 3, 10, 5, 0, 200, 198,
		1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 9, 1, 0, 0, 0, 202, 207, 3, 12,
		6, 0, 203, 204, 5, 38, 0, 0, 204, 206, 3, 12, 6, 0, 205, 203, 1, 0, 0,
		0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208,
		11, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 213, 3, 14, 7, 0, 211, 213,
		5, 55, 0, 0, 212, 210, 1, 0, 0, 0, 212, 211, 1, 0, 0, 0, 213, 13, 1, 0,
		0, 0, 214, 215, 5, 55, 0, 0, 215, 216, 5, 36, 0, 0, 216, 217, 3, 16, 8,
		0, 217, 224, 5, 37, 0, 0, 218, 219, 5, 36, 0, 0, 219, 220, 3, 16, 8, 0,
		220, 221, 5, 37, 0, 0, 221, 223, 1, 0, 0, 0, 222, 218, 1, 0, 0, 0, 223,
		226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 15, 1,
		0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 3, 4, 2, 0, 228, 17, 1, 0, 0,
		0, 229, 232, 5, 2, 0, 0, 230, 233, 3, 18, 9, 0, 231, 233, 3, 4, 2, 0, 232,
		230, 1, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235,
		5, 4, 0, 0, 235, 19, 1, 0, 0, 0, 236, 237, 7, 4, 0, 0, 237, 21, 1, 0, 0,
		0, 238, 241, 3, 24, 12, 0, 239, 241, 3, 28, 14, 0, 240, 238, 1, 0, 0, 0,
		240, 239, 1, 0, 0, 0, 241, 23, 1, 0, 0, 0, 242, 243, 5, 25, 0, 0, 243,
		248, 3, 26, 13, 0, 244, 245, 5, 3, 0, 0, 245, 247, 3, 26, 13, 0, 246, 244,
		1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0,
		0, 0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 252, 5, 26, 0, 0,
		252, 256, 1, 0, 0, 0, 253, 254, 5, 25, 0, 0, 254, 256, 5, 26, 0, 0, 255,
		242, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 25, 1, 0, 0, 0, 257, 258, 5,
		50, 0, 0, 258, 259, 5, 24, 0, 0, 259, 260, 3, 30, 15, 0, 260, 27, 1, 0,
		0, 0, 261, 262, 5, 36, 0, 0, 262, 267, 3, 30, 15, 0, 263, 264, 5, 3, 0,
		0, 264, 266, 3, 30, 15, 0, 265, 263, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0,
		267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270, 1, 0, 0, 0, 269,
		267, 1, 0, 0, 0, 270, 271, 5, 37, 0, 0, 271, 275, 1, 0, 0, 0, 272, 273,
		5, 36, 0, 0, 273, 275, 5, 37, 0, 0, 274, 261, 1, 0, 0, 0, 274, 272, 1,
		0, 0, 0, 275, 29, 1, 0, 0, 0, 276, 279, 3, 22, 11, 0, 277, 279, 3, 4, 2,
		0, 278, 276, 1, 0, 0, 0, 278, 277, 1, 0, 0, 0, 279, 31, 1, 0, 0, 0, 17,
		34, 59, 176, 178, 193, 196, 200, 207, 212, 224, 232, 240, 248, 255, 267,
		274, 278,
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
	TafexprParserEOF                       = antlr.TokenEOF
	TafexprParserT__0                      = 1
	TafexprParserT__1                      = 2
	TafexprParserT__2                      = 3
	TafexprParserT__3                      = 4
	TafexprParserT__4                      = 5
	TafexprParserT__5                      = 6
	TafexprParserT__6                      = 7
	TafexprParserT__7                      = 8
	TafexprParserT__8                      = 9
	TafexprParserT__9                      = 10
	TafexprParserT__10                     = 11
	TafexprParserT__11                     = 12
	TafexprParserT__12                     = 13
	TafexprParserT__13                     = 14
	TafexprParserT__14                     = 15
	TafexprParserT__15                     = 16
	TafexprParserT__16                     = 17
	TafexprParserT__17                     = 18
	TafexprParserT__18                     = 19
	TafexprParserT__19                     = 20
	TafexprParserT__20                     = 21
	TafexprParserT__21                     = 22
	TafexprParserT__22                     = 23
	TafexprParserT__23                     = 24
	TafexprParserT__24                     = 25
	TafexprParserT__25                     = 26
	TafexprParserT__26                     = 27
	TafexprParserMUL                       = 28
	TafexprParserDIV                       = 29
	TafexprParserMOD                       = 30
	TafexprParserADD                       = 31
	TafexprParserSUB                       = 32
	TafexprParserINTEGER                   = 33
	TafexprParserDOUBLE                    = 34
	TafexprParserWHITESPACE                = 35
	TafexprParserLBR                       = 36
	TafexprParserRBR                       = 37
	TafexprParserCON                       = 38
	TafexprParserNULL_TOKEN                = 39
	TafexprParserLESSER_THAN               = 40
	TafexprParserLESSER_THAN_EQUAL         = 41
	TafexprParserEQUAL                     = 42
	TafexprParserUNEQUAL                   = 43
	TafexprParserGREATER_THAN              = 44
	TafexprParserGREATER_THAN_EQUAL        = 45
	TafexprParserLOGICAL_AND               = 46
	TafexprParserLOGICAL_OR                = 47
	TafexprParserLOGICAL_NOT               = 48
	TafexprParserDOLLAR                    = 49
	TafexprParserSTRING                    = 50
	TafexprParserUnterminatedStringLiteral = 51
	TafexprParserBOOLEAN                   = 52
	TafexprParserNUMBER                    = 53
	TafexprParserVARIABLE_NAME             = 54
	TafexprParserPROP                      = 55
	TafexprParserEXCEPT_QUOTE              = 56
	TafexprParserJSON_NUMBER               = 57
	TafexprParserWS                        = 58
	TafexprParserUNKNOWN                   = 59
)

// TafexprParser rules.
const (
	TafexprParserRULE_taf_expression          = 0
	TafexprParserRULE_libfunc                 = 1
	TafexprParserRULE_expression              = 2
	TafexprParserRULE_var_expression          = 3
	TafexprParserRULE_indx_expr               = 4
	TafexprParserRULE_var_path                = 5
	TafexprParserRULE_jsonpath_expr           = 6
	TafexprParserRULE_identifierWithQualifier = 7
	TafexprParserRULE_index_expression        = 8
	TafexprParserRULE_parenthesisExpression   = 9
	TafexprParserRULE_any                     = 10
	TafexprParserRULE_json                    = 11
	TafexprParserRULE_obj                     = 12
	TafexprParserRULE_pair                    = 13
	TafexprParserRULE_arr                     = 14
	TafexprParserRULE_value                   = 15
)

// ITaf_expressionContext is an interface to support dynamic dispatch.
type ITaf_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Libfunc() ILibfuncContext
	Expression() IExpressionContext

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

func (s *Taf_expressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TafexprParserEOF, 0)
}

func (s *Taf_expressionContext) Libfunc() ILibfuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILibfuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILibfuncContext)
}

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
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(32)
			p.Libfunc()
		}

	case 2:
		{
			p.SetState(33)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(36)
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

// ILibfuncContext is an interface to support dynamic dispatch.
type ILibfuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLibfuncContext differentiates from other interfaces.
	IsLibfuncContext()
}

type LibfuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibfuncContext() *LibfuncContext {
	var p = new(LibfuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_libfunc
	return p
}

func InitEmptyLibfuncContext(p *LibfuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_libfunc
}

func (*LibfuncContext) IsLibfuncContext() {}

func NewLibfuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibfuncContext {
	var p = new(LibfuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_libfunc

	return p
}

func (s *LibfuncContext) GetParser() antlr.Parser { return s.parser }

func (s *LibfuncContext) CopyAll(ctx *LibfuncContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LibfuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibfuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HandleRandomDoubleInRangeContext struct {
	LibfuncContext
}

func NewHandleRandomDoubleInRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleRandomDoubleInRangeContext {
	var p = new(HandleRandomDoubleInRangeContext)

	InitEmptyLibfuncContext(&p.LibfuncContext)
	p.parser = parser
	p.CopyAll(ctx.(*LibfuncContext))

	return p
}

func (s *HandleRandomDoubleInRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleRandomDoubleInRangeContext) AllExpression() []IExpressionContext {
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

func (s *HandleRandomDoubleInRangeContext) Expression(i int) IExpressionContext {
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

func (s *HandleRandomDoubleInRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleRandomDoubleInRange(s)
	}
}

func (s *HandleRandomDoubleInRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleRandomDoubleInRange(s)
	}
}

func (p *TafexprParser) Libfunc() (localctx ILibfuncContext) {
	localctx = NewLibfuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TafexprParserRULE_libfunc)
	localctx = NewHandleRandomDoubleInRangeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(TafexprParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(TafexprParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.expression(0)
	}
	{
		p.SetState(41)
		p.Match(TafexprParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.expression(0)
	}
	{
		p.SetState(43)
		p.Match(TafexprParserT__3)
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

type HandleFindOneByXPATHContext struct {
	ExpressionContext
}

func NewHandleFindOneByXPATHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFindOneByXPATHContext {
	var p = new(HandleFindOneByXPATHContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleFindOneByXPATHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFindOneByXPATHContext) AllExpression() []IExpressionContext {
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

func (s *HandleFindOneByXPATHContext) Expression(i int) IExpressionContext {
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

func (s *HandleFindOneByXPATHContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleFindOneByXPATHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFindOneByXPATH(s)
	}
}

func (s *HandleFindOneByXPATHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFindOneByXPATH(s)
	}
}

type HandleTrimContext struct {
	ExpressionContext
}

func NewHandleTrimContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleTrimContext {
	var p = new(HandleTrimContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleTrimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleTrimContext) Expression() IExpressionContext {
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

func (s *HandleTrimContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleTrimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleTrim(s)
	}
}

func (s *HandleTrimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleTrim(s)
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

type HandleFindOneDoubleByXPATHContext struct {
	ExpressionContext
}

func NewHandleFindOneDoubleByXPATHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFindOneDoubleByXPATHContext {
	var p = new(HandleFindOneDoubleByXPATHContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleFindOneDoubleByXPATHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFindOneDoubleByXPATHContext) AllExpression() []IExpressionContext {
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

func (s *HandleFindOneDoubleByXPATHContext) Expression(i int) IExpressionContext {
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

func (s *HandleFindOneDoubleByXPATHContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleFindOneDoubleByXPATHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFindOneDoubleByXPATH(s)
	}
}

func (s *HandleFindOneDoubleByXPATHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFindOneDoubleByXPATH(s)
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

type HandleFindOneStringByXPATHContext struct {
	ExpressionContext
}

func NewHandleFindOneStringByXPATHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFindOneStringByXPATHContext {
	var p = new(HandleFindOneStringByXPATHContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleFindOneStringByXPATHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFindOneStringByXPATHContext) AllExpression() []IExpressionContext {
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

func (s *HandleFindOneStringByXPATHContext) Expression(i int) IExpressionContext {
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

func (s *HandleFindOneStringByXPATHContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleFindOneStringByXPATHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFindOneStringByXPATH(s)
	}
}

func (s *HandleFindOneStringByXPATHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFindOneStringByXPATH(s)
	}
}

type HandleToStringContext struct {
	ExpressionContext
}

func NewHandleToStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleToStringContext {
	var p = new(HandleToStringContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleToStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleToStringContext) Expression() IExpressionContext {
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

func (s *HandleToStringContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleToStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleToString(s)
	}
}

func (s *HandleToStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleToString(s)
	}
}

type HandleLibfuncContext struct {
	ExpressionContext
}

func NewHandleLibfuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleLibfuncContext {
	var p = new(HandleLibfuncContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleLibfuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleLibfuncContext) Libfunc() ILibfuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILibfuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILibfuncContext)
}

func (s *HandleLibfuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleLibfunc(s)
	}
}

func (s *HandleLibfuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleLibfunc(s)
	}
}

type HandleFindByXPATHContext struct {
	ExpressionContext
}

func NewHandleFindByXPATHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFindByXPATHContext {
	var p = new(HandleFindByXPATHContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleFindByXPATHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFindByXPATHContext) AllExpression() []IExpressionContext {
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

func (s *HandleFindByXPATHContext) Expression(i int) IExpressionContext {
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

func (s *HandleFindByXPATHContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleFindByXPATHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFindByXPATH(s)
	}
}

func (s *HandleFindByXPATHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFindByXPATH(s)
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

type HandleExtractOneByREGEXContext struct {
	ExpressionContext
}

func NewHandleExtractOneByREGEXContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleExtractOneByREGEXContext {
	var p = new(HandleExtractOneByREGEXContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleExtractOneByREGEXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleExtractOneByREGEXContext) AllExpression() []IExpressionContext {
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

func (s *HandleExtractOneByREGEXContext) Expression(i int) IExpressionContext {
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

func (s *HandleExtractOneByREGEXContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleExtractOneByREGEXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleExtractOneByREGEX(s)
	}
}

func (s *HandleExtractOneByREGEXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleExtractOneByREGEX(s)
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

type HandleToBooleanContext struct {
	ExpressionContext
}

func NewHandleToBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleToBooleanContext {
	var p = new(HandleToBooleanContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleToBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleToBooleanContext) Expression() IExpressionContext {
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

func (s *HandleToBooleanContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleToBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleToBoolean(s)
	}
}

func (s *HandleToBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleToBoolean(s)
	}
}

type HandleTrimLeftContext struct {
	ExpressionContext
}

func NewHandleTrimLeftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleTrimLeftContext {
	var p = new(HandleTrimLeftContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleTrimLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleTrimLeftContext) Expression() IExpressionContext {
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

func (s *HandleTrimLeftContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleTrimLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleTrimLeft(s)
	}
}

func (s *HandleTrimLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleTrimLeft(s)
	}
}

type HandleFindOneBooleanByXPATHContext struct {
	ExpressionContext
}

func NewHandleFindOneBooleanByXPATHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFindOneBooleanByXPATHContext {
	var p = new(HandleFindOneBooleanByXPATHContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleFindOneBooleanByXPATHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFindOneBooleanByXPATHContext) AllExpression() []IExpressionContext {
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

func (s *HandleFindOneBooleanByXPATHContext) Expression(i int) IExpressionContext {
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

func (s *HandleFindOneBooleanByXPATHContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleFindOneBooleanByXPATHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFindOneBooleanByXPATH(s)
	}
}

func (s *HandleFindOneBooleanByXPATHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFindOneBooleanByXPATH(s)
	}
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

type HandleLengthContext struct {
	ExpressionContext
}

func NewHandleLengthContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleLengthContext {
	var p = new(HandleLengthContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleLengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleLengthContext) Expression() IExpressionContext {
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

func (s *HandleLengthContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleLengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleLength(s)
	}
}

func (s *HandleLengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleLength(s)
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

type HandleFindOneIntegerByXPATHContext struct {
	ExpressionContext
}

func NewHandleFindOneIntegerByXPATHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleFindOneIntegerByXPATHContext {
	var p = new(HandleFindOneIntegerByXPATHContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleFindOneIntegerByXPATHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleFindOneIntegerByXPATHContext) AllExpression() []IExpressionContext {
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

func (s *HandleFindOneIntegerByXPATHContext) Expression(i int) IExpressionContext {
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

func (s *HandleFindOneIntegerByXPATHContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleFindOneIntegerByXPATHContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleFindOneIntegerByXPATH(s)
	}
}

func (s *HandleFindOneIntegerByXPATHContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleFindOneIntegerByXPATH(s)
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

type HandleToDoubleContext struct {
	ExpressionContext
}

func NewHandleToDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleToDoubleContext {
	var p = new(HandleToDoubleContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleToDoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleToDoubleContext) Expression() IExpressionContext {
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

func (s *HandleToDoubleContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleToDoubleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleToDouble(s)
	}
}

func (s *HandleToDoubleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleToDouble(s)
	}
}

type HandleEndsWithContext struct {
	ExpressionContext
}

func NewHandleEndsWithContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleEndsWithContext {
	var p = new(HandleEndsWithContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleEndsWithContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleEndsWithContext) AllExpression() []IExpressionContext {
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

func (s *HandleEndsWithContext) Expression(i int) IExpressionContext {
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

func (s *HandleEndsWithContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleEndsWithContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleEndsWith(s)
	}
}

func (s *HandleEndsWithContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleEndsWith(s)
	}
}

type HandleContainsStringContext struct {
	ExpressionContext
}

func NewHandleContainsStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleContainsStringContext {
	var p = new(HandleContainsStringContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleContainsStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleContainsStringContext) AllExpression() []IExpressionContext {
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

func (s *HandleContainsStringContext) Expression(i int) IExpressionContext {
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

func (s *HandleContainsStringContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleContainsStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleContainsString(s)
	}
}

func (s *HandleContainsStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleContainsString(s)
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

func (s *LogicalOperationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TafexprParserEQUAL, 0)
}

func (s *LogicalOperationContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(TafexprParserGREATER_THAN, 0)
}

func (s *LogicalOperationContext) GREATER_THAN_EQUAL() antlr.TerminalNode {
	return s.GetToken(TafexprParserGREATER_THAN_EQUAL, 0)
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

type HandleToIntegerContext struct {
	ExpressionContext
}

func NewHandleToIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleToIntegerContext {
	var p = new(HandleToIntegerContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleToIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleToIntegerContext) Expression() IExpressionContext {
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

func (s *HandleToIntegerContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleToIntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleToInteger(s)
	}
}

func (s *HandleToIntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleToInteger(s)
	}
}

type HandleTrimRightContext struct {
	ExpressionContext
}

func NewHandleTrimRightContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleTrimRightContext {
	var p = new(HandleTrimRightContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleTrimRightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleTrimRightContext) Expression() IExpressionContext {
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

func (s *HandleTrimRightContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleTrimRightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleTrimRight(s)
	}
}

func (s *HandleTrimRightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleTrimRight(s)
	}
}

type HandleStartsWithContext struct {
	ExpressionContext
}

func NewHandleStartsWithContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleStartsWithContext {
	var p = new(HandleStartsWithContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleStartsWithContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleStartsWithContext) AllExpression() []IExpressionContext {
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

func (s *HandleStartsWithContext) Expression(i int) IExpressionContext {
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

func (s *HandleStartsWithContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleStartsWithContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleStartsWith(s)
	}
}

func (s *HandleStartsWithContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleStartsWith(s)
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

type HandleReplaceAllStringOccurrencesContext struct {
	ExpressionContext
}

func NewHandleReplaceAllStringOccurrencesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HandleReplaceAllStringOccurrencesContext {
	var p = new(HandleReplaceAllStringOccurrencesContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HandleReplaceAllStringOccurrencesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HandleReplaceAllStringOccurrencesContext) AllExpression() []IExpressionContext {
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

func (s *HandleReplaceAllStringOccurrencesContext) Expression(i int) IExpressionContext {
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

func (s *HandleReplaceAllStringOccurrencesContext) CON() antlr.TerminalNode {
	return s.GetToken(TafexprParserCON, 0)
}

func (s *HandleReplaceAllStringOccurrencesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterHandleReplaceAllStringOccurrences(s)
	}
}

func (s *HandleReplaceAllStringOccurrencesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitHandleReplaceAllStringOccurrences(s)
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, TafexprParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TafexprParserSUB:
		localctx = NewHandleNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(46)
			p.Match(TafexprParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.expression(34)
		}

	case TafexprParserLOGICAL_NOT:
		localctx = NewHandleLogicalNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(TafexprParserLOGICAL_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.expression(33)
		}

	case TafexprParserT__0:
		localctx = NewHandleLibfuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Libfunc()
		}

	case TafexprParserINTEGER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(51)
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
			p.SetState(52)
			p.Match(TafexprParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserT__1:
		localctx = NewOrderedEvaluationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.ParenthesisExpression()
		}

	case TafexprParserVARIABLE_NAME:
		localctx = NewHandleVarExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Var_expression()
		}

	case TafexprParserBOOLEAN:
		localctx = NewHandleBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
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
			p.SetState(56)
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
			p.SetState(57)
			p.Match(TafexprParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TafexprParserT__24, TafexprParserLBR:
		localctx = NewHandleJsonContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Json()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewHandleLogicalContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
					goto errorExit
				}
				{
					p.SetState(62)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*HandleLogicalContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TafexprParserLOGICAL_AND || _la == TafexprParserLOGICAL_OR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*HandleLogicalContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(63)
					p.expression(33)
				}

			case 2:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
					goto errorExit
				}
				{
					p.SetState(65)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1879048192) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(66)
					p.expression(32)
				}

			case 3:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 30)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 30)", ""))
					goto errorExit
				}
				{
					p.SetState(68)

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
					p.SetState(69)
					p.expression(31)
				}

			case 4:
				localctx = NewLogicalOperationContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
					goto errorExit
				}
				{
					p.SetState(71)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicalOperationContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69269232549888) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicalOperationContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(72)
					p.expression(30)
				}

			case 5:
				localctx = NewHandleLengthContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				{
					p.SetState(74)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(75)
					p.Match(TafexprParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 6:
				localctx = NewHandleFindOneByXPATHContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				{
					p.SetState(77)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(78)
					p.Match(TafexprParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(79)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.expression(0)
				}
				{
					p.SetState(81)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 7:
				localctx = NewHandleFindOneStringByXPATHContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				{
					p.SetState(84)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(85)
					p.Match(TafexprParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(87)
					p.expression(0)
				}
				{
					p.SetState(88)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 8:
				localctx = NewHandleFindOneDoubleByXPATHContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)
					p.Match(TafexprParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(93)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(94)
					p.expression(0)
				}
				{
					p.SetState(95)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 9:
				localctx = NewHandleFindOneIntegerByXPATHContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(98)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(99)
					p.Match(TafexprParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(100)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(101)
					p.expression(0)
				}
				{
					p.SetState(102)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewHandleFindOneBooleanByXPATHContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(105)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.Match(TafexprParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(107)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(108)
					p.expression(0)
				}
				{
					p.SetState(109)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 11:
				localctx = NewHandleFindByXPATHContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.Match(TafexprParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(114)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(115)
					p.expression(0)
				}
				{
					p.SetState(116)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 12:
				localctx = NewHandleExtractOneByREGEXContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
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
					p.Match(TafexprParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(121)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(122)
					p.expression(0)
				}
				{
					p.SetState(123)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 13:
				localctx = NewHandleReplaceAllStringOccurrencesContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(126)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(127)
					p.Match(TafexprParserT__12)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(128)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(129)
					p.expression(0)
				}
				{
					p.SetState(130)
					p.Match(TafexprParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)
					p.expression(0)
				}
				{
					p.SetState(132)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 14:
				localctx = NewHandleToStringContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(135)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(136)
					p.Match(TafexprParserT__13)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 15:
				localctx = NewHandleToBooleanContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(138)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(139)
					p.Match(TafexprParserT__14)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 16:
				localctx = NewHandleToIntegerContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(141)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(142)
					p.Match(TafexprParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 17:
				localctx = NewHandleToDoubleContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(144)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(145)
					p.Match(TafexprParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 18:
				localctx = NewHandleContainsStringContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(147)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(148)
					p.Match(TafexprParserT__17)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(149)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(150)
					p.expression(0)
				}
				{
					p.SetState(151)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 19:
				localctx = NewHandleStartsWithContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(154)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(155)
					p.Match(TafexprParserT__18)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(156)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(157)
					p.expression(0)
				}
				{
					p.SetState(158)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 20:
				localctx = NewHandleEndsWithContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(161)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(162)
					p.Match(TafexprParserT__19)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(163)
					p.Match(TafexprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(164)
					p.expression(0)
				}
				{
					p.SetState(165)
					p.Match(TafexprParserT__3)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 21:
				localctx = NewHandleTrimLeftContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(168)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(169)
					p.Match(TafexprParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 22:
				localctx = NewHandleTrimRightContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(171)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(172)
					p.Match(TafexprParserT__21)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 23:
				localctx = NewHandleTrimContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TafexprParserRULE_expression)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(174)
					p.Match(TafexprParserCON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(175)
					p.Match(TafexprParserT__22)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, TafexprParserRULE_var_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(TafexprParserVARIABLE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
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
	p.EnterRule(localctx, 8, TafexprParserRULE_indx_expr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(184)
			p.Match(TafexprParserLBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Index_expression()
		}
		{
			p.SetState(186)
			p.Match(TafexprParserRBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(187)
					p.Match(TafexprParserLBR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(188)
					p.Index_expression()
				}
				{
					p.SetState(189)
					p.Match(TafexprParserRBR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(198)
			p.Match(TafexprParserCON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
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
	p.EnterRule(localctx, 10, TafexprParserRULE_var_path)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Jsonpath_expr()
	}
	p.SetState(207)
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
				p.SetState(203)
				p.Match(TafexprParserCON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(204)
				p.Jsonpath_expr()
			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 12, TafexprParserRULE_jsonpath_expr)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.IdentifierWithQualifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
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
	p.EnterRule(localctx, 14, TafexprParserRULE_identifierWithQualifier)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(TafexprParserPROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(TafexprParserLBR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Index_expression()
	}
	{
		p.SetState(217)
		p.Match(TafexprParserRBR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(218)
				p.Match(TafexprParserLBR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(219)
				p.Index_expression()
			}
			{
				p.SetState(220)
				p.Match(TafexprParserRBR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 16, TafexprParserRULE_index_expression)
	localctx = NewIndexExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
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
	p.EnterRule(localctx, 18, TafexprParserRULE_parenthesisExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(TafexprParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(230)
			p.ParenthesisExpression()
		}

	case 2:
		{
			p.SetState(231)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(234)
		p.Match(TafexprParserT__3)
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

// IAnyContext is an interface to support dynamic dispatch.
type IAnyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBR() antlr.TerminalNode
	RBR() antlr.TerminalNode

	// IsAnyContext differentiates from other interfaces.
	IsAnyContext()
}

type AnyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyContext() *AnyContext {
	var p = new(AnyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_any
	return p
}

func InitEmptyAnyContext(p *AnyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TafexprParserRULE_any
}

func (*AnyContext) IsAnyContext() {}

func NewAnyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyContext {
	var p = new(AnyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TafexprParserRULE_any

	return p
}

func (s *AnyContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyContext) LBR() antlr.TerminalNode {
	return s.GetToken(TafexprParserLBR, 0)
}

func (s *AnyContext) RBR() antlr.TerminalNode {
	return s.GetToken(TafexprParserRBR, 0)
}

func (s *AnyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.EnterAny(s)
	}
}

func (s *AnyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TafexprListener); ok {
		listenerT.ExitAny(s)
	}
}

func (p *TafexprParser) Any() (localctx IAnyContext) {
	localctx = NewAnyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TafexprParserRULE_any)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&206410088448) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.EnterRule(localctx, 22, TafexprParserRULE_json)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TafexprParserT__24:
		localctx = NewHandleObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Obj()
		}

	case TafexprParserLBR:
		localctx = NewHandleArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
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
	p.EnterRule(localctx, 24, TafexprParserRULE_obj)
	var _la int

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewHandleObjectDataContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Match(TafexprParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Pair()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TafexprParserT__2 {
			{
				p.SetState(244)
				p.Match(TafexprParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(245)
				p.Pair()
			}

			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(251)
			p.Match(TafexprParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewHandleEmptyObjectDataContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(TafexprParserT__24)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(TafexprParserT__25)
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
	p.EnterRule(localctx, 26, TafexprParserRULE_pair)
	localctx = NewHandleObjectPairContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(TafexprParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(TafexprParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
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
	p.EnterRule(localctx, 28, TafexprParserRULE_arr)
	var _la int

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.Match(TafexprParserLBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Value()
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TafexprParserT__2 {
			{
				p.SetState(263)
				p.Match(TafexprParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(264)
				p.Value()
			}

			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(270)
			p.Match(TafexprParserRBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.Match(TafexprParserLBR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
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
	p.EnterRule(localctx, 30, TafexprParserRULE_value)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewHandleJJContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Json()
		}

	case 2:
		localctx = NewHandleFooContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
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
	case 2:
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
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 30)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 22:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
