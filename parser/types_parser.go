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
		"", "'['", "']'", "'$'", "'('", "')'", "'|'", "'{'", "'}'", "'<'", "'>'",
		"'...'", "','", "':'", "'int32'", "'uint32'", "'int64'", "'uint64'",
		"'string'", "'bool'", "'bytes'", "'double'", "'float'", "'char'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ATOMIC_INT32",
		"ATOMIC_UINT32", "ATOMIC_INT64", "ATOMIC_UINT64", "ATOMIC_STRING", "ATOMIC_BOOL",
		"ATOMIC_BYTES", "ATOMIC_DOUBLE", "ATOMIC_FLOAT", "ATOMIC_CHAR", "IDENT",
		"FIXED_FIELDNAME", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"start", "type", "tupleTypeList", "unionTypeList", "fieldList", "fieldName",
		"atomic",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 148, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 43, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4, 2, 51, 8, 2, 11, 2, 12, 2, 52, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 4, 2, 60, 8, 2, 11, 2, 12, 2, 61, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 68, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 76, 8,
		3, 11, 3, 12, 3, 77, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 85, 8, 3, 11,
		3, 12, 3, 86, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 93, 8, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 106, 8, 4, 11, 4,
		12, 4, 107, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4,
		119, 8, 4, 11, 4, 12, 4, 120, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 130, 8, 4, 1, 5, 1, 5, 3, 5, 134, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 146, 8, 6, 1, 6, 0, 0, 7, 0, 2,
		4, 6, 8, 10, 12, 0, 0, 172, 0, 14, 1, 0, 0, 0, 2, 42, 1, 0, 0, 0, 4, 67,
		1, 0, 0, 0, 6, 92, 1, 0, 0, 0, 8, 129, 1, 0, 0, 0, 10, 133, 1, 0, 0, 0,
		12, 145, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 16, 5, 0, 0, 1, 16, 1, 1,
		0, 0, 0, 17, 18, 5, 1, 0, 0, 18, 19, 3, 2, 1, 0, 19, 20, 5, 2, 0, 0, 20,
		43, 1, 0, 0, 0, 21, 43, 5, 24, 0, 0, 22, 23, 5, 3, 0, 0, 23, 43, 5, 24,
		0, 0, 24, 25, 5, 4, 0, 0, 25, 26, 3, 4, 2, 0, 26, 27, 5, 5, 0, 0, 27, 43,
		1, 0, 0, 0, 28, 29, 5, 6, 0, 0, 29, 30, 3, 6, 3, 0, 30, 31, 5, 6, 0, 0,
		31, 43, 1, 0, 0, 0, 32, 33, 5, 7, 0, 0, 33, 34, 3, 8, 4, 0, 34, 35, 5,
		8, 0, 0, 35, 43, 1, 0, 0, 0, 36, 37, 5, 24, 0, 0, 37, 38, 5, 9, 0, 0, 38,
		39, 3, 4, 2, 0, 39, 40, 5, 10, 0, 0, 40, 43, 1, 0, 0, 0, 41, 43, 3, 12,
		6, 0, 42, 17, 1, 0, 0, 0, 42, 21, 1, 0, 0, 0, 42, 22, 1, 0, 0, 0, 42, 24,
		1, 0, 0, 0, 42, 28, 1, 0, 0, 0, 42, 32, 1, 0, 0, 0, 42, 36, 1, 0, 0, 0,
		42, 41, 1, 0, 0, 0, 43, 3, 1, 0, 0, 0, 44, 45, 3, 2, 1, 0, 45, 46, 5, 11,
		0, 0, 46, 68, 1, 0, 0, 0, 47, 48, 3, 2, 1, 0, 48, 49, 5, 12, 0, 0, 49,
		51, 1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 50, 1, 0, 0,
		0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 3, 2, 1, 0, 55, 68,
		1, 0, 0, 0, 56, 57, 3, 2, 1, 0, 57, 58, 5, 12, 0, 0, 58, 60, 1, 0, 0, 0,
		59, 56, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 3, 2, 1, 0, 64, 65, 5, 11, 0, 0, 65,
		68, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 44, 1, 0, 0, 0, 67, 50, 1, 0, 0,
		0, 67, 59, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 5, 1, 0, 0, 0, 69, 70, 3,
		2, 1, 0, 70, 71, 5, 11, 0, 0, 71, 93, 1, 0, 0, 0, 72, 73, 3, 2, 1, 0, 73,
		74, 5, 6, 0, 0, 74, 76, 1, 0, 0, 0, 75, 72, 1, 0, 0, 0, 76, 77, 1, 0, 0,
		0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80,
		3, 2, 1, 0, 80, 93, 1, 0, 0, 0, 81, 82, 3, 2, 1, 0, 82, 83, 5, 6, 0, 0,
		83, 85, 1, 0, 0, 0, 84, 81, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 84, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 3, 2, 1, 0, 89,
		90, 5, 11, 0, 0, 90, 93, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 69, 1, 0,
		0, 0, 92, 75, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 7,
		1, 0, 0, 0, 94, 95, 5, 3, 0, 0, 95, 96, 5, 24, 0, 0, 96, 97, 5, 13, 0,
		0, 97, 98, 3, 2, 1, 0, 98, 99, 5, 11, 0, 0, 99, 130, 1, 0, 0, 0, 100, 101,
		3, 10, 5, 0, 101, 102, 5, 13, 0, 0, 102, 103, 3, 2, 1, 0, 103, 104, 5,
		12, 0, 0, 104, 106, 1, 0, 0, 0, 105, 100, 1, 0, 0, 0, 106, 107, 1, 0, 0,
		0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		110, 3, 10, 5, 0, 110, 111, 5, 13, 0, 0, 111, 112, 3, 2, 1, 0, 112, 130,
		1, 0, 0, 0, 113, 114, 3, 10, 5, 0, 114, 115, 5, 13, 0, 0, 115, 116, 3,
		2, 1, 0, 116, 117, 5, 12, 0, 0, 117, 119, 1, 0, 0, 0, 118, 113, 1, 0, 0,
		0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		122, 1, 0, 0, 0, 122, 123, 5, 3, 0, 0, 123, 124, 5, 24, 0, 0, 124, 125,
		5, 13, 0, 0, 125, 126, 3, 2, 1, 0, 126, 127, 5, 11, 0, 0, 127, 130, 1,
		0, 0, 0, 128, 130, 1, 0, 0, 0, 129, 94, 1, 0, 0, 0, 129, 105, 1, 0, 0,
		0, 129, 118, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 9, 1, 0, 0, 0, 131,
		134, 5, 25, 0, 0, 132, 134, 5, 24, 0, 0, 133, 131, 1, 0, 0, 0, 133, 132,
		1, 0, 0, 0, 134, 11, 1, 0, 0, 0, 135, 146, 5, 14, 0, 0, 136, 146, 5, 15,
		0, 0, 137, 146, 5, 16, 0, 0, 138, 146, 5, 17, 0, 0, 139, 146, 5, 18, 0,
		0, 140, 146, 5, 19, 0, 0, 141, 146, 5, 20, 0, 0, 142, 146, 5, 21, 0, 0,
		143, 146, 5, 22, 0, 0, 144, 146, 5, 23, 0, 0, 145, 135, 1, 0, 0, 0, 145,
		136, 1, 0, 0, 0, 145, 137, 1, 0, 0, 0, 145, 138, 1, 0, 0, 0, 145, 139,
		1, 0, 0, 0, 145, 140, 1, 0, 0, 0, 145, 141, 1, 0, 0, 0, 145, 142, 1, 0,
		0, 0, 145, 143, 1, 0, 0, 0, 145, 144, 1, 0, 0, 0, 146, 13, 1, 0, 0, 0,
		12, 42, 52, 61, 67, 77, 86, 92, 107, 120, 129, 133, 145,
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
	TypesParserEOF             = antlr.TokenEOF
	TypesParserT__0            = 1
	TypesParserT__1            = 2
	TypesParserT__2            = 3
	TypesParserT__3            = 4
	TypesParserT__4            = 5
	TypesParserT__5            = 6
	TypesParserT__6            = 7
	TypesParserT__7            = 8
	TypesParserT__8            = 9
	TypesParserT__9            = 10
	TypesParserT__10           = 11
	TypesParserT__11           = 12
	TypesParserT__12           = 13
	TypesParserATOMIC_INT32    = 14
	TypesParserATOMIC_UINT32   = 15
	TypesParserATOMIC_INT64    = 16
	TypesParserATOMIC_UINT64   = 17
	TypesParserATOMIC_STRING   = 18
	TypesParserATOMIC_BOOL     = 19
	TypesParserATOMIC_BYTES    = 20
	TypesParserATOMIC_DOUBLE   = 21
	TypesParserATOMIC_FLOAT    = 22
	TypesParserATOMIC_CHAR     = 23
	TypesParserIDENT           = 24
	TypesParserFIXED_FIELDNAME = 25
	TypesParserWHITESPACE      = 26
)

