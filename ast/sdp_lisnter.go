package ast

import (
	"path"
	"sdp-lsp/ast/parsing"
	"strconv"
	"strings"

	ant "github.com/antlr4-go/antlr/v4"
	"github.com/isolk/go-lsp/lsp/defines"
)

// 一个SdpListener对应一个文件
type SdpListener struct {
	*SdpErrorListener
	parsing.BaseSdpListener
	sdpParser *parsing.Sdp

	// 以下是自定义的属性
	fileName    string
	errors      []defines.Diagnostic
	modules     map[string]*SdpModule // 模块名，对应的模块。用于检查是否有重复规则
	tokenDatas  []TokenData
	hoverDatas  []HoverData
	includeList map[string]*defines.Range
}

// 传入的文件名，必须是绝对路径
func NewSdpListener(fileName, content string) *SdpListener {
	stream := ant.NewCommonTokenStream(parsing.NewLex(ant.NewInputStream(content)), 0)
	sdpParser := parsing.NewSdp(stream)
	sdpParser.RemoveErrorListeners()
	l := &SdpListener{
		fileName:         fileName,
		modules:          make(map[string]*SdpModule),
		SdpErrorListener: &SdpErrorListener{},
		sdpParser:        sdpParser,
		includeList:      make(map[string]*defines.Range),
	}
	sdpParser.AddParseListener(l)
	sdpParser.AddErrorListener(l.SdpErrorListener)
	return l
}

func (v *SdpListener) Parse() {
	v.clear()
	v.sdpParser.Sdp()
}

func (v *SdpListener) GetSyntaxErrors() []defines.Diagnostic {
	return append(v.errors, v.SdpErrorListener.errors...)
}

func (v *SdpListener) GetTokenDatas() []TokenData {
	return v.tokenDatas
}

func (v *SdpListener) getHoverDatas() []HoverData {
	return v.hoverDatas
}

func (v *SdpListener) clear() {
	v.errors = nil
	v.tokenDatas = nil
	v.SdpErrorListener.errors = nil
	v.modules = make(map[string]*SdpModule)
	v.includeList = make(map[string]*defines.Range)
	v.sdpParser.GetInputStream().Seek(0)
}

func (v *SdpListener) getIncludeList() []string {
	list := make([]string, len(v.includeList))
	for k := range v.includeList {
		list = append(list, k)
	}
	return list
}

func (v *SdpListener) searchTypeLocation(module, typeName string) []defines.LocationLink {
	if v.modules[module] == nil {
		return nil
	}

	moduleData := v.modules[module]
	if moduleData == nil {
		return nil
	}

	// struct
	res := []defines.LocationLink{}
	for _, v := range moduleData.Structs[typeName] {
		res = append(res, defines.LocationLink{
			TargetUri: defines.DocumentUri(v.FileName),
			TargetRange: defines.Range{
				Start: defines.Position{
					Line:      uint(v.Line),
					Character: uint(v.Column),
				},
				End: defines.Position{
					Line:      uint(v.EndLine),
					Character: uint(v.EndColumn),
				},
			},
			TargetSelectionRange: defines.Range{
				Start: defines.Position{
					Line:      uint(v.Line),
					Character: uint(v.Column),
				},
				End: defines.Position{
					Line:      uint(v.EndLine),
					Character: uint(v.EndColumn),
				},
			},
		})
	}

	// enum
	for _, v := range moduleData.Enum[typeName] {
		res = append(res, defines.LocationLink{
			TargetUri: defines.DocumentUri(v.FileName),
			TargetRange: defines.Range{
				Start: defines.Position{
					Line:      uint(v.Line),
					Character: uint(v.Column),
				},
				End: defines.Position{
					Line:      uint(v.EndLine),
					Character: uint(v.EndColumn),
				},
			},
		})
	}
	return res
}

