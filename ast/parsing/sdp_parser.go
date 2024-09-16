// Code generated from ./Sdp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Sdp
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

type Sdp struct {
	*antlr.BaseParser
}

var SdpParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sdpParserInit() {
	staticData := &SdpParserStaticData
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
		"sdp", "includeStatement", "moduleStatement", "structStatement", "fieldStatement",
		"serviceStatement", "serviceMethodStatement", "methodParamStatement",
		"enumStatement", "enumField", "enumLastField", "type", "vectorType",
		"mapType", "moduleType",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 177, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 5, 0,
		32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 55, 8, 2, 10, 2, 12, 2, 58, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 67, 8, 3, 10, 3, 12, 3, 70, 9, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 81, 8, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 89, 8, 5, 10, 5, 12, 5, 92, 9, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 101, 8, 6, 1, 6, 1, 6, 1, 6, 1, 7, 3, 7,
		107, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 113, 8, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 118, 8, 7, 10, 7, 12, 7, 121, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 127,
		8, 8, 10, 8, 12, 8, 130, 9, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9,
		138, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 159, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 0, 0, 15,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 3, 1, 0, 10,
		11, 2, 0, 24, 25, 41, 42, 1, 0, 27, 28, 188, 0, 33, 1, 0, 0, 0, 2, 44,
		1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6, 62, 1, 0, 0, 0, 8, 74, 1, 0, 0, 0, 10,
		84, 1, 0, 0, 0, 12, 96, 1, 0, 0, 0, 14, 106, 1, 0, 0, 0, 16, 122, 1, 0,
		0, 0, 18, 134, 1, 0, 0, 0, 20, 141, 1, 0, 0, 0, 22, 158, 1, 0, 0, 0, 24,
		160, 1, 0, 0, 0, 26, 165, 1, 0, 0, 0, 28, 172, 1, 0, 0, 0, 30, 32, 3, 2,
		1, 0, 31, 30, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34,
		1, 0, 0, 0, 34, 39, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 38, 3, 4, 2, 0,
		37, 36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1,
		0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 0, 0, 1, 43,
		1, 1, 0, 0, 0, 44, 45, 5, 6, 0, 0, 45, 46, 5, 7, 0, 0, 46, 47, 5, 41, 0,
		0, 47, 3, 1, 0, 0, 0, 48, 49, 5, 5, 0, 0, 49, 50, 5, 43, 0, 0, 50, 56,
		5, 29, 0, 0, 51, 55, 3, 6, 3, 0, 52, 55, 3, 10, 5, 0, 53, 55, 3, 16, 8,
		0, 54, 51, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58,
		1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0,
		58, 56, 1, 0, 0, 0, 59, 60, 5, 30, 0, 0, 60, 61, 5, 35, 0, 0, 61, 5, 1,
		0, 0, 0, 62, 63, 5, 8, 0, 0, 63, 64, 5, 43, 0, 0, 64, 68, 5, 29, 0, 0,
		65, 67, 3, 8, 4, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1,
		0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71,
		72, 5, 30, 0, 0, 72, 73, 5, 35, 0, 0, 73, 7, 1, 0, 0, 0, 74, 75, 5, 42,
		0, 0, 75, 76, 7, 0, 0, 0, 76, 77, 3, 22, 11, 0, 77, 80, 5, 43, 0, 0, 78,
		79, 5, 38, 0, 0, 79, 81, 7, 1, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0,
		0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 5, 35, 0, 0, 83, 9, 1, 0, 0, 0, 84, 85,
		5, 26, 0, 0, 85, 86, 5, 43, 0, 0, 86, 90, 5, 29, 0, 0, 87, 89, 3, 12, 6,
		0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91,
		1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 94, 5, 30, 0, 0,
		94, 95, 5, 35, 0, 0, 95, 11, 1, 0, 0, 0, 96, 97, 3, 22, 11, 0, 97, 98,
		5, 43, 0, 0, 98, 100, 5, 33, 0, 0, 99, 101, 3, 14, 7, 0, 100, 99, 1, 0,
		0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 5, 34, 0, 0,
		103, 104, 5, 35, 0, 0, 104, 13, 1, 0, 0, 0, 105, 107, 7, 2, 0, 0, 106,
		105, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109,
		3, 22, 11, 0, 109, 119, 5, 43, 0, 0, 110, 112, 5, 37, 0, 0, 111, 113, 7,
		2, 0, 0, 112, 111, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0,
		0, 114, 115, 3, 22, 11, 0, 115, 116, 5, 43, 0, 0, 116, 118, 1, 0, 0, 0,
		117, 110, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 15, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5,
		9, 0, 0, 123, 124, 5, 43, 0, 0, 124, 128, 5, 29, 0, 0, 125, 127, 3, 18,
		9, 0, 126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0,
		128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131,
		132, 5, 30, 0, 0, 132, 133, 5, 35, 0, 0, 133, 17, 1, 0, 0, 0, 134, 137,
		5, 43, 0, 0, 135, 136, 5, 38, 0, 0, 136, 138, 5, 42, 0, 0, 137, 135, 1,
		0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 5, 37, 0,
		0, 140, 19, 1, 0, 0, 0, 141, 142, 5, 43, 0, 0, 142, 143, 5, 37, 0, 0, 143,
		21, 1, 0, 0, 0, 144, 159, 5, 18, 0, 0, 145, 159, 5, 14, 0, 0, 146, 159,
		5, 15, 0, 0, 147, 159, 5, 16, 0, 0, 148, 159, 5, 17, 0, 0, 149, 159, 5,
		19, 0, 0, 150, 159, 5, 20, 0, 0, 151, 159, 5, 23, 0, 0, 152, 159, 5, 21,
		0, 0, 153, 159, 5, 22, 0, 0, 154, 159, 5, 43, 0, 0, 155, 159, 3, 28, 14,
		0, 156, 159, 3, 24, 12, 0, 157, 159, 3, 26, 13, 0, 158, 144, 1, 0, 0, 0,
		158, 145, 1, 0, 0, 0, 158, 146, 1, 0, 0, 0, 158, 147, 1, 0, 0, 0, 158,
		148, 1, 0, 0, 0, 158, 149, 1, 0, 0, 0, 158, 150, 1, 0, 0, 0, 158, 151,
		1, 0, 0, 0, 158, 152, 1, 0, 0, 0, 158, 153, 1, 0, 0, 0, 158, 154, 1, 0,
		0, 0, 158, 155, 1, 0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 157, 1, 0, 0, 0,
		159, 23, 1, 0, 0, 0, 160, 161, 5, 12, 0, 0, 161, 162, 5, 39, 0, 0, 162,
		163, 3, 22, 11, 0, 163, 164, 5, 40, 0, 0, 164, 25, 1, 0, 0, 0, 165, 166,
		5, 13, 0, 0, 166, 167, 5, 39, 0, 0, 167, 168, 3, 22, 11, 0, 168, 169, 5,
		37, 0, 0, 169, 170, 3, 22, 11, 0, 170, 171, 5, 40, 0, 0, 171, 27, 1, 0,
		0, 0, 172, 173, 5, 43, 0, 0, 173, 174, 5, 44, 0, 0, 174, 175, 5, 43, 0,
		0, 175, 29, 1, 0, 0, 0, 14, 33, 39, 54, 56, 68, 80, 90, 100, 106, 112,
		119, 128, 137, 158,
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

// SdpInit initializes any static state used to implement Sdp. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSdp(). You can call this function if you wish to initialize the static state ahead
// of time.
func SdpInit() {
	staticData := &SdpParserStaticData
	staticData.once.Do(sdpParserInit)
}

// NewSdp produces a new parser instance for the optional input antlr.TokenStream.
func NewSdp(input antlr.TokenStream) *Sdp {
	SdpInit()
	this := new(Sdp)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SdpParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Sdp.g4"

	return this
}

// Sdp tokens.
const (
	SdpEOF                 = antlr.TokenEOF
	SdpWS                  = 1
	SdpNEW_LINE            = 2
	SdpLINE_COMMENT        = 3
	SdpBLOCK_COMMENT       = 4
	SdpMODULE              = 5
	SdpHASH                = 6
	SdpINCLUDE             = 7
	SdpSTRUCT              = 8
	SdpENUM                = 9
	SdpOPTIONAL            = 10
	SdpREQUIRED            = 11
	SdpVECTOR              = 12
	SdpMAP                 = 13
	SdpULONG               = 14
	SdpLONG                = 15
	SdpFLOAT               = 16
	SdpBYTE                = 17
	SdpUINT                = 18
	SdpINT                 = 19
	SdpDOUBLE              = 20
	SdpSTRING              = 21
	SdpVOID                = 22
	SdpBOOL                = 23
	SdpTRUE                = 24
	SdpFALSE               = 25
	SdpINTERFACE           = 26
	SdpIN                  = 27
	SdpOUT                 = 28
	SdpLEFT_BRACE          = 29
	SdpRIGHT_BRACE         = 30
	SdpLEFT_BRACKET        = 31
	SdpRIGHT_BRACKET       = 32
	SdpLEFT_PAREN          = 33
	SdpRIGHT_PAREN         = 34
	SdpSEMICOLON           = 35
	SdpCOLON               = 36
	SdpCOMMA               = 37
	SdpEQUAL               = 38
	SdpLEFT_ANGLE_BRACKET  = 39
	SdpRIGHT_ANGLE_BRACKET = 40
	SdpSTRING_LITERAL      = 41
	SdpNUMBER              = 42
	SdpIDENTIFIER          = 43
	SdpDOUBLE_COLON        = 44
)

// Sdp rules.
const (
	SdpRULE_sdp                    = 0
	SdpRULE_includeStatement       = 1
	SdpRULE_moduleStatement        = 2
	SdpRULE_structStatement        = 3
	SdpRULE_fieldStatement         = 4
	SdpRULE_serviceStatement       = 5
	SdpRULE_serviceMethodStatement = 6
	SdpRULE_methodParamStatement   = 7
	SdpRULE_enumStatement          = 8
	SdpRULE_enumField              = 9
	SdpRULE_enumLastField          = 10
	SdpRULE_type                   = 11
	SdpRULE_vectorType             = 12
	SdpRULE_mapType                = 13
	SdpRULE_moduleType             = 14
)

// ISdpContext is an interface to support dynamic dispatch.
type ISdpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllIncludeStatement() []IIncludeStatementContext
	IncludeStatement(i int) IIncludeStatementContext
	AllModuleStatement() []IModuleStatementContext
	ModuleStatement(i int) IModuleStatementContext

	// IsSdpContext differentiates from other interfaces.
	IsSdpContext()
}

type SdpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySdpContext() *SdpContext {
	var p = new(SdpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_sdp
	return p
}

func InitEmptySdpContext(p *SdpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_sdp
}

func (*SdpContext) IsSdpContext() {}

func NewSdpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SdpContext {
	var p = new(SdpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_sdp

	return p
}

func (s *SdpContext) GetParser() antlr.Parser { return s.parser }

func (s *SdpContext) EOF() antlr.TerminalNode {
	return s.GetToken(SdpEOF, 0)
}

func (s *SdpContext) AllIncludeStatement() []IIncludeStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIncludeStatementContext); ok {
			len++
		}
	}

	tst := make([]IIncludeStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIncludeStatementContext); ok {
			tst[i] = t.(IIncludeStatementContext)
			i++
		}
	}

	return tst
}

func (s *SdpContext) IncludeStatement(i int) IIncludeStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludeStatementContext); ok {
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

	return t.(IIncludeStatementContext)
}

func (s *SdpContext) AllModuleStatement() []IModuleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModuleStatementContext); ok {
			len++
		}
	}

	tst := make([]IModuleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModuleStatementContext); ok {
			tst[i] = t.(IModuleStatementContext)
			i++
		}
	}

	return tst
}

func (s *SdpContext) ModuleStatement(i int) IModuleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleStatementContext); ok {
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

	return t.(IModuleStatementContext)
}

