package codeanalysis

type NumberExpressionSyntax struct {
	numberToken SyntaxToken
}

func NewNumberExpressionSyntax(numberToken SyntaxToken) *NumberExpressionSyntax {
	return &NumberExpressionSyntax{numberToken: numberToken}
}

func (n *NumberExpressionSyntax) Kind() SyntaxKind {
	return NumberExpression
}

func (n *NumberExpressionSyntax) NumberToken() SyntaxToken {
	return n.numberToken
}

func (n *NumberExpressionSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{&n.numberToken}
}
