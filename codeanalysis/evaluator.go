package codeanalysis

import (
	"fmt"

	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

type Evaluator struct {
	root syntax.ExpressionSyntax
}

func NewEvaluator(root syntax.ExpressionSyntax) *Evaluator {
	return &Evaluator{root: root}
}

func (e *Evaluator) Evaluate() int {
	return e.evaluateExpression(e.root)
}

func (e *Evaluator) evaluateExpression(node syntax.ExpressionSyntax) int {
	if n, ok := node.(*syntax.LiteralExpressionSyntax); ok {
		token := n.LiteralToken()
		return token.Value().(int)
	}

	if u, ok := node.(*syntax.UnaryExpressionSyntax); ok {
		operand := e.evaluateExpression(u.Operand())

		operatorToken := u.OperatorToken()
		if operatorToken.Kind() == syntax.PlusToken {
			return operand
		} else if operatorToken.Kind() == syntax.MinusToken {
			return -operand
		} else {
			panic(fmt.Sprintf("Unexcpected unary operator %v", operatorToken.Kind()))
		}
	}

	if b, ok := node.(*syntax.BinaryExpressionSyntax); ok {
		left := e.evaluateExpression(b.Left())
		right := e.evaluateExpression(b.Right())

		operatorToken := b.OperatorToken()
		if operatorToken.Kind() == syntax.PlusToken {
			return left + right
		} else if operatorToken.Kind() == syntax.MinusToken {
			return left - right
		} else if operatorToken.Kind() == syntax.StarToken {
			return left * right
		} else if operatorToken.Kind() == syntax.SlashToken {
			return left / right
		} else {
			panic(fmt.Sprintf("Unexcpected binary operator %v", operatorToken.Kind()))
		}
	}

	if p, ok := node.(*syntax.ParenthesizedExpressionSyntax); ok {
		return e.evaluateExpression(p.Expression())
	}

	panic(fmt.Sprintf("Unexcpected node %v", node.Kind()))
}