// TypesParser rules.
const (
	TypesParserRULE_start         = 0
	TypesParserRULE_type          = 1
	TypesParserRULE_tupleTypeList = 2
	TypesParserRULE_unionTypeList = 3
	TypesParserRULE_fieldList     = 4
	TypesParserRULE_fieldName     = 5
	TypesParserRULE_atomic        = 6
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
		p.SetState(14)

		var _x = p.Type()

		localctx.(*StartContext).value = _x
	}
	{
		p.SetState(15)
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

type NamedContext struct {
	*TypeContext
	name  antlr.Token
	types ITupleTypeListContext
}

func NewNamedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedContext {
	var p = new(NamedContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *NamedContext) GetName() antlr.Token { return s.name }

func (s *NamedContext) SetName(v antlr.Token) { s.name = v }

func (s *NamedContext) GetTypes() ITupleTypeListContext { return s.types }

func (s *NamedContext) SetTypes(v ITupleTypeListContext) { s.types = v }

func (s *NamedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TypesParserIDENT, 0)
}

func (s *NamedContext) TupleTypeList() ITupleTypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleTypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleTypeListContext)
}

func (s *NamedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitNamed(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeVarContext struct {
	*TypeContext
	name antlr.Token
}

func NewTypeVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeVarContext {
	var p = new(TypeVarContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *TypeVarContext) GetName() antlr.Token { return s.name }

func (s *TypeVarContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeVarContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TypesParserIDENT, 0)
}

func (s *TypeVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTypeVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type PackVarContext struct {
	*TypeContext
	name antlr.Token
}

func NewPackVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PackVarContext {
	var p = new(PackVarContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *PackVarContext) GetName() antlr.Token { return s.name }

func (s *PackVarContext) SetName(v antlr.Token) { s.name = v }

func (s *PackVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackVarContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TypesParserIDENT, 0)
}

func (s *PackVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitPackVar(s)

	default:
		return t.VisitChildren(s)
	}
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
	types ITupleTypeListContext
}

func NewTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleContext {
	var p = new(TupleContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *TupleContext) GetTypes() ITupleTypeListContext { return s.types }

func (s *TupleContext) SetTypes(v ITupleTypeListContext) { s.types = v }

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) TupleTypeList() ITupleTypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleTypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleTypeListContext)
}

