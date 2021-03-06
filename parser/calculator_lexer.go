// Code generated from .\Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 55, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 6, 9, 37, 10,
	9, 13, 9, 14, 9, 38, 3, 9, 3, 9, 6, 9, 43, 10, 9, 13, 9, 14, 9, 44, 5,
	9, 47, 10, 9, 3, 10, 6, 10, 50, 10, 10, 13, 10, 14, 10, 51, 3, 10, 3, 10,
	2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	3, 2, 3, 5, 2, 11, 12, 15, 15, 34, 34, 2, 58, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3,
	21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 27, 3, 2, 2, 2,
	11, 29, 3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15, 33, 3, 2, 2, 2, 17, 36, 3,
	2, 2, 2, 19, 49, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 4, 3, 2, 2, 2, 23,
	24, 7, 43, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 45, 2, 2, 26, 8, 3, 2, 2,
	2, 27, 28, 7, 47, 2, 2, 28, 10, 3, 2, 2, 2, 29, 30, 7, 44, 2, 2, 30, 12,
	3, 2, 2, 2, 31, 32, 7, 49, 2, 2, 32, 14, 3, 2, 2, 2, 33, 34, 7, 96, 2,
	2, 34, 16, 3, 2, 2, 2, 35, 37, 4, 50, 59, 2, 36, 35, 3, 2, 2, 2, 37, 38,
	3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 46, 3, 2, 2, 2,
	40, 42, 7, 48, 2, 2, 41, 43, 4, 50, 59, 2, 42, 41, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2,
	46, 40, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 18, 3, 2, 2, 2, 48, 50, 9,
	2, 2, 2, 49, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 8, 10, 2, 2, 54, 20, 3, 2,
	2, 2, 7, 2, 38, 44, 46, 51, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'^'",
}

var lexerSymbolicNames = []string{
	"", "LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "POW", "NUMBER",
	"WS",
}

var lexerRuleNames = []string{
	"LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "POW", "NUMBER", "WS",
}

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {

	l := new(CalculatorLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerLPAREN = 1
	CalculatorLexerRPAREN = 2
	CalculatorLexerPLUS   = 3
	CalculatorLexerMINUS  = 4
	CalculatorLexerTIMES  = 5
	CalculatorLexerDIV    = 6
	CalculatorLexerPOW    = 7
	CalculatorLexerNUMBER = 8
	CalculatorLexerWS     = 9
)
