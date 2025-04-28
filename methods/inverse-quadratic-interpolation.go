package methods

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func InverseQuadraticInterpolation(a, b, prec float64, maxIterations int, function func(float64) float64) (float64, int, error) {
	iterations := 0
	c := (a + b) / 2

	for iterations < maxIterations {
		if math.Abs(function(a)) < prec {
			return a, iterations, nil
		}

		if math.Abs(function(b)) < prec {
			return b, iterations, nil
		}

		if math.Abs(function(c)) < prec {
			return c, iterations, nil
		}

		denom := (function(a) - function(b)) * (function(a) - function(c)) * (function(b) - function(c))

		if denom == 0 {
			return 0, 0, errors.New("denominator equals zero during interpolation")
		}

		x := (function(b)*function(c))/((function(a)-function(b))*(function(a)-function(c)))*a + (function(a)*function(c))/((function(b)-function(a))*(function(b)-function(c)))*b + (function(a)*function(b))/((function(c)-function(a))*(function(c)-function(b)))*c
		a, b, c = b, c, x

		if math.Abs(function(x)) < prec {
			return x, iterations, nil
		}

		iterations++
	}

	return 0, 0, errors.New("number of max iterations exceeded")
}

func PrintInverseQuadraticInterpolation(realRoot, a, b, prec float64, maxInterations int, function func(float64) float64) {
	start := time.Now()
	foundRoot, iterations, err := InverseQuadraticInterpolation(a, b, prec, maxInterations, function)

	if err != nil {
		fmt.Printf("error while calculating root by Inverse Quadratic Interporlation method in range %.2f to %.2f\n: %s", a, b, err.Error())
	}

	end := time.Now()
	duration := end.Sub(start)
	absoluteError := math.Abs(realRoot - foundRoot)
	relativeError := math.Abs(absoluteError/realRoot) * 100

	fmt.Printf("root %.2f found by Inverse Quadratic Interporlation method in range %.2f to %.2f with %d iterations and %v needed. absolute error: %.2f, relative error: %.2f%%\n", foundRoot, a, b, iterations, duration, absoluteError, relativeError)
}
