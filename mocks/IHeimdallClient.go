// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	bor "github.com/maticnetwork/bor/consensus/bor"
	mock "github.com/stretchr/testify/mock"
)

// IHeimdallClient is an autogenerated mock type for the IHeimdallClient type
type IHeimdallClient struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: path, query
func (_m *IHeimdallClient) Fetch(path string, query string) (*bor.ResponseWithHeight, error) {
	ret := _m.Called(path, query)

	var r0 *bor.ResponseWithHeight
	if rf, ok := ret.Get(0).(func(string, string) *bor.ResponseWithHeight); ok {
		r0 = rf(path, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bor.ResponseWithHeight)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchWithRetry provides a mock function with given fields: path, query
func (_m *IHeimdallClient) FetchWithRetry(path string, query string) (*bor.ResponseWithHeight, error) {
	ret := _m.Called(path, query)

	var r0 *bor.ResponseWithHeight
	if rf, ok := ret.Get(0).(func(string, string) *bor.ResponseWithHeight); ok {
		r0 = rf(path, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bor.ResponseWithHeight)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}