package interpreter

import (
	"bytes"
	"testing"

	"github.com/cornelmarck/crafting-interpreters/golox/ast"
	"github.com/cornelmarck/crafting-interpreters/golox/token"

	"github.com/stretchr/testify/assert"
)

func TestInterpret(t *testing.T) {
	print := func(expr ast.Expression) *ast.PrintStatement {
		return &ast.PrintStatement{Expression: expr}
	}

	for _, tc := range []struct {
		name     string
		execFunc func() []ast.Statement
		expected string
		err      string
	}{
		{
			name: "print literal",
			execFunc: func() []ast.Statement {
				expr := ast.StringExpression{Value: "hello world"}
				return []ast.Statement{print(expr)}
			},
			expected: "hello world",
		}, {
			name: "print expression",
			execFunc: func() []ast.Statement {
				expr := &ast.BinaryExpression{
					Operator: token.Plus,
					Left:     ast.NumberExpression{Value: 2},
					Right:    ast.NumberExpression{Value: 1},
				}
				return []ast.Statement{print(expr)}
			},
			expected: "3",
			// }, {
			// 	name: "assign variable",
			// 	execFunc: func() []ast.Statement {
			// 		var lines []ast.Statement

			// 		lines = append(lines, &ast.ExpressionStatement{
			// 			&ast.AssignExpression{
			// 				Name:  "x",
			// 				Value: "1",
			// 			},
			// 		})
			// 		lines = append(lines, print(ast.VariableExpression{}))
			// 		return lines
			// },
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			i := New(&buf)

			for _, statement := range tc.execFunc() {
				err := i.Interpret(statement)
				if tc.err != "" {
					assert.ErrorContains(t, err, tc.err)
				}
				assert.NoError(t, err)
			}
			expected := tc.expected + "\n"
			assert.Equal(t, expected, buf.String())
		})
	}
}
