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
			tokenType, _ := m.StringToToken(c)
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
	prevTokenType, _ := m.StringToToken(tokens[0])
	if len(tokens) == 1 {
		return prevTokenType == types.Value, currentCharLength
	}
	if prevTokenType == types.Operator || prevTokenType == types.RParen {
		return false, currentCharLength
	}
	currentCharLength += len(tokens[0])

	for i := 1; i < len(tokens); i++ {
		tokenType, _ := m.StringToToken(tokens[i])
		// Whether the previous token could be the end of an expression
		prevIsEndable := prevTokenType == types.Value || prevTokenType == types.RParen
		// Whether the current token can take place after an endable previous token
		currFollows := tokenType == types.Operator || tokenType == types.RParen
		if (prevIsEndable && !currFollows) || (!prevIsEndable && currFollows) {
			return false, currentCharLength
		}

		currentCharLength += len(tokens[i])

		prevTokenType = tokenType
	}

	validEnd := prevTokenType == types.Value || prevTokenType == types.RParen
	if !validEnd {
		currentCharLength -= len(tokens[len(tokens)-1])
	}

	return validEnd, currentCharLength
}

func verifyValidVariables(tokens []string, variables map[string]struct{}, m *types.MathGroup) {
	for _, t := range tokens {
		tokenType, _ := m.StringToToken(t)
		if tokenType == types.Variable {
			if _, ok := variables[t]; !ok {
				panic("Variable " + t + " is not recognized.")
			}
		}
	}
}

func parseExpression(expression string, m *types.MathGroup) ParsedExpression {
	tokens := ParsedExpression([]string{})
	for i := 0; i < len(expression); {
		token, nextIndex := getNextTokenString(expression, i, m)
		tokens = append(tokens, token)
		i = nextIndex
	}
	return tokens
}

// Change to postfix for slight increase in speed for repeated calculations
func toPostfix(tokens ParsedExpression, m *types.MathGroup) ParsedExpression {
	output := ParsedExpression([]string{})
	operations := stack.New()

	for _, t := range tokens {
		tokenType, token := m.StringToToken(t)
		switch tokenType {
		case types.Value:
			output = append(output, t)
		case types.SingleFunction:
			operations.Push(t)
		case types.Operator:
			for operations.Size() > 0 {
				prevType, prevKeyWord := m.StringToToken(operations.Top().(string))
				if prevType == types.LParen || m.PushCurrentOp(prevKeyWord, prevType, token) {
					break
				}
				output = append(output, m.TokenToString(operations.Pop().(types.Keyword)))
			}
			operations.Push(t)
		case types.LParen:
			operations.Push(t)
		case types.RParen:
			foundMatchingParen := false
			for operations.Size() > 0 {
				prevTokenString := operations.Pop().(string)
				prevType, _ := m.StringToToken(prevTokenString)
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
		prevType, _ := m.StringToToken(prevTokenString)
		if prevType == types.LParen {
			panic("Expression has unmatched parentheses.")
		}
		output = append(output, prevTokenString)
	}

	return output
}

func Parse(expression string, variables map[string]struct{}, m *types.MathGroup) ParsedExpression {
	tokens := parseExpression(expression, m)
	if isValid, i := isValidExpression(tokens, m); !isValid {
		panic("Expression is not valid\n" + expression + strings.Repeat(" ", i) + "^")
	}
	verifyValidVariables(tokens, variables, m)
	return toPostfix(tokens, m)
}
