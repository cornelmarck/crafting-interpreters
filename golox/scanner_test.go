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
			Tokens: []TokenType{Star, EOF},
		},
		{
			Name:   "with whitespace",
			Src:    "\n * + *   \n   ",
			Tokens: []TokenType{Star, Plus, Star, EOF},
		},
		{
			Name:   "empty string",
			Src:    "",
			Tokens: []TokenType{EOF},
		},
		{
			Name:   "two char tokens",
			Src:    "\t >= \n\n <==",
			Tokens: []TokenType{GreaterEqual, LessEqual, Equal, EOF},
		},
		{
			Name:   "keyword",
			Src:    "while",
			Tokens: []TokenType{While, EOF},
		},
		{
			Name:     "string literal",
			Src:      "beans and toast",
			Tokens:   []TokenType{Identifier, And, Identifier, EOF},
			Literals: []string{"beans", "", "toast", ""},
		},
	} {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			s := NewScanner()
			res := s.Scan([]byte(tc.Src))

			// token types
			var expectedTokens []string
			var actualTokens []string
			for _, t := range tc.Tokens {
				expectedTokens = append(expectedTokens, tokens[t])
			}
			for _, t := range res {
				actualTokens = append(actualTokens, tokens[t.Type])
			}
			assert.Equal(t, expectedTokens, actualTokens)

			// literal values
			if len(tc.Literals) > 0 {
				var actualLiterals []string
				for _, t := range res {
					actualLiterals = append(actualLiterals, t.Literal)
				}
				assert.Equal(t, tc.Literals, actualLiterals)
			}
		})
	}
}
