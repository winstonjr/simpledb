package main

import (
	"bufio"
	"fmt"
	"github.com/winstonjr/simpledb/internal/compiler"
	"os"
)

type REPL interface {
	StartREPL()
}

type replStruct struct {
}

func New() REPL {
	return &replStruct{}
}

func (r *replStruct) StartREPL() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error reading input %v \n", err)
			continue
		}

		input = input[:len(input)]

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		r.evaluate(input)
	}
}

func (r *replStruct) evaluate(str string) {
	comp := compiler.NewCompiler(str)
	t := comp.Call()
	fmt.Println("Output ...", t)
}

func main() {
	cmd := New()
	cmd.StartREPL()

	fmt.Println("Goodbye ...")
	os.Exit(0)
}
