// Package real is an implementation of the real number system with common operators and functions.
package real

import (
	"math"
	"strconv"

	"github.com/yasteen/go-parse/evaluate"
	"github.com/yasteen/go-parse/parse"
	"github.com/yasteen/go-parse/types"
)

// Common functions and operations defined for the real group
const (
	Add types.Keyword = iota
	Subtract
	Multiply
	Divide
	Power
	Sin
	Cos
	Tan
	Log
	Exp
)

var realTokenMap = map[types.Keyword]types.KeywordData{
	Add: {Symbol: "+", TokenType: types.Operator,
		Apply: func(params ...interface{}) interface{} {
			return params[0].(float64) + params[1].(float64)
		}},
	Subtract: {Symbol: "-", TokenType: types.Operator,
		Apply: func(params ...interface{}) interface{} {
			return params[0].(float64) - params[1].(float64)
		}},
	Multiply: {Symbol: "*", TokenType: types.Operator,
		Apply: func(params ...interface{}) interface{} {
			return params[0].(float64) * params[1].(float64)
		}},
	Divide: {Symbol: "/", TokenType: types.Operator,
		Apply: func(params ...interface{}) interface{} {
			if params[1].(float64) == 0 {
				// TODO: Handle division by 0
				panic("Attempted division by 0")
			}
			return params[0].(float64) / params[1].(float64)
		}},
	Power: {Symbol: "^", TokenType: types.Operator,
		Apply: func(params ...interface{}) interface{} {
			return math.Pow(params[0].(float64), params[1].(float64))
		}},
	Sin: {Symbol: "sin", TokenType: types.SingleFunction,
		Apply: func(params ...interface{}) interface{} {
			return math.Sin(params[0].(float64))
		}},
	Cos: {Symbol: "cos", TokenType: types.SingleFunction,
		Apply: func(params ...interface{}) interface{} {
			return math.Cos(params[0].(float64))
		}},
	Tan: {Symbol: "tan", TokenType: types.SingleFunction,
		Apply: func(params ...interface{}) interface{} {
			return math.Tan(params[0].(float64))
		}},
	Log: {Symbol: "log", TokenType: types.SingleFunction,
		Apply: func(params ...interface{}) interface{} {
			return math.Log(params[0].(float64))
		}},
	Exp: {Symbol: "exp", TokenType: types.SingleFunction,
		Apply: func(params ...interface{}) interface{} {
			return math.Exp(params[0].(float64))
		}},
}

var realStringToToken = map[string]types.Keyword{
	"+":   Add,
	"-":   Subtract,
	"*":   Multiply,
	"/":   Divide,
	"^":   Power,
	"sin": Sin,
	"cos": Cos,
	"tan": Tan,
	"log": Log,
	"exp": Exp,
}

var realOperatorPrecedence = map[types.Keyword]int{
	Add:      1,
	Subtract: 1,
	Multiply: 2,
	Divide:   2,
	Power:    3,
}

func getReal(s string) (interface{}, bool) {
	if num, err := strconv.ParseFloat(s, 64); err == nil {
		return num, true
	}
	return 0, false
}

// Real represents real number system (float64) and some defined operations/functions
var Real = types.NewMathGroup(realTokenMap, realStringToToken, realOperatorPrecedence, getReal)

// NewRealInterval constructs a new real interval.
func NewRealInterval(start float64, step float64, end float64) *types.Interval {
	if step == 0 || start > end {
		panic("Invalid interval")
	}
	return &types.Interval{
		Start: start,
		Step:  step,
		End:   end,
		Next: func(cur interface{}) interface{} {
			next := cur.(float64) + step
			if next > end {
				return nil
			}
			return next
		},
	}
}

// MapValues evaluates an expression for all real values specified by the interval.
func MapValues(expression string, interval types.Interval, varName string) ([]float64, error) {
	parsedExpression, err := parse.Parse(expression, varName, Real)
	if err != nil {
		return nil, err
	}
	result, err := evaluate.Evaluate(parsedExpression, interval, Real)
	if err != nil {
		return nil, err
	}
	ret := []float64{}
	for _, val := range result {
		ret = append(ret, val.(float64))
	}
	return ret, nil
}
