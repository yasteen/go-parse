package parse

import (
	"testing"
)

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

func testIsValidExpressionHelper(input []string, expected bool, t *testing.T) {
	if isValidExpression(input) != expected {
		str := ""
		for _, s := range input {
			str += s
		}
		t.Errorf("Test isValidExpression failed for %s.", str)
	}
}

func TestGetNextTokenWithString(t *testing.T) {
	testExpressionToTokens("sin(x)", []string{"sin", "(", "x", ")"}, t)
	testExpressionToTokens("x * log((5 +3) / 2)", []string{"x", "*", "log", "(", "(", "5", "+", "3", ")", "/", "2", ")"}, t)
}

func TestIsValidExpression(t *testing.T) {
	testIsValidExpressionHelper([]string{"x"}, true, t)
	testIsValidExpressionHelper([]string{"(", "x", ")"}, true, t)
	testIsValidExpressionHelper([]string{"sin", "x"}, true, t)
	testIsValidExpressionHelper([]string{"1", "+", "3", "*", "sin", "y"}, true, t)
	testIsValidExpressionHelper([]string{"(", "x", "^", "2", "-", "9", ")", "+", "x"}, true, t)

	testIsValidExpressionHelper([]string{"("}, false, t)
	testIsValidExpressionHelper([]string{")"}, false, t)
	testIsValidExpressionHelper([]string{"+"}, false, t)
	testIsValidExpressionHelper([]string{"sin"}, false, t)
	testIsValidExpressionHelper([]string{"sin", "+", "x"}, false, t)
	testIsValidExpressionHelper([]string{"4", "^"}, false, t)
}
