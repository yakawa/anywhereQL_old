package lexer

import (
	"github.com/anywhereQL/anywhereQL/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(i string) *Lexer {
	l := &Lexer{
		input:        i,
		position:     0,
		readPosition: 0,
	}

	return l
}

func (l *Lexer) Tokenize() ([]token.Token, error) {
	t := make([]token.Token, 100)

	for {
		tok, err := l.nextToken()
		if err != nil {
			e := make([]token.Token, 1)
			return e, err
		}
		t = append(t, tok)
		if tok.Type == token.EOF_TOKEN {
			break
		}
	}

	return t, nil
}

func (l *Lexer) nextToken() (token.Token, error) {
	t := token.Token{Type: token.ILLEGAL_TOKEN, Literal: ""}
	if l.ch == 0 {
		return token.Token{Type: token.EOF_TOKEN, Literal: ""}, nil
	} else if isSQLSpecialCharacter(l.ch) {
		tok, s, err := l.readSpecialCharacterToken()
		if err != nil {
			return t, err
		}
		t.Type = tok
		t.Literal = s
	} else if isDigit(l.ch) {
		s, err := l.readNumber()
		if err != nil {
			return t, err
		}
		t.Type = token.NUMBER_TOKEN
		t.Literal = s
	}
	l.readChar()
	return t, nil
}

func (l *Lexer) readChar() {
	if l.readPosition > len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipSpace() {
	for l.ch == ' ' {
		l.readChar()
	}
}
