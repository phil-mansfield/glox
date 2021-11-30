package scanner

import (
	"fmt"
)

type TokenType int
const (
	LEFT_PAREN TokenType = iota // Speical characters
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER // Literals
	STRING
	NUMBER
	AND // Keywords
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF // EOF
)

type Token struct {
	Type TokenType
	Lexeme string
	Literal interface{} // ??
	Line int
}

func (tok *Token) String() string {
	tokStr := ""
	switch tok.Type {
	case LEFT_PAREN: tokStr = "LEFT_PAREN"
	case RIGHT_PAREN: tokStr = "RIGHT_PAREN"
	case LEFT_BRACE: tokStr = "LEFT_BRACE"
	case RIGHT_BRACE: tokStr = "RIGHT_BRACE"
	case COMMA: tokStr = "COMMA"
	case DOT: tokStr = "DOT"
	case MINUS: tokStr = "MINUS"
	case PLUS: tokStr = "PLUS"
	case SEMICOLON: tokStr = "SEMICOLON"
	case SLASH: tokStr = "SLASH"
	case STAR: tokStr = "STAR"
	case BANG: tokStr = "BANG"
	case BANG_EQUAL: tokStr = "BANG_EQUAL"
	case EQUAL: tokStr = "EQUAL"
	case EQUAL_EQUAL: tokStr = "EQUAL_EQUAL"
	case GREATER: tokStr = "GREATER"
	case GREATER_EQUAL: tokStr = "GREATER_EQUAL"
	case LESS: tokStr = "LESS"
	case LESS_EQUAL: tokStr = "LESS_EQUAL"
	case IDENTIFIER: tokStr = "IDENTIFIER"
	case STRING: tokStr = "STRING"
	case NUMBER: tokStr = "NUMBER"
	case AND: tokStr = "AND"
	case CLASS: tokStr = "CLASS"
	case ELSE: tokStr = "ELSE"
	case FALSE: tokStr = "FALSE"
	case FUN: tokStr = "FUN"
	case FOR: tokStr = "FOR"
	case IF: tokStr = "IF"
	case NIL: tokStr = "NIL"
	case OR: tokStr = "OR"
	case PRINT: tokStr = "PRINT"
	case RETURN: tokStr = "RETURN"
	case SUPER: tokStr = "SUPER"
	case THIS: tokStr = "THIS"
	case TRUE: tokStr = "TRUE"
	case VAR: tokStr = "VAR"
	case WHILE: tokStr = "WHILE"
	case EOF: tokStr = "EOF"
	}

	return fmt.Sprintf("{%s: ('%s', line %d), %v}",
		tokStr, tok.Lexeme, tok.Line, tok.Literal)
}

func (tok *Token) GoString() string {
	return tok.String()
}