func (s *TupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnionContext struct {
	*TypeContext
	types IUnionTypeListContext
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *UnionContext) GetTypes() IUnionTypeListContext { return s.types }

func (s *UnionContext) SetTypes(v IUnionTypeListContext) { s.types = v }

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) UnionTypeList() IUnionTypeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionTypeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionTypeListContext)
}

func (s *UnionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitUnion(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructContext struct {
	*TypeContext
	fields IFieldListContext
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *StructContext) GetFields() IFieldListContext { return s.fields }

func (s *StructContext) SetFields(v IFieldListContext) { s.fields = v }

func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitStruct(s)

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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.Match(TypesParserT__0)
		}
		{
			p.SetState(18)

			var _x = p.Type()

			localctx.(*ListContext).elemType = _x
		}
		{
			p.SetState(19)
			p.Match(TypesParserT__1)
		}

	case 2:
		localctx = NewTypeVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)

			var _m = p.Match(TypesParserIDENT)

			localctx.(*TypeVarContext).name = _m
		}

	case 3:
		localctx = NewPackVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Match(TypesParserT__2)
		}
		{
			p.SetState(23)

			var _m = p.Match(TypesParserIDENT)

			localctx.(*PackVarContext).name = _m
		}

	case 4:
		localctx = NewTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(24)
			p.Match(TypesParserT__3)
		}
		{
			p.SetState(25)

			var _x = p.TupleTypeList()

			localctx.(*TupleContext).types = _x
		}
		{
			p.SetState(26)
			p.Match(TypesParserT__4)
		}

	case 5:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(28)
			p.Match(TypesParserT__5)
		}
		{
			p.SetState(29)

			var _x = p.UnionTypeList()

			localctx.(*UnionContext).types = _x
		}
		{
			p.SetState(30)
			p.Match(TypesParserT__5)
		}

	case 6:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(32)
			p.Match(TypesParserT__6)
		}
		{
			p.SetState(33)

			var _x = p.FieldList()

			localctx.(*StructContext).fields = _x
		}
		{
			p.SetState(34)
			p.Match(TypesParserT__7)
		}

	case 7:
		localctx = NewNamedContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(36)

			var _m = p.Match(TypesParserIDENT)

			localctx.(*NamedContext).name = _m
		}
		{
			p.SetState(37)
			p.Match(TypesParserT__8)
		}
		{
			p.SetState(38)

			var _x = p.TupleTypeList()

			localctx.(*NamedContext).types = _x
		}
		{
			p.SetState(39)
			p.Match(TypesParserT__9)
		}

	case 8:
		localctx = NewAtomicTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(41)

			var _x = p.Atomic()

			localctx.(*AtomicTypeContext).atom = _x
		}

	}

	return localctx
}

// ITupleTypeListContext is an interface to support dynamic dispatch.
type ITupleTypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleTypeListContext differentiates from other interfaces.
	IsTupleTypeListContext()
}

type TupleTypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleTypeListContext() *TupleTypeListContext {
	var p = new(TupleTypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_tupleTypeList
	return p
}

func (*TupleTypeListContext) IsTupleTypeListContext() {}

func NewTupleTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleTypeListContext {
	var p = new(TupleTypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_tupleTypeList

	return p
}

func (s *TupleTypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleTypeListContext) CopyFrom(ctx *TupleTypeListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TupleTypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TupleTypeListOnlyExpansionContext struct {
	*TupleTypeListContext
	pattern ITypeContext
}

func NewTupleTypeListOnlyExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleTypeListOnlyExpansionContext {
	var p = new(TupleTypeListOnlyExpansionContext)

	p.TupleTypeListContext = NewEmptyTupleTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TupleTypeListContext))

	return p
}

func (s *TupleTypeListOnlyExpansionContext) GetPattern() ITypeContext { return s.pattern }

func (s *TupleTypeListOnlyExpansionContext) SetPattern(v ITypeContext) { s.pattern = v }

func (s *TupleTypeListOnlyExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeListOnlyExpansionContext) Type() ITypeContext {
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

func (s *TupleTypeListOnlyExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTupleTypeListOnlyExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleTypeListExpansionContext struct {
	*TupleTypeListContext
	_type   ITypeContext
	types   []ITypeContext
	pattern ITypeContext
}

func NewTupleTypeListExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleTypeListExpansionContext {
	var p = new(TupleTypeListExpansionContext)

	p.TupleTypeListContext = NewEmptyTupleTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TupleTypeListContext))

	return p
}

func (s *TupleTypeListExpansionContext) Get_type() ITypeContext { return s._type }

func (s *TupleTypeListExpansionContext) GetPattern() ITypeContext { return s.pattern }

func (s *TupleTypeListExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *TupleTypeListExpansionContext) SetPattern(v ITypeContext) { s.pattern = v }

func (s *TupleTypeListExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *TupleTypeListExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *TupleTypeListExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeListExpansionContext) AllType() []ITypeContext {
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

func (s *TupleTypeListExpansionContext) Type(i int) ITypeContext {
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

func (s *TupleTypeListExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTupleTypeListExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleTypeListNoExpansionContext struct {
	*TupleTypeListContext
	_type     ITypeContext
	types     []ITypeContext
	finalType ITypeContext
}

func NewTupleTypeListNoExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleTypeListNoExpansionContext {
	var p = new(TupleTypeListNoExpansionContext)

	p.TupleTypeListContext = NewEmptyTupleTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TupleTypeListContext))

	return p
}

func (s *TupleTypeListNoExpansionContext) Get_type() ITypeContext { return s._type }

func (s *TupleTypeListNoExpansionContext) GetFinalType() ITypeContext { return s.finalType }

func (s *TupleTypeListNoExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *TupleTypeListNoExpansionContext) SetFinalType(v ITypeContext) { s.finalType = v }

func (s *TupleTypeListNoExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *TupleTypeListNoExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *TupleTypeListNoExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeListNoExpansionContext) AllType() []ITypeContext {
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

func (s *TupleTypeListNoExpansionContext) Type(i int) ITypeContext {
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

func (s *TupleTypeListNoExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTupleTypeListNoExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleTypeListEmptyContext struct {
	*TupleTypeListContext
}

func NewTupleTypeListEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleTypeListEmptyContext {
	var p = new(TupleTypeListEmptyContext)

	p.TupleTypeListContext = NewEmptyTupleTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TupleTypeListContext))

	return p
}

func (s *TupleTypeListEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeListEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitTupleTypeListEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) TupleTypeList() (localctx ITupleTypeListContext) {
	this := p
	_ = this

	localctx = NewTupleTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TypesParserRULE_tupleTypeList)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTupleTypeListOnlyExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)

			var _x = p.Type()

			localctx.(*TupleTypeListOnlyExpansionContext).pattern = _x
		}
		{
			p.SetState(45)
			p.Match(TypesParserT__10)
		}

	case 2:
		localctx = NewTupleTypeListNoExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(47)

					var _x = p.Type()

					localctx.(*TupleTypeListNoExpansionContext)._type = _x
				}
				localctx.(*TupleTypeListNoExpansionContext).types = append(localctx.(*TupleTypeListNoExpansionContext).types, localctx.(*TupleTypeListNoExpansionContext)._type)
				{
					p.SetState(48)
					p.Match(TypesParserT__11)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
		}
		{
			p.SetState(54)

			var _x = p.Type()

			localctx.(*TupleTypeListNoExpansionContext).finalType = _x
		}

	case 3:
		localctx = NewTupleTypeListExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(56)

					var _x = p.Type()

					localctx.(*TupleTypeListExpansionContext)._type = _x
				}
				localctx.(*TupleTypeListExpansionContext).types = append(localctx.(*TupleTypeListExpansionContext).types, localctx.(*TupleTypeListExpansionContext)._type)
				{
					p.SetState(57)
					p.Match(TypesParserT__11)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
		}
		{
			p.SetState(63)

			var _x = p.Type()

			localctx.(*TupleTypeListExpansionContext).pattern = _x
		}
		{
			p.SetState(64)
			p.Match(TypesParserT__10)
		}

	case 4:
		localctx = NewTupleTypeListEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)

	}

	return localctx
}

