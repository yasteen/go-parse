package evaluate

import (
	"strconv"

	"github.com/karalabe/cookiejar/collections/stack"
	"github.com/yasteen/go-parse/parse"
	"github.com/yasteen/go-parse/types"
)

type Interval struct {
	min  float64
	step float64
	max  float64
}

func Evaluate(expression parse.ParsedExpression, variables map[string]Interval) {

}

func evaluateOnce(expression parse.ParsedExpression, variables map[string]float64) {
	numbers := stack.New()
	for _, t := range expression {
		token, tokenType := types.StringToToken(t)
		var number float64
		switch tokenType {
		case types.Value:
			if token == types.Number {
				number, _ = strconv.ParseFloat(t, 64)
			} else {
				number = variables[t]
			}
		case types.Operator:
			val1 := numbers.Pop()
			val2 := numbers.Pop()
		case types.SingleFunction:
		default:
			panic("Invalid token.")
		}
		numbers.Push(number)
	}
}
