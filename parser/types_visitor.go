// Code generated from Types.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Types

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TypesParser.
type TypesVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TypesParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by TypesParser#List.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by TypesParser#Tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by TypesParser#AtomicType.
	VisitAtomicType(ctx *AtomicTypeContext) interface{}

	// Visit a parse tree produced by TypesParser#TypeListOnlyExpansion.
	VisitTypeListOnlyExpansion(ctx *TypeListOnlyExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#TypeListNoExpansion.
	VisitTypeListNoExpansion(ctx *TypeListNoExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#TypeListExpansion.
	VisitTypeListExpansion(ctx *TypeListExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#TypeListEmpty.
	VisitTypeListEmpty(ctx *TypeListEmptyContext) interface{}

	// Visit a parse tree produced by TypesParser#atomic.
	VisitAtomic(ctx *AtomicContext) interface{}
}
