package interpreter

import (
	"errors"
	"fmt"
	"io"

	"github.com/cornelmarck/crafting-interpreters/golox/ast"
)

type Interpreter struct {
	env     environment
	printer io.Writer
}

func New(printer io.Writer) *Interpreter {
	return &Interpreter{
		env:     environment{},
		printer: printer,
	}
}

func (i *Interpreter) Interpret(statements ...ast.Statement) error {
	for _, s := range statements {
		if err := i.execute(s); err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) execute(statement ast.Statement) error {
	switch node := statement.(type) {
	case *ast.PrintStatement:
		return i.executePrint(node)
	case *ast.ExpressionStatement:
		_, err := evaluateExpression(node.Expression, i.env)
		return err
	case *ast.VariableDeclaration:
		value, err := evaluateExpression(node.Initializer, i.env)
		if err != nil {
			return err
		}
		i.env.set(node.Name, value)
		return nil
	default:
		return errors.New("unknown statement")
	}
}

func (i *Interpreter) executePrint(node *ast.PrintStatement) error {
	value, err := evaluateExpression(node.Expression, i.env)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(i.printer, value)
	return err
}

func (i *Interpreter) declareVar() error {
	return nil
}
