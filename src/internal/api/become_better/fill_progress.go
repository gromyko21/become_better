package api

import (
	"context"
	"fmt"

	gen "become_better/src/gen/become_better"
	"become_better/src/internal/models"

	"github.com/google/uuid"
)

func (s *MainService) FillProgress(ctx context.Context, fillData *gen.FillProgressRequest) (*gen.EmptyResponse, error) {
	categoryID, err := uuid.Parse(fillData.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("не удалось определить category_id(%s), как uuid: %v", fillData.CategoryId, err)
	}

	userID, err := uuid.Parse(fillData.UserId)
	if err != nil {
		return nil, fmt.Errorf("не удалось определить category_id(%s), как uuid: %v", fillData.UserId, err)
	}

	progress := models.Progress{
		CategoryID: categoryID,
		UserID: userID,
		Description: fillData.Description,
		Result: fillData.Result,
		Date: fillData.Date,
	}

	err = s.ProgressInterface.FillProgress(ctx, s.App.Postgres.Pool, &progress)
	if err != nil {
		return nil, err
	}
	
	return &gen.EmptyResponse{}, nil
}