// Generated by mockery with a minor update as mockery confuses import paths
package mocks

import (
	types "github.com/canonical/microceph/microceph/api/types"
	state "github.com/canonical/microcluster/state" // mockery gets confused about import paths here
	mock "github.com/stretchr/testify/mock"
)

// OSDQueryInterface is an autogenerated mock type for the OSDQueryInterface type
type OSDQueryInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: s, osd
func (_m *OSDQueryInterface) Delete(s *state.State, osd int64) error {
	ret := _m.Called(s, osd)

	var r0 error
	if rf, ok := ret.Get(0).(func(*state.State, int64) error); ok {
		r0 = rf(s, osd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HaveOSD provides a mock function with given fields: s, osd
func (_m *OSDQueryInterface) HaveOSD(s *state.State, osd int64) bool {
	ret := _m.Called(s, osd)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*state.State, int64) bool); ok {
		r0 = rf(s, osd)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// List provides a mock function with given fields: s
func (_m *OSDQueryInterface) List(s *state.State) (types.Disks, error) {
	ret := _m.Called(s)

	var r0 types.Disks
	var r1 error
	if rf, ok := ret.Get(0).(func(*state.State) (types.Disks, error)); ok {
		return rf(s)
	}
	if rf, ok := ret.Get(0).(func(*state.State) types.Disks); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Disks)
		}
	}

	if rf, ok := ret.Get(1).(func(*state.State) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Path provides a mock function with given fields: s, osd
func (_m *OSDQueryInterface) Path(s *state.State, osd int64) (string, error) {
	ret := _m.Called(s, osd)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*state.State, int64) (string, error)); ok {
		return rf(s, osd)
	}
	if rf, ok := ret.Get(0).(func(*state.State, int64) string); ok {
		r0 = rf(s, osd)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*state.State, int64) error); ok {
		r1 = rf(s, osd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOSDQueryInterface creates a new instance of OSDQueryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOSDQueryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *OSDQueryInterface {
	mock := &OSDQueryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
