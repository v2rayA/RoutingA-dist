// Code generated from routingA.g4 by ANTLR 4.10.1. DO NOT EDIT.

package routingA // routingA
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type routingAParser struct {
	*antlr.BaseParser
}

var routingaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func routingaParserInit() {
	staticData := &routingaParserStaticData
	staticData.literalNames = []string{
		"", "'{'", "'}'", "'@'", "'=='", "'!='", "'&&'", "':'", "'='", "'('",
		"')'", "','", "'->'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "WHITESPACE", "COMMENT_BLOCK",
		"COMMENT_LINE_SHARP", "ID", "NON_ID", "QUOTE_STRING",
	}
	staticData.ruleNames = []string{
		"start", "bare_literal", "quote_literal", "literal", "input", "programStructureBlcok",
		"expression", "optConditionStatement", "optFunctionPrototypeExpressionAnd",
		"declaration", "assignmentExpression", "functionPrototype", "parameterList",
		"nonEmptyParameterList", "parameter", "routingRule", "functionPrototypeExpression",
		"functionPrototypeExpressionList", "routingRuleOrDeclarationList",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 170, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 48, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 53, 8,
		4, 1, 4, 1, 4, 5, 4, 57, 8, 4, 10, 4, 12, 4, 60, 9, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 72, 8, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 83, 8, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 89, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 97, 8,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 3, 12, 110, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5,
		13, 118, 8, 13, 10, 13, 12, 13, 121, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 130, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 142, 8, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 150, 8, 16, 10, 16, 12, 16, 153, 9,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 159, 8, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 168, 8, 18, 1, 18, 0, 3, 8, 26, 32,
		19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 0, 1, 1, 0, 16, 17, 168, 0, 38, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0, 4, 43,
		1, 0, 0, 0, 6, 47, 1, 0, 0, 0, 8, 52, 1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12,
		71, 1, 0, 0, 0, 14, 82, 1, 0, 0, 0, 16, 88, 1, 0, 0, 0, 18, 96, 1, 0, 0,
		0, 20, 98, 1, 0, 0, 0, 22, 102, 1, 0, 0, 0, 24, 109, 1, 0, 0, 0, 26, 111,
		1, 0, 0, 0, 28, 129, 1, 0, 0, 0, 30, 141, 1, 0, 0, 0, 32, 143, 1, 0, 0,
		0, 34, 158, 1, 0, 0, 0, 36, 167, 1, 0, 0, 0, 38, 39, 3, 8, 4, 0, 39, 40,
		5, 0, 0, 1, 40, 1, 1, 0, 0, 0, 41, 42, 7, 0, 0, 0, 42, 3, 1, 0, 0, 0, 43,
		44, 5, 18, 0, 0, 44, 5, 1, 0, 0, 0, 45, 48, 3, 4, 2, 0, 46, 48, 3, 2, 1,
		0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48, 7, 1, 0, 0, 0, 49, 50, 6,
		4, -1, 0, 50, 53, 3, 10, 5, 0, 51, 53, 1, 0, 0, 0, 52, 49, 1, 0, 0, 0,
		52, 51, 1, 0, 0, 0, 53, 58, 1, 0, 0, 0, 54, 55, 10, 2, 0, 0, 55, 57, 3,
		10, 5, 0, 56, 54, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58,
		59, 1, 0, 0, 0, 59, 9, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62, 3, 14, 7,
		0, 62, 63, 3, 16, 8, 0, 63, 64, 3, 12, 6, 0, 64, 11, 1, 0, 0, 0, 65, 72,
		3, 18, 9, 0, 66, 72, 3, 30, 15, 0, 67, 68, 5, 1, 0, 0, 68, 69, 3, 36, 18,
		0, 69, 70, 5, 2, 0, 0, 70, 72, 1, 0, 0, 0, 71, 65, 1, 0, 0, 0, 71, 66,
		1, 0, 0, 0, 71, 67, 1, 0, 0, 0, 72, 13, 1, 0, 0, 0, 73, 74, 5, 3, 0, 0,
		74, 75, 5, 16, 0, 0, 75, 76, 5, 4, 0, 0, 76, 83, 3, 6, 3, 0, 77, 78, 5,
		3, 0, 0, 78, 79, 5, 16, 0, 0, 79, 80, 5, 5, 0, 0, 80, 83, 3, 6, 3, 0, 81,
		83, 1, 0, 0, 0, 82, 73, 1, 0, 0, 0, 82, 77, 1, 0, 0, 0, 82, 81, 1, 0, 0,
		0, 83, 15, 1, 0, 0, 0, 84, 85, 3, 32, 16, 0, 85, 86, 5, 6, 0, 0, 86, 89,
		1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 84, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0,
		89, 17, 1, 0, 0, 0, 90, 91, 5, 16, 0, 0, 91, 92, 5, 7, 0, 0, 92, 97, 3,
		20, 10, 0, 93, 94, 5, 16, 0, 0, 94, 95, 5, 7, 0, 0, 95, 97, 3, 6, 3, 0,
		96, 90, 1, 0, 0, 0, 96, 93, 1, 0, 0, 0, 97, 19, 1, 0, 0, 0, 98, 99, 5,
		16, 0, 0, 99, 100, 5, 8, 0, 0, 100, 101, 3, 22, 11, 0, 101, 21, 1, 0, 0,
		0, 102, 103, 5, 16, 0, 0, 103, 104, 5, 9, 0, 0, 104, 105, 3, 24, 12, 0,
		105, 106, 5, 10, 0, 0, 106, 23, 1, 0, 0, 0, 107, 110, 3, 26, 13, 0, 108,
		110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 25, 1,
		0, 0, 0, 111, 112, 6, 13, -1, 0, 112, 113, 3, 28, 14, 0, 113, 119, 1, 0,
		0, 0, 114, 115, 10, 1, 0, 0, 115, 116, 5, 11, 0, 0, 116, 118, 3, 28, 14,
		0, 117, 114, 1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119,
		120, 1, 0, 0, 0, 120, 27, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5,
		16, 0, 0, 123, 124, 5, 7, 0, 0, 124, 130, 3, 6, 3, 0, 125, 126, 5, 16,
		0, 0, 126, 127, 5, 8, 0, 0, 127, 130, 3, 6, 3, 0, 128, 130, 3, 6, 3, 0,
		129, 122, 1, 0, 0, 0, 129, 125, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130,
		29, 1, 0, 0, 0, 131, 132, 3, 32, 16, 0, 132, 133, 5, 12, 0, 0, 133, 134,
		3, 2, 1, 0, 134, 142, 1, 0, 0, 0, 135, 136, 5, 1, 0, 0, 136, 137, 3, 34,
		17, 0, 137, 138, 5, 2, 0, 0, 138, 139, 5, 12, 0, 0, 139, 140, 3, 2, 1,
		0, 140, 142, 1, 0, 0, 0, 141, 131, 1, 0, 0, 0, 141, 135, 1, 0, 0, 0, 142,
		31, 1, 0, 0, 0, 143, 144, 6, 16, -1, 0, 144, 145, 3, 22, 11, 0, 145, 151,
		1, 0, 0, 0, 146, 147, 10, 1, 0, 0, 147, 148, 5, 6, 0, 0, 148, 150, 3, 32,
		16, 2, 149, 146, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 33, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 159,
		3, 32, 16, 0, 155, 156, 3, 32, 16, 0, 156, 157, 3, 34, 17, 0, 157, 159,
		1, 0, 0, 0, 158, 154, 1, 0, 0, 0, 158, 155, 1, 0, 0, 0, 159, 35, 1, 0,
		0, 0, 160, 161, 3, 30, 15, 0, 161, 162, 3, 36, 18, 0, 162, 168, 1, 0, 0,
		0, 163, 164, 3, 18, 9, 0, 164, 165, 3, 36, 18, 0, 165, 168, 1, 0, 0, 0,
		166, 168, 1, 0, 0, 0, 167, 160, 1, 0, 0, 0, 167, 163, 1, 0, 0, 0, 167,
		166, 1, 0, 0, 0, 168, 37, 1, 0, 0, 0, 14, 47, 52, 58, 71, 82, 88, 96, 109,
		119, 129, 141, 151, 158, 167,
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

// routingAParserInit initializes any static state used to implement routingAParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewroutingAParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RoutingAParserInit() {
	staticData := &routingaParserStaticData
	staticData.once.Do(routingaParserInit)
}

// NewroutingAParser produces a new parser instance for the optional input antlr.TokenStream.
func NewroutingAParser(input antlr.TokenStream) *routingAParser {
	RoutingAParserInit()
	this := new(routingAParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &routingaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "routingA.g4"

	return this
}

// routingAParser tokens.
const (
	routingAParserEOF                = antlr.TokenEOF
	routingAParserT__0               = 1
	routingAParserT__1               = 2
	routingAParserT__2               = 3
	routingAParserT__3               = 4
	routingAParserT__4               = 5
	routingAParserT__5               = 6
	routingAParserT__6               = 7
	routingAParserT__7               = 8
	routingAParserT__8               = 9
	routingAParserT__9               = 10
	routingAParserT__10              = 11
	routingAParserT__11              = 12
	routingAParserWHITESPACE         = 13
	routingAParserCOMMENT_BLOCK      = 14
	routingAParserCOMMENT_LINE_SHARP = 15
	routingAParserID                 = 16
	routingAParserNON_ID             = 17
	routingAParserQUOTE_STRING       = 18
)

// routingAParser rules.
const (
	routingAParserRULE_start                             = 0
	routingAParserRULE_bare_literal                      = 1
	routingAParserRULE_quote_literal                     = 2
	routingAParserRULE_literal                           = 3
	routingAParserRULE_input                             = 4
	routingAParserRULE_programStructureBlcok             = 5
	routingAParserRULE_expression                        = 6
	routingAParserRULE_optConditionStatement             = 7
	routingAParserRULE_optFunctionPrototypeExpressionAnd = 8
	routingAParserRULE_declaration                       = 9
	routingAParserRULE_assignmentExpression              = 10
	routingAParserRULE_functionPrototype                 = 11
	routingAParserRULE_parameterList                     = 12
	routingAParserRULE_nonEmptyParameterList             = 13
	routingAParserRULE_parameter                         = 14
	routingAParserRULE_routingRule                       = 15
	routingAParserRULE_functionPrototypeExpression       = 16
	routingAParserRULE_functionPrototypeExpressionList   = 17
	routingAParserRULE_routingRuleOrDeclarationList      = 18
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Input() IInputContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(routingAParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *routingAParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, routingAParserRULE_start)

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
		p.SetState(38)
		p.input(0)
	}
	{
		p.SetState(39)
		p.Match(routingAParserEOF)
	}

	return localctx
}

// IBare_literalContext is an interface to support dynamic dispatch.
type IBare_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBare_literalContext differentiates from other interfaces.
	IsBare_literalContext()
}

