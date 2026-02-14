package lexer

import (
	"tpl/token"
	"tpl/utils"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {

	case '=':
		tok = token.NewToken(token.ASSIGN, string(l.ch))

	case '+':
		tok = token.NewToken(token.PLUS, string(l.ch))

	case '-':
		tok = token.NewToken(token.MINUS, string(l.ch))

	case '*':
		tok = token.NewToken(token.ASTERISK, string(l.ch))

	case '/':
		tok = token.NewToken(token.SLASH, string(l.ch))

	case '<':
		tok = token.NewToken(token.LT, string(l.ch))

	case '>':
		tok = token.NewToken(token.GT, string(l.ch))

	case '(':
		tok = token.NewToken(token.LPAREN, string(l.ch))

	case ')':
		tok = token.NewToken(token.RPAREN, string(l.ch))

	case '{':
		tok = token.NewToken(token.LBRACE, string(l.ch))

	case '}':
		tok = token.NewToken(token.RBRACE, string(l.ch))

	case 0:
		tok = token.NewToken(token.EOF, "")

	default:
		if utils.IsLetter(l.ch) {
			literal := l.readIdentifier()
			return token.NewToken(token.LookupIdent(literal), literal)
		} else if utils.IsDigit(l.ch) {
			literal := l.readNumber()
			return token.NewToken(token.INT, literal)
		} else {
			tok = token.NewToken(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for utils.IsWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for utils.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for utils.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
