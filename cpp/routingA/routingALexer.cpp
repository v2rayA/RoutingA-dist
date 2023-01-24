
// Generated from routingA.g4 by ANTLR 4.11.1


#include "routingALexer.h"


using namespace antlr4;

using namespace routingA;


using namespace antlr4;

namespace {

struct RoutingALexerStaticData final {
  RoutingALexerStaticData(std::vector<std::string> ruleNames,
                          std::vector<std::string> channelNames,
                          std::vector<std::string> modeNames,
                          std::vector<std::string> literalNames,
                          std::vector<std::string> symbolicNames)
      : ruleNames(std::move(ruleNames)), channelNames(std::move(channelNames)),
        modeNames(std::move(modeNames)), literalNames(std::move(literalNames)),
        symbolicNames(std::move(symbolicNames)),
        vocabulary(this->literalNames, this->symbolicNames) {}

  RoutingALexerStaticData(const RoutingALexerStaticData&) = delete;
  RoutingALexerStaticData(RoutingALexerStaticData&&) = delete;
  RoutingALexerStaticData& operator=(const RoutingALexerStaticData&) = delete;
  RoutingALexerStaticData& operator=(RoutingALexerStaticData&&) = delete;

  std::vector<antlr4::dfa::DFA> decisionToDFA;
  antlr4::atn::PredictionContextCache sharedContextCache;
  const std::vector<std::string> ruleNames;
  const std::vector<std::string> channelNames;
  const std::vector<std::string> modeNames;
  const std::vector<std::string> literalNames;
  const std::vector<std::string> symbolicNames;
  const antlr4::dfa::Vocabulary vocabulary;
  antlr4::atn::SerializedATNView serializedATN;
  std::unique_ptr<antlr4::atn::ATN> atn;
};

::antlr4::internal::OnceFlag routingalexerLexerOnceFlag;
RoutingALexerStaticData *routingalexerLexerStaticData = nullptr;

void routingalexerLexerInitialize() {
  assert(routingalexerLexerStaticData == nullptr);
  auto staticData = std::make_unique<RoutingALexerStaticData>(
    std::vector<std::string>{
      "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
      "T__9", "T__10", "T__11", "SAFE_ID_HEAD_CHAR", "SAFE_NONID_HEAD_CHAR", 
      "SAFE_INTERMEDIATE_CHAR", "SAFE_CHAR", "DOUBLE_QUOTE_STRING", "SINGLE_QUOTE_STRING", 
      "WHITESPACE", "COMMENT_BLOCK", "COMMENT_LINE_SHARP", "ID", "NON_ID", 
      "QUOTE_STRING"
    },
    std::vector<std::string>{
      "DEFAULT_TOKEN_CHANNEL", "HIDDEN"
    },
    std::vector<std::string>{
      "DEFAULT_MODE"
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
  	4,0,18,166,6,-1,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
  	6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,
  	7,14,2,15,7,15,2,16,7,16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,
  	7,21,2,22,7,22,2,23,7,23,1,0,1,0,1,1,1,1,1,2,1,2,1,3,1,3,1,3,1,4,1,4,
  	1,4,1,5,1,5,1,6,1,6,1,7,1,7,1,8,1,8,1,9,1,9,1,10,1,10,1,10,1,11,1,11,
  	1,11,1,12,1,12,1,13,1,13,1,14,1,14,1,15,1,15,1,15,3,15,87,8,15,1,16,1,
  	16,1,16,1,16,5,16,93,8,16,10,16,12,16,96,9,16,1,16,1,16,1,17,1,17,1,17,
  	1,17,5,17,104,8,17,10,17,12,17,107,9,17,1,17,1,17,1,18,4,18,112,8,18,
  	11,18,12,18,113,1,18,1,18,1,19,1,19,1,19,1,19,5,19,122,8,19,10,19,12,
  	19,125,9,19,1,19,1,19,1,19,1,19,1,19,1,20,1,20,5,20,134,8,20,10,20,12,
  	20,137,9,20,1,20,4,20,140,8,20,11,20,12,20,141,1,20,3,20,145,8,20,1,20,
  	1,20,1,21,1,21,5,21,151,8,21,10,21,12,21,154,9,21,1,22,1,22,5,22,158,
  	8,22,10,22,12,22,161,9,22,1,23,1,23,3,23,165,8,23,4,94,105,123,135,0,
  	24,1,1,3,2,5,3,7,4,9,5,11,6,13,7,15,8,17,9,19,10,21,11,23,12,25,0,27,
  	0,29,0,31,0,33,0,35,0,37,13,39,14,41,15,43,16,45,17,47,18,1,0,5,3,0,65,
  	90,95,95,97,122,5,0,33,33,42,43,45,57,92,92,94,94,2,0,36,36,64,64,3,0,
  	9,10,13,13,32,32,2,0,10,10,13,13,173,0,1,1,0,0,0,0,3,1,0,0,0,0,5,1,0,
  	0,0,0,7,1,0,0,0,0,9,1,0,0,0,0,11,1,0,0,0,0,13,1,0,0,0,0,15,1,0,0,0,0,
  	17,1,0,0,0,0,19,1,0,0,0,0,21,1,0,0,0,0,23,1,0,0,0,0,37,1,0,0,0,0,39,1,
  	0,0,0,0,41,1,0,0,0,0,43,1,0,0,0,0,45,1,0,0,0,0,47,1,0,0,0,1,49,1,0,0,
  	0,3,51,1,0,0,0,5,53,1,0,0,0,7,55,1,0,0,0,9,58,1,0,0,0,11,61,1,0,0,0,13,
  	63,1,0,0,0,15,65,1,0,0,0,17,67,1,0,0,0,19,69,1,0,0,0,21,71,1,0,0,0,23,
  	74,1,0,0,0,25,77,1,0,0,0,27,79,1,0,0,0,29,81,1,0,0,0,31,86,1,0,0,0,33,
  	88,1,0,0,0,35,99,1,0,0,0,37,111,1,0,0,0,39,117,1,0,0,0,41,131,1,0,0,0,
  	43,148,1,0,0,0,45,155,1,0,0,0,47,164,1,0,0,0,49,50,5,123,0,0,50,2,1,0,
  	0,0,51,52,5,125,0,0,52,4,1,0,0,0,53,54,5,64,0,0,54,6,1,0,0,0,55,56,5,
  	61,0,0,56,57,5,61,0,0,57,8,1,0,0,0,58,59,5,33,0,0,59,60,5,61,0,0,60,10,
  	1,0,0,0,61,62,5,58,0,0,62,12,1,0,0,0,63,64,5,61,0,0,64,14,1,0,0,0,65,
  	66,5,40,0,0,66,16,1,0,0,0,67,68,5,41,0,0,68,18,1,0,0,0,69,70,5,44,0,0,
  	70,20,1,0,0,0,71,72,5,45,0,0,72,73,5,62,0,0,73,22,1,0,0,0,74,75,5,38,
  	0,0,75,76,5,38,0,0,76,24,1,0,0,0,77,78,7,0,0,0,78,26,1,0,0,0,79,80,7,
  	1,0,0,80,28,1,0,0,0,81,82,7,2,0,0,82,30,1,0,0,0,83,87,3,25,12,0,84,87,
  	3,27,13,0,85,87,3,29,14,0,86,83,1,0,0,0,86,84,1,0,0,0,86,85,1,0,0,0,87,
  	32,1,0,0,0,88,94,5,34,0,0,89,90,5,92,0,0,90,93,5,34,0,0,91,93,9,0,0,0,
  	92,89,1,0,0,0,92,91,1,0,0,0,93,96,1,0,0,0,94,95,1,0,0,0,94,92,1,0,0,0,
  	95,97,1,0,0,0,96,94,1,0,0,0,97,98,5,34,0,0,98,34,1,0,0,0,99,105,5,39,
  	0,0,100,101,5,92,0,0,101,104,5,39,0,0,102,104,9,0,0,0,103,100,1,0,0,0,
  	103,102,1,0,0,0,104,107,1,0,0,0,105,106,1,0,0,0,105,103,1,0,0,0,106,108,
  	1,0,0,0,107,105,1,0,0,0,108,109,5,39,0,0,109,36,1,0,0,0,110,112,7,3,0,
  	0,111,110,1,0,0,0,112,113,1,0,0,0,113,111,1,0,0,0,113,114,1,0,0,0,114,
  	115,1,0,0,0,115,116,6,18,0,0,116,38,1,0,0,0,117,118,5,47,0,0,118,119,
  	5,42,0,0,119,123,1,0,0,0,120,122,9,0,0,0,121,120,1,0,0,0,122,125,1,0,
  	0,0,123,124,1,0,0,0,123,121,1,0,0,0,124,126,1,0,0,0,125,123,1,0,0,0,126,
  	127,5,42,0,0,127,128,5,47,0,0,128,129,1,0,0,0,129,130,6,19,0,0,130,40,
  	1,0,0,0,131,135,5,35,0,0,132,134,9,0,0,0,133,132,1,0,0,0,134,137,1,0,
  	0,0,135,136,1,0,0,0,135,133,1,0,0,0,136,144,1,0,0,0,137,135,1,0,0,0,138,
  	140,7,4,0,0,139,138,1,0,0,0,140,141,1,0,0,0,141,139,1,0,0,0,141,142,1,
  	0,0,0,142,145,1,0,0,0,143,145,5,0,0,1,144,139,1,0,0,0,144,143,1,0,0,0,
  	145,146,1,0,0,0,146,147,6,20,0,0,147,42,1,0,0,0,148,152,3,25,12,0,149,
  	151,3,31,15,0,150,149,1,0,0,0,151,154,1,0,0,0,152,150,1,0,0,0,152,153,
  	1,0,0,0,153,44,1,0,0,0,154,152,1,0,0,0,155,159,3,27,13,0,156,158,3,31,
  	15,0,157,156,1,0,0,0,158,161,1,0,0,0,159,157,1,0,0,0,159,160,1,0,0,0,
  	160,46,1,0,0,0,161,159,1,0,0,0,162,165,3,33,16,0,163,165,3,35,17,0,164,
  	162,1,0,0,0,164,163,1,0,0,0,165,48,1,0,0,0,14,0,86,92,94,103,105,113,
  	123,135,141,144,152,159,164,1,6,0,0
  };
  staticData->serializedATN = antlr4::atn::SerializedATNView(serializedATNSegment, sizeof(serializedATNSegment) / sizeof(serializedATNSegment[0]));

  antlr4::atn::ATNDeserializer deserializer;
  staticData->atn = deserializer.deserialize(staticData->serializedATN);

  const size_t count = staticData->atn->getNumberOfDecisions();
  staticData->decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    staticData->decisionToDFA.emplace_back(staticData->atn->getDecisionState(i), i);
  }
  routingalexerLexerStaticData = staticData.release();
}

}

