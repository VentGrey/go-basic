package main

import (
	"fmt"
	"math"
)

func main() {
	var rad float64
	fmt.Println("Ingrese el radio del círculo a evaluar")

	fmt.Printf("El área del círclulo es: %.2f\n", math.Pi * (rad * rad))
	fmt.Printf("La circunferencia es: %.2f\n", math.Pi * (2 * rad))
}
