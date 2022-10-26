// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/photon-storage/photon-proto/consensus (interfaces: Node_StreamChainHeadServer,Node_StreamAttestationsServer,Node_StreamBlocksServer,Node_StreamValidatorsInfoServer,Node_StreamIndexedAttestationsServer,Node_WaitForActivationServer,Node_WaitForChainStartServer,Node_StreamDutiesServer)

// Package mockgen is a generated GoMock package.
package mockgen

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pbc "github.com/photon-storage/photon-proto/consensus"
	metadata "google.golang.org/grpc/metadata"
)

// MockNode_StreamChainHeadServer is a mock of Node_StreamChainHeadServer interface.
type MockNode_StreamChainHeadServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_StreamChainHeadServerMockRecorder
}

// MockNode_StreamChainHeadServerMockRecorder is the mock recorder for MockNode_StreamChainHeadServer.
type MockNode_StreamChainHeadServerMockRecorder struct {
	mock *MockNode_StreamChainHeadServer
}

// NewMockNode_StreamChainHeadServer creates a new mock instance.
func NewMockNode_StreamChainHeadServer(ctrl *gomock.Controller) *MockNode_StreamChainHeadServer {
	mock := &MockNode_StreamChainHeadServer{ctrl: ctrl}
	mock.recorder = &MockNode_StreamChainHeadServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_StreamChainHeadServer) EXPECT() *MockNode_StreamChainHeadServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_StreamChainHeadServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_StreamChainHeadServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_StreamChainHeadServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNode_StreamChainHeadServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_StreamChainHeadServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_StreamChainHeadServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_StreamChainHeadServer) Send(arg0 *pbc.ChainHead) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_StreamChainHeadServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_StreamChainHeadServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_StreamChainHeadServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_StreamChainHeadServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_StreamChainHeadServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_StreamChainHeadServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_StreamChainHeadServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_StreamChainHeadServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_StreamChainHeadServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_StreamChainHeadServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_StreamChainHeadServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_StreamChainHeadServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_StreamChainHeadServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_StreamChainHeadServer)(nil).SetTrailer), arg0)
}

// MockNode_StreamAttestationsServer is a mock of Node_StreamAttestationsServer interface.
type MockNode_StreamAttestationsServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_StreamAttestationsServerMockRecorder
}

// MockNode_StreamAttestationsServerMockRecorder is the mock recorder for MockNode_StreamAttestationsServer.
type MockNode_StreamAttestationsServerMockRecorder struct {
	mock *MockNode_StreamAttestationsServer
}

// NewMockNode_StreamAttestationsServer creates a new mock instance.
func NewMockNode_StreamAttestationsServer(ctrl *gomock.Controller) *MockNode_StreamAttestationsServer {
	mock := &MockNode_StreamAttestationsServer{ctrl: ctrl}
	mock.recorder = &MockNode_StreamAttestationsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_StreamAttestationsServer) EXPECT() *MockNode_StreamAttestationsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_StreamAttestationsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_StreamAttestationsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_StreamAttestationsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNode_StreamAttestationsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_StreamAttestationsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_StreamAttestationsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_StreamAttestationsServer) Send(arg0 *pbc.Attestation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_StreamAttestationsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_StreamAttestationsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_StreamAttestationsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_StreamAttestationsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_StreamAttestationsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_StreamAttestationsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_StreamAttestationsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_StreamAttestationsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_StreamAttestationsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_StreamAttestationsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_StreamAttestationsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_StreamAttestationsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_StreamAttestationsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_StreamAttestationsServer)(nil).SetTrailer), arg0)
}

// MockNode_StreamBlocksServer is a mock of Node_StreamBlocksServer interface.
type MockNode_StreamBlocksServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_StreamBlocksServerMockRecorder
}

// MockNode_StreamBlocksServerMockRecorder is the mock recorder for MockNode_StreamBlocksServer.
type MockNode_StreamBlocksServerMockRecorder struct {
	mock *MockNode_StreamBlocksServer
}

