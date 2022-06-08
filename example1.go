// example1.go
package main

import (
	"fmt"

	"example/gvilums/types-parser/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bleenco/bms/protobuf/bms/protobuf/pipelang"
)

type typesVisitor struct{}

func NewTypesVisitor() parser.TypesVisitor {
	return &typesVisitor{}
}

// Required

func (v *typesVisitor) Visit(tree antlr.ParseTree) interface{}            { return tree.Accept(v) }
func (v *typesVisitor) VisitChildren(node antlr.RuleNode) interface{}     { return nil }
func (v *typesVisitor) VisitTerminal(node antlr.TerminalNode) interface{} { return nil }
func (v *typesVisitor) VisitErrorNode(node antlr.ErrorNode) interface{}   { return nil }

// TYPE

func (v *typesVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return ctx.GetValue().Accept(v)
}

func (v *typesVisitor) VisitList(ctx *parser.ListContext) interface{} {
	out := ctx.GetElemType().Accept(v)
	if out == nil {
		return nil
	}
	t := out.(pipelang.PipelangType)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_List{List: &pipelang.List{ElemType: &t}}}
}

func (v *typesVisitor) VisitTuple(ctx *parser.TupleContext) interface{} {
	out := ctx.GetTypes().Accept(v)
	if out == nil {
		return nil
	}
	tl := out.(pipelang.TypeList)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Tuple{Tuple: &pipelang.Tuple{Types: &tl}}}
}

func (v *typesVisitor) VisitUnion(ctx *parser.UnionContext) interface{} {
	out := ctx.GetTypes().Accept(v)
	if out == nil {
		return nil
	}
	tl := out.(pipelang.TypeList)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Union{Union: &pipelang.Union{Types: &tl}}}
}

func (v *typesVisitor) VisitStruct(ctx *parser.StructContext) interface{} {
	out := ctx.GetFields().Accept(v)
	if out == nil {
		return nil
	}
	fl := out.(pipelang.FieldList)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Struct{Struct: &pipelang.Struct{Fields: &fl}}}
}

func (v *typesVisitor) VisitNamed(ctx *parser.NamedContext) interface{} {
	name := ctx.GetName().GetText()
	out := ctx.GetTypes().Accept(v)
	if out == nil {
		return nil
	}
	types := out.(pipelang.TypeList)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Named{Named: &pipelang.NamedType{Name: name, Arguments: &types}}}
}

func (v *typesVisitor) VisitPackVar(ctx *parser.PackVarContext) interface{} {
	name := ctx.GetName().GetText()
	return pipelang.PipelangType{Type: &pipelang.PipelangType_PackVariable{PackVariable: &pipelang.TypePack{Name: name}}}
}

func (v *typesVisitor) VisitTypeVar(ctx *parser.TypeVarContext) interface{} {
	name := ctx.GetName().GetText()
	return pipelang.PipelangType{Type: &pipelang.PipelangType_TypeVariable{TypeVariable: &pipelang.TypeVariable{Name: name}}}
}

func (v *typesVisitor) VisitAtomicType(ctx *parser.AtomicTypeContext) interface{} {
	out := ctx.GetAtom().Accept(v).(pipelang.Atomic)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Atomic{Atomic: out}}
}

// TYPE LIST

func (v *typesVisitor) VisitTupleTypeListOnlyExpansion(ctx *parser.TupleTypeListOnlyExpansionContext) interface{} {
	out := ctx.GetPattern().Accept(v)
	if out == nil {
		return nil
	}
	pattern := out.(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: &pattern}
}

func (v *typesVisitor) VisitTupleTypeListNoExpansion(ctx *parser.TupleTypeListNoExpansionContext) interface{} {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types)+1)
	for i, t := range types {
		out := t.Accept(v)
		if out == nil {
			return nil
		}
		tp := out.(pipelang.PipelangType)
		fixed[i] = &tp
	}
	out := ctx.GetFinalType().Accept(v)
	if out == nil {
		return nil
	}
	tp := out.(pipelang.PipelangType)
	fixed[len(fixed)-1] = &tp
	return pipelang.TypeList{Fixed: fixed, Expansion: nil}
}

