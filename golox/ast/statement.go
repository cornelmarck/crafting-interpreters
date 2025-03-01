package ast

type Statement interface {
	Node
	statementNode()
}

type PrintStatement struct {
	Expression Expression
}

func (n PrintStatement) Type() NodeType {
	return Print
}

func (ps PrintStatement) statementNode() {}

type VariableDeclaration struct {
	Name        string
	Initializer Expression
}

func (vd VariableDeclaration) Type() NodeType {
	return Var
}

func (vd VariableDeclaration) statementNode() {}

type ExpressionStatement struct {
	Expression Expression
}

func (es ExpressionStatement) Type() NodeType {
	return Expression_
}

func (es ExpressionStatement) statementNode() {}
