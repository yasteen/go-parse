package types

// Interval represents a range of values within a number/value system.
type Interval[T any] struct {
	Start T
	End   T
	Step  T
	// Return Next value in the interval. Done is true if outside of interval
	Next func(cur T) (next T, done bool)
}
