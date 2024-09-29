package ast

import (
	tok "golox/token"
)

// Parser is a recursive descent parser

type Parser struct {
	tokens []tok.Token
	offset int
}

func NewParser(tokens []tok.Token) *Parser {
	return &Parser{
		tokens: tokens,
		offset: 0,
	}
}

func (p *Parser) Parse() Node {
	return p.parseExpression()
}

func (p *Parser) parseExpression() Node {
	return p.parseEquality()
}

func (p *Parser) parseEquality() Node {
	node := p.parseComparison()
	for p.matchToken(tok.EqualEqual, tok.BangEqual) {
		op := p.currentToken()
		p.next()
		node = &BinaryNode{
			Operator: op,
			Left:     node,
			Right:    p.parseComparison(),
		}
	}
	return node
}

func (p *Parser) parseComparison() Node {
	node := p.parseTerm()
	for p.matchToken(tok.Greater, tok.GreaterEqual, tok.Less, tok.LessEqual) {
		op := p.currentToken()
		p.next()
		node = &BinaryNode{
			Operator: op,
			Left:     node,
			Right:    p.parseTerm(),
		}
	}
	return node
}

func (p *Parser) parseTerm() Node {
	node := p.parseFactor()
	for p.matchToken(tok.Minus, tok.Plus) {
		op := p.currentToken()
		p.next()
		node = &BinaryNode{
			Operator: op,
			Left:     node,
			Right:    p.parseFactor(),
		}
	}
	return node
}

func (p *Parser) parseFactor() Node {
	node := p.parseUrnary()
	for p.matchToken(tok.Slash, tok.Star) {
		op := p.currentToken()
		p.next()
		node = &BinaryNode{
			Operator: op,
			Left:     node,
			Right:    p.parseFactor(),
		}
	}
	return node
}

func (p *Parser) parseUrnary() Node {
	if p.matchToken(tok.Bang, tok.Minus) {
		op := p.currentToken()
		p.next()
		return &UrnaryNode{
			Operator: op,
			Right:    p.parseUrnary(),
		}
	}
	return p.parsePrimary()
}

func (p *Parser) parsePrimary() Node {
	c := p.currentToken()
	p.next()
	switch c.Type {
	case tok.False:
		return BooleanNode{Value: false}
	case tok.True:
		return BooleanNode{Value: true}
	case tok.Nil:
		return NilNode{}
	case tok.Number:
		return NumberNode{c.Literal.(float64)}
	case tok.String:
		return StringNode{c.Literal.(string)}
	case tok.LeftParen:
		node := p.parseExpression()
		p.next()
		// TODO: assert that the closing parenthesis is present instead of
		// simply skipping over it
		return &GroupingNode{Expression: node}
	case tok.EOF:
		return nil
	default:
		// TODO: return an error instead
		return nil
	}
}

// helper functions

func (p *Parser) next() {
	p.offset += 1
}

func (p *Parser) currentToken() tok.Token {
	return p.tokens[p.offset]
}

func (p *Parser) matchToken(types ...tok.Type) bool {
	c := p.currentToken()
	for _, t := range types {
		if c.Type == t {
			return true
		}
	}
	return false
}
