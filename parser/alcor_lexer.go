// Code generated from parser/AlcorLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AlcorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AlcorLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func alcorlexerLexerInit() {
	staticData := &AlcorLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "STAR", "DIV", "DECIMAL_LIT", "BINARY_LIT", "OCTAL_LIT",
		"HEX_LIT", "SEP",
	}
	staticData.RuleNames = []string{
		"PLUS", "MINUS", "STAR", "DIV", "DECIMAL_LIT", "BINARY_LIT", "OCTAL_LIT",
		"HEX_LIT", "SEP", "DECIMALS", "OCTAL_DIGIT", "HEX_DIGIT", "BIN_DIGIT",
		"EXPONENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 104, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 41, 8, 4, 1, 4, 5,
		4, 44, 8, 4, 10, 4, 12, 4, 47, 9, 4, 3, 4, 49, 8, 4, 1, 5, 1, 5, 1, 5,
		3, 5, 54, 8, 5, 1, 5, 4, 5, 57, 8, 5, 11, 5, 12, 5, 58, 1, 6, 1, 6, 1,
		6, 3, 6, 64, 8, 6, 1, 6, 4, 6, 67, 8, 6, 11, 6, 12, 6, 68, 1, 7, 1, 7,
		1, 7, 3, 7, 74, 8, 7, 1, 7, 4, 7, 77, 8, 7, 11, 7, 12, 7, 78, 1, 8, 1,
		8, 1, 9, 1, 9, 3, 9, 85, 8, 9, 1, 9, 5, 9, 88, 8, 9, 10, 9, 12, 9, 91,
		9, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 3, 13, 101,
		8, 13, 1, 13, 1, 13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 1, 0, 10, 1, 0, 49,
		57, 1, 0, 48, 57, 2, 0, 66, 66, 98, 98, 1, 0, 48, 49, 2, 0, 79, 79, 111,
		111, 1, 0, 48, 55, 2, 0, 88, 88, 120, 120, 3, 0, 48, 57, 65, 70, 97, 102,
		2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 110, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		1, 29, 1, 0, 0, 0, 3, 31, 1, 0, 0, 0, 5, 33, 1, 0, 0, 0, 7, 35, 1, 0, 0,
		0, 9, 48, 1, 0, 0, 0, 11, 50, 1, 0, 0, 0, 13, 60, 1, 0, 0, 0, 15, 70, 1,
		0, 0, 0, 17, 80, 1, 0, 0, 0, 19, 82, 1, 0, 0, 0, 21, 92, 1, 0, 0, 0, 23,
		94, 1, 0, 0, 0, 25, 96, 1, 0, 0, 0, 27, 98, 1, 0, 0, 0, 29, 30, 5, 43,
		0, 0, 30, 2, 1, 0, 0, 0, 31, 32, 5, 45, 0, 0, 32, 4, 1, 0, 0, 0, 33, 34,
		5, 42, 0, 0, 34, 6, 1, 0, 0, 0, 35, 36, 5, 47, 0, 0, 36, 8, 1, 0, 0, 0,
		37, 49, 5, 48, 0, 0, 38, 45, 7, 0, 0, 0, 39, 41, 5, 95, 0, 0, 40, 39, 1,
		0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44, 7, 1, 0, 0, 43,
		40, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0,
		0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 37, 1, 0, 0, 0, 48, 38,
		1, 0, 0, 0, 49, 10, 1, 0, 0, 0, 50, 51, 5, 48, 0, 0, 51, 56, 7, 2, 0, 0,
		52, 54, 5, 95, 0, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 1,
		0, 0, 0, 55, 57, 7, 3, 0, 0, 56, 53, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58,
		56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 12, 1, 0, 0, 0, 60, 61, 5, 48,
		0, 0, 61, 66, 7, 4, 0, 0, 62, 64, 5, 95, 0, 0, 63, 62, 1, 0, 0, 0, 63,
		64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 7, 5, 0, 0, 66, 63, 1, 0, 0,
		0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 14,
		1, 0, 0, 0, 70, 71, 5, 48, 0, 0, 71, 76, 7, 6, 0, 0, 72, 74, 5, 95, 0,
		0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77,
		7, 7, 0, 0, 76, 73, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0,
		78, 79, 1, 0, 0, 0, 79, 16, 1, 0, 0, 0, 80, 81, 5, 32, 0, 0, 81, 18, 1,
		0, 0, 0, 82, 89, 7, 1, 0, 0, 83, 85, 5, 95, 0, 0, 84, 83, 1, 0, 0, 0, 84,
		85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 88, 7, 1, 0, 0, 87, 84, 1, 0, 0,
		0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 20,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 7, 5, 0, 0, 93, 22, 1, 0, 0, 0,
		94, 95, 7, 7, 0, 0, 95, 24, 1, 0, 0, 0, 96, 97, 7, 3, 0, 0, 97, 26, 1,
		0, 0, 0, 98, 100, 7, 8, 0, 0, 99, 101, 7, 9, 0, 0, 100, 99, 1, 0, 0, 0,
		100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 3, 19, 9, 0, 103,
		28, 1, 0, 0, 0, 13, 0, 40, 45, 48, 53, 58, 63, 68, 73, 78, 84, 89, 100,
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

// AlcorLexerInit initializes any static state used to implement AlcorLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAlcorLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlcorLexerInit() {
	staticData := &AlcorLexerLexerStaticData
	staticData.once.Do(alcorlexerLexerInit)
}

// NewAlcorLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAlcorLexer(input antlr.CharStream) *AlcorLexer {
	AlcorLexerInit()
	l := new(AlcorLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AlcorLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "AlcorLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AlcorLexer tokens.
const (
	AlcorLexerPLUS        = 1
	AlcorLexerMINUS       = 2
	AlcorLexerSTAR        = 3
	AlcorLexerDIV         = 4
	AlcorLexerDECIMAL_LIT = 5
	AlcorLexerBINARY_LIT  = 6
	AlcorLexerOCTAL_LIT   = 7
	AlcorLexerHEX_LIT     = 8
	AlcorLexerSEP         = 9
)
