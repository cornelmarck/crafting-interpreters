package interpreter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cornelmarck/crafting-interpreters/golox/ast"
	"github.com/cornelmarck/crafting-interpreters/golox/token"
)

var (
	indentString = strings.Repeat("\t", 4)
)

func TestInterpreter(t *testing.T) {
	for _, tc := range []struct {
		name     string
		code     string
		expected string
		err      string
	}{
		{
			name: "print lines",
			code: `
				print "hello world!";
				print "how are you?";
			`,
			expected: `
				hello world!
				how are you?
			`,
		}, {
			name: "add ints",
			code: `
				print 1 + 2;
			`,
			expected: `3`,
		}, {
			name: "declare variable",
			code: `
				var x = 1;
				print x;
			`,
			expected: `1`,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			scanner := token.NewScanner([]byte(tc.code))
			tokens := scanner.Scan()

			parser := ast.NewParser(tokens)
			statements, err := parser.Parse()
			if err != nil {
				t.Fatalf("error parsing code: %v", err)
			}

			var buf bytes.Buffer
			interpreter := New(&buf)

			err = interpreter.Interpret(statements...)
			if tc.err == "" && err != nil {
				t.Fatalf("unexpected error: %v", err)
			} else if tc.err != "" && !strings.Contains(err.Error(), tc.err) {
				t.Fatalf("expected %s error, got %v", tc.err, err)
			}

			expectedStdOut := parseExpectedStdOut(tc.expected)
			if expectedStdOut != buf.String() {
				t.Fatalf("expected:\n%s\ngot:\n%s", expectedStdOut, buf.String())
			}
		})
	}
}

func parseExpectedStdOut(s string) string {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, indentString)
	}
	return strings.Join(lines, "\n") + "\n"
}
