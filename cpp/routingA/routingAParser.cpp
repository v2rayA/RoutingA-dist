
// Generated from routingA.g4 by ANTLR 4.10.1


#include "routingAListener.h"

#include "routingAParser.h"


using namespace antlrcpp;
using namespace routingA;

using namespace antlr4;

namespace {

struct RoutingAParserStaticData final {
  RoutingAParserStaticData(std::vector<std::string> ruleNames,
                        std::vector<std::string> literalNames,
                        std::vector<std::string> symbolicNames)
      : ruleNames(std::move(ruleNames)), literalNames(std::move(literalNames)),
        symbolicNames(std::move(symbolicNames)),
        vocabulary(this->literalNames, this->symbolicNames) {}

  RoutingAParserStaticData(const RoutingAParserStaticData&) = delete;
  RoutingAParserStaticData(RoutingAParserStaticData&&) = delete;
  RoutingAParserStaticData& operator=(const RoutingAParserStaticData&) = delete;
  RoutingAParserStaticData& operator=(RoutingAParserStaticData&&) = delete;

  std::vector<antlr4::dfa::DFA> decisionToDFA;
  antlr4::atn::PredictionContextCache sharedContextCache;
  const std::vector<std::string> ruleNames;
  const std::vector<std::string> literalNames;
  const std::vector<std::string> symbolicNames;
  const antlr4::dfa::Vocabulary vocabulary;
  antlr4::atn::SerializedATNView serializedATN;
  std::unique_ptr<antlr4::atn::ATN> atn;
};

std::once_flag routingaParserOnceFlag;
RoutingAParserStaticData *routingaParserStaticData = nullptr;

void routingaParserInitialize() {
  assert(routingaParserStaticData == nullptr);
  auto staticData = std::make_unique<RoutingAParserStaticData>(
    std::vector<std::string>{
      "start", "bare_literal", "quote_literal", "literal", "input", "programStructureBlcok", 
      "expression", "optConditionStatement", "optFunctionPrototypeExpressionAnd", 
      "declaration", "assignmentExpression", "functionPrototype", "parameterList", 
      "nonEmptyParameterList", "parameter", "routingRule", "functionPrototypeExpression", 
      "functionPrototypeExpressionList", "routingRuleOrDeclarationList"
    },
    std::vector<std::string>{
      "", "'{'", "'}'", "'@'", "'=='", "'!='", "'&&'", "':'", "'='", "'('", 
      "')'", "','", "'->'"
    },
    std::vector<std::string>{
      "", "", "", "", "", "", "", "", "", "", "", "", "", "WHITESPACE", 
      "COMMENT_BLOCK", "COMMENT_LINE_SHARP", "ID", "NON_ID", "QUOTE_STRING"
    }
  );
  static const int32_t serializedATNSegment[] = {
  	4,1,18,170,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,
  	7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,
  	14,2,15,7,15,2,16,7,16,2,17,7,17,2,18,7,18,1,0,1,0,1,0,1,1,1,1,1,2,1,
  	2,1,3,1,3,3,3,48,8,3,1,4,1,4,1,4,3,4,53,8,4,1,4,1,4,5,4,57,8,4,10,4,12,
  	4,60,9,4,1,5,1,5,1,5,1,5,1,6,1,6,1,6,1,6,1,6,1,6,3,6,72,8,6,1,7,1,7,1,
  	7,1,7,1,7,1,7,1,7,1,7,1,7,3,7,83,8,7,1,8,1,8,1,8,1,8,3,8,89,8,8,1,9,1,
  	9,1,9,1,9,1,9,1,9,3,9,97,8,9,1,10,1,10,1,10,1,10,1,11,1,11,1,11,1,11,
  	1,11,1,12,1,12,3,12,110,8,12,1,13,1,13,1,13,1,13,1,13,1,13,5,13,118,8,
  	13,10,13,12,13,121,9,13,1,14,1,14,1,14,1,14,1,14,1,14,1,14,3,14,130,8,
  	14,1,15,1,15,1,15,1,15,1,15,1,15,1,15,1,15,1,15,1,15,3,15,142,8,15,1,
  	16,1,16,1,16,1,16,1,16,1,16,5,16,150,8,16,10,16,12,16,153,9,16,1,17,1,
  	17,1,17,1,17,3,17,159,8,17,1,18,1,18,1,18,1,18,1,18,1,18,1,18,3,18,168,
  	8,18,1,18,0,3,8,26,32,19,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,
  	34,36,0,1,1,0,16,17,168,0,38,1,0,0,0,2,41,1,0,0,0,4,43,1,0,0,0,6,47,1,
  	0,0,0,8,52,1,0,0,0,10,61,1,0,0,0,12,71,1,0,0,0,14,82,1,0,0,0,16,88,1,
  	0,0,0,18,96,1,0,0,0,20,98,1,0,0,0,22,102,1,0,0,0,24,109,1,0,0,0,26,111,
  	1,0,0,0,28,129,1,0,0,0,30,141,1,0,0,0,32,143,1,0,0,0,34,158,1,0,0,0,36,
  	167,1,0,0,0,38,39,3,8,4,0,39,40,5,0,0,1,40,1,1,0,0,0,41,42,7,0,0,0,42,
  	3,1,0,0,0,43,44,5,18,0,0,44,5,1,0,0,0,45,48,3,4,2,0,46,48,3,2,1,0,47,
  	45,1,0,0,0,47,46,1,0,0,0,48,7,1,0,0,0,49,50,6,4,-1,0,50,53,3,10,5,0,51,
  	53,1,0,0,0,52,49,1,0,0,0,52,51,1,0,0,0,53,58,1,0,0,0,54,55,10,2,0,0,55,
  	57,3,10,5,0,56,54,1,0,0,0,57,60,1,0,0,0,58,56,1,0,0,0,58,59,1,0,0,0,59,
  	9,1,0,0,0,60,58,1,0,0,0,61,62,3,14,7,0,62,63,3,16,8,0,63,64,3,12,6,0,
  	64,11,1,0,0,0,65,72,3,18,9,0,66,72,3,30,15,0,67,68,5,1,0,0,68,69,3,36,
  	18,0,69,70,5,2,0,0,70,72,1,0,0,0,71,65,1,0,0,0,71,66,1,0,0,0,71,67,1,
  	0,0,0,72,13,1,0,0,0,73,74,5,3,0,0,74,75,5,16,0,0,75,76,5,4,0,0,76,83,
  	3,6,3,0,77,78,5,3,0,0,78,79,5,16,0,0,79,80,5,5,0,0,80,83,3,6,3,0,81,83,
  	1,0,0,0,82,73,1,0,0,0,82,77,1,0,0,0,82,81,1,0,0,0,83,15,1,0,0,0,84,85,
  	3,32,16,0,85,86,5,6,0,0,86,89,1,0,0,0,87,89,1,0,0,0,88,84,1,0,0,0,88,
  	87,1,0,0,0,89,17,1,0,0,0,90,91,5,16,0,0,91,92,5,7,0,0,92,97,3,20,10,0,
  	93,94,5,16,0,0,94,95,5,7,0,0,95,97,3,6,3,0,96,90,1,0,0,0,96,93,1,0,0,
  	0,97,19,1,0,0,0,98,99,5,16,0,0,99,100,5,8,0,0,100,101,3,22,11,0,101,21,
  	1,0,0,0,102,103,5,16,0,0,103,104,5,9,0,0,104,105,3,24,12,0,105,106,5,
  	10,0,0,106,23,1,0,0,0,107,110,3,26,13,0,108,110,1,0,0,0,109,107,1,0,0,
  	0,109,108,1,0,0,0,110,25,1,0,0,0,111,112,6,13,-1,0,112,113,3,28,14,0,
  	113,119,1,0,0,0,114,115,10,1,0,0,115,116,5,11,0,0,116,118,3,28,14,0,117,
  	114,1,0,0,0,118,121,1,0,0,0,119,117,1,0,0,0,119,120,1,0,0,0,120,27,1,
  	0,0,0,121,119,1,0,0,0,122,123,5,16,0,0,123,124,5,7,0,0,124,130,3,6,3,
  	0,125,126,5,16,0,0,126,127,5,8,0,0,127,130,3,6,3,0,128,130,3,6,3,0,129,
  	122,1,0,0,0,129,125,1,0,0,0,129,128,1,0,0,0,130,29,1,0,0,0,131,132,3,
  	32,16,0,132,133,5,12,0,0,133,134,3,2,1,0,134,142,1,0,0,0,135,136,5,1,
  	0,0,136,137,3,34,17,0,137,138,5,2,0,0,138,139,5,12,0,0,139,140,3,2,1,
  	0,140,142,1,0,0,0,141,131,1,0,0,0,141,135,1,0,0,0,142,31,1,0,0,0,143,
  	144,6,16,-1,0,144,145,3,22,11,0,145,151,1,0,0,0,146,147,10,1,0,0,147,
  	148,5,6,0,0,148,150,3,32,16,2,149,146,1,0,0,0,150,153,1,0,0,0,151,149,
  	1,0,0,0,151,152,1,0,0,0,152,33,1,0,0,0,153,151,1,0,0,0,154,159,3,32,16,
  	0,155,156,3,32,16,0,156,157,3,34,17,0,157,159,1,0,0,0,158,154,1,0,0,0,
  	158,155,1,0,0,0,159,35,1,0,0,0,160,161,3,30,15,0,161,162,3,36,18,0,162,
  	168,1,0,0,0,163,164,3,18,9,0,164,165,3,36,18,0,165,168,1,0,0,0,166,168,
  	1,0,0,0,167,160,1,0,0,0,167,163,1,0,0,0,167,166,1,0,0,0,168,37,1,0,0,
  	0,14,47,52,58,71,82,88,96,109,119,129,141,151,158,167
  };
  staticData->serializedATN = antlr4::atn::SerializedATNView(serializedATNSegment, sizeof(serializedATNSegment) / sizeof(serializedATNSegment[0]));

  antlr4::atn::ATNDeserializer deserializer;
  staticData->atn = deserializer.deserialize(staticData->serializedATN);

  const size_t count = staticData->atn->getNumberOfDecisions();
  staticData->decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    staticData->decisionToDFA.emplace_back(staticData->atn->getDecisionState(i), i);
  }
  routingaParserStaticData = staticData.release();
}

}

