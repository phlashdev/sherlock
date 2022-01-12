package binding

import (
	"fmt"

	"github.com/phlashdev/sherlock/codeanalysis/diagnostic"
	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

type binder struct {
	diagnostics []diagnostic.Diagnostic
}

func NewBinder() *binder {
	return &binder{}
}

func (b *binder) Diagnostics() []diagnostic.Diagnostic {
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
		report := reportUndefinedUnaryOperator(operatorToken.Span(), operatorToken.Text(), boundOperand.ResultType())
		b.diagnostics = append(b.diagnostics, report)
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
		report := reportUndefinedBinaryOperator(operatorToken.Span(), operatorToken.Text(), boundLeft.ResultType(), boundRight.ResultType())
		b.diagnostics = append(b.diagnostics, report)
		return boundLeft
	}

	return newBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}
