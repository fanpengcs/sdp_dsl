// Code generated from ./Sdp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Sdp
import "github.com/antlr4-go/antlr/v4"

// BaseSdpListener is a complete listener for a parse tree produced by Sdp.
type BaseSdpListener struct{}

var _ SdpListener = &BaseSdpListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSdpListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSdpListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSdpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSdpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSdp is called when production sdp is entered.
func (s *BaseSdpListener) EnterSdp(ctx *SdpContext) {}

// ExitSdp is called when production sdp is exited.
func (s *BaseSdpListener) ExitSdp(ctx *SdpContext) {}

// EnterIncludeStatement is called when production includeStatement is entered.
func (s *BaseSdpListener) EnterIncludeStatement(ctx *IncludeStatementContext) {}

// ExitIncludeStatement is called when production includeStatement is exited.
func (s *BaseSdpListener) ExitIncludeStatement(ctx *IncludeStatementContext) {}

// EnterModuleStatement is called when production moduleStatement is entered.
func (s *BaseSdpListener) EnterModuleStatement(ctx *ModuleStatementContext) {}

// ExitModuleStatement is called when production moduleStatement is exited.
func (s *BaseSdpListener) ExitModuleStatement(ctx *ModuleStatementContext) {}

// EnterStructStatement is called when production structStatement is entered.
func (s *BaseSdpListener) EnterStructStatement(ctx *StructStatementContext) {}

// ExitStructStatement is called when production structStatement is exited.
func (s *BaseSdpListener) ExitStructStatement(ctx *StructStatementContext) {}

// EnterFieldStatement is called when production fieldStatement is entered.
func (s *BaseSdpListener) EnterFieldStatement(ctx *FieldStatementContext) {}

// ExitFieldStatement is called when production fieldStatement is exited.
func (s *BaseSdpListener) ExitFieldStatement(ctx *FieldStatementContext) {}

// EnterServiceStatement is called when production serviceStatement is entered.
func (s *BaseSdpListener) EnterServiceStatement(ctx *ServiceStatementContext) {}

// ExitServiceStatement is called when production serviceStatement is exited.
func (s *BaseSdpListener) ExitServiceStatement(ctx *ServiceStatementContext) {}

// EnterServiceMethodStatement is called when production serviceMethodStatement is entered.
func (s *BaseSdpListener) EnterServiceMethodStatement(ctx *ServiceMethodStatementContext) {}

// ExitServiceMethodStatement is called when production serviceMethodStatement is exited.
func (s *BaseSdpListener) ExitServiceMethodStatement(ctx *ServiceMethodStatementContext) {}

// EnterMethodParamStatement is called when production methodParamStatement is entered.
func (s *BaseSdpListener) EnterMethodParamStatement(ctx *MethodParamStatementContext) {}

// ExitMethodParamStatement is called when production methodParamStatement is exited.
func (s *BaseSdpListener) ExitMethodParamStatement(ctx *MethodParamStatementContext) {}

// EnterEnumStatement is called when production enumStatement is entered.
func (s *BaseSdpListener) EnterEnumStatement(ctx *EnumStatementContext) {}

// ExitEnumStatement is called when production enumStatement is exited.
func (s *BaseSdpListener) ExitEnumStatement(ctx *EnumStatementContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BaseSdpListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BaseSdpListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterEnumLastField is called when production enumLastField is entered.
func (s *BaseSdpListener) EnterEnumLastField(ctx *EnumLastFieldContext) {}

// ExitEnumLastField is called when production enumLastField is exited.
func (s *BaseSdpListener) ExitEnumLastField(ctx *EnumLastFieldContext) {}

// EnterType is called when production type is entered.
func (s *BaseSdpListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseSdpListener) ExitType(ctx *TypeContext) {}

// EnterVectorType is called when production vectorType is entered.
func (s *BaseSdpListener) EnterVectorType(ctx *VectorTypeContext) {}

// ExitVectorType is called when production vectorType is exited.
func (s *BaseSdpListener) ExitVectorType(ctx *VectorTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseSdpListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseSdpListener) ExitMapType(ctx *MapTypeContext) {}

// EnterModuleType is called when production moduleType is entered.
func (s *BaseSdpListener) EnterModuleType(ctx *ModuleTypeContext) {}

// ExitModuleType is called when production moduleType is exited.
func (s *BaseSdpListener) ExitModuleType(ctx *ModuleTypeContext) {}
