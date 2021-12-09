package binding

import (
	"reflect"
)

type BoundUnaryExpression struct {
	operatorKind boundUnaryOperatorKind
	operand      BoundExpression
}

func newBoundUnaryExpression(operatorKind boundUnaryOperatorKind, operand BoundExpression) *BoundUnaryExpression {
	return &BoundUnaryExpression{
		operatorKind: operatorKind,
		operand:      operand,
	}
}

func (b *BoundUnaryExpression) Kind() boundNodeKind {
	return UnaryExpression
}

func (b *BoundUnaryExpression) Type() reflect.Type {
	return b.operand.Type()
}

func (b *BoundUnaryExpression) OperatorKind() boundUnaryOperatorKind {
	return b.operatorKind
}

func (b *BoundUnaryExpression) Operand() BoundExpression {
	return b.operand
}
