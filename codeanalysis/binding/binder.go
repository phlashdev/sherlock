package binding

import (
	"fmt"

	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

type binder struct {
	diagnostics []string
}

func NewBinder() *binder {
	return &binder{}
}

func (b *binder) Diagnostics() []string {
	return b.diagnostics
}

func (b *binder) BindExpression(expressionSyntax syntax.ExpressionSyntax) BoundExpression {
	switch expressionSyntax.Kind() {
	case syntax.LiteralExpression:
		literal, _ := expressionSyntax.(*syntax.LiteralExpressionSyntax)
		return b.bindLiteralExpression(*literal)
	case syntax.UnaryExpression:
		unary, _ := expressionSyntax.(*syntax.UnaryExpressionSyntax)
		return b.bindUnaryExpression(*unary)
	case syntax.BinaryExpression:
		binary, _ := expressionSyntax.(*syntax.BinaryExpressionSyntax)
		return b.bindBinaryExpression(*binary)
	case syntax.ParenthesizedExpression:
		parenthesized, _ := expressionSyntax.(*syntax.ParenthesizedExpressionSyntax)
		return b.BindExpression(parenthesized.Expression())
	default:
		panic(fmt.Sprintf("Unexpected syntax %v", expressionSyntax.Kind()))
	}
}

func (b *binder) bindLiteralExpression(expressionSyntax syntax.LiteralExpressionSyntax) BoundExpression {
	if expressionSyntax.Value() == nil {
		return newBoundLiteralExpression(0)
	}
	return newBoundLiteralExpression(expressionSyntax.Value())
}

func (b *binder) bindUnaryExpression(expressionSyntax syntax.UnaryExpressionSyntax) BoundExpression {
	boundOperand := b.BindExpression(expressionSyntax.Operand())
	operatorToken := expressionSyntax.OperatorToken()

	boundOperator, err := bindUnaryOperator(operatorToken.Kind(), boundOperand.ResultType())
	if err != nil {
		message := fmt.Sprintf("Unary operator %q is not defined for type %v", operatorToken.Text(), boundOperand.ResultType())
		b.diagnostics = append(b.diagnostics, message)
		return boundOperand
	}

	return newBoundUnaryExpression(boundOperator, boundOperand)
}

func (b *binder) bindBinaryExpression(expressionSyntax syntax.BinaryExpressionSyntax) BoundExpression {
	boundLeft := b.BindExpression(expressionSyntax.Left())
	boundRight := b.BindExpression(expressionSyntax.Right())
	operatorToken := expressionSyntax.OperatorToken()

	boundOperator, err := bindBinaryOperator(operatorToken.Kind(), boundLeft.ResultType(), boundRight.ResultType())
	if err != nil {
		message := fmt.Sprintf("Binary operator %q is not defined for types %v and %v", operatorToken.Text(), boundLeft.ResultType(), boundRight.ResultType())
		b.diagnostics = append(b.diagnostics, message)
		return boundLeft
	}

	return newBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}
