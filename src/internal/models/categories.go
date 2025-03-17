package models

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type CategoriesModelInterface interface {
	GetCategories(ctx context.Context, pool *pgxpool.Pool) ([]Category, error)
	AddCategory(ctx context.Context, pool *pgxpool.Pool, category *Category) (*Category, error)
	CategoryTypeByID(ctx context.Context, pool *pgxpool.Pool, categoryID uuid.UUID) (int32, error)
}

type CategoriesModelImpl struct{}

func (c *CategoriesModelImpl) GetCategories(ctx context.Context, pool *pgxpool.Pool) ([]Category, error) {
	var categories []Category

	sql, args, err := sq.
		Select("id, main_category, name, description").
		From("categories").
		ToSql()
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
		Columns("id, main_category, name, description", "progress_type").
		Values(category.ID, category.MainCategory, category.Name, category.Description, category.ProgressType).
		PlaceholderFormat(sq.Dollar)

	// Генерация SQL-запроса и списка аргументов
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		logrus.Error("Failed to generate SQL query:", err)
		return nil, err
	}

	// Выполняем запрос
	_, err = pool.Exec(ctx, sqlQuery, args...)
	if err != nil {
		logrus.Error("Failed to execute query:", err)
		return nil, err
	}

	return category, err
}

func (c *CategoriesModelImpl) CategoryTypeByID(ctx context.Context, pool *pgxpool.Pool, categoryID uuid.UUID) (int32, error) {
	var progressType int32

	query, args, err := sq.Select("progress_type").
		From("categories").
		Where(sq.Eq{"id": categoryID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		logrus.Error("Ошибка построения SQL-запроса CategoryByID:", err)
		return progressType, err
	}

	// Выполнение запроса
	err = pool.QueryRow(ctx, query, args...).Scan(&progressType)
	if err != nil {
		logrus.Error("Ошибка выполнения SQL-запроса CategoryByID:", err)
		return progressType, err
	}
	return progressType, nil
}
