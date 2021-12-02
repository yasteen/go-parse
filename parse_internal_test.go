package parse

import "testing"

func testExpressionToTokens(input string, expected []string, t *testing.T) {
	curStringIndex := 0
	for i := 0; i < len(input); {
		var tokenString string
		tokenString, i = getNextTokenString(input, i)
		if curStringIndex >= len(expected) {
			t.Errorf("Test for '%s' produced too many tokens.", input)
		}
		if tokenString != expected[curStringIndex] {
			t.Errorf("Unexpected token. Expected '%s', got '%s'", expected[curStringIndex], tokenString)
		}
		curStringIndex++
	}
}
func TestGetNextTokenWithString(t *testing.T) {
	testExpressionToTokens("sin(x)", []string{"sin", "(", "x", ")"}, t)
	testExpressionToTokens("x * log((5 +3) / 2)", []string{"x", "*", "log", "(", "(", "5", "+", "3", ")", "/", "2", ")"}, t)
}
