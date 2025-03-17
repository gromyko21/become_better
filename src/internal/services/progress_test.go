package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	gen "become_better/src/gen/become_better"

	"become_better/src/internal/models"
	"become_better/src/internal/models/mocks"
)

func TestFillProgress(t *testing.T) {
	// TODO: починить тест
	connString := ""
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}
	defer pool.Close()

	// now := time.Now()
	// ID := uuid.New()

	tests := []struct {
		name                           string
		progress                       models.Progress
		mockCategoriesResponseError    error
		mockCategoriesResponse         int32
		mockProgressModelResponseError error
		expectedResult                 error
		expectedError                  bool
	}{
		{
			name: "fail in validateProgressDate",
			// progress: models.Progress{Date: "fff"},
			mockCategoriesResponseError: nil,
			expectedResult:              fmt.Errorf("не удалось определить дату в формате DD.MM.YYYY(%s): parsing time \"fff\" as \"02.01.2006\": cannot parse \"fff\" as \"02\"", "fff"),
			expectedError:               true,
		},
		{
			name: "fail in CategoryByID",
			// progress: models.Progress{Date: fmt.Sprintf("%02d.%02d.%d", now.Day(), now.Month(), now.Year())},
			mockCategoriesResponseError: fmt.Errorf("some error"),
			expectedResult:              fmt.Errorf("some error"),
			expectedError:               true,
		},
		{
			name: "fail in ID category",
			// progress: models.Progress{Date: fmt.Sprintf("%02d.%02d.%d", now.Day(), now.Month(), now.Year())},
			mockCategoriesResponseError:    nil,
			mockProgressModelResponseError: nil,
			expectedResult:                 fmt.Errorf("category with ID 00000000-0000-0000-0000-000000000000 doesn't exists"),
			expectedError:                  true,
		},
		{
			name: "success",
			// progress: models.Progress{Date: fmt.Sprintf("%02d.%02d.%d", now.Day(), now.Month(), now.Year()), CategoryID: ID},
			mockCategoriesResponseError:    nil,
			mockCategoriesResponse:         models.MinuteCategoryType,
			mockProgressModelResponseError: nil,
			expectedResult:                 nil,
			expectedError:                  false,
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
				ProgressModelInterface:   mockProgressModelInterface,
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
		name                           string
		progressID                     uuid.UUID
		userID                         uuid.UUID
		mockProgressModelResponseError error
		expectedError                  bool
	}{
		{
			name:                           "success",
			progressID:                     progressID,
			userID:                         userID,
			mockProgressModelResponseError: nil,
			expectedError:                  false,
		},
		{
			name:                           "error",
			progressID:                     progressID,
			userID:                         userID,
			mockProgressModelResponseError: fmt.Errorf("some error"),
			expectedError:                  true,
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

func TestPrepareDate(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "valid date",
			input:       "15.03.2024",
			expected:    "2024-03-15",
			expectError: false,
		},
		{
			name:        "invalid format",
			input:       "2024-03-15",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty string",
			input:       "",
			expected:    "",
			expectError: true,
		},
		{
			name:        "invalid date",
			input:       "32.12.2024", // Неверный день
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := prepareDate(tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestProgressToGetProgressResponse(t *testing.T) {
	validUUID := uuid.New()
	tests := []struct {
		name        string
		progress    []*models.Progress
		filter      models.ProgressFilter
		countRows   int32
		expected    *gen.GetProgressResponse
		expectError bool
	}{
		{
			name: "valid progress data",
			progress: []*models.Progress{
				{
					ID:           validUUID,
					CategoryID:   validUUID,
					UserID:       validUUID,
					Date:         time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
					ProgressType: models.MinuteCategoryType,
					Result:       100,
					Description:  "Test progress",
				},
			},
			filter: models.ProgressFilter{
				Page:  1,
				Limit: 10,
			},
			countRows: 15,
			expected: &gen.GetProgressResponse{
				Page:       1,
				Limit:      10,
				CountPages: 2,
				Progress: []*gen.Progress{
					{
						ID:                validUUID.String(),
						CategoryId:        validUUID.String(),
						UserId:            validUUID.String(),
						Date:              "2024-03-15 00:00:00 +0000 UTC",
						ProgressType:      "some_type",
						ResultInt:         100,
						ResultDescription: "Test progress",
					},
				},
			},
			expectError: false,
		},
		{
			name: "invalid progress type",
			progress: []*models.Progress{
				{
					ID:           validUUID,
					CategoryID:   validUUID,
					UserID:       validUUID,
					Date:         time.Now(),
					ProgressType: 999, // Несуществующий тип
					Result:       50,
					Description:  "Invalid type test",
				},
			},
			filter: models.ProgressFilter{
				Page:  1,
				Limit: 10,
			},
			countRows:   5,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "empty progress list",
			progress:    []*models.Progress{},
			filter:      models.ProgressFilter{Page: 1, Limit: 5},
			countRows:   0,
			expected:    &gen.GetProgressResponse{Page: 1, Limit: 5, CountPages: 0, Progress: []*gen.Progress{}},
			expectError: false,
		},
	}

	// Добавляем тестовые значения в ProgressTypesMap
	models.ProgressTypesMap[1] = "some_type"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ProgressToGetProgressResponse(tt.progress, tt.filter, tt.countRows)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Page, result.Page)
				assert.Equal(t, tt.expected.Limit, result.Limit)
				assert.Equal(t, tt.expected.CountPages, result.CountPages)
				assert.Equal(t, len(tt.expected.Progress), len(result.Progress))

				if len(result.Progress) > 0 {
					assert.Equal(t, tt.expected.Progress[0].ID, result.Progress[0].ID)
					assert.Equal(t, tt.expected.Progress[0].CategoryId, result.Progress[0].CategoryId)
					assert.Equal(t, tt.expected.Progress[0].UserId, result.Progress[0].UserId)
					assert.Equal(t, tt.expected.Progress[0].Date, result.Progress[0].Date)
					assert.Equal(t, tt.expected.Progress[0].ProgressType, result.Progress[0].ProgressType)
					assert.Equal(t, tt.expected.Progress[0].ResultInt, result.Progress[0].ResultInt)
					assert.Equal(t, tt.expected.Progress[0].ResultDescription, result.Progress[0].ResultDescription)
				}
			}
		})
	}
}

