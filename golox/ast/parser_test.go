package ast

import (
	"golox/token"
	"testing"

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
			Name: "2 / 4 + 1",
			Tokens: []token.Token{
				{Type: token.Number, Literal: float64(2)},
				{Type: token.Slash},
				{Type: token.Number, Literal: float64(4)},
				{Type: token.Plus},
				{Type: token.Number, Literal: float64(1)},
				{Type: token.EOF},
			},
			Expected: &BinaryNode{
				Operator: token.Token{Type: token.Plus},
				Left: &BinaryNode{
					Operator: token.Token{Type: token.Slash},
					Left:     NumberNode{Value: 2},
					Right:    NumberNode{Value: 4},
				},
				Right: NumberNode{Value: 1},
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
				{Type: token.EOF},
			},
			Expected: &BinaryNode{
				Operator: token.Token{Type: token.Slash},
				Left:     NumberNode{Value: float64(4)},
				Right: &GroupingNode{
					Expression: &BinaryNode{
						Operator: token.Token{Type: token.Minus},
						Left:     NumberNode{Value: float64(3)},
						Right:    NumberNode{Value: float64(2)},
					},
				},
			},
		},
	} {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			p := NewParser(tc.Tokens)
			assert.Equal(t, tc.Expected, p.Parse())
		})
	}
}
