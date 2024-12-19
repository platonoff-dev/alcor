package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/platonoff-dev/alcor/parser"
)

type TreeShapeListener struct {
	*parser.BaseAlcorListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input := antlr.NewInputStream("1234")
	lexer := parser.NewAlcorLexer(input)

	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewAlcorParser(tokenStream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
}
