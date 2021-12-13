# equation-parser
Parses and calculates expressions and equations.

Supports parsing equations in different number systems (currently implemented for real and complex numbers).

## Example Usage
Run the following:
```go
real.MapValues("exp(x) - 3^2", real.NewRealInterval(0, 0.1, 5), "x")
```
to get the corresponding `y` values for `y = exp(x) - 3^2` in the interval `[0, 5]`, where `x` increments by `0.1` each time.