// IUnionTypeListContext is an interface to support dynamic dispatch.
type IUnionTypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeListContext differentiates from other interfaces.
	IsUnionTypeListContext()
}

type UnionTypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeListContext() *UnionTypeListContext {
	var p = new(UnionTypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_unionTypeList
	return p
}

func (*UnionTypeListContext) IsUnionTypeListContext() {}

func NewUnionTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeListContext {
	var p = new(UnionTypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_unionTypeList

	return p
}

func (s *UnionTypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeListContext) CopyFrom(ctx *UnionTypeListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *UnionTypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnionTypeListNoExpansionContext struct {
	*UnionTypeListContext
	_type     ITypeContext
	types     []ITypeContext
	finalType ITypeContext
}

func NewUnionTypeListNoExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionTypeListNoExpansionContext {
	var p = new(UnionTypeListNoExpansionContext)

	p.UnionTypeListContext = NewEmptyUnionTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnionTypeListContext))

	return p
}

func (s *UnionTypeListNoExpansionContext) Get_type() ITypeContext { return s._type }

func (s *UnionTypeListNoExpansionContext) GetFinalType() ITypeContext { return s.finalType }

func (s *UnionTypeListNoExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *UnionTypeListNoExpansionContext) SetFinalType(v ITypeContext) { s.finalType = v }

func (s *UnionTypeListNoExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *UnionTypeListNoExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *UnionTypeListNoExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeListNoExpansionContext) AllType() []ITypeContext {
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

func (s *UnionTypeListNoExpansionContext) Type(i int) ITypeContext {
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

func (s *UnionTypeListNoExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitUnionTypeListNoExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnionTypeListOnlyExpansionContext struct {
	*UnionTypeListContext
	pattern ITypeContext
}

func NewUnionTypeListOnlyExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionTypeListOnlyExpansionContext {
	var p = new(UnionTypeListOnlyExpansionContext)

	p.UnionTypeListContext = NewEmptyUnionTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnionTypeListContext))

	return p
}

func (s *UnionTypeListOnlyExpansionContext) GetPattern() ITypeContext { return s.pattern }

func (s *UnionTypeListOnlyExpansionContext) SetPattern(v ITypeContext) { s.pattern = v }

func (s *UnionTypeListOnlyExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeListOnlyExpansionContext) Type() ITypeContext {
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

func (s *UnionTypeListOnlyExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitUnionTypeListOnlyExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnionTypeListExpansionContext struct {
	*UnionTypeListContext
	_type   ITypeContext
	types   []ITypeContext
	pattern ITypeContext
}

func NewUnionTypeListExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionTypeListExpansionContext {
	var p = new(UnionTypeListExpansionContext)

	p.UnionTypeListContext = NewEmptyUnionTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnionTypeListContext))

	return p
}

func (s *UnionTypeListExpansionContext) Get_type() ITypeContext { return s._type }

func (s *UnionTypeListExpansionContext) GetPattern() ITypeContext { return s.pattern }

func (s *UnionTypeListExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *UnionTypeListExpansionContext) SetPattern(v ITypeContext) { s.pattern = v }

func (s *UnionTypeListExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *UnionTypeListExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *UnionTypeListExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeListExpansionContext) AllType() []ITypeContext {
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

func (s *UnionTypeListExpansionContext) Type(i int) ITypeContext {
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

func (s *UnionTypeListExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitUnionTypeListExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnionTypeListEmptyContext struct {
	*UnionTypeListContext
}

func NewUnionTypeListEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionTypeListEmptyContext {
	var p = new(UnionTypeListEmptyContext)

	p.UnionTypeListContext = NewEmptyUnionTypeListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnionTypeListContext))

	return p
}

func (s *UnionTypeListEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeListEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitUnionTypeListEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) UnionTypeList() (localctx IUnionTypeListContext) {
	this := p
	_ = this

	localctx = NewUnionTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TypesParserRULE_unionTypeList)

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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnionTypeListOnlyExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)

			var _x = p.Type()

			localctx.(*UnionTypeListOnlyExpansionContext).pattern = _x
		}
		{
			p.SetState(70)
			p.Match(TypesParserT__10)
		}

	case 2:
		localctx = NewUnionTypeListNoExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(72)

					var _x = p.Type()

					localctx.(*UnionTypeListNoExpansionContext)._type = _x
				}
				localctx.(*UnionTypeListNoExpansionContext).types = append(localctx.(*UnionTypeListNoExpansionContext).types, localctx.(*UnionTypeListNoExpansionContext)._type)
				{
					p.SetState(73)
					p.Match(TypesParserT__5)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}
		{
			p.SetState(79)

			var _x = p.Type()

			localctx.(*UnionTypeListNoExpansionContext).finalType = _x
		}

	case 3:
		localctx = NewUnionTypeListExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(81)

					var _x = p.Type()

					localctx.(*UnionTypeListExpansionContext)._type = _x
				}
				localctx.(*UnionTypeListExpansionContext).types = append(localctx.(*UnionTypeListExpansionContext).types, localctx.(*UnionTypeListExpansionContext)._type)
				{
					p.SetState(82)
					p.Match(TypesParserT__5)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}
		{
			p.SetState(88)

			var _x = p.Type()

			localctx.(*UnionTypeListExpansionContext).pattern = _x
		}
		{
			p.SetState(89)
			p.Match(TypesParserT__10)
		}

	case 4:
		localctx = NewUnionTypeListEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)

	}

	return localctx
}

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_fieldList
	return p
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) CopyFrom(ctx *FieldListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldListNoExpansionContext struct {
	*FieldListContext
	_fieldName IFieldNameContext
	names      []IFieldNameContext
	_type      ITypeContext
	types      []ITypeContext
	finalName  IFieldNameContext
	finalType  ITypeContext
}

func NewFieldListNoExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldListNoExpansionContext {
	var p = new(FieldListNoExpansionContext)

	p.FieldListContext = NewEmptyFieldListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldListContext))

	return p
}

