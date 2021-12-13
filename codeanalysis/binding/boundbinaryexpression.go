package binding

import (
	"reflect"
)

type BoundBinaryExpression struct {
	left  BoundExpression
	op    boundBinaryOperator
	right BoundExpression
}

func newBoundBinaryExpression(left BoundExpression, op boundBinaryOperator, right BoundExpression) *BoundBinaryExpression {
	return &BoundBinaryExpression{
		left:  left,
		op:    op,
		right: right,
	}
}

func (b *BoundBinaryExpression) Kind() boundNodeKind {
	return BinaryExpression
}

func (b *BoundBinaryExpression) Type() reflect.Type {
	return b.left.Type()
}

func (b *BoundBinaryExpression) Left() BoundExpression {
	return b.left
}

func (b *BoundBinaryExpression) Op() boundBinaryOperator {
	return b.op
}

func (b *BoundBinaryExpression) Right() BoundExpression {
	return b.right
}