routingAParser::routingAParser(TokenStream *input) : routingAParser(input, antlr4::atn::ParserATNSimulatorOptions()) {}

routingAParser::routingAParser(TokenStream *input, const antlr4::atn::ParserATNSimulatorOptions &options) : Parser(input) {
  routingAParser::initialize();
  _interpreter = new atn::ParserATNSimulator(this, *routingaParserStaticData->atn, routingaParserStaticData->decisionToDFA, routingaParserStaticData->sharedContextCache, options);
}

routingAParser::~routingAParser() {
  delete _interpreter;
}

const atn::ATN& routingAParser::getATN() const {
  return *routingaParserStaticData->atn;
}

std::string routingAParser::getGrammarFileName() const {
  return "routingA.g4";
}

const std::vector<std::string>& routingAParser::getRuleNames() const {
  return routingaParserStaticData->ruleNames;
}

const dfa::Vocabulary& routingAParser::getVocabulary() const {
  return routingaParserStaticData->vocabulary;
}

antlr4::atn::SerializedATNView routingAParser::getSerializedATN() const {
  return routingaParserStaticData->serializedATN;
}


//----------------- StartContext ------------------------------------------------------------------

routingAParser::StartContext::StartContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::InputContext* routingAParser::StartContext::input() {
  return getRuleContext<routingAParser::InputContext>(0);
}

