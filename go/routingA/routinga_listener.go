// Code generated from routingA.g4 by ANTLR 4.10.1. DO NOT EDIT.

package routingA // routingA
import "github.com/antlr/antlr4/runtime/Go/antlr"

// routingAListener is a complete listener for a parse tree produced by routingAParser.
type routingAListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterBare_literal is called when entering the bare_literal production.
	EnterBare_literal(c *Bare_literalContext)

	// EnterQuote_literal is called when entering the quote_literal production.
	EnterQuote_literal(c *Quote_literalContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInput is called when entering the input production.
	EnterInput(c *InputContext)

	// EnterProgramStructureBlcok is called when entering the programStructureBlcok production.
	EnterProgramStructureBlcok(c *ProgramStructureBlcokContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOptConditionStatement is called when entering the optConditionStatement production.
	EnterOptConditionStatement(c *OptConditionStatementContext)

	// EnterOptFunctionPrototypeExpressionAnd is called when entering the optFunctionPrototypeExpressionAnd production.
	EnterOptFunctionPrototypeExpressionAnd(c *OptFunctionPrototypeExpressionAndContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterFunctionPrototype is called when entering the functionPrototype production.
	EnterFunctionPrototype(c *FunctionPrototypeContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterNonEmptyParameterList is called when entering the nonEmptyParameterList production.
	EnterNonEmptyParameterList(c *NonEmptyParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterRoutingRule is called when entering the routingRule production.
	EnterRoutingRule(c *RoutingRuleContext)

	// EnterFunctionPrototypeExpression is called when entering the functionPrototypeExpression production.
	EnterFunctionPrototypeExpression(c *FunctionPrototypeExpressionContext)

	// EnterFunctionPrototypeExpressionList is called when entering the functionPrototypeExpressionList production.
	EnterFunctionPrototypeExpressionList(c *FunctionPrototypeExpressionListContext)

	// EnterRoutingRuleOrDeclarationList is called when entering the routingRuleOrDeclarationList production.
	EnterRoutingRuleOrDeclarationList(c *RoutingRuleOrDeclarationListContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitBare_literal is called when exiting the bare_literal production.
	ExitBare_literal(c *Bare_literalContext)

	// ExitQuote_literal is called when exiting the quote_literal production.
	ExitQuote_literal(c *Quote_literalContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInput is called when exiting the input production.
	ExitInput(c *InputContext)

	// ExitProgramStructureBlcok is called when exiting the programStructureBlcok production.
	ExitProgramStructureBlcok(c *ProgramStructureBlcokContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOptConditionStatement is called when exiting the optConditionStatement production.
	ExitOptConditionStatement(c *OptConditionStatementContext)

	// ExitOptFunctionPrototypeExpressionAnd is called when exiting the optFunctionPrototypeExpressionAnd production.
	ExitOptFunctionPrototypeExpressionAnd(c *OptFunctionPrototypeExpressionAndContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitFunctionPrototype is called when exiting the functionPrototype production.
	ExitFunctionPrototype(c *FunctionPrototypeContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitNonEmptyParameterList is called when exiting the nonEmptyParameterList production.
	ExitNonEmptyParameterList(c *NonEmptyParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitRoutingRule is called when exiting the routingRule production.
	ExitRoutingRule(c *RoutingRuleContext)

	// ExitFunctionPrototypeExpression is called when exiting the functionPrototypeExpression production.
	ExitFunctionPrototypeExpression(c *FunctionPrototypeExpressionContext)

	// ExitFunctionPrototypeExpressionList is called when exiting the functionPrototypeExpressionList production.
	ExitFunctionPrototypeExpressionList(c *FunctionPrototypeExpressionListContext)

	// ExitRoutingRuleOrDeclarationList is called when exiting the routingRuleOrDeclarationList production.
	ExitRoutingRuleOrDeclarationList(c *RoutingRuleOrDeclarationListContext)
}
