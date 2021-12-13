package syntax

type BinaryExpressionSyntax struct {
	left          ExpressionSyntax
	operatorToken SyntaxToken
	right         ExpressionSyntax
}

func NewBinaryExpressionSyntax(left ExpressionSyntax, operatorToken SyntaxToken, right ExpressionSyntax) *BinaryExpressionSyntax {
	return &BinaryExpressionSyntax{
		left:          left,
		operatorToken: operatorToken,
		right:         right,
	}
}

func (b *BinaryExpressionSyntax) Kind() SyntaxKind {
	return BinaryExpression
}

func (b *BinaryExpressionSyntax) Left() ExpressionSyntax {
	return b.left
}

func (b *BinaryExpressionSyntax) OperatorToken() SyntaxToken {
	return b.operatorToken
}

func (b *BinaryExpressionSyntax) Right() ExpressionSyntax {
	return b.right
}

func (b *BinaryExpressionSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{
		b.left,
		&b.operatorToken,
		b.right,
	}
}