tree::TerminalNode* routingAParser::StartContext::EOF() {
  return getToken(routingAParser::EOF, 0);
}


size_t routingAParser::StartContext::getRuleIndex() const {
  return routingAParser::RuleStart;
}

void routingAParser::StartContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterStart(this);
}

void routingAParser::StartContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitStart(this);
}

routingAParser::StartContext* routingAParser::start() {
  StartContext *_localctx = _tracker.createInstance<StartContext>(_ctx, getState());
  enterRule(_localctx, 0, routingAParser::RuleStart);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(38);
    input(0);
    setState(39);
    match(routingAParser::EOF);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- Bare_literalContext ------------------------------------------------------------------

routingAParser::Bare_literalContext::Bare_literalContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* routingAParser::Bare_literalContext::ID() {
  return getToken(routingAParser::ID, 0);
}

tree::TerminalNode* routingAParser::Bare_literalContext::NON_ID() {
  return getToken(routingAParser::NON_ID, 0);
}


size_t routingAParser::Bare_literalContext::getRuleIndex() const {
  return routingAParser::RuleBare_literal;
}

void routingAParser::Bare_literalContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterBare_literal(this);
}

void routingAParser::Bare_literalContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitBare_literal(this);
}

routingAParser::Bare_literalContext* routingAParser::bare_literal() {
  Bare_literalContext *_localctx = _tracker.createInstance<Bare_literalContext>(_ctx, getState());
  enterRule(_localctx, 2, routingAParser::RuleBare_literal);
  size_t _la = 0;

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(41);
    _la = _input->LA(1);
    if (!(_la == routingAParser::ID

    || _la == routingAParser::NON_ID)) {
    _errHandler->recoverInline(this);
    }
    else {
      _errHandler->reportMatch(this);
      consume();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- Quote_literalContext ------------------------------------------------------------------

routingAParser::Quote_literalContext::Quote_literalContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* routingAParser::Quote_literalContext::QUOTE_STRING() {
  return getToken(routingAParser::QUOTE_STRING, 0);
}


size_t routingAParser::Quote_literalContext::getRuleIndex() const {
  return routingAParser::RuleQuote_literal;
}

void routingAParser::Quote_literalContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterQuote_literal(this);
}

void routingAParser::Quote_literalContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitQuote_literal(this);
}

routingAParser::Quote_literalContext* routingAParser::quote_literal() {
  Quote_literalContext *_localctx = _tracker.createInstance<Quote_literalContext>(_ctx, getState());
  enterRule(_localctx, 4, routingAParser::RuleQuote_literal);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(43);
    match(routingAParser::QUOTE_STRING);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- LiteralContext ------------------------------------------------------------------

routingAParser::LiteralContext::LiteralContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::Quote_literalContext* routingAParser::LiteralContext::quote_literal() {
  return getRuleContext<routingAParser::Quote_literalContext>(0);
}

routingAParser::Bare_literalContext* routingAParser::LiteralContext::bare_literal() {
  return getRuleContext<routingAParser::Bare_literalContext>(0);
}


size_t routingAParser::LiteralContext::getRuleIndex() const {
  return routingAParser::RuleLiteral;
}

void routingAParser::LiteralContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterLiteral(this);
}

