package book

import (
	"context"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
)

type BookRepository interface {
	Get(ctx context.Context) (result []book.BookModel, err error)
	FindById(ctx context.Context, id string) (result book.BookModel, err error)
	FindByName(ctx context.Context, name string) (result book.BookModel, err error)
	Create(ctx context.Context, author author.AuthorModel, book book.BookModel) (result book.BookModel, err error)
	Update(ctx context.Context, author author.AuthorModel, book book.BookModel) (result book.BookModel, err error)
	Delete(ctx context.Context, id string) (status bool, err error)
}
