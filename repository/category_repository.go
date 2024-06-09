package repository

import (
	"context"
	"database/sql"
	"resfull_api/model/domain"
)

type CategoryRepository interface{
	Save(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	Update(ctx context.Context,tx *sql.Tx, Category domain.Category) domain.Category
	Delete(ctx context.Context,tx *sql.Tx, Category domain.Category) 
	FindById(ctx context.Context,tx *sql.Tx, categoryId int)( domain.Category, error)
	FindAll(ctx context.Context,tx *sql.Tx,) []domain.Category
}