// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package routingA // routingA
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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
		"", "'{'", "'}'", "'@'", "'=='", "'!='", "':'", "'='", "'!'", "'('",
		"')'", "','", "'->'", "'&&'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "WHITESPACE",
		"COMMENT_BLOCK", "COMMENT_LINE_SHARP", "ID", "NON_ID", "QUOTE_STRING",
	}
	staticData.ruleNames = []string{
		"start", "bare_literal", "quote_literal", "literal", "input", "programStructureBlcok",
		"expression", "optConditionStatement", "declaration", "assignmentExpression",
		"functionPrototype", "optParameterList", "nonEmptyParameterList", "parameter",
		"routingRule", "routingRuleLeft", "routingRuleLeftList", "optFunctionPrototypeExpressionAnd",
		"functionPrototypeExpression", "routingRuleOrDeclarationList", "routingRuleList",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 192, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 52, 8, 3, 1, 4,
		1, 4, 1, 4, 3, 4, 57, 8, 4, 1, 4, 1, 4, 5, 4, 61, 8, 4, 10, 4, 12, 4, 64,
		9, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 75, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 86, 8, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 94, 8, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 10, 3, 10, 101, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 3, 11, 110, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5,
		12, 118, 8, 12, 10, 12, 12, 12, 121, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 3, 13, 130, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 141, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 151, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 157, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 163, 8,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 171, 8, 18, 10, 18,
		12, 18, 174, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 184, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 190, 8, 20, 1,
		20, 0, 3, 8, 24, 36, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 0, 1, 1, 0, 17, 18, 192, 0, 42, 1, 0, 0,
		0, 2, 45, 1, 0, 0, 0, 4, 47, 1, 0, 0, 0, 6, 51, 1, 0, 0, 0, 8, 56, 1, 0,
		0, 0, 10, 65, 1, 0, 0, 0, 12, 74, 1, 0, 0, 0, 14, 85, 1, 0, 0, 0, 16, 93,
		1, 0, 0, 0, 18, 95, 1, 0, 0, 0, 20, 100, 1, 0, 0, 0, 22, 109, 1, 0, 0,
		0, 24, 111, 1, 0, 0, 0, 26, 129, 1, 0, 0, 0, 28, 140, 1, 0, 0, 0, 30, 150,
		1, 0, 0, 0, 32, 156, 1, 0, 0, 0, 34, 162, 1, 0, 0, 0, 36, 164, 1, 0, 0,
		0, 38, 183, 1, 0, 0, 0, 40, 189, 1, 0, 0, 0, 42, 43, 3, 8, 4, 0, 43, 44,
		5, 0, 0, 1, 44, 1, 1, 0, 0, 0, 45, 46, 7, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47,
		48, 5, 19, 0, 0, 48, 5, 1, 0, 0, 0, 49, 52, 3, 4, 2, 0, 50, 52, 3, 2, 1,
		0, 51, 49, 1, 0, 0, 0, 51, 50, 1, 0, 0, 0, 52, 7, 1, 0, 0, 0, 53, 54, 6,
		4, -1, 0, 54, 57, 3, 10, 5, 0, 55, 57, 1, 0, 0, 0, 56, 53, 1, 0, 0, 0,
		56, 55, 1, 0, 0, 0, 57, 62, 1, 0, 0, 0, 58, 59, 10, 2, 0, 0, 59, 61, 3,
		10, 5, 0, 60, 58, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62,
		63, 1, 0, 0, 0, 63, 9, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 3, 14, 7,
		0, 66, 67, 3, 12, 6, 0, 67, 11, 1, 0, 0, 0, 68, 75, 3, 16, 8, 0, 69, 75,
		3, 28, 14, 0, 70, 71, 5, 1, 0, 0, 71, 72, 3, 38, 19, 0, 72, 73, 5, 2, 0,
		0, 73, 75, 1, 0, 0, 0, 74, 68, 1, 0, 0, 0, 74, 69, 1, 0, 0, 0, 74, 70,
		1, 0, 0, 0, 75, 13, 1, 0, 0, 0, 76, 77, 5, 3, 0, 0, 77, 78, 5, 17, 0, 0,
		78, 79, 5, 4, 0, 0, 79, 86, 3, 6, 3, 0, 80, 81, 5, 3, 0, 0, 81, 82, 5,
		17, 0, 0, 82, 83, 5, 5, 0, 0, 83, 86, 3, 6, 3, 0, 84, 86, 1, 0, 0, 0, 85,
		76, 1, 0, 0, 0, 85, 80, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 15, 1, 0, 0,
		0, 87, 88, 5, 17, 0, 0, 88, 89, 5, 6, 0, 0, 89, 94, 3, 18, 9, 0, 90, 91,
		5, 17, 0, 0, 91, 92, 5, 6, 0, 0, 92, 94, 3, 6, 3, 0, 93, 87, 1, 0, 0, 0,
		93, 90, 1, 0, 0, 0, 94, 17, 1, 0, 0, 0, 95, 96, 5, 17, 0, 0, 96, 97, 5,
		7, 0, 0, 97, 98, 3, 20, 10, 0, 98, 19, 1, 0, 0, 0, 99, 101, 5, 8, 0, 0,
		100, 99, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103,
		5, 17, 0, 0, 103, 104, 5, 9, 0, 0, 104, 105, 3, 22, 11, 0, 105, 106, 5,
		10, 0, 0, 106, 21, 1, 0, 0, 0, 107, 110, 3, 24, 12, 0, 108, 110, 1, 0,
		0, 0, 109, 107, 1, 0, 0, 0, 109, 108, 1, 0, 0, 0, 110, 23, 1, 0, 0, 0,
		111, 112, 6, 12, -1, 0, 112, 113, 3, 26, 13, 0, 113, 119, 1, 0, 0, 0, 114,
		115, 10, 1, 0, 0, 115, 116, 5, 11, 0, 0, 116, 118, 3, 26, 13, 0, 117, 114,
		1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0,
		0, 0, 120, 25, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5, 17, 0, 0,
		123, 124, 5, 6, 0, 0, 124, 130, 3, 6, 3, 0, 125, 126, 5, 17, 0, 0, 126,
		127, 5, 7, 0, 0, 127, 130, 3, 6, 3, 0, 128, 130, 3, 6, 3, 0, 129, 122,
		1, 0, 0, 0, 129, 125, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 27, 1, 0,
		0, 0, 131, 132, 3, 30, 15, 0, 132, 133, 5, 12, 0, 0, 133, 134, 3, 2, 1,
		0, 134, 141, 1, 0, 0, 0, 135, 136, 3, 34, 17, 0, 136, 137, 5, 1, 0, 0,
		137, 138, 3, 40, 20, 0, 138, 139, 5, 2, 0, 0, 139, 141, 1, 0, 0, 0, 140,
		131, 1, 0, 0, 0, 140, 135, 1, 0, 0, 0, 141, 29, 1, 0, 0, 0, 142, 143, 3,
		34, 17, 0, 143, 144, 5, 1, 0, 0, 144, 145, 3, 32, 16, 0, 145, 146, 5, 2,
		0, 0, 146, 151, 1, 0, 0, 0, 147, 148, 3, 34, 17, 0, 148, 149, 3, 36, 18,
		0, 149, 151, 1, 0, 0, 0, 150, 142, 1, 0, 0, 0, 150, 147, 1, 0, 0, 0, 151,
		31, 1, 0, 0, 0, 152, 157, 3, 30, 15, 0, 153, 154, 3, 30, 15, 0, 154, 155,
		3, 32, 16, 0, 155, 157, 1, 0, 0, 0, 156, 152, 1, 0, 0, 0, 156, 153, 1,
		0, 0, 0, 157, 33, 1, 0, 0, 0, 158, 159, 3, 36, 18, 0, 159, 160, 5, 13,
		0, 0, 160, 163, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162, 158, 1, 0, 0, 0,
		162, 161, 1, 0, 0, 0, 163, 35, 1, 0, 0, 0, 164, 165, 6, 18, -1, 0, 165,
		166, 3, 20, 10, 0, 166, 172, 1, 0, 0, 0, 167, 168, 10, 1, 0, 0, 168, 169,
		5, 13, 0, 0, 169, 171, 3, 36, 18, 2, 170, 167, 1, 0, 0, 0, 171, 174, 1,
		0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 37, 1, 0, 0,
		0, 174, 172, 1, 0, 0, 0, 175, 184, 3, 28, 14, 0, 176, 184, 3, 16, 8, 0,
		177, 178, 3, 28, 14, 0, 178, 179, 3, 38, 19, 0, 179, 184, 1, 0, 0, 0, 180,
		181, 3, 16, 8, 0, 181, 182, 3, 38, 19, 0, 182, 184, 1, 0, 0, 0, 183, 175,
		1, 0, 0, 0, 183, 176, 1, 0, 0, 0, 183, 177, 1, 0, 0, 0, 183, 180, 1, 0,
		0, 0, 184, 39, 1, 0, 0, 0, 185, 190, 3, 28, 14, 0, 186, 187, 3, 28, 14,
		0, 187, 188, 3, 40, 20, 0, 188, 190, 1, 0, 0, 0, 189, 185, 1, 0, 0, 0,
		189, 186, 1, 0, 0, 0, 190, 41, 1, 0, 0, 0, 17, 51, 56, 62, 74, 85, 93,
		100, 109, 119, 129, 140, 150, 156, 162, 172, 183, 189,
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
	this.GrammarFileName = "java-escape"

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
	routingAParserT__12              = 13
	routingAParserWHITESPACE         = 14
	routingAParserCOMMENT_BLOCK      = 15
	routingAParserCOMMENT_LINE_SHARP = 16
	routingAParserID                 = 17
	routingAParserNON_ID             = 18
	routingAParserQUOTE_STRING       = 19
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
	routingAParserRULE_declaration                       = 8
	routingAParserRULE_assignmentExpression              = 9
	routingAParserRULE_functionPrototype                 = 10
	routingAParserRULE_optParameterList                  = 11
	routingAParserRULE_nonEmptyParameterList             = 12
	routingAParserRULE_parameter                         = 13
	routingAParserRULE_routingRule                       = 14
	routingAParserRULE_routingRuleLeft                   = 15
	routingAParserRULE_routingRuleLeftList               = 16
	routingAParserRULE_optFunctionPrototypeExpressionAnd = 17
	routingAParserRULE_functionPrototypeExpression       = 18
	routingAParserRULE_routingRuleOrDeclarationList      = 19
	routingAParserRULE_routingRuleList                   = 20
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
		p.SetState(42)
		p.input(0)
	}
	{
		p.SetState(43)
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
		p.SetState(45)
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
		p.SetState(47)
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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case routingAParserQUOTE_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Quote_literal()
		}

	case routingAParserID, routingAParserNON_ID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(54)
			p.ProgramStructureBlcok()
		}

	case 2:

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(62)
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
			p.SetState(58)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(59)
				p.ProgramStructureBlcok()
			}

		}
		p.SetState(64)
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
		p.SetState(65)
		p.OptConditionStatement()
	}
	{
		p.SetState(66)
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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.RoutingRule()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Match(routingAParserT__0)
		}
		{
			p.SetState(71)
			p.RoutingRuleOrDeclarationList()
		}
		{
			p.SetState(72)
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Match(routingAParserT__2)
		}
		{
			p.SetState(77)
			p.Match(routingAParserID)
		}
		{
			p.SetState(78)
			p.Match(routingAParserT__3)
		}
		{
			p.SetState(79)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(routingAParserT__2)
		}
		{
			p.SetState(81)
			p.Match(routingAParserID)
		}
		{
			p.SetState(82)
			p.Match(routingAParserT__4)
		}
		{
			p.SetState(83)
			p.Literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

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
	p.EnterRule(localctx, 16, routingAParserRULE_declaration)

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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(routingAParserID)
		}
		{
			p.SetState(88)
			p.Match(routingAParserT__5)
		}
		{
			p.SetState(89)
			p.AssignmentExpression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(routingAParserID)
		}
		{
			p.SetState(91)
			p.Match(routingAParserT__5)
		}
		{
			p.SetState(92)
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
	p.EnterRule(localctx, 18, routingAParserRULE_assignmentExpression)

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
		p.SetState(95)
		p.Match(routingAParserID)
	}
	{
		p.SetState(96)
		p.Match(routingAParserT__6)
	}
	{
		p.SetState(97)
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

func (s *FunctionPrototypeContext) OptParameterList() IOptParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptParameterListContext)
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
	p.EnterRule(localctx, 20, routingAParserRULE_functionPrototype)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == routingAParserT__7 {
		{
			p.SetState(99)
			p.Match(routingAParserT__7)
		}

	}
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
		p.OptParameterList()
	}
	{
		p.SetState(105)
		p.Match(routingAParserT__9)
	}

	return localctx
}