func (s *SdpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SdpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SdpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterSdp(s)
	}
}

func (s *SdpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitSdp(s)
	}
}

func (s *SdpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitSdp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) Sdp() (localctx ISdpContext) {
	localctx = NewSdpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SdpRULE_sdp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SdpHASH {
		{
			p.SetState(30)
			p.IncludeStatement()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SdpMODULE {
		{
			p.SetState(36)
			p.ModuleStatement()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(SdpEOF)
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

// IIncludeStatementContext is an interface to support dynamic dispatch.
type IIncludeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	INCLUDE() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsIncludeStatementContext differentiates from other interfaces.
	IsIncludeStatementContext()
}

type IncludeStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeStatementContext() *IncludeStatementContext {
	var p = new(IncludeStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_includeStatement
	return p
}

func InitEmptyIncludeStatementContext(p *IncludeStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_includeStatement
}

func (*IncludeStatementContext) IsIncludeStatementContext() {}

func NewIncludeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeStatementContext {
	var p = new(IncludeStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_includeStatement

	return p
}

func (s *IncludeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeStatementContext) HASH() antlr.TerminalNode {
	return s.GetToken(SdpHASH, 0)
}

func (s *IncludeStatementContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(SdpINCLUDE, 0)
}

func (s *IncludeStatementContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SdpSTRING_LITERAL, 0)
}

func (s *IncludeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterIncludeStatement(s)
	}
}

func (s *IncludeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitIncludeStatement(s)
	}
}

func (s *IncludeStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitIncludeStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) IncludeStatement() (localctx IIncludeStatementContext) {
	localctx = NewIncludeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SdpRULE_includeStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(SdpHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(SdpINCLUDE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(SdpSTRING_LITERAL)
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

// IModuleStatementContext is an interface to support dynamic dispatch.
type IModuleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODULE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LEFT_BRACE() antlr.TerminalNode
	RIGHT_BRACE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllStructStatement() []IStructStatementContext
	StructStatement(i int) IStructStatementContext
	AllServiceStatement() []IServiceStatementContext
	ServiceStatement(i int) IServiceStatementContext
	AllEnumStatement() []IEnumStatementContext
	EnumStatement(i int) IEnumStatementContext

	// IsModuleStatementContext differentiates from other interfaces.
	IsModuleStatementContext()
}

type ModuleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleStatementContext() *ModuleStatementContext {
	var p = new(ModuleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_moduleStatement
	return p
}

func InitEmptyModuleStatementContext(p *ModuleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_moduleStatement
}

func (*ModuleStatementContext) IsModuleStatementContext() {}

func NewModuleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleStatementContext {
	var p = new(ModuleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_moduleStatement

	return p
}

func (s *ModuleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleStatementContext) MODULE() antlr.TerminalNode {
	return s.GetToken(SdpMODULE, 0)
}

func (s *ModuleStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *ModuleStatementContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpLEFT_BRACE, 0)
}

func (s *ModuleStatementContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpRIGHT_BRACE, 0)
}

func (s *ModuleStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SdpSEMICOLON, 0)
}

func (s *ModuleStatementContext) AllStructStatement() []IStructStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructStatementContext); ok {
			len++
		}
	}

	tst := make([]IStructStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructStatementContext); ok {
			tst[i] = t.(IStructStatementContext)
			i++
		}
	}

	return tst
}

