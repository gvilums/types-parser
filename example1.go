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

func parseTuple(input string) (pipelang.Tuple, string, error) {

	return pipelang.Tuple{Types: &pipelang.TypeList{}}, input, nil
}

func parseTypeList(input string) (pipelang.TypeList, string, error) {
	return pipelang.TypeList{}, input, nil
}

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
	*parser.BaseTypesVisitor
}

// TYPE

func (v *typesVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.Visit(ctx.GetValue())
}

func (v *typesVisitor) VisitList(ctx *parser.ListContext) pipelang.List {
	out := v.Visit(ctx.GetElemType()).(pipelang.PipelangType)
	return pipelang.List{ElemType: &out}
}

func (v *typesVisitor) VisitTuple(ctx *parser.TupleContext) pipelang.Tuple {
	out := v.Visit(ctx.GetTypes()).(pipelang.TypeList)
	return pipelang.Tuple{Types: &out}
}

func (v *typesVisitor) VisitAtomicType(ctx *parser.TupleContext) pipelang.PipelangType {
	out := v.Visit(ctx.GetTypes()).(pipelang.TypeList)
	x := pipelang.PipelangType{Type: }
	return pipelang.PipelangType{Type: pipelang.TypeVariable{name: "test"}}
	return pipelang.Tuple{Types: &out}
}

// TYPE LIST

func (v *typesVisitor) VisitTypeListOnlyExpansion(ctx *parser.TypeListOnlyExpansionContext) pipelang.TypeList {
	pattern := v.Visit(ctx.GetPattern()).(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: &pattern}
}

func (v *typesVisitor) VisitTypeListNoExpansion(ctx *parser.TypeListNoExpansionContext) pipelang.TypeList {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types))
	for i, t := range types {
		out := v.Visit(t).(pipelang.PipelangType)
		fixed[i] = &out
	}
	return pipelang.TypeList{Fixed: fixed, Expansion: nil}
}

func (v *typesVisitor) VisitTypeListExpansion(ctx *parser.TypeListExpansionContext) pipelang.TypeList {
	types := ctx.GetTypes()
	fixed := make([]*pipelang.PipelangType, len(types))
	for i, t := range types {
		out := v.Visit(t).(pipelang.PipelangType)
		fixed[i] = &out
	}
	pattern := v.Visit(ctx.GetPattern()).(pipelang.PipelangType)
	return pipelang.TypeList{Fixed: fixed, Expansion: &pattern}
}

func (v *typesVisitor) VisitTypeListEmpty(ctx *parser.TypeListEmptyContext) pipelang.TypeList {
	return pipelang.TypeList{Fixed: make([]*pipelang.PipelangType, 0), Expansion: nil}
}

func (v *typesVisitor) VisitAtomic(ctx *parser.AtomicContext) pipelang.Atomic {
	switch ctx.GetName().GetTokenType() {
	case parser.TypesParserATOMIC_INT32:
		return pipelang.Atomic_INT32
	case parser.TypesParserATOMIC_UINT32:
		return pipelang.Atomic_UINT32
	default:
		panic(fmt.Sprintf("unexpected atomic: %s", ctx.GetName().GetText()))
	}
}

func parse_type(input string) *pipelang.PipelangType {
	is := antlr.NewInputStream(input)

	lexer := parser.NewTypesLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTypesParser(stream)

	var visitor typesVisitor

	result := visitor.Visit(p.Start())
	if result != nil {
		out := result.(pipelang.PipelangType)
		return &out
	} else {
		return nil
	}
}

func main() {
	result := parse_type("(int32, [uint32],)")
	fmt.Println(result)
}
