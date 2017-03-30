package main

import (
	"fmt"
	"os"

	"github.com/brandon0172/go-todo-rest/go-todo-rest/web"
	graceful "gopkg.in/tylerb/graceful.v1"
)

func main() {
	fmt.Println("Go todo REST api")
	s := web.NewServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "1337"
	}
	fmt.Println("Starting server on: 127.0.0.1:" + port)
	graceful.Run(":"+port, 0, s)
}