// IOptParameterListContext is an interface to support dynamic dispatch.
type IOptParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptParameterListContext differentiates from other interfaces.
	IsOptParameterListContext()
}

type OptParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptParameterListContext() *OptParameterListContext {
	var p = new(OptParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_optParameterList
	return p
}

func (*OptParameterListContext) IsOptParameterListContext() {}

func NewOptParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptParameterListContext {
	var p = new(OptParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_optParameterList

	return p
}

func (s *OptParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *OptParameterListContext) NonEmptyParameterList() INonEmptyParameterListContext {
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

func (s *OptParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterOptParameterList(s)
	}
}

func (s *OptParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitOptParameterList(s)
	}
}

func (p *routingAParser) OptParameterList() (localctx IOptParameterListContext) {
	this := p
	_ = this

	localctx = NewOptParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, routingAParserRULE_optParameterList)

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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, routingAParserRULE_nonEmptyParameterList, _p)

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
	p.EnterRule(localctx, 26, routingAParserRULE_parameter)

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
			p.Match(routingAParserT__5)
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
			p.Match(routingAParserT__6)
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

func (s *RoutingRuleContext) RoutingRuleLeft() IRoutingRuleLeftContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleLeftContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleLeftContext)
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

func (s *RoutingRuleContext) OptFunctionPrototypeExpressionAnd() IOptFunctionPrototypeExpressionAndContext {
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

func (s *RoutingRuleContext) RoutingRuleList() IRoutingRuleListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleListContext)
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
	p.EnterRule(localctx, 28, routingAParserRULE_routingRule)

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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.RoutingRuleLeft()
		}
		{
			p.SetState(132)
			p.Match(routingAParserT__11)
		}
		{
			p.SetState(133)
			p.Bare_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.OptFunctionPrototypeExpressionAnd()
		}
		{
			p.SetState(136)
			p.Match(routingAParserT__0)
		}
		{
			p.SetState(137)
			p.RoutingRuleList()
		}
		{
			p.SetState(138)
			p.Match(routingAParserT__1)
		}

	}

	return localctx
}

