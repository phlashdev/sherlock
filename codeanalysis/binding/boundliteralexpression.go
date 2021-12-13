package binding

import (
	"reflect"
)

type BoundLiteralExpression struct {
	value interface{}
}

func newBoundLiteralExpression(value interface{}) *BoundLiteralExpression {
	return &BoundLiteralExpression{value: value}
}

func (b *BoundLiteralExpression) Kind() boundNodeKind {
	return LiteralExpression
}

func (b *BoundLiteralExpression) ResultType() reflect.Type {
	return reflect.TypeOf(b.value)
}

func (b *BoundLiteralExpression) Value() interface{} {
	return b.value
}
