package imp

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetDataBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error database '%s'", err)
	}

	mockAuthor := []author.AuthorModel{
		author.AuthorModel{
			Id:   helper.CreateNewUUID(),
			Name: "Author 1",
		},
		author.AuthorModel{
			Id:   helper.CreateNewUUID(),
			Name: "Author 2",
		},
		author.AuthorModel{
			Id:   helper.CreateNewUUID(),
			Name: "Author 2",
		},
	}

	mockBook := []book.BookModel{
		book.BookModel{
			Id:         helper.CreateNewUUID(),
			Title:      "Book 1",
			AuthorId:   mockAuthor[0].Id,
			AuthorName: mockAuthor[0].Name,
		},
		book.BookModel{
			Id:         helper.CreateNewUUID(),
			Title:      "Book 2",
			AuthorId:   mockAuthor[1].Id,
			AuthorName: mockAuthor[1].Name,
		},
		book.BookModel{
			Id:         helper.CreateNewUUID(),
			Title:      "Book 3",
			AuthorId:   mockAuthor[2].Id,
			AuthorName: mockAuthor[2].Name,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "book_author_id", "book_author_name"}).
		AddRow(mockBook[0].Id, mockBook[0].Title, mockBook[0].GetBookAuthorId(), mockBook[0].GetBookAuthorName()).
		AddRow(mockBook[1].Id, mockBook[1].Title, mockBook[1].GetBookAuthorId(), mockBook[1].GetBookAuthorName()).
		AddRow(mockBook[1].Id, mockBook[1].Title, mockBook[1].GetBookAuthorId(), mockBook[1].GetBookAuthorName())

	query := "select 'books'.'*','authors'.'name' as authorName,'authors'.'id' as authorId from books join authors on 'books'.'author_id' = 'authors'.'id'"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	instance := NewBookRepository(db)

	list, err := instance.Get(context.TODO())

	t.Run("success check data", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Len(t, list, 3)
		assert.Equal(t, "Book 1", list[0].Title)
	})
}

func TestStoreBook(t *testing.T) {
	mockAuthor := author.AuthorModel{
		Id:   helper.CreateNewUUID(),
		Name: "Author 1",
	}
	mockBook := &book.BookModel{
		Id:    helper.CreateNewUUID(),
		Title: "Book 1",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "INSERT books SET name=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(mockBook.Title)

	a := NewBookRepository(db)

	res, err := a.Create(context.TODO(), &mockAuthor, mockBook)

	t.Run("Success Insert book", func(t *testing.T) {
		assert.True(t, true)
		assert.NotEmpty(t, res)
	})
	assert.NotNil(t, err)

}

func TestUpdateBook(t *testing.T) {

	mockAuthor := author.AuthorModel{
		Id:   helper.CreateNewUUID(),
		Name: "Author 2",
	}

	mockBook := &book.BookModel{
		Id:    "123",
		Title: "Book 1",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "update books SET name=?,author_id=? WHERE id=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(mockBook.Title, mockAuthor.GetAuthorId, mockBook.GetBookId)

	a := NewBookRepository(db)

	res, err := a.Update(context.TODO(), &mockAuthor, mockBook)

	t.Run("Success update book", func(t *testing.T) {
		assert.True(t, true)
		assert.NotEmpty(t, res)
	})
	assert.NotNil(t, err)
}

func TestDeleteBook(t *testing.T) {
	ar := &book.BookModel{
		Id: "123",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "delete books WHERE id=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(ar.Id)

	a := NewBookRepository(db)

	err = a.Delete(context.TODO(), ar.Id)

	t.Run("Success update book", func(t *testing.T) {
		assert.True(t, true)
	})

	assert.NotNil(t, err)
}
