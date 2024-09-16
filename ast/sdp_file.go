package ast

import (
	"sdp-lsp/ast/formatter"

	"github.com/antlr4-go/antlr/v4"
	"github.com/isolk/go-lsp/lsp/defines"
)

type SdpFile struct {
	fileName string                // 文件名
	content  string                // 文件内容
	lister   *SdpListener          // 语法树
	version  int                   // 版本
	visitor  *formatter.SdpVisitor // 格式化访问器，用于格式化代码
}

func NewSdpFile(fileName, content string) *SdpFile {
	return &SdpFile{
		fileName: fileName,
		content:  content,
		version:  0,
		lister:   NewSdpListener(fileName, content),
		visitor:  formatter.NewStringSdpVisitor(content),
	}
}

// 清除语法树，重新解析
func (s *SdpFile) parse() {
	s.lister.Parse()
	s.visitor.Parse()
}

func (s *SdpFile) getFormatCode() (string, error) {
	return s.visitor.GetFormatCode()
}

func (s *SdpFile) getSyntaxErrors() []defines.Diagnostic {
	return s.lister.GetSyntaxErrors()
}

func (s *SdpFile) getTokenDatas() []TokenData {
	return s.lister.GetTokenDatas()
}

func (s *SdpFile) getHoverData(line, column uint) (content, module string) {
	hoverDatas := s.lister.getHoverDatas()
	for _, v := range hoverDatas {
		if v.Start.Line != line {
			continue
		}
		if v.Start.Character <= column && v.End.Character >= column {
			content = v.Content
			module = v.Module
			break
		}
	}
	return
}

func (s *SdpFile) getIncludeList() []string {
	return s.lister.getIncludeList()
}

func (s *SdpFile) searchType(module, typeName string) string {
	return s.visitor.SearchType(module, typeName)
}

func (s *SdpFile) searchTypeLocation(module, typeName string) []defines.LocationLink {
	return s.lister.searchTypeLocation(module, typeName)
}

func (s *SdpFile) getAllTypes() []TypeContent {
	res := []TypeContent{}
	for moduleName, module := range s.lister.modules {
		for name, _ := range module.Enum {
			res = append(res, TypeContent{Content: name, Module: moduleName, FileName: s.fileName})
		}
		for name, _ := range module.Structs {
			res = append(res, TypeContent{Content: name, Module: moduleName, FileName: s.fileName})
		}
	}
	return res
}

// 根据位置获取内容
func (s *SdpFile) getContent(line, column uint) (content, moduleName string) {
	// 遍历所有的token，找到对应的位置
	stream := s.lister.sdpParser.GetInputStream()
	cStream := stream.(*antlr.CommonTokenStream)
	for _, v := range cStream.GetAllTokens() {
		if v.GetLine() != int(line) {
			continue
		}

		endColumn := v.GetColumn() + len(v.GetText())
		if v.GetColumn() <= int(column) && endColumn >= int(column) {
			content = v.GetText()

			// 获取模块
			for mName, module := range s.lister.modules {
				// 小于开始行
				if line < module.Range.Start.Line {
					continue
				}
				// 等于开始行，但是小于开始字符
				if line == module.Range.Start.Line && column < module.Range.Start.Character {
					continue
				}
				// 大于结束行
				if line > module.Range.End.Line {
					continue
				}
				// 等于结束行，但是大于结束字符
				if line == module.Range.End.Line && column > module.Range.End.Character {
					continue
				}
				moduleName = mName
			}
			break
		}
	}
	return
}

func (s *SdpFile) getIncludeFileName(line, column uint) string {
	for fileName, includeRange := range s.lister.includeList {
		if line < includeRange.Start.Line || line > includeRange.End.Line {
			continue
		}
		if line == includeRange.Start.Line && column < includeRange.Start.Character {
			continue
		}
		if line == includeRange.End.Line && column > includeRange.End.Character {
			continue
		}
		return fileName
	}
	return ""
}
