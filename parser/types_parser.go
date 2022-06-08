// Code generated from Types.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Types

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TypesParser struct {
	*antlr.BaseParser
}

var typesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func typesParserInit() {
	staticData := &typesParserStaticData
	staticData.literalNames = []string{
		"", "'...'", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'int32'",
		"'uint32'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACKET", "CLOSE_BRACKET",
		"OPEN_BRACE", "CLOSE_BRACE", "ATOMIC_INT32", "ATOMIC_UINT32", "IDENT",
		"WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "type", "typeList", "atomic",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 52, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 21,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 29, 8, 2, 11, 2, 12, 2,
		30, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 38, 8, 2, 11, 2, 12, 2, 39, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 46, 8, 2, 1, 3, 1, 3, 3, 3, 50, 8, 3, 1, 3,
		0, 0, 4, 0, 2, 4, 6, 0, 0, 55, 0, 8, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4,
		45, 1, 0, 0, 0, 6, 49, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 0, 0, 1,
		10, 1, 1, 0, 0, 0, 11, 12, 5, 5, 0, 0, 12, 13, 3, 2, 1, 0, 13, 14, 5, 6,
		0, 0, 14, 21, 1, 0, 0, 0, 15, 16, 5, 3, 0, 0, 16, 17, 3, 4, 2, 0, 17, 18,
		5, 4, 0, 0, 18, 21, 1, 0, 0, 0, 19, 21, 3, 6, 3, 0, 20, 11, 1, 0, 0, 0,
		20, 15, 1, 0, 0, 0, 20, 19, 1, 0, 0, 0, 21, 3, 1, 0, 0, 0, 22, 23, 3, 2,
		1, 0, 23, 24, 5, 1, 0, 0, 24, 46, 1, 0, 0, 0, 25, 26, 3, 2, 1, 0, 26, 27,
		5, 2, 0, 0, 27, 29, 1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0,
		30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 3,
		2, 1, 0, 33, 46, 1, 0, 0, 0, 34, 35, 3, 2, 1, 0, 35, 36, 5, 2, 0, 0, 36,
		38, 1, 0, 0, 0, 37, 34, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 37, 1, 0, 0,
		0, 39, 40, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 3, 2, 1, 0, 42, 43,
		5, 1, 0, 0, 43, 46, 1, 0, 0, 0, 44, 46, 1, 0, 0, 0, 45, 22, 1, 0, 0, 0,
		45, 28, 1, 0, 0, 0, 45, 37, 1, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 5, 1, 0,
		0, 0, 47, 50, 5, 9, 0, 0, 48, 50, 5, 10, 0, 0, 49, 47, 1, 0, 0, 0, 49,
		48, 1, 0, 0, 0, 50, 7, 1, 0, 0, 0, 5, 20, 30, 39, 45, 49,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TypesParserInit initializes any static state used to implement TypesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTypesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TypesParserInit() {
	staticData := &typesParserStaticData
	staticData.once.Do(typesParserInit)
}

// NewTypesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTypesParser(input antlr.TokenStream) *TypesParser {
	TypesParserInit()
	this := new(TypesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &typesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Types.g4"

	return this
}

// TypesParser tokens.
const (
	TypesParserEOF           = antlr.TokenEOF
	TypesParserT__0          = 1
	TypesParserT__1          = 2
	TypesParserOPEN_PAREN    = 3
	TypesParserCLOSE_PAREN   = 4
	TypesParserOPEN_BRACKET  = 5
	TypesParserCLOSE_BRACKET = 6
	TypesParserOPEN_BRACE    = 7
	TypesParserCLOSE_BRACE   = 8
	TypesParserATOMIC_INT32  = 9
	TypesParserATOMIC_UINT32 = 10
	TypesParserIDENT         = 11
	TypesParserWHITESPACE    = 12
)

// TypesParser rules.
const (
	TypesParserRULE_start    = 0
	TypesParserRULE_type     = 1
	TypesParserRULE_typeList = 2
	TypesParserRULE_atomic   = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() ITypeContext

	// SetValue sets the value rule contexts.
	SetValue(ITypeContext)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  ITypeContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) GetValue() ITypeContext { return s.value }

func (s *StartContext) SetValue(v ITypeContext) { s.value = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TypesParserEOF, 0)
}

func (s *StartContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TypesParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)

		var _x = p.Type()

		localctx.(*StartContext).value = _x
	}
	{
		p.SetState(9)
		p.Match(TypesParserEOF)
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyFrom(ctx *TypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListContext struct {
	*TypeContext
	elemType ITypeContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *ListContext) GetElemType() ITypeContext { return s.elemType }

func (s *ListContext) SetElemType(v ITypeContext) { s.elemType = v }

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(TypesParserOPEN_BRACKET, 0)
}

func (s *ListContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(TypesParserCLOSE_BRACKET, 0)
}

func (s *ListContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomicTypeContext struct {
	*TypeContext
	atom IAtomicContext
}

func NewAtomicTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomicTypeContext {
	var p = new(AtomicTypeContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *AtomicTypeContext) GetAtom() IAtomicContext { return s.atom }

func (s *AtomicTypeContext) SetAtom(v IAtomicContext) { s.atom = v }

func (s *AtomicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicTypeContext) Atomic() IAtomicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomicContext)
}

func (s *AtomicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitAtomicType(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleContext struct {
	*TypeContext
	types ITypeListContext
}

func NewTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleContext {
	var p = new(TupleContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *TupleContext) GetTypes() ITypeListContext { return s.types }

func (s *TupleContext) SetTypes(v ITypeListContext) { s.types = v }

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(TypesParserOPEN_PAREN, 0)
}

func (s *TupleContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(TypesParserCLOSE_PAREN, 0)
}

func (s *TupleContext) TypeList() ITypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *TupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) Type() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TypesParserRULE_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(20)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TypesParserOPEN_BRACKET:
		localctx = NewListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(11)
			p.Match(TypesParserOPEN_BRACKET)
		}
		{
			p.SetState(12)

			var _x = p.Type()

			localctx.(*ListContext).elemType = _x
		}
		{
			p.SetState(13)
			p.Match(TypesParserCLOSE_BRACKET)
		}

	case TypesParserOPEN_PAREN:
		localctx = NewTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Match(TypesParserOPEN_PAREN)
		}
		{
			p.SetState(16)

			var _x = p.TypeList()

			localctx.(*TupleContext).types = _x
		}
		{
			p.SetState(17)
			p.Match(TypesParserCLOSE_PAREN)
		}

	case TypesParserATOMIC_INT32, TypesParserATOMIC_UINT32:
		localctx = NewAtomicTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(19)

			var _x = p.Atomic()

			localctx.(*AtomicTypeContext).atom = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) CopyFrom(ctx *TypeListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeListExpansionContext struct {
	*TypeListContext
	_type   ITypeContext
	types   []ITypeContext
	pattern ITypeContext
}

func NewTypeListExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeListExpansionContext {
	var p = new(TypeListExpansionContext)

	p.TypeListContext = NewEmptyTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeListContext))

	return p
}

