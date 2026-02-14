package utils

func IsLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		ch == '_'
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func IsWhitespace(ch byte) bool {
	return ch == ' ' ||
		ch == '\n' ||
		ch == '\t' ||
		ch == '\r'
}
