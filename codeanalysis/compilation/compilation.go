package compilation

import (
	"github.com/phlashdev/sherlock/codeanalysis/binding"
	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

type Compilation struct {
	syntax syntax.SyntaxTree
}

func NewCompilation(syntax syntax.SyntaxTree) *Compilation {
	return &Compilation{
		syntax: syntax,
	}
}

func (c *Compilation) Syntax() syntax.SyntaxTree {
	return c.syntax
}

func (c *Compilation) Evaluate() EvaluationResult {
	binder := binding.NewBinder()
	boundExpression := binder.BindExpression(c.syntax.Root())

	diagnostics := append(c.syntax.Diagnostics(), binder.Diagnostics()...)
	if len(diagnostics) > 0 {
		return *NewEvaluationResult(diagnostics, nil)
	}

	evaluator := NewEvaluator(boundExpression)
	value := evaluator.Evaluate()
	return *NewEvaluationResult(diagnostics, value)
}
