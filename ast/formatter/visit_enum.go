package formatter

import (
	"sdp-lsp/ast/parsing"
	"strings"

	ant "github.com/antlr4-go/antlr/v4"
)

func (v *SdpVisitor) VisitEnumStatement(ctx *parsing.EnumStatementContext) string {
	if v.hasError {
		panic("has error")
	}

	maxInfo := v.CaculateSpaceNumInEnum(ctx)
	str := ctx.ENUM().GetText() + " " + ctx.IDENTIFIER().GetText() + "\n{"
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *parsing.EnumFieldContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 4)
			str += v.VisitEnumField(t, maxInfo)
		case *ant.TerminalNodeImpl:
			if t.GetSymbol().GetTokenType() == parsing.SdpRIGHT_BRACE {
				str += v.VisitBeforComment(t.GetSymbol(), ctx.GetParser().GetTokenStream(), 4)
			}
		}
	}
	str += "};"
	return str
}

func (v *SdpVisitor) VisitEnumField(ctx *parsing.EnumFieldContext, maxInfo FiledMaxInfo) string {
	if v.hasError {
		panic("has error")
	}

	str := "    " + ctx.IDENTIFIER().GetText()
	if ctx.NUMBER() != nil {
		str += strings.Repeat(" ", maxInfo.MaxNameLen-len(ctx.IDENTIFIER().GetText()))
		str += " = " + ctx.NUMBER().GetText()
		str += ","
		comment := v.VisitRightComment(ctx.GetStop(), ctx.GetParser().GetTokenStream())
		if comment != "" {
			str += strings.Repeat(" ", maxInfo.MaxNumberLen-len(ctx.NUMBER().GetText()))
			str += " " + comment
		}
	} else {
		str += ","
		str += strings.Repeat(" ", maxInfo.MaxNameLen-len(ctx.IDENTIFIER().GetText())+3)
		comment := v.VisitRightComment(ctx.GetStop(), ctx.GetParser().GetTokenStream())
		if comment != "" {
			str += strings.Repeat(" ", maxInfo.MaxNumberLen)
			str += " " + comment
		}
	}

	return str
}

func (v *SdpVisitor) CaculateSpaceNumInEnum(ctx *parsing.EnumStatementContext) FiledMaxInfo {
	if v.hasError {
		panic("has error")
	}

	maxInfo := FiledMaxInfo{}
	fileds := ctx.AllEnumField()
	for _, filed := range fileds {
		nameLen := len(filed.IDENTIFIER().GetText())
		if nameLen > maxInfo.MaxNameLen {
			maxInfo.MaxNameLen = nameLen
		}

		if filed.NUMBER() == nil {
			continue
		}

		numLen := len(filed.NUMBER().GetText())
		if numLen > maxInfo.MaxNumberLen {
			maxInfo.MaxNumberLen = numLen
		}
	}
	return maxInfo
}