func (s *FieldListNoExpansionContext) Get_fieldName() IFieldNameContext { return s._fieldName }

func (s *FieldListNoExpansionContext) Get_type() ITypeContext { return s._type }

func (s *FieldListNoExpansionContext) GetFinalName() IFieldNameContext { return s.finalName }

func (s *FieldListNoExpansionContext) GetFinalType() ITypeContext { return s.finalType }

func (s *FieldListNoExpansionContext) Set_fieldName(v IFieldNameContext) { s._fieldName = v }

func (s *FieldListNoExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *FieldListNoExpansionContext) SetFinalName(v IFieldNameContext) { s.finalName = v }

func (s *FieldListNoExpansionContext) SetFinalType(v ITypeContext) { s.finalType = v }

func (s *FieldListNoExpansionContext) GetNames() []IFieldNameContext { return s.names }

func (s *FieldListNoExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *FieldListNoExpansionContext) SetNames(v []IFieldNameContext) { s.names = v }

func (s *FieldListNoExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *FieldListNoExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListNoExpansionContext) AllFieldName() []IFieldNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldNameContext); ok {
			len++
		}
	}

	tst := make([]IFieldNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldNameContext); ok {
			tst[i] = t.(IFieldNameContext)
			i++
		}
	}

	return tst
}

func (s *FieldListNoExpansionContext) FieldName(i int) IFieldNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
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

	return t.(IFieldNameContext)
}

func (s *FieldListNoExpansionContext) AllType() []ITypeContext {
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

func (s *FieldListNoExpansionContext) Type(i int) ITypeContext {
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

func (s *FieldListNoExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitFieldListNoExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldListOnlyExpansionContext struct {
	*FieldListContext
	namePattern antlr.Token
	typePattern ITypeContext
}

func NewFieldListOnlyExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldListOnlyExpansionContext {
	var p = new(FieldListOnlyExpansionContext)

	p.FieldListContext = NewEmptyFieldListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldListContext))

	return p
}

func (s *FieldListOnlyExpansionContext) GetNamePattern() antlr.Token { return s.namePattern }

func (s *FieldListOnlyExpansionContext) SetNamePattern(v antlr.Token) { s.namePattern = v }

func (s *FieldListOnlyExpansionContext) GetTypePattern() ITypeContext { return s.typePattern }

func (s *FieldListOnlyExpansionContext) SetTypePattern(v ITypeContext) { s.typePattern = v }

func (s *FieldListOnlyExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListOnlyExpansionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TypesParserIDENT, 0)
}

func (s *FieldListOnlyExpansionContext) Type() ITypeContext {
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

func (s *FieldListOnlyExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitFieldListOnlyExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldListEmptyContext struct {
	*FieldListContext
}

func NewFieldListEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldListEmptyContext {
	var p = new(FieldListEmptyContext)

	p.FieldListContext = NewEmptyFieldListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldListContext))

	return p
}

