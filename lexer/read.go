package lexer

import (
	"fmt"

	"github.com/anywhereQL/anywhereQL/token"
)

type LexerReadError struct {
	Msg string
	Ch  byte
}

func (err *LexerReadError) Error() string {
	return fmt.Sprintf("Error: %s : (%s)", err.Msg, string(err.Ch))
}

func (l *Lexer) readSpecialCharacterToken() (token.TokenType, string, error) {
	switch l.ch {
	case ' ':
		return token.SPACE_TOKEN, " ", nil
	case '"':
		return token.DOUBLE_QUOTE_TOKEN, "\"", nil
	case '%':
		return token.PERCENT_TOKEN, "%", nil
	case '&':
		return token.AMPERSAND_TOKEN, "&", nil
	case '(':
		return token.LEFT_PAREN_TOKEN, "(", nil
	case ')':
		return token.RIGHT_PAREN_TOKEN, ")", nil
	case '*':
		return token.ASTERISK_TOKEN, "*", nil
	case '+':
		return token.PLUS_SIGN_TOKEN, "+", nil
	case ',':
		return token.COMMA_TOKEN, ",", nil
	case '-':
		n := l.peekChar()
		if n == '-' {
			c, err := l.readComment()
			return token.COMMENT_TOKEN, c, err
		}
		return token.MINUS_SIGN_TOKEN, "-", nil
	case '.':
		if isDigit(l.peekChar()) {
			s, err := l.readNumber()
			return token.NUMBER_TOKEN, s, err
		} else if l.peekChar() == '.' {
			return token.DOUBLE_PERIOD_TOKEN, "..", nil
		}
		return token.PERIOD_TOKEN, ".", nil
	case '/':
		return token.SOLIDAS_TOKEN, "/", nil
	case ':':
		return token.COLON_TOKEN, ":", nil
	case ';':
		return token.SEMICOLON_TOKEN, ";", nil
	case '<':
		n := l.peekChar()
		if n == '>' {
			return token.NOT_EQUALS_OPERATOR_TOKEN, "<>", nil
		} else if n == '=' {
			return token.LESS_THAN_OR_EQUALS_OPERATOR_TOKEN, "<=", nil
		}
		return token.LESS_THAN_OPERATOR_TOKEN, "<", nil
	case '=':
		return token.EQUALS_OPERATOR_TOKEN, "=", nil
	case '>':
		if l.peekChar() == '=' {
			return token.GREATER_THAN_OPERATOR_TOKEN, ">=", nil
		}
		return token.GREATER_THAN_OPERATOR_TOKEN, ">", nil
	case '?':
		return token.QUESTION_MARK_TOKEN, "?", nil
	case '_':
		return token.UNDERSCORE_TOKEN, "_", nil
	case '|':
		if l.peekChar() == '|' {
			return token.CONCATENATION_OPERATOR_TOKEN, "||", nil
		}
		return token.VERTICAL_BAR_TOKEN, "|", nil
	case '[':
		return token.LEFT_BRACKET_TOKEN, "[", nil
	case ']':
		return token.RIGHT_BRACKET_TOKEN, "]", nil
	}
	return token.ILLEGAL_TOKEN, "", &LexerReadError{Msg: "Character is not Special Character", Ch: l.ch}
}

func (l *Lexer) readComment() (string, error) {
	/*
		<comment> ::= <comment introducer> [ <comment character>... ] <newline>
		<comment introducer> ::= <minus sign><minus sign>[<minus sign>...]
		<comment character> ::= <nonquote character> | <quote>
		<newline> ::= !! implementation-defined end-of-line indicator
	*/
	l.readChar()
	l.readChar()
	for l.ch == '-' {
		l.readChar()
	}
	l.skipSpace()
	pos := l.position
	for isNewline(l.ch) {
		l.readChar()
	}
	return l.input[pos : l.position+1], nil
}

func (l *Lexer) readNumber() (string, error) {
	/*
		<unsigned numeric literal> ::= <exact numeric literal> | <approximate numeric literal>
		<exact numeric literal> ::= <unsigned integer> [ <period> [ <unsigned integer> ] ] | <period> <unsigned integer>
		<approximate numeric literal> ::= <mantissa> E <exponent>
		<mantissa> ::= <exact numeric literal>
		<exponent> ::= <signed integer>
		<signed integer> ::= [ <sign> ] <unsigned integer>
		<sign> ::= <plus sign> | <minus sign>
		<unsigned integer> ::= <digit>...

		<bit string literal> ::= B <quote> [ <bit>... ] <quote> [ { <separator>... <quote> [ <bit>... ] <quote> }... ]
		<hex string literal> ::= X <quote> [ <hexit>... ] <quote> [ { <separator>... <quote> [ <hexit>... ] <quote> }... ]
		<bit> ::= 0 | 1
		<hexit> ::= <digit> | A | B | C | D | E | F | a | b | c | d | e | f
	*/
	pos := l.position
	if l.ch == 'B' {
	} else if l.ch == 'X' {
	} else {
		intPart := false
		if isDigit(l.ch) {
			intPart = true
			for isDigit(l.ch) {
				l.readChar()
			}
		}
		if l.ch == '.' {
			l.readChar()
			if !(intPart == false && isDigit(l.ch)) {
				return l.input[pos:l.position], nil
			}
			for isDigit(l.ch) {
				l.readChar()
			}
		} else if !(l.ch == 'E' || l.ch == 'e') {
			return l.input[pos:l.position], nil
		}
		n := l.peekChar()
		if (l.ch == 'E' || l.ch == 'e') && (isDigit(n) || n == '+' || n == '-') {
			l.readChar()
			if l.ch == '+' || l.ch == '-' {
				l.readChar()
			}
			for isDigit(l.ch) {
				l.readChar()
			}
		}
	}
	return l.input[pos:l.position], nil
}
