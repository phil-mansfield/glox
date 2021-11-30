package scanner

import (
	"fmt"
	"strconv"
	
	"github.com/phil-mansfield/glox/glox1/error"
)

var (
	keywordMap = map[string]TokenType {
		"and": AND, "class": CLASS, "else": ELSE, "false": FALSE, "for": FOR,
			"fun": FUN, "if": IF, "nil": NIL, "or": OR, "print": PRINT,
			"return": RETURN, "super": SUPER, "this": THIS, "true": TRUE,
			"while": WHILE,
	}
)

type Scanner struct {
	start, current, line int
	text string
	tokens []Token
}

func NewScanner(text string) *Scanner {
	return &Scanner{ start: 0, current: 0, line: 1, text: text, tokens: nil }
}

func (s *Scanner) Scan() []Token {
	for !s.atEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, Token{ EOF, "", nil, s.line})
	return s.tokens
}

func (s *Scanner) atEnd() bool {
	return s.current >= len(s.text)
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case ' ', '\r', '\t':
	case '\n': s.line++
	case '(': s.addToken(LEFT_PAREN, nil)
	case ')': s.addToken(RIGHT_PAREN, nil)
	case '{': s.addToken(LEFT_BRACE, nil)
	case '}': s.addToken(RIGHT_BRACE, nil)
	case ',': s.addToken(COMMA, nil)
	case '.': s.addToken(DOT, nil)
	case '-': s.addToken(MINUS, nil)
	case '+': s.addToken(PLUS, nil)
	case ';': s.addToken(SEMICOLON, nil)
	case '*': s.addToken(STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.atEnd() { s.advance() }
		} else {
			s.addToken(SLASH, nil)
		}
	case '"': s.string()		
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifierOrReserved()
		} else {
			error.Error(s.line, fmt.Sprintf("Unexpeted character, %q", c))
		}
	}
}

func (s *Scanner) advance() byte {
	c := s.text[s.current]
	s.current++
	return c
}

func (s *Scanner) addToken(t TokenType, literal interface{}) {
	text := s.text[s.start: s.current]
	s.tokens = append(s.tokens, Token{t, text, literal, s.line})
}

func (s *Scanner) match(c byte) bool {
	if s.atEnd() { return false }
	if s.text[s.current] != c { return false }
	
	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.atEnd() { return 0 }
	return s.text[s.current]
}

func (s *Scanner) peekAhead(n int) byte {
	if s.current + n >= len(s.text) { return 0 }
	return s.text[s.current + n]
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.atEnd() {
		if s.peek() == '\n' { s.line++ }
		s.advance()
	}

	if s.atEnd() {
		error.Error(s.line, "Unterminated string")
	} else {
		s.advance()
		s.addToken(STRING, s.text[s.start+1: s.current-1])
	}
}

func (s *Scanner) number() {
	for isDigit(s.peek()) { s.advance() }
	if s.peek() == '.' {
		s.advance()
		for isDigit(s.peek()) { s.advance() }
	}

	if c := s.peek(); c == 'e' || c == 'E' {
		if isDigit(s.peekAhead(1)) {
			s.advance()
			for isDigit(s.peek()) { s.advance() }
		} else if c1, c2 := s.peekAhead(1), s.peekAhead(2); (c1 == '-' ||
			c1 == '+') && isDigit(c2) {
			s.advance()
			s.advance()
			for isDigit(s.peek()) { s.advance() }
		}
	}

	x, err := strconv.ParseFloat(s.text[s.start: s.current], 64)
	if err != nil {
		error.Error(s.line, "[Internal Error] glox recognizes '%s' as a " +
			"number, but Go's float parser does not recognize it.")
	}
	
	s.addToken(NUMBER, x)
}

func (s *Scanner) identifierOrReserved() {
	for isIdentifierByte(s.peek()) {
		s.advance()
	}

	name := s.text[s.start: s.current]
	t, ok := keywordMap[name]
	if !ok { t = IDENTIFIER }
	
	s.addToken(t, nil)
	
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isIdentifierByte(c byte) bool {
	return isDigit(c) || isAlpha(c) || c == '_' || c == '?' || c == '!'
}
