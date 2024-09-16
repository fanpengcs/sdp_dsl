package ast

import "github.com/isolk/go-lsp/lsp/defines"

// TokenTypes:     []string{"namespace", "type", "struct", "typeParameter", "parameter", "variable", "property", "enumMember", "event", "function", "method", "macro", "keyword", "modifier", "comment", "string", "number", "regexp", "operator", "decorator"},
const (
	TokenType_Namespace = 0
	TokenType_Type      = 1
	TokenType_Struct    = 2
	TokenType_TypeParam = 3
	TokenType_Parameter = 4
	TokenType_Variable  = 5
	TokenType_Property  = 6
	TokenType_Enum      = 7
	TokenType_Event     = 8
	TokenType_Function  = 9
	TokenType_Method    = 10
	TokenType_Macro     = 11
	TokenType_Keyword   = 12
	TokenType_Modifier  = 13
	TokenType_Comment   = 14
	TokenType_String    = 15
	TokenType_Number    = 16
	TokenType_Regexp    = 17
	TokenType_Operator  = 18
	TokenType_Decorator = 19
)

// TokenModifiers: []string{"declaration", "definition", "readonly", "static", "deprecated", "abstract", "async", "modification", "documentation", "defaultLibrary"},
const (
	TokenModifier_None           = 0
	TokenModifier_Declaration    = 0x1
	TokenModifier_Definition     = 0x2
	TokenModifier_ReadOnly       = 0x4
	TokenModifier_Static         = 0x8
	TokenModifier_Deprecated     = 0x10
	TokenModifier_Abstract       = 0x20
	TokenModifier_Async          = 0x40
	TokenModifier_Modification   = 0x80
	TokenModifier_Documentation  = 0x100
	TokenModifier_DefaultLibrary = 0x200
)

const (
	SemanticError_RepeatedStructDefine  = "Repeated struct define"
	SemanticError_RepeatedEnumDefine    = "Repeated enum define"
	SemanticError_RepeatedServiceDefine = "Repeated service define"

	SemanticError_RepeatedStructFieldNumberDefine = "Repeated struct field number define"
	SemanticError_RepeatedStructFieldNameDefine   = "Repeated struct field name define"

	SemanticError_RepeatedEnumFieldNumberDefine = "Repeated enum field number define"
	SemanticError_RepeatedEnumFieldNameDefine   = "Repeated enum field name define"

	SemanticError_RepeatedServiceMethodDefine = "Repeated service method define"

	SemanticError_TypeNotDefine = "Type not define"
)

var (
	Source = "sdp-lsp"

	DiagnosticSeverityError = defines.DiagnosticSeverityError
	DiagnosticSeverityWarn  = defines.DiagnosticSeverityWarning
	DiagnosticSeverityInfo  = defines.DiagnosticSeverityInformation
	DiagnosticSeverityHint  = defines.DiagnosticSeverityHint
)

type TokenData struct {
	Line      int
	Column    int
	Length    int
	TokenType int
	Modifier  int
}

type HoverData struct {
	defines.Range
	Content string
	Module  string
}

// module包括struct、service、enum
type SdpModule struct {
	Range    defines.Range
	Structs  map[string][]SdpPosionInfo
	Services map[string][]SdpPosionInfo
	Enum     map[string][]SdpPosionInfo
}

func NewSdpModule() *SdpModule {
	return &SdpModule{
		Structs:  map[string][]SdpPosionInfo{},
		Services: map[string][]SdpPosionInfo{},
		Enum:     map[string][]SdpPosionInfo{},
	}
}

type SdpStruct struct {
	Fields     map[string][]SdpPosionInfo
	FiledsName map[string][]SdpPosionInfo
}

func NewSdpStruct() *SdpStruct {
	return &SdpStruct{
		Fields:     map[string][]SdpPosionInfo{},
		FiledsName: map[string][]SdpPosionInfo{},
	}
}

type SdpService struct {
	Methods map[string][]SdpPosionInfo
}

func NewSdpService() *SdpService {
	return &SdpService{
		Methods: map[string][]SdpPosionInfo{},
	}
}

type SdpEnum struct {
	LastNumber int
	Fileds     map[int][]SdpPosionInfo
	FiledsName map[string][]SdpPosionInfo
}

func NewSdpEnum() *SdpEnum {
	return &SdpEnum{
		LastNumber: -1,
		Fileds:     map[int][]SdpPosionInfo{},
		FiledsName: map[string][]SdpPosionInfo{},
	}
}

// 主要存储重当复定义时，不同位置的信息
type SdpPosionInfo struct {
	FileName  string
	Line      int
	Column    int
	EndLine   int
	EndColumn int
	Data      interface{} // 当前位置的子类型数据
}

var keywords = []string{
	"module",
	"struct",
	"interface",
	"enum",
	"include",
	"optional",
	"required",
	"vector",
	"map",
	"unsigned long",
	"unsigned int",
	"long",
	"int",
	"string",
	"bool",
	"double",
	"float",
	"void",
	"true",
	"false",
	"null",
	"in",
	"out",
}
