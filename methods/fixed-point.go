package methods

import (
	"fmt"
	"math"
	"time"
)

func FixedPoint(a, b, prec float64, function, functionFi func(float64) float64) (float64, int) {
	x := (a + b) / 2
	iterations := 0

	for math.Abs(function(x)) > math.Pow(10, -8) {
		iterations++
		x = functionFi(x)
	}

	return x, iterations
}

func PrintFixedPoint(realRoot, a, b, prec float64, function, functionFi func(float64) float64) {
	start := time.Now()
	foundRoot, iterations := FixedPoint(a, b, prec, function, functionFi)
	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Fixed Point method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
