package lexer

import (
	"fmt"
	"testing"

	"github.com/xmonader/monkey-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := NewLexer(input)
	fmt.Println(l)
	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println("TOK: ", tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d]: Expected %v and Got %v", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d]: Expected literal %s and got %s ", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// func TestGetNextTokenSimple(t *testing.T) {
// 	input := "=+(){},;"
// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.ASSIGN, "="},
// 		{token.PLUS, "+"},
// 		{token.LPAREN, "("},
// 		{token.RPAREN, ")"},
// 		{token.LBRACE, "{"},
// 		{token.RBRACE, "}"},
// 		{token.COMMA, ","},
// 		{token.SEMICOLON, ";"},
// 	}

// 	l := NewLexer(input)
// 	for i, tt := range tests {
// 		tok := l.NextToken()
// 		fmt.Println("TOK: ", tok)
// 		if tok.Type != tt.expectedType {
// 			t.Fatalf("test[%d]: Expected %v and Got %v", i, tt.expectedType, tok.Type)
// 		}
// 		if tok.Literal != tt.expectedLiteral {
// 			t.Fatalf("test[%d]: Expected literal %s and got %s ", i, tt.expectedLiteral, tok.Literal)
// 		}
// 	}
// }