void routingAParser::LiteralContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitLiteral(this);
}

routingAParser::LiteralContext* routingAParser::literal() {
  LiteralContext *_localctx = _tracker.createInstance<LiteralContext>(_ctx, getState());
  enterRule(_localctx, 6, routingAParser::RuleLiteral);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(47);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case routingAParser::QUOTE_STRING: {
        enterOuterAlt(_localctx, 1);
        setState(45);
        quote_literal();
        break;
      }

      case routingAParser::ID:
      case routingAParser::NON_ID: {
        enterOuterAlt(_localctx, 2);
        setState(46);
        bare_literal();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- InputContext ------------------------------------------------------------------

routingAParser::InputContext::InputContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::ProgramStructureBlcokContext* routingAParser::InputContext::programStructureBlcok() {
  return getRuleContext<routingAParser::ProgramStructureBlcokContext>(0);
}

routingAParser::InputContext* routingAParser::InputContext::input() {
  return getRuleContext<routingAParser::InputContext>(0);
}


size_t routingAParser::InputContext::getRuleIndex() const {
  return routingAParser::RuleInput;
}

void routingAParser::InputContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterInput(this);
}

void routingAParser::InputContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitInput(this);
}


routingAParser::InputContext* routingAParser::input() {
   return input(0);
}

routingAParser::InputContext* routingAParser::input(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  routingAParser::InputContext *_localctx = _tracker.createInstance<InputContext>(_ctx, parentState);
  routingAParser::InputContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 8;
  enterRecursionRule(_localctx, 8, routingAParser::RuleInput, precedence);

    

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(52);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 1, _ctx)) {
    case 1: {
      setState(50);
      programStructureBlcok();
      break;
    }

    case 2: {
      break;
    }

    default:
      break;
    }
    _ctx->stop = _input->LT(-1);
    setState(58);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 2, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<InputContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleInput);
        setState(54);

        if (!(precpred(_ctx, 2))) throw FailedPredicateException(this, "precpred(_ctx, 2)");
        setState(55);
        programStructureBlcok(); 
      }
      setState(60);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 2, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ProgramStructureBlcokContext ------------------------------------------------------------------

routingAParser::ProgramStructureBlcokContext::ProgramStructureBlcokContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::OptConditionStatementContext* routingAParser::ProgramStructureBlcokContext::optConditionStatement() {
  return getRuleContext<routingAParser::OptConditionStatementContext>(0);
}

routingAParser::OptFunctionPrototypeExpressionAndContext* routingAParser::ProgramStructureBlcokContext::optFunctionPrototypeExpressionAnd() {
  return getRuleContext<routingAParser::OptFunctionPrototypeExpressionAndContext>(0);
}

routingAParser::ExpressionContext* routingAParser::ProgramStructureBlcokContext::expression() {
  return getRuleContext<routingAParser::ExpressionContext>(0);
}


size_t routingAParser::ProgramStructureBlcokContext::getRuleIndex() const {
  return routingAParser::RuleProgramStructureBlcok;
}

void routingAParser::ProgramStructureBlcokContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterProgramStructureBlcok(this);
}

void routingAParser::ProgramStructureBlcokContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitProgramStructureBlcok(this);
}

routingAParser::ProgramStructureBlcokContext* routingAParser::programStructureBlcok() {
  ProgramStructureBlcokContext *_localctx = _tracker.createInstance<ProgramStructureBlcokContext>(_ctx, getState());
  enterRule(_localctx, 10, routingAParser::RuleProgramStructureBlcok);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(61);
    optConditionStatement();
    setState(62);
    optFunctionPrototypeExpressionAnd();
    setState(63);
    expression();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExpressionContext ------------------------------------------------------------------

routingAParser::ExpressionContext::ExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::DeclarationContext* routingAParser::ExpressionContext::declaration() {
  return getRuleContext<routingAParser::DeclarationContext>(0);
}

routingAParser::RoutingRuleContext* routingAParser::ExpressionContext::routingRule() {
  return getRuleContext<routingAParser::RoutingRuleContext>(0);
}

routingAParser::RoutingRuleOrDeclarationListContext* routingAParser::ExpressionContext::routingRuleOrDeclarationList() {
  return getRuleContext<routingAParser::RoutingRuleOrDeclarationListContext>(0);
}


size_t routingAParser::ExpressionContext::getRuleIndex() const {
  return routingAParser::RuleExpression;
}

void routingAParser::ExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExpression(this);
}

