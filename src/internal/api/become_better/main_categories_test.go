package api

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"become_better/src/config"
	gen "become_better/src/gen/become_better"

	database "become_better/src/db"
	"become_better/src/internal/api/become_better/mocks"
	"become_better/src/internal/models"
)

func TestMainCategories(t *testing.T) {
	uuidID := uuid.New()

	tests := []struct {
		name           string
		mockResponse   []models.Category
		mockError      error
		expectedResult *gen.MainCategoriesResponse
		expectedError  bool
	}{
		{
			name: "success",
			mockResponse: []models.Category{
				{ID: uuidID, Name: "Category1", MainCategory: models.CategoryStudy},
			},
			mockError: nil,
			expectedResult: &gen.MainCategoriesResponse{
				MainCategories: []*gen.MainCategoriesMessage{
					{Id: uuidID.String(), Name: "Category1", MainCategory: "Учеба"},
				},
			},
			expectedError: false,
		},
		{
			name: "fail",
			mockResponse: []models.Category{},
			mockError: errors.New("some error"),
			expectedResult: &gen.MainCategoriesResponse{},
			expectedError: true,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			
			mockCategoriesService := new(mocks.MainCategoriesInterface)
			mockCategoriesService.On("MainCategories", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockResponse, tt.mockError)

			mainService := &MainService{
				MainCategoriesInterface: mockCategoriesService,
				Ctx: context.Background(),
				App: config.App{Postgres: &database.Postgres{}},
			}

			// Вызов метода
			res, err := mainService.MainCategories(context.Background(), &gen.MainCategoriesRequest{})

			// Проверка результата
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, res)
			}

			// Проверяем вызов моков
			mockCategoriesService.AssertExpectations(t)
		})
	}
}

func TestCategoriesToProto(t *testing.T) {
	id := uuid.New()
	in := []models.Category{
		{
			ID: id,
			MainCategory: models.CategoryStudy,
			Description: "desc",
			Name: "test",
		},
	}
	out := []*gen.MainCategoriesMessage{
		{
			Id: id.String(),
			Name: "test",
			Description: "desc",
			MainCategory: "Учеба",
		},
	}
	response := CategoriesToProto(in)

	assert.Equal(t, out, response)
}