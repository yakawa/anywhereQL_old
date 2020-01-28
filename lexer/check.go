package lexer

var (
	digit                       = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	sqlSpecialCharacter         = [...]byte{' ', '"', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '>', '=', '?', '_', '|'}
	sqlEmbededLanguageCharacter = [...]byte{'[', ']'}
	separatorCharacter          = [...]byte{' ', '\t', '\r', '\n'}
	newlineCharacter            = [...]byte{'\r', '\n'}
)

func isSQLSpecialCharacter(ch byte) bool {
	for _, v := range sqlSpecialCharacter {
		if v == ch {
			return true
		}
	}
	for _, v := range sqlEmbededLanguageCharacter {
		if v == ch {
			return true
		}
	}
	return false
}

func isSeparatorCharacter(ch byte) bool {
	for _, v := range separatorCharacter {
		if v == ch {
			return true
		}
	}
	return false
}

func isNewline(ch byte) bool {
	for _, v := range newlineCharacter {
		if v == ch {
			return true
		}
	}
	return false
}

func isDigit(ch byte) bool {
	for _, v := range digit {
		if v == ch {
			return true
		}
	}
	return false
}
