// Package evaluate is used for evaluating expressions
package evaluate

import (
	"errors"

	"github.com/karalabe/cookiejar/collections/stack"
	"github.com/yasteen/go-parse/parsexp"
	"github.com/yasteen/go-parse/types"
)

// For now, assume there's only one variable
// TODO: Finish implementation for any number of variables, and rethink output type.

// Evaluate evaluates the given expression within the given domain.
func Evaluate[T any](expression parsexp.ParsedExpression, domain types.Interval[T], m *types.MathGroup[T]) ([]T, error) {
	result := []T{}
	done := false
	for current := domain.Start; !done; current, done = domain.Next(current) {
		val, err := Once(expression, current, m)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

// Once evaluates the given expression using a given variable under the context of the given mathematical group.
func Once[T any](expression parsexp.ParsedExpression, variable T, m *types.MathGroup[T]) (T, error) {
	values := stack.New()
	for _, t := range expression {
		tokenType, keyword := m.StringToTokenType(t)
		var value T
		switch tokenType {
		case types.Value:
			value, _ = m.GetValue(t)
		case types.Variable:
			value = variable
		case types.Operator:
			val2 := values.Pop().(T)
			val1 := values.Pop().(T)
			value = m.ApplyKeyword(keyword, val1, val2)
		case types.SingleFunction:
			val := values.Pop().(T)
			value = m.ApplyKeyword(keyword, val)
		default:
			return variable, errors.New("invalid token")
		}
		values.Push(value)
	}
	if values.Size() != 1 {
		return variable, errors.New("expression is invalid")
	}
	return values.Pop().(T), nil
}