func (s *ModuleStatementContext) StructStatement(i int) IStructStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructStatementContext); ok {
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

	return t.(IStructStatementContext)
}

func (s *ModuleStatementContext) AllServiceStatement() []IServiceStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceStatementContext); ok {
			len++
		}
	}

	tst := make([]IServiceStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceStatementContext); ok {
			tst[i] = t.(IServiceStatementContext)
			i++
		}
	}

	return tst
}

func (s *ModuleStatementContext) ServiceStatement(i int) IServiceStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceStatementContext); ok {
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

	return t.(IServiceStatementContext)
}

func (s *ModuleStatementContext) AllEnumStatement() []IEnumStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumStatementContext); ok {
			len++
		}
	}

	tst := make([]IEnumStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumStatementContext); ok {
			tst[i] = t.(IEnumStatementContext)
			i++
		}
	}

	return tst
}

func (s *ModuleStatementContext) EnumStatement(i int) IEnumStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumStatementContext); ok {
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

	return t.(IEnumStatementContext)
}

func (s *ModuleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterModuleStatement(s)
	}
}

func (s *ModuleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitModuleStatement(s)
	}
}

func (s *ModuleStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitModuleStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) ModuleStatement() (localctx IModuleStatementContext) {
	localctx = NewModuleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SdpRULE_moduleStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(SdpMODULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(SdpLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67109632) != 0 {
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SdpSTRUCT:
			{
				p.SetState(51)
				p.StructStatement()
			}

		case SdpINTERFACE:
			{
				p.SetState(52)
				p.ServiceStatement()
			}

		case SdpENUM:
			{
				p.SetState(53)
				p.EnumStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(SdpRIGHT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(SdpSEMICOLON)
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

// IStructStatementContext is an interface to support dynamic dispatch.
type IStructStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LEFT_BRACE() antlr.TerminalNode
	RIGHT_BRACE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllFieldStatement() []IFieldStatementContext
	FieldStatement(i int) IFieldStatementContext

	// IsStructStatementContext differentiates from other interfaces.
	IsStructStatementContext()
}

type StructStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructStatementContext() *StructStatementContext {
	var p = new(StructStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_structStatement
	return p
}

func InitEmptyStructStatementContext(p *StructStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_structStatement
}

func (*StructStatementContext) IsStructStatementContext() {}

func NewStructStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructStatementContext {
	var p = new(StructStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_structStatement

	return p
}

func (s *StructStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StructStatementContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SdpSTRUCT, 0)
}

func (s *StructStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *StructStatementContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpLEFT_BRACE, 0)
}

func (s *StructStatementContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpRIGHT_BRACE, 0)
}

func (s *StructStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SdpSEMICOLON, 0)
}

func (s *StructStatementContext) AllFieldStatement() []IFieldStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldStatementContext); ok {
			len++
		}
	}

	tst := make([]IFieldStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldStatementContext); ok {
			tst[i] = t.(IFieldStatementContext)
			i++
		}
	}

	return tst
}

func (s *StructStatementContext) FieldStatement(i int) IFieldStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldStatementContext); ok {
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

	return t.(IFieldStatementContext)
}

func (s *StructStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterStructStatement(s)
	}
}

