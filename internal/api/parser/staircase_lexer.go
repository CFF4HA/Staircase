// Code generated from staircase.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type staircaseLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var StaircaseLexerLexerStaticData struct {
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

func staircaselexerLexerInit() {
	staticData := &StaircaseLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'keywords'", "'['", "']'", "'@'", "'click'", "'collect'", "'input'",
		"'parent'", "'child'", "'drilldown'", "'sibling'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "CLICK", "COLLECT", "INPUT", "PARENT", "CHILD",
		"DRILLDOWN", "SIBLING", "URL", "STRING", "IDENTIFIER", "NUMBER", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "CLICK", "COLLECT", "INPUT", "PARENT",
		"CHILD", "DRILLDOWN", "SIBLING", "URL", "STRING", "IDENTIFIER", "NUMBER",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 153, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 106, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		4, 11, 113, 8, 11, 11, 11, 12, 11, 114, 1, 12, 1, 12, 5, 12, 119, 8, 12,
		10, 12, 12, 12, 122, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 128, 8,
		13, 10, 13, 12, 13, 131, 9, 13, 1, 14, 4, 14, 134, 8, 14, 11, 14, 12, 14,
		135, 1, 14, 1, 14, 5, 14, 140, 8, 14, 10, 14, 12, 14, 143, 9, 14, 3, 14,
		145, 8, 14, 1, 15, 4, 15, 148, 8, 15, 11, 15, 12, 15, 149, 1, 15, 1, 15,
		0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 1, 0, 6, 7, 0, 35,
		35, 45, 57, 61, 61, 63, 63, 65, 90, 95, 95, 97, 122, 3, 0, 10, 10, 13,
		13, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 160, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 33,
		1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 5, 44, 1, 0, 0, 0, 7, 46, 1, 0, 0, 0, 9,
		48, 1, 0, 0, 0, 11, 54, 1, 0, 0, 0, 13, 62, 1, 0, 0, 0, 15, 68, 1, 0, 0,
		0, 17, 75, 1, 0, 0, 0, 19, 81, 1, 0, 0, 0, 21, 91, 1, 0, 0, 0, 23, 99,
		1, 0, 0, 0, 25, 116, 1, 0, 0, 0, 27, 125, 1, 0, 0, 0, 29, 133, 1, 0, 0,
		0, 31, 147, 1, 0, 0, 0, 33, 34, 5, 107, 0, 0, 34, 35, 5, 101, 0, 0, 35,
		36, 5, 121, 0, 0, 36, 37, 5, 119, 0, 0, 37, 38, 5, 111, 0, 0, 38, 39, 5,
		114, 0, 0, 39, 40, 5, 100, 0, 0, 40, 41, 5, 115, 0, 0, 41, 2, 1, 0, 0,
		0, 42, 43, 5, 91, 0, 0, 43, 4, 1, 0, 0, 0, 44, 45, 5, 93, 0, 0, 45, 6,
		1, 0, 0, 0, 46, 47, 5, 64, 0, 0, 47, 8, 1, 0, 0, 0, 48, 49, 5, 99, 0, 0,
		49, 50, 5, 108, 0, 0, 50, 51, 5, 105, 0, 0, 51, 52, 5, 99, 0, 0, 52, 53,
		5, 107, 0, 0, 53, 10, 1, 0, 0, 0, 54, 55, 5, 99, 0, 0, 55, 56, 5, 111,
		0, 0, 56, 57, 5, 108, 0, 0, 57, 58, 5, 108, 0, 0, 58, 59, 5, 101, 0, 0,
		59, 60, 5, 99, 0, 0, 60, 61, 5, 116, 0, 0, 61, 12, 1, 0, 0, 0, 62, 63,
		5, 105, 0, 0, 63, 64, 5, 110, 0, 0, 64, 65, 5, 112, 0, 0, 65, 66, 5, 117,
		0, 0, 66, 67, 5, 116, 0, 0, 67, 14, 1, 0, 0, 0, 68, 69, 5, 112, 0, 0, 69,
		70, 5, 97, 0, 0, 70, 71, 5, 114, 0, 0, 71, 72, 5, 101, 0, 0, 72, 73, 5,
		110, 0, 0, 73, 74, 5, 116, 0, 0, 74, 16, 1, 0, 0, 0, 75, 76, 5, 99, 0,
		0, 76, 77, 5, 104, 0, 0, 77, 78, 5, 105, 0, 0, 78, 79, 5, 108, 0, 0, 79,
		80, 5, 100, 0, 0, 80, 18, 1, 0, 0, 0, 81, 82, 5, 100, 0, 0, 82, 83, 5,
		114, 0, 0, 83, 84, 5, 105, 0, 0, 84, 85, 5, 108, 0, 0, 85, 86, 5, 108,
		0, 0, 86, 87, 5, 100, 0, 0, 87, 88, 5, 111, 0, 0, 88, 89, 5, 119, 0, 0,
		89, 90, 5, 110, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92, 5, 115, 0, 0, 92, 93,
		5, 105, 0, 0, 93, 94, 5, 98, 0, 0, 94, 95, 5, 108, 0, 0, 95, 96, 5, 105,
		0, 0, 96, 97, 5, 110, 0, 0, 97, 98, 5, 103, 0, 0, 98, 22, 1, 0, 0, 0, 99,
		100, 5, 104, 0, 0, 100, 101, 5, 116, 0, 0, 101, 102, 5, 116, 0, 0, 102,
		103, 5, 112, 0, 0, 103, 105, 1, 0, 0, 0, 104, 106, 5, 115, 0, 0, 105, 104,
		1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 5, 58,
		0, 0, 108, 109, 5, 47, 0, 0, 109, 110, 5, 47, 0, 0, 110, 112, 1, 0, 0,
		0, 111, 113, 7, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114,
		112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 24, 1, 0, 0, 0, 116, 120, 5,
		34, 0, 0, 117, 119, 8, 1, 0, 0, 118, 117, 1, 0, 0, 0, 119, 122, 1, 0, 0,
		0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122,
		120, 1, 0, 0, 0, 123, 124, 5, 34, 0, 0, 124, 26, 1, 0, 0, 0, 125, 129,
		7, 2, 0, 0, 126, 128, 7, 3, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0,
		0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 28, 1, 0, 0, 0,
		131, 129, 1, 0, 0, 0, 132, 134, 7, 4, 0, 0, 133, 132, 1, 0, 0, 0, 134,
		135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 144,
		1, 0, 0, 0, 137, 141, 5, 46, 0, 0, 138, 140, 7, 4, 0, 0, 139, 138, 1, 0,
		0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0,
		142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 137, 1, 0, 0, 0, 144,
		145, 1, 0, 0, 0, 145, 30, 1, 0, 0, 0, 146, 148, 7, 5, 0, 0, 147, 146, 1,
		0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0,
		0, 150, 151, 1, 0, 0, 0, 151, 152, 6, 15, 0, 0, 152, 32, 1, 0, 0, 0, 9,
		0, 105, 114, 120, 129, 135, 141, 144, 149, 1, 6, 0, 0,
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

// staircaseLexerInit initializes any static state used to implement staircaseLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewstaircaseLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func StaircaseLexerInit() {
	staticData := &StaircaseLexerLexerStaticData
	staticData.once.Do(staircaselexerLexerInit)
}

// NewstaircaseLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewstaircaseLexer(input antlr.CharStream) *staircaseLexer {
	StaircaseLexerInit()
	l := new(staircaseLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &StaircaseLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "staircase.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// staircaseLexer tokens.
const (
	staircaseLexerT__0       = 1
	staircaseLexerT__1       = 2
	staircaseLexerT__2       = 3
	staircaseLexerT__3       = 4
	staircaseLexerCLICK      = 5
	staircaseLexerCOLLECT    = 6
	staircaseLexerINPUT      = 7
	staircaseLexerPARENT     = 8
	staircaseLexerCHILD      = 9
	staircaseLexerDRILLDOWN  = 10
	staircaseLexerSIBLING    = 11
	staircaseLexerURL        = 12
	staircaseLexerSTRING     = 13
	staircaseLexerIDENTIFIER = 14
	staircaseLexerNUMBER     = 15
	staircaseLexerWS         = 16
)
