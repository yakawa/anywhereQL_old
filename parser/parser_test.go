package parser

import (
	"testing"

	//"github.com/anywhereQL/anywhereQL/ast"
	"github.com/anywhereQL/anywhereQL/lexer"
)

func TestDirectSelectStatment(t *testing.T) {
	input := "SELECT * FROM table1;"
	l := lexer.New(input)
	tokens, err := l.Tokenize()
	if err != nil {
		t.Errorf("Lexer Error: %s\n", err)
	}

	p := New(tokens)
	p.ParseDirectSelectStatement()
}
