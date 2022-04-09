package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muhangga/helper"
	"github.com/muhangga/model/domain"
)

type CategoryRepositoryImpl struct {}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category (name) values (?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = id
	return category
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name) 
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "select id, name from category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category

		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}