// IRoutingRuleLeftContext is an interface to support dynamic dispatch.
type IRoutingRuleLeftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoutingRuleLeftContext differentiates from other interfaces.
	IsRoutingRuleLeftContext()
}

type RoutingRuleLeftContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoutingRuleLeftContext() *RoutingRuleLeftContext {
	var p = new(RoutingRuleLeftContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_routingRuleLeft
	return p
}

func (*RoutingRuleLeftContext) IsRoutingRuleLeftContext() {}

func NewRoutingRuleLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoutingRuleLeftContext {
	var p = new(RoutingRuleLeftContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_routingRuleLeft

	return p
}

func (s *RoutingRuleLeftContext) GetParser() antlr.Parser { return s.parser }

func (s *RoutingRuleLeftContext) OptFunctionPrototypeExpressionAnd() IOptFunctionPrototypeExpressionAndContext {
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

func (s *RoutingRuleLeftContext) RoutingRuleLeftList() IRoutingRuleLeftListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleLeftListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleLeftListContext)
}

func (s *RoutingRuleLeftContext) FunctionPrototypeExpression() IFunctionPrototypeExpressionContext {
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

func (s *RoutingRuleLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutingRuleLeftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoutingRuleLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterRoutingRuleLeft(s)
	}
}

func (s *RoutingRuleLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitRoutingRuleLeft(s)
	}
}

