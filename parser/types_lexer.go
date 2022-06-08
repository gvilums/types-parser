// Code generated from Types.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TypesLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var typeslexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func typeslexerLexerInit() {
	staticData := &typeslexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_BRACKET", "CLOSE_BRACKET",
		"OPEN_BRACE", "CLOSE_BRACE", "ATOMIC_INT32", "ATOMIC_UINT32", "IDENT",
		"WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 69, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 4, 10, 59, 8, 10, 11, 10, 12, 10, 60, 1, 11, 4, 11, 64, 8, 11, 11,
		11, 12, 11, 65, 1, 11, 1, 11, 0, 0, 12, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 1, 0, 3, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13,
		13, 32, 32, 70, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0,
		0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0,
		0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 29, 1, 0, 0, 0, 5, 31, 1, 0,
		0, 0, 7, 33, 1, 0, 0, 0, 9, 35, 1, 0, 0, 0, 11, 37, 1, 0, 0, 0, 13, 39,
		1, 0, 0, 0, 15, 41, 1, 0, 0, 0, 17, 43, 1, 0, 0, 0, 19, 49, 1, 0, 0, 0,
		21, 56, 1, 0, 0, 0, 23, 63, 1, 0, 0, 0, 25, 26, 5, 46, 0, 0, 26, 27, 5,
		46, 0, 0, 27, 28, 5, 46, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30, 5, 44, 0, 0,
		30, 4, 1, 0, 0, 0, 31, 32, 5, 40, 0, 0, 32, 6, 1, 0, 0, 0, 33, 34, 5, 41,
		0, 0, 34, 8, 1, 0, 0, 0, 35, 36, 5, 91, 0, 0, 36, 10, 1, 0, 0, 0, 37, 38,
		5, 93, 0, 0, 38, 12, 1, 0, 0, 0, 39, 40, 5, 123, 0, 0, 40, 14, 1, 0, 0,
		0, 41, 42, 5, 125, 0, 0, 42, 16, 1, 0, 0, 0, 43, 44, 5, 105, 0, 0, 44,
		45, 5, 110, 0, 0, 45, 46, 5, 116, 0, 0, 46, 47, 5, 51, 0, 0, 47, 48, 5,
		50, 0, 0, 48, 18, 1, 0, 0, 0, 49, 50, 5, 117, 0, 0, 50, 51, 5, 105, 0,
		0, 51, 52, 5, 110, 0, 0, 52, 53, 5, 116, 0, 0, 53, 54, 5, 51, 0, 0, 54,
		55, 5, 50, 0, 0, 55, 20, 1, 0, 0, 0, 56, 58, 7, 0, 0, 0, 57, 59, 7, 1,
		0, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61,
		1, 0, 0, 0, 61, 22, 1, 0, 0, 0, 62, 64, 7, 2, 0, 0, 63, 62, 1, 0, 0, 0,
		64, 65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1,
		0, 0, 0, 67, 68, 6, 11, 0, 0, 68, 24, 1, 0, 0, 0, 3, 0, 60, 65, 1, 6, 0,
		0,
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

// TypesLexerInit initializes any static state used to implement TypesLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTypesLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TypesLexerInit() {
	staticData := &typeslexerLexerStaticData
	staticData.once.Do(typeslexerLexerInit)
}

// NewTypesLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTypesLexer(input antlr.CharStream) *TypesLexer {
	TypesLexerInit()
	l := new(TypesLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &typeslexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Types.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TypesLexer tokens.
const (
	TypesLexerT__0          = 1
	TypesLexerT__1          = 2
	TypesLexerOPEN_PAREN    = 3
	TypesLexerCLOSE_PAREN   = 4
	TypesLexerOPEN_BRACKET  = 5
	TypesLexerCLOSE_BRACKET = 6
	TypesLexerOPEN_BRACE    = 7
	TypesLexerCLOSE_BRACE   = 8
	TypesLexerATOMIC_INT32  = 9
	TypesLexerATOMIC_UINT32 = 10
	TypesLexerIDENT         = 11
	TypesLexerWHITESPACE    = 12
)
