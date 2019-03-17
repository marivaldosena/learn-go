package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"net/http"
)

func server() {
	ln, err := net.Listen("tcp", ":8089")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received ", msg)
	}

	c.Close()
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:8089")

	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello World"
	fmt.Println("Sending ", msg)

	err = gob.NewEncoder(c).Encode(msg)

	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		res,
		`<doctype html>
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
				<h1>Hello World!</h1>
			</body>
		</html>`,
	)
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)

	http.HandleFunc("/", hello)
	fmt.Println("Running server on http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
