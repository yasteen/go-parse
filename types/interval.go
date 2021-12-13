package types

type Interval struct {
	Start interface{}
	End   interface{}
	Step  interface{}
	// Return Next value in the interval. Terminates if nil.
	Next func(cur interface{}) interface{}
}

func (i *Interval) NextValue(current interface{}) interface{} {
	return i.Next(current)
}