// NewMockNode_StreamBlocksServer creates a new mock instance.
func NewMockNode_StreamBlocksServer(ctrl *gomock.Controller) *MockNode_StreamBlocksServer {
	mock := &MockNode_StreamBlocksServer{ctrl: ctrl}
	mock.recorder = &MockNode_StreamBlocksServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_StreamBlocksServer) EXPECT() *MockNode_StreamBlocksServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_StreamBlocksServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_StreamBlocksServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_StreamBlocksServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNode_StreamBlocksServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_StreamBlocksServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_StreamBlocksServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_StreamBlocksServer) Send(arg0 *pbc.SignedBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_StreamBlocksServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_StreamBlocksServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_StreamBlocksServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_StreamBlocksServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_StreamBlocksServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_StreamBlocksServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_StreamBlocksServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_StreamBlocksServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_StreamBlocksServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_StreamBlocksServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_StreamBlocksServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_StreamBlocksServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_StreamBlocksServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_StreamBlocksServer)(nil).SetTrailer), arg0)
}

// MockNode_StreamValidatorsInfoServer is a mock of Node_StreamValidatorsInfoServer interface.
type MockNode_StreamValidatorsInfoServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_StreamValidatorsInfoServerMockRecorder
}

// MockNode_StreamValidatorsInfoServerMockRecorder is the mock recorder for MockNode_StreamValidatorsInfoServer.
type MockNode_StreamValidatorsInfoServerMockRecorder struct {
	mock *MockNode_StreamValidatorsInfoServer
}

// NewMockNode_StreamValidatorsInfoServer creates a new mock instance.
func NewMockNode_StreamValidatorsInfoServer(ctrl *gomock.Controller) *MockNode_StreamValidatorsInfoServer {
	mock := &MockNode_StreamValidatorsInfoServer{ctrl: ctrl}
	mock.recorder = &MockNode_StreamValidatorsInfoServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_StreamValidatorsInfoServer) EXPECT() *MockNode_StreamValidatorsInfoServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) Recv() (*pbc.ValidatorChangeSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*pbc.ValidatorChangeSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) Send(arg0 *pbc.ValidatorInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_StreamValidatorsInfoServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_StreamValidatorsInfoServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_StreamValidatorsInfoServer)(nil).SetTrailer), arg0)
}

// MockNode_StreamIndexedAttestationsServer is a mock of Node_StreamIndexedAttestationsServer interface.
type MockNode_StreamIndexedAttestationsServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_StreamIndexedAttestationsServerMockRecorder
}

// MockNode_StreamIndexedAttestationsServerMockRecorder is the mock recorder for MockNode_StreamIndexedAttestationsServer.
type MockNode_StreamIndexedAttestationsServerMockRecorder struct {
	mock *MockNode_StreamIndexedAttestationsServer
}

// NewMockNode_StreamIndexedAttestationsServer creates a new mock instance.
func NewMockNode_StreamIndexedAttestationsServer(ctrl *gomock.Controller) *MockNode_StreamIndexedAttestationsServer {
	mock := &MockNode_StreamIndexedAttestationsServer{ctrl: ctrl}
	mock.recorder = &MockNode_StreamIndexedAttestationsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_StreamIndexedAttestationsServer) EXPECT() *MockNode_StreamIndexedAttestationsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_StreamIndexedAttestationsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_StreamIndexedAttestationsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_StreamIndexedAttestationsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNode_StreamIndexedAttestationsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_StreamIndexedAttestationsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_StreamIndexedAttestationsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_StreamIndexedAttestationsServer) Send(arg0 *pbc.IndexedAttestation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_StreamIndexedAttestationsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_StreamIndexedAttestationsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_StreamIndexedAttestationsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_StreamIndexedAttestationsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_StreamIndexedAttestationsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_StreamIndexedAttestationsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_StreamIndexedAttestationsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_StreamIndexedAttestationsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_StreamIndexedAttestationsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_StreamIndexedAttestationsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_StreamIndexedAttestationsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_StreamIndexedAttestationsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_StreamIndexedAttestationsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_StreamIndexedAttestationsServer)(nil).SetTrailer), arg0)
}

// MockNode_WaitForActivationServer is a mock of Node_WaitForActivationServer interface.
type MockNode_WaitForActivationServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_WaitForActivationServerMockRecorder
}

// MockNode_WaitForActivationServerMockRecorder is the mock recorder for MockNode_WaitForActivationServer.
type MockNode_WaitForActivationServerMockRecorder struct {
	mock *MockNode_WaitForActivationServer
}

