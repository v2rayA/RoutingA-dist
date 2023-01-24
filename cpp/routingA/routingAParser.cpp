
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
      "expression", "optConditionStatement", "declaration", "assignmentExpression", 
      "functionPrototype", "optParameterList", "nonEmptyParameterList", 
      "parameter", "routingRule", "routingRuleLeft", "routingRuleLeftList", 
      "optFunctionPrototypeExpressionAnd", "functionPrototypeExpression", 
      "routingRuleOrDeclarationList", "routingRuleList"
    },
    std::vector<std::string>{
      "", "'{'", "'}'", "'@'", "'=='", "'!='", "':'", "'='", "'('", "')'", 
      "','", "'->'", "'&&'"
    },
    std::vector<std::string>{
      "", "", "", "", "", "", "", "", "", "", "", "", "", "WHITESPACE", 
      "COMMENT_BLOCK", "COMMENT_LINE_SHARP", "ID", "NON_ID", "QUOTE_STRING"
    }
  );
  static const int32_t serializedATNSegment[] = {
  	4,1,18,189,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,
  	7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,
  	14,2,15,7,15,2,16,7,16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,7,20,1,0,1,
  	0,1,0,1,1,1,1,1,2,1,2,1,3,1,3,3,3,52,8,3,1,4,1,4,1,4,3,4,57,8,4,1,4,1,
  	4,5,4,61,8,4,10,4,12,4,64,9,4,1,5,1,5,1,5,1,6,1,6,1,6,1,6,1,6,1,6,3,6,
  	75,8,6,1,7,1,7,1,7,1,7,1,7,1,7,1,7,1,7,1,7,3,7,86,8,7,1,8,1,8,1,8,1,8,
  	1,8,1,8,3,8,94,8,8,1,9,1,9,1,9,1,9,1,10,1,10,1,10,1,10,1,10,1,11,1,11,
  	3,11,107,8,11,1,12,1,12,1,12,1,12,1,12,1,12,5,12,115,8,12,10,12,12,12,
  	118,9,12,1,13,1,13,1,13,1,13,1,13,1,13,1,13,3,13,127,8,13,1,14,1,14,1,
  	14,1,14,1,14,1,14,1,14,1,14,1,14,3,14,138,8,14,1,15,1,15,1,15,1,15,1,
  	15,1,15,1,15,1,15,3,15,148,8,15,1,16,1,16,1,16,1,16,3,16,154,8,16,1,17,
  	1,17,1,17,1,17,3,17,160,8,17,1,18,1,18,1,18,1,18,1,18,1,18,5,18,168,8,
  	18,10,18,12,18,171,9,18,1,19,1,19,1,19,1,19,1,19,1,19,1,19,1,19,3,19,
  	181,8,19,1,20,1,20,1,20,1,20,3,20,187,8,20,1,20,0,3,8,24,36,21,0,2,4,
  	6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,0,1,1,0,16,17,188,
  	0,42,1,0,0,0,2,45,1,0,0,0,4,47,1,0,0,0,6,51,1,0,0,0,8,56,1,0,0,0,10,65,
  	1,0,0,0,12,74,1,0,0,0,14,85,1,0,0,0,16,93,1,0,0,0,18,95,1,0,0,0,20,99,
  	1,0,0,0,22,106,1,0,0,0,24,108,1,0,0,0,26,126,1,0,0,0,28,137,1,0,0,0,30,
  	147,1,0,0,0,32,153,1,0,0,0,34,159,1,0,0,0,36,161,1,0,0,0,38,180,1,0,0,
  	0,40,186,1,0,0,0,42,43,3,8,4,0,43,44,5,0,0,1,44,1,1,0,0,0,45,46,7,0,0,
  	0,46,3,1,0,0,0,47,48,5,18,0,0,48,5,1,0,0,0,49,52,3,4,2,0,50,52,3,2,1,
  	0,51,49,1,0,0,0,51,50,1,0,0,0,52,7,1,0,0,0,53,54,6,4,-1,0,54,57,3,10,
  	5,0,55,57,1,0,0,0,56,53,1,0,0,0,56,55,1,0,0,0,57,62,1,0,0,0,58,59,10,
  	2,0,0,59,61,3,10,5,0,60,58,1,0,0,0,61,64,1,0,0,0,62,60,1,0,0,0,62,63,
  	1,0,0,0,63,9,1,0,0,0,64,62,1,0,0,0,65,66,3,14,7,0,66,67,3,12,6,0,67,11,
  	1,0,0,0,68,75,3,16,8,0,69,75,3,28,14,0,70,71,5,1,0,0,71,72,3,38,19,0,
  	72,73,5,2,0,0,73,75,1,0,0,0,74,68,1,0,0,0,74,69,1,0,0,0,74,70,1,0,0,0,
  	75,13,1,0,0,0,76,77,5,3,0,0,77,78,5,16,0,0,78,79,5,4,0,0,79,86,3,6,3,
  	0,80,81,5,3,0,0,81,82,5,16,0,0,82,83,5,5,0,0,83,86,3,6,3,0,84,86,1,0,
  	0,0,85,76,1,0,0,0,85,80,1,0,0,0,85,84,1,0,0,0,86,15,1,0,0,0,87,88,5,16,
  	0,0,88,89,5,6,0,0,89,94,3,18,9,0,90,91,5,16,0,0,91,92,5,6,0,0,92,94,3,
  	6,3,0,93,87,1,0,0,0,93,90,1,0,0,0,94,17,1,0,0,0,95,96,5,16,0,0,96,97,
  	5,7,0,0,97,98,3,20,10,0,98,19,1,0,0,0,99,100,5,16,0,0,100,101,5,8,0,0,
  	101,102,3,22,11,0,102,103,5,9,0,0,103,21,1,0,0,0,104,107,3,24,12,0,105,
  	107,1,0,0,0,106,104,1,0,0,0,106,105,1,0,0,0,107,23,1,0,0,0,108,109,6,
  	12,-1,0,109,110,3,26,13,0,110,116,1,0,0,0,111,112,10,1,0,0,112,113,5,
  	10,0,0,113,115,3,26,13,0,114,111,1,0,0,0,115,118,1,0,0,0,116,114,1,0,
  	0,0,116,117,1,0,0,0,117,25,1,0,0,0,118,116,1,0,0,0,119,120,5,16,0,0,120,
  	121,5,6,0,0,121,127,3,6,3,0,122,123,5,16,0,0,123,124,5,7,0,0,124,127,
  	3,6,3,0,125,127,3,6,3,0,126,119,1,0,0,0,126,122,1,0,0,0,126,125,1,0,0,
  	0,127,27,1,0,0,0,128,129,3,30,15,0,129,130,5,11,0,0,130,131,3,2,1,0,131,
  	138,1,0,0,0,132,133,3,34,17,0,133,134,5,1,0,0,134,135,3,40,20,0,135,136,
  	5,2,0,0,136,138,1,0,0,0,137,128,1,0,0,0,137,132,1,0,0,0,138,29,1,0,0,
  	0,139,140,3,34,17,0,140,141,5,1,0,0,141,142,3,32,16,0,142,143,5,2,0,0,
  	143,148,1,0,0,0,144,145,3,34,17,0,145,146,3,36,18,0,146,148,1,0,0,0,147,
  	139,1,0,0,0,147,144,1,0,0,0,148,31,1,0,0,0,149,154,3,30,15,0,150,151,
  	3,30,15,0,151,152,3,32,16,0,152,154,1,0,0,0,153,149,1,0,0,0,153,150,1,
  	0,0,0,154,33,1,0,0,0,155,156,3,36,18,0,156,157,5,12,0,0,157,160,1,0,0,
  	0,158,160,1,0,0,0,159,155,1,0,0,0,159,158,1,0,0,0,160,35,1,0,0,0,161,
  	162,6,18,-1,0,162,163,3,20,10,0,163,169,1,0,0,0,164,165,10,1,0,0,165,
  	166,5,12,0,0,166,168,3,36,18,2,167,164,1,0,0,0,168,171,1,0,0,0,169,167,
  	1,0,0,0,169,170,1,0,0,0,170,37,1,0,0,0,171,169,1,0,0,0,172,181,3,28,14,
  	0,173,181,3,16,8,0,174,175,3,28,14,0,175,176,3,38,19,0,176,181,1,0,0,
  	0,177,178,3,16,8,0,178,179,3,38,19,0,179,181,1,0,0,0,180,172,1,0,0,0,
  	180,173,1,0,0,0,180,174,1,0,0,0,180,177,1,0,0,0,181,39,1,0,0,0,182,187,
  	3,28,14,0,183,184,3,28,14,0,184,185,3,40,20,0,185,187,1,0,0,0,186,182,
  	1,0,0,0,186,183,1,0,0,0,187,41,1,0,0,0,16,51,56,62,74,85,93,106,116,126,
  	137,147,153,159,169,180,186
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
    setState(42);
    input(0);
    setState(43);
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
    setState(45);
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
    setState(47);
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
    setState(51);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case routingAParser::QUOTE_STRING: {
        enterOuterAlt(_localctx, 1);
        setState(49);
        quote_literal();
        break;
      }

      case routingAParser::ID:
      case routingAParser::NON_ID: {
        enterOuterAlt(_localctx, 2);
        setState(50);
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
    setState(56);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 1, _ctx)) {
    case 1: {
      setState(54);
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
    setState(62);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 2, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<InputContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleInput);
        setState(58);

        if (!(precpred(_ctx, 2))) throw FailedPredicateException(this, "precpred(_ctx, 2)");
        setState(59);
        programStructureBlcok(); 
      }
      setState(64);
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
    setState(65);
    optConditionStatement();
    setState(66);
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
    setState(74);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 3, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(68);
      declaration();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(69);
      routingRule();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(70);
      match(routingAParser::T__0);
      setState(71);
      routingRuleOrDeclarationList();
      setState(72);
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
    setState(85);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 4, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(76);
      match(routingAParser::T__2);
      setState(77);
      match(routingAParser::ID);
      setState(78);
      match(routingAParser::T__3);
      setState(79);
      literal();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(80);
      match(routingAParser::T__2);
      setState(81);
      match(routingAParser::ID);
      setState(82);
      match(routingAParser::T__4);
      setState(83);
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
  enterRule(_localctx, 16, routingAParser::RuleDeclaration);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(93);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 5, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(87);
      match(routingAParser::ID);
      setState(88);
      match(routingAParser::T__5);
      setState(89);
      assignmentExpression();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(90);
      match(routingAParser::ID);
      setState(91);
      match(routingAParser::T__5);
      setState(92);
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
  enterRule(_localctx, 18, routingAParser::RuleAssignmentExpression);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(95);
    match(routingAParser::ID);
    setState(96);
    match(routingAParser::T__6);
    setState(97);
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

