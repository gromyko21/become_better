package api

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"become_better/src/config"
	database "become_better/src/db"
	gen "become_better/src/gen/become_better"
	"become_better/src/internal/models"
	"become_better/src/internal/services/mocks"
)

func TestFillProgress(t *testing.T) {
	uuidID := uuid.New()

	tests := []*struct {
		name          string
		fillData      gen.FillProgressRequest
		mockResponse  error
		expectedError bool
	}{
		{
			name: "success",
			fillData: gen.FillProgressRequest{
				CategoryId:  uuidID.String(),
				UserId:      uuidID.String(),
				Description: "test",
				Result:      10,
				Date:        "02.02.2025",
			},
			mockResponse:  nil,
			expectedError: false,
		},
		{
			name: "error categoryId",
			fillData: gen.FillProgressRequest{
				CategoryId:  "123",
				UserId:      uuidID.String(),
				Description: "test",
				Result:      10,
				Date:        "02.02.2025",
			},
			mockResponse:  fmt.Errorf(""),
			expectedError: true,
		},
		{
			name: "error categoryId",
			fillData: gen.FillProgressRequest{
				CategoryId:  uuidID.String(),
				UserId:      "123",
				Description: "test",
				Result:      10,
				Date:        "02.02.2025",
			},
			mockResponse:  fmt.Errorf(""),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockProgressService := new(mocks.ProgressInterface)
			mockProgressService.On("FillProgress", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockResponse)

			mainService := &MainService{
				ProgressInterface: mockProgressService,
				Ctx:               context.Background(),
				App:               config.App{Postgres: &database.Postgres{}},
			}

			// Вызов метода
			_, err := mainService.FillProgress(context.Background(), &tt.fillData)

			// Проверка результата
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				mockProgressService.AssertExpectations(t)
			}
		})
	}
}

func TestDeleteProgress(t *testing.T) {
	uuidID := uuid.New()

	tests := []*struct {
		name          string
		deleteProgress     gen.DeleteProgressRequest
		mockResponse  error
		expectedError bool
	}{
		{
			name: "success",
			deleteProgress: gen.DeleteProgressRequest{
				ProgressId: uuidID.String(),
				UserId: uuidID.String(),
			},
			mockResponse:  nil,
			expectedError: false,
		},
		{
			name: "error UserID",
			deleteProgress: gen.DeleteProgressRequest{
				UserId:      "",
				ProgressId:      uuidID.String(),
			},
			mockResponse:  fmt.Errorf(""),
			expectedError: true,
		},
		{
			name: "error ProgressId",
			deleteProgress: gen.DeleteProgressRequest{
				UserId:      uuidID.String(),
				ProgressId:  "",
			},
			mockResponse:  fmt.Errorf(""),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockProgressService := new(mocks.ProgressInterface)
			mockProgressService.On("DeleteProgress", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockResponse)

			mainService := &MainService{
				ProgressInterface: mockProgressService,
				Ctx:               context.Background(),
				App:               config.App{Postgres: &database.Postgres{}},
			}

			// Вызов метода
			_, err := mainService.DeleteProgress(context.Background(), &tt.deleteProgress)

			// Проверка результата
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				mockProgressService.AssertExpectations(t)
			}
		})
	}
}

func TestGetProgress(t *testing.T) {
	uuidID := uuid.New()

	tests := []*struct {
		name         string
		getProgress  gen.GetProgressRequest
		mockResponse []*models.Progress
		mockCount    int32
		mockError    error
		expectedErr  bool
	}{
		{
			name: "success",
			getProgress: gen.GetProgressRequest{
				CategoryId: uuidID.String(),
				UserId:     uuidID.String(),
				Page:       1,
				Limit:      10,
			},
			mockResponse: []*models.Progress{},
			mockCount:    0,
			mockError:    nil,
			expectedErr:  false,
		},
		{
			name: "invalid UserID",
			getProgress: gen.GetProgressRequest{
				CategoryId: uuidID.String(),
				UserId:     "invalid-uuid",
			},
			mockResponse: nil,
			mockCount:    0,
			mockError:    fmt.Errorf("can't define user_id"),
			expectedErr:  true,
		},
		{
			name: "invalid CategoryID",
			getProgress: gen.GetProgressRequest{
				CategoryId: "invalid-uuid",
				UserId:     uuidID.String(),
			},
			mockResponse: nil,
			mockCount:    0,
			mockError:    fmt.Errorf("can't define categoryID"),
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockProgressService := new(mocks.ProgressInterface)
			mockProgressService.On("GetProgress", mock.Anything, mock.Anything, mock.Anything).
				Return(tt.mockResponse, tt.mockCount, tt.mockError)

			mainService := &MainService{
				ProgressInterface: mockProgressService,
				Ctx:               context.Background(),
				App:               config.App{Postgres: &database.Postgres{}},
			}

			_, err := mainService.GetProgress(context.Background(), &tt.getProgress)

			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				mockProgressService.AssertExpectations(t)
			}
		})
	}
}
