package types

// keywords include Operators and SingleFunctions
type Keyword int

type TokenType int

const (
	Value    TokenType = iota // A literal value
	Variable                  // A variable acting as a placeholder for a literal
	LParen
	RParen
	Operator       // A math operator
	SingleFunction // A one-parameter function
)

// Data relating to a Keyword.
// TokenType == Operator: Expect 2 arguments to Apply
// TokenType == SingleFunction: Expect 1 argument to Apply
type KeywordData struct {
	Symbol    string
	TokenType TokenType
	Apply     func(...interface{}) interface{}
}

// Data structure representing a mathematical system.
type MathGroup struct {
	keywordMap         map[Keyword]KeywordData
	keywordStringMap   map[string]Keyword
	operatorPrecedence map[Keyword]int // For operators
	getValue           func(string) (interface{}, bool)
}

// Constructor for MathGroup
// TODO: Add verification for the three maps
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

// Operation precedence helper function for conversion from infix to postfix notation
func (m *MathGroup) PushCurrentOp(prev Keyword, prevType TokenType, current Keyword) bool {
	return prevType != SingleFunction && m.operatorPrecedence[current] > m.operatorPrecedence[prev]
}

func (m *MathGroup) KeywordToString(keyword Keyword) (s string) {
	if keywordData, exists := m.keywordMap[keyword]; exists {
		return keywordData.Symbol
	}
	return ""
}
func (m *MathGroup) ApplyKeyword(keyword Keyword, args ...interface{}) interface{} {
	return m.keywordMap[keyword].Apply(args...)
}

// Returns TokenType for given string. Keyword is added if applicable.
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
	if _, isValue := m.getValue(s); isValue {
		return Value, 0
	}

	return Variable, 0
}
