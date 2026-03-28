// Code generated from staircase.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // staircase

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

type staircaseParser struct {
	*antlr.BaseParser
}

var StaircaseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func staircaseParserInit() {
	staticData := &StaircaseParserStaticData
	staticData.LiteralNames = []string{
		"", "'click'", "'collect'", "'input'", "'parent'", "'child'", "'drilldown'",
		"'sibling'", "'keywords'", "'['", "']'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "CLICK", "COLLECT", "INPUT", "PARENT", "CHILD", "DRILLDOWN", "SIBLING",
		"KEYWORDS", "LBRACK", "RBRACK", "AT", "URL", "STRING", "INT", "WS",
	}
	staticData.RuleNames = []string{
		"job", "rule_navigation", "rule_navigation_sequence", "rule_keyword_list",
		"rule_site", "rule_type_declaration", "rule_in_value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 50, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 3, 0, 17, 8, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 3, 0, 23, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 31, 8,
		1, 1, 2, 4, 2, 34, 8, 2, 11, 2, 12, 2, 35, 1, 3, 4, 3, 39, 8, 3, 11, 3,
		12, 3, 40, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0,
		2, 4, 6, 8, 10, 12, 0, 2, 1, 0, 4, 7, 1, 0, 1, 3, 47, 0, 14, 1, 0, 0, 0,
		2, 26, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 42, 1, 0, 0,
		0, 10, 45, 1, 0, 0, 0, 12, 47, 1, 0, 0, 0, 14, 16, 3, 10, 5, 0, 15, 17,
		3, 8, 4, 0, 16, 15, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0,
		18, 19, 5, 8, 0, 0, 19, 20, 3, 6, 3, 0, 20, 22, 3, 4, 2, 0, 21, 23, 3,
		12, 6, 0, 22, 21, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24,
		25, 5, 0, 0, 1, 25, 1, 1, 0, 0, 0, 26, 30, 7, 0, 0, 0, 27, 28, 5, 9, 0,
		0, 28, 29, 5, 14, 0, 0, 29, 31, 5, 10, 0, 0, 30, 27, 1, 0, 0, 0, 30, 31,
		1, 0, 0, 0, 31, 3, 1, 0, 0, 0, 32, 34, 3, 2, 1, 0, 33, 32, 1, 0, 0, 0,
		34, 35, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 5, 1, 0,
		0, 0, 37, 39, 5, 13, 0, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40,
		38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 7, 1, 0, 0, 0, 42, 43, 5, 11, 0,
		0, 43, 44, 5, 12, 0, 0, 44, 9, 1, 0, 0, 0, 45, 46, 7, 1, 0, 0, 46, 11,
		1, 0, 0, 0, 47, 48, 5, 13, 0, 0, 48, 13, 1, 0, 0, 0, 5, 16, 22, 30, 35,
		40,
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

// staircaseParserInit initializes any static state used to implement staircaseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewstaircaseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StaircaseParserInit() {
	staticData := &StaircaseParserStaticData
	staticData.once.Do(staircaseParserInit)
}

// NewstaircaseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewstaircaseParser(input antlr.TokenStream) *staircaseParser {
	StaircaseParserInit()
	this := new(staircaseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &StaircaseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "staircase.g4"

	return this
}

// staircaseParser tokens.
const (
	staircaseParserEOF       = antlr.TokenEOF
	staircaseParserCLICK     = 1
	staircaseParserCOLLECT   = 2
	staircaseParserINPUT     = 3
	staircaseParserPARENT    = 4
	staircaseParserCHILD     = 5
	staircaseParserDRILLDOWN = 6
	staircaseParserSIBLING   = 7
	staircaseParserKEYWORDS  = 8
	staircaseParserLBRACK    = 9
	staircaseParserRBRACK    = 10
	staircaseParserAT        = 11
	staircaseParserURL       = 12
	staircaseParserSTRING    = 13
	staircaseParserINT       = 14
	staircaseParserWS        = 15
)

// staircaseParser rules.
const (
	staircaseParserRULE_job                      = 0
	staircaseParserRULE_rule_navigation          = 1
	staircaseParserRULE_rule_navigation_sequence = 2
	staircaseParserRULE_rule_keyword_list        = 3
	staircaseParserRULE_rule_site                = 4
	staircaseParserRULE_rule_type_declaration    = 5
	staircaseParserRULE_rule_in_value            = 6
)

// IJobContext is an interface to support dynamic dispatch.
type IJobContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Rule_type_declaration() IRule_type_declarationContext
	KEYWORDS() antlr.TerminalNode
	Rule_keyword_list() IRule_keyword_listContext
	Rule_navigation_sequence() IRule_navigation_sequenceContext
	EOF() antlr.TerminalNode
	Rule_site() IRule_siteContext
	Rule_in_value() IRule_in_valueContext

	// IsJobContext differentiates from other interfaces.
	IsJobContext()
}

type JobContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJobContext() *JobContext {
	var p = new(JobContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_job
	return p
}

func InitEmptyJobContext(p *JobContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_job
}

func (*JobContext) IsJobContext() {}

func NewJobContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JobContext {
	var p = new(JobContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = staircaseParserRULE_job

	return p
}

func (s *JobContext) GetParser() antlr.Parser { return s.parser }

func (s *JobContext) Rule_type_declaration() IRule_type_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRule_type_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRule_type_declarationContext)
}

func (s *JobContext) KEYWORDS() antlr.TerminalNode {
	return s.GetToken(staircaseParserKEYWORDS, 0)
}

func (s *JobContext) Rule_keyword_list() IRule_keyword_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRule_keyword_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRule_keyword_listContext)
}

func (s *JobContext) Rule_navigation_sequence() IRule_navigation_sequenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRule_navigation_sequenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRule_navigation_sequenceContext)
}

func (s *JobContext) EOF() antlr.TerminalNode {
	return s.GetToken(staircaseParserEOF, 0)
}

func (s *JobContext) Rule_site() IRule_siteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRule_siteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRule_siteContext)
}

func (s *JobContext) Rule_in_value() IRule_in_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRule_in_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRule_in_valueContext)
}

func (s *JobContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JobContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JobContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.EnterJob(s)
	}
}

func (s *JobContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.ExitJob(s)
	}
}

