package methods

import "math"

func FalsePosition(a float64, b float64, function func(float64) float64) float64 {
	x := (a*function(b) - b*function(a)) / (function(b) - function(a))

	for {
		if math.Abs(function(x)) <= math.Pow(10, -8) {
			break
		}

		if function(a)*function(x) < 0 {
			b = x
		} else {
			a = x
		}

		x = (a*function(b) - b*function(a)) / (function(b) - function(a))
	}

	return x
}
