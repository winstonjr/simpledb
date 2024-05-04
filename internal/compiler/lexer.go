package compiler

import "unicode"

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
	tt    TokenType
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
		return Token{tt: TokenEOF, value: ""}
	}

	c := l.input[l.position]
	switch {
	case c == ' ' || c == '\t' || c == '\n' || c == '\r':
		l.consumeWhitespace()
		return l.NextToken()
	case c == ',' || c == ';' || c == '*' || c == '(' || c == ')' || c == '+' || c == '-' || c == '\\':
		l.position++
		return Token{tt: TokenSymbol, value: string(c)}
	case unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)):
		return l.consumeIdentifier()
	default:
		return Token{tt: TokenError, value: string(c)}
	}
}

func (l *LexerSimple) consumeWhitespace() {
	for l.position < len(l.input) &&
		(l.input[l.position] == ' ' || l.input[l.position] == '\t' || l.input[l.position] == '\n' || l.input[l.position] == '\r') {
		l.position++
	}
}

func (l *LexerSimple) consumeIdentifier() Token {
	start := l.position
	for l.position < len(l.input) &&
		(unicode.IsNumber(rune(l.input[l.position])) || unicode.IsLetter(rune(l.input[l.position]))) {
		l.position++
	}

	value := l.input[start:l.position]
	return Token{tt: TokenIdentifier, value: value}
}
