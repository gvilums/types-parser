// Code generated from Types.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Types

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseTypesVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTypesVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitAtomicType(ctx *AtomicTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTypeListOnlyExpansion(ctx *TypeListOnlyExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTypeListNoExpansion(ctx *TypeListNoExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTypeListExpansion(ctx *TypeListExpansionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitTypeListEmpty(ctx *TypeListEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTypesVisitor) VisitAtomic(ctx *AtomicContext) interface{} {
	return v.VisitChildren(ctx)
}
