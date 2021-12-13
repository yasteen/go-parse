package evaluate

import (
	"strconv"

	"github.com/karalabe/cookiejar/collections/stack"
	"github.com/yasteen/go-parse/parse"
	"github.com/yasteen/go-parse/types"
)

// For now, assume there's only one variable
// TODO: Finish implementation for any number of variables, and rethink output type.
func Evaluate(expression parse.ParsedExpression, domain types.Interval, m *types.MathGroup) []interface{} {
	result := []interface{}{}
	for current := domain.Start; current != nil; current = domain.NextValue(current) {
		result = append(result, evaluateOnce(expression, current, m))
	}
	return result
}

func evaluateOnce(expression parse.ParsedExpression, variable interface{}, m *types.MathGroup) interface{} {
	values := stack.New()
	for _, t := range expression {
		tokenType, keyword := m.StringToTokenType(t)
		var value interface{}
		switch tokenType {
		case types.Value:
			value, _ = strconv.ParseFloat(t, 64)
		case types.Variable:
			value = variable
		case types.Operator:
			val1 := values.Pop()
			val2 := values.Pop()
			value = m.ApplyKeyword(keyword, val1, val2)
		case types.SingleFunction:
			val := values.Pop()
			value = m.ApplyKeyword(keyword, val)
		default:
			panic("Invalid token.")
		}
		values.Push(value)
	}
	if values.Size() != 1 {
		panic("Expression is invalid.")
	}
	return values.Pop()
}
