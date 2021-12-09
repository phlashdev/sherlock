package binding

import (
	"reflect"
)

type BoundBinaryExpression struct {
	left         BoundExpression
	operatorKind boundBinaryOperatorKind
	right        BoundExpression
}

func newBoundBinaryExpression(left BoundExpression, operatorKind boundBinaryOperatorKind, right BoundExpression) *BoundBinaryExpression {
	return &BoundBinaryExpression{
		left:         left,
		operatorKind: operatorKind,
		right:        right,
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

func (b *BoundBinaryExpression) OperatorKind() boundBinaryOperatorKind {
	return b.operatorKind
}

func (b *BoundBinaryExpression) Right() BoundExpression {
	return b.right
}
