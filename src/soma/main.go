package main

import "fmt"

// Soma dois numeros inteiros
func Soma(a int, b int) int {
	return a + b
}

func main() {
	fmt.Printf("Resultado: %d\n", Soma(5, 5))
}
