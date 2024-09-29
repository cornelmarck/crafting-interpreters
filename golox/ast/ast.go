package ast

import "golox/token"

type Node interface {
	Type() NodeType
}

type NodeType int

const (
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
)

// Literals
type BooleanNode struct {
	Value bool
}

func (n BooleanNode) Type() NodeType {
	return Boolean
}

type NilNode struct{}

func (n NilNode) Type() NodeType {
	return Nil
}

type NumberNode struct {
	Value float64
}

func (n NumberNode) Type() NodeType {
	return Number
}

type StringNode struct {
	Value string
}

func (n StringNode) Type() NodeType {
	return String
}

// Assign
type AssignNode struct {
	Name  string
	Value any
}

func (n *AssignNode) Type() NodeType {
	return Assign
}

// Binary
type BinaryNode struct {
	Operator token.Token
	Left     Node
	Right    Node
}

func (n *BinaryNode) Type() NodeType {
	return Binary
}

// Urnary
type UrnaryNode struct {
	Operator token.Token
	Right    Node
}

func (n *UrnaryNode) Type() NodeType {
	return Urnary
}

// Grouping
type GroupingNode struct {
	Expression Node
}

func (n *GroupingNode) Type() NodeType {
	return Grouping
}