void routingAParser::ExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExpression(this);
}

routingAParser::ExpressionContext* routingAParser::expression() {
  ExpressionContext *_localctx = _tracker.createInstance<ExpressionContext>(_ctx, getState());
  enterRule(_localctx, 12, routingAParser::RuleExpression);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(71);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 3, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(65);
      declaration();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(66);
      routingRule();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(67);
      match(routingAParser::T__0);
      setState(68);
      routingRuleOrDeclarationList();
      setState(69);
      match(routingAParser::T__1);
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- OptConditionStatementContext ------------------------------------------------------------------

routingAParser::OptConditionStatementContext::OptConditionStatementContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* routingAParser::OptConditionStatementContext::ID() {
  return getToken(routingAParser::ID, 0);
}

routingAParser::LiteralContext* routingAParser::OptConditionStatementContext::literal() {
  return getRuleContext<routingAParser::LiteralContext>(0);
}


size_t routingAParser::OptConditionStatementContext::getRuleIndex() const {
  return routingAParser::RuleOptConditionStatement;
}

void routingAParser::OptConditionStatementContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterOptConditionStatement(this);
}

void routingAParser::OptConditionStatementContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitOptConditionStatement(this);
}

routingAParser::OptConditionStatementContext* routingAParser::optConditionStatement() {
  OptConditionStatementContext *_localctx = _tracker.createInstance<OptConditionStatementContext>(_ctx, getState());
  enterRule(_localctx, 14, routingAParser::RuleOptConditionStatement);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(82);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 4, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(73);
      match(routingAParser::T__2);
      setState(74);
      match(routingAParser::ID);
      setState(75);
      match(routingAParser::T__3);
      setState(76);
      literal();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(77);
      match(routingAParser::T__2);
      setState(78);
      match(routingAParser::ID);
      setState(79);
      match(routingAParser::T__4);
      setState(80);
      literal();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);

      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- OptFunctionPrototypeExpressionAndContext ------------------------------------------------------------------

routingAParser::OptFunctionPrototypeExpressionAndContext::OptFunctionPrototypeExpressionAndContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::FunctionPrototypeExpressionContext* routingAParser::OptFunctionPrototypeExpressionAndContext::functionPrototypeExpression() {
  return getRuleContext<routingAParser::FunctionPrototypeExpressionContext>(0);
}


size_t routingAParser::OptFunctionPrototypeExpressionAndContext::getRuleIndex() const {
  return routingAParser::RuleOptFunctionPrototypeExpressionAnd;
}

void routingAParser::OptFunctionPrototypeExpressionAndContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterOptFunctionPrototypeExpressionAnd(this);
}

void routingAParser::OptFunctionPrototypeExpressionAndContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitOptFunctionPrototypeExpressionAnd(this);
}

routingAParser::OptFunctionPrototypeExpressionAndContext* routingAParser::optFunctionPrototypeExpressionAnd() {
  OptFunctionPrototypeExpressionAndContext *_localctx = _tracker.createInstance<OptFunctionPrototypeExpressionAndContext>(_ctx, getState());
  enterRule(_localctx, 16, routingAParser::RuleOptFunctionPrototypeExpressionAnd);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(88);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 5, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(84);
      functionPrototypeExpression(0);
      setState(85);
      match(routingAParser::T__5);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);

      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DeclarationContext ------------------------------------------------------------------

routingAParser::DeclarationContext::DeclarationContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* routingAParser::DeclarationContext::ID() {
  return getToken(routingAParser::ID, 0);
}

routingAParser::AssignmentExpressionContext* routingAParser::DeclarationContext::assignmentExpression() {
  return getRuleContext<routingAParser::AssignmentExpressionContext>(0);
}

routingAParser::LiteralContext* routingAParser::DeclarationContext::literal() {
  return getRuleContext<routingAParser::LiteralContext>(0);
}


size_t routingAParser::DeclarationContext::getRuleIndex() const {
  return routingAParser::RuleDeclaration;
}

void routingAParser::DeclarationContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterDeclaration(this);
}

void routingAParser::DeclarationContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitDeclaration(this);
}

routingAParser::DeclarationContext* routingAParser::declaration() {
  DeclarationContext *_localctx = _tracker.createInstance<DeclarationContext>(_ctx, getState());
  enterRule(_localctx, 18, routingAParser::RuleDeclaration);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(96);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 6, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(90);
      match(routingAParser::ID);
      setState(91);
      match(routingAParser::T__6);
      setState(92);
      assignmentExpression();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(93);
      match(routingAParser::ID);
      setState(94);
      match(routingAParser::T__6);
      setState(95);
      literal();
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- AssignmentExpressionContext ------------------------------------------------------------------

routingAParser::AssignmentExpressionContext::AssignmentExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* routingAParser::AssignmentExpressionContext::ID() {
  return getToken(routingAParser::ID, 0);
}

routingAParser::FunctionPrototypeContext* routingAParser::AssignmentExpressionContext::functionPrototype() {
  return getRuleContext<routingAParser::FunctionPrototypeContext>(0);
}


size_t routingAParser::AssignmentExpressionContext::getRuleIndex() const {
  return routingAParser::RuleAssignmentExpression;
}

void routingAParser::AssignmentExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAssignmentExpression(this);
}

