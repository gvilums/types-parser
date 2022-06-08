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
		"", "'['", "']'", "'{{'", "'}}'", "'('", "')'", "'|'", "'{'", "'}'",
		"'<'", "'>'", "'...'", "','", "':'", "'int32'", "'uint32'", "'int64'",
		"'uint64'", "'string'", "'bool'", "'bytes'", "'double'", "'float'",
		"'char'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "ATOMIC_INT32",
		"ATOMIC_UINT32", "ATOMIC_INT64", "ATOMIC_UINT64", "ATOMIC_STRING", "ATOMIC_BOOL",
		"ATOMIC_BYTES", "ATOMIC_DOUBLE", "ATOMIC_FLOAT", "ATOMIC_CHAR", "IDENT",
		"FIXED_FIELDNAME", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "ATOMIC_INT32", "ATOMIC_UINT32",
		"ATOMIC_INT64", "ATOMIC_UINT64", "ATOMIC_STRING", "ATOMIC_BOOL", "ATOMIC_BYTES",
		"ATOMIC_DOUBLE", "ATOMIC_FLOAT", "ATOMIC_CHAR", "IDENT", "FIXED_FIELDNAME",
		"WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 167, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 5, 24, 152, 8, 24, 10, 24, 12, 24, 155, 9, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 26, 4, 26, 162, 8, 26, 11, 26, 12, 26, 163, 1, 26, 1, 26,
		0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 1,
		0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 3, 0, 9, 10, 13, 13, 32, 32, 168, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1, 55, 1, 0, 0, 0, 3, 57, 1, 0, 0,
		0, 5, 59, 1, 0, 0, 0, 7, 62, 1, 0, 0, 0, 9, 65, 1, 0, 0, 0, 11, 67, 1,
		0, 0, 0, 13, 69, 1, 0, 0, 0, 15, 71, 1, 0, 0, 0, 17, 73, 1, 0, 0, 0, 19,
		75, 1, 0, 0, 0, 21, 77, 1, 0, 0, 0, 23, 79, 1, 0, 0, 0, 25, 83, 1, 0, 0,
		0, 27, 85, 1, 0, 0, 0, 29, 87, 1, 0, 0, 0, 31, 93, 1, 0, 0, 0, 33, 100,
		1, 0, 0, 0, 35, 106, 1, 0, 0, 0, 37, 113, 1, 0, 0, 0, 39, 120, 1, 0, 0,
		0, 41, 125, 1, 0, 0, 0, 43, 131, 1, 0, 0, 0, 45, 138, 1, 0, 0, 0, 47, 144,
		1, 0, 0, 0, 49, 149, 1, 0, 0, 0, 51, 156, 1, 0, 0, 0, 53, 161, 1, 0, 0,
		0, 55, 56, 5, 91, 0, 0, 56, 2, 1, 0, 0, 0, 57, 58, 5, 93, 0, 0, 58, 4,
		1, 0, 0, 0, 59, 60, 5, 123, 0, 0, 60, 61, 5, 123, 0, 0, 61, 6, 1, 0, 0,
		0, 62, 63, 5, 125, 0, 0, 63, 64, 5, 125, 0, 0, 64, 8, 1, 0, 0, 0, 65, 66,
		5, 40, 0, 0, 66, 10, 1, 0, 0, 0, 67, 68, 5, 41, 0, 0, 68, 12, 1, 0, 0,
		0, 69, 70, 5, 124, 0, 0, 70, 14, 1, 0, 0, 0, 71, 72, 5, 123, 0, 0, 72,
		16, 1, 0, 0, 0, 73, 74, 5, 125, 0, 0, 74, 18, 1, 0, 0, 0, 75, 76, 5, 60,
		0, 0, 76, 20, 1, 0, 0, 0, 77, 78, 5, 62, 0, 0, 78, 22, 1, 0, 0, 0, 79,
		80, 5, 46, 0, 0, 80, 81, 5, 46, 0, 0, 81, 82, 5, 46, 0, 0, 82, 24, 1, 0,
		0, 0, 83, 84, 5, 44, 0, 0, 84, 26, 1, 0, 0, 0, 85, 86, 5, 58, 0, 0, 86,
		28, 1, 0, 0, 0, 87, 88, 5, 105, 0, 0, 88, 89, 5, 110, 0, 0, 89, 90, 5,
		116, 0, 0, 90, 91, 5, 51, 0, 0, 91, 92, 5, 50, 0, 0, 92, 30, 1, 0, 0, 0,
		93, 94, 5, 117, 0, 0, 94, 95, 5, 105, 0, 0, 95, 96, 5, 110, 0, 0, 96, 97,
		5, 116, 0, 0, 97, 98, 5, 51, 0, 0, 98, 99, 5, 50, 0, 0, 99, 32, 1, 0, 0,
		0, 100, 101, 5, 105, 0, 0, 101, 102, 5, 110, 0, 0, 102, 103, 5, 116, 0,
		0, 103, 104, 5, 54, 0, 0, 104, 105, 5, 52, 0, 0, 105, 34, 1, 0, 0, 0, 106,
		107, 5, 117, 0, 0, 107, 108, 5, 105, 0, 0, 108, 109, 5, 110, 0, 0, 109,
		110, 5, 116, 0, 0, 110, 111, 5, 54, 0, 0, 111, 112, 5, 52, 0, 0, 112, 36,
		1, 0, 0, 0, 113, 114, 5, 115, 0, 0, 114, 115, 5, 116, 0, 0, 115, 116, 5,
		114, 0, 0, 116, 117, 5, 105, 0, 0, 117, 118, 5, 110, 0, 0, 118, 119, 5,
		103, 0, 0, 119, 38, 1, 0, 0, 0, 120, 121, 5, 98, 0, 0, 121, 122, 5, 111,
		0, 0, 122, 123, 5, 111, 0, 0, 123, 124, 5, 108, 0, 0, 124, 40, 1, 0, 0,
		0, 125, 126, 5, 98, 0, 0, 126, 127, 5, 121, 0, 0, 127, 128, 5, 116, 0,
		0, 128, 129, 5, 101, 0, 0, 129, 130, 5, 115, 0, 0, 130, 42, 1, 0, 0, 0,
		131, 132, 5, 100, 0, 0, 132, 133, 5, 111, 0, 0, 133, 134, 5, 117, 0, 0,
		134, 135, 5, 98, 0, 0, 135, 136, 5, 108, 0, 0, 136, 137, 5, 101, 0, 0,
		137, 44, 1, 0, 0, 0, 138, 139, 5, 102, 0, 0, 139, 140, 5, 108, 0, 0, 140,
		141, 5, 111, 0, 0, 141, 142, 5, 97, 0, 0, 142, 143, 5, 116, 0, 0, 143,
		46, 1, 0, 0, 0, 144, 145, 5, 99, 0, 0, 145, 146, 5, 104, 0, 0, 146, 147,
		5, 97, 0, 0, 147, 148, 5, 114, 0, 0, 148, 48, 1, 0, 0, 0, 149, 153, 7,
		0, 0, 0, 150, 152, 7, 1, 0, 0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0,
		0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 50, 1, 0, 0, 0, 155,
		153, 1, 0, 0, 0, 156, 157, 5, 34, 0, 0, 157, 158, 3, 49, 24, 0, 158, 159,
		5, 34, 0, 0, 159, 52, 1, 0, 0, 0, 160, 162, 7, 2, 0, 0, 161, 160, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0,
		164, 165, 1, 0, 0, 0, 165, 166, 6, 26, 0, 0, 166, 54, 1, 0, 0, 0, 3, 0,
		153, 163, 1, 6, 0, 0,
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
	TypesLexerT__0            = 1
	TypesLexerT__1            = 2
	TypesLexerT__2            = 3
	TypesLexerT__3            = 4
	TypesLexerT__4            = 5
	TypesLexerT__5            = 6
	TypesLexerT__6            = 7
	TypesLexerT__7            = 8
	TypesLexerT__8            = 9
	TypesLexerT__9            = 10
	TypesLexerT__10           = 11
	TypesLexerT__11           = 12
	TypesLexerT__12           = 13
	TypesLexerT__13           = 14
	TypesLexerATOMIC_INT32    = 15
	TypesLexerATOMIC_UINT32   = 16
	TypesLexerATOMIC_INT64    = 17
	TypesLexerATOMIC_UINT64   = 18
	TypesLexerATOMIC_STRING   = 19
	TypesLexerATOMIC_BOOL     = 20
	TypesLexerATOMIC_BYTES    = 21
	TypesLexerATOMIC_DOUBLE   = 22
	TypesLexerATOMIC_FLOAT    = 23
	TypesLexerATOMIC_CHAR     = 24
	TypesLexerIDENT           = 25
	TypesLexerFIXED_FIELDNAME = 26
	TypesLexerWHITESPACE      = 27
)
