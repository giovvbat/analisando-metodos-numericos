package functions

import "math"

func First(x float64) float64 {
	return math.Pow(x, 5) - (10.0 / 9 * math.Pow(x, 3)) + (5.0 / 21 * x)
}

func Second(x float64) float64 {
	return math.Pow(x, 2) + math.Log(x)
}

func Third(x float64) float64 {
	return math.Pow(x, 2) + x - 6
}

func ThirdFi(x float64) float64 {
	return 6 / (x + 1)
}

func Derivative(x float64, function func(float64) float64) float64 {
	eps := math.Pow(10, -8)
	return (function(x+eps) - function(x)) / eps
}