// 引用逻辑处理
func (v *SdpListener) ExitIncludeStatement(ctx *parsing.IncludeStatementContext) {
	// 处理任何逻辑失败时，直接忽略当前include规则
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	fileText := ctx.STRING_LITERAL().GetText()
	includeFileName := fileText[1 : len(fileText)-1]
	baseDir := v.fileName[:strings.LastIndex(v.fileName, "/")+1]
	absIncludeFileName := baseDir + includeFileName
	absIncludeFileName = path.Clean(absIncludeFileName)

	// 引用过了，忽略
	if v.includeList[absIncludeFileName] != nil {
		return
	}

	v.includeList[absIncludeFileName] = &defines.Range{
		Start: defines.Position{
			Line:      uint(ctx.STRING_LITERAL().GetSymbol().GetLine()),
			Character: uint(ctx.STRING_LITERAL().GetSymbol().GetColumn()),
		},
		End: defines.Position{
			Line:      uint(ctx.GetStop().GetLine()),
			Character: uint(ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText())),
		},
	}

	// 已经存在，添加引用即可
	if manager.getFile(absIncludeFileName) != nil {
		manager.addRef(v.fileName, absIncludeFileName)
		manager.includeManager.recaculateRealRef()
		return
	}

	// 不存在，添加文件
	if manager.addFile(absIncludeFileName) == nil {
		manager.addRef(v.fileName, absIncludeFileName)
		manager.includeManager.recaculateRealRef()
	}
}

// ------------------------以下逻辑，会在进入规则的时候，预处理，避免进入内部规则时，需要判断是否为nil------------------------
func (v *SdpListener) EnterModuleStatement(ctx *parsing.ModuleStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	moduleName := ctx.GetParser().GetTokenStream().LT(2).GetText()
	module := v.modules[moduleName]
	if module == nil {
		mod := NewSdpModule()
		mod.Range = defines.Range{
			Start: defines.Position{
				Line:      uint(ctx.GetStart().GetLine()),
				Character: uint(ctx.GetStart().GetColumn()),
			},
		}
		v.modules[moduleName] = mod
	}
}

func (v *SdpListener) EnterStructStatement(ctx *parsing.StructStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	moduleContext := ctx.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]

	structNameToken := ctx.GetParser().GetTokenStream().LT(2)
	structName := structNameToken.GetText()

	sdpPosInfo := &SdpPosionInfo{
		FileName:  v.fileName,
		Line:      structNameToken.GetLine(),
		Column:    structNameToken.GetColumn(),
		EndLine:   structNameToken.GetLine(),
		EndColumn: structNameToken.GetColumn() + len(structName),
		Data:      NewSdpStruct(),
	}
	sdpModule.Structs[structName] = append(sdpModule.Structs[structName], *sdpPosInfo)
}

func (v *SdpListener) EnterServiceStatement(ctx *parsing.ServiceStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	moduleContext := ctx.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]

	serviceNameToken := ctx.GetParser().GetTokenStream().LT(2)
	serviceName := serviceNameToken.GetText()

	sdpPosInfo := &SdpPosionInfo{
		FileName:  v.fileName,
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
		EndLine:   serviceNameToken.GetLine(),
		EndColumn: serviceNameToken.GetColumn() + len(serviceName),
		Data:      NewSdpService(),
	}
	sdpModule.Services[serviceName] = append(sdpModule.Services[serviceName], *sdpPosInfo)
}

func (v *SdpListener) EnterEnumStatement(ctx *parsing.EnumStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	moduleContext := ctx.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]

	enumNameToken := ctx.GetParser().GetTokenStream().LT(2)
	enumName := enumNameToken.GetText()

	sdpPosInfo := &SdpPosionInfo{
		FileName:  v.fileName,
		Line:      ctx.GetStart().GetLine(),
		Column:    ctx.GetStart().GetColumn(),
		EndLine:   enumNameToken.GetLine(),
		EndColumn: enumNameToken.GetColumn() + len(enumName),
		Data:      NewSdpEnum(),
	}
	sdpModule.Enum[enumName] = append(sdpModule.Enum[enumName], *sdpPosInfo)
}

