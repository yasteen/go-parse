// Package run is used as a wrapper to parse/evaluate
// an expression within a given math system
package run

import (
	"github.com/yasteen/go-parse/evaluate"
	"github.com/yasteen/go-parse/parsexp"
	"github.com/yasteen/go-parse/types"
)

// RunnableMathGroup is a MathGroup
type RunnableMathGroup[T any] types.MathGroup[T]

// GetRunnableMathGroup converts a MathGroup pointer into a RunnnbleMathGroup
func GetRunnableMathGroup[T any](group *types.MathGroup[T]) RunnableMathGroup[T] {
	return RunnableMathGroup[T](*group)
}

// MapValues parses and evaluates an expression within an interval in a MathGroup
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