type Bare_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBare_literalContext() *Bare_literalContext {
	var p = new(Bare_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_bare_literal
	return p
}

func (*Bare_literalContext) IsBare_literalContext() {}

func NewBare_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bare_literalContext {
	var p = new(Bare_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_bare_literal

	return p
}

func (s *Bare_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Bare_literalContext) ID() antlr.TerminalNode {
	return s.GetToken(routingAParserID, 0)
}

func (s *Bare_literalContext) NON_ID() antlr.TerminalNode {
	return s.GetToken(routingAParserNON_ID, 0)
}

func (s *Bare_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bare_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bare_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterBare_literal(s)
	}
}

func (s *Bare_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitBare_literal(s)
	}
}

func (p *routingAParser) Bare_literal() (localctx IBare_literalContext) {
	this := p
	_ = this

	localctx = NewBare_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, routingAParserRULE_bare_literal)
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
		p.SetState(41)
		_la = p.GetTokenStream().LA(1)

		if !(_la == routingAParserID || _la == routingAParserNON_ID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IQuote_literalContext is an interface to support dynamic dispatch.
type IQuote_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuote_literalContext differentiates from other interfaces.
	IsQuote_literalContext()
}

type Quote_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuote_literalContext() *Quote_literalContext {
	var p = new(Quote_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_quote_literal
	return p
}

func (*Quote_literalContext) IsQuote_literalContext() {}

func NewQuote_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quote_literalContext {
	var p = new(Quote_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_quote_literal

	return p
}

func (s *Quote_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Quote_literalContext) QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(routingAParserQUOTE_STRING, 0)
}

func (s *Quote_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quote_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Quote_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterQuote_literal(s)
	}
}