func (v *SdpListener) ExitModuleStatement(ctx *parsing.ModuleStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	// 语法高亮
	identifier := ctx.IDENTIFIER()
	v.tokenDatas = append(v.tokenDatas, TokenData{
		Line:      identifier.GetSymbol().GetLine(),
		Column:    identifier.GetSymbol().GetColumn(),
		Length:    len(identifier.GetText()),
		TokenType: TokenType_Namespace,
		Modifier:  TokenModifier_None,
	})

	moduleName := identifier.GetText()
	sdpModule := v.modules[moduleName]
	sdpModule.Range.End = defines.Position{
		Line:      uint(ctx.GetStop().GetLine()),
		Character: uint(ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText())),
	}

	// 处理当前子规则的重复错误（enum，struct,interface）
	for _, sdpStructs := range sdpModule.Structs {
		if len(sdpStructs) <= 1 {
			continue
		}
		for i := 0; i < len(sdpStructs); i++ {
			v.errors = append(v.errors, defines.Diagnostic{
				Range: defines.Range{
					Start: defines.Position{
						Line:      uint(sdpStructs[i].Line),
						Character: uint(sdpStructs[i].Column),
					},
					End: defines.Position{
						Line:      uint(sdpStructs[i].EndLine),
						Character: uint(sdpStructs[i].EndColumn),
					},
				},
				Severity: &DiagnosticSeverityError,
				Source:   &Source,
				Message:  SemanticError_RepeatedStructDefine,
			})
		}
	}

	for serviceName, sdpServices := range sdpModule.Services {
		if len(sdpServices) <= 1 {
			continue
		}

		for i := 0; i < len(sdpServices); i++ {
			v.errors = append(v.errors, defines.Diagnostic{
				Range: defines.Range{
					Start: defines.Position{
						Line:      uint(sdpServices[i].Line),
						Character: uint(sdpServices[i].Column),
					},
					End: defines.Position{
						Line:      uint(sdpServices[i].Line),
						Character: uint(sdpServices[i].Column + len(serviceName)),
					},
				},
				Severity: &DiagnosticSeverityError,
				Source:   &Source,
				Message:  SemanticError_RepeatedServiceDefine,
			})
		}
	}

	for enumName, sdpEnums := range sdpModule.Enum {
		if len(sdpEnums) <= 1 {
			continue
		}

		for i := 0; i < len(sdpEnums); i++ {
			v.errors = append(v.errors, defines.Diagnostic{
				Range: defines.Range{
					Start: defines.Position{
						Line:      uint(sdpEnums[i].Line),
						Character: uint(sdpEnums[i].Column),
					},
					End: defines.Position{
						Line:      uint(sdpEnums[i].EndLine),
						Character: uint(sdpEnums[i].EndColumn + len(enumName)),
					},
				},
				Severity: &DiagnosticSeverityError,
				Source:   &Source,
				Message:  SemanticError_RepeatedEnumDefine,
			})
		}
	}
}

func (v *SdpListener) ExitFieldStatement(ctx *parsing.FieldStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	// 语法高亮
	identifier := ctx.IDENTIFIER()
	v.tokenDatas = append(v.tokenDatas, TokenData{
		Line:      identifier.GetSymbol().GetLine(),
		Column:    identifier.GetSymbol().GetColumn(),
		Length:    len(identifier.GetText()),
		TokenType: TokenType_Variable,
		Modifier:  TokenModifier_None,
	})

	// 填充Field数据
	structContext := ctx.GetParent().(*parsing.StructStatementContext)
	moduleContext := structContext.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	structName := structContext.IDENTIFIER().GetText()
	sdpStructs := v.modules[moduleName].Structs[structName]
	structLine := structContext.IDENTIFIER().GetSymbol().GetLine()
	structColumn := structContext.IDENTIFIER().GetSymbol().GetColumn()
	for _, v := range sdpStructs {
		if v.Line == structLine && v.Column == structColumn {
			sdpStruct, _ := v.Data.(*SdpStruct)
			sdpPos := SdpPosionInfo{
				FileName:  v.FileName,
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()),
				Data:      nil,
			}
			sdpStruct.Fields[ctx.NUMBER(0).GetText()] = append(sdpStruct.Fields[ctx.NUMBER(0).GetText()], sdpPos)
			sdpStruct.FiledsName[ctx.IDENTIFIER().GetText()] = append(sdpStruct.FiledsName[ctx.IDENTIFIER().GetText()], sdpPos)
		}
	}
}

