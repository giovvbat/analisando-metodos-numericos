package methods

import "math"

func FixedPoint(a float64, b float64, function func(float64) float64, fi func(float64) float64) float64 {
	x := (a + b) / 2

	for {
		if math.Abs(function(x)) <= math.Pow(10, -8) {
			break
		}

		x = fi(x)
	}

	return x
}