routingAParser::OptParameterListContext* routingAParser::FunctionPrototypeContext::optParameterList() {
  return getRuleContext<routingAParser::OptParameterListContext>(0);
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
  enterRule(_localctx, 20, routingAParser::RuleFunctionPrototype);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(99);
    match(routingAParser::ID);
    setState(100);
    match(routingAParser::T__7);
    setState(101);
    optParameterList();
    setState(102);
    match(routingAParser::T__8);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- OptParameterListContext ------------------------------------------------------------------

routingAParser::OptParameterListContext::OptParameterListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::NonEmptyParameterListContext* routingAParser::OptParameterListContext::nonEmptyParameterList() {
  return getRuleContext<routingAParser::NonEmptyParameterListContext>(0);
}


size_t routingAParser::OptParameterListContext::getRuleIndex() const {
  return routingAParser::RuleOptParameterList;
}

void routingAParser::OptParameterListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterOptParameterList(this);
}

void routingAParser::OptParameterListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitOptParameterList(this);
}

routingAParser::OptParameterListContext* routingAParser::optParameterList() {
  OptParameterListContext *_localctx = _tracker.createInstance<OptParameterListContext>(_ctx, getState());
  enterRule(_localctx, 22, routingAParser::RuleOptParameterList);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(106);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case routingAParser::ID:
      case routingAParser::NON_ID:
      case routingAParser::QUOTE_STRING: {
        enterOuterAlt(_localctx, 1);
        setState(104);
        nonEmptyParameterList(0);
        break;
      }

      case routingAParser::T__8: {
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
  size_t startState = 24;
  enterRecursionRule(_localctx, 24, routingAParser::RuleNonEmptyParameterList, precedence);

    

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
    setState(109);
    parameter();
    _ctx->stop = _input->LT(-1);
    setState(116);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 7, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<NonEmptyParameterListContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleNonEmptyParameterList);
        setState(111);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(112);
        match(routingAParser::T__9);
        setState(113);
        parameter(); 
      }
      setState(118);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 7, _ctx);
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
  enterRule(_localctx, 26, routingAParser::RuleParameter);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(126);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 8, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(119);
      match(routingAParser::ID);
      setState(120);
      match(routingAParser::T__5);
      setState(121);
      literal();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(122);
      match(routingAParser::ID);
      setState(123);
      match(routingAParser::T__6);
      setState(124);
      literal();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(125);
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

