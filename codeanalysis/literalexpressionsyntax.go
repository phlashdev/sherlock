package codeanalysis

type LiteralExpressionSyntax struct {
	literalToken SyntaxToken
}

func NewNumberExpressionSyntax(literalToken SyntaxToken) *LiteralExpressionSyntax {
	return &LiteralExpressionSyntax{literalToken: literalToken}
}

func (n *LiteralExpressionSyntax) Kind() SyntaxKind {
	return NumberExpression
}

func (n *LiteralExpressionSyntax) LiteralToken() SyntaxToken {
	return n.literalToken
}

func (n *LiteralExpressionSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{&n.literalToken}
}