func (s *Quote_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitQuote_literal(s)
	}
}

func (p *routingAParser) Quote_literal() (localctx IQuote_literalContext) {
	this := p
	_ = this

	localctx = NewQuote_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, routingAParserRULE_quote_literal)

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
		p.SetState(43)
		p.Match(routingAParserQUOTE_STRING)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) Quote_literal() IQuote_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuote_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuote_literalContext)
}

func (s *LiteralContext) Bare_literal() IBare_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBare_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBare_literalContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *routingAParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, routingAParserRULE_literal)

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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case routingAParserQUOTE_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Quote_literal()
		}

	case routingAParserID, routingAParserNON_ID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Bare_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_input
	return p
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) ProgramStructureBlcok() IProgramStructureBlcokContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramStructureBlcokContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramStructureBlcokContext)
}

func (s *InputContext) Input() IInputContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInputContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInputContext)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterInput(s)
	}
}

func (s *InputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitInput(s)
	}
}

func (p *routingAParser) Input() (localctx IInputContext) {
	return p.input(0)
}

func (p *routingAParser) input(_p int) (localctx IInputContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewInputContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInputContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, routingAParserRULE_input, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(50)
			p.ProgramStructureBlcok()
		}

	case 2:

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewInputContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, routingAParserRULE_input)
			p.SetState(54)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(55)
				p.ProgramStructureBlcok()
			}

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IProgramStructureBlcokContext is an interface to support dynamic dispatch.
type IProgramStructureBlcokContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramStructureBlcokContext differentiates from other interfaces.
	IsProgramStructureBlcokContext()
}

type ProgramStructureBlcokContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramStructureBlcokContext() *ProgramStructureBlcokContext {
	var p = new(ProgramStructureBlcokContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_programStructureBlcok
	return p
}

func (*ProgramStructureBlcokContext) IsProgramStructureBlcokContext() {}

func NewProgramStructureBlcokContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramStructureBlcokContext {
	var p = new(ProgramStructureBlcokContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_programStructureBlcok

	return p
}

func (s *ProgramStructureBlcokContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramStructureBlcokContext) OptConditionStatement() IOptConditionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptConditionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptConditionStatementContext)
}

func (s *ProgramStructureBlcokContext) OptFunctionPrototypeExpressionAnd() IOptFunctionPrototypeExpressionAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptFunctionPrototypeExpressionAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptFunctionPrototypeExpressionAndContext)
}

func (s *ProgramStructureBlcokContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramStructureBlcokContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramStructureBlcokContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramStructureBlcokContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterProgramStructureBlcok(s)
	}
}

func (s *ProgramStructureBlcokContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitProgramStructureBlcok(s)
	}
}

func (p *routingAParser) ProgramStructureBlcok() (localctx IProgramStructureBlcokContext) {
	this := p
	_ = this

	localctx = NewProgramStructureBlcokContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, routingAParserRULE_programStructureBlcok)

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
		p.SetState(61)
		p.OptConditionStatement()
	}
	{
		p.SetState(62)
		p.OptFunctionPrototypeExpressionAnd()
	}
	{
		p.SetState(63)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ExpressionContext) RoutingRule() IRoutingRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleContext)
}

func (s *ExpressionContext) RoutingRuleOrDeclarationList() IRoutingRuleOrDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleOrDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleOrDeclarationListContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *routingAParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, routingAParserRULE_expression)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.RoutingRule()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(routingAParserT__0)
		}
		{
			p.SetState(68)
			p.RoutingRuleOrDeclarationList()
		}
		{
			p.SetState(69)
			p.Match(routingAParserT__1)
		}

	}

	return localctx
}

// IOptConditionStatementContext is an interface to support dynamic dispatch.
type IOptConditionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptConditionStatementContext differentiates from other interfaces.
	IsOptConditionStatementContext()
}

type OptConditionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptConditionStatementContext() *OptConditionStatementContext {
	var p = new(OptConditionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_optConditionStatement
	return p
}

func (*OptConditionStatementContext) IsOptConditionStatementContext() {}

func NewOptConditionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptConditionStatementContext {
	var p = new(OptConditionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_optConditionStatement

	return p
}

func (s *OptConditionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptConditionStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(routingAParserID, 0)
}

func (s *OptConditionStatementContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OptConditionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptConditionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptConditionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterOptConditionStatement(s)
	}
}

func (s *OptConditionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitOptConditionStatement(s)
	}
}

func (p *routingAParser) OptConditionStatement() (localctx IOptConditionStatementContext) {
	this := p
	_ = this

	localctx = NewOptConditionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, routingAParserRULE_optConditionStatement)

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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(routingAParserT__2)
		}
		{
			p.SetState(74)
			p.Match(routingAParserID)
		}
		{
			p.SetState(75)
			p.Match(routingAParserT__3)
		}
		{
			p.SetState(76)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(routingAParserT__2)
		}
		{
			p.SetState(78)
			p.Match(routingAParserID)
		}
		{
			p.SetState(79)
			p.Match(routingAParserT__4)
		}
		{
			p.SetState(80)
			p.Literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	}

	return localctx
}

// IOptFunctionPrototypeExpressionAndContext is an interface to support dynamic dispatch.
type IOptFunctionPrototypeExpressionAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptFunctionPrototypeExpressionAndContext differentiates from other interfaces.
	IsOptFunctionPrototypeExpressionAndContext()
}

type OptFunctionPrototypeExpressionAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptFunctionPrototypeExpressionAndContext() *OptFunctionPrototypeExpressionAndContext {
	var p = new(OptFunctionPrototypeExpressionAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_optFunctionPrototypeExpressionAnd
	return p
}

func (*OptFunctionPrototypeExpressionAndContext) IsOptFunctionPrototypeExpressionAndContext() {}

func NewOptFunctionPrototypeExpressionAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptFunctionPrototypeExpressionAndContext {
	var p = new(OptFunctionPrototypeExpressionAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_optFunctionPrototypeExpressionAnd

	return p
}

func (s *OptFunctionPrototypeExpressionAndContext) GetParser() antlr.Parser { return s.parser }

func (s *OptFunctionPrototypeExpressionAndContext) FunctionPrototypeExpression() IFunctionPrototypeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionPrototypeExpressionContext)
}

func (s *OptFunctionPrototypeExpressionAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptFunctionPrototypeExpressionAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptFunctionPrototypeExpressionAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterOptFunctionPrototypeExpressionAnd(s)
	}
}

func (s *OptFunctionPrototypeExpressionAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitOptFunctionPrototypeExpressionAnd(s)
	}
}

func (p *routingAParser) OptFunctionPrototypeExpressionAnd() (localctx IOptFunctionPrototypeExpressionAndContext) {
	this := p
	_ = this

	localctx = NewOptFunctionPrototypeExpressionAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, routingAParserRULE_optFunctionPrototypeExpressionAnd)

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

	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.functionPrototypeExpression(0)
		}
		{
			p.SetState(85)
			p.Match(routingAParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(routingAParserID, 0)
}

func (s *DeclarationContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *DeclarationContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *routingAParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, routingAParserRULE_declaration)

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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(routingAParserID)
		}
		{
			p.SetState(91)
			p.Match(routingAParserT__6)
		}
		{
			p.SetState(92)
			p.AssignmentExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(routingAParserID)
		}
		{
			p.SetState(94)
			p.Match(routingAParserT__6)
		}
		{
			p.SetState(95)
			p.Literal()
		}

	}

	return localctx
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_assignmentExpression
	return p
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(routingAParserID, 0)
}

func (s *AssignmentExpressionContext) FunctionPrototype() IFunctionPrototypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionPrototypeContext)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (p *routingAParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	this := p
	_ = this

	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, routingAParserRULE_assignmentExpression)

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
		p.SetState(98)
		p.Match(routingAParserID)
	}
	{
		p.SetState(99)
		p.Match(routingAParserT__7)
	}
	{
		p.SetState(100)
		p.FunctionPrototype()
	}

	return localctx
}

