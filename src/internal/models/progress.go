package models

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type ProgressModelInterface interface {
	AddProgress(ctx context.Context, pool *pgxpool.Pool, progress *Progress) (error)  
}

type ProgressModelImpl struct {}

func (c *CategoriesModelImpl) AddProgress(ctx context.Context, pool *pgxpool.Pool, progress *Progress) (error)  {

	query := sq.
		Insert("progress").
		Columns("id, user_id, category_id, progress_type", "result_int", "result_description", "date").
		Values(progress.ID, progress.UserID, progress.CategoryID, progress.ProgressType, 
			progress.Result, progress.Description, progress.Date).
		PlaceholderFormat(sq.Dollar)

	// Генерация SQL-запроса и списка аргументов
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		logrus.Error("AddProgress failed to generate SQL query:", err)
	}

	// Выполняем запрос
	_, err = pool.Exec(ctx, sqlQuery, args...)
	if err != nil {
		logrus.Error("AddProgress failed to execute query:", err)
	}

	return err
}
