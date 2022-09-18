package imp

import (
	"context"
	"database/sql"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
	"log"

	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
)

const (
	queryResult = `select 'books.*','authors.name' as authorName,'authors.id' as authorId from books join authors on 'books.author_id' = 'authors.id'`
)

type repository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context) (result []book.BookModel, err error) {
	rows, err := r.db.QueryContext(ctx, queryResult)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println(errRow)
		}
	}()

	result = make([]book.BookModel, 0)
	for rows.Next() {
		t := book.BookModel{}
		rows.Scan(
			&t.Id,
			&t.Title,
			&t.AuthorId,
			&t.AuthorName,
		)
		result = append(result, t)
	}

	return result, nil
}

func (r *repository) FindById(ctx context.Context, id string) (result book.BookModel, err error) {
	query := queryResult + ` WHERE books.id=?`

	statement, err := r.db.PrepareContext(ctx, string(query))

	if err != nil {
		return book.BookModel{}, err
	}

	row := statement.QueryRowContext(ctx, id)

	result = book.BookModel{}

	err = row.Scan(
		&result.Id,
		&result.Title,
		&result.AuthorId,
		&result.AuthorName,
	)

	return
}

func (r *repository) Create(ctx context.Context, dataEntity book.BookModel) (book.BookModel, error) {
	uuid := helper.CreateNewUUID()
	query := `INSERT books SET id=?,title=?,author_id=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return dataEntity, err
	}

	_, err = stmt.ExecContext(ctx, uuid, dataEntity.Title, dataEntity.AuthorId)
	if err != nil {
		return dataEntity, err
	}

	dataEntity.Id = uuid
	return dataEntity, nil
}

func (r *repository) Update(ctx context.Context, dataEntity book.BookModel) (book.BookModel, error) {
	query := `update  books SET title=?,author_id=? WHERE id=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return dataEntity, err
	}

	_, err = stmt.ExecContext(ctx, dataEntity.Title, dataEntity.GetBookAuthorId(), dataEntity.Id)
	if err != nil {
		return dataEntity, err
	}
	return dataEntity, nil
}

func (r *repository) Delete(ctx context.Context, id string) (status bool, err error) {
	query := "DELETE FROM books WHERE id = ?"
	stmt, err := r.db.PrepareContext(ctx, string(query))
	if err != nil {
		return false, err
	}

	result, err := stmt.ExecContext(ctx, id)

	if err != nil {
		return false, err
	}

	rowsAfected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAfected != 1 {
		return false, nil
	}

	return true, nil
}
