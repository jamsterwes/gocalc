// Code generated from .\Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Calculator

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 57, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 7, 3, 22, 10, 3, 12, 3, 14, 3, 25,
	11, 3, 3, 4, 3, 4, 3, 4, 7, 4, 30, 10, 4, 12, 4, 14, 4, 33, 11, 4, 3, 5,
	3, 5, 3, 5, 7, 5, 38, 10, 5, 12, 5, 14, 5, 41, 11, 5, 3, 6, 3, 6, 3, 6,
	5, 6, 46, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 53, 10, 7, 3, 8, 3,
	8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 4, 3, 2, 5, 6, 3, 2, 7, 8,
	2, 54, 2, 16, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 34,
	3, 2, 2, 2, 10, 45, 3, 2, 2, 2, 12, 52, 3, 2, 2, 2, 14, 54, 3, 2, 2, 2,
	16, 17, 5, 4, 3, 2, 17, 3, 3, 2, 2, 2, 18, 23, 5, 6, 4, 2, 19, 20, 9, 2,
	2, 2, 20, 22, 5, 6, 4, 2, 21, 19, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21,
	3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2,
	26, 31, 5, 8, 5, 2, 27, 28, 9, 3, 2, 2, 28, 30, 5, 8, 5, 2, 29, 27, 3,
	2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32,
	7, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 39, 5, 10, 6, 2, 35, 36, 7, 9, 2,
	2, 36, 38, 5, 10, 6, 2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 9, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2,
	42, 43, 7, 6, 2, 2, 43, 46, 5, 12, 7, 2, 44, 46, 5, 12, 7, 2, 45, 42, 3,
	2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 11, 3, 2, 2, 2, 47, 48, 7, 3, 2, 2, 48,
	49, 5, 4, 3, 2, 49, 50, 7, 4, 2, 2, 50, 53, 3, 2, 2, 2, 51, 53, 5, 14,
	8, 2, 52, 47, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 13, 3, 2, 2, 2, 54, 55,
	7, 10, 2, 2, 55, 15, 3, 2, 2, 2, 7, 23, 31, 39, 45, 52,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "'^'",
}
var symbolicNames = []string{
	"", "LPAREN", "RPAREN", "PLUS", "MINUS", "TIMES", "DIV", "POW", "NUMBER",
	"WS",
}

var ruleNames = []string{
	"root", "expr", "mult_expr", "pow_expr", "sign_atom", "atom", "constant",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CalculatorParser struct {
	*antlr.BaseParser
}

func NewCalculatorParser(input antlr.TokenStream) *CalculatorParser {
	this := new(CalculatorParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calculator.g4"

	return this
}

// CalculatorParser tokens.
const (
	CalculatorParserEOF    = antlr.TokenEOF
	CalculatorParserLPAREN = 1
	CalculatorParserRPAREN = 2
	CalculatorParserPLUS   = 3
	CalculatorParserMINUS  = 4
	CalculatorParserTIMES  = 5
	CalculatorParserDIV    = 6
	CalculatorParserPOW    = 7
	CalculatorParserNUMBER = 8
	CalculatorParserWS     = 9
)

// CalculatorParser rules.
const (
	CalculatorParserRULE_root      = 0
	CalculatorParserRULE_expr      = 1
	CalculatorParserRULE_mult_expr = 2
	CalculatorParserRULE_pow_expr  = 3
	CalculatorParserRULE_sign_atom = 4
	CalculatorParserRULE_atom      = 5
	CalculatorParserRULE_constant  = 6
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalculatorParserRULE_root)

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
		p.Expr()
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllMult_expr() []IMult_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMult_exprContext)(nil)).Elem())
	var tst = make([]IMult_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMult_exprContext)
		}
	}

	return tst
}

func (s *ExprContext) Mult_expr(i int) IMult_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMult_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMult_exprContext)
}

func (s *ExprContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(CalculatorParserPLUS)
}

func (s *ExprContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(CalculatorParserPLUS, i)
}

func (s *ExprContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(CalculatorParserMINUS)
}

