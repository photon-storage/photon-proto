// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/photon-storage/photon-proto/keymanager (interfaces: KeyManagerClient)

// Package mockgen is a generated GoMock package.
package mockgen

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pbk "github.com/photon-storage/photon-proto/keymanager"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockKeyManagerClient is a mock of KeyManagerClient interface.
type MockKeyManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyManagerClientMockRecorder
}

// MockKeyManagerClientMockRecorder is the mock recorder for MockKeyManagerClient.
type MockKeyManagerClientMockRecorder struct {
	mock *MockKeyManagerClient
}

// NewMockKeyManagerClient creates a new mock instance.
func NewMockKeyManagerClient(ctrl *gomock.Controller) *MockKeyManagerClient {
	mock := &MockKeyManagerClient{ctrl: ctrl}
	mock.recorder = &MockKeyManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyManagerClient) EXPECT() *MockKeyManagerClientMockRecorder {
	return m.recorder
}

// ListPublicKeys mocks base method.
func (m *MockKeyManagerClient) ListPublicKeys(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*pbk.ListPublicKeysResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPublicKeys", varargs...)
	ret0, _ := ret[0].(*pbk.ListPublicKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublicKeys indicates an expected call of ListPublicKeys.
func (mr *MockKeyManagerClientMockRecorder) ListPublicKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublicKeys", reflect.TypeOf((*MockKeyManagerClient)(nil).ListPublicKeys), varargs...)
}

// Sign mocks base method.
func (m *MockKeyManagerClient) Sign(arg0 context.Context, arg1 *pbk.SignRequest, arg2 ...grpc.CallOption) (*pbk.SignResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sign", varargs...)
	ret0, _ := ret[0].(*pbk.SignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockKeyManagerClientMockRecorder) Sign(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockKeyManagerClient)(nil).Sign), varargs...)
}