// NewMockNode_WaitForActivationServer creates a new mock instance.
func NewMockNode_WaitForActivationServer(ctrl *gomock.Controller) *MockNode_WaitForActivationServer {
	mock := &MockNode_WaitForActivationServer{ctrl: ctrl}
	mock.recorder = &MockNode_WaitForActivationServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_WaitForActivationServer) EXPECT() *MockNode_WaitForActivationServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_WaitForActivationServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_WaitForActivationServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_WaitForActivationServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNode_WaitForActivationServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_WaitForActivationServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_WaitForActivationServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_WaitForActivationServer) Send(arg0 *pbc.ValidatorActivationResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_WaitForActivationServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_WaitForActivationServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_WaitForActivationServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_WaitForActivationServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_WaitForActivationServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_WaitForActivationServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_WaitForActivationServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_WaitForActivationServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_WaitForActivationServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_WaitForActivationServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_WaitForActivationServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_WaitForActivationServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_WaitForActivationServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_WaitForActivationServer)(nil).SetTrailer), arg0)
}

// MockNode_WaitForChainStartServer is a mock of Node_WaitForChainStartServer interface.
type MockNode_WaitForChainStartServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_WaitForChainStartServerMockRecorder
}

// MockNode_WaitForChainStartServerMockRecorder is the mock recorder for MockNode_WaitForChainStartServer.
type MockNode_WaitForChainStartServerMockRecorder struct {
	mock *MockNode_WaitForChainStartServer
}

// NewMockNode_WaitForChainStartServer creates a new mock instance.
func NewMockNode_WaitForChainStartServer(ctrl *gomock.Controller) *MockNode_WaitForChainStartServer {
	mock := &MockNode_WaitForChainStartServer{ctrl: ctrl}
	mock.recorder = &MockNode_WaitForChainStartServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_WaitForChainStartServer) EXPECT() *MockNode_WaitForChainStartServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_WaitForChainStartServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_WaitForChainStartServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_WaitForChainStartServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNode_WaitForChainStartServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_WaitForChainStartServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_WaitForChainStartServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_WaitForChainStartServer) Send(arg0 *pbc.ChainStartResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_WaitForChainStartServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_WaitForChainStartServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_WaitForChainStartServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_WaitForChainStartServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_WaitForChainStartServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_WaitForChainStartServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_WaitForChainStartServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_WaitForChainStartServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_WaitForChainStartServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_WaitForChainStartServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_WaitForChainStartServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_WaitForChainStartServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_WaitForChainStartServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_WaitForChainStartServer)(nil).SetTrailer), arg0)
}

// MockNode_StreamDutiesServer is a mock of Node_StreamDutiesServer interface.
type MockNode_StreamDutiesServer struct {
	ctrl     *gomock.Controller
	recorder *MockNode_StreamDutiesServerMockRecorder
}

// MockNode_StreamDutiesServerMockRecorder is the mock recorder for MockNode_StreamDutiesServer.
type MockNode_StreamDutiesServerMockRecorder struct {
	mock *MockNode_StreamDutiesServer
}

// NewMockNode_StreamDutiesServer creates a new mock instance.
func NewMockNode_StreamDutiesServer(ctrl *gomock.Controller) *MockNode_StreamDutiesServer {
	mock := &MockNode_StreamDutiesServer{ctrl: ctrl}
	mock.recorder = &MockNode_StreamDutiesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode_StreamDutiesServer) EXPECT() *MockNode_StreamDutiesServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNode_StreamDutiesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNode_StreamDutiesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNode_StreamDutiesServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNode_StreamDutiesServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNode_StreamDutiesServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNode_StreamDutiesServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNode_StreamDutiesServer) Send(arg0 *pbc.DutiesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNode_StreamDutiesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNode_StreamDutiesServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNode_StreamDutiesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNode_StreamDutiesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNode_StreamDutiesServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNode_StreamDutiesServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNode_StreamDutiesServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNode_StreamDutiesServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNode_StreamDutiesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNode_StreamDutiesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNode_StreamDutiesServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNode_StreamDutiesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNode_StreamDutiesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNode_StreamDutiesServer)(nil).SetTrailer), arg0)
}
