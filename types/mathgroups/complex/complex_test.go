package complex_test

import (
	"math"
	"testing"

	"github.com/yasteen/go-parse/types/mathgroups/complex"
)

var MIN_THRESHOLD = math.Pow10(-10)

func equalEnough(a float64, b float64) bool {
	return math.Abs(a-b) < MIN_THRESHOLD
}

func testMapValuesHelper(expression string, input complex.ComplexNumber, expected complex.ComplexNumber, t *testing.T) {
	c, err := complex.MapValues(expression, *complex.NewComplexInterval(input, 1, input), "x")

	if err != nil {
		t.Error(err)
	}

	if !equalEnough(c[0].Re, expected.Re) || !equalEnough(c[0].Im, expected.Im) {
		t.Error("Failed addition on expression", expression, "- Expected:", expected, "Got:", c[0])
	}
}

func TestMapValues(t *testing.T) {
	testMapValuesHelper("x + 2_3", complex.ComplexNumber{5, 4}, complex.ComplexNumber{7, 7}, t)
	testMapValuesHelper("x - 2_3", complex.ComplexNumber{5, 4}, complex.ComplexNumber{3, 1}, t)
	testMapValuesHelper("x * 2_3", complex.ComplexNumber{5, 4}, complex.ComplexNumber{-2, 23}, t)
	testMapValuesHelper("x / 2_3", complex.ComplexNumber{5, 4}, complex.ComplexNumber{22. / 13., -7. / 13.}, t)

	testMapValuesHelper("3 * 2_3", complex.ComplexNumber{0, 0}, complex.ComplexNumber{6, 9}, t)
	testMapValuesHelper("3i * 2_3", complex.ComplexNumber{0, 0}, complex.ComplexNumber{-9, 6}, t)

	testMapValuesHelper("(3i + 2_3) * x", complex.ComplexNumber{3, 2}, complex.ComplexNumber{-6, 22}, t)
	testMapValuesHelper("exp(i * x)", complex.ComplexNumber{math.Pi, 0}, complex.ComplexNumber{-1, 0}, t)
}
