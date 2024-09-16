package formatter

import (
	"fmt"
	"sdp-lsp/ast/parsing"

	ant "github.com/antlr4-go/antlr/v4"
)

func (v *SdpVisitor) VisitModuleStatement(ctx *parsing.ModuleStatementContext) string {
	if v.hasError {
		panic("has error")
	}

	str := fmt.Sprintf("%s %s\n%s", ctx.MODULE().GetText(), ctx.IDENTIFIER().GetText(), ctx.LEFT_BRACE().GetText())
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *parsing.StructStatementContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 0)
			str += v.VisitStructStatement(t)
		case *parsing.ServiceStatementContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 0)
			str += v.VisitServiceStatement(t)
		case *parsing.EnumStatementContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 0)
			str += v.VisitEnumStatement(t)
		case *ant.TerminalNodeImpl:
			if t.GetSymbol().GetTokenType() == parsing.SdpRIGHT_BRACE {
				str += v.VisitBeforComment(t.GetSymbol(), ctx.GetParser().GetTokenStream(), 0)
			}
		}
	}
	str += fmt.Sprintf("%v%v", ctx.RIGHT_BRACE(), ctx.SEMICOLON())
	return str
}
