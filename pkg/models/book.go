package models

import (
	"github.com/jinzhu/gorm"
	"github.com/notcoderguy/go-crud-server/pkg/config"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name 	  string `json:"name"`
	Author    string `json:"author"`
	Publication string `json:"publication"`
	Price     float32 `json:"price"`
 }

func init() {
	config.Connect()
	DB = config.GetDB()
	DB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	DB.NewRecord(b)
	DB.Create(&b)
	return b
}

func GetBooks() []Book {
	var books []Book
	DB.Find(&books)
	return books
}

func GetBook(id int64) (Book, error) {
	var book Book
	err := DB.Where("ID = ?", id).First(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func DeleteBook(id  int64) Book {
	var book Book
	DB.Where("ID = ?", id).Delete(&book)
	return book
}