func (s *StructStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitStructStatement(s)
	}
}

func (s *StructStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitStructStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) StructStatement() (localctx IStructStatementContext) {
	localctx = NewStructStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SdpRULE_structStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(SdpSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(SdpLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SdpNUMBER {
		{
			p.SetState(65)
			p.FieldStatement()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
		p.Match(SdpRIGHT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(SdpSEMICOLON)
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

// IFieldStatementContext is an interface to support dynamic dispatch.
type IFieldStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	Type_() ITypeContext
	IDENTIFIER() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	OPTIONAL() antlr.TerminalNode
	REQUIRED() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsFieldStatementContext differentiates from other interfaces.
	IsFieldStatementContext()
}

type FieldStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldStatementContext() *FieldStatementContext {
	var p = new(FieldStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_fieldStatement
	return p
}

func InitEmptyFieldStatementContext(p *FieldStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_fieldStatement
}

func (*FieldStatementContext) IsFieldStatementContext() {}

func NewFieldStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldStatementContext {
	var p = new(FieldStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_fieldStatement

	return p
}

func (s *FieldStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldStatementContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(SdpNUMBER)
}

func (s *FieldStatementContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SdpNUMBER, i)
}

func (s *FieldStatementContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FieldStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *FieldStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SdpSEMICOLON, 0)
}

func (s *FieldStatementContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(SdpOPTIONAL, 0)
}

func (s *FieldStatementContext) REQUIRED() antlr.TerminalNode {
	return s.GetToken(SdpREQUIRED, 0)
}

func (s *FieldStatementContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SdpEQUAL, 0)
}

func (s *FieldStatementContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SdpSTRING_LITERAL, 0)
}

func (s *FieldStatementContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SdpTRUE, 0)
}

func (s *FieldStatementContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SdpFALSE, 0)
}

func (s *FieldStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterFieldStatement(s)
	}
}

func (s *FieldStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitFieldStatement(s)
	}
}

func (s *FieldStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitFieldStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) FieldStatement() (localctx IFieldStatementContext) {
	localctx = NewFieldStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SdpRULE_fieldStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(SdpNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SdpOPTIONAL || _la == SdpREQUIRED) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(76)
		p.Type_()
	}
	{
		p.SetState(77)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SdpEQUAL {
		{
			p.SetState(78)
			p.Match(SdpEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6597120098304) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(82)
		p.Match(SdpSEMICOLON)
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

// IServiceStatementContext is an interface to support dynamic dispatch.
type IServiceStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTERFACE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LEFT_BRACE() antlr.TerminalNode
	RIGHT_BRACE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllServiceMethodStatement() []IServiceMethodStatementContext
	ServiceMethodStatement(i int) IServiceMethodStatementContext

	// IsServiceStatementContext differentiates from other interfaces.
	IsServiceStatementContext()
}

type ServiceStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceStatementContext() *ServiceStatementContext {
	var p = new(ServiceStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_serviceStatement
	return p
}

func InitEmptyServiceStatementContext(p *ServiceStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_serviceStatement
}

func (*ServiceStatementContext) IsServiceStatementContext() {}

func NewServiceStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceStatementContext {
	var p = new(ServiceStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_serviceStatement

	return p
}

func (s *ServiceStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceStatementContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(SdpINTERFACE, 0)
}

func (s *ServiceStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *ServiceStatementContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpLEFT_BRACE, 0)
}

func (s *ServiceStatementContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpRIGHT_BRACE, 0)
}

func (s *ServiceStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SdpSEMICOLON, 0)
}

func (s *ServiceStatementContext) AllServiceMethodStatement() []IServiceMethodStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceMethodStatementContext); ok {
			len++
		}
	}

	tst := make([]IServiceMethodStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceMethodStatementContext); ok {
			tst[i] = t.(IServiceMethodStatementContext)
			i++
		}
	}

	return tst
}

func (s *ServiceStatementContext) ServiceMethodStatement(i int) IServiceMethodStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceMethodStatementContext); ok {
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

	return t.(IServiceMethodStatementContext)
}

func (s *ServiceStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterServiceStatement(s)
	}
}

func (s *ServiceStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitServiceStatement(s)
	}
}

