package api

import (
	"context"
	
	gen "become_better/gen/become_better"
)

func (s *HelloService) MainCategories(ctx context.Context, req *gen.MainCategoriesRequest) (*gen.MainCategoriesResponse, error) {
    return &gen.MainCategoriesResponse{
        MainCategories: []*gen.MainCategories{},
    }, nil
}