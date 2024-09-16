package formatter

import (
	"sdp-lsp/ast/parsing"

	ant "github.com/antlr4-go/antlr/v4"
)

func NewStringSdpVisitor(content string) *SdpVisitor {
	stream := ant.NewCommonTokenStream(parsing.NewLex(ant.NewInputStream(content)), 0)
	sdpParser := parsing.NewSdp(stream)
	sdpParser.RemoveErrorListeners()
	s := SdpVisitor{
		formattedType: map[string]map[string]string{},
		sdpParser:     sdpParser,
	}
	sdpParser.AddErrorListener(&s)
	return &s
}

type SdpVisitor struct {
	parsing.BaseSdpVisitor
	SdpErrorListener

	formattedType    map[string]map[string]string // module->type->content
	sdpParser        *parsing.Sdp
	formattedContent string
}

// 开始格式化解析
func (i *SdpVisitor) Parse() {
	i.clearData()

	sdpContext := i.sdpParser.Sdp()
	i.formattedContent = i.VisitSdp(sdpContext.(*parsing.SdpContext))
	if i.hasError {
		i.clearData()
	}
}

func (i *SdpVisitor) GetFormatCode() (string, error) {
	return i.formattedContent, i.error
}

func (i *SdpVisitor) SearchType(module, typeName string) string {
	if i.formattedType[module] == nil {
		return ""
	}
	return i.formattedType[module][typeName]
}

func (i *SdpVisitor) clearData() {
	i.formattedType = map[string]map[string]string{}
	i.formattedContent = ""
	i.hasError = false
	i.error = nil
}
