// Code generated by MockGen. DO NOT EDIT.
// Source: ./statefactory/statefactory.go

// Package mock_statefactory is a generated GoMock package.
package mock_statefactory

import (
	gomock "github.com/golang/mock/gomock"
	trx "github.com/iotexproject/iotex-core/blockchain/trx"
	common "github.com/iotexproject/iotex-core/common"
	statefactory "github.com/iotexproject/iotex-core/statefactory"
	big "math/big"
	reflect "reflect"
)

// MockStateFactory is a mock of StateFactory interface
type MockStateFactory struct {
	ctrl     *gomock.Controller
	recorder *MockStateFactoryMockRecorder
}

// MockStateFactoryMockRecorder is the mock recorder for MockStateFactory
type MockStateFactoryMockRecorder struct {
	mock *MockStateFactory
}

// NewMockStateFactory creates a new mock instance
func NewMockStateFactory(ctrl *gomock.Controller) *MockStateFactory {
	mock := &MockStateFactory{ctrl: ctrl}
	mock.recorder = &MockStateFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateFactory) EXPECT() *MockStateFactoryMockRecorder {
	return m.recorder
}

// CreateState mocks base method
func (m *MockStateFactory) CreateState(arg0 string, arg1 uint64) (*statefactory.State, error) {
	ret := m.ctrl.Call(m, "CreateState", arg0, arg1)
	ret0, _ := ret[0].(*statefactory.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateState indicates an expected call of CreateState
func (mr *MockStateFactoryMockRecorder) CreateState(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateState", reflect.TypeOf((*MockStateFactory)(nil).CreateState), arg0, arg1)
}

// Balance mocks base method
func (m *MockStateFactory) Balance(arg0 string) (*big.Int, error) {
	ret := m.ctrl.Call(m, "Balance", arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balance indicates an expected call of Balance
func (mr *MockStateFactoryMockRecorder) Balance(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockStateFactory)(nil).Balance), arg0)
}

// UpdateStatesWithTransfer mocks base method
func (m *MockStateFactory) UpdateStatesWithTransfer(arg0 []*trx.Tx) error {
	ret := m.ctrl.Call(m, "UpdateStatesWithTransfer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatesWithTransfer indicates an expected call of UpdateStatesWithTransfer
func (mr *MockStateFactoryMockRecorder) UpdateStatesWithTransfer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatesWithTransfer", reflect.TypeOf((*MockStateFactory)(nil).UpdateStatesWithTransfer), arg0)
}

// SetNonce mocks base method
func (m *MockStateFactory) SetNonce(arg0 string, arg1 uint64) error {
	ret := m.ctrl.Call(m, "SetNonce", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNonce indicates an expected call of SetNonce
func (mr *MockStateFactoryMockRecorder) SetNonce(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNonce", reflect.TypeOf((*MockStateFactory)(nil).SetNonce), arg0, arg1)
}

// Nonce mocks base method
func (m *MockStateFactory) Nonce(arg0 string) (uint64, error) {
	ret := m.ctrl.Call(m, "Nonce", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Nonce indicates an expected call of Nonce
func (mr *MockStateFactoryMockRecorder) Nonce(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockStateFactory)(nil).Nonce), arg0)
}

// RootHash mocks base method
func (m *MockStateFactory) RootHash() common.Hash32B {
	ret := m.ctrl.Call(m, "RootHash")
	ret0, _ := ret[0].(common.Hash32B)
	return ret0
}

// RootHash indicates an expected call of RootHash
func (mr *MockStateFactoryMockRecorder) RootHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHash", reflect.TypeOf((*MockStateFactory)(nil).RootHash))
}