func (v *SdpListener) ExitStructStatement(ctx *parsing.StructStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	id := ctx.IDENTIFIER()
	v.tokenDatas = append(v.tokenDatas, TokenData{
		Line:      id.GetSymbol().GetLine(),
		Column:    id.GetSymbol().GetColumn(),
		Length:    len(id.GetText()),
		TokenType: TokenType_Struct,
		Modifier:  TokenModifier_None,
	})

	// 处理当前子规则的重复错误Filed
	moduleContext := ctx.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]
	structName := id.GetText()
	sdpStructs := sdpModule.Structs[structName]
	for i, sdpPosStruct := range sdpStructs {
		startCo := ctx.IDENTIFIER().GetSymbol().GetColumn()
		startLine := ctx.IDENTIFIER().GetSymbol().GetLine()
		if sdpPosStruct.Column != startCo || sdpPosStruct.Line != startLine || sdpPosStruct.EndLine != id.GetSymbol().GetLine() {
			continue
		}
		sdpStruct := sdpPosStruct.Data.(*SdpStruct)
		for _, fileds := range sdpStruct.Fields {
			if len(fileds) <= 1 {
				continue
			}

			for i := 0; i < len(fileds); i++ {
				v.errors = append(v.errors, defines.Diagnostic{
					Range: defines.Range{
						Start: defines.Position{
							Line:      uint(fileds[i].Line),
							Character: uint(fileds[i].Column),
						},
						End: defines.Position{
							Line:      uint(fileds[i].EndLine),
							Character: uint(fileds[i].EndColumn),
						},
					},
					Severity: &DiagnosticSeverityError,
					Source:   &Source,
					Message:  SemanticError_RepeatedStructFieldNumberDefine,
				})
			}
		}

		for _, fileds := range sdpStruct.FiledsName {
			if len(fileds) <= 1 {
				continue
			}

			for i := 0; i < len(fileds); i++ {
				v.errors = append(v.errors, defines.Diagnostic{
					Range: defines.Range{
						Start: defines.Position{
							Line:      uint(fileds[i].Line),
							Character: uint(fileds[i].Column),
						},
						End: defines.Position{
							Line:      uint(fileds[i].EndLine),
							Character: uint(fileds[i].EndColumn),
						},
					},
					Severity: &DiagnosticSeverityError,
					Source:   &Source,
					Message:  SemanticError_RepeatedStructFieldNameDefine,
				})
			}
		}

		start := ctx.GetStart().GetTokenIndex()
		end := ctx.GetStop().GetTokenIndex()
		originText := ctx.GetParser().GetTokenStream().GetTextFromInterval(ant.Interval{Start: start, Stop: end})
		sdpStructs[i].Data = originText
	}
}

func (v *SdpListener) ExitServiceStatement(ctx *parsing.ServiceStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	// 语法高亮
	id := ctx.IDENTIFIER()
	v.tokenDatas = append(v.tokenDatas, TokenData{
		Line:      id.GetSymbol().GetLine(),
		Column:    id.GetSymbol().GetColumn(),
		Length:    len(id.GetText()),
		TokenType: TokenType_Type,
		Modifier:  TokenModifier_None,
	})

	// 处理当前子规则的重复错误（method）
	moduleContext := ctx.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]
	serviceName := id.GetText()
	sdpServices := sdpModule.Services[serviceName]
	for _, sdpPosService := range sdpServices {
		if sdpPosService.Column != ctx.GetStart().GetColumn() || sdpPosService.Line != ctx.GetStart().GetLine() || sdpPosService.EndLine != id.GetSymbol().GetLine() {
			continue
		}
		sdpService := sdpPosService.Data.(*SdpService)
		for _, method := range sdpService.Methods {
			if len(method) <= 1 {
				continue
			}

			for i := 0; i < len(method); i++ {
				v.errors = append(v.errors, defines.Diagnostic{
					Range: defines.Range{
						Start: defines.Position{
							Line:      uint(method[i].Line),
							Character: uint(method[i].Column),
						},
						End: defines.Position{
							Line:      uint(method[i].EndLine),
							Character: uint(method[i].EndColumn),
						},
					},
					Severity: &DiagnosticSeverityError,
					Source:   &Source,
					Message:  SemanticError_RepeatedServiceMethodDefine,
				})
			}
		}
	}
}

