package parse_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/yasteen/go-parse/parse"
	"github.com/yasteen/go-parse/types/mathgroups/real"
)

func testGetNextTokenStringHelper(input string, expected []string, t *testing.T) {
	curStringIndex := 0
	for i := 0; i < len(input); {
		var tokenString string
		tokenString, i = parse.GetNextTokenString(input, i, real.Real)
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

func testIsLocallyValidHelper(input []string, expected bool, t *testing.T) {
	if isValid, i := parse.IsLocallyValid(input, real.Real); isValid != expected {
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
func TestIsLocallyValid(t *testing.T) {
	testIsLocallyValidHelper([]string{"x"}, true, t)
	testIsLocallyValidHelper([]string{"(", "x", ")"}, true, t)
	testIsLocallyValidHelper([]string{"sin", "x"}, true, t)
	testIsLocallyValidHelper([]string{"1", "+", "3", "*", "sin", "y"}, true, t)
	testIsLocallyValidHelper([]string{"(", "x", "^", "2", "-", "9", ")", "+", "x"}, true, t)

	testIsLocallyValidHelper([]string{"("}, false, t)
	testIsLocallyValidHelper([]string{")"}, false, t)
	testIsLocallyValidHelper([]string{"+"}, false, t)
	testIsLocallyValidHelper([]string{"sin"}, false, t)
	testIsLocallyValidHelper([]string{"sin", "+", "x"}, false, t)
	testIsLocallyValidHelper([]string{"4", "^"}, false, t)
}

func testToPostfixHelper(input []string, expected []string, unmatchedParen bool) error {
	output, err := parse.ToPostfix(input, real.Real)

	if unmatchedParen {
		if err == nil {
			return errors.New("Failed to detect unmatched parentheses.")
		}
	} else {

		if err != nil {
			return err
		}

		if len(output) != len(expected) {
			return errors.New(fmt.Sprintln("Output and expected output length do not match. Expected", expected, "Produced", output))
		}
		for i := 0; i < len(expected); i++ {
			if output[i] != expected[i] {
				return errors.New(fmt.Sprintln("Output and expected output do not match at index", i, "- Expected", expected[i], "Produced", output[i]))
			}
		}
	}
	return nil
}

func TestToPostfixHelper(t *testing.T) {
	var err error
	err = testToPostfixHelper([]string{"x", "+", "4"}, []string{"x", "4", "+"}, false)
	if err != nil {
		t.Error(err)
	}
	err = testToPostfixHelper([]string{"(", "(", "x", "+", "4", ")", "^", "5", ")"}, []string{"x", "4", "+", "5", "^"}, false)
	if err != nil {
		t.Error(err)
	}
	err = testToPostfixHelper([]string{"(", "(", "x", "+", "4", "^", "5", ")"}, []string{}, true)
	if err != nil {
		t.Error(err)
	}
	err = testToPostfixHelper([]string{"(", "x", "+", "4"}, []string{}, true)
	if err != nil {
		t.Error(err)
	}
}
