package binding

import (
	"reflect"
)

type BoundUnaryExpression struct {
	op      boundUnaryOperator
	operand BoundExpression
}

func newBoundUnaryExpression(op boundUnaryOperator, operand BoundExpression) *BoundUnaryExpression {
	return &BoundUnaryExpression{
		op:      op,
		operand: operand,
	}
}

func (b *BoundUnaryExpression) Kind() boundNodeKind {
	return UnaryExpression
}

func (b *BoundUnaryExpression) Type() reflect.Type {
	return b.operand.Type()
}

func (b *BoundUnaryExpression) Op() boundUnaryOperator {
	return b.op
}

func (b *BoundUnaryExpression) Operand() BoundExpression {
	return b.operand
}