func (v *typesVisitor) VisitTupleTypeListExpansion(ctx *parser.TupleTypeListExpansionContext) interface{} {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types))
	for i, t := range types {
		out := t.Accept(v)
		if out == nil {
			return nil
		}
		tp := out.(pipelang.PipelangType)
		fixed[i] = &tp
	}
	out := ctx.GetPattern().Accept(v)
	if out == nil {
		return nil
	}
	pattern := out.(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: fixed, Expansion: &pattern}
}

func (v *typesVisitor) VisitTupleTypeListEmpty(ctx *parser.TupleTypeListEmptyContext) interface{} {
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: nil}
}

// TYPE LIST

func (v *typesVisitor) VisitUnionTypeListOnlyExpansion(ctx *parser.UnionTypeListOnlyExpansionContext) interface{} {
	out := ctx.GetPattern().Accept(v)
	if out == nil {
		return nil
	}
	pattern := out.(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: &pattern}
}

func (v *typesVisitor) VisitUnionTypeListNoExpansion(ctx *parser.UnionTypeListNoExpansionContext) interface{} {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types)+1)
	for i, t := range types {
		out := t.Accept(v)
		if out == nil {
			return nil
		}
		tp := out.(pipelang.PipelangType)
		fixed[i] = &tp
	}
	out := ctx.GetFinalType().Accept(v)
	if out == nil {
		return nil
	}
	tp := out.(pipelang.PipelangType)
	fixed[len(fixed)-1] = &tp
	return pipelang.TypeList{Fixed: fixed, Expansion: nil}
}

func (v *typesVisitor) VisitUnionTypeListExpansion(ctx *parser.UnionTypeListExpansionContext) interface{} {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types))
	for i, t := range types {
		out := t.Accept(v)
		if out == nil {
			return nil
		}
		tp := out.(pipelang.PipelangType)
		fixed[i] = &tp
	}
	out := ctx.GetPattern().Accept(v)
	if out == nil {
		return nil
	}
	pattern := out.(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: fixed, Expansion: &pattern}
}

func (v *typesVisitor) VisitUnionTypeListEmpty(ctx *parser.UnionTypeListEmptyContext) interface{} {
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: nil}
}

// FIELD LIST

func makeField(n interface{}, t interface{}) *pipelang.FieldList_Field {
	if n == nil || t == nil {
		return nil
	}
	tp := t.(pipelang.PipelangType)
	switch np := n.(type) {
	case pipelang.FieldList_Field_Fixed:
		return &pipelang.FieldList_Field{FieldName: &np, FieldType: &tp}
	case pipelang.FieldList_Field_Variable:
		return &pipelang.FieldList_Field{FieldName: &np, FieldType: &tp}
	default:
		panic("unexpected field name type")
	}
}

func makeFieldPattern(n string, t interface{}) *pipelang.FieldList_FieldExpansion {
	if t == nil {
		return nil
	}
	return &pipelang.FieldList_FieldExpansion{NamePack: n, Pattern: t.(*pipelang.PipelangType)}
}

func (v *typesVisitor) VisitFieldListOnlyExpansion(ctx *parser.FieldListOnlyExpansionContext) interface{} {
	name_pattern := ctx.GetNamePattern().GetText()
	type_pattern := ctx.GetTypePattern().Accept(v)
	exp := makeFieldPattern(name_pattern, type_pattern)
	if exp == nil {
		return nil
	}
	return pipelang.FieldList{Fixed: make([]*pipelang.FieldList_Field, 0), Expansion: exp}
}

