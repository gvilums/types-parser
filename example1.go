// example1.go
package main

import (
	"errors"
	"fmt"
	"strconv"

	"example/gvilums/types-parser/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bleenco/bms/protobuf/bms/protobuf/pipelang"
)

type errorListener struct {
	*antlr.DefaultErrorListener
	errors string
}

func newErrorListener() errorListener {
	return errorListener{DefaultErrorListener: &antlr.DefaultErrorListener{}, errors: ""}
}

func (listener *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	listener.errors += "line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg + "\n"
}

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

func (v *typesVisitor) VisitTypeNonUnion(ctx *parser.TypeNonUnionContext) interface{} {
	return ctx.GetT().Accept(v)
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

func (v *typesVisitor) VisitStruct(ctx *parser.StructContext) interface{} {
	out := ctx.GetFields().Accept(v)
	if out == nil {
		return nil
	}
	fl := out.(pipelang.FieldList)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Struct{Struct: &pipelang.Struct{Fields: &fl}}}
}

func (v *typesVisitor) VisitNamedSingle(ctx *parser.NamedSingleContext) interface{} {
	name := ctx.GetName().GetText()
	out := ctx.GetElemType().Accept(v)
	if out == nil {
		return nil
	}
	typeArray := make([]*pipelang.PipelangType, 1)
	outPl := out.(pipelang.PipelangType)
	typeArray[0] = &outPl

	return pipelang.PipelangType{Type: &pipelang.PipelangType_Named{Named: &pipelang.NamedType{Name: name, Arguments: &pipelang.TypeList{Fixed: typeArray, Expansion: nil}}}}
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

func (v *typesVisitor) VisitTypeUnion(ctx *parser.TypeUnionContext) interface{} {
	out := ctx.GetTypes().Accept(v)
	if out == nil {
		return nil
	}
	tl := out.(pipelang.TypeList)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Union{Union: &pipelang.Union{Types: &tl}}}
}

func (v *typesVisitor) VisitParenthesized(ctx *parser.ParenthesizedContext) interface{} {
	return ctx.GetInnerType().Accept(v)
}

// TUPLE TYPE LIST

func (v *typesVisitor) VisitTupleTypeListNoExpansion(ctx *parser.TupleTypeListNoExpansionContext) interface{} {
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

// UNION TYPE LIST

func (v *typesVisitor) VisitUnionTypeListNoExpansion(ctx *parser.UnionTypeListNoExpansionContext) interface{} {
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
	pat := t.(pipelang.PipelangType)
	return &pipelang.FieldList_FieldExpansion{NamePack: n, Pattern: &pat}
}

func (v *typesVisitor) VisitFieldListNoExpansion(ctx *parser.FieldListNoExpansionContext) interface{} {
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
	name := ctx.GetName().GetText()
	withoutQuotes := name[1 : len(name)-1]
	return pipelang.FieldList_Field_Fixed{Fixed: withoutQuotes}
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

func ParseType(input string) (pipelang.PipelangType, error) {
	is := antlr.NewInputStream(input)

	lexer := parser.NewTypesLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTypesParser(stream)

	listener := newErrorListener()

	p.RemoveErrorListeners()
	p.AddErrorListener(&listener)

	visitor := NewTypesVisitor()

	tree := p.Start()
	result := visitor.Visit(tree)
	if result != nil && listener.errors == "" {
		return result.(pipelang.PipelangType), nil
	} else {
		return pipelang.PipelangType{}, errors.New(listener.errors)
	}
}

func main() {
	result, err := ParseType("int32 | (a | Maybe<int32>)")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(&result)
	}
}
