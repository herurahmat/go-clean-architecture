package imp

import (
	"context"
	"database/sql"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"log"
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
	ett := author.AuthorModel{}
	rows, err := r.db.QueryContext(ctx, "select * from ?", ett.GetTableName())
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
	entity := author.AuthorModel{}
	query := "select * from ? WHERE id = ?"

	statement, err := r.db.PrepareContext(ctx, string(query))

	if err != nil {
		return author.AuthorModel{}, err
	}

	row := statement.QueryRowContext(ctx, entity.GetTableName(), id)

	result = author.AuthorModel{}

	err = row.Scan(
		&result.Id,
		&result.Name,
	)

	return
}

func (r *repository) FindByName(ctx context.Context, name string) (result author.AuthorModel, err error) {
	entity := author.AuthorModel{}
	query := "select * from ? WHERE name LIKE '%?%'"

	statement, err := r.db.PrepareContext(ctx, string(query))

	if err != nil {
		return author.AuthorModel{}, err
	}

	row := statement.QueryRowContext(ctx, entity.GetTableName(), name)

	result = author.AuthorModel{}

	err = row.Scan(
		&result.Id,
		&result.Name,
	)

	return
}

func (r *repository) Create(ctx context.Context, author *author.AuthorModel) (err error) {

	query := `INSERT  ? SET name=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, author.GetTableName(), author.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(ctx context.Context, author *author.AuthorModel) (status bool, err error) {

	query := `update  ? SET name=? WHERE id=?`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, author.GetTableName(), author.Name, author.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) Delete(ctx context.Context, id string) (status bool, err error) {
	entity := author.AuthorModel{}
	query := "delete ? WHERE id=?"
	stmt, err := r.db.PrepareContext(ctx, string(query))
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, entity.GetTableName(), id)
	if err != nil {
		return false, err
	}
	return true, nil
}