func (s *ServiceStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitServiceStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) ServiceStatement() (localctx IServiceStatementContext) {
	localctx = NewServiceStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SdpRULE_serviceStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(SdpINTERFACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(SdpLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796109795328) != 0 {
		{
			p.SetState(87)
			p.ServiceMethodStatement()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
		p.Match(SdpRIGHT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(SdpSEMICOLON)
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

// IServiceMethodStatementContext is an interface to support dynamic dispatch.
type IServiceMethodStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	IDENTIFIER() antlr.TerminalNode
	LEFT_PAREN() antlr.TerminalNode
	RIGHT_PAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	MethodParamStatement() IMethodParamStatementContext

	// IsServiceMethodStatementContext differentiates from other interfaces.
	IsServiceMethodStatementContext()
}

type ServiceMethodStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceMethodStatementContext() *ServiceMethodStatementContext {
	var p = new(ServiceMethodStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_serviceMethodStatement
	return p
}

func InitEmptyServiceMethodStatementContext(p *ServiceMethodStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_serviceMethodStatement
}

func (*ServiceMethodStatementContext) IsServiceMethodStatementContext() {}

func NewServiceMethodStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceMethodStatementContext {
	var p = new(ServiceMethodStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_serviceMethodStatement

	return p
}

func (s *ServiceMethodStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceMethodStatementContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ServiceMethodStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *ServiceMethodStatementContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(SdpLEFT_PAREN, 0)
}

func (s *ServiceMethodStatementContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(SdpRIGHT_PAREN, 0)
}

func (s *ServiceMethodStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SdpSEMICOLON, 0)
}

func (s *ServiceMethodStatementContext) MethodParamStatement() IMethodParamStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodParamStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodParamStatementContext)
}

func (s *ServiceMethodStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceMethodStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceMethodStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterServiceMethodStatement(s)
	}
}

func (s *ServiceMethodStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitServiceMethodStatement(s)
	}
}

func (s *ServiceMethodStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitServiceMethodStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) ServiceMethodStatement() (localctx IServiceMethodStatementContext) {
	localctx = NewServiceMethodStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SdpRULE_serviceMethodStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Type_()
	}
	{
		p.SetState(97)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(SdpLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796512448512) != 0 {
		{
			p.SetState(99)
			p.MethodParamStatement()
		}

	}
	{
		p.SetState(102)
		p.Match(SdpRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(SdpSEMICOLON)
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

// IMethodParamStatementContext is an interface to support dynamic dispatch.
type IMethodParamStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllIN() []antlr.TerminalNode
	IN(i int) antlr.TerminalNode
	AllOUT() []antlr.TerminalNode
	OUT(i int) antlr.TerminalNode

	// IsMethodParamStatementContext differentiates from other interfaces.
	IsMethodParamStatementContext()
}

type MethodParamStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodParamStatementContext() *MethodParamStatementContext {
	var p = new(MethodParamStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_methodParamStatement
	return p
}

func InitEmptyMethodParamStatementContext(p *MethodParamStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_methodParamStatement
}

func (*MethodParamStatementContext) IsMethodParamStatementContext() {}

func NewMethodParamStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodParamStatementContext {
	var p = new(MethodParamStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_methodParamStatement

	return p
}

func (s *MethodParamStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodParamStatementContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *MethodParamStatementContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *MethodParamStatementContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SdpIDENTIFIER)
}

func (s *MethodParamStatementContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, i)
}

func (s *MethodParamStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SdpCOMMA)
}

func (s *MethodParamStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SdpCOMMA, i)
}

func (s *MethodParamStatementContext) AllIN() []antlr.TerminalNode {
	return s.GetTokens(SdpIN)
}

func (s *MethodParamStatementContext) IN(i int) antlr.TerminalNode {
	return s.GetToken(SdpIN, i)
}

func (s *MethodParamStatementContext) AllOUT() []antlr.TerminalNode {
	return s.GetTokens(SdpOUT)
}

func (s *MethodParamStatementContext) OUT(i int) antlr.TerminalNode {
	return s.GetToken(SdpOUT, i)
}

func (s *MethodParamStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodParamStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodParamStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterMethodParamStatement(s)
	}
}

func (s *MethodParamStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitMethodParamStatement(s)
	}
}

