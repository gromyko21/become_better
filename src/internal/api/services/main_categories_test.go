package services

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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
		mockResponse   []models.Category
		mockError      error
		expectedResult []models.Category
		expectedError  bool
	}{
		{
			name:           "fail",
			mockResponse:   []models.Category{},
			mockError:      errors.New("some error"),
			expectedResult: []models.Category{},
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

func TestAddCategories(t *testing.T) {

	connString := ""
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}
	defer pool.Close()
	id := uuid.New()

	tests := []struct {
		name           string
		category       models.Category
		mockResponse   models.Category
		mockError      error
		expectedResult *models.Category
		expectedError  bool
	}{
		{
			name:           "fail",
			mockResponse:   models.Category{},
			mockError:      errors.New("some error"),
			expectedResult: &models.Category{},
			expectedError:  true,
		},
		{
			name:           "success",
			mockResponse:   models.Category{
				ID: id,
				Name: "name",
				Description: "desc",
				MainCategory: 1,
			},
			mockError:      nil,
			expectedResult: &models.Category{
				ID: id,
				Name: "name",
				Description: "desc",
				MainCategory: 1,
			},
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			mockCategoriesModelInterface := new(mocks.CategoriesModelInterface)
			mockCategoriesModelInterface.On("AddCategory", mock.Anything, mock.Anything, &tt.category).
				Return(&tt.mockResponse, tt.mockError)

			categoriesService := CategoriesServiceImpl{
				CategoriesModelInterface: mockCategoriesModelInterface,
			}

			// Вызов метода
			res, err := categoriesService.AddCategories(ctx, pool, &tt.category)

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
