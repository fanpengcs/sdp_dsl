package formatter

import (
	"fmt"
	"sdp-lsp/ast/parsing"

	ant "github.com/antlr4-go/antlr/v4"
)

func (v *SdpVisitor) VisitSdp(ctx *parsing.SdpContext) string {
	str := ""
	defer func() {
		if r := recover(); r != nil {
			v.hasError = true
			v.error = fmt.Errorf("VisitSdp error: %v", r)
		}
	}()

	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *parsing.IncludeStatementContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 0)
			str += v.VisitIncludeStatement(t)
		case *parsing.ModuleStatementContext:
			str += v.VisitBeforComment(t.GetStart(), ctx.GetParser().GetTokenStream(), 0)
			str += v.VisitModuleStatement(t)
		case *ant.TerminalNodeImpl:
			str += v.VisitBeforComment(t.GetSymbol(), ctx.GetParser().GetTokenStream(), 0)
		}
	}
	return str
}

func (v *SdpVisitor) VisitIncludeStatement(ctx *parsing.IncludeStatementContext) string {
	return fmt.Sprintf("%s%s %s", ctx.HASH().GetText(), ctx.INCLUDE().GetText(), ctx.STRING_LITERAL().GetText())
}

func (v *SdpVisitor) VisitErrorNode(_ ant.ErrorNode) interface{} {
	return nil
}
