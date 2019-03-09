package main

import "fmt"

func main() {
	var fahr float64
	fmt.Println("Fahr: ")
	fmt.Scanf("%f", &fahr)
	fmt.Println("Result: ", fahrToCelsius(fahr))
}

func fahrToCelsius(fahr float64) float64 {
	result := (fahr - 32) * 5 / 9
	return result
}
