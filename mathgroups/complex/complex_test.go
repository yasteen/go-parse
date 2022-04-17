package complex_test

import (
	"math"
	"testing"

	"github.com/yasteen/go-parse/mathgroups/complex"
	"github.com/yasteen/go-parse/run"
)

var MIN_THRESHOLD = math.Pow10(-10)

func equalEnough(a float64, b float64) bool {
	return math.Abs(a-b) < MIN_THRESHOLD
}

func testMapValuesHelper(expression string, input complex.Number, expected complex.Number, t *testing.T) {
	runnableComplex := run.GetRunnableMathGroup(complex.Complex)
	c, err := runnableComplex.MapValues(expression, *complex.NewComplexInterval(input, complex.Number{1, 0}, input), "x")

	if err != nil {
		t.Error(err)
	}

	if !equalEnough(c[0].Re, expected.Re) || !equalEnough(c[0].Im, expected.Im) {
		t.Error("Failed addition on expression", expression, "- Expected:", expected, "Got:", c[0])
	}
}

func TestMapValues(t *testing.T) {
	testMapValuesHelper("x + 2_3", complex.Number{5, 4}, complex.Number{7, 7}, t)
	testMapValuesHelper("x - 2_3", complex.Number{5, 4}, complex.Number{3, 1}, t)
	testMapValuesHelper("x * 2_3", complex.Number{5, 4}, complex.Number{-2, 23}, t)
	testMapValuesHelper("x / 2_3", complex.Number{5, 4}, complex.Number{22. / 13., -7. / 13.}, t)

	testMapValuesHelper("3 * 2_3", complex.Number{0, 0}, complex.Number{6, 9}, t)
	testMapValuesHelper("3i * 2_3", complex.Number{0, 0}, complex.Number{-9, 6}, t)

	testMapValuesHelper("(3i + 2_3) * x", complex.Number{3, 2}, complex.Number{-6, 22}, t)
	testMapValuesHelper("exp(i * x)", complex.Number{math.Pi, 0}, complex.Number{-1, 0}, t)
}
