package models

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"	
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProgressModelInterface interface {
	AddProgress(ctx context.Context, pool *pgxpool.Pool, progress *Progress) (error)  
	DeleteProgress(ctx context.Context, pool *pgxpool.Pool, progressID, userID uuid.UUID) (error)
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

func (c *CategoriesModelImpl) DeleteProgress(ctx context.Context, pool *pgxpool.Pool, progressID, userID uuid.UUID) (error) {

	query := sq.
		Delete("progress").
		Where(sq.Eq{"id": progressID, "user_id": userID}).
		PlaceholderFormat(sq.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		logrus.Error("DeleteProgress failed to generate SQL query:", err)
	}
	
	_, err = pool.Exec(ctx, sqlQuery, args...)
	if err != nil {
		logrus.Error("DeleteProgress failed to execute query:", err)
	}

	return err
}