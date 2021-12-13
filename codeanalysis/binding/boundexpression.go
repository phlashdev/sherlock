package binding

import (
	"reflect"
)

type BoundExpression interface {
	boundNode
	ResultType() reflect.Type
}