func (s *ExprContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(CalculatorParserMINUS, i)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalculatorParserRULE_expr)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Mult_expr()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CalculatorParserPLUS || _la == CalculatorParserMINUS {
		{
			p.SetState(17)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CalculatorParserPLUS || _la == CalculatorParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(18)
			p.Mult_expr()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMult_exprContext is an interface to support dynamic dispatch.
type IMult_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMult_exprContext differentiates from other interfaces.
	IsMult_exprContext()
}

type Mult_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMult_exprContext() *Mult_exprContext {
	var p = new(Mult_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_mult_expr
	return p
}

func (*Mult_exprContext) IsMult_exprContext() {}

func NewMult_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mult_exprContext {
	var p = new(Mult_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_mult_expr

	return p
}

func (s *Mult_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Mult_exprContext) AllPow_expr() []IPow_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPow_exprContext)(nil)).Elem())
	var tst = make([]IPow_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPow_exprContext)
		}
	}

	return tst
}

func (s *Mult_exprContext) Pow_expr(i int) IPow_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPow_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPow_exprContext)
}

func (s *Mult_exprContext) AllTIMES() []antlr.TerminalNode {
	return s.GetTokens(CalculatorParserTIMES)
}

func (s *Mult_exprContext) TIMES(i int) antlr.TerminalNode {
	return s.GetToken(CalculatorParserTIMES, i)
}

func (s *Mult_exprContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(CalculatorParserDIV)
}

func (s *Mult_exprContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(CalculatorParserDIV, i)
}

func (s *Mult_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mult_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mult_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitMult_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Mult_expr() (localctx IMult_exprContext) {
	localctx = NewMult_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalculatorParserRULE_mult_expr)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Pow_expr()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CalculatorParserTIMES || _la == CalculatorParserDIV {
		{
			p.SetState(25)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CalculatorParserTIMES || _la == CalculatorParserDIV) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(26)
			p.Pow_expr()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPow_exprContext is an interface to support dynamic dispatch.
type IPow_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPow_exprContext differentiates from other interfaces.
	IsPow_exprContext()
}

type Pow_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPow_exprContext() *Pow_exprContext {
	var p = new(Pow_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_pow_expr
	return p
}

func (*Pow_exprContext) IsPow_exprContext() {}

func NewPow_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pow_exprContext {
	var p = new(Pow_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_pow_expr

	return p
}

func (s *Pow_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Pow_exprContext) AllSign_atom() []ISign_atomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISign_atomContext)(nil)).Elem())
	var tst = make([]ISign_atomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISign_atomContext)
		}
	}

	return tst
}

func (s *Pow_exprContext) Sign_atom(i int) ISign_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISign_atomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISign_atomContext)
}

func (s *Pow_exprContext) AllPOW() []antlr.TerminalNode {
	return s.GetTokens(CalculatorParserPOW)
}

func (s *Pow_exprContext) POW(i int) antlr.TerminalNode {
	return s.GetToken(CalculatorParserPOW, i)
}

func (s *Pow_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pow_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pow_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitPow_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Pow_expr() (localctx IPow_exprContext) {
	localctx = NewPow_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalculatorParserRULE_pow_expr)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Sign_atom()
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CalculatorParserPOW {
		{
			p.SetState(33)
			p.Match(CalculatorParserPOW)
		}
		{
			p.SetState(34)
			p.Sign_atom()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISign_atomContext is an interface to support dynamic dispatch.
type ISign_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSign_atomContext differentiates from other interfaces.
	IsSign_atomContext()
}

type Sign_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySign_atomContext() *Sign_atomContext {
	var p = new(Sign_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_sign_atom
	return p
}

func (*Sign_atomContext) IsSign_atomContext() {}

func NewSign_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sign_atomContext {
	var p = new(Sign_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_sign_atom

	return p
}

func (s *Sign_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Sign_atomContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMINUS, 0)
}

func (s *Sign_atomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Sign_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sign_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sign_atomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitSign_atom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Sign_atom() (localctx ISign_atomContext) {
	localctx = NewSign_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalculatorParserRULE_sign_atom)

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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Match(CalculatorParserMINUS)
		}
		{
			p.SetState(41)
			p.Atom()
		}

	case CalculatorParserLPAREN, CalculatorParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLPAREN, 0)
}

func (s *AtomContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalculatorParserRPAREN, 0)
}

func (s *AtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalculatorParserRULE_atom)

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

	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(CalculatorParserLPAREN)
		}
		{
			p.SetState(46)
			p.Expr()
		}
		{
			p.SetState(47)
			p.Match(CalculatorParserRPAREN)
		}

	case CalculatorParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Constant()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNUMBER, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalculatorParserRULE_constant)

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
		p.SetState(52)
		p.Match(CalculatorParserNUMBER)
	}

	return localctx
}
