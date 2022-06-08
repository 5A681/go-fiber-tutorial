package models

type Book struct {
	Id       int
	Bookname string
	Author   string
}

func (Book) TableName() string {
	return "book"
}