void routingAParser::AssignmentExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAssignmentExpression(this);
}

routingAParser::AssignmentExpressionContext* routingAParser::assignmentExpression() {
  AssignmentExpressionContext *_localctx = _tracker.createInstance<AssignmentExpressionContext>(_ctx, getState());
  enterRule(_localctx, 20, routingAParser::RuleAssignmentExpression);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(98);
    match(routingAParser::ID);
    setState(99);
    match(routingAParser::T__7);
    setState(100);
    functionPrototype();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FunctionPrototypeContext ------------------------------------------------------------------

routingAParser::FunctionPrototypeContext::FunctionPrototypeContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* routingAParser::FunctionPrototypeContext::ID() {
  return getToken(routingAParser::ID, 0);
}

routingAParser::ParameterListContext* routingAParser::FunctionPrototypeContext::parameterList() {
  return getRuleContext<routingAParser::ParameterListContext>(0);
}


size_t routingAParser::FunctionPrototypeContext::getRuleIndex() const {
  return routingAParser::RuleFunctionPrototype;
}

void routingAParser::FunctionPrototypeContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFunctionPrototype(this);
}

void routingAParser::FunctionPrototypeContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFunctionPrototype(this);
}

routingAParser::FunctionPrototypeContext* routingAParser::functionPrototype() {
  FunctionPrototypeContext *_localctx = _tracker.createInstance<FunctionPrototypeContext>(_ctx, getState());
  enterRule(_localctx, 22, routingAParser::RuleFunctionPrototype);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(102);
    match(routingAParser::ID);
    setState(103);
    match(routingAParser::T__8);
    setState(104);
    parameterList();
    setState(105);
    match(routingAParser::T__9);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ParameterListContext ------------------------------------------------------------------

routingAParser::ParameterListContext::ParameterListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::NonEmptyParameterListContext* routingAParser::ParameterListContext::nonEmptyParameterList() {
  return getRuleContext<routingAParser::NonEmptyParameterListContext>(0);
}


size_t routingAParser::ParameterListContext::getRuleIndex() const {
  return routingAParser::RuleParameterList;
}

void routingAParser::ParameterListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterParameterList(this);
}

void routingAParser::ParameterListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitParameterList(this);
}

routingAParser::ParameterListContext* routingAParser::parameterList() {
  ParameterListContext *_localctx = _tracker.createInstance<ParameterListContext>(_ctx, getState());
  enterRule(_localctx, 24, routingAParser::RuleParameterList);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(109);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case routingAParser::ID:
      case routingAParser::NON_ID:
      case routingAParser::QUOTE_STRING: {
        enterOuterAlt(_localctx, 1);
        setState(107);
        nonEmptyParameterList(0);
        break;
      }

      case routingAParser::T__9: {
        enterOuterAlt(_localctx, 2);

        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- NonEmptyParameterListContext ------------------------------------------------------------------

routingAParser::NonEmptyParameterListContext::NonEmptyParameterListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::ParameterContext* routingAParser::NonEmptyParameterListContext::parameter() {
  return getRuleContext<routingAParser::ParameterContext>(0);
}

routingAParser::NonEmptyParameterListContext* routingAParser::NonEmptyParameterListContext::nonEmptyParameterList() {
  return getRuleContext<routingAParser::NonEmptyParameterListContext>(0);
}


size_t routingAParser::NonEmptyParameterListContext::getRuleIndex() const {
  return routingAParser::RuleNonEmptyParameterList;
}

void routingAParser::NonEmptyParameterListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterNonEmptyParameterList(this);
}

void routingAParser::NonEmptyParameterListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitNonEmptyParameterList(this);
}


routingAParser::NonEmptyParameterListContext* routingAParser::nonEmptyParameterList() {
   return nonEmptyParameterList(0);
}

routingAParser::NonEmptyParameterListContext* routingAParser::nonEmptyParameterList(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  routingAParser::NonEmptyParameterListContext *_localctx = _tracker.createInstance<NonEmptyParameterListContext>(_ctx, parentState);
  routingAParser::NonEmptyParameterListContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 26;
  enterRecursionRule(_localctx, 26, routingAParser::RuleNonEmptyParameterList, precedence);

    

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(112);
    parameter();
    _ctx->stop = _input->LT(-1);
    setState(119);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 8, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<NonEmptyParameterListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleNonEmptyParameterList);
        setState(114);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(115);
        match(routingAParser::T__10);
        setState(116);
        parameter(); 
      }
      setState(121);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 8, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- ParameterContext ------------------------------------------------------------------

