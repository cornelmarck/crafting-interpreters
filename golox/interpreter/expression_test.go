package interpreter

import (
	"testing"

	"github.com/cornelmarck/crafting-interpreters/golox/ast"
	"github.com/cornelmarck/crafting-interpreters/golox/token"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateExpression(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		Name       string
		Expression ast.Expression
		Expected   any
		Error      string
	}{
		{
			Name: "success; 4 / (3 - 2)",
			Expression: &ast.BinaryExpression{
				Operator: token.Slash,
				Left:     ast.NumberExpression{Value: float64(4)},
				Right: &ast.GroupingExpression{
					Expression: &ast.BinaryExpression{
						Operator: token.Minus,
						Left:     ast.NumberExpression{Value: float64(3)},
						Right:    ast.NumberExpression{Value: float64(2)},
					},
				},
			},
			Expected: float64(4),
		}, {
			Name: "success; 4.5 - -7 == 11.5",
			Expression: &ast.BinaryExpression{
				Operator: token.EqualEqual,
				Left: &ast.BinaryExpression{
					Operator: token.Minus,
					Left:     ast.NumberExpression{Value: float64(4.5)},
					Right: &ast.UrnaryExpression{
						Operator: token.Minus,
						Right:    ast.NumberExpression{Value: float64(7)},
					},
				},
				Right: ast.NumberExpression{Value: float64(11.5)},
			},
			Expected: true,
		}, {
			Name: "invalid operands; 'hello' + 3",
			Expression: &ast.BinaryExpression{
				Operator: token.Plus,
				Left:     ast.StringExpression{Value: "hello"},
				Right:    ast.NumberExpression{Value: float64(3)},
			},
			Error: "invalid operand",
		}, {
			Name: "divide by zero",
			Expression: &ast.BinaryExpression{
				Operator: token.Slash,
				Left:     ast.NumberExpression{Value: float64(0)},
				Right:    ast.NumberExpression{Value: float64(0.0)},
			},
			Error: "divide by zero",
		},
	} {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			res, err := evaluateExpression(tc.Expression, environment{})
			if tc.Error == "" {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, res)
			} else {
				assert.ErrorContains(t, err, tc.Error)
			}
		})
	}
}
