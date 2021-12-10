package binding

import (
	"fmt"
	"reflect"

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
	boundOperatorKind := bindUnaryOperatorKind(operatorToken.Kind(), boundOperand.Type())

	if boundOperatorKind == UnknownUnaryOperator {
		message := fmt.Sprintf("Unary operator %q is not defined for type %v", operatorToken.Text(), boundOperand.Type())
		b.diagnostics = append(b.diagnostics, message)
		return boundOperand
	}

	return newBoundUnaryExpression(boundOperatorKind, boundOperand)
}

func (b *binder) bindBinaryExpression(expressionSyntax syntax.BinaryExpressionSyntax) BoundExpression {
	boundLeft := b.BindExpression(expressionSyntax.Left())
	boundRight := b.BindExpression(expressionSyntax.Right())
	operatorToken := expressionSyntax.OperatorToken()
	boundOperatorKind := bindBinaryOperatorKind(operatorToken.Kind(), boundLeft.Type(), boundRight.Type())

	if boundOperatorKind == UnknownBinaryOperator {
		message := fmt.Sprintf("Binary operator %q is not defined for types %v and %v", operatorToken.Text(), boundLeft.Type(), boundRight.Type())
		b.diagnostics = append(b.diagnostics, message)
		return boundLeft
	}

	return newBoundBinaryExpression(boundLeft, boundOperatorKind, boundRight)
}

func bindUnaryOperatorKind(kind syntax.SyntaxKind, operandType reflect.Type) boundUnaryOperatorKind {
	if operandType.Kind() != reflect.Int {
		return UnknownUnaryOperator
	}

	switch kind {
	case syntax.PlusToken:
		return Identity
	case syntax.MinusToken:
		return Negation
	default:
		panic(fmt.Sprintf("Unexpected unary operator %v", kind))
	}
}

func bindBinaryOperatorKind(kind syntax.SyntaxKind, leftType reflect.Type, rightType reflect.Type) boundBinaryOperatorKind {
	if leftType.Kind() != reflect.Int || rightType.Kind() != reflect.Int {
		return UnknownBinaryOperator
	}

	switch kind {
	case syntax.PlusToken:
		return Addition
	case syntax.MinusToken:
		return Subtraction
	case syntax.StarToken:
		return Multiplication
	case syntax.SlashToken:
		return Division
	default:
		panic(fmt.Sprintf("Unexpected binary operator %v", kind))
	}
}
