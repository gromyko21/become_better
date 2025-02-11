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

func Test_validateProgressDate(t *testing.T) {
	tests := []struct {
		name        string
		inputDate   string
		expectError bool
	}{
		{
			name:        "success",
			inputDate:   time.Now().Format("02.01.2006"),
			expectError: false,
		},
		{
			name:        "Error future date",
			inputDate:   time.Now().AddDate(0, 0, 1).Format("02.01.2006"),
			expectError: true,
		},
		{
			name:        "error incorrect date",
			inputDate:   "2024-03-20",
			expectError: true,
		},
		{
			name:        "error empty date",
			inputDate:   "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProgressDate(tt.inputDate)
			if (err != nil) != tt.expectError {
				t.Errorf("validateProgressDate() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestDeleteProgress(t *testing.T) {
	connString := ""
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}
	defer pool.Close()

	progressID := uuid.New()
	userID := uuid.New()

	tests := []struct {
		name                          string
		progressID                    uuid.UUID
		userID                        uuid.UUID
		mockProgressModelResponseError error
		expectedError                 bool
	}{
		{
			name:                          "success",
			progressID:                    progressID,
			userID:                        userID,
			mockProgressModelResponseError: nil,
			expectedError:                 false,
		},
		{
			name:                          "error",
			progressID:                    progressID,
			userID:                        userID,
			mockProgressModelResponseError: fmt.Errorf("some error"),
			expectedError:                 true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			mockProgressModelInterface := new(mocks.ProgressModelInterface)
			mockProgressModelInterface.On("DeleteProgress", mock.Anything, mock.Anything, tt.progressID, tt.userID).
				Return(tt.mockProgressModelResponseError)

			progressService := ProgressService{
				ProgressModelInterface: mockProgressModelInterface,
			}

			// Вызов метода
			err := progressService.DeleteProgress(ctx, pool, tt.progressID, tt.userID)

			// Проверка результата
			if tt.expectedError {
				assert.Error(t, err)
				assert.Equal(t, tt.mockProgressModelResponseError, err)
			} else {
				assert.NoError(t, err)
				mockProgressModelInterface.AssertExpectations(t)
			}
		})
	}
}
