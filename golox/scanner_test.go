package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanTokens(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		Name     string
		Src      string
		Tokens   []TokenType
		Literals []string
	}{
		{
			Name:   "single character",
			Src:    "*",
			Tokens: []TokenType{Star},
		},
		{
			Name:   "with whitespace",
			Src:    "\n * + *   \n   ",
			Tokens: []TokenType{Star, Plus, Star},
		},
		{
			Name:   "empty string",
			Src:    "",
			Tokens: []TokenType{},
		},
		{
			Name:   "two char tokens",
			Src:    "\t >= \n\n <==",
			Tokens: []TokenType{GreaterEqual, LessEqual, Equal},
		},
		{
			Name:   "keyword",
			Src:    "while",
			Tokens: []TokenType{While},
		},
		{
			Name:     "identifier and keyword",
			Src:      "beans and toast",
			Tokens:   []TokenType{Identifier, And, Identifier},
			Literals: []string{"beans", "", "toast"},
		},
		{
			Name:     "string literal",
			Src:      `beans = "toast"`,
			Tokens:   []TokenType{Identifier, Equal, String},
			Literals: []string{"beans", "", "toast"},
		},
	} {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			s := NewScanner()
			res := s.Scan([]byte(tc.Src))

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
				expectedLiterals = append(expectedLiterals, "") // EOF

				var actualLiterals []string
				for _, t := range res {
					actualLiterals = append(actualLiterals, t.Literal)

				}
				assert.Equal(t, expectedLiterals, actualLiterals)
			}
		})
	}
}
