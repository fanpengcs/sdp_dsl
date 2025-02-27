// Code generated from ./Lex.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Lex struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LexLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lexLexerInit() {
	staticData := &LexLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'module'", "'#'", "'include'", "'struct'", "'enum'",
		"'optional'", "'required'", "'vector'", "'map'", "'unsigned long'",
		"'long'", "'float'", "'byte'", "'unsigned int'", "'int'", "'double'",
		"'string'", "'void'", "'bool'", "'true'", "'false'", "'interface'",
		"'in'", "'out'", "'{'", "'}'", "'['", "']'", "'('", "')'", "';'", "':'",
		"','", "'='", "'<'", "'>'", "", "", "", "'::'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NEW_LINE", "LINE_COMMENT", "BLOCK_COMMENT", "MODULE", "HASH",
		"INCLUDE", "STRUCT", "ENUM", "OPTIONAL", "REQUIRED", "VECTOR", "MAP",
		"ULONG", "LONG", "FLOAT", "BYTE", "UINT", "INT", "DOUBLE", "STRING",
		"VOID", "BOOL", "TRUE", "FALSE", "INTERFACE", "IN", "OUT", "LEFT_BRACE",
		"RIGHT_BRACE", "LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_PAREN", "RIGHT_PAREN",
		"SEMICOLON", "COLON", "COMMA", "EQUAL", "LEFT_ANGLE_BRACKET", "RIGHT_ANGLE_BRACKET",
		"STRING_LITERAL", "NUMBER", "IDENTIFIER", "DOUBLE_COLON",
	}
	staticData.RuleNames = []string{
		"WS", "NEW_LINE", "LINE_COMMENT", "BLOCK_COMMENT", "MODULE", "HASH",
		"INCLUDE", "STRUCT", "ENUM", "OPTIONAL", "REQUIRED", "VECTOR", "MAP",
		"ULONG", "LONG", "FLOAT", "BYTE", "UINT", "INT", "DOUBLE", "STRING",
		"VOID", "BOOL", "TRUE", "FALSE", "INTERFACE", "IN", "OUT", "LEFT_BRACE",
		"RIGHT_BRACE", "LEFT_BRACKET", "RIGHT_BRACKET", "LEFT_PAREN", "RIGHT_PAREN",
		"SEMICOLON", "COLON", "COMMA", "EQUAL", "LEFT_ANGLE_BRACKET", "RIGHT_ANGLE_BRACKET",
		"STRING_LITERAL", "NUMBER", "IDENTIFIER", "DOUBLE_COLON",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 339, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 4, 0, 91, 8, 0, 11, 0, 12, 0, 92,
		1, 0, 1, 0, 1, 1, 3, 1, 98, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 108, 8, 2, 10, 2, 12, 2, 111, 9, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 119, 8, 3, 10, 3, 12, 3, 122, 9, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 40, 1, 40, 4, 40, 312, 8, 40, 11, 40, 12, 40, 313, 1, 40, 1, 40, 1,
		41, 3, 41, 319, 8, 41, 1, 41, 1, 41, 3, 41, 323, 8, 41, 1, 41, 4, 41, 326,
		8, 41, 11, 41, 12, 41, 327, 1, 42, 1, 42, 5, 42, 332, 8, 42, 10, 42, 12,
		42, 335, 9, 42, 1, 43, 1, 43, 1, 43, 2, 120, 313, 0, 44, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 1, 0, 7, 3, 0, 9, 9, 12, 13, 32, 32,
		1, 0, 13, 13, 2, 0, 10, 10, 13, 13, 1, 0, 45, 45, 1, 0, 48, 57, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 347, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0,
		0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0,
		0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0,
		0, 0, 0, 87, 1, 0, 0, 0, 1, 90, 1, 0, 0, 0, 3, 97, 1, 0, 0, 0, 5, 103,
		1, 0, 0, 0, 7, 114, 1, 0, 0, 0, 9, 128, 1, 0, 0, 0, 11, 135, 1, 0, 0, 0,
		13, 137, 1, 0, 0, 0, 15, 145, 1, 0, 0, 0, 17, 152, 1, 0, 0, 0, 19, 157,
		1, 0, 0, 0, 21, 166, 1, 0, 0, 0, 23, 175, 1, 0, 0, 0, 25, 182, 1, 0, 0,
		0, 27, 186, 1, 0, 0, 0, 29, 200, 1, 0, 0, 0, 31, 205, 1, 0, 0, 0, 33, 211,
		1, 0, 0, 0, 35, 216, 1, 0, 0, 0, 37, 229, 1, 0, 0, 0, 39, 233, 1, 0, 0,
		0, 41, 240, 1, 0, 0, 0, 43, 247, 1, 0, 0, 0, 45, 252, 1, 0, 0, 0, 47, 257,
		1, 0, 0, 0, 49, 262, 1, 0, 0, 0, 51, 268, 1, 0, 0, 0, 53, 278, 1, 0, 0,
		0, 55, 281, 1, 0, 0, 0, 57, 285, 1, 0, 0, 0, 59, 287, 1, 0, 0, 0, 61, 289,
		1, 0, 0, 0, 63, 291, 1, 0, 0, 0, 65, 293, 1, 0, 0, 0, 67, 295, 1, 0, 0,
		0, 69, 297, 1, 0, 0, 0, 71, 299, 1, 0, 0, 0, 73, 301, 1, 0, 0, 0, 75, 303,
		1, 0, 0, 0, 77, 305, 1, 0, 0, 0, 79, 307, 1, 0, 0, 0, 81, 309, 1, 0, 0,
		0, 83, 318, 1, 0, 0, 0, 85, 329, 1, 0, 0, 0, 87, 336, 1, 0, 0, 0, 89, 91,
		7, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0,
		92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 6, 0, 0, 0, 95, 2, 1, 0,
		0, 0, 96, 98, 7, 1, 0, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 100, 5, 10, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 6, 1,
		1, 0, 102, 4, 1, 0, 0, 0, 103, 104, 5, 47, 0, 0, 104, 105, 5, 47, 0, 0,
		105, 109, 1, 0, 0, 0, 106, 108, 8, 2, 0, 0, 107, 106, 1, 0, 0, 0, 108,
		111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112,
		1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 113, 6, 2, 1, 0, 113, 6, 1, 0, 0,
		0, 114, 115, 5, 47, 0, 0, 115, 116, 5, 42, 0, 0, 116, 120, 1, 0, 0, 0,
		117, 119, 9, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 120,
		1, 0, 0, 0, 123, 124, 5, 42, 0, 0, 124, 125, 5, 47, 0, 0, 125, 126, 1,
		0, 0, 0, 126, 127, 6, 3, 1, 0, 127, 8, 1, 0, 0, 0, 128, 129, 5, 109, 0,
		0, 129, 130, 5, 111, 0, 0, 130, 131, 5, 100, 0, 0, 131, 132, 5, 117, 0,
		0, 132, 133, 5, 108, 0, 0, 133, 134, 5, 101, 0, 0, 134, 10, 1, 0, 0, 0,
		135, 136, 5, 35, 0, 0, 136, 12, 1, 0, 0, 0, 137, 138, 5, 105, 0, 0, 138,
		139, 5, 110, 0, 0, 139, 140, 5, 99, 0, 0, 140, 141, 5, 108, 0, 0, 141,
		142, 5, 117, 0, 0, 142, 143, 5, 100, 0, 0, 143, 144, 5, 101, 0, 0, 144,
		14, 1, 0, 0, 0, 145, 146, 5, 115, 0, 0, 146, 147, 5, 116, 0, 0, 147, 148,
		5, 114, 0, 0, 148, 149, 5, 117, 0, 0, 149, 150, 5, 99, 0, 0, 150, 151,
		5, 116, 0, 0, 151, 16, 1, 0, 0, 0, 152, 153, 5, 101, 0, 0, 153, 154, 5,
		110, 0, 0, 154, 155, 5, 117, 0, 0, 155, 156, 5, 109, 0, 0, 156, 18, 1,
		0, 0, 0, 157, 158, 5, 111, 0, 0, 158, 159, 5, 112, 0, 0, 159, 160, 5, 116,
		0, 0, 160, 161, 5, 105, 0, 0, 161, 162, 5, 111, 0, 0, 162, 163, 5, 110,
		0, 0, 163, 164, 5, 97, 0, 0, 164, 165, 5, 108, 0, 0, 165, 20, 1, 0, 0,
		0, 166, 167, 5, 114, 0, 0, 167, 168, 5, 101, 0, 0, 168, 169, 5, 113, 0,
		0, 169, 170, 5, 117, 0, 0, 170, 171, 5, 105, 0, 0, 171, 172, 5, 114, 0,
		0, 172, 173, 5, 101, 0, 0, 173, 174, 5, 100, 0, 0, 174, 22, 1, 0, 0, 0,
		175, 176, 5, 118, 0, 0, 176, 177, 5, 101, 0, 0, 177, 178, 5, 99, 0, 0,
		178, 179, 5, 116, 0, 0, 179, 180, 5, 111, 0, 0, 180, 181, 5, 114, 0, 0,
		181, 24, 1, 0, 0, 0, 182, 183, 5, 109, 0, 0, 183, 184, 5, 97, 0, 0, 184,
		185, 5, 112, 0, 0, 185, 26, 1, 0, 0, 0, 186, 187, 5, 117, 0, 0, 187, 188,
		5, 110, 0, 0, 188, 189, 5, 115, 0, 0, 189, 190, 5, 105, 0, 0, 190, 191,
		5, 103, 0, 0, 191, 192, 5, 110, 0, 0, 192, 193, 5, 101, 0, 0, 193, 194,
		5, 100, 0, 0, 194, 195, 5, 32, 0, 0, 195, 196, 5, 108, 0, 0, 196, 197,
		5, 111, 0, 0, 197, 198, 5, 110, 0, 0, 198, 199, 5, 103, 0, 0, 199, 28,
		1, 0, 0, 0, 200, 201, 5, 108, 0, 0, 201, 202, 5, 111, 0, 0, 202, 203, 5,
		110, 0, 0, 203, 204, 5, 103, 0, 0, 204, 30, 1, 0, 0, 0, 205, 206, 5, 102,
		0, 0, 206, 207, 5, 108, 0, 0, 207, 208, 5, 111, 0, 0, 208, 209, 5, 97,
		0, 0, 209, 210, 5, 116, 0, 0, 210, 32, 1, 0, 0, 0, 211, 212, 5, 98, 0,
		0, 212, 213, 5, 121, 0, 0, 213, 214, 5, 116, 0, 0, 214, 215, 5, 101, 0,
		0, 215, 34, 1, 0, 0, 0, 216, 217, 5, 117, 0, 0, 217, 218, 5, 110, 0, 0,
		218, 219, 5, 115, 0, 0, 219, 220, 5, 105, 0, 0, 220, 221, 5, 103, 0, 0,
		221, 222, 5, 110, 0, 0, 222, 223, 5, 101, 0, 0, 223, 224, 5, 100, 0, 0,
		224, 225, 5, 32, 0, 0, 225, 226, 5, 105, 0, 0, 226, 227, 5, 110, 0, 0,
		227, 228, 5, 116, 0, 0, 228, 36, 1, 0, 0, 0, 229, 230, 5, 105, 0, 0, 230,
		231, 5, 110, 0, 0, 231, 232, 5, 116, 0, 0, 232, 38, 1, 0, 0, 0, 233, 234,
		5, 100, 0, 0, 234, 235, 5, 111, 0, 0, 235, 236, 5, 117, 0, 0, 236, 237,
		5, 98, 0, 0, 237, 238, 5, 108, 0, 0, 238, 239, 5, 101, 0, 0, 239, 40, 1,
		0, 0, 0, 240, 241, 5, 115, 0, 0, 241, 242, 5, 116, 0, 0, 242, 243, 5, 114,
		0, 0, 243, 244, 5, 105, 0, 0, 244, 245, 5, 110, 0, 0, 245, 246, 5, 103,
		0, 0, 246, 42, 1, 0, 0, 0, 247, 248, 5, 118, 0, 0, 248, 249, 5, 111, 0,
		0, 249, 250, 5, 105, 0, 0, 250, 251, 5, 100, 0, 0, 251, 44, 1, 0, 0, 0,
		252, 253, 5, 98, 0, 0, 253, 254, 5, 111, 0, 0, 254, 255, 5, 111, 0, 0,
		255, 256, 5, 108, 0, 0, 256, 46, 1, 0, 0, 0, 257, 258, 5, 116, 0, 0, 258,
		259, 5, 114, 0, 0, 259, 260, 5, 117, 0, 0, 260, 261, 5, 101, 0, 0, 261,
		48, 1, 0, 0, 0, 262, 263, 5, 102, 0, 0, 263, 264, 5, 97, 0, 0, 264, 265,
		5, 108, 0, 0, 265, 266, 5, 115, 0, 0, 266, 267, 5, 101, 0, 0, 267, 50,
		1, 0, 0, 0, 268, 269, 5, 105, 0, 0, 269, 270, 5, 110, 0, 0, 270, 271, 5,
		116, 0, 0, 271, 272, 5, 101, 0, 0, 272, 273, 5, 114, 0, 0, 273, 274, 5,
		102, 0, 0, 274, 275, 5, 97, 0, 0, 275, 276, 5, 99, 0, 0, 276, 277, 5, 101,
		0, 0, 277, 52, 1, 0, 0, 0, 278, 279, 5, 105, 0, 0, 279, 280, 5, 110, 0,
		0, 280, 54, 1, 0, 0, 0, 281, 282, 5, 111, 0, 0, 282, 283, 5, 117, 0, 0,
		283, 284, 5, 116, 0, 0, 284, 56, 1, 0, 0, 0, 285, 286, 5, 123, 0, 0, 286,
		58, 1, 0, 0, 0, 287, 288, 5, 125, 0, 0, 288, 60, 1, 0, 0, 0, 289, 290,
		5, 91, 0, 0, 290, 62, 1, 0, 0, 0, 291, 292, 5, 93, 0, 0, 292, 64, 1, 0,
		0, 0, 293, 294, 5, 40, 0, 0, 294, 66, 1, 0, 0, 0, 295, 296, 5, 41, 0, 0,
		296, 68, 1, 0, 0, 0, 297, 298, 5, 59, 0, 0, 298, 70, 1, 0, 0, 0, 299, 300,
		5, 58, 0, 0, 300, 72, 1, 0, 0, 0, 301, 302, 5, 44, 0, 0, 302, 74, 1, 0,
		0, 0, 303, 304, 5, 61, 0, 0, 304, 76, 1, 0, 0, 0, 305, 306, 5, 60, 0, 0,
		306, 78, 1, 0, 0, 0, 307, 308, 5, 62, 0, 0, 308, 80, 1, 0, 0, 0, 309, 311,
		5, 34, 0, 0, 310, 312, 9, 0, 0, 0, 311, 310, 1, 0, 0, 0, 312, 313, 1, 0,
		0, 0, 313, 314, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0,
		315, 316, 5, 34, 0, 0, 316, 82, 1, 0, 0, 0, 317, 319, 7, 3, 0, 0, 318,
		317, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 321,
		5, 48, 0, 0, 321, 323, 5, 120, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1,
		0, 0, 0, 323, 325, 1, 0, 0, 0, 324, 326, 7, 4, 0, 0, 325, 324, 1, 0, 0,
		0, 326, 327, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328,
		84, 1, 0, 0, 0, 329, 333, 7, 5, 0, 0, 330, 332, 7, 6, 0, 0, 331, 330, 1,
		0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0,
		0, 334, 86, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 337, 5, 58, 0, 0, 337,
		338, 5, 58, 0, 0, 338, 88, 1, 0, 0, 0, 10, 0, 92, 97, 109, 120, 313, 318,
		322, 327, 333, 2, 6, 0, 0, 0, 1, 0,
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

// LexInit initializes any static state used to implement Lex. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLex(). You can call this function if you wish to initialize the static state ahead
// of time.
func LexInit() {
	staticData := &LexLexerStaticData
	staticData.once.Do(lexLexerInit)
}

// NewLex produces a new lexer instance for the optional input antlr.CharStream.
func NewLex(input antlr.CharStream) *Lex {
	LexInit()
	l := new(Lex)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LexLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lex.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lex tokens.
const (
	LexWS                  = 1
	LexNEW_LINE            = 2
	LexLINE_COMMENT        = 3
	LexBLOCK_COMMENT       = 4
	LexMODULE              = 5
	LexHASH                = 6
	LexINCLUDE             = 7
	LexSTRUCT              = 8
	LexENUM                = 9
	LexOPTIONAL            = 10
	LexREQUIRED            = 11
	LexVECTOR              = 12
	LexMAP                 = 13
	LexULONG               = 14
	LexLONG                = 15
	LexFLOAT               = 16
	LexBYTE                = 17
	LexUINT                = 18
	LexINT                 = 19
	LexDOUBLE              = 20
	LexSTRING              = 21
	LexVOID                = 22
	LexBOOL                = 23
	LexTRUE                = 24
	LexFALSE               = 25
	LexINTERFACE           = 26
	LexIN                  = 27
	LexOUT                 = 28
	LexLEFT_BRACE          = 29
	LexRIGHT_BRACE         = 30
	LexLEFT_BRACKET        = 31
	LexRIGHT_BRACKET       = 32
	LexLEFT_PAREN          = 33
	LexRIGHT_PAREN         = 34
	LexSEMICOLON           = 35
	LexCOLON               = 36
	LexCOMMA               = 37
	LexEQUAL               = 38
	LexLEFT_ANGLE_BRACKET  = 39
	LexRIGHT_ANGLE_BRACKET = 40
	LexSTRING_LITERAL      = 41
	LexNUMBER              = 42
	LexIDENTIFIER          = 43
	LexDOUBLE_COLON        = 44
)