func (s *MethodParamStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitMethodParamStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) MethodParamStatement() (localctx IMethodParamStatementContext) {
	localctx = NewMethodParamStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SdpRULE_methodParamStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SdpIN || _la == SdpOUT {
		{
			p.SetState(105)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SdpIN || _la == SdpOUT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(108)
		p.Type_()
	}
	{
		p.SetState(109)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SdpCOMMA {
		{
			p.SetState(110)
			p.Match(SdpCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SdpIN || _la == SdpOUT {
			{
				p.SetState(111)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SdpIN || _la == SdpOUT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(114)
			p.Type_()
		}
		{
			p.SetState(115)
			p.Match(SdpIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IEnumStatementContext is an interface to support dynamic dispatch.
type IEnumStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LEFT_BRACE() antlr.TerminalNode
	RIGHT_BRACE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllEnumField() []IEnumFieldContext
	EnumField(i int) IEnumFieldContext

	// IsEnumStatementContext differentiates from other interfaces.
	IsEnumStatementContext()
}

type EnumStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumStatementContext() *EnumStatementContext {
	var p = new(EnumStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_enumStatement
	return p
}

func InitEmptyEnumStatementContext(p *EnumStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_enumStatement
}

func (*EnumStatementContext) IsEnumStatementContext() {}

func NewEnumStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumStatementContext {
	var p = new(EnumStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_enumStatement

	return p
}

func (s *EnumStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumStatementContext) ENUM() antlr.TerminalNode {
	return s.GetToken(SdpENUM, 0)
}

func (s *EnumStatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *EnumStatementContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpLEFT_BRACE, 0)
}

func (s *EnumStatementContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(SdpRIGHT_BRACE, 0)
}

func (s *EnumStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SdpSEMICOLON, 0)
}

func (s *EnumStatementContext) AllEnumField() []IEnumFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumFieldContext); ok {
			len++
		}
	}

	tst := make([]IEnumFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumFieldContext); ok {
			tst[i] = t.(IEnumFieldContext)
			i++
		}
	}

	return tst
}

func (s *EnumStatementContext) EnumField(i int) IEnumFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
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

	return t.(IEnumFieldContext)
}

func (s *EnumStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterEnumStatement(s)
	}
}

func (s *EnumStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitEnumStatement(s)
	}
}

func (s *EnumStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitEnumStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) EnumStatement() (localctx IEnumStatementContext) {
	localctx = NewEnumStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SdpRULE_enumStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(SdpENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(SdpLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SdpIDENTIFIER {
		{
			p.SetState(125)
			p.EnumField()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(131)
		p.Match(SdpRIGHT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(SdpSEMICOLON)
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

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *EnumFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SdpCOMMA, 0)
}

func (s *EnumFieldContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SdpEQUAL, 0)
}

func (s *EnumFieldContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SdpNUMBER, 0)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (s *EnumFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SdpRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SdpEQUAL {
		{
			p.SetState(135)
			p.Match(SdpEQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(SdpNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(139)
		p.Match(SdpCOMMA)
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

// IEnumLastFieldContext is an interface to support dynamic dispatch.
type IEnumLastFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsEnumLastFieldContext differentiates from other interfaces.
	IsEnumLastFieldContext()
}

type EnumLastFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumLastFieldContext() *EnumLastFieldContext {
	var p = new(EnumLastFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_enumLastField
	return p
}

func InitEmptyEnumLastFieldContext(p *EnumLastFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_enumLastField
}

func (*EnumLastFieldContext) IsEnumLastFieldContext() {}

func NewEnumLastFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumLastFieldContext {
	var p = new(EnumLastFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_enumLastField

	return p
}

func (s *EnumLastFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumLastFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *EnumLastFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SdpCOMMA, 0)
}

func (s *EnumLastFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumLastFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumLastFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterEnumLastField(s)
	}
}

func (s *EnumLastFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitEnumLastField(s)
	}
}

func (s *EnumLastFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitEnumLastField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) EnumLastField() (localctx IEnumLastFieldContext) {
	localctx = NewEnumLastFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SdpRULE_enumLastField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(SdpCOMMA)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UINT() antlr.TerminalNode
	ULONG() antlr.TerminalNode
	LONG() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BYTE() antlr.TerminalNode
	INT() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	VOID() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ModuleType() IModuleTypeContext
	VectorType() IVectorTypeContext
	MapType() IMapTypeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) UINT() antlr.TerminalNode {
	return s.GetToken(SdpUINT, 0)
}

func (s *TypeContext) ULONG() antlr.TerminalNode {
	return s.GetToken(SdpULONG, 0)
}

func (s *TypeContext) LONG() antlr.TerminalNode {
	return s.GetToken(SdpLONG, 0)
}

func (s *TypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SdpFLOAT, 0)
}

func (s *TypeContext) BYTE() antlr.TerminalNode {
	return s.GetToken(SdpBYTE, 0)
}

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(SdpINT, 0)
}

func (s *TypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(SdpDOUBLE, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SdpBOOL, 0)
}

func (s *TypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(SdpSTRING, 0)
}

func (s *TypeContext) VOID() antlr.TerminalNode {
	return s.GetToken(SdpVOID, 0)
}

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, 0)
}

func (s *TypeContext) ModuleType() IModuleTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModuleTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModuleTypeContext)
}

func (s *TypeContext) VectorType() IVectorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorTypeContext)
}

