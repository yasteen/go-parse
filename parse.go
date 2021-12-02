package main

// Returns a string with the next token.
func getNextTokenString(equation string, index int) (tokenString string, nextIndex int) {
	tokenString = ""
	for ; index < len(equation); index++ {
		c := string(equation[index])
		if c == " " {
			continue
		} else if c == TokenToString(RParen) || c == TokenToString(LParen) {
			if tokenString == "" {
				tokenString = c
			}
			index++
			break
		}
		tokenString += c
	}
	return tokenString, nextIndex
}

func Parse(equation string, variables map[string]float64) {
	tokens := []string{}
	for i := 0; i < len(equation); {
		token, nextIndex := getNextTokenString(equation, i)
		tokens = append(tokens, token)
		i = nextIndex
	}
}
