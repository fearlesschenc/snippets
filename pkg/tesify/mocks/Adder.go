// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Adder is an autogenerated mock type for the Adder type
type Adder struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0, _a1
func (_m *Adder) Add(_a0 int, _a1 int) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
