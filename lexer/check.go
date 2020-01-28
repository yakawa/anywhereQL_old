package lexer

var (
	digit                       = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	simpleLatinUpperCaseLetter  = [...]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	simpleLatinLowerCaseLetter  = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	sqlSpecialCharacter         = [...]byte{' ', '"', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '>', '=', '?', '_', '|'}
	sqlEmbededLanguageCharacter = [...]byte{'[', ']'}
	separatorCharacter          = [...]byte{' ', '\t', '\r', '\n', 0}
	newlineCharacter            = [...]byte{'\r', '\n', 0}
)

func (l *Lexer) isSQLSpecialCharacter() bool {
	for _, v := range sqlSpecialCharacter {
		if v == l.ch {
			return true
		}
	}
	for _, v := range sqlEmbededLanguageCharacter {
		if v == l.ch {
			return true
		}
	}
	return false
}

func (l *Lexer) isSeparator() bool {
	for _, v := range separatorCharacter {
		if v == l.ch {
			return true
		}
	}
	if l.ch == '-' && l.peekChar() == '-' {
		return true
	}
	return false
}

func (l *Lexer) isPeekSeparator() bool {
	for _, v := range separatorCharacter {
		if v == l.peekChar() {
			return true
		}
	}
	if l.peekChar() == '-' && l.peek2Char() == '-' {
		return true
	}
	return false
}

func (l *Lexer) isNewline() bool {
	for _, v := range newlineCharacter {
		if v == l.ch {
			return true
		}
	}
	return false
}

func (l *Lexer) isDigit() bool {
	ch := l.ch
	if ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9' {
		return true
	}
	return false
}

func (l *Lexer) isPeekDigit() bool {
	ch := l.peekChar()
	if ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9' {
		return true
	}
	return false
}

func (l *Lexer) isBit() bool {
	if l.ch == '0' || l.ch == '1' {
		return true
	}
	return false
}

func (l *Lexer) isHexit() bool {
	ch := l.ch
	if l.isDigit() == true {
		return true
	}
	if ch == 'A' || ch == 'B' || ch == 'C' || ch == 'D' || ch == 'E' || ch == 'F' || ch == 'a' || ch == 'b' || ch == 'c' || ch == 'd' || ch == 'e' || ch == 'f' {
		return true
	}
	return false
}

func (l *Lexer) isSimpleLatin() bool {
	for _, v := range simpleLatinUpperCaseLetter {
		if v == l.ch {
			return true
		}
	}
	for _, v := range simpleLatinLowerCaseLetter {
		if v == l.ch {
			return true
		}
	}
	return false
}

func (l *Lexer) isNonquotedCharacter() bool {
	if l.isSimpleLatin() == true {
		return true
	}
	if l.isDigit() == true {
		return true
	}
	if l.isSQLSpecialCharacter() == true {
		if l.ch == '\'' {
			if l.peekChar() == '\'' {
				l.readChar()
				return true
			} else {
				return false
			}
		}
		return true
	}
	return false
}

func (l *Lexer) isNondoublequotedCharacter() bool {
	if l.isSimpleLatin() == true {
		return true
	}
	if l.isDigit() == true {
		return true
	}
	if l.isSQLSpecialCharacter() == true {
		if l.ch == '"' {
			if l.peekChar() == '"' {
				l.readChar()
				return true
			} else {
				return false
			}
		}
		return true
	}
	return false
}
