package methods

import (
	"fmt"
	"math"
	"time"
)

func Bisection(a, b, prec float64, function func(float64) float64) (float64, int) {
	x := (a + b) / 2
	iterations := 0

	for math.Abs(function(x)) > prec {
		if function(a)*function(x) < 0 {
			b = x
		} else {
			a = x
		}

		x = (a + b) / 2
		iterations++
	}

	return x, iterations
}

func PrintBisection(realRoot, a, b, prec float64, function func(float64) float64) {
	start := time.Now()
	foundRoot, iterations := Bisection(a, b, prec, function)
	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Bisection method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