func (s *TypeListExpansionContext) Get_type() ITypeContext { return s._type }

func (s *TypeListExpansionContext) GetPattern() ITypeContext { return s.pattern }

func (s *TypeListExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *TypeListExpansionContext) SetPattern(v ITypeContext) { s.pattern = v }

func (s *TypeListExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *TypeListExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *TypeListExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListExpansionContext) AllType() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeListExpansionContext) Type(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeListExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTypeListExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeListEmptyContext struct {
	*TypeListContext
}

func NewTypeListEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeListEmptyContext {
	var p = new(TypeListEmptyContext)

	p.TypeListContext = NewEmptyTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeListContext))

	return p
}

func (s *TypeListEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTypeListEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeListOnlyExpansionContext struct {
	*TypeListContext
	pattern ITypeContext
}

func NewTypeListOnlyExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeListOnlyExpansionContext {
	var p = new(TypeListOnlyExpansionContext)

	p.TypeListContext = NewEmptyTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeListContext))

	return p
}

func (s *TypeListOnlyExpansionContext) GetPattern() ITypeContext { return s.pattern }

func (s *TypeListOnlyExpansionContext) SetPattern(v ITypeContext) { s.pattern = v }

func (s *TypeListOnlyExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListOnlyExpansionContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeListOnlyExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTypeListOnlyExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeListNoExpansionContext struct {
	*TypeListContext
	_type     ITypeContext
	types     []ITypeContext
	finalType ITypeContext
}

func NewTypeListNoExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeListNoExpansionContext {
	var p = new(TypeListNoExpansionContext)

	p.TypeListContext = NewEmptyTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeListContext))

	return p
}

