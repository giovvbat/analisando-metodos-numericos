package main

import (
	"analisando-metodos-numericos/functions"
	"analisando-metodos-numericos/methods"
	"fmt"
)

func main() {
	a, b, realRoot := 1.0, 1.4, 2.0

	for {
		if b > 3 {
			break
		}

		if functions.Third(a)*functions.Third(b) < 0 {
			methods.PrintBisection(realRoot, a, b, functions.Third)
			methods.PrintFalsePosition(realRoot, a, b, functions.Third)
			methods.PrintFixedPoint(realRoot, a, b, functions.Third, functions.ThirdFi)
			methods.PrintNewton(realRoot, a, b, functions.Third)
			methods.PrintSecant(realRoot, a, b, functions.Third)
			methods.PrintInverseQuadraticInterpolation(realRoot, a, b, 10.000, functions.Third)
		} else {
			fmt.Printf("no roots in range %.2f to %.2f\n", a, b)
		}

		a = a + 0.4
		b = b + 0.4
	}
}