func (s *FieldListEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitFieldListEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldListExpansionContext struct {
	*FieldListContext
	_fieldName  IFieldNameContext
	names       []IFieldNameContext
	_type       ITypeContext
	types       []ITypeContext
	namePattern antlr.Token
	typePattern ITypeContext
}

func NewFieldListExpansionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldListExpansionContext {
	var p = new(FieldListExpansionContext)

	p.FieldListContext = NewEmptyFieldListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldListContext))

	return p
}

func (s *FieldListExpansionContext) GetNamePattern() antlr.Token { return s.namePattern }

func (s *FieldListExpansionContext) SetNamePattern(v antlr.Token) { s.namePattern = v }

func (s *FieldListExpansionContext) Get_fieldName() IFieldNameContext { return s._fieldName }

func (s *FieldListExpansionContext) Get_type() ITypeContext { return s._type }

func (s *FieldListExpansionContext) GetTypePattern() ITypeContext { return s.typePattern }

func (s *FieldListExpansionContext) Set_fieldName(v IFieldNameContext) { s._fieldName = v }

func (s *FieldListExpansionContext) Set_type(v ITypeContext) { s._type = v }

func (s *FieldListExpansionContext) SetTypePattern(v ITypeContext) { s.typePattern = v }

func (s *FieldListExpansionContext) GetNames() []IFieldNameContext { return s.names }

func (s *FieldListExpansionContext) GetTypes() []ITypeContext { return s.types }

func (s *FieldListExpansionContext) SetNames(v []IFieldNameContext) { s.names = v }

func (s *FieldListExpansionContext) SetTypes(v []ITypeContext) { s.types = v }

func (s *FieldListExpansionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListExpansionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TypesParserIDENT, 0)
}

func (s *FieldListExpansionContext) AllType() []ITypeContext {
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

func (s *FieldListExpansionContext) Type(i int) ITypeContext {
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

func (s *FieldListExpansionContext) AllFieldName() []IFieldNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldNameContext); ok {
			len++
		}
	}

	tst := make([]IFieldNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldNameContext); ok {
			tst[i] = t.(IFieldNameContext)
			i++
		}
	}

	return tst
}

func (s *FieldListExpansionContext) FieldName(i int) IFieldNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
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

	return t.(IFieldNameContext)
}

func (s *FieldListExpansionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitFieldListExpansion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) FieldList() (localctx IFieldListContext) {
	this := p
	_ = this

	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TypesParserRULE_fieldList)
	var _la int

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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFieldListOnlyExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(TypesParserT__2)
		}
		{
			p.SetState(95)

			var _m = p.Match(TypesParserIDENT)

			localctx.(*FieldListOnlyExpansionContext).namePattern = _m
		}
		{
			p.SetState(96)
			p.Match(TypesParserT__12)
		}
		{
			p.SetState(97)

			var _x = p.Type()

			localctx.(*FieldListOnlyExpansionContext).typePattern = _x
		}
		{
			p.SetState(98)
			p.Match(TypesParserT__10)
		}

	case 2:
		localctx = NewFieldListNoExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(100)

					var _x = p.FieldName()

					localctx.(*FieldListNoExpansionContext)._fieldName = _x
				}
				localctx.(*FieldListNoExpansionContext).names = append(localctx.(*FieldListNoExpansionContext).names, localctx.(*FieldListNoExpansionContext)._fieldName)
				{
					p.SetState(101)
					p.Match(TypesParserT__12)
				}
				{
					p.SetState(102)

					var _x = p.Type()

					localctx.(*FieldListNoExpansionContext)._type = _x
				}
				localctx.(*FieldListNoExpansionContext).types = append(localctx.(*FieldListNoExpansionContext).types, localctx.(*FieldListNoExpansionContext)._type)
				{
					p.SetState(103)
					p.Match(TypesParserT__11)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
		}
		{
			p.SetState(109)

			var _x = p.FieldName()

			localctx.(*FieldListNoExpansionContext).finalName = _x
		}
		{
			p.SetState(110)
			p.Match(TypesParserT__12)
		}
		{
			p.SetState(111)

			var _x = p.Type()

			localctx.(*FieldListNoExpansionContext).finalType = _x
		}

	case 3:
		localctx = NewFieldListExpansionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == TypesParserIDENT || _la == TypesParserFIXED_FIELDNAME {
			{
				p.SetState(113)

				var _x = p.FieldName()

				localctx.(*FieldListExpansionContext)._fieldName = _x
			}
			localctx.(*FieldListExpansionContext).names = append(localctx.(*FieldListExpansionContext).names, localctx.(*FieldListExpansionContext)._fieldName)
			{
				p.SetState(114)
				p.Match(TypesParserT__12)
			}
			{
				p.SetState(115)

				var _x = p.Type()

				localctx.(*FieldListExpansionContext)._type = _x
			}
			localctx.(*FieldListExpansionContext).types = append(localctx.(*FieldListExpansionContext).types, localctx.(*FieldListExpansionContext)._type)
			{
				p.SetState(116)
				p.Match(TypesParserT__11)
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(122)
			p.Match(TypesParserT__2)
		}
		{
			p.SetState(123)

			var _m = p.Match(TypesParserIDENT)

			localctx.(*FieldListExpansionContext).namePattern = _m
		}
		{
			p.SetState(124)
			p.Match(TypesParserT__12)
		}
		{
			p.SetState(125)

			var _x = p.Type()

			localctx.(*FieldListExpansionContext).typePattern = _x
		}
		{
			p.SetState(126)
			p.Match(TypesParserT__10)
		}

	case 4:
		localctx = NewFieldListEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)

	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TypesParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TypesParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) CopyFrom(ctx *FieldNameContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldNameFixedContext struct {
	*FieldNameContext
	name antlr.Token
}

func NewFieldNameFixedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldNameFixedContext {
	var p = new(FieldNameFixedContext)

	p.FieldNameContext = NewEmptyFieldNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldNameContext))

	return p
}