routingAParser::ParameterContext::ParameterContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* routingAParser::ParameterContext::ID() {
  return getToken(routingAParser::ID, 0);
}

routingAParser::LiteralContext* routingAParser::ParameterContext::literal() {
  return getRuleContext<routingAParser::LiteralContext>(0);
}


size_t routingAParser::ParameterContext::getRuleIndex() const {
  return routingAParser::RuleParameter;
}

void routingAParser::ParameterContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterParameter(this);
}

void routingAParser::ParameterContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitParameter(this);
}

routingAParser::ParameterContext* routingAParser::parameter() {
  ParameterContext *_localctx = _tracker.createInstance<ParameterContext>(_ctx, getState());
  enterRule(_localctx, 28, routingAParser::RuleParameter);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(129);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 9, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(122);
      match(routingAParser::ID);
      setState(123);
      match(routingAParser::T__6);
      setState(124);
      literal();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(125);
      match(routingAParser::ID);
      setState(126);
      match(routingAParser::T__7);
      setState(127);
      literal();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(128);
      literal();
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- RoutingRuleContext ------------------------------------------------------------------

routingAParser::RoutingRuleContext::RoutingRuleContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::FunctionPrototypeExpressionContext* routingAParser::RoutingRuleContext::functionPrototypeExpression() {
  return getRuleContext<routingAParser::FunctionPrototypeExpressionContext>(0);
}

routingAParser::Bare_literalContext* routingAParser::RoutingRuleContext::bare_literal() {
  return getRuleContext<routingAParser::Bare_literalContext>(0);
}

routingAParser::FunctionPrototypeExpressionListContext* routingAParser::RoutingRuleContext::functionPrototypeExpressionList() {
  return getRuleContext<routingAParser::FunctionPrototypeExpressionListContext>(0);
}


size_t routingAParser::RoutingRuleContext::getRuleIndex() const {
  return routingAParser::RuleRoutingRule;
}

void routingAParser::RoutingRuleContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRoutingRule(this);
}

void routingAParser::RoutingRuleContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRoutingRule(this);
}

routingAParser::RoutingRuleContext* routingAParser::routingRule() {
  RoutingRuleContext *_localctx = _tracker.createInstance<RoutingRuleContext>(_ctx, getState());
  enterRule(_localctx, 30, routingAParser::RuleRoutingRule);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(141);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case routingAParser::ID: {
        enterOuterAlt(_localctx, 1);
        setState(131);
        functionPrototypeExpression(0);
        setState(132);
        match(routingAParser::T__11);
        setState(133);
        bare_literal();
        break;
      }

      case routingAParser::T__0: {
        enterOuterAlt(_localctx, 2);
        setState(135);
        match(routingAParser::T__0);
        setState(136);
        functionPrototypeExpressionList();
        setState(137);
        match(routingAParser::T__1);
        setState(138);
        match(routingAParser::T__11);
        setState(139);
        bare_literal();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FunctionPrototypeExpressionContext ------------------------------------------------------------------

routingAParser::FunctionPrototypeExpressionContext::FunctionPrototypeExpressionContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::FunctionPrototypeContext* routingAParser::FunctionPrototypeExpressionContext::functionPrototype() {
  return getRuleContext<routingAParser::FunctionPrototypeContext>(0);
}

std::vector<routingAParser::FunctionPrototypeExpressionContext *> routingAParser::FunctionPrototypeExpressionContext::functionPrototypeExpression() {
  return getRuleContexts<routingAParser::FunctionPrototypeExpressionContext>();
}

routingAParser::FunctionPrototypeExpressionContext* routingAParser::FunctionPrototypeExpressionContext::functionPrototypeExpression(size_t i) {
  return getRuleContext<routingAParser::FunctionPrototypeExpressionContext>(i);
}


size_t routingAParser::FunctionPrototypeExpressionContext::getRuleIndex() const {
  return routingAParser::RuleFunctionPrototypeExpression;
}

void routingAParser::FunctionPrototypeExpressionContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFunctionPrototypeExpression(this);
}

void routingAParser::FunctionPrototypeExpressionContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFunctionPrototypeExpression(this);
}


routingAParser::FunctionPrototypeExpressionContext* routingAParser::functionPrototypeExpression() {
   return functionPrototypeExpression(0);
}

