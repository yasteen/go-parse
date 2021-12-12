package parse

import (
	"strings"

	"github.com/karalabe/cookiejar/collections/stack"
	"github.com/yasteen/go-parse/types"
)

// A valid equation in postfix notation
type ParsedExpression []string

// Returns a string with the next token.
func getNextTokenString(expression string, index int, m *types.MathGroup) (tokenString string, nextIndex int) {
	tokenString = ""
	for ; index < len(expression); index++ {
		c := string(expression[index])
		if c == " " {
			if tokenString == "" {
				continue
			}
			for index < len(expression) && c == " " {
				index++
			}
			break
		} else {
			tokenType, _ := m.StringToTokenType(c)
			if tokenType == types.RParen || tokenType == types.LParen {
				if tokenString == "" {
					tokenString = c
					index++
				}
				break
			}
		}
		tokenString += c
	}
	return tokenString, index
}

func isValidExpression(tokens []string, m *types.MathGroup) (bool, int) {
	currentCharLength := 0
	if len(tokens) == 0 {
		return true, currentCharLength
	}
	prevTokenType, _ := m.StringToTokenType(tokens[0])
	if len(tokens) == 1 {
		return prevTokenType == types.Value || prevTokenType == types.Variable, currentCharLength
	}
	if prevTokenType == types.Operator || prevTokenType == types.RParen {
		return false, currentCharLength
	}
	currentCharLength += len(tokens[0])

	for i := 1; i < len(tokens); i++ {
		tokenType, _ := m.StringToTokenType(tokens[i])
		// Whether the previous token could be the end of an expression
		prevIsEndable := prevTokenType == types.Value || prevTokenType == types.Variable || prevTokenType == types.RParen
		// Whether the current token can take place after an endable previous token
		currFollows := tokenType == types.Operator || tokenType == types.RParen
		if (prevIsEndable && !currFollows) || (!prevIsEndable && currFollows) {
			return false, currentCharLength
		}

		currentCharLength += len(tokens[i])

		prevTokenType = tokenType
	}

	validEnd := prevTokenType == types.Value || prevTokenType == types.Variable || prevTokenType == types.RParen
	if !validEnd {
		currentCharLength -= len(tokens[len(tokens)-1])
	}

	return validEnd, currentCharLength
}

func assertVariableTokensAreValid(tokens []string, variableName string, m *types.MathGroup) {
	for _, t := range tokens {
		tokenType, _ := m.StringToTokenType(t)
		if tokenType == types.Variable {
			if t != variableName {
				panic("Token " + t + " is not recognized.")
			}
		}
	}
}

func parseExpression(expression string, m *types.MathGroup) ParsedExpression {
	tokens := ParsedExpression([]string{})
	for i := 0; i < len(expression); {
		tokenString, nextIndex := getNextTokenString(expression, i, m)
		tokens = append(tokens, tokenString)
		i = nextIndex
	}
	return tokens
}

// Change to postfix for slight increase in speed for repeated calculations
func toPostfix(tokens ParsedExpression, m *types.MathGroup) ParsedExpression {
	output := ParsedExpression([]string{})
	operations := stack.New()

	for _, t := range tokens {
		tokenType, keyword := m.StringToTokenType(t)
		switch tokenType {
		case types.Value:
			fallthrough
		case types.Variable:
			output = append(output, t)
		case types.SingleFunction:
			operations.Push(t)
		case types.Operator:
			for operations.Size() > 0 {
				prevType, prevKeyWord := m.StringToTokenType(operations.Top().(string))
				if prevType == types.LParen || m.PushCurrentOp(prevKeyWord, prevType, keyword) {
					break
				}
				output = append(output, m.KeywordToString(operations.Pop().(types.Keyword)))
			}
			operations.Push(t)
		case types.LParen:
			operations.Push(t)
		case types.RParen:
			foundMatchingParen := false
			for operations.Size() > 0 {
				prevTokenString := operations.Pop().(string)
				prevType, _ := m.StringToTokenType(prevTokenString)
				if prevType == types.LParen {
					foundMatchingParen = true
					break
				}
				output = append(output, prevTokenString)
			}
			if !foundMatchingParen {
				panic("Expression has unmatched parentheses.")
			}
		}
	}
	for operations.Size() > 0 {
		prevTokenString := operations.Pop().(string)
		prevType, _ := m.StringToTokenType(prevTokenString)
		if prevType == types.LParen {
			panic("Expression has unmatched parentheses.")
		}
		output = append(output, prevTokenString)
	}

	return output
}

func Parse(expression string, variableName string, m *types.MathGroup) ParsedExpression {
	tokens := parseExpression(expression, m)
	assertVariableTokensAreValid(tokens, variableName, m)
	if isValid, i := isValidExpression(tokens, m); !isValid {
		panic("Expression is not valid\n" + expression + strings.Repeat(" ", i) + "^")
	}
	return toPostfix(tokens, m)
}
