package parse

import (
	"strconv"
	"strings"
	"testing"

	"github.com/yasteen/go-parse/types/mathgroups/real"
)

func testGetNextTokenStringHelper(input string, expected []string, t *testing.T) {
	curStringIndex := 0
	for i := 0; i < len(input); {
		var tokenString string
		tokenString, i = getNextTokenString(input, i, real.Real)
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
	testGetNextTokenStringHelper("sin(x)", []string{"sin", "(", "x", ")"}, t)
	testGetNextTokenStringHelper("x * log((5 +3) / 2)", []string{"x", "*", "log", "(", "(", "5", "+", "3", ")", "/", "2", ")"}, t)
}

func testIsValidExpressionHelper(input []string, expected bool, t *testing.T) {
	if isValid, i := isValidExpression(input, real.Real); isValid != expected {
		str := ""
		for _, s := range input {
			str += s
		}
		errorString := "Test isValidExpression failed:\n" + str
		errorString += ", expected:" + strconv.FormatBool(expected) + "\n"
		if expected {
			errorString += strings.Repeat(" ", i) + "^"
		}
		t.Error(errorString)
	}
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

func testToPostfixHelper(input []string, expected []string, unmatchedParen bool, t *testing.T) {
	defer func() {
		recover()
	}()
	output := toPostfix(input, real.Real)

	if unmatchedParen {
		// Should not reach here if we expect a panic
		t.Error("Failed to detect unmatched parentheses.")
	} else {
		if len(output) != len(expected) {
			t.Error("Output and expected output length do not match. Expected", expected, "Produced", output)
		}
		for i := 0; i < len(expected); i++ {
			if output[i] != expected[i] {
				t.Error("Output and expected output do not match at index", i, "- Expected", expected[i], "Produced", output[i])
			}
		}
	}
}

func TestToPostfixHelper(t *testing.T) {
	testToPostfixHelper([]string{"x", "+", "4"}, []string{"x", "4", "+"}, false, t)
	testToPostfixHelper([]string{"(", "(", "x", "+", "4", ")", "^", "5", ")"}, []string{"x", "4", "+", "5", "^"}, false, t)

	testToPostfixHelper([]string{"(", "(", "x", "+", "4", "^", "5", ")"}, []string{}, true, t)
	testToPostfixHelper([]string{"(", "x", "+", "4"}, []string{}, true, t)
}
