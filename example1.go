// example1.go
package main

import (
	"fmt"

	"example/gvilums/types-parser/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bleenco/bms/protobuf/bms/protobuf/pipelang"
)

// TODO WHITESPACE HANDLING

// func parseType(input string) (pipelang.PipelangType, string, error) {
// 	if input[0] == '(' {
// 		tuple := parseTuple(input)
// 		return pipelang.PipelangType{Type: &pipelang.PipelangType_Tuple{Tuple: &tuple}}, input, nil
// 	} else {
// 		list := parseList(input)
// 		return pipelang.PipelangType{Type: &pipelang.PipelangType_List{List: &list}}, input, nil
// 	}
// }

// func parseTuple(input string) (pipelang.Tuple, string, error) {

// 	return pipelang.Tuple{Types: &pipelang.TypeList{}}, input, nil
// }

// func parseTypeList(input string) (pipelang.TypeList, string, error) {
// 	return pipelang.TypeList{}, input, nil
// }

// func parseList(input string) (pipelang.List, string, error) {
// 	if input[0] != '[' {
// 		return pipelang.List{}, input, errors.New("expected opening '[' for list type")
// 	}
// 	elem_type, remainder, err := parseType(input[1:])
// 	if err != nil {
// 		return pipelang.List{}, input, err
// 	}
// 	if remainder[0] != ']' {
// 		return pipelang.List{}, input, errors.New("expected closing ']' in list type")
// 	}
// 	return pipelang.List{ElemType: &elem_type}, remainder[1:], nil
// }

// func parseAtomic(input string) (pipelang.List, string, error) {

// }

type typesVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func NewTypesVisitor() typesVisitor {
	return typesVisitor{BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{}}
}

// TYPE

func (v *typesVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return ctx.GetValue().Accept(v)
}

func (v *typesVisitor) VisitList(ctx *parser.ListContext) pipelang.PipelangType {
	out := ctx.GetElemType().Accept(v).(pipelang.PipelangType)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_List{List: &pipelang.List{ElemType: &out}}}
}

func (v *typesVisitor) VisitTuple(ctx *parser.TupleContext) pipelang.PipelangType {
	out := ctx.GetTypes().Accept(v).(pipelang.TypeList)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Tuple{Tuple: &pipelang.Tuple{Types: &out}}}
}

func (v *typesVisitor) VisitAtomicType(ctx *parser.AtomicTypeContext) pipelang.PipelangType {
	out := ctx.GetAtom().Accept(v).(pipelang.Atomic)
	return pipelang.PipelangType{Type: &pipelang.PipelangType_Atomic{Atomic: out}}
}

// TYPE LIST

func (v *typesVisitor) VisitTypeListOnlyExpansion(ctx *parser.TypeListOnlyExpansionContext) pipelang.TypeList {
	pattern := ctx.GetPattern().Accept(v).(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: &pattern}
}

func (v *typesVisitor) VisitTypeListNoExpansion(ctx *parser.TypeListNoExpansionContext) pipelang.TypeList {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types))
	for i, t := range types {
		out := t.Accept(v).(pipelang.PipelangType)
		fixed[i] = &out
	}
	return pipelang.TypeList{Fixed: fixed, Expansion: nil}
}

func (v *typesVisitor) VisitTypeListExpansion(ctx *parser.TypeListExpansionContext) pipelang.TypeList {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types))
	for i, t := range types {
		out := t.Accept(v).(pipelang.PipelangType)
		fixed[i] = &out
	}
	pattern := ctx.GetPattern().Accept(v).(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: fixed, Expansion: &pattern}
}

func (v *typesVisitor) VisitTypeListEmpty(ctx *parser.TypeListEmptyContext) pipelang.TypeList {
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: nil}
}

func (v *typesVisitor) VisitAtomic(ctx *parser.AtomicContext) interface{} {
	switch ctx.GetName().GetTokenType() {
	case parser.TypesParserATOMIC_INT32:
		return pipelang.Atomic_INT32
	case parser.TypesParserATOMIC_UINT32:
		return pipelang.Atomic_UINT32
	default:
		panic(fmt.Sprintf("unexpected atomic: %s", ctx.GetName().GetText()))
	}
}

// func (v *typesVisitor) VisitAtomic(ctx *parser.AtomicContext) interface{} {
// 	return v.VisitChildren(ctx)
// }

func parse_type(input string) *pipelang.PipelangType {
	is := antlr.NewInputStream(input)

	lexer := parser.NewTypesLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTypesParser(stream)

	visitor := typesVisitor{BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{}}

	// visitor := parser.BaseTypesVisitor{BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{}}

	var opaque interface{} = visitor

	cast_visitor := opaque.(parser.TypesVisitor)

	tree := p.Start()
	result := cast_visitor.Visit(tree)
	if result != nil {
		out := result.(pipelang.PipelangType)
		return &out
	} else {
		return nil
	}
}

func main() {
	result := parse_type("int32")
	fmt.Println(result)
}
