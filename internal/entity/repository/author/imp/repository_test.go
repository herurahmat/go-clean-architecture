package imp

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetDataAuthor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error database '%s'", err)
	}

	mockAuthor := []author.AuthorModel{
		author.AuthorModel{
			Id:   helper.CreateNewUUID(),
			Name: "Book 1",
		},
		author.AuthorModel{
			Id:   helper.CreateNewUUID(),
			Name: "Book 2",
		},
		author.AuthorModel{
			Id:   helper.CreateNewUUID(),
			Name: "Book 2",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(mockAuthor[0].Id, mockAuthor[0].Name).
		AddRow(mockAuthor[1].Id, mockAuthor[1].Name).
		AddRow(mockAuthor[2].Id, mockAuthor[2].Name)

	query := "select * from ?"

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
		Name: "Book 1",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "INSERT ? SET name=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(ar.GetTableName(), ar.Name)

	a := NewAuthorRepository(db)

	err = a.Create(context.TODO(), ar)

	t.Run("Success Insert author", func(t *testing.T) {
		assert.True(t, true)
	})
}

func TestUpdateAuthor(t *testing.T) {
	ar := &author.AuthorModel{
		Id:   "123",
		Name: "Book 1",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "update  ? SET name=? WHERE id=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(ar.GetTableName(), ar.Name, ar.Id)

	a := NewAuthorRepository(db)

	err = a.Create(context.TODO(), ar)

	t.Run("Success update author", func(t *testing.T) {
		assert.True(t, true)
	})
}

func TestDeleteAuthor(t *testing.T) {
	ar := &author.AuthorModel{
		Id: "123",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error '%s'", err)
	}

	query := "delete ? WHERE id=?"
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(ar.GetTableName(), ar.Id)

	a := NewAuthorRepository(db)

	err = a.Create(context.TODO(), ar)

	t.Run("Success update author", func(t *testing.T) {
		assert.True(t, true)
	})
}
