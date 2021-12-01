package codeanalysis

import (
	"fmt"
)

type parser struct {
	tokens      []SyntaxToken
	position    int
	diagnostics []string
}

func NewParser(text string) *parser {
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

	parser := parser{
		tokens:      tokens,
		diagnostics: lexer.Diagnostics(),
	}
	return &parser
}

func (p *parser) Diagnostics() []string {
	return p.diagnostics
}

func (p *parser) peek(offset int) SyntaxToken {
	index := p.position + offset
	if index >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[index]
}

func (p *parser) current() SyntaxToken {
	return p.peek(0)
}

func (p *parser) nextToken() SyntaxToken {
	current := p.current()
	p.position++
	return current
}

func (p *parser) matchToken(kind SyntaxKind) SyntaxToken {
	current := p.current()
	if current.Kind() == kind {
		return p.nextToken()
	}

	p.diagnostics = append(p.diagnostics, fmt.Sprintf("ERROR: Unexpected token <%v>, expected <%v>", current.kind, kind))
	return *NewSyntaxToken(kind, current.Position(), "", nil)
}

func (p *parser) Parse() SyntaxTree {
	expression := p.parseExpression()
	endOfFileToken := p.matchToken(EndOfFileToken)
	return *NewSyntaxTree(p.diagnostics, expression, endOfFileToken)
}

func (p *parser) parseExpression() ExpressionSyntax {
	return p.parseTerm()
}

func (p *parser) parseTerm() ExpressionSyntax {
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

func (p *parser) parseFactor() ExpressionSyntax {
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

func (p *parser) parsePrimaryExpression() ExpressionSyntax {
	if p.current().kind == OpenParenthesisToken {
		left := p.nextToken()
		expression := p.parseExpression()
		right := p.matchToken(CloseParenthesisToken)
		return NewParenthesizedExpressionSyntax(left, expression, right)
	}

	numberToken := p.matchToken(NumberToken)
	return NewNumberExpressionSyntax(numberToken)
}
