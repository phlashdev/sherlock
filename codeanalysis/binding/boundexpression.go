package binding

import (
	"reflect"
)

type BoundExpression interface {
	boundNode
	Type() reflect.Type
}
