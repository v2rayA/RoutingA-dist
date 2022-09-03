
// Generated from routingA.g4 by ANTLR 4.10.1

#pragma once


#include "antlr4-runtime.h"
#include "routingAParser.h"


namespace routingA {

/**
 * This interface defines an abstract listener for a parse tree produced by routingAParser.
 */
class  routingAListener : public antlr4::tree::ParseTreeListener {
public:

  virtual void enterStart(routingAParser::StartContext *ctx) = 0;
  virtual void exitStart(routingAParser::StartContext *ctx) = 0;

  virtual void enterBare_literal(routingAParser::Bare_literalContext *ctx) = 0;
  virtual void exitBare_literal(routingAParser::Bare_literalContext *ctx) = 0;

  virtual void enterQuote_literal(routingAParser::Quote_literalContext *ctx) = 0;
  virtual void exitQuote_literal(routingAParser::Quote_literalContext *ctx) = 0;

  virtual void enterLiteral(routingAParser::LiteralContext *ctx) = 0;
  virtual void exitLiteral(routingAParser::LiteralContext *ctx) = 0;

  virtual void enterInput(routingAParser::InputContext *ctx) = 0;
  virtual void exitInput(routingAParser::InputContext *ctx) = 0;

  virtual void enterProgramStructureBlcok(routingAParser::ProgramStructureBlcokContext *ctx) = 0;
  virtual void exitProgramStructureBlcok(routingAParser::ProgramStructureBlcokContext *ctx) = 0;

  virtual void enterExpression(routingAParser::ExpressionContext *ctx) = 0;
  virtual void exitExpression(routingAParser::ExpressionContext *ctx) = 0;

  virtual void enterOptConditionStatement(routingAParser::OptConditionStatementContext *ctx) = 0;
  virtual void exitOptConditionStatement(routingAParser::OptConditionStatementContext *ctx) = 0;

  virtual void enterOptFunctionPrototypeExpressionAnd(routingAParser::OptFunctionPrototypeExpressionAndContext *ctx) = 0;
  virtual void exitOptFunctionPrototypeExpressionAnd(routingAParser::OptFunctionPrototypeExpressionAndContext *ctx) = 0;

  virtual void enterDeclaration(routingAParser::DeclarationContext *ctx) = 0;
  virtual void exitDeclaration(routingAParser::DeclarationContext *ctx) = 0;

  virtual void enterAssignmentExpression(routingAParser::AssignmentExpressionContext *ctx) = 0;
  virtual void exitAssignmentExpression(routingAParser::AssignmentExpressionContext *ctx) = 0;

  virtual void enterFunctionPrototype(routingAParser::FunctionPrototypeContext *ctx) = 0;
  virtual void exitFunctionPrototype(routingAParser::FunctionPrototypeContext *ctx) = 0;

  virtual void enterParameterList(routingAParser::ParameterListContext *ctx) = 0;
  virtual void exitParameterList(routingAParser::ParameterListContext *ctx) = 0;

  virtual void enterNonEmptyParameterList(routingAParser::NonEmptyParameterListContext *ctx) = 0;
  virtual void exitNonEmptyParameterList(routingAParser::NonEmptyParameterListContext *ctx) = 0;

  virtual void enterParameter(routingAParser::ParameterContext *ctx) = 0;
  virtual void exitParameter(routingAParser::ParameterContext *ctx) = 0;

  virtual void enterRoutingRule(routingAParser::RoutingRuleContext *ctx) = 0;
  virtual void exitRoutingRule(routingAParser::RoutingRuleContext *ctx) = 0;

  virtual void enterFunctionPrototypeExpression(routingAParser::FunctionPrototypeExpressionContext *ctx) = 0;
  virtual void exitFunctionPrototypeExpression(routingAParser::FunctionPrototypeExpressionContext *ctx) = 0;

  virtual void enterFunctionPrototypeExpressionList(routingAParser::FunctionPrototypeExpressionListContext *ctx) = 0;
  virtual void exitFunctionPrototypeExpressionList(routingAParser::FunctionPrototypeExpressionListContext *ctx) = 0;

  virtual void enterRoutingRuleOrDeclarationList(routingAParser::RoutingRuleOrDeclarationListContext *ctx) = 0;
  virtual void exitRoutingRuleOrDeclarationList(routingAParser::RoutingRuleOrDeclarationListContext *ctx) = 0;


};

}  // namespace routingA
