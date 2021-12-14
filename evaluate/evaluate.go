// Package evaluate is used for evaluating expressions
package evaluate

import (
	"errors"

	"github.com/karalabe/cookiejar/collections/stack"
	"github.com/yasteen/go-parse/parse"
	"github.com/yasteen/go-parse/types"
)

// For now, assume there's only one variable
// TODO: Finish implementation for any number of variables, and rethink output type.

// Evaluate evaluates the given expression within the given domain.
func Evaluate(expression parse.ParsedExpression, domain types.Interval, m *types.MathGroup) ([]interface{}, error) {
	result := []interface{}{}
	for current := domain.Start; current != nil; current = domain.NextValue(current) {
		val, err := Once(expression, current, m)
		if err != nil {
			return result, err
		}
		result = append(result, val)
	}
	return result, nil
}

// Once evaluates the given expression using a given variable under the context of the given mathematical group.
func Once(expression parse.ParsedExpression, variable interface{}, m *types.MathGroup) (interface{}, error) {
	values := stack.New()
	for _, t := range expression {
		tokenType, keyword := m.StringToTokenType(t)
		var value interface{}
		switch tokenType {
		case types.Value:
			value, _ = m.GetValue(t)
		case types.Variable:
			value = variable
		case types.Operator:
			val2 := values.Pop()
			val1 := values.Pop()
			value = m.ApplyKeyword(keyword, val1, val2)
		case types.SingleFunction:
			val := values.Pop()
			value = m.ApplyKeyword(keyword, val)
		default:
			return nil, errors.New("invalid token")
		}
		values.Push(value)
	}
	if values.Size() != 1 {
		return nil, errors.New("expression is invalid")
	}
	return values.Pop(), nil
}
