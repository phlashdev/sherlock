package codeanalysis

import (
	"fmt"
)

type Parser struct {
	tokens      []SyntaxToken
	position    int
	diagnostics []string
}

func NewParser(text string) *Parser {
	tokens := make([]SyntaxToken, 0, 10)

	lexer := NewLexer(text)
	for {
		token := lexer.NextToken()

		if token.Kind() != WhitespaceToken && token.Kind() != BadToken {
			tokens = append(tokens, token)
		}

		if token.Kind() == EndOfFileToken {
			break
		}
	}

	parser := Parser{
		tokens:      tokens,
		diagnostics: lexer.Diagnostics(),
	}
	return &parser
}

func (p *Parser) Diagnostics() []string {
	return p.diagnostics
}

func (p *Parser) peek(offset int) SyntaxToken {
	index := p.position + offset
	if index >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[index]
}

func (p *Parser) current() SyntaxToken {
	return p.peek(0)
}

func (p *Parser) nextToken() SyntaxToken {
	current := p.current()
	p.position++
	return current
}

func (p *Parser) match(kind SyntaxKind) SyntaxToken {
	current := p.current()
	if current.Kind() == kind {
		return p.nextToken()
	}

	p.diagnostics = append(p.diagnostics, fmt.Sprintf("ERROR: Unexpected token <%v>, expected <%v>", current.kind, kind))
	return *NewSyntaxToken(kind, current.Position(), "", nil)
}

func (p *Parser) parseExpression() ExpressionSyntax {
	return p.parseTerm()
}

func (p *Parser) Parse() SyntaxTree {
	expression := p.parseTerm()
	endOfFileToken := p.match(EndOfFileToken)
	return *NewSyntaxTree(p.diagnostics, expression, endOfFileToken)
}

func (p *Parser) parseTerm() ExpressionSyntax {
	left := p.parseFactor()

	for {
		current := p.current()
		if !(current.Kind() == PlusToken || current.Kind() == MinusToken || current.Kind() == StarToken || current.Kind() == SlashToken) {
			break
		}

		operatorToken := p.nextToken()
		right := p.parseFactor()
		left = NewBinaryExpressionSyntax(left, operatorToken, right)
	}

	return left
}

func (p *Parser) parseFactor() ExpressionSyntax {
	left := p.parsePrimaryExpression()

	for {
		current := p.current()
		if !(current.Kind() == StarToken || current.Kind() == SlashToken) {
			break
		}

		operatorToken := p.nextToken()
		right := p.parsePrimaryExpression()
		left = NewBinaryExpressionSyntax(left, operatorToken, right)
	}

	return left
}

func (p *Parser) parsePrimaryExpression() ExpressionSyntax {
	if p.current().kind == OpenParenthesisToken {
		left := p.nextToken()
		expression := p.parseExpression()
		right := p.match(CloseParenthesisToken)
		return NewParenthesizedExpressionSyntax(left, expression, right)
	}

	numberToken := p.match(NumberToken)
	return NewNumberExpressionSyntax(numberToken)
}
