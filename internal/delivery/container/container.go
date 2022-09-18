package container

import (
	"github.com/herurahmat/go-clean-architecture/infrastructure/mysql"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	authorRepo "github.com/herurahmat/go-clean-architecture/internal/entity/repository/author/imp"
	bookRepo "github.com/herurahmat/go-clean-architecture/internal/entity/repository/book/imp"
	useCaseAuthor "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/author"
	useCaseAuthorImpl "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/author/impl"
	useCaseBook "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/book"
	useCaseBookImpl "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/book/impl"
	"log"
)

type Container struct {
	AuthorService useCaseAuthor.AuthorService
	BookService   useCaseBook.BookService
	Config        *config.Config
}

func New() *Container {
	c, err := config.New()
	if err != nil {
		log.Panic(err)
	}

	db, err := mysql.NewDatabase(c)
	if err != nil {
		log.Panic(err)
	}
	authorRepository := authorRepo.NewAuthorRepository(db)
	authorService := useCaseAuthorImpl.NewServiceAuthor(authorRepository, db, c)

	bookRepo := bookRepo.NewBookRepository(db)
	bookService := useCaseBookImpl.NewServiceBook(bookRepo, db, c)

	return &Container{
		AuthorService: authorService,
		BookService:   bookService,
		Config:        c,
	}
}