func (p *routingAParser) RoutingRuleLeft() (localctx IRoutingRuleLeftContext) {
	this := p
	_ = this

	localctx = NewRoutingRuleLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, routingAParserRULE_routingRuleLeft)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.OptFunctionPrototypeExpressionAnd()
		}
		{
			p.SetState(143)
			p.Match(routingAParserT__0)
		}
		{
			p.SetState(144)
			p.RoutingRuleLeftList()
		}
		{
			p.SetState(145)
			p.Match(routingAParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.OptFunctionPrototypeExpressionAnd()
		}
		{
			p.SetState(148)
			p.functionPrototypeExpression(0)
		}

	}

	return localctx
}

// IRoutingRuleLeftListContext is an interface to support dynamic dispatch.
type IRoutingRuleLeftListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoutingRuleLeftListContext differentiates from other interfaces.
	IsRoutingRuleLeftListContext()
}

type RoutingRuleLeftListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoutingRuleLeftListContext() *RoutingRuleLeftListContext {
	var p = new(RoutingRuleLeftListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_routingRuleLeftList
	return p
}

func (*RoutingRuleLeftListContext) IsRoutingRuleLeftListContext() {}

func NewRoutingRuleLeftListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoutingRuleLeftListContext {
	var p = new(RoutingRuleLeftListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_routingRuleLeftList

	return p
}

func (s *RoutingRuleLeftListContext) GetParser() antlr.Parser { return s.parser }

func (s *RoutingRuleLeftListContext) RoutingRuleLeft() IRoutingRuleLeftContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleLeftContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleLeftContext)
}

