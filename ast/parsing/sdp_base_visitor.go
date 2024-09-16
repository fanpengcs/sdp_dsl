// Code generated from ./Sdp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Sdp
import "github.com/antlr4-go/antlr/v4"

type BaseSdpVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSdpVisitor) VisitSdp(ctx *SdpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitIncludeStatement(ctx *IncludeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitModuleStatement(ctx *ModuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitStructStatement(ctx *StructStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitFieldStatement(ctx *FieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitServiceStatement(ctx *ServiceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitServiceMethodStatement(ctx *ServiceMethodStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitMethodParamStatement(ctx *MethodParamStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitEnumStatement(ctx *EnumStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitEnumField(ctx *EnumFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitEnumLastField(ctx *EnumLastFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitVectorType(ctx *VectorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSdpVisitor) VisitModuleType(ctx *ModuleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
