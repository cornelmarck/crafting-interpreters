package token

import (
	"fmt"
)

type Token struct {
	Type    Type
	Literal any
	Pos     Position
}

type Position struct {
	Offset int // absolute offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1
}

type Type int

const (
	// Special tokens
	Illegal Type = iota
	EOF

	// Single-character tokens
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One or two character tokens
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals
	Identifier
	String
	Number

	// Keywords
	keyword_beg
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While
	keyword_end
)

var tokens = [...]string{
	Illegal: "illegal",
	EOF:     "eof",

	LeftParen:  "(",
	RightParen: ")",
	LeftBrace:  "{",
	RightBrace: "}",
	Comma:      ",",
	Dot:        ".",
	Minus:      "-",
	Plus:       "+",
	Semicolon:  ";",
	Slash:      "/",
	Star:       "*",

	Bang:         "!",
	BangEqual:    "!=",
	Equal:        "=",
	EqualEqual:   "==",
	Greater:      ">",
	GreaterEqual: ">=",
	Less:         "<",
	LessEqual:    "<=",

	Identifier: "identifier",
	String:     "string",
	Number:     "number",

	And:    "and",
	Class:  "class",
	Else:   "else",
	False:  "false",
	For:    "for",
	Fun:    "fun",
	If:     "if",
	Nil:    "nil",
	Or:     "or",
	Print:  "print",
	Return: "return",
	Super:  "super",
	This:   "this",
	True:   "true",
	Var:    "var",
	While:  "while",
}

func (t Type) String() string {
	if 0 <= t && t < Type(len(tokens)) {
		return tokens[t]
	}
	return fmt.Sprintf("token(%d)", t)
}

var keywords map[string]Type

func init() {
	keywords = make(map[string]Type, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup if token is a keyword, defaulting to an identifier if not found
func Lookup(ident string) Type {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return Identifier
}
