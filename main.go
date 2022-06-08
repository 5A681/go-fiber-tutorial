package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TableName overrides the table name used by User to `profiles`
func (Book) TableName() string {
	return "book"
}

type Book struct {
	Id       int
	Bookname string
	Author   string
}

func main() {
	dsn := "host=localhost user=postgres password=60041308116 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Table("test").AutoMigrate(&Book{})

	var book Book

	db.First(&book, 1)
	fmt.Println(book)
}
