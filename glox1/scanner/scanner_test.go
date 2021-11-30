package scanner

import (
	"testing"
)

func TestSymbolScanner(t *testing.T) {
	text := `(()){} // grouping symbols

* + - / // arithmetic operators
// blank comment line
!=!<>>=<=;== // other operators
`
	expTok := []Token{
		{LEFT_PAREN, "(", nil, 1},
		{LEFT_PAREN, "(", nil, 1},
		{RIGHT_PAREN, ")", nil, 1},
		{RIGHT_PAREN, ")", nil, 1},
		{LEFT_BRACE, "{", nil, 1},
		{RIGHT_BRACE, "}", nil, 1},
		{STAR, "*", nil, 3},
		{PLUS, "+", nil, 3},
		{MINUS, "-", nil, 3},
		{SLASH, "/", nil, 3},
		{BANG_EQUAL, "!=", nil, 5},
		{BANG, "!", nil, 5},
		{LESS, "<", nil, 5},
		{GREATER, ">", nil, 5},
		{GREATER_EQUAL, ">=", nil, 5},
		{LESS_EQUAL, "<=", nil, 5},
		{SEMICOLON, ";", nil, 5},
		{EQUAL_EQUAL, "==", nil, 5},
		{EOF, "", nil, 6},
	}

	scanner := NewScanner(text)
	tok := scanner.Scan()
	
	if len(tok) != len(expTok) {
		t.Errorf("Expected length %d for tokens, got %d.",
			len(expTok), len(tok))
	}

	n := len(tok)
	if len(expTok) < n { n = len(expTok) }

	for i := 0; i < n; i++ {
		if tok[i].Type != expTok[i].Type ||
			tok[i].Lexeme != expTok[i].Lexeme ||
			tok[i].Literal != expTok[i].Literal ||
			tok[i].Line != expTok[i].Line {
			t.Errorf("tok[%d]: expected %v, got %v.", i, expTok[i], tok[i])
		}
	}
}

func TestLiteralScanner(t *testing.T) {
	text := `123 12.0 11.
1e2 1.0e2 1.0e+2 2.5e-1 1.e2
"" "hello"`
	expTok := []Token{
		{NUMBER, "123", 123.0, 1},
		{NUMBER, "12.0", 12.0, 1},
		{NUMBER, "11.", 11.0, 1},
		{NUMBER, "1e2", 100.0, 2},
		{NUMBER, "1.0e2", 100.0, 2},
		{NUMBER, "1.0e+2", 100.0, 2},
		{NUMBER, "2.5e-1", 0.25, 2},
		{NUMBER, "1.e2", 100.0, 2},
		{STRING, `""`, "", 3},
		{STRING, `"hello"`, "hello", 3},
		{EOF, "", nil, 3},
	}

	scanner := NewScanner(text)
	tok := scanner.Scan()
	
	if len(tok) != len(expTok) {
		t.Errorf("Expected length %d for tokens, got %d.",
			len(expTok), len(tok))
	}

	n := len(tok)
	if len(expTok) < n { n = len(expTok) }

	for i := 0; i < n; i++ {
		if tok[i].Type != expTok[i].Type ||
			tok[i].Lexeme != expTok[i].Lexeme ||
			tok[i].Literal != expTok[i].Literal ||
			tok[i].Line != expTok[i].Line {
			t.Errorf("tok[%d]: expected %v, got %v.", i, expTok[i], tok[i])
		}
	}
}

func TestIdentifierScanner(t *testing.T) {
	text := `and class else false for
    fun if nil or print
return super this true while
myvar
my_var?!
m2222m`
	expTok := []Token{
		{AND, "and", nil, 1},
		{CLASS, "class", nil, 1},
		{ELSE, "else", nil, 1},
		{FALSE, "false", nil, 1},
		{FOR, "for", nil, 1},
		{FUN, "fun", nil, 2},
		{IF, "if", nil, 2},
		{NIL, "nil", nil, 2},
		{OR, "or", nil, 2},
		{PRINT, "print", nil, 2},
		{RETURN, "return", nil, 3},
		{SUPER, "super", nil, 3},
		{THIS, "this", nil, 3},
		{TRUE, "true", nil, 3},
		{WHILE, "while", nil, 3},
		{IDENTIFIER, "myvar", nil, 4},
		{IDENTIFIER, "my_var?!", nil, 5},
		{IDENTIFIER, "m2222m", nil, 6},
		{EOF, "", nil, 6},
	}

	scanner := NewScanner(text)
	tok := scanner.Scan()
	
	if len(tok) != len(expTok) {
		t.Errorf("Expected length %d for tokens, got %d.",
			len(expTok), len(tok))
	}

	n := len(tok)
	if len(expTok) < n { n = len(expTok) }

	for i := 0; i < n; i++ {
		if tok[i].Type != expTok[i].Type ||
			tok[i].Lexeme != expTok[i].Lexeme ||
			tok[i].Literal != expTok[i].Literal ||
			tok[i].Line != expTok[i].Line {
			t.Errorf("tok[%d]: expected %v, got %v.", i, expTok[i], tok[i])
		}
	}
}