func (v *SdpListener) ExitServiceMethodStatement(ctx *parsing.ServiceMethodStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	id := ctx.IDENTIFIER()
	v.tokenDatas = append(v.tokenDatas, TokenData{
		Line:      id.GetSymbol().GetLine(),
		Column:    id.GetSymbol().GetColumn(),
		Length:    len(id.GetText()),
		TokenType: TokenType_Method,
		Modifier:  TokenModifier_None,
	})

	// 填充Method数据
	serviceContext := ctx.GetParent().(*parsing.ServiceStatementContext)
	moduleContext := serviceContext.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	serviceName := serviceContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]
	sdpServices := sdpModule.Services[serviceName]
	serviceLine := serviceContext.GetStart().GetLine()
	serviceColumn := serviceContext.GetStart().GetColumn()
	for _, v := range sdpServices {
		if v.Line == serviceLine && v.Column == serviceColumn {
			sdpService, _ := v.Data.(*SdpService)
			sdpPos := SdpPosionInfo{
				FileName:  v.FileName,
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()),
				Data:      nil,
			}
			sdpService.Methods[ctx.IDENTIFIER().GetText()] = append(sdpService.Methods[ctx.IDENTIFIER().GetText()], sdpPos)
		}
	}
}

func (v *SdpListener) ExitEnumField(ctx *parsing.EnumFieldContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	// 语法高亮
	identifier := ctx.IDENTIFIER()
	v.tokenDatas = append(v.tokenDatas, TokenData{
		Line:      identifier.GetSymbol().GetLine(),
		Column:    identifier.GetSymbol().GetColumn(),
		Length:    len(identifier.GetText()),
		TokenType: TokenType_Enum,
		Modifier:  TokenModifier_None,
	})

	// 填充Enum数据
	enumContext := ctx.GetParent().(*parsing.EnumStatementContext)
	moduleContext := enumContext.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	enumName := enumContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]
	sdpEnums := sdpModule.Enum[enumName]
	enumLine := enumContext.GetStart().GetLine()
	enumColumn := enumContext.GetStart().GetColumn()
	for _, v := range sdpEnums {
		if v.Line == enumLine && v.Column == enumColumn {
			sdpEnum, _ := v.Data.(*SdpEnum)
			sdpPos := SdpPosionInfo{
				FileName:  v.FileName,
				Line:      ctx.GetStart().GetLine(),
				Column:    ctx.GetStart().GetColumn(),
				EndLine:   ctx.GetStop().GetLine(),
				EndColumn: ctx.GetStop().GetColumn() + len(ctx.GetStop().GetText()),
				Data:      nil,
			}
			num := 0
			if ctx.NUMBER() == nil {
				num = sdpEnum.LastNumber + 1
			} else {
				num = parseNumberValue(ctx.NUMBER().GetText())
			}
			sdpEnum.Fileds[num] = append(sdpEnum.Fileds[num], sdpPos)

			name := ctx.IDENTIFIER().GetText()
			sdpEnum.FiledsName[name] = append(sdpEnum.FiledsName[name], sdpPos)
			sdpEnum.LastNumber = num
		}
	}
}

func parseNumberValue(numStr string) int {
	if strings.HasPrefix(numStr, "0x") {
		val, _ := strconv.ParseInt(numStr[2:], 16, 64)
		return int(val)
	}
	val, _ := strconv.Atoi(numStr)
	return val
}

