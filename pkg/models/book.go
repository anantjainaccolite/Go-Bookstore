package models

import (
	"github.com/anantjain28/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//automigrate =keeps the schema of our database upto Date

func init() {
	//this fucntion executes prior to the main()
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
	//passing the refernce of the struct book will create a bokk table with column name as the var in struct

}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var b Book
	db := db.Where("ID=?", Id).Find(&b)
	return &b, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
