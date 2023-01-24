
// Generated from routingA.g4 by ANTLR 4.11.1

#pragma once


#include "antlr4-runtime.h"
#include "routingAListener.h"


namespace routingA {

/**
 * This class provides an empty implementation of routingAListener,
 * which can be extended to create a listener which only needs to handle a subset
 * of the available methods.
 */
class  routingABaseListener : public routingAListener {
public:

  virtual void enterStart(routingAParser::StartContext * /*ctx*/) override { }
  virtual void exitStart(routingAParser::StartContext * /*ctx*/) override { }

  virtual void enterBare_literal(routingAParser::Bare_literalContext * /*ctx*/) override { }
  virtual void exitBare_literal(routingAParser::Bare_literalContext * /*ctx*/) override { }

  virtual void enterQuote_literal(routingAParser::Quote_literalContext * /*ctx*/) override { }
  virtual void exitQuote_literal(routingAParser::Quote_literalContext * /*ctx*/) override { }

  virtual void enterLiteral(routingAParser::LiteralContext * /*ctx*/) override { }
  virtual void exitLiteral(routingAParser::LiteralContext * /*ctx*/) override { }

  virtual void enterInput(routingAParser::InputContext * /*ctx*/) override { }
  virtual void exitInput(routingAParser::InputContext * /*ctx*/) override { }

  virtual void enterProgramStructureBlcok(routingAParser::ProgramStructureBlcokContext * /*ctx*/) override { }
  virtual void exitProgramStructureBlcok(routingAParser::ProgramStructureBlcokContext * /*ctx*/) override { }

  virtual void enterExpression(routingAParser::ExpressionContext * /*ctx*/) override { }
  virtual void exitExpression(routingAParser::ExpressionContext * /*ctx*/) override { }

  virtual void enterOptConditionStatement(routingAParser::OptConditionStatementContext * /*ctx*/) override { }
  virtual void exitOptConditionStatement(routingAParser::OptConditionStatementContext * /*ctx*/) override { }

  virtual void enterDeclaration(routingAParser::DeclarationContext * /*ctx*/) override { }
  virtual void exitDeclaration(routingAParser::DeclarationContext * /*ctx*/) override { }

  virtual void enterAssignmentExpression(routingAParser::AssignmentExpressionContext * /*ctx*/) override { }
  virtual void exitAssignmentExpression(routingAParser::AssignmentExpressionContext * /*ctx*/) override { }

  virtual void enterFunctionPrototype(routingAParser::FunctionPrototypeContext * /*ctx*/) override { }
  virtual void exitFunctionPrototype(routingAParser::FunctionPrototypeContext * /*ctx*/) override { }

  virtual void enterOptParameterList(routingAParser::OptParameterListContext * /*ctx*/) override { }
  virtual void exitOptParameterList(routingAParser::OptParameterListContext * /*ctx*/) override { }

  virtual void enterNonEmptyParameterList(routingAParser::NonEmptyParameterListContext * /*ctx*/) override { }
  virtual void exitNonEmptyParameterList(routingAParser::NonEmptyParameterListContext * /*ctx*/) override { }

  virtual void enterParameter(routingAParser::ParameterContext * /*ctx*/) override { }
  virtual void exitParameter(routingAParser::ParameterContext * /*ctx*/) override { }

  virtual void enterRoutingRule(routingAParser::RoutingRuleContext * /*ctx*/) override { }
  virtual void exitRoutingRule(routingAParser::RoutingRuleContext * /*ctx*/) override { }

  virtual void enterRoutingRuleLeft(routingAParser::RoutingRuleLeftContext * /*ctx*/) override { }
  virtual void exitRoutingRuleLeft(routingAParser::RoutingRuleLeftContext * /*ctx*/) override { }

  virtual void enterRoutingRuleLeftList(routingAParser::RoutingRuleLeftListContext * /*ctx*/) override { }
  virtual void exitRoutingRuleLeftList(routingAParser::RoutingRuleLeftListContext * /*ctx*/) override { }

  virtual void enterOptFunctionPrototypeExpressionAnd(routingAParser::OptFunctionPrototypeExpressionAndContext * /*ctx*/) override { }
  virtual void exitOptFunctionPrototypeExpressionAnd(routingAParser::OptFunctionPrototypeExpressionAndContext * /*ctx*/) override { }

  virtual void enterFunctionPrototypeExpression(routingAParser::FunctionPrototypeExpressionContext * /*ctx*/) override { }
  virtual void exitFunctionPrototypeExpression(routingAParser::FunctionPrototypeExpressionContext * /*ctx*/) override { }

  virtual void enterRoutingRuleOrDeclarationList(routingAParser::RoutingRuleOrDeclarationListContext * /*ctx*/) override { }
  virtual void exitRoutingRuleOrDeclarationList(routingAParser::RoutingRuleOrDeclarationListContext * /*ctx*/) override { }

  virtual void enterRoutingRuleList(routingAParser::RoutingRuleListContext * /*ctx*/) override { }
  virtual void exitRoutingRuleList(routingAParser::RoutingRuleListContext * /*ctx*/) override { }


  virtual void enterEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void exitEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void visitTerminal(antlr4::tree::TerminalNode * /*node*/) override { }
  virtual void visitErrorNode(antlr4::tree::ErrorNode * /*node*/) override { }

};

}  // namespace routingA
