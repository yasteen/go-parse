// Package parse is used for parsing expressions.
package parse

import (
	"errors"
	"strings"

	"github.com/karalabe/cookiejar/collections/stack"
	"github.com/yasteen/go-parse/types"
)

// ParsedExpression represents a valid equation in infix or postfix notation
type ParsedExpression []string

// GetNextTokenString returns a string with the next token.
func GetNextTokenString(expression string, index int, m *types.MathGroup) (tokenString string, nextIndex int) {
	tokenString = ""
	for ; index < len(expression); index++ {
		c := string(expression[index])
		if c == " " {
			if tokenString == "" {
				continue
			}
			for index < len(expression) && string(expression[index]) == " " {
				index++
			}
			break
		} else {
			tokenType, _ := m.StringToTokenType(c)
			if tokenType == types.RParen || tokenType == types.LParen || tokenType == types.Operator {
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

// IsLocallyValid verifies whether each token is valid with reference to its neighbours.
func IsLocallyValid(tokens []string, m *types.MathGroup) (bool, int) {
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

// Returns true if all tokens classified as a variable match the given variable namea variable match the given variable name.
func areTokensValid(tokens []string, variableName string, m *types.MathGroup) (bool, string) {
	for _, t := range tokens {
		tokenType, _ := m.StringToTokenType(t)
		if tokenType == types.Variable {
			if t != variableName {
				return false, t
			}
		}
	}
	return true, ""
}

// Converts an expression into a list of strings, split by token.
func parseExpression(expression string, m *types.MathGroup) (ParsedExpression, error) {
	tokens := ParsedExpression([]string{})
	for i := 0; i < len(expression); {
		tokenString, nextIndex := GetNextTokenString(expression, i, m)
		tokens = append(tokens, tokenString)
		i = nextIndex
	}
	return tokens, nil
}

// ToPostfix converts a ParsedExpression from infix notation to postfix notation.
// This change to postfix is useful for slightly optimizing speed in repeated calculations.
func ToPostfix(tokens ParsedExpression, m *types.MathGroup) (ParsedExpression, error) {
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
				if prevType == types.LParen || m.HasHigherPriority(keyword, prevKeyWord, prevType) {
					break
				}
				output = append(output, operations.Pop().(string))
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
				return nil, errors.New("expression has unmatched parentheses")
			}
		}
	}
	for operations.Size() > 0 {
		prevTokenString := operations.Pop().(string)
		prevType, _ := m.StringToTokenType(prevTokenString)
		if prevType == types.LParen {
			println("Bye")
			return nil, errors.New("expression has unmatched parentheses")
		}
		output = append(output, prevTokenString)
	}

	return output, nil
}

// Parse takes in the expression given, and parses it into in postfix form.
// This expression can be used in the evaluate module.
func Parse(expression string, variableName string, m *types.MathGroup) (ParsedExpression, error) {
	tokens, err := parseExpression(expression, m)
	if err != nil {
		return nil, err
	}
	if valid, t := areTokensValid(tokens, variableName, m); !valid {
		return nil, errors.New("Token " + t + " is not recognized.")
	}
	if isValid, i := IsLocallyValid(tokens, m); !isValid {
		return nil, errors.New("Expression is not valid\n" + expression + "\n" + strings.Repeat(" ", i) + "^")
	}
	finalExpr, err := ToPostfix(tokens, m)
	if err != nil {
		return nil, err
	}
	return finalExpr, nil
}
