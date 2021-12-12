package parse

import (
	"strings"

	"github.com/karalabe/cookiejar/collections/stack"
)

// A valid equation in postfix notation
type ParsedExpression []string

// Returns a string with the next token.
func getNextTokenString(expression string, index int) (tokenString string, nextIndex int) {
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
		} else if c == TokenToString(RParen) || c == TokenToString(LParen) {
			if tokenString == "" {
				tokenString = c
				index++
			}
			break
		}
		tokenString += c
	}
	return tokenString, index
}

func isValidExpression(tokens []string) (bool, int) {
	currentCharLength := 0
	if len(tokens) == 0 {
		return true, currentCharLength
	}
	prevToken, prevTokenType := StringToToken(tokens[0])
	if len(tokens) == 1 {
		return prevTokenType == Value, currentCharLength
	}
	if prevTokenType == Operator || prevToken == RParen {
		return false, currentCharLength
	}
	currentCharLength += len(tokens[0])

	for i := 1; i < len(tokens); i++ {
		token, tokenType := StringToToken(tokens[i])
		// Whether the previous token could be the end of an expression
		prevIsEndable := prevTokenType == Value || prevToken == RParen
		// Whether the current token can take place after an endable previous token
		currFollows := tokenType == Operator || token == RParen
		if (prevIsEndable && !currFollows) || (!prevIsEndable && currFollows) {
			return false, currentCharLength
		}

		currentCharLength += len(tokens[i])

		prevToken = token
		prevTokenType = tokenType
	}

	validEnd := prevTokenType == Value || prevToken == RParen
	if !validEnd {
		currentCharLength -= len(tokens[len(tokens)-1])
	}

	return validEnd, currentCharLength
}

func verifyValidVariables(tokens []string, variables map[string]struct{}) {
	for _, t := range tokens {
		token, _ := StringToToken(t)
		if token == Variable {
			if _, ok := variables[t]; !ok {
				panic("Variable " + t + " is not recognized.")
			}
		}
	}
}

func parseExpression(expression string) ParsedExpression {
	tokens := ParsedExpression([]string{})
	for i := 0; i < len(expression); {
		token, nextIndex := getNextTokenString(expression, i)
		tokens = append(tokens, token)
		i = nextIndex
	}
	return tokens
}

// Change to postfix for slight increase in speed for repeated calculations
func toPostfix(tokens ParsedExpression) ParsedExpression {
	output := ParsedExpression([]string{})
	operations := stack.New()

	for _, t := range tokens {
		token, tokenType := StringToToken(t)
		switch tokenType {
		case Value:
			output = append(output, t)
		case SingleFunction:
			operations.Push(token)
		case Operator:
			for operations.Size() > 0 {
				op := operations.Top().(Token)
				if op == LParen || PushCurrentOp(op, token) {
					break
				}
				output = append(output, TokenToString(operations.Pop().(Token)))
			}
			operations.Push(token)
		case Paren:
			if token == LParen {
				operations.Push(token)
			} else {
				foundMatchingParen := false
				for operations.Size() > 0 {
					op := operations.Pop().(Token)
					if op == LParen {
						foundMatchingParen = true
						break
					}
					output = append(output, TokenToString(op))
				}
				if !foundMatchingParen {
					panic("Expression has unmatched parentheses.")
				}
			}
		}
	}
	for operations.Size() > 0 {
		op := operations.Pop().(Token)
		if op == LParen {
			panic("Expression has unmatched parentheses.")
		}
		output = append(output, TokenToString(op))
	}

	return output
}

func Parse(expression string, variables map[string]struct{}) ParsedExpression {
	tokens := parseExpression(expression)
	if isValid, i := isValidExpression(tokens); !isValid {
		panic("Expression is not valid\n" + expression + strings.Repeat(" ", i) + "^")
	}
	verifyValidVariables(tokens, variables)
	return toPostfix(tokens)
}