func (s *TypeListNoExpansionContext) Get_type() ITypeContext { return s._type }

func (s *TypeListNoExpansionContext) GetFinalType() ITypeContext { return s.finalType }

func (s *TypeListNoExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *TypeListNoExpansionContext) SetFinalType(v ITypeContext) { s.finalType = v }

func (s *TypeListNoExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *TypeListNoExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *TypeListNoExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListNoExpansionContext) AllType() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *TypeListNoExpansionContext) Type(i int) ITypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *TypeListNoExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTypeListNoExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) TypeList() (localctx ITypeListContext) {
	this := p
	_ = this

	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TypesParserRULE_typeList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeListOnlyExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)

			var _x = p.Type()

			localctx.(*TypeListOnlyExpansionContext).pattern = _x
		}
		{
			p.SetState(23)
			p.Match(TypesParserT__0)
		}

	case 2:
		localctx = NewTypeListNoExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(25)

					var _x = p.Type()

					localctx.(*TypeListNoExpansionContext)._type = _x
				}
				localctx.(*TypeListNoExpansionContext).types = append(localctx.(*TypeListNoExpansionContext).types, localctx.(*TypeListNoExpansionContext)._type)
				{
					p.SetState(26)
					p.Match(TypesParserT__1)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}
		{
			p.SetState(32)

			var _x = p.Type()

			localctx.(*TypeListNoExpansionContext).finalType = _x
		}

	case 3:
		localctx = NewTypeListExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(34)

					var _x = p.Type()

					localctx.(*TypeListExpansionContext)._type = _x
				}
				localctx.(*TypeListExpansionContext).types = append(localctx.(*TypeListExpansionContext).types, localctx.(*TypeListExpansionContext)._type)
				{
					p.SetState(35)
					p.Match(TypesParserT__1)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}
		{
			p.SetState(41)

			var _x = p.Type()

			localctx.(*TypeListExpansionContext).pattern = _x
		}
		{
			p.SetState(42)
			p.Match(TypesParserT__0)
		}

	case 4:
		localctx = NewTypeListEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)

	}

	return localctx
}

// IAtomicContext is an interface to support dynamic dispatch.
type IAtomicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsAtomicContext differentiates from other interfaces.
	IsAtomicContext()
}

type AtomicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyAtomicContext() *AtomicContext {
	var p = new(AtomicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_atomic
	return p
}

func (*AtomicContext) IsAtomicContext() {}

func NewAtomicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicContext {
	var p = new(AtomicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_atomic

	return p
}

func (s *AtomicContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicContext) GetName() antlr.Token { return s.name }

func (s *AtomicContext) SetName(v antlr.Token) { s.name = v }

func (s *AtomicContext) ATOMIC_INT32() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_INT32, 0)
}

func (s *AtomicContext) ATOMIC_UINT32() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_UINT32, 0)
}

func (s *AtomicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitAtomic(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) Atomic() (localctx IAtomicContext) {
	this := p
	_ = this

	localctx = NewAtomicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TypesParserRULE_atomic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TypesParserATOMIC_INT32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)

			var _m = p.Match(TypesParserATOMIC_INT32)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_UINT32:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)

			var _m = p.Match(TypesParserATOMIC_UINT32)

			localctx.(*AtomicContext).name = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
