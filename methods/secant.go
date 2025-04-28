package methods

import (
	"fmt"
	"math"
	"time"
)

func Secant(a, b, prec float64, function func(float64) float64) (float64, int) {
	x := (a*function(b) - b*function(a)) / (function(b) - function(a))
	iterations := 0

	for math.Abs(function(x)) > prec {
		iterations++
		a = b
		b = x
		x = (a*function(b) - b*function(a)) / (function(b) - function(a))
	}

	return x, iterations
}

func PrintSecant(realRoot, a, b, prec float64, function func(float64) float64) {
	start := time.Now()
	foundRoot, iterations := Secant(a, b, prec, function)
	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Secant method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
