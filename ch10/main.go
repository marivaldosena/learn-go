package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println("Type Enter to exit...", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	fmt.Scanln(&input)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Type Enter to exit...", msg1)
			case msg2 := <-c2:
				fmt.Println("Type Enter to exit...", msg2)
			}

			/* select é mais usado para timeouts */
			/*
				select {
				case msg1 := <- c1:
					fmt.Println("Message 1", msg1)
				case msg2 := <- c2:
					fmt.Println("Message 2", msg2)
				case <- time.After(time.Second):
					fmt.Println("timeout")
				default:
					fmt.Println("Nothing ready")
				}
			*/
		}
	}()

	fmt.Scanln(&input)
}
