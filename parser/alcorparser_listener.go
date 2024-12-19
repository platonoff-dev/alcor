// Code generated from parser/AlcorParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // AlcorParser
import "github.com/antlr4-go/antlr/v4"

// AlcorParserListener is a complete listener for a parse tree produced by AlcorParser.
type AlcorParserListener interface {
	antlr.ParseTreeListener

	// EnterNumberLit is called when entering the numberLit production.
	EnterNumberLit(c *NumberLitContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// ExitNumberLit is called when exiting the numberLit production.
	ExitNumberLit(c *NumberLitContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)
}
