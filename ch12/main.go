package main

import (
	"fmt"

	s "github.com/marivaldosena/learn-go/ch12/str_utils"
)

const word = "test"

/*
	Implmenta o package str_utils que utiliza as funções do
	package str.
*/
func main() {
	// Contains
	fmt.Println(s.Contains(word, "es"))
	// Count
	fmt.Println(s.Count(word, "t"))
	// BeginsWith
	fmt.Println(s.BeginsWith(word, "te"))
	// EndsWith
	fmt.Println(s.EndsWith(word, "st"))
	// Index
	fmt.Println(s.Index(word, "e"))
	// Concat
	fmt.Println(s.Concat([]string{"a", "b"}, "-"))
	// Repeat
	fmt.Println(s.Repeat("a", 5))
	// Upper
	fmt.Println(s.Upper(word))
}
