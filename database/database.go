package database

import (
	"fmt"

	"github.com/go-fiber-tutorial/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {

	var err error
	dsn := "host=localhost user=postgres password=60041308116 dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("-- Database Conected -- ")

	return nil
}

func GetBooks() []models.Book {
	var book []models.Book
	db.Table("book").AutoMigrate(&models.Book{})

	db.Where([]int64{1, 2}).Find(&book)
	return book
}
func GetBook() models.Book {
	var book models.Book
	return book
}
func UpdateBook() {

}