routingAParser::RoutingRuleLeftContext* routingAParser::RoutingRuleContext::routingRuleLeft() {
  return getRuleContext<routingAParser::RoutingRuleLeftContext>(0);
}

routingAParser::Bare_literalContext* routingAParser::RoutingRuleContext::bare_literal() {
  return getRuleContext<routingAParser::Bare_literalContext>(0);
}

routingAParser::OptFunctionPrototypeExpressionAndContext* routingAParser::RoutingRuleContext::optFunctionPrototypeExpressionAnd() {
  return getRuleContext<routingAParser::OptFunctionPrototypeExpressionAndContext>(0);
}

routingAParser::RoutingRuleListContext* routingAParser::RoutingRuleContext::routingRuleList() {
  return getRuleContext<routingAParser::RoutingRuleListContext>(0);
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
  enterRule(_localctx, 28, routingAParser::RuleRoutingRule);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(137);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 9, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(128);
      routingRuleLeft();
      setState(129);
      match(routingAParser::T__10);
      setState(130);
      bare_literal();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(132);
      optFunctionPrototypeExpressionAnd();
      setState(133);
      match(routingAParser::T__0);
      setState(134);
      routingRuleList();
      setState(135);
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

//----------------- RoutingRuleLeftContext ------------------------------------------------------------------

routingAParser::RoutingRuleLeftContext::RoutingRuleLeftContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::OptFunctionPrototypeExpressionAndContext* routingAParser::RoutingRuleLeftContext::optFunctionPrototypeExpressionAnd() {
  return getRuleContext<routingAParser::OptFunctionPrototypeExpressionAndContext>(0);
}

routingAParser::RoutingRuleLeftListContext* routingAParser::RoutingRuleLeftContext::routingRuleLeftList() {
  return getRuleContext<routingAParser::RoutingRuleLeftListContext>(0);
}

routingAParser::FunctionPrototypeExpressionContext* routingAParser::RoutingRuleLeftContext::functionPrototypeExpression() {
  return getRuleContext<routingAParser::FunctionPrototypeExpressionContext>(0);
}


size_t routingAParser::RoutingRuleLeftContext::getRuleIndex() const {
  return routingAParser::RuleRoutingRuleLeft;
}

void routingAParser::RoutingRuleLeftContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRoutingRuleLeft(this);
}

