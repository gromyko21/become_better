package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	gen "become_better/src/gen/become_better"
	"become_better/src/internal/api/models"
)

type MainCategoriesInterface interface {
	MainCategories(ctx context.Context, pool *pgxpool.Pool) ([]*gen.MainCategories, error) 
}

type CategoriesServiceImpl struct{
	models.CategoriesModelInterface
}

func (c *CategoriesServiceImpl) MainCategories(ctx context.Context, pool *pgxpool.Pool) ([]*gen.MainCategories, error) {

	categories, err := c.CategoriesModelInterface.GetCategories(ctx, pool)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return CategoriesToProto(categories), nil
}

func CategoriesToProto(categories []models.Categories) []*gen.MainCategories {
	var response []*gen.MainCategories
	for _, category := range categories {
		response = append(response,
			&gen.MainCategories{
				Id:          category.ID.String(),
				Name:        category.Name,
				Description: category.Description,
				MainCategory: models.MainCategoriesMap[category.MainCategory],
			},
		)
	}
	return response
}