func (v *typesVisitor) VisitFieldListNoExpansion(ctx *parser.FieldListNoExpansionContext) interface{} {
	names := ctx.GetNames()
	types := ctx.GetTypes()
	if len(names) != len(types) {
		panic("internal error: names list and types list in FieldList have different lengths")
	}
	fixed := make([]*pipelang.FieldList_Field, len(types)+1)
	for i := range types {
		name := names[i].Accept(v)
		out := types[i].Accept(v)
		f := makeField(name, out)
		if f == nil {
			return f
		}
		fixed[i] = f
	}
	f := makeField(ctx.GetFinalName().Accept(v), ctx.GetFinalType().Accept(v))
	if f == nil {
		return nil
	}
	fixed[len(fixed)-1] = f
	return pipelang.FieldList{Fixed: fixed, Expansion: nil}
}

func (v *typesVisitor) VisitFieldListExpansion(ctx *parser.FieldListExpansionContext) interface{} {
	names := ctx.GetNames()
	types := ctx.GetTypes()
	if len(names) != len(types) {
		panic("internal error: names list and types list in FieldList have different lengths")
	}
	fixed := make([]*pipelang.FieldList_Field, len(types))
	for i := range types {
		name := names[i].Accept(v)
		out := types[i].Accept(v)
		f := makeField(name, out)
		if f == nil {
			return nil
		}
		fixed[i] = f
	}
	exp := makeFieldPattern(ctx.GetNamePattern().GetText(), ctx.GetTypePattern().Accept(v))
	if exp == nil {
		return nil
	}
	return pipelang.FieldList{Fixed: fixed, Expansion: exp}
}

func (v *typesVisitor) VisitFieldListEmpty(ctx *parser.FieldListEmptyContext) interface{} {
	return pipelang.FieldList{Fixed: make([]*pipelang.FieldList_Field, 0), Expansion: nil}
}

// FIELD NAME

func (v *typesVisitor) VisitFieldNameFixed(ctx *parser.FieldNameFixedContext) interface{} {
	return pipelang.FieldList_Field_Fixed{Fixed: ctx.GetName().GetText()}
}

func (v *typesVisitor) VisitFieldNameVariable(ctx *parser.FieldNameVariableContext) interface{} {
	return pipelang.FieldList_Field_Variable{Variable: ctx.GetName().GetText()}
}

// ATOMIC

func (v *typesVisitor) VisitAtomic(ctx *parser.AtomicContext) interface{} {
	switch ctx.GetName().GetTokenType() {
	case parser.TypesParserATOMIC_INT32:
		return pipelang.Atomic_INT32
	case parser.TypesParserATOMIC_UINT32:
		return pipelang.Atomic_UINT32
	case parser.TypesParserATOMIC_INT64:
		return pipelang.Atomic_INT64
	case parser.TypesParserATOMIC_UINT64:
		return pipelang.Atomic_UINT64
	case parser.TypesParserATOMIC_BOOL:
		return pipelang.Atomic_BOOL
	case parser.TypesParserATOMIC_STRING:
		return pipelang.Atomic_STRING
	case parser.TypesParserATOMIC_CHAR:
		return pipelang.Atomic_CHAR
	case parser.TypesParserATOMIC_BYTES:
		return pipelang.Atomic_BYTES
	case parser.TypesParserATOMIC_DOUBLE:
		return pipelang.Atomic_DOUBLE
	case parser.TypesParserATOMIC_FLOAT:
		return pipelang.Atomic_FLOAT
	default:
		panic(fmt.Sprintf("unexpected atomic: %s", ctx.GetName().GetText()))
	}
}

// Public Interface

func ParseType(input string) *pipelang.PipelangType {
	is := antlr.NewInputStream(input)

	lexer := parser.NewTypesLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTypesParser(stream)

	visitor := NewTypesVisitor()

	tree := p.Start()
	result := visitor.Visit(tree)
	if result != nil {
		out := result.(pipelang.PipelangType)
		return &out
	} else {
		return nil
	}
}

func main() {
	result := ParseType("(Variant<int32, string, ts...>, char, {{ts}}...)")
	fmt.Println(result)
}
