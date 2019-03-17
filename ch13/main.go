package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
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

func openFileWithIOUtil(arquivo string) {
	bs, err := ioutil.ReadFile(arquivo)

	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func main() {
	caminho, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	arquivo := path.Join(caminho, "arquivo.txt")
	openFile(arquivo)
	openFileWithIOUtil(arquivo)
}
