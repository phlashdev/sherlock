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
		operand := e.evaluateExpression(u.Operand()).(int)

		switch u.OperatorKind() {
		case binding.Identity:
			return operand
		case binding.Negation:
			return -operand
		default:
			panic(fmt.Sprintf("Unexcpected unary operator %v", u.OperatorKind()))
		}
	}

	if b, ok := node.(*binding.BoundBinaryExpression); ok {
		left := e.evaluateExpression(b.Left()).(int)
		right := e.evaluateExpression(b.Right()).(int)

		switch b.OperatorKind() {
		case binding.Addition:
			return left + right
		case binding.Subtraction:
			return left - right
		case binding.Multiplication:
			return left * right
		case binding.Division:
			return left / right
		default:
			panic(fmt.Sprintf("Unexcpected binary operator %v", b.OperatorKind()))
		}
	}

	panic(fmt.Sprintf("Unexcpected node %v", node.Kind()))
}
