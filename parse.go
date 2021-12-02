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

func Parse(equation string, variables map[string]float64) {
	tokens := []string{}
	for i := 0; i < len(equation); {
		token, nextIndex := getNextTokenString(equation, i)
		tokens = append(tokens, token)
		i = nextIndex
	}
}
