package main

import (
	"analisando-metodos-numericos/functions"
	"analisando-metodos-numericos/methods"
	"fmt"
)

func main() {
	// exemplo de aplicação de um dos métodos:
	fmt.Printf("%.5f\n", methods.Secant(-5, -1.5, functions.Third))
}
