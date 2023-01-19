package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookName   string `json:"bookName"`
	BookAuthor string `json:"bookAuthor"`
}
