package imp

import (
	"context"
	"database/sql"
	"log"

	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
)

type repository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context) (result []author.AuthorModel, err error) {
	rows, err := r.db.QueryContext(ctx, "select * from authors")
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

	result = make([]author.AuthorModel, 0)
	for rows.Next() {
		t := author.AuthorModel{}
		rows.Scan(
			&t.Id,
			&t.Name,
		)
		result = append(result, t)
	}

	return result, nil
}

func (r *repository) FindById(ctx context.Context, id string) (result author.AuthorModel, err error) {
	query := "select * from authors WHERE id = ?"

	statement, err := r.db.PrepareContext(ctx, string(query))

	if err != nil {
		return author.AuthorModel{}, err
	}

	row := statement.QueryRowContext(ctx, id)

	result = author.AuthorModel{}

	err = row.Scan(
		&result.Id,
		&result.Name,
	)

	return
}

func (r *repository) FindByName(ctx context.Context, name string) (result author.AuthorModel, err error) {
	query := "select * from authors WHERE name LIKE '%?%'"

	statement, err := r.db.PrepareContext(ctx, string(query))

	if err != nil {
		return author.AuthorModel{}, err
	}

	row := statement.QueryRowContext(ctx, name)

	result = author.AuthorModel{}

	err = row.Scan(
		&result.Id,
		&result.Name,
	)

	return
}

func (r *repository) Create(ctx context.Context, dataEntity *author.AuthorModel) (*author.AuthorModel, error) {
	query := `INSERT authors SET name=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return dataEntity, err
	}

	_, err = stmt.ExecContext(ctx, dataEntity.Name)
	if err != nil {
		return dataEntity, err
	}
	return dataEntity, nil
}

func (r *repository) Update(ctx context.Context, dataEntity *author.AuthorModel) (*author.AuthorModel, error) {
	query := `update  authors SET name=? WHERE id=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return dataEntity, err
	}

	_, err = stmt.ExecContext(ctx, dataEntity.Name, dataEntity.Id)
	if err != nil {
		return dataEntity, err
	}
	return dataEntity, nil
}

func (r *repository) Delete(ctx context.Context, id string) (err error) {
	query := "delete authors WHERE id=?"
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
