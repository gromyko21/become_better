package models

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type CategoriesModelInterface interface {
	GetCategories(ctx context.Context, pool *pgxpool.Pool) ([]Categories, error)
}

type CategoriesModelImpl struct {}

func (c *CategoriesModelImpl) GetCategories(ctx context.Context, pool *pgxpool.Pool) ([]Categories, error) {
	var categories []Categories

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
