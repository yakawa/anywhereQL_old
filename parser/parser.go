package parser

import (
	"github.com/anywhereQL/anywhereQL/ast"
	"github.com/anywhereQL/anywhereQL/token"
)

type Parser struct {
	tokens []token.Token

	curTokenPos  int
	peekTokenPos int
}

func New(tokens []token.Token) *Parser {
	p := &Parser{tokens: tokens}
	p.curTokenPos = 0
	p.peekTokenPos = 1
	return p
}

func (p *Parser) ParseDirectSelectStatement() *ast.DirectSelectStatement {
	return nil
}
