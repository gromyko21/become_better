package services

import (
	"context"
	"fmt"
	"time"

	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"become_better/src/internal/models"
	"become_better/src/internal/models/mocks"
)

func TestFillProgress(t *testing.T) {

	connString := ""
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}
	defer pool.Close()

	now := time.Now()
	ID := uuid.New()

	tests := []struct {
		name           string
		progress       models.Progress
		mockCategoriesResponseError   error
		mockCategoriesResponse   int32
		mockProgressModelResponseError   error
		expectedResult error
		expectedError  bool
	}{
		{
			name:           "fail in validateProgressDate",
			progress: models.Progress{Date: "fff"},
			mockCategoriesResponseError:   nil,
			expectedResult: fmt.Errorf("не удалось определить дату в формате DD.MM.YYYY(%s): parsing time \"fff\" as \"02.01.2006\": cannot parse \"fff\" as \"02\"", "fff"),
			expectedError:  true,
		},
		{
			name:           "fail in CategoryByID",
			progress: models.Progress{Date: fmt.Sprintf("%02d.%02d.%d", now.Day(), now.Month(), now.Year())},
			mockCategoriesResponseError:   fmt.Errorf("some error"),
			expectedResult: fmt.Errorf("some error"),
			expectedError:  true,
		},
		{
			name:           "fail in ID category",
			progress: models.Progress{Date: fmt.Sprintf("%02d.%02d.%d", now.Day(), now.Month(), now.Year())},
			mockCategoriesResponseError: nil,
			mockProgressModelResponseError: nil,
			expectedResult: fmt.Errorf("category with ID 00000000-0000-0000-0000-000000000000 doesn't exists"),
			expectedError:  true,
		},
		{
			name:           "success",
			progress: models.Progress{Date: fmt.Sprintf("%02d.%02d.%d", now.Day(), now.Month(), now.Year()), CategoryID: ID},
			mockCategoriesResponseError: nil,
			mockCategoriesResponse: models.MinuteCategoryType,
			mockProgressModelResponseError: nil,
			expectedResult: nil,
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			

			mockCategoriesModelInterface := new(mocks.CategoriesModelInterface)
			mockCategoriesModelInterface.On("CategoryTypeByID", mock.Anything, mock.Anything, tt.progress.CategoryID).
				Return(tt.mockCategoriesResponse, tt.mockCategoriesResponseError)

			mockProgressModelInterface := new(mocks.ProgressModelInterface)
			mockProgressModelInterface.On("AddProgress", mock.Anything, mock.Anything, &tt.progress).
				Return(tt.mockProgressModelResponseError)

			progressService := ProgressService{
				CategoriesModelInterface: mockCategoriesModelInterface,
				ProgressModelInterface: mockProgressModelInterface,
			}

			// Вызов метода
			err := progressService.FillProgress(ctx, pool, &tt.progress)

			// Проверка результата
			if tt.expectedError {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedResult, err)
			} else {
				assert.NoError(t, err)
				mockCategoriesModelInterface.AssertExpectations(t)
				mockProgressModelInterface.AssertExpectations(t)
			}
		})
	}
}
