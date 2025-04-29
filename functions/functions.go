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

func Fourth(x float64) float64 {
	return 2*math.Pow(x, 4) + 4*math.Pow(x, 3) + 3*math.Pow(x, 2) - 10*x - 15
}

func SecondFourth(x float64) float64 {
	return math.Pow(x, 5) - 2*math.Pow(x, 4) - 9*math.Pow(x, 3) + 22*math.Pow(x, 2) + 4*x - 24
}

func Derivative(x float64, function func(float64) float64) float64 {
	eps := math.Pow(10, -8)
	return (function(x+eps) - function(x)) / eps
}
