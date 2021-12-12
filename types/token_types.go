package types

import (
	"fmt"
	"unicode"
)

// keywords include Operators and SingleFunctions
type Keyword int
type TokenType int

const (
	Value TokenType = iota // A numerical value
	Variable
	LParen
	RParen
	Operator       // A math operator
	SingleFunction // A one-parameter function
)

type KeywordData struct {
	Symbol    string
	TokenType TokenType
	Apply     func(interface{}, interface{}) interface{}
}
type MathGroup struct {
	keywordMap         map[Keyword]KeywordData
	keywordStringMap   map[string]Keyword
	operatorPrecedence map[Keyword]int // For operators
	getValue           func(string) (interface{}, bool)
}

func New() *MathGroup {
	return &MathGroup{
		keywordMap:         map[Keyword]KeywordData{},
		keywordStringMap:   map[string]Keyword{},
		operatorPrecedence: map[Keyword]int{},
		getValue: func(s string) (interface{}, bool) {
			return 0, false
		},
	}
}

func NewMathGroup(
	keywordMap map[Keyword]KeywordData,
	keywordStringMap map[string]Keyword,
	operatorPrecedence map[Keyword]int,
	getValue func(string) (interface{}, bool),
) *MathGroup {
	return &MathGroup{
		keywordMap:         keywordMap,
		keywordStringMap:   keywordStringMap,
		operatorPrecedence: operatorPrecedence,
		getValue:           getValue,
	}
}

func (m *MathGroup) PushCurrentOp(prev Keyword, prevType TokenType, current Keyword) bool {
	return prevType != SingleFunction && m.operatorPrecedence[current] > m.operatorPrecedence[prev]
}

func (m *MathGroup) GetTokenType(keyword Keyword) (TokenType, bool) {
	if keywordData, exists := m.keywordMap[keyword]; exists {
		return keywordData.TokenType, true
	}
	return 0, false
}

func (m *MathGroup) GetKeywordString(s string) (Keyword, bool) {
	if keyword, exists := m.keywordStringMap[s]; exists {
		return keyword, true
	}
	return 0, false
}

func (m *MathGroup) TokenToString(keyword Keyword) (s string) {
	if _, exists := m.GetTokenType(keyword); exists {
		return m.keywordMap[keyword].Symbol
	}
	return ""
}

// Returns TokenType for given string. Keyword is added if applicable
func (m *MathGroup) StringToToken(s string) (TokenType, Keyword) {

	if s == "(" {
		return LParen, 0
	}
	if s == ")" {
		return RParen, 0
	}

	keyword, ok := m.GetKeywordString(s)
	tokenType, ok2 := m.GetTokenType(keyword)
	if ok && ok2 {
		return tokenType, keyword
	}

	// Is number
	if _, isValue := m.getValue(s); isValue {
		return Value, 0
	}

	// Is variable
	var first rune
	for _, r := range s {
		first = r
		break
	}
	if len(s) == 1 && unicode.IsLetter(first) {
		return Variable, 0
	}

	panic(fmt.Sprintf("Invalid token '%s' not recognized.", s))
}
