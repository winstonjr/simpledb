package compiler

import "github.com/winstonjr/simpledb/internal/model"

type Compiler interface {
	Call() *model.Bytecode
}

type compiler struct {
	lexer     Lexer
	parser    Parser
	generator Generator
}

func NewCompiler(sqlText string) Compiler {
	lexer := NewLexer(sqlText)
	parser := NewParser(lexer)
	gen := NewGenerator()

	return &compiler{
		lexer:     lexer,
		parser:    parser,
		generator: gen,
	}
}

func (comp *compiler) Call() *model.Bytecode {
	comp.parser.ParseStatement()
	return comp.generator.GenerateCode(comp.parser.GetTokens())
}
