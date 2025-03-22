package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	gen "become_better/src/gen/become_better"

	"become_better/src/internal/models"
	"become_better/src/internal/utils"
)

type ProgressInterface interface {
	FillProgress(ctx context.Context, pool *pgxpool.Pool, progress *models.FillProgress) error
	DeleteProgress(ctx context.Context, pool *pgxpool.Pool, progressID, userID uuid.UUID) error
	GetProgress(ctx context.Context, pool *pgxpool.Pool, filter models.ProgressFilter) ([]*models.Progress, int32, error)
	GetProgressByCategory(ctx context.Context, pool *pgxpool.Pool, filter models.ProgressByCategoryFilter) (*models.ProgressByCategory, error)
}

type ProgressService struct {
	models.ProgressModelInterface
	models.CategoriesModelInterface
}

func (p *ProgressService) FillProgress(ctx context.Context, pool *pgxpool.Pool, progress *models.FillProgress) error {
	// TODO:  вернуть и переопределить
	// err := validateProgressDate(progress.Date)
	// if err != nil {
	// 	logrus.Error(err)
	// 	return err
	// }
	categoryType, err := p.CategoriesModelInterface.CategoryTypeByID(ctx, pool, progress.CategoryID)

	if err != nil {
		return err
	}

	if categoryType == 0 {
		return fmt.Errorf("category with ID %s doesn't exists", progress.CategoryID.String())
	}

	progress.ProgressType = categoryType
	progress.ID = uuid.New()

	err = p.ProgressModelInterface.AddProgress(ctx, pool, progress)
	if err != nil {
		logrus.Error(err)

		return fmt.Errorf("не удалось создать новый прогресс: %v", err)
	}

	return nil
}

func (p *ProgressService) DeleteProgress(ctx context.Context, pool *pgxpool.Pool, progressID, userID uuid.UUID) error {

	err := p.ProgressModelInterface.DeleteProgress(ctx, pool, progressID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProgressService) GetProgress(ctx context.Context, pool *pgxpool.Pool, filter models.ProgressFilter) ([]*models.Progress, int32, error) {

	dateFrom, err := prepareDate(filter.DateFrom)
	if err != nil {
		logrus.Error(err)
		return nil, 0, err
	}

	dateTo, err := prepareDate(filter.DateTo)
	if err != nil {
		logrus.Error(err)
		return nil, 0, err
	}

	// Проверяем, что начальная дата не больше конечной
	// TODO добавить проверку не по строкам, а по дате
	if dateFrom > dateTo {
		err := fmt.Errorf("dateFrom (%s) cannot be greater than dateTo (%s)", dateFrom, dateTo)
		logrus.Error(err)
		return nil, 0, err
	}

	progress, countPages, err := p.ProgressModelInterface.GetProgress(ctx, pool, filter)
	fmt.Println(progress, countPages, err, filter)
	if err != nil {
		fmt.Println("GetProgressGetProgress")
		logrus.Error(err)
		return nil, 0, err
	}

	return progress, countPages, nil
}

func (p *ProgressService) GetProgressByCategory(
	ctx context.Context,
	pool *pgxpool.Pool,
	filter models.ProgressByCategoryFilter) (*models.ProgressByCategory, error) {

	dateFrom, err := validateProgressDate(filter.DateFrom)
	if err != nil {
		return nil, err
	}

	dateTo, err := validateProgressDate(filter.DateTo)
	if err != nil {
		return nil, err
	}
	if dateTo.Before(*dateFrom) {
		return nil, fmt.Errorf("date_from(%d) have to be later or equal then date_to(%d)", dateFrom, dateTo)
	}

	progress, err := p.ProgressModelInterface.GetProgressByCategory(ctx, pool, &filter)
	if err != nil {
		return nil, err
	}

	progress.CountDays = daysDateDiff(*dateFrom, *dateTo)
	return progress, nil
}

// Map progress to gen.GetProgressResponse
func ProgressToGetProgressResponse(
	progress []*models.Progress,
	filter models.ProgressFilter,
	countRows int32) (*gen.GetProgressResponse, error) {
	response := &gen.GetProgressResponse{Page: filter.Page, Limit: filter.Limit, CountPages: int32(utils.TotalPages(countRows, filter.Limit))}
	genProgress := []*gen.Progress{}
	for _, v := range progress {
		_, ok := models.ProgressTypesMap[v.ProgressType]
		if !ok {
			return nil, fmt.Errorf("category type with such ID - %v doesn't exist", v.ProgressType)
		}

		p := &gen.Progress{
			ID:                v.ID.String(),
			CategoryId:        v.CategoryID.String(),
			UserId:            v.UserID.String(),
			Date:              v.Date.String(),
			ProgressType:      models.ProgressTypesMap[v.ProgressType],
			ResultInt:         v.Result,
			ResultDescription: v.Description,
		}
		genProgress = append(genProgress, p)
	}

	response.Progress = genProgress

	return response, nil
}

// convert date from 02.01.2006 to 2006-01-02
func prepareDate(dateS string) (string, error) {
	date, err := validateProgressDate(dateS)
	if err != nil {
		logrus.Error(err)
		return "", err
	}

	preparedDate := fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())

	return preparedDate, nil
}

func validateProgressDate(progressDate string) (*time.Time, error) {
	layout := "02.01.2006"
	dateField, err := time.Parse(layout, progressDate)
	if err != nil {
		return nil, fmt.Errorf("не удалось определить дату в формате DD.MM.YYYY(%s): %v", progressDate, err)
	}

	now := time.Now()
	onlyDate := time.Date(now.Year(), now.Month(), now.Day(), 24, 59, 0, 0, now.Location())

	if dateField.After(onlyDate) {
		return nil, fmt.Errorf("дата(%s) не может быть старше сегодняшнего дня", progressDate)
	}

	return &dateField, nil
}

func daysDateDiff(first, second time.Time) int32 {
	date1 := time.Date(first.Year(), first.Month(), first.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(second.Year(), second.Month(), second.Day(), 0, 0, 0, 0, time.UTC)

	return int32(date2.Sub(date1).Hours() / 24)
}
