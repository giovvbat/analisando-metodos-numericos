package methods

import (
	"fmt"
	"math"
	"time"
)

func Bisection(a float64, b float64, function func(float64) float64) (float64, int) {
	x := (a + b) / 2
	iterations := 0

	for {
		if math.Abs(function(x)) <= math.Pow(10, -8) {
			break
		}

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

func PrintBisection(realRoot float64, a float64, b float64, function func(float64) float64) {
	start := time.Now()
	foundRoot, iterations := Bisection(a, b, function)
	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Bisection method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
