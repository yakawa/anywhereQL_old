package parser

import (
	"testing"

	//"github.com/anywhereQL/anywhereQL/ast"
	"github.com/anywhereQL/anywhereQL/lexer"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		input string
		//expected []ast.Node
	}{
		{"SELECT *"},
		{"SELECT t"},
		{"SELECT t.*"},
		{"SELECT t.a"},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		tokens, _ := l.Tokenize()
		p := New(tokens)
		p.GetSQL()
	}
}
