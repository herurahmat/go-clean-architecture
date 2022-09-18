package book

import (
	"context"
	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
)

type BookService interface {
	GetBook(ctx context.Context) (result []book.BookModel, err error)
	FindBookById(ctx context.Context, id string) (result book.BookModel, err error)
	CreateBook(ctx context.Context, book book.BookModel) (result book.BookModel, err error)
	UpdateBook(ctx context.Context, book book.BookModel) (result book.BookModel, err error)
	DeleteBook(ctx context.Context, id string) (status bool, err error)
}
