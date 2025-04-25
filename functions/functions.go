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

func ThirdDerivative(x float64) float64 {
	return 2*x + 1
}
