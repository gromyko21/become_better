package main_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"become_better/config"
	gen "become_better/gen/become_better"
	api "become_better/internal/api/become_better"

	database "become_better/db"
	"become_better/internal/api/become_better/mocks"
)

func TestMainCategories(t *testing.T) {

	tests := []struct {
		name           string
		mockResponse   []*gen.MainCategories
		mockError      error
		expectedResult *gen.MainCategoriesResponse
		expectedError  bool
	}{
		{
			name: "success",
			mockResponse: []*gen.MainCategories{
				{Id: "1", Name: "Category1"},
				{Id: "2", Name: "Category2"},
			},
			mockError: nil,
			expectedResult: &gen.MainCategoriesResponse{
				MainCategories: []*gen.MainCategories{
					&gen.MainCategories{Id: "1", Name:"Category1",},
					&gen.MainCategories{Id: "2", Name:"Category2",},
				},

			},
			expectedError: false,
		},
		{
			name: "fail",
			mockResponse: []*gen.MainCategories{},
			mockError: errors.New("some error"),
			expectedResult: &gen.MainCategoriesResponse{},
			expectedError: true,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCategoriesService := new(mocks.MainCategoriesInterface)
			mockCategoriesService.On("MainCategories", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockResponse, tt.mockError)

			mainService := &api.MainService{
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