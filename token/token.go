package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

const (
	// Special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA  = ","
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	THALAPATHY = "THALAPATHY"
	IAMWAITING = "IAMWAITING"
	ALLISWELL  = "ALLISWELL"
)

var keywords = map[string]TokenType{
	"thalapathy": THALAPATHY,
	"IamWaiting": IAMWAITING,
	"alliswell":  ALLISWELL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
