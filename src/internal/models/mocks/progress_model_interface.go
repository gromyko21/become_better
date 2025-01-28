// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	models "become_better/src/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
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
