// Code generated by mockery. DO NOT EDIT.

package mockseth

import (
	context "context"

	etherman "github.com/0xPolygonHermez/zkevm-node/etherman"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// BlockRetriever is an autogenerated mock type for the BlockRetriever type
type BlockRetriever struct {
	mock.Mock
}

type BlockRetriever_Expecter struct {
	mock *mock.Mock
}

func (_m *BlockRetriever) EXPECT() *BlockRetriever_Expecter {
	return &BlockRetriever_Expecter{mock: &_m.Mock}
}

// RetrieveFullBlockForEvent provides a mock function with given fields: ctx, vLog
func (_m *BlockRetriever) RetrieveFullBlockForEvent(ctx context.Context, vLog types.Log) (*etherman.Block, error) {
	ret := _m.Called(ctx, vLog)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveFullBlockForEvent")
	}

	var r0 *etherman.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Log) (*etherman.Block, error)); ok {
		return rf(ctx, vLog)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Log) *etherman.Block); ok {
		r0 = rf(ctx, vLog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*etherman.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Log) error); ok {
		r1 = rf(ctx, vLog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockRetriever_RetrieveFullBlockForEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveFullBlockForEvent'
type BlockRetriever_RetrieveFullBlockForEvent_Call struct {
	*mock.Call
}

// RetrieveFullBlockForEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - vLog types.Log
func (_e *BlockRetriever_Expecter) RetrieveFullBlockForEvent(ctx interface{}, vLog interface{}) *BlockRetriever_RetrieveFullBlockForEvent_Call {
	return &BlockRetriever_RetrieveFullBlockForEvent_Call{Call: _e.mock.On("RetrieveFullBlockForEvent", ctx, vLog)}
}

func (_c *BlockRetriever_RetrieveFullBlockForEvent_Call) Run(run func(ctx context.Context, vLog types.Log)) *BlockRetriever_RetrieveFullBlockForEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.Log))
	})
	return _c
}

func (_c *BlockRetriever_RetrieveFullBlockForEvent_Call) Return(_a0 *etherman.Block, _a1 error) *BlockRetriever_RetrieveFullBlockForEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BlockRetriever_RetrieveFullBlockForEvent_Call) RunAndReturn(run func(context.Context, types.Log) (*etherman.Block, error)) *BlockRetriever_RetrieveFullBlockForEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewBlockRetriever creates a new instance of BlockRetriever. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockRetriever(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockRetriever {
	mock := &BlockRetriever{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}