func (s *RoutingRuleLeftListContext) RoutingRuleLeftList() IRoutingRuleLeftListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleLeftListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleLeftListContext)
}

func (s *RoutingRuleLeftListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutingRuleLeftListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoutingRuleLeftListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterRoutingRuleLeftList(s)
	}
}

func (s *RoutingRuleLeftListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitRoutingRuleLeftList(s)
	}
}

func (p *routingAParser) RoutingRuleLeftList() (localctx IRoutingRuleLeftListContext) {
	this := p
	_ = this

	localctx = NewRoutingRuleLeftListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, routingAParserRULE_routingRuleLeftList)

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

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.RoutingRuleLeft()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.RoutingRuleLeft()
		}
		{
			p.SetState(154)
			p.RoutingRuleLeftList()
		}

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
	p.EnterRule(localctx, 34, routingAParserRULE_optFunctionPrototypeExpressionAnd)

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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.functionPrototypeExpression(0)
		}
		{
			p.SetState(159)
			p.Match(routingAParserT__12)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)

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
	_startState := 36
	p.EnterRecursionRule(localctx, 36, routingAParserRULE_functionPrototypeExpression, _p)

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
		p.SetState(165)
		p.FunctionPrototype()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewFunctionPrototypeExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, routingAParserRULE_functionPrototypeExpression)
			p.SetState(167)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(168)
				p.Match(routingAParserT__12)
			}
			{
				p.SetState(169)
				p.functionPrototypeExpression(2)
			}

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 38, routingAParserRULE_routingRuleOrDeclarationList)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.RoutingRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.RoutingRule()
		}
		{
			p.SetState(178)
			p.RoutingRuleOrDeclarationList()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(180)
			p.Declaration()
		}
		{
			p.SetState(181)
			p.RoutingRuleOrDeclarationList()
		}

	}

	return localctx
}

// IRoutingRuleListContext is an interface to support dynamic dispatch.
type IRoutingRuleListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoutingRuleListContext differentiates from other interfaces.
	IsRoutingRuleListContext()
}

type RoutingRuleListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoutingRuleListContext() *RoutingRuleListContext {
	var p = new(RoutingRuleListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = routingAParserRULE_routingRuleList
	return p
}

func (*RoutingRuleListContext) IsRoutingRuleListContext() {}

func NewRoutingRuleListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoutingRuleListContext {
	var p = new(RoutingRuleListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = routingAParserRULE_routingRuleList

	return p
}

func (s *RoutingRuleListContext) GetParser() antlr.Parser { return s.parser }

func (s *RoutingRuleListContext) RoutingRule() IRoutingRuleContext {
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

func (s *RoutingRuleListContext) RoutingRuleList() IRoutingRuleListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoutingRuleListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoutingRuleListContext)
}

func (s *RoutingRuleListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoutingRuleListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoutingRuleListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.EnterRoutingRuleList(s)
	}
}

func (s *RoutingRuleListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(routingAListener); ok {
		listenerT.ExitRoutingRuleList(s)
	}
}

func (p *routingAParser) RoutingRuleList() (localctx IRoutingRuleListContext) {
	this := p
	_ = this

	localctx = NewRoutingRuleListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, routingAParserRULE_routingRuleList)

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

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.RoutingRule()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.RoutingRule()
		}
		{
			p.SetState(187)
			p.RoutingRuleList()
		}

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

	case 12:
		var t *NonEmptyParameterListContext = nil
		if localctx != nil {
			t = localctx.(*NonEmptyParameterListContext)
		}
		return p.NonEmptyParameterList_Sempred(t, predIndex)

	case 18:
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
