// An implementation of the complex number system with common operators and functions
package complex

import (
	"math"
	"strconv"
	"strings"

	"github.com/yasteen/go-parse/evaluate"
	"github.com/yasteen/go-parse/parse"
	"github.com/yasteen/go-parse/types"
)

type ComplexNumber struct {
	Re float64
	Im float64
}

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

// Helper function to convert from Cartesian to Polar form
func cartesianToPolar(re float64, im float64) (mod float64, arg float64) {
	if re == 0 && im == 0 {
		panic("Arg is undefined.")
	}
	mod = re*re + im*im
	if re > 0 || im != 0 {
		arg = 2 * math.Atan(im/math.Sqrt(mod)+re)
	} else {
		arg = math.Pi
	}
	return mod, arg
}

func opAdd(params ...interface{}) interface{} {
	return ComplexNumber{
		Re: params[0].(ComplexNumber).Re + params[1].(ComplexNumber).Re,
		Im: params[0].(ComplexNumber).Im + params[1].(ComplexNumber).Im,
	}
}
func opSubtract(params ...interface{}) interface{} {
	return ComplexNumber{
		Re: params[0].(ComplexNumber).Re - params[1].(ComplexNumber).Re,
		Im: params[0].(ComplexNumber).Im - params[1].(ComplexNumber).Im,
	}
}
func opMultiply(params ...interface{}) interface{} {
	xr := params[0].(ComplexNumber).Re
	xi := params[0].(ComplexNumber).Im
	yr := params[1].(ComplexNumber).Re
	yi := params[1].(ComplexNumber).Im
	return ComplexNumber{
		Re: xr*yr - xi*yi,
		Im: xr*yi + xi*yr,
	}
}
func opDivide(params ...interface{}) interface{} {
	xr := params[0].(ComplexNumber).Re
	xi := params[0].(ComplexNumber).Im
	yr := params[1].(ComplexNumber).Re
	yi := params[1].(ComplexNumber).Im
	if yr == 0 && yi == 0 {
		// TODO: Handle division by 0
		panic("Attempted division by 0")
	}
	return ComplexNumber{
		Re: (xr*yr + xi*yi) / (yr*yr + yi*yi),
		Im: (xi*yr - xr*yi) / (yr*yr + yi*yi),
	}
}
func fnLog(params ...interface{}) interface{} {
	re := params[0].(ComplexNumber).Re
	im := params[0].(ComplexNumber).Im
	mod, arg := cartesianToPolar(re, im)
	return ComplexNumber{
		Re: math.Log(mod),
		Im: arg,
	}
}
func fnExp(params ...interface{}) interface{} {
	re := params[0].(ComplexNumber).Re
	im := params[0].(ComplexNumber).Im
	exp := math.Exp(re)
	return ComplexNumber{
		Re: exp * math.Cos(im),
		Im: exp * math.Sin(im),
	}
}
func fnSin(params ...interface{}) interface{} {
	re := params[0].(ComplexNumber).Re
	im := params[1].(ComplexNumber).Im
	first := fnExp(ComplexNumber{-im, re}).(ComplexNumber)  // e^(iz)
	second := fnExp(ComplexNumber{im, -re}).(ComplexNumber) // e^(-iz)
	return ComplexNumber{
		Re: (first.Im - second.Im) / 2,
		Im: (second.Re - first.Re) / 2,
	}
}
func fnCos(params ...interface{}) interface{} {
	re := params[0].(ComplexNumber).Re
	im := params[1].(ComplexNumber).Im
	first := fnExp(ComplexNumber{-im, re}).(ComplexNumber)  // e^(iz)
	second := fnExp(ComplexNumber{im, -re}).(ComplexNumber) // e^(-iz)
	return ComplexNumber{
		Re: (first.Re - second.Re) / 2,
		Im: (first.Im - second.Im) / 2,
	}
}

var complexTokenMap = map[types.Keyword]types.KeywordData{
	Add:      {Symbol: "+", TokenType: types.Operator, Apply: opAdd},
	Subtract: {Symbol: "-", TokenType: types.Operator, Apply: opSubtract},
	Multiply: {Symbol: "*", TokenType: types.Operator, Apply: opMultiply},
	Divide:   {Symbol: "/", TokenType: types.Operator, Apply: opDivide},
	Power: {Symbol: "^", TokenType: types.Operator,
		Apply: func(params ...interface{}) interface{} {
			return fnExp(opMultiply(params[1], fnLog(params[0])))
		},
	},
	Sin: {Symbol: "sin", TokenType: types.SingleFunction, Apply: fnSin},
	Cos: {Symbol: "cos", TokenType: types.SingleFunction, Apply: fnCos},
	Tan: {Symbol: "tan", TokenType: types.SingleFunction,
		Apply: func(params ...interface{}) interface{} {
			return opDivide(fnSin(params[0]), fnCos(params[0]))
		},
	},
	Log: {Symbol: "log", TokenType: types.SingleFunction, Apply: fnLog},
	Exp: {Symbol: "exp", TokenType: types.SingleFunction, Apply: fnExp},
}

var complexStringToToken = map[string]types.Keyword{
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

var complexOperatorPrecedence = map[types.Keyword]int{
	Add:      1,
	Subtract: 1,
	Multiply: 2,
	Divide:   2,
	Power:    3,
}

func getComplex(s string) (interface{}, bool) {
	if strings.Contains(s, "_") {
		nums := strings.Split(s, "_")
		if len(nums) == 2 {
			re, err := strconv.ParseFloat(nums[0], 64)
			im, err2 := strconv.ParseFloat(nums[1], 64)
			if err == nil && err2 == nil {
				return ComplexNumber{re, im}, true
			}
		}
		return 0, false
	}
	if num, err := strconv.ParseFloat(s, 64); err == nil {
		return ComplexNumber{num, 0}, true
	}
	if s == "i" {
		return ComplexNumber{0, 1}, true
	}
	if len(s) != 0 && s[len(s)-1] == 'i' {
		if num, err := strconv.ParseFloat(s[:len(s)-1], 64); err == nil {
			return ComplexNumber{0, num}, true
		}
	}
	return 0, false
}

// The complex number system (float64, float64) and some defined operations/functions
var Complex = types.NewMathGroup(complexTokenMap, complexStringToToken, complexOperatorPrecedence, getComplex)

// Constructs a new complex interval (top right to bottom left corner in Cartesian form)
func NewComplexInterval(start ComplexNumber, step float64, end ComplexNumber) *types.Interval {
	if step == 0 || (start.Re < end.Re || start.Im < end.Im) {
		panic("Invalid Interval")
	}
	return &types.Interval{
		Start: start,
		Step:  step,
		End:   end,
		Next: func(cur interface{}) interface{} {
			nextRe := cur.(ComplexNumber).Re + step
			nextIm := cur.(ComplexNumber).Im
			if nextRe > end.Re {
				nextRe = start.Re
				nextIm += step
				if nextIm > end.Im {
					return nil
				}
				return ComplexNumber{nextRe, nextIm}
			}
			return ComplexNumber{nextRe, nextIm}
		},
	}
}

// Evaluates an expression for all complex values specified by the interval.
func MapValues(expression string, interval types.Interval, varName string) ([]ComplexNumber, error) {
	parsedExpression, err := parse.Parse(expression, varName, Complex)
	if err != nil {
		return nil, err
	}
	result, err := evaluate.Evaluate(parsedExpression, interval, Complex)
	if err != nil {
		return nil, err
	}
	ret := []ComplexNumber{}
	for _, val := range result {
		ret = append(ret, val.(ComplexNumber))
	}
	return ret, nil
}
