package ast

type Node interface {
	Type() NodeType
}

type NodeType int

const (
	// Expressions
	Assign NodeType = iota
	Binary
	Boolean
	Call
	Get
	Grouping
	Logical
	Nil
	Number
	Set
	Super
	String
	This
	Urnary
	Variable

	// Statements
	Block
	Class
	Expression_
	Function
	If
	Print
	Return
	Var
	While
)
