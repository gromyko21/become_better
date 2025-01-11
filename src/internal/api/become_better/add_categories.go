package api

import (
	"context"
	"fmt"

	gen "become_better/src/gen/become_better"
	"become_better/src/internal/api/models"

)

func (s *MainService) AddCategories(ctx context.Context, newCategory *gen.AddCategoryMessage) (*gen.MainCategoriesMessage, error) {


	_, ok := models.MainCategoriesMap[newCategory.MainCategory]
	if !ok {
		return nil, fmt.Errorf("category with such ID - %v doesn't exist", newCategory.MainCategory)
	}

	category := models.Category{
		MainCategory: newCategory.MainCategory,
		Name: newCategory.Name,
		Description: newCategory.Description,
	}
	createdCategory, err := s.MainCategoriesInterface.AddCategories(ctx, s.App.Postgres.Pool, &category)
	if err != nil {
		return nil, err
	}

	return &gen.MainCategoriesMessage{
		Id: createdCategory.ID.String(),
		Name: createdCategory.Name,
		Description: createdCategory.Description,
		MainCategory: models.MainCategoriesMap[createdCategory.MainCategory],
	}, nil
}