package ast

import "github.com/cornelmarck/crafting-interpreters/golox/token"

type Expression interface {
	Node
	expressionNode()
}

type BooleanExpression struct {
	Value bool
}

func (be BooleanExpression) expressionNode() {}

func (be BooleanExpression) Type() NodeType {
	return Boolean
}

type NilExpression struct{}

func (n NilExpression) Type() NodeType {
	return Nil
}

func (ne NilExpression) expressionNode() {}

type NumberExpression struct {
	Value float64
}

func (ne NumberExpression) Type() NodeType {
	return Number
}

func (ne NumberExpression) expressionNode() {}

type StringExpression struct {
	Value string
}

func (se StringExpression) Type() NodeType {
	return String
}

func (se StringExpression) expressionNode() {}

type AssignExpression struct {
	Name  string
	Value any
}

func (n *AssignExpression) Type() NodeType {
	return Assign
}

func (ae *AssignExpression) expressionNode() {}

// Binary
type BinaryExpression struct {
	Operator token.Type
	Left     Expression
	Right    Expression
}

func (n *BinaryExpression) Type() NodeType {
	return Binary
}

func (be *BinaryExpression) expressionNode() {}

type UrnaryExpression struct {
	Operator token.Type
	Right    Expression
}

func (ue *UrnaryExpression) Type() NodeType {
	return Urnary
}

func (ue *UrnaryExpression) expressionNode() {}

type GroupingExpression struct {
	Expression Expression
}

func (ge *GroupingExpression) Type() NodeType {
	return Grouping
}

func (ge *GroupingExpression) expressionNode() {}

type VariableExpression struct {
	Name string
}

func (ve VariableExpression) Type() NodeType {
	return Variable
}

func (ve VariableExpression) expressionNode() {}
