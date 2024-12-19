// Code generated from parser/AlcorParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AlcorParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AlcorParser struct {
	*antlr.BaseParser
}

var AlcorParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func alcorparserParserInit() {
	staticData := &AlcorParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "STAR", "DIV", "DECIMAL_LIT", "BINARY_LIT", "OCTAL_LIT",
		"HEX_LIT", "SEP",
	}
	staticData.RuleNames = []string{
		"numberLit", "expr", "prog",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 20, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 16, 8, 2, 11, 2, 12, 2, 17, 1, 2, 0,
		0, 3, 0, 2, 4, 0, 2, 1, 0, 5, 8, 1, 0, 1, 4, 17, 0, 6, 1, 0, 0, 0, 2, 8,
		1, 0, 0, 0, 4, 15, 1, 0, 0, 0, 6, 7, 7, 0, 0, 0, 7, 1, 1, 0, 0, 0, 8, 9,
		3, 0, 0, 0, 9, 10, 5, 9, 0, 0, 10, 11, 7, 1, 0, 0, 11, 12, 5, 9, 0, 0,
		12, 13, 3, 0, 0, 0, 13, 3, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 1, 0,
		0, 0, 16, 17, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 5,
		1, 0, 0, 0, 1, 17,
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

// AlcorParserInit initializes any static state used to implement AlcorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAlcorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlcorParserInit() {
	staticData := &AlcorParserParserStaticData
	staticData.once.Do(alcorparserParserInit)
}

// NewAlcorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAlcorParser(input antlr.TokenStream) *AlcorParser {
	AlcorParserInit()
	this := new(AlcorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AlcorParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AlcorParser.g4"

	return this
}

// AlcorParser tokens.
const (
	AlcorParserEOF         = antlr.TokenEOF
	AlcorParserPLUS        = 1
	AlcorParserMINUS       = 2
	AlcorParserSTAR        = 3
	AlcorParserDIV         = 4
	AlcorParserDECIMAL_LIT = 5
	AlcorParserBINARY_LIT  = 6
	AlcorParserOCTAL_LIT   = 7
	AlcorParserHEX_LIT     = 8
	AlcorParserSEP         = 9
)

// AlcorParser rules.
const (
	AlcorParserRULE_numberLit = 0
	AlcorParserRULE_expr      = 1
	AlcorParserRULE_prog      = 2
)

// INumberLitContext is an interface to support dynamic dispatch.
type INumberLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_LIT() antlr.TerminalNode
	BINARY_LIT() antlr.TerminalNode
	OCTAL_LIT() antlr.TerminalNode
	HEX_LIT() antlr.TerminalNode

	// IsNumberLitContext differentiates from other interfaces.
	IsNumberLitContext()
}

type NumberLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLitContext() *NumberLitContext {
	var p = new(NumberLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlcorParserRULE_numberLit
	return p
}

func InitEmptyNumberLitContext(p *NumberLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlcorParserRULE_numberLit
}

func (*NumberLitContext) IsNumberLitContext() {}

func NewNumberLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLitContext {
	var p = new(NumberLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlcorParserRULE_numberLit

	return p
}

func (s *NumberLitContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLitContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(AlcorParserDECIMAL_LIT, 0)
}

func (s *NumberLitContext) BINARY_LIT() antlr.TerminalNode {
	return s.GetToken(AlcorParserBINARY_LIT, 0)
}

func (s *NumberLitContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(AlcorParserOCTAL_LIT, 0)
}

func (s *NumberLitContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(AlcorParserHEX_LIT, 0)
}

func (s *NumberLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlcorParserListener); ok {
		listenerT.EnterNumberLit(s)
	}
}

func (s *NumberLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlcorParserListener); ok {
		listenerT.ExitNumberLit(s)
	}
}

func (p *AlcorParser) NumberLit() (localctx INumberLitContext) {
	localctx = NewNumberLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AlcorParserRULE_numberLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&480) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumberLit() []INumberLitContext
	NumberLit(i int) INumberLitContext
	AllSEP() []antlr.TerminalNode
	SEP(i int) antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	STAR() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlcorParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlcorParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlcorParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllNumberLit() []INumberLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberLitContext); ok {
			len++
		}
	}

	tst := make([]INumberLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberLitContext); ok {
			tst[i] = t.(INumberLitContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) NumberLit(i int) INumberLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberLitContext); ok {
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

	return t.(INumberLitContext)
}

func (s *ExprContext) AllSEP() []antlr.TerminalNode {
	return s.GetTokens(AlcorParserSEP)
}

func (s *ExprContext) SEP(i int) antlr.TerminalNode {
	return s.GetToken(AlcorParserSEP, i)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(AlcorParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(AlcorParserMINUS, 0)
}

func (s *ExprContext) STAR() antlr.TerminalNode {
	return s.GetToken(AlcorParserSTAR, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(AlcorParserDIV, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlcorParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlcorParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *AlcorParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AlcorParserRULE_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.NumberLit()
	}
	{
		p.SetState(9)
		p.Match(AlcorParserSEP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(10)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(11)
		p.Match(AlcorParserSEP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(12)
		p.NumberLit()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlcorParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlcorParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlcorParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlcorParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlcorParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *AlcorParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AlcorParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&480) != 0) {
		{
			p.SetState(14)
			p.Expr()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
