package formatter

import (
	"sdp-lsp/ast/parsing"

	ant "github.com/antlr4-go/antlr/v4"
)

func (v *SdpVisitor) VisitType(ctx *parsing.TypeContext) string {
	if v.hasError {
		panic("has error")
	}

	str := ""
	ch := ctx.GetChild(0)
	switch t := ch.(type) {
	case *ant.TerminalNodeImpl:
		str = t.GetText()
	case *parsing.VectorTypeContext:
		str = v.VisitVectorType(t)
	case *parsing.MapTypeContext:
		str = v.VisitMapType(t)
	case *parsing.ModuleTypeContext:
		str = v.VisitModuleType(t)
	}
	return str
}

func (v *SdpVisitor) VisitVectorType(ctx *parsing.VectorTypeContext) string {
	if v.hasError {
		panic("has error")
	}

	str := ""
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *ant.TerminalNodeImpl:
			str += t.GetText()
		case *parsing.TypeContext:
			ch1 := t.GetChild(0)
			// 看子类型是不是基本类型，如果是，就不需要双边加一个空格
			if _, ok := ch1.(*ant.TerminalNodeImpl); ok {
				str += v.VisitType(t)
			} else if _, ok := ch1.(*parsing.ModuleTypeContext); ok {
				str += v.VisitType(t)
			} else {
				str += " " + v.VisitType(t) + " "
			}
		}
	}
	return str
}

func (v *SdpVisitor) VisitMapType(ctx *parsing.MapTypeContext) string {
	if v.hasError {
		panic("has error")
	}

	str := ""
	for i := 0; i < ctx.GetChildCount(); i++ {
		ch := ctx.GetChild(i)
		switch t := ch.(type) {
		case *ant.TerminalNodeImpl:
			str += t.GetText()
		case *parsing.TypeContext:
			ch1 := t.GetChild(0)
			// 看子类型是不是基本类型，如果是，就不需要双边加一个空格
			if _, ok := ch1.(*ant.TerminalNodeImpl); ok {
				str += v.VisitType(t)
			} else if _, ok := ch1.(*parsing.ModuleTypeContext); ok {
				str += v.VisitType(t)
			} else {
				str += " " + v.VisitType(t) + " "
			}
		}
	}
	return str
}

func (v *SdpVisitor) VisitModuleType(ctx *parsing.ModuleTypeContext) string {
	if v.hasError {
		panic("has error")
	}

	return ctx.IDENTIFIER(0).GetText() + "::" + ctx.IDENTIFIER(1).GetText()
}
