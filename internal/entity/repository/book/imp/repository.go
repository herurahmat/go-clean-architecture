package imp

import (
	"context"
	"database/sql"
	"log"

	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
)

const (
	queryResult = `select 'books'.'*','authors'.'name' as authorName,'authors'.'id' as authorId from books join authors on 'books'.'author_id' = 'authors'.'id'`
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
	query := queryResult + ` WHERE id=?`

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

func (r *repository) FindByName(ctx context.Context, name string) (result book.BookModel, err error) {
	query := queryResult + ` WHERE name LIKE '%?%'`

	statement, err := r.db.PrepareContext(ctx, string(query))

	if err != nil {
		return book.BookModel{}, err
	}

	row := statement.QueryRowContext(ctx, name)

	result = book.BookModel{}

	err = row.Scan(
		&result.Id,
		&result.Title,
		&result.AuthorId,
		&result.AuthorName,
	)

	return
}

func (r *repository) Create(ctx context.Context, dataAuthor *author.AuthorModel, dataEntity *book.BookModel) (*book.BookModel, error) {
	query := `INSERT books SET name=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return dataEntity, err
	}

	_, err = stmt.ExecContext(ctx, dataEntity.Title)
	if err != nil {
		return dataEntity, err
	}
	return dataEntity, nil
}

func (r *repository) Update(ctx context.Context, dataAuthor *author.AuthorModel, dataEntity *book.BookModel) (*book.BookModel, error) {
	query := `update  books SET name=?,author_id=? WHERE id=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return dataEntity, err
	}

	_, err = stmt.ExecContext(ctx, dataEntity.Title, dataAuthor.GetAuthorId(), dataEntity.Id)
	if err != nil {
		return dataEntity, err
	}
	return dataEntity, nil
}

func (r *repository) Delete(ctx context.Context, id string) (err error) {
	query := "delete books WHERE id=?"
	stmt, err := r.db.PrepareContext(ctx, string(query))
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
