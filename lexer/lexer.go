package lexer

import (
	"fmt"
	"unicode"

	"github.com/xmonader/monkey-lang/token"
)

type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // next char position to read
	ch           byte // current char
}

func NewLexer(input string) *Lexer {
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

func (l *Lexer) NextToken() *token.Token {
	var tok *token.Token
	l.skipWhitespace()
	fmt.Println("CHAR IS :", l.ch)
	switch l.ch {

	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case 0:
		tok = token.NewToken(token.EOF, 0)
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok = new(token.Token)
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifierType(tok.Literal)
			return tok
		} else if unicode.IsDigit(rune(l.ch)) {
			tok = new(token.Token)
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	fmt.Println("TOK: ", tok)
	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_' || ch == '?'
}