routingAParser::FunctionPrototypeExpressionContext* routingAParser::functionPrototypeExpression(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  routingAParser::FunctionPrototypeExpressionContext *_localctx = _tracker.createInstance<FunctionPrototypeExpressionContext>(_ctx, parentState);
  routingAParser::FunctionPrototypeExpressionContext *previousContext = _localctx;
  (void)previousContext; // Silence compiler, in case the context is not used by generated code.
  size_t startState = 32;
  enterRecursionRule(_localctx, 32, routingAParser::RuleFunctionPrototypeExpression, precedence);

    

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(144);
    functionPrototype();
    _ctx->stop = _input->LT(-1);
    setState(151);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 11, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<FunctionPrototypeExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleFunctionPrototypeExpression);
        setState(146);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(147);
        match(routingAParser::T__5);
        setState(148);
        functionPrototypeExpression(2); 
      }
      setState(153);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 11, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- FunctionPrototypeExpressionListContext ------------------------------------------------------------------

routingAParser::FunctionPrototypeExpressionListContext::FunctionPrototypeExpressionListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::FunctionPrototypeExpressionContext* routingAParser::FunctionPrototypeExpressionListContext::functionPrototypeExpression() {
  return getRuleContext<routingAParser::FunctionPrototypeExpressionContext>(0);
}

routingAParser::FunctionPrototypeExpressionListContext* routingAParser::FunctionPrototypeExpressionListContext::functionPrototypeExpressionList() {
  return getRuleContext<routingAParser::FunctionPrototypeExpressionListContext>(0);
}


size_t routingAParser::FunctionPrototypeExpressionListContext::getRuleIndex() const {
  return routingAParser::RuleFunctionPrototypeExpressionList;
}

void routingAParser::FunctionPrototypeExpressionListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFunctionPrototypeExpressionList(this);
}

void routingAParser::FunctionPrototypeExpressionListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFunctionPrototypeExpressionList(this);
}

routingAParser::FunctionPrototypeExpressionListContext* routingAParser::functionPrototypeExpressionList() {
  FunctionPrototypeExpressionListContext *_localctx = _tracker.createInstance<FunctionPrototypeExpressionListContext>(_ctx, getState());
  enterRule(_localctx, 34, routingAParser::RuleFunctionPrototypeExpressionList);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(158);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 12, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(154);
      functionPrototypeExpression(0);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(155);
      functionPrototypeExpression(0);
      setState(156);
      functionPrototypeExpressionList();
      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- RoutingRuleOrDeclarationListContext ------------------------------------------------------------------

routingAParser::RoutingRuleOrDeclarationListContext::RoutingRuleOrDeclarationListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::RoutingRuleContext* routingAParser::RoutingRuleOrDeclarationListContext::routingRule() {
  return getRuleContext<routingAParser::RoutingRuleContext>(0);
}

routingAParser::RoutingRuleOrDeclarationListContext* routingAParser::RoutingRuleOrDeclarationListContext::routingRuleOrDeclarationList() {
  return getRuleContext<routingAParser::RoutingRuleOrDeclarationListContext>(0);
}

routingAParser::DeclarationContext* routingAParser::RoutingRuleOrDeclarationListContext::declaration() {
  return getRuleContext<routingAParser::DeclarationContext>(0);
}


size_t routingAParser::RoutingRuleOrDeclarationListContext::getRuleIndex() const {
  return routingAParser::RuleRoutingRuleOrDeclarationList;
}

void routingAParser::RoutingRuleOrDeclarationListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRoutingRuleOrDeclarationList(this);
}

void routingAParser::RoutingRuleOrDeclarationListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRoutingRuleOrDeclarationList(this);
}

routingAParser::RoutingRuleOrDeclarationListContext* routingAParser::routingRuleOrDeclarationList() {
  RoutingRuleOrDeclarationListContext *_localctx = _tracker.createInstance<RoutingRuleOrDeclarationListContext>(_ctx, getState());
  enterRule(_localctx, 36, routingAParser::RuleRoutingRuleOrDeclarationList);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(167);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(160);
      routingRule();
      setState(161);
      routingRuleOrDeclarationList();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(163);
      declaration();
      setState(164);
      routingRuleOrDeclarationList();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);

      break;
    }

    default:
      break;
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

bool routingAParser::sempred(RuleContext *context, size_t ruleIndex, size_t predicateIndex) {
  switch (ruleIndex) {
    case 4: return inputSempred(antlrcpp::downCast<InputContext *>(context), predicateIndex);
    case 13: return nonEmptyParameterListSempred(antlrcpp::downCast<NonEmptyParameterListContext *>(context), predicateIndex);
    case 16: return functionPrototypeExpressionSempred(antlrcpp::downCast<FunctionPrototypeExpressionContext *>(context), predicateIndex);

  default:
    break;
  }
  return true;
}

bool routingAParser::inputSempred(InputContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 0: return precpred(_ctx, 2);

  default:
    break;
  }
  return true;
}

bool routingAParser::nonEmptyParameterListSempred(NonEmptyParameterListContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 1: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

bool routingAParser::functionPrototypeExpressionSempred(FunctionPrototypeExpressionContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 2: return precpred(_ctx, 1);

  default:
    break;
  }
  return true;
}

void routingAParser::initialize() {
  std::call_once(routingaParserOnceFlag, routingaParserInitialize);
}
