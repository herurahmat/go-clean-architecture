package author

import (
	"github.com/herurahmat/go-clean-architecture/internal/helper"
)

type AuthorModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (am *AuthorModel) GetAuthorId() string {
	return am.Id
}

func (am *AuthorModel) GetAuthorName() string {
	return am.Name
}

func CreateNewAuthor(name string) *AuthorModel {
	return &AuthorModel{
		Id:   helper.CreateNewUUID(),
		Name: name,
	}
}

func (am *AuthorModel) GetTableName() string {
	return "authors"
}
