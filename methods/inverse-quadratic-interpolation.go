package methods

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func InverseQuadraticInterpolation(a, b float64, maxIterations int, function func(float64) float64) (float64, int, error) {
	eps := math.Pow(10, -8)
	iterations := 0
	c := (a + b) / 2

	for {
		if iterations >= maxIterations {
			break
		}

		if math.Abs(function(a)) < eps {
			return a, iterations, nil
		}

		if math.Abs(function(b)) < eps {
			return b, iterations, nil
		}

		if math.Abs(function(c)) < eps {
			return c, iterations, nil
		}

		denom := (function(a) - function(b)) * (function(a) - function(c)) * (function(b) - function(c))

		if denom == 0 {
			return 0, 0, errors.New("denominator equals zero during interpolation")
		}

		x := (function(b)*function(c))/((function(a)-function(b))*(function(a)-function(c)))*a + (function(a)*function(c))/((function(b)-function(a))*(function(b)-function(c)))*b + (function(a)*function(b))/((function(c)-function(a))*(function(c)-function(b)))*c
		a, b, c = b, c, x

		if math.Abs(function(x)) < eps {
			return x, iterations, nil
		}

		iterations++
	}

	return 0, 0, errors.New("number of max iterations exceeded")
}

func PrintInverseQuadraticInterpolation(realRoot, a, b float64, maxInterations int, function func(float64) float64) {
	start := time.Now()
	foundRoot, iterations, err := InverseQuadraticInterpolation(a, b, maxInterations, function)

	if err != nil {
		fmt.Printf("error while calculating root by Inverse Quadratic Interporlation method in range %.2f to %.2f\n: %s", a, b, err.Error())
	}

	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Inverse Quadratic Interporlation method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
