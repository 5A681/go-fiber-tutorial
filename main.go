package main

import (
	"fmt"

	"github.com/go-fiber-tutorial/database"
)

func main() {
	err := database.Init()
	if err != nil {
		panic(err)
	}
	book := database.GetBooks()
	fmt.Println(book)
}
