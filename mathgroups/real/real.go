// Package real is an implementation of the real number system with common operators and functions.
package real

import (
	"math"
	"strconv"

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

var realTokenMap = map[types.Keyword]types.KeywordData[float64]{
	Add: {Symbol: "+", TokenType: types.Operator,
		Apply: func(params ...float64) float64 {
			return params[0] + params[1]
		}},
	Subtract: {Symbol: "-", TokenType: types.Operator,
		Apply: func(params ...float64) float64 {
			return params[0] - params[1]
		}},
	Multiply: {Symbol: "*", TokenType: types.Operator,
		Apply: func(params ...float64) float64 {
			return params[0] * params[1]
		}},
	Divide: {Symbol: "/", TokenType: types.Operator,
		Apply: func(params ...float64) float64 {
			if params[1] == 0 {
				panic("Attempted division by 0")
			}
			return params[0] / params[1]
		}},
	Power: {Symbol: "^", TokenType: types.Operator,
		Apply: func(params ...float64) float64 {
			return math.Pow(params[0], params[1])
		}},
	Sin: {Symbol: "sin", TokenType: types.SingleFunction,
		Apply: func(params ...float64) float64 {
			return math.Sin(params[0])
		}},
	Cos: {Symbol: "cos", TokenType: types.SingleFunction,
		Apply: func(params ...float64) float64 {
			return math.Cos(params[0])
		}},
	Tan: {Symbol: "tan", TokenType: types.SingleFunction,
		Apply: func(params ...float64) float64 {
			return math.Tan(params[0])
		}},
	Log: {Symbol: "log", TokenType: types.SingleFunction,
		Apply: func(params ...float64) float64 {
			return math.Log(params[0])
		}},
	Exp: {Symbol: "exp", TokenType: types.SingleFunction,
		Apply: func(params ...float64) float64 {
			return math.Exp(params[0])
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

func getReal(s string) (float64, bool) {
	if num, err := strconv.ParseFloat(s, 64); err == nil {
		return num, true
	}
	return 0, false
}

// Real represents real number system (float64) and some defined operations/functions
var Real = types.NewMathGroup(realTokenMap, realStringToToken, realOperatorPrecedence, getReal)

// NewInterval constructs a new real interval.
func NewInterval(start float64, step float64, end float64) *types.Interval[float64] {
	if step == 0 || start > end {
		panic("Invalid interval")
	}
	return &types.Interval[float64]{
		Start: start,
		Step:  step,
		End:   end,
		Next: func(cur float64) (float64, bool) {
			next := cur + step
			if next > end {
				return end, true
			}
			return next, false
		},
	}
}
