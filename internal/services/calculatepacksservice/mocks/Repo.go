// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// FindAvailablePacks provides a mock function with given fields:
func (_m *Repo) FindAvailablePacks() []int {
	ret := _m.Called()

	var r0 []int
	if rf, ok := ret.Get(0).(func() []int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// NewRepo creates a new instance of Repo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repo {
	mock := &Repo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
