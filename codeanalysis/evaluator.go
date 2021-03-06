package codeanalysis

import (
	"fmt"

	"github.com/phlashdev/sherlock/codeanalysis/binding"
)

type Evaluator struct {
	root binding.BoundExpression
}

func NewEvaluator(root binding.BoundExpression) *Evaluator {
	return &Evaluator{root: root}
}

func (e *Evaluator) Evaluate() interface{} {
	return e.evaluateExpression(e.root)
}

func (e *Evaluator) evaluateExpression(node binding.BoundExpression) interface{} {
	if n, ok := node.(*binding.BoundLiteralExpression); ok {
		return n.Value()
	}

	if u, ok := node.(*binding.BoundUnaryExpression); ok {
		operand := e.evaluateExpression(u.Operand())
		op := u.Op()

		switch op.Kind() {
		case binding.Identity:
			return operand.(int)
		case binding.Negation:
			return -operand.(int)
		case binding.LogicalNegation:
			return !operand.(bool)
		default:
			panic(fmt.Sprintf("Unexcpected unary operator %v", op))
		}
	}

	if b, ok := node.(*binding.BoundBinaryExpression); ok {
		left := e.evaluateExpression(b.Left())
		right := e.evaluateExpression(b.Right())
		op := b.Op()

		switch op.Kind() {
		case binding.Addition:
			return left.(int) + right.(int)
		case binding.Subtraction:
			return left.(int) - right.(int)
		case binding.Multiplication:
			return left.(int) * right.(int)
		case binding.Division:
			return left.(int) / right.(int)
		case binding.LogicalAnd:
			return left.(bool) && right.(bool)
		case binding.LogicalOr:
			return left.(bool) || right.(bool)
		case binding.Equals:
			return left == right
		case binding.NotEquals:
			return left != right
		default:
			panic(fmt.Sprintf("Unexcpected binary operator %v", op))
		}
	}

	panic(fmt.Sprintf("Unexcpected node %v", node.Kind()))
}
