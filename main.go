// Package main: A simple interface for evaluating expressions.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/yasteen/go-parse/mathgroups/real"
	"github.com/yasteen/go-parse/run"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What expression do you want to evaluate? (In terms of x)")
	var expression string
	if scanner.Scan() {
		expression = scanner.Text()
	}

	fmt.Println("What value do you want x to be?")
	var val float64
	var err error
	if scanner.Scan() {
		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}
	fmt.Println("Read value", val)

	if err != nil {
		fmt.Println("Bad values :(")
	}

	runnableReal := run.GetRunnableMathGroup(real.Real)
	fmt.Println(runnableReal.MapValues(expression, *real.NewInterval(val, 2, val+1), "x"))
}
