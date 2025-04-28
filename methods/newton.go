package methods

import (
	"analisando-metodos-numericos/functions"
	"fmt"
	"math"
	"time"
)

func Newton(a, b, prec float64, function func(float64) float64) (float64, int) {
	x := (a + b) / 2
	iterations := 0

	for math.Abs(function(x)) > prec {
		iterations++
		x = x - function(x)/functions.Derivative(x, function)
	}

	return x, iterations
}

func PrintNewton(realRoot, a, b, prec float64, function func(float64) float64) {
	start := time.Now()
	foundRoot, iterations := Newton(a, b, prec, function)
	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Newton method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
