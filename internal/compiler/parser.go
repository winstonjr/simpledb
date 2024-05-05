package compiler

type Parser interface {
	ParseStatement()
	GetTokens() []Token
}

type ParserSimple struct {
	lexer  Lexer
	tokens []Token
}

func NewParser(l Lexer) Parser {
	return &ParserSimple{lexer: l, tokens: []Token{}}
}

func (p *ParserSimple) GetTokens() []Token {
	return p.tokens
}

func (p *ParserSimple) ParseStatement() {
	for {
		token := p.lexer.NextToken()
		p.tokens = append(p.tokens, token)

		if token.Type == TokenEOF {
			break
		}
	}
}
