package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/cornelmarck/crafting-interpreters/golox/ast"
	"github.com/cornelmarck/crafting-interpreters/golox/interpreter"
	"github.com/cornelmarck/crafting-interpreters/golox/token"
)

func main() {
	if len(os.Args) == 1 {
		runPrompt()
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		fmt.Println("usage: jlox [script]")
		os.Exit(64)
	}
}

func runPrompt() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (Ctrl+D to stop)")

	for {
		fmt.Print("> ")
		if ok := reader.Scan(); !ok {
			return
		}
		line := reader.Bytes()
		run(line)
	}
}

func runFile(name string) {
	fmt.Println("running file")
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("could not open '%s': %v", name, err)
		os.Exit(2)
	}

	src, err := io.ReadAll(file)
	file.Close()
	if err != nil {
		fmt.Printf("could not copy src: %v", err)
		os.Exit(1)
	}

	run(src)
}

func run(src []byte) {
	scanner := token.NewScanner(src)
	tokens := scanner.Scan()

	parser := ast.NewParser(tokens)
	statements, err := parser.Parse()
	handleError(err)

	interpreter := interpreter.New(os.Stdout)
	err = interpreter.Interpret(statements...)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("runtime error: %v\n", err)
		os.Exit(0)
	}
}
