package real

import (
	"math"
	"strconv"

	"github.com/yasteen/go-parse/types"
)

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
			return params[0].(float64) + params[1].(float64)
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

var Real = types.NewMathGroup(realTokenMap, realStringToToken, realOperatorPrecedence, getReal)
