package syntax

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
		token := lexer.Lex()

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
	expression := p.parseExpression(0)
	endOfFileToken := p.matchToken(EndOfFileToken)
	return *NewSyntaxTree(p.diagnostics, expression, endOfFileToken)
}

func (p *parser) parseExpression(parentPrecedence int) ExpressionSyntax {
	var left ExpressionSyntax
	unaryOperatorPrecedence := getUnaryOperatorPrecedence(p.current().kind)
	if unaryOperatorPrecedence != 0 && unaryOperatorPrecedence >= parentPrecedence {
		operatorToken := p.nextToken()
		operand := p.parseExpression(unaryOperatorPrecedence)
		left = NewUnaryExpressionSyntax(operatorToken, operand)
	} else {
		left = p.parsePrimaryExpression()
	}

	for {
		precedence := getBinaryOperatorPrecedence(p.current().kind)
		if precedence == 0 || precedence <= parentPrecedence {
			break
		}

		operatorToken := p.nextToken()
		right := p.parseExpression(precedence)
		left = NewBinaryExpressionSyntax(left, operatorToken, right)
	}

	return left
}

func (p *parser) parsePrimaryExpression() ExpressionSyntax {
	currentKind := p.current().kind

	switch currentKind {
	case OpenParenthesisToken:
		left := p.nextToken()
		expression := p.parseExpression(0)
		right := p.matchToken(CloseParenthesisToken)
		return NewParenthesizedExpressionSyntax(left, expression, right)
	case TrueKeyword, FalseKeyword:
		keywordToken := p.nextToken()
		value := keywordToken.Kind() == TrueKeyword
		return NewLiteralExpressionSyntaxWithValue(keywordToken, value)
	default:
		numberToken := p.matchToken(NumberToken)
		return NewLiteralExpressionSyntax(numberToken)
	}
}
