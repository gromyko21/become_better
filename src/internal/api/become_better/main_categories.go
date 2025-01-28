package api

import (
	"context"

	gen "become_better/src/gen/become_better"
	"become_better/src/internal/models"
)

func (s *MainService) MainCategories(ctx context.Context, req *gen.MainCategoriesRequest) (*gen.MainCategoriesResponse, error) {
	categories, err := s.MainCategoriesInterface.MainCategories(ctx, s.App.Postgres.Pool)
	if err != nil {
		return nil, err
	}

	return &gen.MainCategoriesResponse{
		MainCategories: CategoriesToProto(categories),
	}, nil
}
func CategoriesToProto(categories []models.Category) []*gen.MainCategoriesMessage {
	var response []*gen.MainCategoriesMessage
	for _, category := range categories {
		response = append(response,
			&gen.MainCategoriesMessage{
				Id:          category.ID.String(),
				Name:        category.Name,
				Description: category.Description,
				MainCategory: models.MainCategoriesMap[category.MainCategory],
			},
		)
	}
	return response
}
