// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	primitives "github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	query "github.com/smartcontractkit/chainlink-common/pkg/types/query"

	read "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/read"

	types "github.com/smartcontractkit/chainlink-common/pkg/types"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

type Reader_Expecter struct {
	mock *mock.Mock
}

func (_m *Reader) EXPECT() *Reader_Expecter {
	return &Reader_Expecter{mock: &_m.Mock}
}

// BatchCall provides a mock function with given fields: address, params, retVal
func (_m *Reader) BatchCall(address common.Address, params any, retVal any) (read.Call, error) {
	ret := _m.Called(address, params, retVal)

	if len(ret) == 0 {
		panic("no return value specified for BatchCall")
	}

	var r0 read.Call
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, any, any) (read.Call, error)); ok {
		return rf(address, params, retVal)
	}
	if rf, ok := ret.Get(0).(func(common.Address, any, any) read.Call); ok {
		r0 = rf(address, params, retVal)
	} else {
		r0 = ret.Get(0).(read.Call)
	}

	if rf, ok := ret.Get(1).(func(common.Address, any, any) error); ok {
		r1 = rf(address, params, retVal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reader_BatchCall_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchCall'
type Reader_BatchCall_Call struct {
	*mock.Call
}

// BatchCall is a helper method to define mock.On call
//   - address common.Address
//   - params any
//   - retVal any
func (_e *Reader_Expecter) BatchCall(address interface{}, params interface{}, retVal interface{}) *Reader_BatchCall_Call {
	return &Reader_BatchCall_Call{Call: _e.mock.On("BatchCall", address, params, retVal)}
}

func (_c *Reader_BatchCall_Call) Run(run func(address common.Address, params any, retVal any)) *Reader_BatchCall_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(any), args[2].(any))
	})
	return _c
}

func (_c *Reader_BatchCall_Call) Return(_a0 read.Call, _a1 error) *Reader_BatchCall_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Reader_BatchCall_Call) RunAndReturn(run func(common.Address, any, any) (read.Call, error)) *Reader_BatchCall_Call {
	_c.Call.Return(run)
	return _c
}

// Bind provides a mock function with given fields: _a0, _a1
func (_m *Reader) Bind(_a0 context.Context, _a1 ...common.Address) error {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Bind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...common.Address) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reader_Bind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bind'
type Reader_Bind_Call struct {
	*mock.Call
}

// Bind is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...common.Address
func (_e *Reader_Expecter) Bind(_a0 interface{}, _a1 ...interface{}) *Reader_Bind_Call {
	return &Reader_Bind_Call{Call: _e.mock.On("Bind",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *Reader_Bind_Call) Run(run func(_a0 context.Context, _a1 ...common.Address)) *Reader_Bind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]common.Address, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(common.Address)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *Reader_Bind_Call) Return(_a0 error) *Reader_Bind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reader_Bind_Call) RunAndReturn(run func(context.Context, ...common.Address) error) *Reader_Bind_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestValueWithHeadData provides a mock function with given fields: ctx, addr, confidence, params, returnVal
func (_m *Reader) GetLatestValueWithHeadData(ctx context.Context, addr common.Address, confidence primitives.ConfidenceLevel, params any, returnVal any) (*types.Head, error) {
	ret := _m.Called(ctx, addr, confidence, params, returnVal)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestValueWithHeadData")
	}

	var r0 *types.Head
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, primitives.ConfidenceLevel, any, any) (*types.Head, error)); ok {
		return rf(ctx, addr, confidence, params, returnVal)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, primitives.ConfidenceLevel, any, any) *types.Head); ok {
		r0 = rf(ctx, addr, confidence, params, returnVal)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Head)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, primitives.ConfidenceLevel, any, any) error); ok {
		r1 = rf(ctx, addr, confidence, params, returnVal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reader_GetLatestValueWithHeadData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestValueWithHeadData'
type Reader_GetLatestValueWithHeadData_Call struct {
	*mock.Call
}

// GetLatestValueWithHeadData is a helper method to define mock.On call
//   - ctx context.Context
//   - addr common.Address
//   - confidence primitives.ConfidenceLevel
//   - params any
//   - returnVal any
func (_e *Reader_Expecter) GetLatestValueWithHeadData(ctx interface{}, addr interface{}, confidence interface{}, params interface{}, returnVal interface{}) *Reader_GetLatestValueWithHeadData_Call {
	return &Reader_GetLatestValueWithHeadData_Call{Call: _e.mock.On("GetLatestValueWithHeadData", ctx, addr, confidence, params, returnVal)}
}

func (_c *Reader_GetLatestValueWithHeadData_Call) Run(run func(ctx context.Context, addr common.Address, confidence primitives.ConfidenceLevel, params any, returnVal any)) *Reader_GetLatestValueWithHeadData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(primitives.ConfidenceLevel), args[3].(any), args[4].(any))
	})
	return _c
}

func (_c *Reader_GetLatestValueWithHeadData_Call) Return(_a0 *types.Head, _a1 error) *Reader_GetLatestValueWithHeadData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Reader_GetLatestValueWithHeadData_Call) RunAndReturn(run func(context.Context, common.Address, primitives.ConfidenceLevel, any, any) (*types.Head, error)) *Reader_GetLatestValueWithHeadData_Call {
	_c.Call.Return(run)
	return _c
}

