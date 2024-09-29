package main

import (
	"golox/ast"
	"golox/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluate(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		Name       string
		Expression ast.Node
		Expected   any
	}{
		{
			Name: "4 - (3 - 2)",
			Expression: &ast.BinaryNode{
				Operator: token.Token{Type: token.Slash},
				Left:     ast.NumberNode{Value: float64(4)},
				Right: &ast.GroupingNode{
					Expression: &ast.BinaryNode{
						Operator: token.Token{Type: token.Minus},
						Left:     ast.NumberNode{Value: float64(3)},
						Right:    ast.NumberNode{Value: float64(2)},
					},
				},
			},
			Expected: float64(3),
		},
	} {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			res, err := Evaluate(tc.Expression)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, res)

		})
	}
}