func (s *FieldNameFixedContext) GetName() antlr.Token { return s.name }

func (s *FieldNameFixedContext) SetName(v antlr.Token) { s.name = v }

func (s *FieldNameFixedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameFixedContext) FIXED_FIELDNAME() antlr.TerminalNode {
	return s.GetToken(TypesParserFIXED_FIELDNAME, 0)
}

func (s *FieldNameFixedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitFieldNameFixed(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldNameVariableContext struct {
	*FieldNameContext
	name antlr.Token
}

func NewFieldNameVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldNameVariableContext {
	var p = new(FieldNameVariableContext)

	p.FieldNameContext = NewEmptyFieldNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldNameContext))

	return p
}

func (s *FieldNameVariableContext) GetName() antlr.Token { return s.name }

func (s *FieldNameVariableContext) SetName(v antlr.Token) { s.name = v }

func (s *FieldNameVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameVariableContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TypesParserIDENT, 0)
}

func (s *FieldNameVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TypesVisitor:
		return t.VisitFieldNameVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TypesParser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TypesParserRULE_fieldName)

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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TypesParserFIXED_FIELDNAME:
		localctx = NewFieldNameFixedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)

			var _m = p.Match(TypesParserFIXED_FIELDNAME)

			localctx.(*FieldNameFixedContext).name = _m
		}

	case TypesParserIDENT:
		localctx = NewFieldNameVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)

			var _m = p.Match(TypesParserIDENT)

			localctx.(*FieldNameVariableContext).name = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *AtomicContext) ATOMIC_INT64() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_INT64, 0)
}

func (s *AtomicContext) ATOMIC_UINT64() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_UINT64, 0)
}

func (s *AtomicContext) ATOMIC_STRING() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_STRING, 0)
}

func (s *AtomicContext) ATOMIC_BOOL() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_BOOL, 0)
}

func (s *AtomicContext) ATOMIC_BYTES() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_BYTES, 0)
}

func (s *AtomicContext) ATOMIC_DOUBLE() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_DOUBLE, 0)
}

func (s *AtomicContext) ATOMIC_FLOAT() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_FLOAT, 0)
}

func (s *AtomicContext) ATOMIC_CHAR() antlr.TerminalNode {
	return s.GetToken(TypesParserATOMIC_CHAR, 0)
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
	p.EnterRule(localctx, 12, TypesParserRULE_atomic)

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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TypesParserATOMIC_INT32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)

			var _m = p.Match(TypesParserATOMIC_INT32)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_UINT32:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)

			var _m = p.Match(TypesParserATOMIC_UINT32)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_INT64:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)

			var _m = p.Match(TypesParserATOMIC_INT64)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_UINT64:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)

			var _m = p.Match(TypesParserATOMIC_UINT64)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)

			var _m = p.Match(TypesParserATOMIC_STRING)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_BOOL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)

			var _m = p.Match(TypesParserATOMIC_BOOL)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_BYTES:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(141)

			var _m = p.Match(TypesParserATOMIC_BYTES)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_DOUBLE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(142)

			var _m = p.Match(TypesParserATOMIC_DOUBLE)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_FLOAT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(143)

			var _m = p.Match(TypesParserATOMIC_FLOAT)

			localctx.(*AtomicContext).name = _m
		}

	case TypesParserATOMIC_CHAR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(144)

			var _m = p.Match(TypesParserATOMIC_CHAR)

			localctx.(*AtomicContext).name = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