routingALexer::routingALexer(CharStream *input) : Lexer(input) {
  routingALexer::initialize();
  _interpreter = new atn::LexerATNSimulator(this, *routingalexerLexerStaticData->atn, routingalexerLexerStaticData->decisionToDFA, routingalexerLexerStaticData->sharedContextCache);
}

routingALexer::~routingALexer() {
  delete _interpreter;
}

std::string routingALexer::getGrammarFileName() const {
  return "routingA.g4";
}

const std::vector<std::string>& routingALexer::getRuleNames() const {
  return routingalexerLexerStaticData->ruleNames;
}

const std::vector<std::string>& routingALexer::getChannelNames() const {
  return routingalexerLexerStaticData->channelNames;
}

const std::vector<std::string>& routingALexer::getModeNames() const {
  return routingalexerLexerStaticData->modeNames;
}

const dfa::Vocabulary& routingALexer::getVocabulary() const {
  return routingalexerLexerStaticData->vocabulary;
}

antlr4::atn::SerializedATNView routingALexer::getSerializedATN() const {
  return routingalexerLexerStaticData->serializedATN;
}

const atn::ATN& routingALexer::getATN() const {
  return *routingalexerLexerStaticData->atn;
}




void routingALexer::initialize() {
  ::antlr4::internal::call_once(routingalexerLexerOnceFlag, routingalexerLexerInitialize);
}