func (s *TypeContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SdpRULE_type)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Match(SdpUINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(SdpULONG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Match(SdpLONG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(147)
			p.Match(SdpFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(148)
			p.Match(SdpBYTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(149)
			p.Match(SdpINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(150)
			p.Match(SdpDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(151)
			p.Match(SdpBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(152)
			p.Match(SdpSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(153)
			p.Match(SdpVOID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(154)
			p.Match(SdpIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(155)
			p.ModuleType()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(156)
			p.VectorType()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(157)
			p.MapType()
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

// IVectorTypeContext is an interface to support dynamic dispatch.
type IVectorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VECTOR() antlr.TerminalNode
	LEFT_ANGLE_BRACKET() antlr.TerminalNode
	Type_() ITypeContext
	RIGHT_ANGLE_BRACKET() antlr.TerminalNode

	// IsVectorTypeContext differentiates from other interfaces.
	IsVectorTypeContext()
}

type VectorTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorTypeContext() *VectorTypeContext {
	var p = new(VectorTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_vectorType
	return p
}

func InitEmptyVectorTypeContext(p *VectorTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_vectorType
}

func (*VectorTypeContext) IsVectorTypeContext() {}

func NewVectorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorTypeContext {
	var p = new(VectorTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_vectorType

	return p
}

func (s *VectorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorTypeContext) VECTOR() antlr.TerminalNode {
	return s.GetToken(SdpVECTOR, 0)
}

func (s *VectorTypeContext) LEFT_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SdpLEFT_ANGLE_BRACKET, 0)
}

func (s *VectorTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VectorTypeContext) RIGHT_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SdpRIGHT_ANGLE_BRACKET, 0)
}

func (s *VectorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterVectorType(s)
	}
}

func (s *VectorTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitVectorType(s)
	}
}

func (s *VectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitVectorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) VectorType() (localctx IVectorTypeContext) {
	localctx = NewVectorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SdpRULE_vectorType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(SdpVECTOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(SdpLEFT_ANGLE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Type_()
	}
	{
		p.SetState(163)
		p.Match(SdpRIGHT_ANGLE_BRACKET)
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

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LEFT_ANGLE_BRACKET() antlr.TerminalNode
	AllType_() []ITypeContext
	Type_(i int) ITypeContext
	COMMA() antlr.TerminalNode
	RIGHT_ANGLE_BRACKET() antlr.TerminalNode

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_mapType
	return p
}

func InitEmptyMapTypeContext(p *MapTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_mapType
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) MAP() antlr.TerminalNode {
	return s.GetToken(SdpMAP, 0)
}

func (s *MapTypeContext) LEFT_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SdpLEFT_ANGLE_BRACKET, 0)
}

func (s *MapTypeContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *MapTypeContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
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

	return t.(ITypeContext)
}

func (s *MapTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SdpCOMMA, 0)
}

func (s *MapTypeContext) RIGHT_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(SdpRIGHT_ANGLE_BRACKET, 0)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SdpRULE_mapType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(SdpMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(SdpLEFT_ANGLE_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Type_()
	}
	{
		p.SetState(168)
		p.Match(SdpCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Type_()
	}
	{
		p.SetState(170)
		p.Match(SdpRIGHT_ANGLE_BRACKET)
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

// IModuleTypeContext is an interface to support dynamic dispatch.
type IModuleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	DOUBLE_COLON() antlr.TerminalNode

	// IsModuleTypeContext differentiates from other interfaces.
	IsModuleTypeContext()
}

type ModuleTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleTypeContext() *ModuleTypeContext {
	var p = new(ModuleTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_moduleType
	return p
}

func InitEmptyModuleTypeContext(p *ModuleTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SdpRULE_moduleType
}

func (*ModuleTypeContext) IsModuleTypeContext() {}

func NewModuleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleTypeContext {
	var p = new(ModuleTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SdpRULE_moduleType

	return p
}

func (s *ModuleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SdpIDENTIFIER)
}

func (s *ModuleTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SdpIDENTIFIER, i)
}

func (s *ModuleTypeContext) DOUBLE_COLON() antlr.TerminalNode {
	return s.GetToken(SdpDOUBLE_COLON, 0)
}

func (s *ModuleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.EnterModuleType(s)
	}
}

func (s *ModuleTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SdpListener); ok {
		listenerT.ExitModuleType(s)
	}
}

func (s *ModuleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SdpVisitor:
		return t.VisitModuleType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Sdp) ModuleType() (localctx IModuleTypeContext) {
	localctx = NewModuleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SdpRULE_moduleType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(SdpIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(SdpDOUBLE_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(SdpIDENTIFIER)
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
