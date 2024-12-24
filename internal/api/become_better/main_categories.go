package api

import (
	"context"

	gen "become_better/gen/become_better"
)

func (s *MainService) MainCategories(ctx context.Context, req *gen.MainCategoriesRequest) (*gen.MainCategoriesResponse, error) {
	response, err := s.MainCategoriesInterface.MainCategories(ctx, s.App.Postgres.Pool)
	if err != nil {
		return nil, err
	}

	return &gen.MainCategoriesResponse{
		MainCategories: response,
	}, nil
}
