package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"

	EQ       = "=="
	NOT_EQ   = "!="
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func NewToken(tt TokenType, ch byte) *Token {
	return &Token{
		Type:    tt,
		Literal: string(ch),
	}
}

func LookupIdentifierType(ident string) TokenType {
	if tok, exists := keywords[ident]; exists {
		return tok
	}
	return IDENT
}
