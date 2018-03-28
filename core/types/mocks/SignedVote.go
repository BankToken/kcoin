// Code generated by mockery v1.0.0
package mocks

import common "github.com/kowala-tech/kUSD/common"
import mock "github.com/stretchr/testify/mock"
import types "github.com/kowala-tech/kUSD/core/types"

// SignedVote is an autogenerated mock type for the SignedVote type
type SignedVote struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *SignedVote) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// Vote provides a mock function with given fields:
func (_m *SignedVote) Vote() *types.Vote {
	ret := _m.Called()

	var r0 *types.Vote
	if rf, ok := ret.Get(0).(func() *types.Vote); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Vote)
		}
	}

	return r0
}