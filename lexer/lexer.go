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
	l.readChar()
	return l
}

func (l *Lexer) Tokenize() ([]token.Token, error) {
	t := make([]token.Token, 0, 100)

	for {
		l.skipSpace()
		tok, err := l.nextToken()
		if err != nil {
			e := make([]token.Token, 0, 100)
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
	} else if l.isSQLSpecialCharacter() {
		tok, s, err := l.readSpecialCharacterToken()
		if err != nil {
			return t, err
		}
		t.Type = tok
		t.Literal = s
	} else if l.isDigit() {
		s, err := l.readNumber()
		if err != nil {
			return t, err
		}
		t.Type = token.NUMBER_TOKEN
		t.Literal = s
	} else {
		tok, s, err := l.readIdentifier()
		if err != nil {
			return t, err
		}
		t.Type = tok
		t.Literal = s
	}
	return t, nil
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
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

func (l *Lexer) peek2Char() byte {
	if l.readPosition+1 >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition+1]
}

func (l *Lexer) skipSpace() {
	for l.isSeparator() {
		if l.ch == 0 {
			return
		}
		if l.ch == '-' {
			return
		}
		l.readChar()
	}
}

func (l *Lexer) skipSeparator() {
	for l.isSeparator() {
		if l.ch == 0 {
			return
		}
		if l.ch == '-' && l.peekChar() == '-' {
			l.readComment()
		} else {
			l.readChar()
		}
	}
	return
}

func (l *Lexer) setPosition(p, r int, c byte) {
	l.position = p
	l.readPosition = r
	l.ch = c
}
