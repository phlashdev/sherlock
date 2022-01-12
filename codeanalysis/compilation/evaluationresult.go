package compilation

import "github.com/phlashdev/sherlock/codeanalysis/diagnostic"

type EvaluationResult struct {
	diagnostics []diagnostic.Diagnostic
	value       interface{}
}

func NewEvaluationResult(diagnostics []diagnostic.Diagnostic, value interface{}) *EvaluationResult {
	return &EvaluationResult{
		diagnostics: diagnostics,
		value:       value,
	}
}

func (e *EvaluationResult) Diagnostics() []diagnostic.Diagnostic {
	return e.diagnostics
}

func (e *EvaluationResult) Value() interface{} {
	return e.value
}