// IFunctionPrototypeContext is an interface to support dynamic dispatch.
type IFunctionPrototypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionPrototypeContext differentiates from other interfaces.
	IsFunctionPrototypeContext()
}

type FunctionPrototypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionPrototypeContext() *FunctionPrototypeContext {
	var p = new(FunctionPrototypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_functionPrototype
	return p
}

func (*FunctionPrototypeContext) IsFunctionPrototypeContext() {}

func NewFunctionPrototypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionPrototypeContext {
	var p = new(FunctionPrototypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_functionPrototype

	return p
}

func (s *FunctionPrototypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionPrototypeContext) ID() antlr.TerminalNode {
	return s.GetToken(routingAParserID, 0)
}

func (s *FunctionPrototypeContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionPrototypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionPrototypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionPrototypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterFunctionPrototype(s)
	}
}

func (s *FunctionPrototypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitFunctionPrototype(s)
	}
}

func (p *routingAParser) FunctionPrototype() (localctx IFunctionPrototypeContext) {
	this := p
	_ = this

	localctx = NewFunctionPrototypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, routingAParserRULE_functionPrototype)

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
		p.SetState(102)
		p.Match(routingAParserID)
	}
	{
		p.SetState(103)
		p.Match(routingAParserT__8)
	}
	{
		p.SetState(104)
		p.ParameterList()
	}
	{
		p.SetState(105)
		p.Match(routingAParserT__9)
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) NonEmptyParameterList() INonEmptyParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonEmptyParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonEmptyParameterListContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *routingAParser) ParameterList() (localctx IParameterListContext) {
	this := p
	_ = this

	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, routingAParserRULE_parameterList)

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

	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case routingAParserID, routingAParserNON_ID, routingAParserQUOTE_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.nonEmptyParameterList(0)
		}

	case routingAParserT__9:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INonEmptyParameterListContext is an interface to support dynamic dispatch.
type INonEmptyParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonEmptyParameterListContext differentiates from other interfaces.
	IsNonEmptyParameterListContext()
}

type NonEmptyParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonEmptyParameterListContext() *NonEmptyParameterListContext {
	var p = new(NonEmptyParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_nonEmptyParameterList
	return p
}

func (*NonEmptyParameterListContext) IsNonEmptyParameterListContext() {}

func NewNonEmptyParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonEmptyParameterListContext {
	var p = new(NonEmptyParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_nonEmptyParameterList

	return p
}

func (s *NonEmptyParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *NonEmptyParameterListContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *NonEmptyParameterListContext) NonEmptyParameterList() INonEmptyParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonEmptyParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonEmptyParameterListContext)
}

func (s *NonEmptyParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonEmptyParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonEmptyParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterNonEmptyParameterList(s)
	}
}

func (s *NonEmptyParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitNonEmptyParameterList(s)
	}
}

func (p *routingAParser) NonEmptyParameterList() (localctx INonEmptyParameterListContext) {
	return p.nonEmptyParameterList(0)
}

func (p *routingAParser) nonEmptyParameterList(_p int) (localctx INonEmptyParameterListContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewNonEmptyParameterListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INonEmptyParameterListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, routingAParserRULE_nonEmptyParameterList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Parameter()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewNonEmptyParameterListContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, routingAParserRULE_nonEmptyParameterList)
			p.SetState(114)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(115)
				p.Match(routingAParserT__10)
			}
			{
				p.SetState(116)
				p.Parameter()
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) ID() antlr.TerminalNode {
	return s.GetToken(routingAParserID, 0)
}

func (s *ParameterContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *routingAParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, routingAParserRULE_parameter)

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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(routingAParserID)
		}
		{
			p.SetState(123)
			p.Match(routingAParserT__6)
		}
		{
			p.SetState(124)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(routingAParserID)
		}
		{
			p.SetState(126)
			p.Match(routingAParserT__7)
		}
		{
			p.SetState(127)
			p.Literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Literal()
		}

	}

	return localctx
}

