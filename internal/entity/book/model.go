package book

import (
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
)

type BookModel struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author *author.AuthorModel
}

func CreateNewBook(title string, author author.AuthorModel) *BookModel {
	return &BookModel{
		Id:     helper.CreateNewUUID(),
		Title:  title,
		Author: &author,
	}
}

func (bm *BookModel) GetBookId() string {
	return bm.Id
}

func (bm *BookModel) GetBookTitle() string {
	return bm.Title
}
