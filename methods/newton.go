package methods

import (
	"analisando-metodos-numericos/functions"
	"math"
)

func Newton(a float64, b float64, function func(float64) float64) float64 {
	x := (a + b) / 2

	for {
		if math.Abs(function(x)) <= math.Pow(10, -8) {
			break
		}

		x = x - function(x)/functions.Derivative(x, function)
	}

	return x
}
