// Package types consists of constants and types representing data relating to a mathematical group/system used for parsing/evaluating
package types

// Keyword consists of Operators and SingleFunctions
type Keyword int

// TokenType represents the type of a token.
type TokenType int

// The possible token types
const (
	Value          TokenType = iota // A literal value
	Variable                        // A variable acting as a placeholder for a literal
	LParen                          // Left parenthesis
	RParen                          // Right parenthesis
	Operator                        // A math operator
	SingleFunction                  // A one-parameter function
)

// KeywordData represents data relating to a Keyword.
// TokenType == Operator: Expect 2 arguments to Apply
// TokenType == SingleFunction: Expect 1 argument to Apply
type KeywordData struct {
	Symbol    string
	TokenType TokenType
	Apply     func(...interface{}) interface{}
}

// MathGroup is a data structure representing a mathematical system.
type MathGroup struct {
	keywordMap         map[Keyword]KeywordData
	keywordStringMap   map[string]Keyword
	operatorPrecedence map[Keyword]int // For operators
	GetValue           func(string) (interface{}, bool)
}

// TODO: Add verification for the three maps

// NewMathGroup is a constructor for MathGroup
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
		GetValue:           getValue,
	}
}

// HasHigherPriority returns true if the current operator has a higher priority.
func (m *MathGroup) HasHigherPriority(current Keyword, ref Keyword, refType TokenType) bool {
	return refType != SingleFunction && m.operatorPrecedence[current] > m.operatorPrecedence[ref]
}

// KeywordToString converts a keyword into its corresponding string.
func (m *MathGroup) KeywordToString(keyword Keyword) (s string) {
	if keywordData, exists := m.keywordMap[keyword]; exists {
		return keywordData.Symbol
	}
	return ""
}

// ApplyKeyword applies an operation or function on the given arguments.
func (m *MathGroup) ApplyKeyword(keyword Keyword, args ...interface{}) interface{} {
	return m.keywordMap[keyword].Apply(args...)
}

// StringToTokenType returns the TokenType for a given string. Keyword is added if applicable.
// Note: If nothing is matched, the default type returned is Variable.
func (m *MathGroup) StringToTokenType(s string) (TokenType, Keyword) {

	if s == "(" {
		return LParen, 0
	}
	if s == ")" {
		return RParen, 0
	}

	keyword, ok := m.keywordStringMap[s]
	keywordData, ok2 := m.keywordMap[keyword]
	if ok && ok2 {
		return keywordData.TokenType, keyword
	}

	// Is a valid value
	if _, isValue := m.GetValue(s); isValue {
		return Value, 0
	}

	return Variable, 0
}
