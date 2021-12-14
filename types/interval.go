package types

// Interval represents a range of values within a number/value system.
type Interval struct {
	Start interface{}
	End   interface{}
	Step  interface{}
	// Return Next value in the interval. Terminates if nil.
	Next func(cur interface{}) interface{}
}

// NextValue gives the next value when iterating through an interval
func (i *Interval) NextValue(current interface{}) interface{} {
	return i.Next(current)
}
