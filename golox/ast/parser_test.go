package ast

import (
	"testing"

	"github.com/cornelmarck/crafting-interpreters/golox/token"

	"github.com/stretchr/testify/assert"
)

func TestASTParser(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		Name     string
		Tokens   []token.Token
		Expected Node
	}{
		{
			Name: "1 - 1 + 2",
			Tokens: []token.Token{
				{Type: token.Number, Literal: float64(1)},
				{Type: token.Minus},
				{Type: token.Number, Literal: float64(1)},
				{Type: token.Plus},
				{Type: token.Number, Literal: float64(2)},
				{Type: token.Semicolon},
				{Type: token.EOF},
			},
			Expected: ExpressionStatement{
				Expression: &BinaryExpression{
					Operator: token.Plus,
					Left: &BinaryExpression{
						Operator: token.Minus,
						Left:     NumberExpression{Value: 1},
						Right:    NumberExpression{Value: 1},
					},
					Right: NumberExpression{Value: 2},
				},
			},
		}, {
			Name: "2 / 4 + 1",
			Tokens: []token.Token{
				{Type: token.Number, Literal: float64(2)},
				{Type: token.Slash},
				{Type: token.Number, Literal: float64(4)},
				{Type: token.Plus},
				{Type: token.Number, Literal: float64(1)},
				{Type: token.Semicolon},
				{Type: token.EOF},
			},
			Expected: ExpressionStatement{
				Expression: &BinaryExpression{
					Operator: token.Plus,
					Left: &BinaryExpression{
						Operator: token.Slash,
						Left:     NumberExpression{Value: 2},
						Right:    NumberExpression{Value: 4},
					},
					Right: NumberExpression{Value: 1},
				},
			},
		}, {
			Name: "4 / (3 - 2)",
			Tokens: []token.Token{
				{Type: token.Number, Literal: float64(4)},
				{Type: token.Slash},
				{Type: token.LeftParen},
				{Type: token.Number, Literal: float64(3)},
				{Type: token.Minus},
				{Type: token.Number, Literal: float64(2)},
				{Type: token.RightParen},
				{Type: token.Semicolon},
				{Type: token.EOF},
			},
			Expected: ExpressionStatement{
				Expression: &BinaryExpression{
					Operator: token.Slash,
					Left:     NumberExpression{Value: float64(4)},
					Right: &GroupingExpression{
						Expression: &BinaryExpression{
							Operator: token.Minus,
							Left:     NumberExpression{Value: float64(3)},
							Right:    NumberExpression{Value: float64(2)},
						},
					},
				},
			},
		}, {
			Name: "var a = 2.3;",
			Tokens: []token.Token{
				{Type: token.Var},
				{Type: token.Identifier, Literal: "a"},
				{Type: token.Equal},
				{Type: token.Number, Literal: float64(2.3)},
				{Type: token.Semicolon},
				{Type: token.EOF},
			},
			Expected: VariableDeclaration{
				Name:        "a",
				Initializer: NumberExpression{Value: 2.3},
			},
		},
	} {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(tc.Tokens)
			res, err := p.Parse()
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, res[0])
		})
	}
}
