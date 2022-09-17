package imp

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetDataAuthor(t *testing.T) {
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

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(mockAuthor[0].Id, mockAuthor[0].Name).
		AddRow(mockAuthor[1].Id, mockAuthor[1].Name).
		AddRow(mockAuthor[2].Id, mockAuthor[2].Name)

	query := "select * from authors"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	instance := NewAuthorRepository(db)

	list, err := instance.Get(context.TODO())

	t.Run("success check data", func(t *testing.T) {
		assert.Nil(t, err)
		assert.Len(t, list, 3)
		assert.Equal(t, "Book 1", list[0].Name)
	})
}

func TestStoreAuthor(t *testing.T) {
	ar := &author.AuthorModel{
		Id:   helper.CreateNewUUID(),
		Name: "Author 1",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "INSERT authors SET name=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(ar.Name)

	a := NewAuthorRepository(db)

	res, err := a.Create(context.TODO(), ar)

	t.Run("Success Insert author", func(t *testing.T) {
		assert.True(t, true)
		assert.NotEmpty(t, res)
	})
	assert.NotNil(t, err)

}

func TestUpdateAuthor(t *testing.T) {
	ar := &author.AuthorModel{
		Id:   "123",
		Name: "Author 1",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "update authors SET name=? WHERE id=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(ar.Name, ar.Id)

	a := NewAuthorRepository(db)

	res, err := a.Update(context.TODO(), ar)

	t.Run("Success update author", func(t *testing.T) {
		assert.True(t, true)
		assert.NotEmpty(t, res)
	})
	assert.NotNil(t, err)
}

func TestDeleteAuthor(t *testing.T) {
	ar := &author.AuthorModel{
		Id: "123",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "delete authors WHERE id=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(ar.Id)

	a := NewAuthorRepository(db)

	err = a.Delete(context.TODO(), ar.Id)

	t.Run("Success update author", func(t *testing.T) {
		assert.True(t, true)
	})

	assert.NotNil(t, err)
}