func (v *SdpListener) ExitEnumStatement(ctx *parsing.EnumStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	id := ctx.IDENTIFIER()
	v.tokenDatas = append(v.tokenDatas, TokenData{
		Line:      id.GetSymbol().GetLine(),
		Column:    id.GetSymbol().GetColumn(),
		Length:    len(id.GetText()),
		TokenType: TokenType_Enum,
		Modifier:  TokenModifier_None,
	})

	// 处理当前子规则的重复错误（enum）
	moduleContext := ctx.GetParent().(*parsing.ModuleStatementContext)
	moduleName := moduleContext.IDENTIFIER().GetText()
	sdpModule := v.modules[moduleName]
	enumName := id.GetText()
	sdpEnums := sdpModule.Enum[enumName]
	for i, sdpPosEnum := range sdpEnums {
		if sdpPosEnum.Column != ctx.GetStart().GetColumn() || sdpPosEnum.Line != ctx.GetStart().GetLine() || sdpPosEnum.EndLine != id.GetSymbol().GetLine() {
			continue
		}
		sdpEnum := sdpPosEnum.Data.(*SdpEnum)
		for _, enumValue := range sdpEnum.Fileds {
			if len(enumValue) <= 1 {
				continue
			}

			for i := 0; i < len(enumValue); i++ {
				v.errors = append(v.errors, defines.Diagnostic{
					Range: defines.Range{
						Start: defines.Position{
							Line:      uint(enumValue[i].Line),
							Character: uint(enumValue[i].Column),
						},
						End: defines.Position{
							Line:      uint(enumValue[i].EndLine),
							Character: uint(enumValue[i].EndColumn),
						},
					},
					Severity: &DiagnosticSeverityError,
					Source:   &Source,
					Message:  SemanticError_RepeatedEnumFieldNumberDefine,
				})
			}
		}

		for _, enumName := range sdpEnum.FiledsName {
			if len(enumName) <= 1 {
				continue
			}

			for i := 0; i < len(enumName); i++ {
				v.errors = append(v.errors, defines.Diagnostic{
					Range: defines.Range{
						Start: defines.Position{
							Line:      uint(enumName[i].Line),
							Character: uint(enumName[i].Column),
						},
						End: defines.Position{
							Line:      uint(enumName[i].EndLine),
							Character: uint(enumName[i].EndColumn),
						},
					},
					Severity: &DiagnosticSeverityError,
					Source:   &Source,
					Message:  SemanticError_RepeatedEnumFieldNameDefine,
				})
			}
		}

		start := ctx.GetStart().GetTokenIndex()
		end := ctx.GetStop().GetTokenIndex()
		originText := ctx.GetParser().GetTokenStream().GetTextFromInterval(ant.Interval{Start: start, Stop: end})
		sdpEnums[i].Data = originText
	}
}

func (s *SdpListener) ExitType(ctx *parsing.TypeContext) {
	id := ctx.IDENTIFIER()
	if id == nil {
		// 不是基本类型，不处理
		return
	}
	typeName := id.GetText()

	// 处理语法高亮
	s.tokenDatas = append(s.tokenDatas, TokenData{
		Line:      id.GetSymbol().GetLine(),
		Column:    id.GetSymbol().GetColumn(),
		Length:    len(typeName),
		TokenType: TokenType_Type,
		Modifier:  TokenModifier_None,
	})

	// 处理悬停提示
	moduleName := ""
	for parent := ctx.GetParent(); parent != nil; parent = parent.GetParent() {
		moduleContext, ok := parent.(*parsing.ModuleStatementContext)
		if ok {
			moduleName = moduleContext.IDENTIFIER().GetText()
			break
		}
	}
	s.hoverDatas = append(s.hoverDatas, HoverData{
		Range: defines.Range{
			Start: defines.Position{
				Line:      uint(id.GetSymbol().GetLine()),
				Character: uint(id.GetSymbol().GetColumn()),
			},
			End: defines.Position{
				Line:      uint(id.GetSymbol().GetLine()),
				Character: uint(id.GetSymbol().GetColumn() + len(typeName)),
			},
		},
		Content: typeName,
		Module:  moduleName,
	})

	// 处理type找不到的错误，先从当前文件找，找不到再从引用文件找
	res := s.searchTypeLocation(moduleName, typeName)
	if len(res) != 0 {
		return
	}

	// 从引用文件找
	res = manager.sdpFileManager.searchTypeLocation(moduleName, typeName)
	if len(res) != 0 {
		for _, v := range res {
			if manager.getRealRef(s.fileName)[string(v.TargetUri)] {
				return
			}
		}
	}

	s.errors = append(s.errors, defines.Diagnostic{
		Range: defines.Range{
			Start: defines.Position{
				Line:      uint(id.GetSymbol().GetLine()),
				Character: uint(id.GetSymbol().GetColumn()),
			},
			End: defines.Position{
				Line:      uint(id.GetSymbol().GetLine()),
				Character: uint(id.GetSymbol().GetColumn() + len(typeName)),
			},
		},
		Severity: &DiagnosticSeverityError,
		Source:   &Source,
		Message:  SemanticError_TypeNotDefine,
	})
}

