package evaluate_test

import (
	"testing"

	"github.com/yasteen/go-parse/evaluate"
	"github.com/yasteen/go-parse/mathgroups/real"
	"github.com/yasteen/go-parse/parsexp"
)

func TestEvaluateOnce(t *testing.T) {
	testEvaluateOnceHelper(
		parsexp.ParsedExpression{"x"},
		43.2,
		43.2,
		t,
	)
	testEvaluateOnceHelper(
		parsexp.ParsedExpression{"x", "4", "-"},
		43.2,
		39.2,
		t,
	)
	testEvaluateOnceHelper(
		parsexp.ParsedExpression{"x", "x", "2", "-", "*", "3", "^"},
		4,
		512,
		t,
	)

}

func testEvaluateOnceHelper(expression parsexp.ParsedExpression, variable float64, expected float64, t *testing.T) {
	value, err := evaluate.Once(expression, variable, real.Real)

	if err != nil {
		t.Error(err)
	}

	if value != expected {
		t.Error("EvaluateOnce failed. Expected:", expected, "Result:", value)
	}
}
