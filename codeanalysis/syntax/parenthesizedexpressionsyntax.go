package syntax

type ParenthesizedExpressionSyntax struct {
	openParenthesisToken  SyntaxToken
	expression            ExpressionSyntax
	closeParenthesisToken SyntaxToken
}

func NewParenthesizedExpressionSyntax(openParenthesisToken SyntaxToken, expression ExpressionSyntax, closeParenthesisToken SyntaxToken) *ParenthesizedExpressionSyntax {
	return &ParenthesizedExpressionSyntax{
		openParenthesisToken:  openParenthesisToken,
		expression:            expression,
		closeParenthesisToken: closeParenthesisToken,
	}
}

func (p *ParenthesizedExpressionSyntax) Kind() SyntaxKind {
	return ParenthesizedExpression
}

func (p *ParenthesizedExpressionSyntax) OpenParenthesisToken() SyntaxToken {
	return p.openParenthesisToken
}

func (p *ParenthesizedExpressionSyntax) Expression() ExpressionSyntax {
	return p.expression
}

func (p *ParenthesizedExpressionSyntax) CloseParenthesisToken() SyntaxToken {
	return p.closeParenthesisToken
}

func (p *ParenthesizedExpressionSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{
		&p.openParenthesisToken,
		p.expression,
		&p.closeParenthesisToken,
	}
}
