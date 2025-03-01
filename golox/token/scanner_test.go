package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanTokens(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		Name     string
		Src      string
		Tokens   []Type
		Literals []any
	}{
		{
			Name:   "single character",
			Src:    "*",
			Tokens: []Type{Star},
		},
		{
			Name:   "with whitespace",
			Src:    "\n * + *   \n   ",
			Tokens: []Type{Star, Plus, Star},
		},
		{
			Name:   "empty string",
			Src:    "",
			Tokens: []Type{},
		},
		{
			Name:   "two char tokens",
			Src:    "\t >= \n\n <==",
			Tokens: []Type{GreaterEqual, LessEqual, Equal},
		},
		{
			Name:   "keyword",
			Src:    "while",
			Tokens: []Type{While},
		},
		{
			Name:     "identifier and keyword",
			Src:      "beans and toast",
			Tokens:   []Type{Identifier, And, Identifier},
			Literals: []any{"beans", nil, "toast"},
		},
		{
			Name:     "string literal",
			Src:      `beans = "toast";`,
			Tokens:   []Type{Identifier, Equal, String, Semicolon},
			Literals: []any{"beans", nil, "toast", nil},
		}, {
			Name:     "variable declaration",
			Src:      "var hello = \"world\";",
			Tokens:   []Type{Var, Identifier, Equal, String, Semicolon},
			Literals: []any{nil, "hello", nil, "world", nil},
		},
	} {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			s := NewScanner([]byte(tc.Src))
			res := s.Scan()

			// token types
			var expectedTokens []string
			for _, t := range tc.Tokens {
				expectedTokens = append(expectedTokens, tokens[t])
			}
			expectedTokens = append(expectedTokens, EOF.String())

			var actualTokens []string
			for _, t := range res {
				actualTokens = append(actualTokens, tokens[t.Type])
			}

			assert.Equal(t, expectedTokens, actualTokens)

			// literals
			if len(tc.Literals) > 0 {
				expectedLiterals := tc.Literals
				expectedLiterals = append(expectedLiterals, nil) // EOF

				var actualLiterals []any
				for _, t := range res {
					actualLiterals = append(actualLiterals, t.Literal)

				}
				assert.Equal(t, expectedLiterals, actualLiterals)
			}
		})
	}
}
