package parse

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType int

const (
	Value          TokenType = iota // A numerical value
	Paren                           // A left or right parenthesis
	Operator                        // A math operator
	SingleFunction                  // A one-parameter function
)

type Token int

const (
	Number Token = iota
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

type _token struct {
	_symbol string
	_type   TokenType
}

var _tokenMap = map[Token]_token{
	Number:   {"", Value},
	Variable: {"", Value},
	Add:      {"+", Operator},
	Subtract: {"-", Operator},
	Multiply: {"*", Operator},
	Divide:   {"/", Operator},
	Power:    {"^", Operator},
	Sin:      {"sin", SingleFunction},
	Cos:      {"cos", SingleFunction},
	Tan:      {"tan", SingleFunction},
	Log:      {"log", SingleFunction},
	Exp:      {"exp", SingleFunction},
	LParen:   {"(", Paren},
	RParen:   {")", Paren},
}

// Missing Value token types because they are unknown
var _stringToToken = map[string]Token{
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

func GetTokenType(token Token) TokenType {
	return _tokenMap[token]._type
}

func TokenToString(token Token) (s string) {
	return _tokenMap[token]._symbol
}

func StringToToken(s string) (token Token, t TokenType) {
	token, ok := _stringToToken[s]

	// Is predefined token
	if ok {
		return token, _tokenMap[token]._type
	}

	// Is number
	if _, err := strconv.Atoi(s); err == nil || s == "e" {
		return Number, Value
	}

	// Is variable
	var first rune
	for _, r := range s {
		first = r
		break
	}
	if len(s) == 1 && unicode.IsLetter(first) {
		return Variable, Value
	}

	panic(fmt.Sprintf("Invalid token '%s' not recognized.", s))
}
