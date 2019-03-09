package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Print(i, " ")
		i = i + 1
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		sayTheNumber(i)
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

}

func sayTheNumber(number int) {
	switch number {
	case 0:
		fmt.Println(number, "Zero")
	case 1:
		fmt.Println(number, "Um")
	case 2:
		fmt.Println(number, "Dois")
	case 3:
		fmt.Println(number, "Três")
	case 4:
		fmt.Println(number, "Quatro")
	case 5:
		fmt.Println(number, "Cinco")
	case 6:
		fmt.Println(number, "Seis")
	case 7:
		fmt.Println(number, "Sete")
	case 8:
		fmt.Println(number, "Oito")
	case 9:
		fmt.Println(number, "Nove")
	case 10:
		fmt.Println(number, "Dez")
	default:
		fmt.Println("Sorry!")
	}
}
