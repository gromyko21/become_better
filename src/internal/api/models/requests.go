package models

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type CategoriesModelInterface interface {
	GetCategories(ctx context.Context, pool *pgxpool.Pool) ([]Category, error)
	AddCategory(ctx context.Context, pool *pgxpool.Pool, category *Category) (*Category, error)  
}

type CategoriesModelImpl struct {}

func (c *CategoriesModelImpl) GetCategories(ctx context.Context, pool *pgxpool.Pool) ([]Category, error) {
	var categories []Category

	sql, args, err := (sq.
		Select("id, main_category, name, description").
		From("categories").
		ToSql())
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	
	err = pgxscan.Select(ctx, pool, &categories, sql, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return categories, err
}

func (c *CategoriesModelImpl) AddCategory(ctx context.Context, pool *pgxpool.Pool, category *Category) (*Category, error) {

	query := sq.
		Insert("categories").
		Columns("id, main_category, name, description").
		Values(category.ID, category.MainCategory, category.Name, category.Description).
		PlaceholderFormat(sq.Dollar)

	// Генерация SQL-запроса и списка аргументов
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		logrus.Fatal("Failed to generate SQL query:", err)
	}

	// Выполняем запрос
	_, err = pool.Exec(ctx, sqlQuery, args...)
	if err != nil {
		logrus.Fatal("Failed to execute query:", err)
	}

	return category, err
}
