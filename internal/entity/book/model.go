package book

import (
	"github.com/herurahmat/go-clean-architecture/internal/helper"
)

type BookModel struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	AuthorId   string `json:"author_id"`
	AuthorName string `json:"author_name"`
}

func CreateNewBook(title string) *BookModel {
	return &BookModel{
		Id:    helper.CreateNewUUID(),
		Title: title,
	}
}

func CreateNewBookWithAuthor(title string, authorId string, authorName string) *BookModel {
	return &BookModel{
		Id:         helper.CreateNewUUID(),
		Title:      title,
		AuthorId:   authorId,
		AuthorName: authorName,
	}
}

func (bm *BookModel) GetBookId() string {
	return bm.Id
}

func (bm *BookModel) GetBookTitle() string {
	return bm.Title
}

func (bm *BookModel) GetBookAuthorId() string {
	return bm.AuthorId
}

func (bm *BookModel) GetBookAuthorName() string {
	return bm.AuthorName
}
