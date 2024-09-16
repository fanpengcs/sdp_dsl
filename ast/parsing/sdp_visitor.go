// Code generated from ./Sdp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Sdp
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Sdp.
type SdpVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Sdp#sdp.
	VisitSdp(ctx *SdpContext) interface{}

	// Visit a parse tree produced by Sdp#includeStatement.
	VisitIncludeStatement(ctx *IncludeStatementContext) interface{}

	// Visit a parse tree produced by Sdp#moduleStatement.
	VisitModuleStatement(ctx *ModuleStatementContext) interface{}

	// Visit a parse tree produced by Sdp#structStatement.
	VisitStructStatement(ctx *StructStatementContext) interface{}

	// Visit a parse tree produced by Sdp#fieldStatement.
	VisitFieldStatement(ctx *FieldStatementContext) interface{}

	// Visit a parse tree produced by Sdp#serviceStatement.
	VisitServiceStatement(ctx *ServiceStatementContext) interface{}

	// Visit a parse tree produced by Sdp#serviceMethodStatement.
	VisitServiceMethodStatement(ctx *ServiceMethodStatementContext) interface{}

	// Visit a parse tree produced by Sdp#methodParamStatement.
	VisitMethodParamStatement(ctx *MethodParamStatementContext) interface{}

	// Visit a parse tree produced by Sdp#enumStatement.
	VisitEnumStatement(ctx *EnumStatementContext) interface{}

	// Visit a parse tree produced by Sdp#enumField.
	VisitEnumField(ctx *EnumFieldContext) interface{}

	// Visit a parse tree produced by Sdp#enumLastField.
	VisitEnumLastField(ctx *EnumLastFieldContext) interface{}

	// Visit a parse tree produced by Sdp#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by Sdp#vectorType.
	VisitVectorType(ctx *VectorTypeContext) interface{}

	// Visit a parse tree produced by Sdp#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by Sdp#moduleType.
	VisitModuleType(ctx *ModuleTypeContext) interface{}
}
