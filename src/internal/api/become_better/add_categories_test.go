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
	"become_better/src/internal/models"
	"become_better/src/internal/services/mocks"
)

func TestAddCategories(t *testing.T) {
	uuidID := uuid.New()

	tests := []*struct {
		name           string
		category       gen.AddCategoryMessage
		mockResponse   models.Category
		mockError      error
		expectedResult *gen.MainCategoriesMessage
		expectedError  bool
	}{
		{
			name:         "success",
			mockResponse: models.Category{ID: uuidID, Name: "Category1", MainCategory: models.CategoryStudy, Description: "desc"},
			category: gen.AddCategoryMessage{
				MainCategory: models.CategoryStudy,
				CategoryType: models.CountCategoryType,
				Name:         "Category1",
				Description:  "desc",
			},
			mockError: nil,
			expectedResult: &gen.MainCategoriesMessage{
				Id:           uuidID.String(),
				Name:         "Category1",
				MainCategory: "Учеба",
				Description:  "desc",
			},
			expectedError: false,
		},
		{
			name:           "fail",
			mockResponse:   models.Category{},
			mockError:      errors.New("some error"),
			expectedResult: &gen.MainCategoriesMessage{},
			expectedError:  true,
		},
		{
			name:           "fail main category",
			category:       gen.AddCategoryMessage{MainCategory: 8800555},
			mockResponse:   models.Category{},
			mockError:      nil,
			expectedResult: &gen.MainCategoriesMessage{},
			expectedError:  true,
		},
		{
			name: "fail category type",
			category: gen.AddCategoryMessage{
				MainCategory: models.CategoryStudy,
				CategoryType: 8800,
			},
			mockResponse:   models.Category{},
			mockError:      nil,
			expectedResult: &gen.MainCategoriesMessage{},
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockCategoriesService := new(mocks.MainCategoriesInterface)
			mockCategoriesService.On("AddCategories", mock.Anything, mock.Anything, mock.Anything).
				Return(&tt.mockResponse, tt.mockError)

			mainService := &MainService{
				MainCategoriesInterface: mockCategoriesService,
				Ctx:                     context.Background(),
				App:                     config.App{Postgres: &database.Postgres{}},
			}

			// Вызов метода
			res, err := mainService.AddCategories(context.Background(), &tt.category)

			// Проверка результата
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, res)

				// Проверяем вызов моков
				mockCategoriesService.AssertExpectations(t)
			}
		})
	}

}
