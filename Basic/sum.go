package main

import (
	"fmt"
)

func main() {
	var num1 float64
	var num2 float64

	fmt.Println("Ingrese el primer número a evaluar")
	fmt.Scan(&num1)

	fmt.Println("Ingrese el segundo número a evaluar")
	fmt.Scan(&num2)

	fmt.Printf("La suma de los números es: %.2f\n", num1 + num2)
	fmt.Printf("La resta de los números es: %.2f\n", num1 - num2)
	fmt.Printf("La multiplicación de los números es: %.2f\n", num1 * num2)
	fmt.Printf("La división de los números es: %.2f\n", num1 / num2)
}
