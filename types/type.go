package types

// ParseSystem is an interface to regulate the input interval and the math group used
// for parsing and evaluating expressions.
// TODO: Implement more rigorous interface for calling functions when generics come to Go
type ParseSystem interface {
	NewInterval()
	MapValues(expression string, interval Interval, varName string)
}