// QueryKey provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *Reader) QueryKey(_a0 context.Context, _a1 common.Address, _a2 query.KeyFilter, _a3 query.LimitAndSort, _a4 any) ([]types.Sequence, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	if len(ret) == 0 {
		panic("no return value specified for QueryKey")
	}

	var r0 []types.Sequence
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, query.KeyFilter, query.LimitAndSort, any) ([]types.Sequence, error)); ok {
		return rf(_a0, _a1, _a2, _a3, _a4)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, query.KeyFilter, query.LimitAndSort, any) []types.Sequence); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Sequence)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, query.KeyFilter, query.LimitAndSort, any) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reader_QueryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryKey'
type Reader_QueryKey_Call struct {
	*mock.Call
}

// QueryKey is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 common.Address
//   - _a2 query.KeyFilter
//   - _a3 query.LimitAndSort
//   - _a4 any
func (_e *Reader_Expecter) QueryKey(_a0 interface{}, _a1 interface{}, _a2 interface{}, _a3 interface{}, _a4 interface{}) *Reader_QueryKey_Call {
	return &Reader_QueryKey_Call{Call: _e.mock.On("QueryKey", _a0, _a1, _a2, _a3, _a4)}
}

func (_c *Reader_QueryKey_Call) Run(run func(_a0 context.Context, _a1 common.Address, _a2 query.KeyFilter, _a3 query.LimitAndSort, _a4 any)) *Reader_QueryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(query.KeyFilter), args[3].(query.LimitAndSort), args[4].(any))
	})
	return _c
}

func (_c *Reader_QueryKey_Call) Return(_a0 []types.Sequence, _a1 error) *Reader_QueryKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Reader_QueryKey_Call) RunAndReturn(run func(context.Context, common.Address, query.KeyFilter, query.LimitAndSort, any) ([]types.Sequence, error)) *Reader_QueryKey_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0
func (_m *Reader) Register(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reader_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type Reader_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Reader_Expecter) Register(_a0 interface{}) *Reader_Register_Call {
	return &Reader_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *Reader_Register_Call) Run(run func(_a0 context.Context)) *Reader_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Reader_Register_Call) Return(_a0 error) *Reader_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reader_Register_Call) RunAndReturn(run func(context.Context) error) *Reader_Register_Call {
	_c.Call.Return(run)
	return _c
}

// SetCodec provides a mock function with given fields: _a0
func (_m *Reader) SetCodec(_a0 types.RemoteCodec) {
	_m.Called(_a0)
}

// Reader_SetCodec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetCodec'
type Reader_SetCodec_Call struct {
	*mock.Call
}

// SetCodec is a helper method to define mock.On call
//   - _a0 types.RemoteCodec
func (_e *Reader_Expecter) SetCodec(_a0 interface{}) *Reader_SetCodec_Call {
	return &Reader_SetCodec_Call{Call: _e.mock.On("SetCodec", _a0)}
}

func (_c *Reader_SetCodec_Call) Run(run func(_a0 types.RemoteCodec)) *Reader_SetCodec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.RemoteCodec))
	})
	return _c
}

func (_c *Reader_SetCodec_Call) Return() *Reader_SetCodec_Call {
	_c.Call.Return()
	return _c
}

func (_c *Reader_SetCodec_Call) RunAndReturn(run func(types.RemoteCodec)) *Reader_SetCodec_Call {
	_c.Call.Return(run)
	return _c
}

// Unbind provides a mock function with given fields: _a0, _a1
func (_m *Reader) Unbind(_a0 context.Context, _a1 ...common.Address) error {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Unbind")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...common.Address) error); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reader_Unbind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unbind'
type Reader_Unbind_Call struct {
	*mock.Call
}

// Unbind is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...common.Address
func (_e *Reader_Expecter) Unbind(_a0 interface{}, _a1 ...interface{}) *Reader_Unbind_Call {
	return &Reader_Unbind_Call{Call: _e.mock.On("Unbind",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *Reader_Unbind_Call) Run(run func(_a0 context.Context, _a1 ...common.Address)) *Reader_Unbind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]common.Address, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(common.Address)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *Reader_Unbind_Call) Return(_a0 error) *Reader_Unbind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reader_Unbind_Call) RunAndReturn(run func(context.Context, ...common.Address) error) *Reader_Unbind_Call {
	_c.Call.Return(run)
	return _c
}

// Unregister provides a mock function with given fields: _a0
func (_m *Reader) Unregister(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Unregister")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reader_Unregister_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unregister'
type Reader_Unregister_Call struct {
	*mock.Call
}

// Unregister is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Reader_Expecter) Unregister(_a0 interface{}) *Reader_Unregister_Call {
	return &Reader_Unregister_Call{Call: _e.mock.On("Unregister", _a0)}
}

func (_c *Reader_Unregister_Call) Run(run func(_a0 context.Context)) *Reader_Unregister_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Reader_Unregister_Call) Return(_a0 error) *Reader_Unregister_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Reader_Unregister_Call) RunAndReturn(run func(context.Context) error) *Reader_Unregister_Call {
	_c.Call.Return(run)
	return _c
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
