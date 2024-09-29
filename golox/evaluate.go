package main

import (
	"errors"
	"golox/ast"
	"golox/token"
)

func Evaluate(expr ast.Node) (any, error) {
	switch node := expr.(type) {
	// primary
	case ast.BooleanNode:
		return node.Value, nil
	case ast.NilNode:
		return nil, nil
	case ast.NumberNode:
		return node.Value, nil
	case ast.StringNode:
		return node.Value, nil
	case *ast.GroupingNode:
		return Evaluate(node.Expression)
	case *ast.UrnaryNode:
		return evalateUrnary(node)
	case *ast.BinaryNode:
		return evaluateBinary(node)
	}

	return nil, nil
}

func evalateUrnary(node *ast.UrnaryNode) (any, error) {
	right, err := Evaluate(node.Right)
	if err != nil {
		return nil, err
	}

	switch node.Operator.Type {
	case token.Bang:
		return !isTruthy(right), nil
	case token.Minus:
		v, ok := castUrnaryOperand[float64](right)
		if !ok {
			return nil, ErrInvalidOperand
		}
		return -v, nil
	}

	return nil, errors.New("unknown urnary operator")
}

func evaluateBinary(node *ast.BinaryNode) (any, error) {
	left, err := Evaluate(node.Left)
	if err != nil {
		return nil, err
	}
	right, err := Evaluate(node.Right)
	if err != nil {
		return nil, err
	}

	// Equality check does not have type restrictions
	switch node.Operator.Type {
	case token.EqualEqual:
		return left == right, nil
	case token.BangEqual:
		return left != right, nil
	}

	// Plus is a special case because of string concatenation
	if node.Operator.Type == token.Plus {
		l, r, ok := castBinaryOperand[string](left, right)
		if ok {
			return l + r, nil
		}
	}

	// All remaining operands are numeric
	l, r, ok := castBinaryOperand[float64](left, right)
	if !ok {
		return nil, ErrInvalidOperand
	}

	switch node.Operator.Type {
	// arithmetic
	case token.Minus:
		return l - r, nil
	case token.Plus:
		return l + r, nil
	case token.Slash:
		if r == .0 {
			return nil, ErrDivideByZero
		}
		return l / r, nil
	case token.Star:
		return l * r, nil
	// comparison
	case token.Greater:
		return l > r, nil
	case token.GreaterEqual:
		return l >= r, nil
	case token.Less:
		return l < r, nil
	case token.LessEqual:
		return l <= r, nil
	}

	return nil, errors.New("unknown binary operator")
}

func isTruthy(value any) bool {
	if value == nil {
		return false
	}
	if v, ok := value.(bool); ok {
		return v
	}
	return true
}

func castUrnaryOperand[V any](value any) (v V, ok bool) {
	v, ok = value.(V)
	return
}

func castBinaryOperand[V any](a any, b any) (va V, vb V, ok bool) {
	va, ok = a.(V)
	if !ok {
		return
	}
	vb, ok = b.(V)
	if !ok {
		return
	}
	return
}
