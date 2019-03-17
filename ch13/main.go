package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
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

func createFile(arquivo string) {
	file, err := os.Create(arquivo)

	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("test from file.")
}

func readCurrentDirectory() {
	dir, err := os.Open(".")

	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func readCurrentDirectoryRecursively() {
	fmt.Println("Current Directory and Subfolders")
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

func main() {
	caminho, err := os.Getwd()

	if err != nil {
		log.Fatalln(err)
	}

	arquivo := path.Join(caminho, "arquivo.txt")
	createFile(arquivo)
	openFile(arquivo)
	openFileWithIOUtil(arquivo)
	readCurrentDirectory()
}
