package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	gen "become_better/src/gen/become_better"
	"become_better/src/internal/api/models"
	"become_better/src/internal/api/services/mocks"
)

func TestMainCategories(t *testing.T) {

	connString := ""
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}
	defer pool.Close()

	tests := []struct {
		name           string
		mockResponse   []models.Categories
		mockError      error
		expectedResult []*gen.MainCategories
		expectedError  bool
	}{
		{
			name:           "fail",
			mockResponse:   []models.Categories{},
			mockError:      errors.New("some error"),
			expectedResult: []*gen.MainCategories{},
			expectedError:  true,
		},
		{
			name:           "success",
			mockResponse:   nil,
			mockError:      nil,
			expectedResult: nil,
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			mockCategoriesModelInterface := new(mocks.CategoriesModelInterface)
			mockCategoriesModelInterface.On("GetCategories", mock.Anything, mock.Anything).
				Return(tt.mockResponse, tt.mockError)

			categoriesService := CategoriesServiceImpl{
				CategoriesModelInterface: mockCategoriesModelInterface,
			}

			// Вызов метода
			res, err := categoriesService.MainCategories(ctx, pool)

			// Проверка результата
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, res)
			}

			// Проверяем вызов моков
			mockCategoriesModelInterface.AssertExpectations(t)
		})
	}
}


func TestCategoriesToProto(t *testing.T) {
	id := uuid.New()
	in := []models.Categories{
		{
			ID: id,
			MainCategory: models.CategoryStudy,
			Description: "desc",
			Name: "test",
		},
	}
	out := []*gen.MainCategories{
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