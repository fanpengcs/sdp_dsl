package formatter

import (
	"fmt"
	"sdp-lsp/ast/parsing"
	"strings"

	ant "github.com/antlr4-go/antlr/v4"
)

func (v *SdpVisitor) VisitStructStatement(ctx *parsing.StructStatementContext) string {
	if v.hasError {
		panic("has error")
	}

	maxInfo := v.CaculateSpaceNumInStruct(ctx)
	str := fmt.Sprintf("%s %s\n%s", ctx.STRUCT(), ctx.IDENTIFIER(), ctx.LEFT_BRACE())
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *parsing.FieldStatementContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 4)
			str += v.VisitFieldStatement(t, maxInfo)
		case *ant.TerminalNodeImpl:
			if t.GetSymbol().GetTokenType() == parsing.SdpRIGHT_BRACE {
				str += v.VisitBeforComment(t.GetSymbol(), ctx.GetParser().GetTokenStream(), 4)
			}
		default:
			panic(fmt.Sprintf("VisitStructStatement error: %v", t))
		}
	}
	str += fmt.Sprintf("%v%v", ctx.RIGHT_BRACE(), ctx.SEMICOLON())
	module := ctx.GetParent().(*parsing.ModuleStatementContext).IDENTIFIER().GetText()
	if v.formattedType[module] == nil {
		v.formattedType[module] = map[string]string{}
	}
	v.formattedType[module][ctx.IDENTIFIER().GetText()] = str
	return str
}

func (v *SdpVisitor) VisitFieldStatement(ctx *parsing.FieldStatementContext, maxInfo FiledMaxInfo) string {
	if v.hasError {
		panic("has error")
	}

	str := fmt.Sprintf("    %s", ctx.NUMBER(0))
	str += strings.Repeat(" ", maxInfo.MaxNumberLen-len(ctx.NUMBER(0).GetText()))

	if ctx.OPTIONAL() != nil {
		str += fmt.Sprintf(" %s ", ctx.OPTIONAL())
	} else {
		str += fmt.Sprintf(" %s ", ctx.REQUIRED())
	}

	typeStr := v.VisitType(ctx.Type_().(*parsing.TypeContext))
	str += fmt.Sprintf("%s%s", typeStr, strings.Repeat(" ", maxInfo.MaxTypeLen-len(typeStr)))

	str += fmt.Sprintf(" %s", ctx.IDENTIFIER())
	if eq := ctx.EQUAL(); eq != nil {
		txt := ""
		if ctx.NUMBER(1) != nil {
			txt = ctx.NUMBER(1).GetText()
		} else if ctx.STRING_LITERAL() != nil {
			txt = ctx.STRING_LITERAL().GetText()
		} else if ctx.TRUE() != nil {
			txt = ctx.TRUE().GetText()
		} else if ctx.FALSE() != nil {
			txt = ctx.FALSE().GetText()
		}
		str += fmt.Sprintf(" = %s", txt)
	}
	str += ctx.SEMICOLON().GetText()

	rightComment := v.VisitRightComment(ctx.SEMICOLON().GetSymbol(), ctx.GetParser().GetTokenStream())
	if rightComment != "" {
		l := maxInfo.MaxNameLen - len(ctx.IDENTIFIER().GetText())
		if ctx.EQUAL() != nil {
			if ctx.NUMBER(1) != nil {
				l -= len(ctx.NUMBER(1).GetText())
			} else if ctx.STRING_LITERAL() != nil {
				l -= len(ctx.STRING_LITERAL().GetText())
			} else if ctx.TRUE() != nil {
				l -= len(ctx.TRUE().GetText())
			} else if ctx.FALSE() != nil {
				l -= len(ctx.FALSE().GetText())
			}
			l -= 3
			str += strings.Repeat(" ", l)
		} else {
			str += strings.Repeat(" ", l)
		}
		str += " " + rightComment
	}
	// str += " " + v.VisitRightComment(ctx.SEMICOLON().GetSymbol(), ctx.GetParser().GetTokenStream())
	return str
}

type FiledMaxInfo struct {
	MaxNumberLen int
	MaxTypeLen   int
	MaxNameLen   int
}

func (v *SdpVisitor) CaculateSpaceNumInStruct(ctx *parsing.StructStatementContext) FiledMaxInfo {
	if v.hasError {
		panic("has error")
	}

	fileds := ctx.AllFieldStatement()
	maxInfo := FiledMaxInfo{}
	for _, filed := range fileds {
		numberLen := len(filed.NUMBER(0).GetText())
		if numberLen > maxInfo.MaxNumberLen {
			maxInfo.MaxNumberLen = numberLen
		}

		typeLen := len(v.VisitType(filed.Type_().(*parsing.TypeContext)))
		if typeLen > maxInfo.MaxTypeLen {
			maxInfo.MaxTypeLen = typeLen
		}

		nameLen := len(filed.IDENTIFIER().GetText())
		if filed.EQUAL() != nil {
			defaultLen := 0
			if filed.NUMBER(1) != nil {
				defaultLen = len(filed.NUMBER(1).GetText())
			} else if filed.STRING_LITERAL() != nil {
				defaultLen = len(filed.STRING_LITERAL().GetText())
			} else if filed.TRUE() != nil {
				defaultLen = len(filed.TRUE().GetText())
			} else if filed.FALSE() != nil {
				defaultLen = len(filed.FALSE().GetText())
			}
			nameLen += defaultLen + 3 // 3 is for " = "
		}
		if nameLen > maxInfo.MaxNameLen {
			maxInfo.MaxNameLen = nameLen
		}
	}
	return maxInfo
}