func TestGetProgress(t *testing.T) {
	validUUID := uuid.New()
	mockModel := new(mocks.ProgressModelInterface)
	service := &ProgressService{ProgressModelInterface: mockModel}

	validFilter := models.ProgressFilter{
		DateFrom: "01.03.2024",
		DateTo:   "10.03.2024",
	}

	invalidFilter := models.ProgressFilter{
		DateFrom: "10.03.2024",
		DateTo:   "01.03.2024",
	}

	mockProgress := []*models.Progress{
		{ID: validUUID, CategoryID: validUUID, UserID: validUUID},
	}

	tests := []struct {
		name        string
		filter      models.ProgressFilter
		mockReturn  []*models.Progress
		mockPages   int32
		mockErr     error
		expectErr   bool
		expectedLen int
	}{
		{
			name:        "valid request",
			filter:      validFilter,
			mockReturn:  mockProgress,
			mockPages:   2,
			mockErr:     nil,
			expectErr:   false,
			expectedLen: 1,
		},
		{
			name:        "invalid date range",
			filter:      invalidFilter,
			mockReturn:  nil,
			mockPages:   0,
			mockErr:     errors.New("dateFrom (10.03.2024) cannot be greater than dateTo (01.03.2024)"),
			expectErr:   true,
			expectedLen: 0,
		},
		{
			name: "database error",
			filter: models.ProgressFilter{
				DateFrom: "03.03.2024",
				DateTo:   "10.03.2024",
			},
			mockReturn:  nil,
			mockPages:   0,
			mockErr:     errors.New("db error"),
			expectErr:   true,
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockModel.On("GetProgress", context.Background(), mock.Anything, tt.filter).
				Return(tt.mockReturn, tt.mockPages, tt.mockErr)

			result, count, err := service.GetProgress(context.Background(), &pgxpool.Pool{}, tt.filter)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Equal(t, int32(0), count)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedLen, len(result))
				assert.Equal(t, tt.mockPages, count)
				mockModel.AssertExpectations(t)
			}

		})
	}
}
