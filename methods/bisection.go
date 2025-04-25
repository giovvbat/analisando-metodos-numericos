package methods

import "math"

func Bisection(a float64, b float64, function func(float64) float64) float64 {
	x := (a + b) / 2

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
	}

	return x
}
