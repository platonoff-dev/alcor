// Code generated from Alcor.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Alcor
import "github.com/antlr4-go/antlr/v4"

// AlcorListener is a complete listener for a parse tree produced by AlcorParser.
type AlcorListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}