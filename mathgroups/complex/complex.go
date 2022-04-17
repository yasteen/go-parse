// Package complex is an implementation of the complex number system with common operators and functions.
package complex

import (
	"math"
	"strconv"
	"strings"

	"github.com/yasteen/go-parse/types"
)

// Number is a struct representing the real and imaginary parts of a complex number
type Number struct {
	Re float64
	Im float64
}

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

func opAdd(params ...Number) Number {
	return Number{
		Re: params[0].Re + params[1].Re,
		Im: params[0].Im + params[1].Im,
	}
}
func opSubtract(params ...Number) Number {
	return Number{
		Re: params[0].Re - params[1].Re,
		Im: params[0].Im - params[1].Im,
	}
}
func opMultiply(params ...Number) Number {
	xr := params[0].Re
	xi := params[0].Im
	yr := params[1].Re
	yi := params[1].Im
	return Number{
		Re: xr*yr - xi*yi,
		Im: xr*yi + xi*yr,
	}
}
func opDivide(params ...Number) Number {
	xr := params[0].Re
	xi := params[0].Im
	yr := params[1].Re
	yi := params[1].Im
	if yr == 0 && yi == 0 {
		panic("Attempted division by 0")
	}
	return Number{
		Re: (xr*yr + xi*yi) / (yr*yr + yi*yi),
		Im: (xi*yr - xr*yi) / (yr*yr + yi*yi),
	}
}
func fnLog(params ...Number) Number {
	re := params[0].Re
	im := params[0].Im
	mod, arg := cartesianToPolar(re, im)
	return Number{
		Re: math.Log(mod),
		Im: arg,
	}
}
func fnExp(params ...Number) Number {
	re := params[0].Re
	im := params[0].Im
	exp := math.Exp(re)
	return Number{
		Re: exp * math.Cos(im),
		Im: exp * math.Sin(im),
	}
}
func fnSin(params ...Number) Number {
	re := params[0].Re
	im := params[1].Im
	first := fnExp(Number{-im, re})  // e^(iz)
	second := fnExp(Number{im, -re}) // e^(-iz)
	return Number{
		Re: (first.Im - second.Im) / 2,
		Im: (second.Re - first.Re) / 2,
	}
}
func fnCos(params ...Number) Number {
	re := params[0].Re
	im := params[1].Im
	first := fnExp(Number{-im, re})  // e^(iz)
	second := fnExp(Number{im, -re}) // e^(-iz)
	return Number{
		Re: (first.Re - second.Re) / 2,
		Im: (first.Im - second.Im) / 2,
	}
}

var complexTokenMap = map[types.Keyword]types.KeywordData[Number]{
	Add:      {Symbol: "+", TokenType: types.Operator, Apply: opAdd},
	Subtract: {Symbol: "-", TokenType: types.Operator, Apply: opSubtract},
	Multiply: {Symbol: "*", TokenType: types.Operator, Apply: opMultiply},
	Divide:   {Symbol: "/", TokenType: types.Operator, Apply: opDivide},
	Power: {Symbol: "^", TokenType: types.Operator,
		Apply: func(params ...Number) Number {
			return fnExp(opMultiply(params[1], fnLog(params[0])))
		},
	},
	Sin: {Symbol: "sin", TokenType: types.SingleFunction, Apply: fnSin},
	Cos: {Symbol: "cos", TokenType: types.SingleFunction, Apply: fnCos},
	Tan: {Symbol: "tan", TokenType: types.SingleFunction,
		Apply: func(params ...Number) Number {
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

func getComplex(s string) (Number, bool) {
	if strings.Contains(s, "_") {
		nums := strings.Split(s, "_")
		if len(nums) == 2 {
			re, err := strconv.ParseFloat(nums[0], 64)
			im, err2 := strconv.ParseFloat(nums[1], 64)
			if err == nil && err2 == nil {
				return Number{re, im}, true
			}
		}
		return Number{0, 0}, false
	}
	if num, err := strconv.ParseFloat(s, 64); err == nil {
		return Number{num, 0}, true
	}
	if s == "i" {
		return Number{0, 1}, true
	}
	if len(s) != 0 && s[len(s)-1] == 'i' {
		if num, err := strconv.ParseFloat(s[:len(s)-1], 64); err == nil {
			return Number{0, num}, true
		}
	}
	return Number{0, 0}, false
}

// Complex represents the complex number system (float64, float64) and some defined operations/functions
var Complex = types.NewMathGroup(complexTokenMap, complexStringToToken, complexOperatorPrecedence, getComplex)

// NewComplexInterval constructs a new complex interval (top right to bottom left corner in Cartesian form)
func NewComplexInterval(start Number, step Number, end Number) *types.Interval[Number] {
	if step.Re == 0 || (start.Re < end.Re || start.Im < end.Im) {
		panic("Invalid Interval")
	}
	return &types.Interval[Number]{
		Start: start,
		Step:  step,
		End:   end,
		Next: func(cur Number) (Number, bool) {
			nextRe := cur.Re + step.Re
			nextIm := cur.Im
			if nextRe > end.Re {
				nextRe = start.Re
				nextIm += step.Re
				if nextIm > end.Im {
					return end, true
				}
				return Number{nextRe, nextIm}, false
			}
			return Number{nextRe, nextIm}, false
		},
	}
}
