package parse

import "strings"

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

func equationToTokens(expression string) []string {
	tokens := []string{}
	for i := 0; i < len(expression); {
		token, nextIndex := getNextTokenString(expression, i)
		tokens = append(tokens, token)
		i = nextIndex
	}
	return tokens
}

func Parse(expression string, variables map[string]float64) {
	tokens := equationToTokens(expression)
	if isValid, i := isValidExpression(tokens); !isValid {
		panic("Expression is not valid\n" + expression + strings.Repeat(" ", i) + "^")
	}

}
