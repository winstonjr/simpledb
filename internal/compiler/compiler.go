package compiler

type Compiler interface {
	Call() []Token
}

type compiler struct {
	lexer  Lexer
	parser Parser
}

func NewCompiler(sqlText string) Compiler {
	lexer := NewLexer(sqlText)
	parser := NewParser(lexer)

	return &compiler{
		lexer:  lexer,
		parser: parser,
	}
}

func (comp *compiler) Call() []Token {
	comp.parser.ParseStatement()
	return comp.parser.GetTokens()
}
