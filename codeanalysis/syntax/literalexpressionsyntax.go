package syntax

type LiteralExpressionSyntax struct {
	literalToken SyntaxToken
	value        interface{}
}

func NewLiteralExpressionSyntax(literalToken SyntaxToken) *LiteralExpressionSyntax {
	return &LiteralExpressionSyntax{
		literalToken: literalToken,
		value:        literalToken.Value(),
	}
}

func NewLiteralExpressionSyntaxWithValue(literalToken SyntaxToken, value interface{}) *LiteralExpressionSyntax {
	return &LiteralExpressionSyntax{
		literalToken: literalToken,
		value:        value,
	}
}

func (l *LiteralExpressionSyntax) Kind() SyntaxKind {
	return LiteralExpression
}

func (l *LiteralExpressionSyntax) LiteralToken() SyntaxToken {
	return l.literalToken
}

func (l *LiteralExpressionSyntax) Value() interface{} {
	return l.value
}

func (l *LiteralExpressionSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{&l.literalToken}
}
