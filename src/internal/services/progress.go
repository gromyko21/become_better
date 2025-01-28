package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"become_better/src/internal/models"
)

type ProgressInterface interface {
	FillProgress(ctx context.Context, pool *pgxpool.Pool, progress *models.Progress) (error)
}

type ProgressService struct {
	models.ProgressModelInterface
	models.CategoriesModelInterface
}

func (p *ProgressService) FillProgress(ctx context.Context, pool *pgxpool.Pool, progress *models.Progress) (error) {

	err := validateProgressDate(progress.Date)
	if err != nil {
		logrus.Error(err)
		return err
	}
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

func validateProgressDate(progressDate string) (error) {
	layout := "02.01.2006" 
	dateField, err := time.Parse(layout, progressDate)
	if err != nil {
		return fmt.Errorf("не удалось определить дату в формате DD.MM.YYYY(%s): %v", progressDate, err)
	}

	now := time.Now()
	onlyDate := time.Date(now.Year(), now.Month(), now.Day(), 24, 59, 0, 0, now.Location())

	if dateField.After(onlyDate) {
		return fmt.Errorf("дата(%s) не может быть старше сегодняшнего дня", progressDate)
	}

	return nil
}