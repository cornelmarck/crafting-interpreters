package ast

import (
	"errors"
	"fmt"

	"github.com/cornelmarck/crafting-interpreters/golox/token"
)

// Parser is a recursive descent parser.
// The main todo is implementing syntax validation and error handling.

type Parser struct {
	tokens  []token.Token
	current token.Token
	offset  int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: tokens[0],
		offset:  0,
	}
}

func (p *Parser) Parse() ([]Statement, error) {
	var statements []Statement
	for !p.eof() {
		statement, err := p.declaration()
		if err != nil {
			return statements, err
		}
		statements = append(statements, statement)
	}
	return statements, nil
}

// Declaration

func (p *Parser) declaration() (Statement, error) {
	if p.match(token.Var) {
		p.next()
		return p.parseVarDeclaration()
	}
	return p.statement()
}

func (p *Parser) parseVarDeclaration() (Statement, error) {
	if !p.match(token.Identifier) {
		return nil, errors.New("missing identifier in variable declaration")
	}
	identifier := p.current
	p.next()

	var initializer Expression
	if p.match(token.Equal) {
		p.next()
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		initializer = expression
	}

	if !p.match(token.Semicolon) {
		return nil, errors.New("expected ';' after variable declaration")
	}
	p.next()
	return VariableDeclaration{
		Name:        identifier.Literal.(string),
		Initializer: initializer,
	}, nil
}

// Statements

func (p *Parser) statement() (Statement, error) {
	if p.match(token.Print) {
		p.next()
		return p.print()
	}
	expression, err := p.expression()
	if err != nil {
		return nil, err
	}
	if !p.match(token.Semicolon) {
		return nil, errors.New("expected ';' after expression statement")
	}
	p.next()

	return ExpressionStatement{
		Expression: expression,
	}, nil
}

func (p *Parser) print() (Statement, error) {
	expression, err := p.expression()
	if err != nil {
		return nil, err
	}

	node := PrintStatement{
		Expression: expression,
	}
	if !p.match(token.Semicolon) {
		return nil, errors.New("expected ';' after print statement")
	}
	p.next()
	return node, nil
}

// Expressions

func (p *Parser) expression() (Expression, error) {
	return p.equality()
}

func (p *Parser) assignment() (Expression, error) {
	return nil, nil
}

func (p *Parser) equality() (Expression, error) {
	left, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(token.EqualEqual, token.BangEqual) {
		operator := p.current.Type
		p.next()

		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		left = &BinaryExpression{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}
	return left, nil
}

func (p *Parser) comparison() (Expression, error) {
	left, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
		operator := p.current.Type
		p.next()

		right, err := p.term()
		if err != nil {
			return nil, err
		}
		left = &BinaryExpression{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}
	return left, nil
}

func (p *Parser) term() (Expression, error) {
	left, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(token.Minus, token.Plus) {
		operator := p.current.Type
		p.next()

		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		left = &BinaryExpression{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}
	return left, nil
}

func (p *Parser) factor() (Expression, error) {
	left, err := p.urnary()
	if err != nil {
		return nil, err
	}

	for p.match(token.Slash, token.Star) {
		operator := p.current.Type
		p.next()

		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		left = &BinaryExpression{
			Operator: operator,
			Left:     left,
			Right:    right,
		}
	}
	return left, nil
}

func (p *Parser) urnary() (Expression, error) {
	if p.match(token.Bang, token.Minus) {
		operator := p.current.Type
		p.next()

		right, err := p.urnary()
		if err != nil {
			return nil, err
		}
		return &UrnaryExpression{
			Operator: operator,
			Right:    right,
		}, nil
	}
	return p.primary()
}

func (p *Parser) primary() (Expression, error) {
	defer p.next()

	switch p.current.Type {
	case token.False:
		return BooleanExpression{Value: false}, nil
	case token.True:
		return BooleanExpression{Value: true}, nil
	case token.Nil:
		return NilExpression{}, nil
	case token.Number:
		return NumberExpression{Value: p.current.Literal.(float64)}, nil
	case token.String:
		return StringExpression{Value: p.current.Literal.(string)}, nil
	case token.LeftParen:
		p.next()
		grouping, err := p.expression()
		if err != nil {
		}

		if !p.match(token.RightParen) {
			return nil, errors.New("expected closing ')' after grouping expression")
		}
		return &GroupingExpression{Expression: grouping}, nil
	case token.Identifier:
		return VariableExpression{Name: p.current.Literal.(string)}, nil
	default:
		return nil, fmt.Errorf("unexpected token: %s", p.current.Type.String())
	}
}

// Helper functions

func (p *Parser) eof() bool {
	return p.current.Type == token.EOF
}

func (p *Parser) next() {
	if p.current.Type != token.EOF {
		p.offset += 1
		p.current = p.tokens[p.offset]
	}
}

func (p *Parser) match(types ...token.Type) bool {
	for _, t := range types {
		if p.current.Type == t {
			return true
		}
	}
	return false
}
