package run

import (
	"github.com/yasteen/go-parse/evaluate"
	"github.com/yasteen/go-parse/parsexp"
	"github.com/yasteen/go-parse/types"
)

type RunnableMathGroup[T any] types.MathGroup[T]

func GetRunnableMathGroup[T any](group *types.MathGroup[T]) RunnableMathGroup[T] {
	return RunnableMathGroup[T](*group)
}

func (group RunnableMathGroup[T]) MapValues(expression string, interval types.Interval[T], varName string) ([]T, error) {
	g := types.MathGroup[T](group)
	parsedExpression, err := parsexp.Parse(expression, varName, &g)
	if err != nil {
		return nil, err
	}
	result, err := evaluate.Evaluate(parsedExpression, interval, &g)
	if err != nil {
		return nil, err
	}
	return result, nil
}