func (s *SdpListener) ExitModuleType(ctx *parsing.ModuleTypeContext) {
	moduleToken := ctx.IDENTIFIER(0)
	moduleName := moduleToken.GetText()

	typeNameToken := ctx.IDENTIFIER(1)
	typeName := typeNameToken.GetText()

	// 处理语法高亮
	s.tokenDatas = append(s.tokenDatas, TokenData{
		Line:      typeNameToken.GetSymbol().GetLine(),
		Column:    typeNameToken.GetSymbol().GetColumn(),
		Length:    len(typeName),
		TokenType: TokenType_Type,
		Modifier:  TokenModifier_None,
	})

	s.tokenDatas = append(s.tokenDatas, TokenData{
		Line:      moduleToken.GetSymbol().GetLine(),
		Column:    moduleToken.GetSymbol().GetColumn(),
		Length:    len(moduleName),
		TokenType: TokenType_Namespace,
		Modifier:  TokenModifier_None,
	})

	// 处理悬停提示
	// moduleName := ""
	// for parent := ctx.GetParent(); parent != nil; parent = parent.GetParent() {
	// 	moduleContext, ok := parent.(*parsing.ModuleStatementContext)
	// 	if ok {
	// 		moduleName = moduleContext.IDENTIFIER().GetText()
	// 		break
	// 	}
	// }
	s.hoverDatas = append(s.hoverDatas, HoverData{
		Range: defines.Range{
			Start: defines.Position{
				Line:      uint(moduleToken.GetSymbol().GetLine()),
				Character: uint(moduleToken.GetSymbol().GetColumn()),
			},
			End: defines.Position{
				Line:      uint(typeNameToken.GetSymbol().GetLine()),
				Character: uint(typeNameToken.GetSymbol().GetColumn() + len(typeName)),
			},
		},
		Content: typeName,
		Module:  moduleName,
	})

	// 处理type找不到的错误，先从当前文件找，找不到再从引用文件找
	res := s.searchTypeLocation(moduleName, typeName)
	if len(res) != 0 {
		return
	}

	// 从引用文件找
	res = manager.sdpFileManager.searchTypeLocation(moduleName, typeName)
	if len(res) != 0 {
		for _, v := range res {
			if manager.getRealRef(s.fileName)[string(v.TargetUri)] {
				return
			}
		}
	}

	s.errors = append(s.errors, defines.Diagnostic{
		Range: defines.Range{
			Start: defines.Position{
				Line:      uint(moduleToken.GetSymbol().GetLine()),
				Character: uint(moduleToken.GetSymbol().GetColumn()),
			},
			End: defines.Position{
				Line:      uint(typeNameToken.GetSymbol().GetLine()),
				Character: uint(typeNameToken.GetSymbol().GetColumn() + len(typeName)),
			},
		},
		Severity: &DiagnosticSeverityError,
		Source:   &Source,
		Message:  SemanticError_TypeNotDefine,
	})
}

// 仅处理语法高亮
func (v *SdpListener) ExitMethodParamStatement(ctx *parsing.MethodParamStatementContext) {
	defer func() {
		if r := recover(); r != nil {
			// 发现错误，不影响后续处理
			// logs.Println(r)
		}
	}()

	for _, id := range ctx.AllIDENTIFIER() {
		v.tokenDatas = append(v.tokenDatas, TokenData{
			Line:      id.GetSymbol().GetLine(),
			Column:    id.GetSymbol().GetColumn(),
			Length:    len(id.GetText()),
			TokenType: TokenType_Parameter,
			Modifier:  TokenModifier_None,
		})
	}
}
