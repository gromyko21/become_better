package api

import (
	"context"
	"fmt"

	gen "become_better/src/gen/become_better"
	"become_better/src/internal/models"
)

func (s *MainService) AddCategories(ctx context.Context, newCategory *gen.AddCategoryMessage) (*gen.MainCategoriesMessage, error) {

	_, ok := models.MainCategoriesMap[newCategory.MainCategory]
	if !ok {
		return nil, fmt.Errorf("category with such ID - %v doesn't exist", newCategory.MainCategory)
	}

	// TODO: разобраться с неймингом. Почему тип прогресса определяем по типу категории?
	_, ok = models.ProgressTypesMap[newCategory.CategoryType]
	if !ok {
		return nil, fmt.Errorf("category type with such ID - %v doesn't exist", newCategory.CategoryType)
	}

	category := models.Category{
		MainCategory: newCategory.MainCategory,
		Name:         newCategory.Name,
		Description:  newCategory.Description,
		ProgressType: newCategory.CategoryType,
	}
	createdCategory, err := s.MainCategoriesInterface.AddCategories(ctx, s.App.Postgres.Pool, &category)
	if err != nil {
		return nil, err
	}

	return &gen.MainCategoriesMessage{
		Id:           createdCategory.ID.String(),
		Name:         createdCategory.Name,
		Description:  createdCategory.Description,
		MainCategory: models.MainCategoriesMap[createdCategory.MainCategory],
	}, nil
}