func (p *staircaseParser) Job() (localctx IJobContext) {
	localctx = NewJobContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, staircaseParserRULE_job)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Rule_type_declaration()
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == staircaseParserAT {
		{
			p.SetState(15)
			p.Rule_site()
		}

	}
	{
		p.SetState(18)
		p.Match(staircaseParserKEYWORDS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Rule_keyword_list()
	}
	{
		p.SetState(20)
		p.Rule_navigation_sequence()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == staircaseParserSTRING {
		{
			p.SetState(21)
			p.Rule_in_value()
		}

	}
	{
		p.SetState(24)
		p.Match(staircaseParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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

// IRule_navigationContext is an interface to support dynamic dispatch.
type IRule_navigationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARENT() antlr.TerminalNode
	CHILD() antlr.TerminalNode
	DRILLDOWN() antlr.TerminalNode
	SIBLING() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	INT() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsRule_navigationContext differentiates from other interfaces.
	IsRule_navigationContext()
}

type Rule_navigationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_navigationContext() *Rule_navigationContext {
	var p = new(Rule_navigationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_navigation
	return p
}

func InitEmptyRule_navigationContext(p *Rule_navigationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_navigation
}

func (*Rule_navigationContext) IsRule_navigationContext() {}

func NewRule_navigationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_navigationContext {
	var p = new(Rule_navigationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = staircaseParserRULE_rule_navigation

	return p
}

func (s *Rule_navigationContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_navigationContext) PARENT() antlr.TerminalNode {
	return s.GetToken(staircaseParserPARENT, 0)
}

func (s *Rule_navigationContext) CHILD() antlr.TerminalNode {
	return s.GetToken(staircaseParserCHILD, 0)
}

func (s *Rule_navigationContext) DRILLDOWN() antlr.TerminalNode {
	return s.GetToken(staircaseParserDRILLDOWN, 0)
}

func (s *Rule_navigationContext) SIBLING() antlr.TerminalNode {
	return s.GetToken(staircaseParserSIBLING, 0)
}

func (s *Rule_navigationContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(staircaseParserLBRACK, 0)
}

func (s *Rule_navigationContext) INT() antlr.TerminalNode {
	return s.GetToken(staircaseParserINT, 0)
}

func (s *Rule_navigationContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(staircaseParserRBRACK, 0)
}

func (s *Rule_navigationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_navigationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_navigationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.EnterRule_navigation(s)
	}
}

func (s *Rule_navigationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.ExitRule_navigation(s)
	}
}

func (p *staircaseParser) Rule_navigation() (localctx IRule_navigationContext) {
	localctx = NewRule_navigationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, staircaseParserRULE_rule_navigation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == staircaseParserLBRACK {
		{
			p.SetState(27)
			p.Match(staircaseParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.Match(staircaseParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Match(staircaseParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IRule_navigation_sequenceContext is an interface to support dynamic dispatch.
type IRule_navigation_sequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRule_navigation() []IRule_navigationContext
	Rule_navigation(i int) IRule_navigationContext

	// IsRule_navigation_sequenceContext differentiates from other interfaces.
	IsRule_navigation_sequenceContext()
}

type Rule_navigation_sequenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_navigation_sequenceContext() *Rule_navigation_sequenceContext {
	var p = new(Rule_navigation_sequenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_navigation_sequence
	return p
}

func InitEmptyRule_navigation_sequenceContext(p *Rule_navigation_sequenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_navigation_sequence
}

func (*Rule_navigation_sequenceContext) IsRule_navigation_sequenceContext() {}

func NewRule_navigation_sequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_navigation_sequenceContext {
	var p = new(Rule_navigation_sequenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = staircaseParserRULE_rule_navigation_sequence

	return p
}

func (s *Rule_navigation_sequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_navigation_sequenceContext) AllRule_navigation() []IRule_navigationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRule_navigationContext); ok {
			len++
		}
	}

	tst := make([]IRule_navigationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRule_navigationContext); ok {
			tst[i] = t.(IRule_navigationContext)
			i++
		}
	}

	return tst
}

func (s *Rule_navigation_sequenceContext) Rule_navigation(i int) IRule_navigationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRule_navigationContext); ok {
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

	return t.(IRule_navigationContext)
}

func (s *Rule_navigation_sequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_navigation_sequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_navigation_sequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.EnterRule_navigation_sequence(s)
	}
}

func (s *Rule_navigation_sequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.ExitRule_navigation_sequence(s)
	}
}

func (p *staircaseParser) Rule_navigation_sequence() (localctx IRule_navigation_sequenceContext) {
	localctx = NewRule_navigation_sequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, staircaseParserRULE_rule_navigation_sequence)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240) != 0) {
		{
			p.SetState(32)
			p.Rule_navigation()
		}

		p.SetState(35)
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

// IRule_keyword_listContext is an interface to support dynamic dispatch.
type IRule_keyword_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsRule_keyword_listContext differentiates from other interfaces.
	IsRule_keyword_listContext()
}

type Rule_keyword_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_keyword_listContext() *Rule_keyword_listContext {
	var p = new(Rule_keyword_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_keyword_list
	return p
}

func InitEmptyRule_keyword_listContext(p *Rule_keyword_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_keyword_list
}

func (*Rule_keyword_listContext) IsRule_keyword_listContext() {}

func NewRule_keyword_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_keyword_listContext {
	var p = new(Rule_keyword_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = staircaseParserRULE_rule_keyword_list

	return p
}

func (s *Rule_keyword_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_keyword_listContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(staircaseParserSTRING)
}

func (s *Rule_keyword_listContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(staircaseParserSTRING, i)
}

func (s *Rule_keyword_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_keyword_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_keyword_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.EnterRule_keyword_list(s)
	}
}

func (s *Rule_keyword_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.ExitRule_keyword_list(s)
	}
}

func (p *staircaseParser) Rule_keyword_list() (localctx IRule_keyword_listContext) {
	localctx = NewRule_keyword_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, staircaseParserRULE_rule_keyword_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == staircaseParserSTRING {
		{
			p.SetState(37)
			p.Match(staircaseParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(40)
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

// IRule_siteContext is an interface to support dynamic dispatch.
type IRule_siteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	URL() antlr.TerminalNode

	// IsRule_siteContext differentiates from other interfaces.
	IsRule_siteContext()
}

type Rule_siteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_siteContext() *Rule_siteContext {
	var p = new(Rule_siteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_site
	return p
}

func InitEmptyRule_siteContext(p *Rule_siteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_site
}

func (*Rule_siteContext) IsRule_siteContext() {}

func NewRule_siteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_siteContext {
	var p = new(Rule_siteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = staircaseParserRULE_rule_site

	return p
}

func (s *Rule_siteContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_siteContext) AT() antlr.TerminalNode {
	return s.GetToken(staircaseParserAT, 0)
}

func (s *Rule_siteContext) URL() antlr.TerminalNode {
	return s.GetToken(staircaseParserURL, 0)
}

func (s *Rule_siteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_siteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_siteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.EnterRule_site(s)
	}
}

func (s *Rule_siteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.ExitRule_site(s)
	}
}

func (p *staircaseParser) Rule_site() (localctx IRule_siteContext) {
	localctx = NewRule_siteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, staircaseParserRULE_rule_site)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(staircaseParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(staircaseParserURL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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

// IRule_type_declarationContext is an interface to support dynamic dispatch.
type IRule_type_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLICK() antlr.TerminalNode
	COLLECT() antlr.TerminalNode
	INPUT() antlr.TerminalNode

	// IsRule_type_declarationContext differentiates from other interfaces.
	IsRule_type_declarationContext()
}

type Rule_type_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_type_declarationContext() *Rule_type_declarationContext {
	var p = new(Rule_type_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_type_declaration
	return p
}

func InitEmptyRule_type_declarationContext(p *Rule_type_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_type_declaration
}

func (*Rule_type_declarationContext) IsRule_type_declarationContext() {}

func NewRule_type_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_type_declarationContext {
	var p = new(Rule_type_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = staircaseParserRULE_rule_type_declaration

	return p
}

func (s *Rule_type_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_type_declarationContext) CLICK() antlr.TerminalNode {
	return s.GetToken(staircaseParserCLICK, 0)
}

func (s *Rule_type_declarationContext) COLLECT() antlr.TerminalNode {
	return s.GetToken(staircaseParserCOLLECT, 0)
}

func (s *Rule_type_declarationContext) INPUT() antlr.TerminalNode {
	return s.GetToken(staircaseParserINPUT, 0)
}

func (s *Rule_type_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_type_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_type_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.EnterRule_type_declaration(s)
	}
}

func (s *Rule_type_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.ExitRule_type_declaration(s)
	}
}

func (p *staircaseParser) Rule_type_declaration() (localctx IRule_type_declarationContext) {
	localctx = NewRule_type_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, staircaseParserRULE_rule_type_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
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

// IRule_in_valueContext is an interface to support dynamic dispatch.
type IRule_in_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsRule_in_valueContext differentiates from other interfaces.
	IsRule_in_valueContext()
}

type Rule_in_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRule_in_valueContext() *Rule_in_valueContext {
	var p = new(Rule_in_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_in_value
	return p
}

func InitEmptyRule_in_valueContext(p *Rule_in_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = staircaseParserRULE_rule_in_value
}

func (*Rule_in_valueContext) IsRule_in_valueContext() {}

func NewRule_in_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rule_in_valueContext {
	var p = new(Rule_in_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = staircaseParserRULE_rule_in_value

	return p
}

func (s *Rule_in_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Rule_in_valueContext) STRING() antlr.TerminalNode {
	return s.GetToken(staircaseParserSTRING, 0)
}

func (s *Rule_in_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rule_in_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rule_in_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.EnterRule_in_value(s)
	}
}

func (s *Rule_in_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(staircaseListener); ok {
		listenerT.ExitRule_in_value(s)
	}
}

func (p *staircaseParser) Rule_in_value() (localctx IRule_in_valueContext) {
	localctx = NewRule_in_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, staircaseParserRULE_rule_in_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(staircaseParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
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
