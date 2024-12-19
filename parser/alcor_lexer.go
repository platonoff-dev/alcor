// Code generated from Alcor.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "NEWLINE", "INT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "NEWLINE", "INT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 41, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 4, 6, 31, 8, 6, 11,
		6, 12, 6, 32, 1, 6, 1, 6, 1, 7, 4, 7, 38, 8, 7, 11, 7, 12, 7, 39, 0, 0,
		8, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 1, 0, 2, 2, 0, 10,
		10, 13, 13, 1, 0, 48, 57, 42, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 19, 1, 0, 0, 0, 5,
		21, 1, 0, 0, 0, 7, 23, 1, 0, 0, 0, 9, 25, 1, 0, 0, 0, 11, 27, 1, 0, 0,
		0, 13, 30, 1, 0, 0, 0, 15, 37, 1, 0, 0, 0, 17, 18, 5, 42, 0, 0, 18, 2,
		1, 0, 0, 0, 19, 20, 5, 47, 0, 0, 20, 4, 1, 0, 0, 0, 21, 22, 5, 43, 0, 0,
		22, 6, 1, 0, 0, 0, 23, 24, 5, 45, 0, 0, 24, 8, 1, 0, 0, 0, 25, 26, 5, 40,
		0, 0, 26, 10, 1, 0, 0, 0, 27, 28, 5, 41, 0, 0, 28, 12, 1, 0, 0, 0, 29,
		31, 7, 0, 0, 0, 30, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 30, 1, 0, 0,
		0, 32, 33, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 6, 6, 0, 0, 35, 14,
		1, 0, 0, 0, 36, 38, 7, 1, 0, 0, 37, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0,
		39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 16, 1, 0, 0, 0, 3, 0, 32, 39,
		1, 6, 0, 0,
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
	l.GrammarFileName = "Alcor.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AlcorLexer tokens.
const (
	AlcorLexerT__0    = 1
	AlcorLexerT__1    = 2
	AlcorLexerT__2    = 3
	AlcorLexerT__3    = 4
	AlcorLexerT__4    = 5
	AlcorLexerT__5    = 6
	AlcorLexerNEWLINE = 7
	AlcorLexerINT     = 8
)
