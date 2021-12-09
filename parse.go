package parse

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

func isValidExpression(tokens []string) bool {
	if len(tokens) == 0 {
		return true
	}
	prevToken, prevTokenType := StringToToken(tokens[0])
	if len(tokens) == 1 {
		return prevTokenType == Value
	}
	if prevTokenType == Operator || prevToken == RParen {
		return false
	}

	for i := 1; i < len(tokens); i++ {
		token, tokenType := StringToToken(tokens[i])
		// Whether the previous token could be the end of an expression
		prevIsEndable := prevTokenType == Value || prevToken == RParen
		// Whether the current token can take place after an endable previous token
		currFollows := tokenType == Operator || token == RParen
		if (prevIsEndable && !currFollows) || (!prevIsEndable && currFollows) {
			return false
		}
		prevToken = token
		prevTokenType = tokenType
	}

	return prevTokenType == Value || prevToken == RParen
}

func Parse(equation string, variables map[string]float64) {
	tokens := []string{}
	for i := 0; i < len(equation); {
		token, nextIndex := getNextTokenString(equation, i)
		tokens = append(tokens, token)
		i = nextIndex
	}
}