void routingAParser::RoutingRuleLeftContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRoutingRuleLeft(this);
}

routingAParser::RoutingRuleLeftContext* routingAParser::routingRuleLeft() {
  RoutingRuleLeftContext *_localctx = _tracker.createInstance<RoutingRuleLeftContext>(_ctx, getState());
  enterRule(_localctx, 30, routingAParser::RuleRoutingRuleLeft);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(147);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 10, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(139);
      optFunctionPrototypeExpressionAnd();
      setState(140);
      match(routingAParser::T__0);
      setState(141);
      routingRuleLeftList();
      setState(142);
      match(routingAParser::T__1);
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(144);
      optFunctionPrototypeExpressionAnd();
      setState(145);
      functionPrototypeExpression(0);
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

//----------------- RoutingRuleLeftListContext ------------------------------------------------------------------

routingAParser::RoutingRuleLeftListContext::RoutingRuleLeftListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::RoutingRuleLeftContext* routingAParser::RoutingRuleLeftListContext::routingRuleLeft() {
  return getRuleContext<routingAParser::RoutingRuleLeftContext>(0);
}

routingAParser::RoutingRuleLeftListContext* routingAParser::RoutingRuleLeftListContext::routingRuleLeftList() {
  return getRuleContext<routingAParser::RoutingRuleLeftListContext>(0);
}


size_t routingAParser::RoutingRuleLeftListContext::getRuleIndex() const {
  return routingAParser::RuleRoutingRuleLeftList;
}

void routingAParser::RoutingRuleLeftListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRoutingRuleLeftList(this);
}

