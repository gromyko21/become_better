package api

import (
	"context"
	"fmt"

	gen "become_better/src/gen/become_better"
	"become_better/src/internal/models"
	"become_better/src/internal/services"

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
		CategoryID:  categoryID,
		UserID:      userID,
		Description: fillData.Description,
		Result:      fillData.Result,
		// TODO: починить
		// Date:        fillData.Date,
	}

	err = s.ProgressInterface.FillProgress(ctx, s.App.Postgres.Pool, &progress)
	if err != nil {
		return nil, err
	}

	return &gen.EmptyResponse{}, nil
}

func (s *MainService) DeleteProgress(ctx context.Context, deleteProgress *gen.DeleteProgressRequest) (*gen.EmptyResponse, error) {
	progressID, err := uuid.Parse(deleteProgress.ProgressId)
	if err != nil {
		return nil, fmt.Errorf("не удалось определить progress_id(%s), как uuid: %v", deleteProgress.ProgressId, err)
	}

	userID, err := uuid.Parse(deleteProgress.UserId)
	if err != nil {
		return nil, fmt.Errorf("не удалось определить user_id(%s), как uuid: %v", deleteProgress.UserId, err)
	}

	err = s.ProgressInterface.DeleteProgress(ctx, s.App.Postgres.Pool, progressID, userID)
	if err != nil {
		return nil, err
	}

	return &gen.EmptyResponse{}, nil
}

func (s *MainService) GetProgress(ctx context.Context, getProgress *gen.GetProgressRequest) (*gen.GetProgressResponse, error) {

	var categoryID *uuid.UUID
	if getProgress.CategoryId != "" {
		parsedCategoryID, err := uuid.Parse(getProgress.CategoryId)
		if err != nil {
			return nil, fmt.Errorf("can't define categoryID(%s), as uuid: %v", getProgress.CategoryId, err)
		}
		categoryID = &parsedCategoryID
	}

	userID, err := uuid.Parse(getProgress.UserId)
	if err != nil {
		return nil, fmt.Errorf("can't define  user_id(%s), as uuid: %v", getProgress.UserId, err)
	}

	filter := models.ProgressFilter{
		CategoryID: categoryID,
		UserID: &userID,
		DateFrom: getProgress.DateFrom,
		DateTo: getProgress.DateTo,
		Page: getProgress.Page,
		Limit: getProgress.Limit,
	}
	progress, countRows, err := s.ProgressInterface.GetProgress(ctx, s.App.Postgres.Pool, filter)
	if err != nil {
		return nil, fmt.Errorf("error when ProgressInterface.GetProgress: %v",  err)
	}

	response, err := services.ProgressToGetProgressResponse(progress, filter, countRows)
	if err != nil {
		return nil, fmt.Errorf("somethink went wrong: %v",  err)
	}	
	return response, nil
}
