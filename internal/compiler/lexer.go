package compiler

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenIdentifier
	TokenKeyword
	TokenSymbol
	TokenWhitespace
)

type Token struct {
	Type  TokenType
	value string
}

type Lexer interface {
	NextToken() Token
}

type LexerSimple struct {
	input    string
	position int
}

func NewLexer(input string) Lexer {
	return &LexerSimple{input: input, position: 0}
}

func (l *LexerSimple) NextToken() Token {
	if l.position >= len(l.input) {
		return Token{Type: TokenEOF, value: ""}
	}

	c := l.input[l.position]
	switch {
	case c == ' ' || c == '\t' || c == '\n', c == '\'':
		l.consumeWhitespace()
		return l.NextToken()
	case c == ',' || c == ';' || c == '*' || c == '(' || c == ')' || c == '+' || c == '-' || c == '\\':
		l.position++
		return Token{Type: TokenSymbol, value: string(c)}
	case unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)):
		return l.consumeIdentifier()
	default:
		return Token{Type: TokenError, value: string(c)}
	}
}

func (l *LexerSimple) consumeWhitespace() {
	for l.position < len(l.input) &&
		(l.input[l.position] == ' ' || l.input[l.position] == '\t' || l.input[l.position] == '\n' || l.input[l.position] == '\r' || l.input[l.position] == '\'') {
		l.position++
	}
}

//func isLetter(c byte) bool {
//	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
//}
//
//func isNumber(c byte) bool {
//	return '0' <= c && c <= '9'
//}

func (l *LexerSimple) consumeIdentifier() Token {
	start := l.position
	for l.position < len(l.input) &&
		(unicode.IsNumber(rune(l.input[l.position])) || unicode.IsLetter(rune(l.input[l.position]))) {
		//(isLetter(l.input[l.position]) || isNumber(l.input[l.position])) {
		l.position++
	}

	value := l.input[start:l.position]
	t := TokenIdentifier
	if strings.ToUpper(value) == "SELECT" ||
		strings.ToUpper(value) == "INSERT" ||
		strings.ToUpper(value) == "INTO" ||
		strings.ToUpper(value) == "VALUES" ||
		strings.ToUpper(value) == "FROM" ||
		strings.ToUpper(value) == "WHERE" {
		t = TokenKeyword
	}
	return Token{Type: t, value: value}
}
