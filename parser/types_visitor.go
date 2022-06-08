// Code generated from Types.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Types

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by TypesParser.
type TypesVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TypesParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by TypesParser#TypeNonUnion.
	VisitTypeNonUnion(ctx *TypeNonUnionContext) interface{}

	// Visit a parse tree produced by TypesParser#TypeUnion.
	VisitTypeUnion(ctx *TypeUnionContext) interface{}

	// Visit a parse tree produced by TypesParser#List.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by TypesParser#TypeVar.
	VisitTypeVar(ctx *TypeVarContext) interface{}

	// Visit a parse tree produced by TypesParser#PackVar.
	VisitPackVar(ctx *PackVarContext) interface{}

	// Visit a parse tree produced by TypesParser#Tuple.
	VisitTuple(ctx *TupleContext) interface{}

	// Visit a parse tree produced by TypesParser#Struct.
	VisitStruct(ctx *StructContext) interface{}

	// Visit a parse tree produced by TypesParser#Named.
	VisitNamed(ctx *NamedContext) interface{}

	// Visit a parse tree produced by TypesParser#AtomicType.
	VisitAtomicType(ctx *AtomicTypeContext) interface{}

	// Visit a parse tree produced by TypesParser#TupleTypeListOnlyExpansion.
	VisitTupleTypeListOnlyExpansion(ctx *TupleTypeListOnlyExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#TupleTypeListNoExpansion.
	VisitTupleTypeListNoExpansion(ctx *TupleTypeListNoExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#TupleTypeListExpansion.
	VisitTupleTypeListExpansion(ctx *TupleTypeListExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#TupleTypeListEmpty.
	VisitTupleTypeListEmpty(ctx *TupleTypeListEmptyContext) interface{}

	// Visit a parse tree produced by TypesParser#UnionTypeListOnlyExpansion.
	VisitUnionTypeListOnlyExpansion(ctx *UnionTypeListOnlyExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#UnionTypeListNoExpansion.
	VisitUnionTypeListNoExpansion(ctx *UnionTypeListNoExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#UnionTypeListExpansion.
	VisitUnionTypeListExpansion(ctx *UnionTypeListExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#UnionTypeListEmpty.
	VisitUnionTypeListEmpty(ctx *UnionTypeListEmptyContext) interface{}

	// Visit a parse tree produced by TypesParser#FieldListOnlyExpansion.
	VisitFieldListOnlyExpansion(ctx *FieldListOnlyExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#FieldListNoExpansion.
	VisitFieldListNoExpansion(ctx *FieldListNoExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#FieldListExpansion.
	VisitFieldListExpansion(ctx *FieldListExpansionContext) interface{}

	// Visit a parse tree produced by TypesParser#FieldListEmpty.
	VisitFieldListEmpty(ctx *FieldListEmptyContext) interface{}

	// Visit a parse tree produced by TypesParser#FieldNameFixed.
	VisitFieldNameFixed(ctx *FieldNameFixedContext) interface{}

	// Visit a parse tree produced by TypesParser#FieldNameVariable.
	VisitFieldNameVariable(ctx *FieldNameVariableContext) interface{}

	// Visit a parse tree produced by TypesParser#atomic.
	VisitAtomic(ctx *AtomicContext) interface{}
}
