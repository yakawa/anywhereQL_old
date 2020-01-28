package lexer

import (
	"fmt"
	"strings"

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
	case '\'':
		if l.peekChar() == 0 {
			return token.QUOTE_TOKEN, "'", nil
		}
		tok, s, err := l.readQuotedToken()
		return tok, s, err
	case '"':
		if l.peekChar() == 0 {
			return token.DOUBLE_QUOTE_TOKEN, "\"", nil
		}
		tok, s, err := l.readDoubleQuotedToken()
		return tok, s, err
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
		if l.isPeekDigit() {
			s, err := l.readNumber()
			return token.NUMBER_TOKEN, s, err
		} else if l.peekChar() == '.' {
			l.readChar()
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
			l.readChar()
			return token.NOT_EQUALS_OPERATOR_TOKEN, "<>", nil
		} else if n == '=' {
			l.readChar()
			return token.LESS_THAN_OR_EQUALS_OPERATOR_TOKEN, "<=", nil
		}
		return token.LESS_THAN_OPERATOR_TOKEN, "<", nil
	case '=':
		return token.EQUALS_OPERATOR_TOKEN, "=", nil
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			return token.GREATER_THAN_OR_EQUALS_OPERATOR_TOKEN, ">=", nil
		}
		return token.GREATER_THAN_OPERATOR_TOKEN, ">", nil
	case '?':
		return token.QUESTION_MARK_TOKEN, "?", nil
	case '_':
		if l.isPeekSeparator() {
			return token.UNDERSCORE_TOKEN, "_", nil
		}
		tok, s, err := l.readIdentifier()
		return tok, s, err
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
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
	for !l.isNewline() {
		l.readChar()
	}
	return l.input[pos:l.position], nil
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
		l.readChar()
		for true {
			if l.ch == '\'' {
				l.readChar()
				for l.isBit() {
					l.readChar()
				}
				if l.ch != '\'' {
					return "", &LexerReadError{Msg: "Illegal Bit String", Ch: 0}
				} else {
					l.readChar()
				}
				if !l.isSeparator() {
					break
				}
				l.skipSeparator()
			} else {
				break
			}
		}
	} else if l.ch == 'X' {
		l.readChar()
		for true {
			if l.ch == '\'' {
				l.readChar()
				for l.isHexit() {
					l.readChar()
				}
				if l.ch != '\'' {
					return "", &LexerReadError{Msg: "Illegal Bit String", Ch: 0}
				} else {
					l.readChar()
				}
				if !l.isSeparator() {
					break
				}
				l.skipSeparator()
			} else {
				break
			}
		}
	} else {
		intPart := false
		if l.isDigit() {
			intPart = true
			for l.isDigit() {
				l.readChar()
			}
		}
		if l.ch == '.' {
			l.readChar()
			decimalPart := false
			for l.isDigit() {
				decimalPart = true
				l.readChar()
			}
			if intPart == false && decimalPart == false {
				return "", &LexerReadError{Msg: "Illegal Number Format", Ch: 0}
			}
		} else if !(l.ch == 'E' || l.ch == 'e') {
			return l.input[pos:l.position], nil
		}
		n := l.peekChar()
		if (l.ch == 'E' || l.ch == 'e') && (l.isPeekDigit() || n == '+' || n == '-') {
			l.readChar()
			if l.ch == '+' || l.ch == '-' {
				l.readChar()
				if !l.isDigit() {
					return "", &LexerReadError{Msg: "Illegal Number Format", Ch: 0}
				}
			}
			for l.isDigit() {
				l.readChar()
			}
		}
	}
	return l.input[pos:l.position], nil
}

func (l *Lexer) readQuotedToken() (token.TokenType, string, error) {
	literal := ""
	for true {
		if l.ch != '\'' {
			break
		}
		pos := l.position
		for true {
			l.readChar()
			if !l.isNonquotedCharacter() {
				break
			}
		}
		if l.ch == '\'' {
			literal += strings.TrimSpace(l.input[pos:l.position])
		} else {
			return token.ILLEGAL_TOKEN, "", &LexerReadError{Msg: "Illegal Quoted string", Ch: 0}
		}
		l.skipSpace()
	}
	return token.IDENTIFIER_TOKEN, literal, nil
}

func (l *Lexer) readDoubleQuotedToken() (token.TokenType, string, error) {
	literal := ""
	for true {
		if l.ch != '"' {
			break
		}
		pos := l.position
		for true {
			l.readChar()
			if !l.isNondoublequotedCharacter() {
				break
			}
		}
		if l.ch == '"' {
			literal += strings.TrimSpace(l.input[pos:l.position])
		} else {
			return token.ILLEGAL_TOKEN, "", &LexerReadError{Msg: "Illegal Double Quoted string", Ch: 0}
		}
		l.skipSpace()
	}
	return token.IDENTIFIER_TOKEN, literal, nil
}

func (l *Lexer) readIdentifier() (token.TokenType, string, error) {
	pos := l.position
	var tok token.TokenType
	var s string
	var err error
	if l.ch == '\'' {
		tok, s, err = l.readQuotedToken()
		tok = token.LookupKeyword(s)
	} else if l.ch == '"' {
		tok, s, err = l.readDoubleQuotedToken()
		tok = token.LookupKeyword(s)
	} else {
		for true {
			if l.isSeparator() {
				break
			}
			if l.ch == '"' {
				l.readDoubleQuotedToken()
				break
			}
			if l.ch == '\'' {
				l.readQuotedToken()
				break
			}
			l.readChar()
		}
		c := strings.TrimSpace(l.input[pos:l.position])
		t := token.LookupKeyword(c)
		return t, c, nil
	}
	return tok, s, err
}
