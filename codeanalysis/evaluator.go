package codeanalysis

import (
	"fmt"
)

type Evaluator struct {
	root ExpressionSyntax
}

func NewEvaluator(root ExpressionSyntax) *Evaluator {
	return &Evaluator{root: root}
}

func (e *Evaluator) Evaluate() int {
	return e.evaluateExpression(e.root)
}

func (e *Evaluator) evaluateExpression(node ExpressionSyntax) int {
	if n, ok := node.(*NumberExpressionSyntax); ok {
		return n.numberToken.value.(int)
	}

	if b, ok := node.(*BinaryExpressionSyntax); ok {
		left := e.evaluateExpression(b.Left())
		right := e.evaluateExpression(b.Right())

		operatorToken := b.OperatorToken()
		if operatorToken.Kind() == PlusToken {
			return left + right
		} else if operatorToken.Kind() == MinusToken {
			return left - right
		} else if operatorToken.Kind() == StarToken {
			return left * right
		} else if operatorToken.Kind() == SlashToken {
			return left / right
		} else {
			panic(fmt.Sprintf("Unexcpected binary operator %v", operatorToken.Kind()))
		}
	}

	if p, ok := node.(*ParenthesizedExpressionSyntax); ok {
		return e.evaluateExpression(p.Expression())
	}

	panic(fmt.Sprintf("Unexcpected node %v", node.Kind()))
}
