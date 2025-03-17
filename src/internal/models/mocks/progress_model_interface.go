// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	models "become_better/src/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"

	uuid "github.com/google/uuid"
)

// ProgressModelInterface is an autogenerated mock type for the ProgressModelInterface type
type ProgressModelInterface struct {
	mock.Mock
}

// AddProgress provides a mock function with given fields: ctx, pool, progress
func (_m *ProgressModelInterface) AddProgress(ctx context.Context, pool *pgxpool.Pool, progress *models.Progress) error {
	ret := _m.Called(ctx, pool, progress)

	if len(ret) == 0 {
		panic("no return value specified for AddProgress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgxpool.Pool, *models.Progress) error); ok {
		r0 = rf(ctx, pool, progress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteProgress provides a mock function with given fields: ctx, pool, progressID, userID
func (_m *ProgressModelInterface) DeleteProgress(ctx context.Context, pool *pgxpool.Pool, progressID uuid.UUID, userID uuid.UUID) error {
	ret := _m.Called(ctx, pool, progressID, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProgress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgxpool.Pool, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, pool, progressID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProgress provides a mock function with given fields: ctx, pool, filter
func (_m *ProgressModelInterface) GetProgress(ctx context.Context, pool *pgxpool.Pool, filter models.ProgressFilter) ([]*models.Progress, int32, error) {
	ret := _m.Called(ctx, pool, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetProgress")
	}

	var r0 []*models.Progress
	var r1 int32
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *pgxpool.Pool, models.ProgressFilter) ([]*models.Progress, int32, error)); ok {
		return rf(ctx, pool, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pgxpool.Pool, models.ProgressFilter) []*models.Progress); ok {
		r0 = rf(ctx, pool, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Progress)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pgxpool.Pool, models.ProgressFilter) int32); ok {
		r1 = rf(ctx, pool, filter)
	} else {
		r1 = ret.Get(1).(int32)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *pgxpool.Pool, models.ProgressFilter) error); ok {
		r2 = rf(ctx, pool, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewProgressModelInterface creates a new instance of ProgressModelInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProgressModelInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProgressModelInterface {
	mock := &ProgressModelInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
