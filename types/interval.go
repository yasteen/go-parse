package types

type Interval struct {
	Start interface{}
	End   interface{}
	Step  interface{}
	// Return next value in the interval. Terminates if nil.
	next func(cur interface{}, step interface{}, end interface{}) interface{}
}

func (i *Interval) Next(current interface{}) interface{} {
	return i.next(current, i.Step, i.End)
}
