package codeanalysis

type UnaryExpressionSyntax struct {
	operatorToken SyntaxToken
	operand       ExpressionSyntax
}

func NewUnaryExpressionSyntax(operatorToken SyntaxToken, operand ExpressionSyntax) *UnaryExpressionSyntax {
	return &UnaryExpressionSyntax{
		operatorToken: operatorToken,
		operand:       operand,
	}
}

func (u *UnaryExpressionSyntax) Kind() SyntaxKind {
	return UnaryExpression
}

func (u *UnaryExpressionSyntax) OperatorToken() SyntaxToken {
	return u.operatorToken
}

func (u *UnaryExpressionSyntax) Operand() ExpressionSyntax {
	return u.operand
}

func (u *UnaryExpressionSyntax) GetChildren() []SyntaxNode {
	return []SyntaxNode{
		&u.operatorToken,
		u.operand,
	}
}
