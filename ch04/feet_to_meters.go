package main

import "fmt"

func main() {
	var feet float64
	fmt.Print("Feet: ")
	fmt.Scanf("%f", &feet)
	fmt.Println("Meters: ", feetToMeters(feet))
}

func feetToMeters(feet float64) float64 {
	result := feet * 0.3048
	return result
}
