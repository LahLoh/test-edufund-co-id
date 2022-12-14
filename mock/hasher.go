// Code generated by mockery v2.14.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// Hasher is an autogenerated mock type for the Hasher type
type Hasher struct {
	mock.Mock
}

// Hash provides a mock function with given fields: str
func (_m *Hasher) Hash(str string) (string, error) {
	ret := _m.Called(str)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Verify provides a mock function with given fields: str, hash
func (_m *Hasher) Verify(str string, hash string) bool {
	ret := _m.Called(str, hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(str, hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewHasher interface {
	mock.TestingT
	Cleanup(func())
}

// NewHasher creates a new instance of Hasher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHasher(t mockConstructorTestingTNewHasher) *Hasher {
	mock := &Hasher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
