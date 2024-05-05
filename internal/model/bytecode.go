package model

import "encoding/json"

type BytecodeOperationType int

const (
	BytecodeOperationTypeINSERT BytecodeOperationType = iota
	BytecodeOperationTypeSELECT
	BytecodeOperationTypeTableName
	BytecodeOperationTypeIdentifier
	BytecodeOperationTypeValue
	BytecodeOperationTypeCount
	BytecodeOperationTypeDELETE
	BytecodeOperationTypeUPDATE
)

type BytecodeValue struct {
	Type        BytecodeOperationType
	Identifier  *string
	IntValue    *int
	StringValue *string
	Count       int
}

func (b BytecodeValue) String() string {
	s, _ := json.Marshal(b)
	return string(s)
}

type Bytecode struct {
	Instructions []BytecodeValue
}
