package types

// TODO: Implement more rigourous interface for calling functions when generics come to Go

type ParseSystem interface {
	NewInterval()
	MapValues(expression string, interval Interval, varName string)
}
