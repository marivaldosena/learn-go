package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	h := crc32.NewIEEE()
	h.Write(bs)

	return h.Sum32(), nil
}

func main() {
	caminho, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	arquivo := path.Join(caminho, "ch15", "test1.txt")

	h1, err := getHash(arquivo)

	if err != nil {
		return
	}

	arquivo = path.Join(caminho, "ch15", "test2.txt")

	h2, err := getHash(arquivo)

	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)

	h := sha1.New()
	h.Write([]byte("test"))

	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
