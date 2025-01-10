package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"become_better/src/internal/api/models"
)

type MainCategoriesInterface interface {
	MainCategories(ctx context.Context, pool *pgxpool.Pool) ([]models.Category, error)
	AddCategories(ctx context.Context, pool *pgxpool.Pool, category *models.Category) (*models.Category, error)  
}

type CategoriesServiceImpl struct{
	models.CategoriesModelInterface
}

func (c *CategoriesServiceImpl) MainCategories(ctx context.Context, pool *pgxpool.Pool) ([]models.Category, error) {

	categories, err := c.CategoriesModelInterface.GetCategories(ctx, pool)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return categories, nil
}

func (c *CategoriesServiceImpl) AddCategories(ctx context.Context, pool *pgxpool.Pool, category *models.Category) (*models.Category, error)  {

	category.ID = uuid.New()
	category, err := c.CategoriesModelInterface.AddCategory(ctx, pool, category)
	if err != nil {
		logrus.Error(err)
		return category, err
	}

	return category, nil
}