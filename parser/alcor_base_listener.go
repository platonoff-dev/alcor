// Code generated from Alcor.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Alcor
import "github.com/antlr4-go/antlr/v4"

// BaseAlcorListener is a complete listener for a parse tree produced by AlcorParser.
type BaseAlcorListener struct{}

var _ AlcorListener = &BaseAlcorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAlcorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAlcorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAlcorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAlcorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseAlcorListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseAlcorListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAlcorListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAlcorListener) ExitExpr(ctx *ExprContext) {}
