package main

import (
	"fmt"
	"os"
)

func main() {
	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(xs))

	x, y := f()
	fmt.Println(x, y)

	fmt.Println(add(1, 2, 3))

	arr := []int{1, 2, 3}
	fmt.Println(add(arr...))

	closure := func(x, y int) int {
		return x + y
	}

	fmt.Println(closure(1, 1))

	i := 0
	increment := func() int {
		i++
		return i
	}

	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEventGenerator()
	fmt.Println(nextEven(), nextEven(), nextEven())

	fmt.Println(factorial(10))

	defer second()
	first()

	openFile("teste.txt")

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")
}

func average(xs []float64) float64 {
	total := 0.0

	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0

	for _, v := range args {
		total += v
	}
	return total
}

func makeEventGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func openFile(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
}
