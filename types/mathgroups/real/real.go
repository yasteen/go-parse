package real

import (
	"strconv"

	"github.com/yasteen/go-parse/types"
)

const (
	Number types.Keyword = iota
	Variable
	Add
	Subtract
	Multiply
	Divide
	Power
	Sin
	Cos
	Tan
	Log
	Exp
	LParen
	RParen
)

var realTokenMap = map[types.Keyword]types.KeywordData{
	Number:   {Symbol: "", TokenType: types.Value},
	Variable: {Symbol: "", TokenType: types.Value},
	Add:      {Symbol: "+", TokenType: types.Operator},
	Subtract: {Symbol: "-", TokenType: types.Operator},
	Multiply: {Symbol: "*", TokenType: types.Operator},
	Divide:   {Symbol: "/", TokenType: types.Operator},
	Power:    {Symbol: "^", TokenType: types.Operator},
	Sin:      {Symbol: "sin", TokenType: types.SingleFunction},
	Cos:      {Symbol: "cos", TokenType: types.SingleFunction},
	Tan:      {Symbol: "tan", TokenType: types.SingleFunction},
	Log:      {Symbol: "log", TokenType: types.SingleFunction},
	Exp:      {Symbol: "exp", TokenType: types.SingleFunction},
	LParen:   {Symbol: "(", TokenType: types.Paren},
	RParen:   {Symbol: ")", TokenType: types.Paren},
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
	"(":   LParen,
	")":   RParen,
}

var realOperatorPrecedence = map[types.Keyword]int{
	Add:      1,
	Subtract: 1,
	Multiply: 2,
	Divide:   2,
	Power:    3,
	Sin:      4,
	Cos:      4,
	Tan:      4,
	Log:      4,
	Exp:      4,
}

func getReal(s string) (interface{}, bool) {

	if num, err := strconv.ParseFloat(s, 64); err == nil {
		return num, true
	}
	return 0, false
}

var Real = types.NewMathGroup(realTokenMap, realStringToToken, realOperatorPrecedence)