// IRoutingRuleContext is an interface to support dynamic dispatch.
type IRoutingRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoutingRuleContext differentiates from other interfaces.
	IsRoutingRuleContext()
}

type RoutingRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoutingRuleContext() *RoutingRuleContext {
	var p = new(RoutingRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_routingRule
	return p
}

func (*RoutingRuleContext) IsRoutingRuleContext() {}

func NewRoutingRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoutingRuleContext {
	var p = new(RoutingRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_routingRule

	return p
}

func (s *RoutingRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RoutingRuleContext) FunctionPrototypeExpression() IFunctionPrototypeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionPrototypeExpressionContext)
}

func (s *RoutingRuleContext) Bare_literal() IBare_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBare_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBare_literalContext)
}

func (s *RoutingRuleContext) FunctionPrototypeExpressionList() IFunctionPrototypeExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionPrototypeExpressionListContext)
}

func (s *RoutingRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutingRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoutingRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterRoutingRule(s)
	}
}

func (s *RoutingRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitRoutingRule(s)
	}
}

func (p *routingAParser) RoutingRule() (localctx IRoutingRuleContext) {
	this := p
	_ = this

	localctx = NewRoutingRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, routingAParserRULE_routingRule)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case routingAParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.functionPrototypeExpression(0)
		}
		{
			p.SetState(132)
			p.Match(routingAParserT__11)
		}
		{
			p.SetState(133)
			p.Bare_literal()
		}

	case routingAParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(routingAParserT__0)
		}
		{
			p.SetState(136)
			p.FunctionPrototypeExpressionList()
		}
		{
			p.SetState(137)
			p.Match(routingAParserT__1)
		}
		{
			p.SetState(138)
			p.Match(routingAParserT__11)
		}
		{
			p.SetState(139)
			p.Bare_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionPrototypeExpressionContext is an interface to support dynamic dispatch.
type IFunctionPrototypeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionPrototypeExpressionContext differentiates from other interfaces.
	IsFunctionPrototypeExpressionContext()
}

type FunctionPrototypeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionPrototypeExpressionContext() *FunctionPrototypeExpressionContext {
	var p = new(FunctionPrototypeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_functionPrototypeExpression
	return p
}

func (*FunctionPrototypeExpressionContext) IsFunctionPrototypeExpressionContext() {}

func NewFunctionPrototypeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionPrototypeExpressionContext {
	var p = new(FunctionPrototypeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_functionPrototypeExpression

	return p
}

func (s *FunctionPrototypeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionPrototypeExpressionContext) FunctionPrototype() IFunctionPrototypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionPrototypeContext)
}

func (s *FunctionPrototypeExpressionContext) AllFunctionPrototypeExpression() []IFunctionPrototypeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionPrototypeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionPrototypeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionPrototypeExpressionContext); ok {
			tst[i] = t.(IFunctionPrototypeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionPrototypeExpressionContext) FunctionPrototypeExpression(i int) IFunctionPrototypeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeExpressionContext); ok {
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

	return t.(IFunctionPrototypeExpressionContext)
}

func (s *FunctionPrototypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionPrototypeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionPrototypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterFunctionPrototypeExpression(s)
	}
}

func (s *FunctionPrototypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitFunctionPrototypeExpression(s)
	}
}

func (p *routingAParser) FunctionPrototypeExpression() (localctx IFunctionPrototypeExpressionContext) {
	return p.functionPrototypeExpression(0)
}

func (p *routingAParser) functionPrototypeExpression(_p int) (localctx IFunctionPrototypeExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFunctionPrototypeExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFunctionPrototypeExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, routingAParserRULE_functionPrototypeExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.FunctionPrototype()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFunctionPrototypeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, routingAParserRULE_functionPrototypeExpression)
			p.SetState(146)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(147)
				p.Match(routingAParserT__5)
			}
			{
				p.SetState(148)
				p.functionPrototypeExpression(2)
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionPrototypeExpressionListContext is an interface to support dynamic dispatch.
type IFunctionPrototypeExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionPrototypeExpressionListContext differentiates from other interfaces.
	IsFunctionPrototypeExpressionListContext()
}

type FunctionPrototypeExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionPrototypeExpressionListContext() *FunctionPrototypeExpressionListContext {
	var p = new(FunctionPrototypeExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_functionPrototypeExpressionList
	return p
}

func (*FunctionPrototypeExpressionListContext) IsFunctionPrototypeExpressionListContext() {}

func NewFunctionPrototypeExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionPrototypeExpressionListContext {
	var p = new(FunctionPrototypeExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_functionPrototypeExpressionList

	return p
}

func (s *FunctionPrototypeExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionPrototypeExpressionListContext) FunctionPrototypeExpression() IFunctionPrototypeExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionPrototypeExpressionContext)
}

func (s *FunctionPrototypeExpressionListContext) FunctionPrototypeExpressionList() IFunctionPrototypeExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionPrototypeExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionPrototypeExpressionListContext)
}

func (s *FunctionPrototypeExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionPrototypeExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionPrototypeExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterFunctionPrototypeExpressionList(s)
	}
}

func (s *FunctionPrototypeExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitFunctionPrototypeExpressionList(s)
	}
}

func (p *routingAParser) FunctionPrototypeExpressionList() (localctx IFunctionPrototypeExpressionListContext) {
	this := p
	_ = this

	localctx = NewFunctionPrototypeExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, routingAParserRULE_functionPrototypeExpressionList)

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

	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.functionPrototypeExpression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.functionPrototypeExpression(0)
		}
		{
			p.SetState(156)
			p.FunctionPrototypeExpressionList()
		}

	}

	return localctx
}

// IRoutingRuleOrDeclarationListContext is an interface to support dynamic dispatch.
type IRoutingRuleOrDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoutingRuleOrDeclarationListContext differentiates from other interfaces.
	IsRoutingRuleOrDeclarationListContext()
}

type RoutingRuleOrDeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoutingRuleOrDeclarationListContext() *RoutingRuleOrDeclarationListContext {
	var p = new(RoutingRuleOrDeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_routingRuleOrDeclarationList
	return p
}

func (*RoutingRuleOrDeclarationListContext) IsRoutingRuleOrDeclarationListContext() {}

func NewRoutingRuleOrDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoutingRuleOrDeclarationListContext {
	var p = new(RoutingRuleOrDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_routingRuleOrDeclarationList

	return p
}

func (s *RoutingRuleOrDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *RoutingRuleOrDeclarationListContext) RoutingRule() IRoutingRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleContext)
}

func (s *RoutingRuleOrDeclarationListContext) RoutingRuleOrDeclarationList() IRoutingRuleOrDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleOrDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleOrDeclarationListContext)
}

func (s *RoutingRuleOrDeclarationListContext) Declaration() IDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *RoutingRuleOrDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutingRuleOrDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoutingRuleOrDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterRoutingRuleOrDeclarationList(s)
	}
}

func (s *RoutingRuleOrDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitRoutingRuleOrDeclarationList(s)
	}
}

func (p *routingAParser) RoutingRuleOrDeclarationList() (localctx IRoutingRuleOrDeclarationListContext) {
	this := p
	_ = this

	localctx = NewRoutingRuleOrDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, routingAParserRULE_routingRuleOrDeclarationList)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.RoutingRule()
		}
		{
			p.SetState(161)
			p.RoutingRuleOrDeclarationList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Declaration()
		}
		{
			p.SetState(164)
			p.RoutingRuleOrDeclarationList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	}

	return localctx
}

func (p *routingAParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *InputContext = nil
		if localctx != nil {
			t = localctx.(*InputContext)
		}
		return p.Input_Sempred(t, predIndex)

	case 13:
		var t *NonEmptyParameterListContext = nil
		if localctx != nil {
			t = localctx.(*NonEmptyParameterListContext)
		}
		return p.NonEmptyParameterList_Sempred(t, predIndex)

	case 16:
		var t *FunctionPrototypeExpressionContext = nil
		if localctx != nil {
			t = localctx.(*FunctionPrototypeExpressionContext)
		}
		return p.FunctionPrototypeExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *routingAParser) Input_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *routingAParser) NonEmptyParameterList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *routingAParser) FunctionPrototypeExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
