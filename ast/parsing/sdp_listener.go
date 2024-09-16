// Code generated from ./Sdp.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Sdp
import "github.com/antlr4-go/antlr/v4"

// SdpListener is a complete listener for a parse tree produced by Sdp.
type SdpListener interface {
	antlr.ParseTreeListener

	// EnterSdp is called when entering the sdp production.
	EnterSdp(c *SdpContext)

	// EnterIncludeStatement is called when entering the includeStatement production.
	EnterIncludeStatement(c *IncludeStatementContext)

	// EnterModuleStatement is called when entering the moduleStatement production.
	EnterModuleStatement(c *ModuleStatementContext)

	// EnterStructStatement is called when entering the structStatement production.
	EnterStructStatement(c *StructStatementContext)

	// EnterFieldStatement is called when entering the fieldStatement production.
	EnterFieldStatement(c *FieldStatementContext)

	// EnterServiceStatement is called when entering the serviceStatement production.
	EnterServiceStatement(c *ServiceStatementContext)

	// EnterServiceMethodStatement is called when entering the serviceMethodStatement production.
	EnterServiceMethodStatement(c *ServiceMethodStatementContext)

	// EnterMethodParamStatement is called when entering the methodParamStatement production.
	EnterMethodParamStatement(c *MethodParamStatementContext)

	// EnterEnumStatement is called when entering the enumStatement production.
	EnterEnumStatement(c *EnumStatementContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterEnumLastField is called when entering the enumLastField production.
	EnterEnumLastField(c *EnumLastFieldContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterVectorType is called when entering the vectorType production.
	EnterVectorType(c *VectorTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterModuleType is called when entering the moduleType production.
	EnterModuleType(c *ModuleTypeContext)

	// ExitSdp is called when exiting the sdp production.
	ExitSdp(c *SdpContext)

	// ExitIncludeStatement is called when exiting the includeStatement production.
	ExitIncludeStatement(c *IncludeStatementContext)

	// ExitModuleStatement is called when exiting the moduleStatement production.
	ExitModuleStatement(c *ModuleStatementContext)

	// ExitStructStatement is called when exiting the structStatement production.
	ExitStructStatement(c *StructStatementContext)

	// ExitFieldStatement is called when exiting the fieldStatement production.
	ExitFieldStatement(c *FieldStatementContext)

	// ExitServiceStatement is called when exiting the serviceStatement production.
	ExitServiceStatement(c *ServiceStatementContext)

	// ExitServiceMethodStatement is called when exiting the serviceMethodStatement production.
	ExitServiceMethodStatement(c *ServiceMethodStatementContext)

	// ExitMethodParamStatement is called when exiting the methodParamStatement production.
	ExitMethodParamStatement(c *MethodParamStatementContext)

	// ExitEnumStatement is called when exiting the enumStatement production.
	ExitEnumStatement(c *EnumStatementContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitEnumLastField is called when exiting the enumLastField production.
	ExitEnumLastField(c *EnumLastFieldContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitVectorType is called when exiting the vectorType production.
	ExitVectorType(c *VectorTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitModuleType is called when exiting the moduleType production.
	ExitModuleType(c *ModuleTypeContext)
}
