// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package routingA

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type routingALexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var routingalexerLexerStaticData struct {
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

func routingalexerLexerInit() {
	staticData := &routingalexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'{'", "'}'", "'@'", "'=='", "'!='", "':'", "'='", "'!'", "'('",
		"')'", "','", "'->'", "'&&'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "WHITESPACE",
		"COMMENT_BLOCK", "COMMENT_LINE_SHARP", "ID", "NON_ID", "QUOTE_STRING",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "SAFE_ID_HEAD_CHAR", "SAFE_NONID_HEAD_CHAR",
		"SAFE_INTERMEDIATE_CHAR", "SAFE_CHAR", "DOUBLE_QUOTE_STRING", "SINGLE_QUOTE_STRING",
		"WHITESPACE", "COMMENT_BLOCK", "COMMENT_LINE_SHARP", "ID", "NON_ID",
		"QUOTE_STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 170, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 3, 16, 91, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17,
		97, 8, 17, 10, 17, 12, 17, 100, 9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 5, 18, 108, 8, 18, 10, 18, 12, 18, 111, 9, 18, 1, 18, 1, 18, 1,
		19, 4, 19, 116, 8, 19, 11, 19, 12, 19, 117, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 5, 20, 126, 8, 20, 10, 20, 12, 20, 129, 9, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 138, 8, 21, 10, 21, 12, 21,
		141, 9, 21, 1, 21, 4, 21, 144, 8, 21, 11, 21, 12, 21, 145, 1, 21, 3, 21,
		149, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22, 155, 8, 22, 10, 22, 12,
		22, 158, 9, 22, 1, 23, 1, 23, 5, 23, 162, 8, 23, 10, 23, 12, 23, 165, 9,
		23, 1, 24, 1, 24, 3, 24, 169, 8, 24, 4, 98, 109, 127, 139, 0, 25, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 14, 41, 15, 43,
		16, 45, 17, 47, 18, 49, 19, 1, 0, 5, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 42, 43, 45, 57, 92, 92, 94, 94, 3, 0, 33, 33, 36, 36, 64, 64, 3, 0,
		9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 177, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 1, 51, 1, 0, 0, 0, 3, 53, 1,
		0, 0, 0, 5, 55, 1, 0, 0, 0, 7, 57, 1, 0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 63,
		1, 0, 0, 0, 13, 65, 1, 0, 0, 0, 15, 67, 1, 0, 0, 0, 17, 69, 1, 0, 0, 0,
		19, 71, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 75, 1, 0, 0, 0, 25, 78, 1,
		0, 0, 0, 27, 81, 1, 0, 0, 0, 29, 83, 1, 0, 0, 0, 31, 85, 1, 0, 0, 0, 33,
		90, 1, 0, 0, 0, 35, 92, 1, 0, 0, 0, 37, 103, 1, 0, 0, 0, 39, 115, 1, 0,
		0, 0, 41, 121, 1, 0, 0, 0, 43, 135, 1, 0, 0, 0, 45, 152, 1, 0, 0, 0, 47,
		159, 1, 0, 0, 0, 49, 168, 1, 0, 0, 0, 51, 52, 5, 123, 0, 0, 52, 2, 1, 0,
		0, 0, 53, 54, 5, 125, 0, 0, 54, 4, 1, 0, 0, 0, 55, 56, 5, 64, 0, 0, 56,
		6, 1, 0, 0, 0, 57, 58, 5, 61, 0, 0, 58, 59, 5, 61, 0, 0, 59, 8, 1, 0, 0,
		0, 60, 61, 5, 33, 0, 0, 61, 62, 5, 61, 0, 0, 62, 10, 1, 0, 0, 0, 63, 64,
		5, 58, 0, 0, 64, 12, 1, 0, 0, 0, 65, 66, 5, 61, 0, 0, 66, 14, 1, 0, 0,
		0, 67, 68, 5, 33, 0, 0, 68, 16, 1, 0, 0, 0, 69, 70, 5, 40, 0, 0, 70, 18,
		1, 0, 0, 0, 71, 72, 5, 41, 0, 0, 72, 20, 1, 0, 0, 0, 73, 74, 5, 44, 0,
		0, 74, 22, 1, 0, 0, 0, 75, 76, 5, 45, 0, 0, 76, 77, 5, 62, 0, 0, 77, 24,
		1, 0, 0, 0, 78, 79, 5, 38, 0, 0, 79, 80, 5, 38, 0, 0, 80, 26, 1, 0, 0,
		0, 81, 82, 7, 0, 0, 0, 82, 28, 1, 0, 0, 0, 83, 84, 7, 1, 0, 0, 84, 30,
		1, 0, 0, 0, 85, 86, 7, 2, 0, 0, 86, 32, 1, 0, 0, 0, 87, 91, 3, 27, 13,
		0, 88, 91, 3, 29, 14, 0, 89, 91, 3, 31, 15, 0, 90, 87, 1, 0, 0, 0, 90,
		88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 34, 1, 0, 0, 0, 92, 98, 5, 34,
		0, 0, 93, 94, 5, 92, 0, 0, 94, 97, 5, 34, 0, 0, 95, 97, 9, 0, 0, 0, 96,
		93, 1, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 99, 1, 0,
		0, 0, 98, 96, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101,
		102, 5, 34, 0, 0, 102, 36, 1, 0, 0, 0, 103, 109, 5, 39, 0, 0, 104, 105,
		5, 92, 0, 0, 105, 108, 5, 39, 0, 0, 106, 108, 9, 0, 0, 0, 107, 104, 1,
		0, 0, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 110, 1, 0, 0,
		0, 109, 107, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112,
		113, 5, 39, 0, 0, 113, 38, 1, 0, 0, 0, 114, 116, 7, 3, 0, 0, 115, 114,
		1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0,
		0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 6, 19, 0, 0, 120, 40, 1, 0, 0, 0,
		121, 122, 5, 47, 0, 0, 122, 123, 5, 42, 0, 0, 123, 127, 1, 0, 0, 0, 124,
		126, 9, 0, 0, 0, 125, 124, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 130, 1, 0, 0, 0, 129, 127, 1, 0,
		0, 0, 130, 131, 5, 42, 0, 0, 131, 132, 5, 47, 0, 0, 132, 133, 1, 0, 0,
		0, 133, 134, 6, 20, 0, 0, 134, 42, 1, 0, 0, 0, 135, 139, 5, 35, 0, 0, 136,
		138, 9, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 148, 1, 0, 0, 0, 141, 139, 1, 0,
		0, 0, 142, 144, 7, 4, 0, 0, 143, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0,
		145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147,
		149, 5, 0, 0, 1, 148, 143, 1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 150,
		1, 0, 0, 0, 150, 151, 6, 21, 0, 0, 151, 44, 1, 0, 0, 0, 152, 156, 3, 27,
		13, 0, 153, 155, 3, 33, 16, 0, 154, 153, 1, 0, 0, 0, 155, 158, 1, 0, 0,
		0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 46, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 159, 163, 3, 29, 14, 0, 160, 162, 3, 33, 16, 0, 161, 160,
		1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 48, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 169, 3, 35, 17, 0,
		167, 169, 3, 37, 18, 0, 168, 166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169,
		50, 1, 0, 0, 0, 14, 0, 90, 96, 98, 107, 109, 117, 127, 139, 145, 148, 156,
		163, 168, 1, 6, 0, 0,
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

// routingALexerInit initializes any static state used to implement routingALexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewroutingALexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RoutingALexerInit() {
	staticData := &routingalexerLexerStaticData
	staticData.once.Do(routingalexerLexerInit)
}

// NewroutingALexer produces a new lexer instance for the optional input antlr.CharStream.
func NewroutingALexer(input antlr.CharStream) *routingALexer {
	RoutingALexerInit()
	l := new(routingALexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &routingalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "routingA.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// routingALexer tokens.
const (
	routingALexerT__0               = 1
	routingALexerT__1               = 2
	routingALexerT__2               = 3
	routingALexerT__3               = 4
	routingALexerT__4               = 5
	routingALexerT__5               = 6
	routingALexerT__6               = 7
	routingALexerT__7               = 8
	routingALexerT__8               = 9
	routingALexerT__9               = 10
	routingALexerT__10              = 11
	routingALexerT__11              = 12
	routingALexerT__12              = 13
	routingALexerWHITESPACE         = 14
	routingALexerCOMMENT_BLOCK      = 15
	routingALexerCOMMENT_LINE_SHARP = 16
	routingALexerID                 = 17
	routingALexerNON_ID             = 18
	routingALexerQUOTE_STRING       = 19
)
