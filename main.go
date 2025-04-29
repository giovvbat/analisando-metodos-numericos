package main

import (
	"analisando-metodos-numericos/functions"
	"analisando-metodos-numericos/methods"
	"fmt"
	"math"
)

func main() {
	a, b, realRoot, prec, maxIterations := 1.0, 1.4, 2.0, math.Pow(10, -8), 10000

	for b <= 3 {
		if functions.Third(a)*functions.Third(b) < 0 {
			methods.PrintBisection(realRoot, a, b, prec, functions.Third)
			methods.PrintFalsePosition(realRoot, a, b, prec, functions.Third)
			methods.PrintFixedPoint(realRoot, a, b, prec, functions.Third, functions.ThirdFi)
			methods.PrintNewton(realRoot, a, b, prec, functions.Third)
			methods.PrintSecant(realRoot, a, b, prec, functions.Third)
			methods.PrintInverseQuadraticInterpolation(realRoot, a, b, prec, maxIterations, functions.Third)
		} else {
			fmt.Printf("no roots in range %.2f to %.2f\n", a, b)
		}

		a = a + 0.4
		b = b + 0.4
	}
}