void routingAParser::RoutingRuleLeftListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRoutingRuleLeftList(this);
}

routingAParser::RoutingRuleLeftListContext* routingAParser::routingRuleLeftList() {
  RoutingRuleLeftListContext *_localctx = _tracker.createInstance<RoutingRuleLeftListContext>(_ctx, getState());
  enterRule(_localctx, 32, routingAParser::RuleRoutingRuleLeftList);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(153);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 11, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(149);
      routingRuleLeft();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(150);
      routingRuleLeft();
      setState(151);
      routingRuleLeftList();
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
  enterRule(_localctx, 34, routingAParser::RuleOptFunctionPrototypeExpressionAnd);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(159);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 12, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(155);
      functionPrototypeExpression(0);
      setState(156);
      match(routingAParser::T__11);
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
  size_t startState = 36;
  enterRecursionRule(_localctx, 36, routingAParser::RuleFunctionPrototypeExpression, precedence);

    

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
    setState(162);
    functionPrototype();
    _ctx->stop = _input->LT(-1);
    setState(169);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        _localctx = _tracker.createInstance<FunctionPrototypeExpressionContext>(parentContext, parentState);
        pushNewRecursionContext(_localctx, startState, RuleFunctionPrototypeExpression);
        setState(164);

        if (!(precpred(_ctx, 1))) throw FailedPredicateException(this, "precpred(_ctx, 1)");
        setState(165);
        match(routingAParser::T__11);
        setState(166);
        functionPrototypeExpression(2); 
      }
      setState(171);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 13, _ctx);
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

routingAParser::DeclarationContext* routingAParser::RoutingRuleOrDeclarationListContext::declaration() {
  return getRuleContext<routingAParser::DeclarationContext>(0);
}

routingAParser::RoutingRuleOrDeclarationListContext* routingAParser::RoutingRuleOrDeclarationListContext::routingRuleOrDeclarationList() {
  return getRuleContext<routingAParser::RoutingRuleOrDeclarationListContext>(0);
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
  enterRule(_localctx, 38, routingAParser::RuleRoutingRuleOrDeclarationList);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(180);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 14, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(172);
      routingRule();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(173);
      declaration();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(174);
      routingRule();
      setState(175);
      routingRuleOrDeclarationList();
      break;
    }

    case 4: {
      enterOuterAlt(_localctx, 4);
      setState(177);
      declaration();
      setState(178);
      routingRuleOrDeclarationList();
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

//----------------- RoutingRuleListContext ------------------------------------------------------------------

routingAParser::RoutingRuleListContext::RoutingRuleListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

routingAParser::RoutingRuleContext* routingAParser::RoutingRuleListContext::routingRule() {
  return getRuleContext<routingAParser::RoutingRuleContext>(0);
}

routingAParser::RoutingRuleListContext* routingAParser::RoutingRuleListContext::routingRuleList() {
  return getRuleContext<routingAParser::RoutingRuleListContext>(0);
}


size_t routingAParser::RoutingRuleListContext::getRuleIndex() const {
  return routingAParser::RuleRoutingRuleList;
}

void routingAParser::RoutingRuleListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterRoutingRuleList(this);
}

void routingAParser::RoutingRuleListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<routingAListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitRoutingRuleList(this);
}

routingAParser::RoutingRuleListContext* routingAParser::routingRuleList() {
  RoutingRuleListContext *_localctx = _tracker.createInstance<RoutingRuleListContext>(_ctx, getState());
  enterRule(_localctx, 40, routingAParser::RuleRoutingRuleList);

#if __cplusplus > 201703L
  auto onExit = finally([=, this] {
#else
  auto onExit = finally([=] {
#endif
    exitRule();
  });
  try {
    setState(186);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 15, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(182);
      routingRule();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(183);
      routingRule();
      setState(184);
      routingRuleList();
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
    case 12: return nonEmptyParameterListSempred(antlrcpp::downCast<NonEmptyParameterListContext *>(context), predicateIndex);
    case 18: return functionPrototypeExpressionSempred(antlrcpp::downCast<FunctionPrototypeExpressionContext *>(context), predicateIndex);

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
