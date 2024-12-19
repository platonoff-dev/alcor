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
	input := antlr.NewInputStream("1 + 0xA")
	lexer := parser.NewAlcorLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewAlcorParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
