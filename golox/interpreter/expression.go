package interpreter

import (
	"errors"
	"fmt"

	"github.com/cornelmarck/crafting-interpreters/golox/ast"
	"github.com/cornelmarck/crafting-interpreters/golox/token"
)

func evaluateExpression(expr ast.Expression, env environment) (any, error) {
	switch node := expr.(type) {
	case ast.BooleanExpression:
		return node.Value, nil
	case ast.NilExpression:
		return nil, nil
	case ast.NumberExpression:
		return node.Value, nil
	case ast.StringExpression:
		return node.Value, nil
	case ast.VariableExpression:
		return env.get(node.Name)
	case *ast.GroupingExpression:
		return evaluateExpression(node.Expression, env)
	case *ast.UrnaryExpression:
		return evaluateUrnary(node, env)
	case *ast.BinaryExpression:
		return evaluateBinary(node, env)
	default:
		return nil, fmt.Errorf("invalid expression: %d", node.Type())
	}
}

func evaluateUrnary(node *ast.UrnaryExpression, env environment) (any, error) {
	right, err := evaluateExpression(node.Right, env)
	if err != nil {
		return nil, err
	}

	switch node.Operator {
	case token.Bang:
		return !isTruthy(right), nil
	case token.Minus:
		v, ok := castUrnaryOperand[float64](right)
		if !ok {
			return nil, errors.New("invalid operand")
		}
		return -v, nil
	}

	return nil, errors.New("unknown urnary operator")
}

func evaluateBinary(node *ast.BinaryExpression, env environment) (any, error) {
	left, err := evaluateExpression(node.Left, env)
	if err != nil {
		return nil, err
	}
	right, err := evaluateExpression(node.Right, env)
	if err != nil {
		return nil, err
	}

	// Equality check does not have type restrictions
	switch node.Operator {
	case token.EqualEqual:
		return left == right, nil
	case token.BangEqual:
		return left != right, nil
	}

	// Plus is a special case because of string concatenation
	if node.Operator == token.Plus {
		l, r, ok := castBinaryOperand[string](left, right)
		if ok {
			return l + r, nil
		}
	}

	// All remaining operands are numeric
	l, r, ok := castBinaryOperand[float64](left, right)
	if !ok {
		return nil, errors.New("invalid operand")
	}

	switch node.Operator {
	// arithmetic
	case token.Minus:
		return l - r, nil
	case token.Plus:
		return l + r, nil
	case token.Slash:
		if r == .0 {
			return nil, errors.New("divide by zero")
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

	return nil, fmt.Errorf("invalid binary operator: %v", node.Operator.String())
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
