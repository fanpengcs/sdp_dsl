package formatter

import (
	"fmt"
	"sdp-lsp/ast/parsing"

	ant "github.com/antlr4-go/antlr/v4"
)

func (v *SdpVisitor) VisitServiceStatement(ctx *parsing.ServiceStatementContext) string {
	if v.hasError {
		panic("has error")
	}

	str := fmt.Sprintf("%s %s\n%s", ctx.INTERFACE().GetText(), ctx.IDENTIFIER().GetText(), ctx.LEFT_BRACE().GetText())
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *parsing.ServiceMethodStatementContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 4)
			str += v.VisitServiceMethodStatement(t)
		case *ant.TerminalNodeImpl:
			if t.GetSymbol().GetTokenType() == parsing.SdpRIGHT_BRACE {
				str += v.VisitBeforComment(t.GetSymbol(), ctx.GetParser().GetTokenStream(), 0)
			}
		}
	}
	str += fmt.Sprintf("%s%s\n", ctx.RIGHT_BRACE(), ctx.SEMICOLON())
	return str
}

func (v *SdpVisitor) VisitServiceMethodStatement(ctx *parsing.ServiceMethodStatementContext) string {
	if v.hasError {
		panic("has error")
	}

	typeStr := fmt.Sprintf("%v", v.VisitType(ctx.Type_().(*parsing.TypeContext)))
	str := fmt.Sprintf("    %s %s%s", typeStr, ctx.IDENTIFIER(), ctx.LEFT_PAREN())
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *parsing.MethodParamStatementContext:
			str += v.VisitMethodParamStatement(t)
		}
	}
	str += fmt.Sprintf("%s%s", ctx.RIGHT_PAREN().GetText(), ctx.SEMICOLON().GetText())
	str += " " + v.VisitRightComment(ctx.SEMICOLON().GetSymbol(), ctx.GetParser().GetTokenStream())
	return str
}

func (v *SdpVisitor) VisitMethodParamStatement(ctx *parsing.MethodParamStatementContext) string {
	if v.hasError {
		panic("has error")
	}

	str := ""
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *parsing.TypeContext:
			str += fmt.Sprintf("%v ", v.VisitType(t))
		case *ant.TerminalNodeImpl:
			if t.GetSymbol().GetTokenType() != parsing.SdpIDENTIFIER {
				str += fmt.Sprintf("%v ", t.GetText())
			} else {
				str += fmt.Sprintf("%v", t.GetText())
			}

		}
	}
	return str
}
