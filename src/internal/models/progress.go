package models

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type ProgressModelInterface interface {
	AddProgress(ctx context.Context, pool *pgxpool.Pool, progress *FillProgress) error
	DeleteProgress(ctx context.Context, pool *pgxpool.Pool, progressID, userID uuid.UUID) error
	GetProgress(ctx context.Context, pool *pgxpool.Pool, filter ProgressFilter) ([]*Progress, int32, error)
}

type ProgressFilter struct {
	CategoryID *uuid.UUID
	UserID     *uuid.UUID
	DateFrom   string
	DateTo     string
	Page       int32
	Limit      int32
}

type ProgressModelImpl struct{}

func (c *CategoriesModelImpl) AddProgress(ctx context.Context, pool *pgxpool.Pool, progress *FillProgress) error {

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

func (c *CategoriesModelImpl) DeleteProgress(ctx context.Context, pool *pgxpool.Pool, progressID, userID uuid.UUID) error {

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

func (c *CategoriesModelImpl) GetProgress(ctx context.Context, pool *pgxpool.Pool, progressFilter ProgressFilter) ([]*Progress, int32, error) {
	var progresses []*Progress
	var totalCount int

	query := sq.
		Select("id, user_id, category_id, progress_type, result_int, result_description, date, COUNT(*) OVER() AS total_count").
		From("progress")

	if progressFilter.DateTo != "" {
		query = query.Where(sq.LtOrEq{"date": progressFilter.DateTo})
	}
	if progressFilter.DateFrom != "" {
		query = query.Where(sq.GtOrEq{"date": progressFilter.DateFrom})
	}
	if progressFilter.UserID != nil {
		query = query.Where(sq.Eq{"user_id": progressFilter.UserID})
	}
	if progressFilter.CategoryID != nil {
		query = query.Where(sq.Eq{"category_id": progressFilter.CategoryID})
	}

	if progressFilter.Page < 1 {
		progressFilter.Page = 1
	}
	if progressFilter.Limit <= 0 {
		progressFilter.Limit = 10
	}
	offset := (progressFilter.Page - 1) * progressFilter.Limit

	query = query.Limit(uint64(progressFilter.Limit)).Offset(uint64(offset))

	query = query.PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		logrus.Error(err)
		return nil, 0, err
	}

	rows, err := pool.Query(ctx, sql, args...)
	if err != nil {
		logrus.Error(err)
		return nil, 0, err
	}
	defer rows.Close()

	firstRow := true
	for rows.Next() {
		var p Progress
		var count int
		err = rows.Scan(&p.ID, &p.UserID, &p.CategoryID, &p.ProgressType, &p.Result, &p.Description, &p.Date, &count)
		if err != nil {
			logrus.Error(err)
			return nil, 0, err
		}

		if firstRow {
			totalCount = count
			firstRow = false
		}

		progresses = append(progresses, &p)
	}

	return progresses, int32(totalCount), nil
}
