package methods

import (
	"fmt"
	"math"
	"time"
)

func Secant(a float64, b float64, function func(float64) float64) (float64, int) {
	x := (a*function(b) - b*function(a)) / (function(b) - function(a))
	iterations := 0

	for {
		if math.Abs(function(x)) <= math.Pow(10, -8) {
			break
		}

		iterations++
		a = b
		b = x
		x = (a*function(b) - b*function(a)) / (function(b) - function(a))
	}

	return x, iterations
}

func PrintSecant(realRoot float64, a float64, b float64, function func(float64) float64) {
	start := time.Now()
	foundRoot, iterations := Secant(a, b, function)
	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Secant method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
