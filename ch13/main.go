package main

import (
	"fmt"
	"os"
)

func openFile(arquivo string) {
	file, err := os.Open(arquivo)

	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func main() {
	openFile("arquivo.txt")
}
