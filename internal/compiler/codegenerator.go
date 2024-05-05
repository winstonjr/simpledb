package compiler

import (
	"github.com/winstonjr/simpledb/internal/model"
	"strconv"
	"strings"
)

type Generator interface {
	GenerateCode(tokens []Token) *model.Bytecode
}

func NewGenerator() Generator {
	return &generator{}
}

type generator struct {
}

func (g *generator) GenerateCode(tokens []Token) *model.Bytecode {
	var bt *model.Bytecode

	if tokens[0].Type == TokenKeyword && strings.ToUpper(tokens[0].value) == "INSERT" {
		bt = g.GenerateCodeInsert(tokens)
	}

	if tokens[0].Type == TokenKeyword && strings.ToUpper(tokens[0].value) == "SELECT" {
		bt = g.GenerateCodeSelect(tokens)
	}

	return bt
}

func stringToStringPtr(s string) *string {
	return &s
}

func stringToIntPtr(s string) *int {
	v, _ := strconv.Atoi(s)
	return &v
}

func isInteger(s string) bool {
	_, e := strconv.Atoi(s)
	return e == nil
}

func (g *generator) GenerateCodeInsert(tokens []Token) *model.Bytecode {
	bt := &model.Bytecode{}

	insts := []model.BytecodeValue{
		{
			Type:       model.BytecodeOperationTypeINSERT,
			Identifier: stringToStringPtr("INSERT"),
		},
	}

	tblName := model.BytecodeValue{
		Type:       model.BytecodeOperationTypeTableName,
		Identifier: stringToStringPtr(tokens[2].value),
	}

	varNames := []model.BytecodeValue{}
	i := 3
	for i < len(tokens) {
		if tokens[i].Type == TokenKeyword && strings.ToUpper(tokens[i].value) == "VALUES" {
			break
		}

		if tokens[i].Type == TokenIdentifier {
			v := model.BytecodeValue{
				Type:       model.BytecodeOperationTypeIdentifier,
				Identifier: stringToStringPtr(tokens[i].value),
			}

			varNames = append(varNames, v)
		}
		i++
	}
	i++

	countVarNames := model.BytecodeValue{
		Type:  model.BytecodeOperationTypeCount,
		Count: len(varNames),
	}

	varVals := []model.BytecodeValue{}
	for i < len(tokens) {
		if tokens[i].Type == TokenSymbol && tokens[i].value == ";" {
			break
		}

		if tokens[i].Type == TokenIdentifier {
			v := model.BytecodeValue{
				Type: model.BytecodeOperationTypeIdentifier,
			}

			if isInteger(tokens[i].value) {
				v.IntValue = stringToIntPtr(tokens[i].value)
			} else {
				v.StringValue = stringToStringPtr(tokens[i].value)
			}

			varVals = append(varVals, v)
		}
		i++
	}

	countVarVals := model.BytecodeValue{
		Type:  model.BytecodeOperationTypeCount,
		Count: len(varVals),
	}

	insts = append(insts, tblName)
	insts = append(insts, countVarNames)
	insts = append(insts, varNames...)
	insts = append(insts, countVarVals)
	insts = append(insts, varVals...)

	bt.Instructions = insts

	return bt
}

func (g *generator) GenerateCodeSelect(tokens []Token) *model.Bytecode {
	bt := &model.Bytecode{}

	insts := []model.BytecodeValue{
		{
			Type:       model.BytecodeOperationTypeSELECT,
			Identifier: stringToStringPtr("SELECT"),
		},
	}

	varNames := []model.BytecodeValue{}
	i := 1
	for i < len(tokens) {
		if tokens[i].Type == TokenKeyword && strings.ToUpper(tokens[i].value) == "FROM" {
			break
		}

		if tokens[i].Type == TokenIdentifier {
			v := model.BytecodeValue{
				Type:       model.BytecodeOperationTypeIdentifier,
				Identifier: stringToStringPtr(tokens[i].value),
			}

			varNames = append(varNames, v)
		}
		i++
	}
	i++

	tblName := model.BytecodeValue{
		Type:       model.BytecodeOperationTypeTableName,
		Identifier: stringToStringPtr(tokens[i].value),
	}

	countVarNames := model.BytecodeValue{
		Type:  model.BytecodeOperationTypeCount,
		Count: len(varNames),
	}

	insts = append(insts, tblName)
	insts = append(insts, countVarNames)
	insts = append(insts, varNames...)

	bt.Instructions = insts

	return bt
}
