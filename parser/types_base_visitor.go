// Code generated from Types.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Types

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTypesVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTypesVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTypeNonUnion(ctx *TypeNonUnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTypeUnion(ctx *TypeUnionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTypeVar(ctx *TypeVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitPackVar(ctx *PackVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitStruct(ctx *StructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitNamed(ctx *NamedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitAtomicType(ctx *AtomicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTupleTypeListOnlyExpansion(ctx *TupleTypeListOnlyExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTupleTypeListNoExpansion(ctx *TupleTypeListNoExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTupleTypeListExpansion(ctx *TupleTypeListExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTupleTypeListEmpty(ctx *TupleTypeListEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitUnionTypeListOnlyExpansion(ctx *UnionTypeListOnlyExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitUnionTypeListNoExpansion(ctx *UnionTypeListNoExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitUnionTypeListExpansion(ctx *UnionTypeListExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitUnionTypeListEmpty(ctx *UnionTypeListEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitFieldListOnlyExpansion(ctx *FieldListOnlyExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitFieldListNoExpansion(ctx *FieldListNoExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitFieldListExpansion(ctx *FieldListExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitFieldListEmpty(ctx *FieldListEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitFieldNameFixed(ctx *FieldNameFixedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitFieldNameVariable(ctx *FieldNameVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitAtomic(ctx *AtomicContext) interface{} {
	return v.VisitChildren(ctx)
}
