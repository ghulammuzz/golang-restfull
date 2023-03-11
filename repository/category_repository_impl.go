package repository

import (
	"context"
	"database/sql"
	"errors"
	"ghulammuzz/golang-restfull/helper"
	"ghulammuzz/golang-restfull/model/domain"
)

type CategoryRepositoryImp struct {
	
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImp{}
}

func (repository *CategoryRepositoryImp) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into category(name) value (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)
	
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category	
}
func (repository *CategoryRepositoryImp) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id =?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category

}
func (repository *CategoryRepositoryImp) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

}
func (repository *CategoryRepositoryImp) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Gaada")
	}

}
func (repository *CategoryRepositoryImp) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err :=tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	
	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}