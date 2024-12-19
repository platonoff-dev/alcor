// Code generated from parser/AlcorParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AlcorParser
import "github.com/antlr4-go/antlr/v4"

// BaseAlcorParserListener is a complete listener for a parse tree produced by AlcorParser.
type BaseAlcorParserListener struct{}

var _ AlcorParserListener = &BaseAlcorParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAlcorParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAlcorParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAlcorParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAlcorParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNumberLit is called when production numberLit is entered.
func (s *BaseAlcorParserListener) EnterNumberLit(ctx *NumberLitContext) {}

// ExitNumberLit is called when production numberLit is exited.
func (s *BaseAlcorParserListener) ExitNumberLit(ctx *NumberLitContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAlcorParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAlcorParserListener) ExitExpr(ctx *ExprContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseAlcorParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseAlcorParserListener) ExitProg(ctx *ProgContext) {}
