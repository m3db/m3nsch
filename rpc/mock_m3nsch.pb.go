// Automatically generated by MockGen. DO NOT EDIT!
// Source: m3nsch.pb.go

package rpc

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Mock of MenschClient interface
type MockMenschClient struct {
	ctrl     *gomock.Controller
	recorder *_MockMenschClientRecorder
}

// Recorder for MockMenschClient (not exported)
type _MockMenschClientRecorder struct {
	mock *MockMenschClient
}

func NewMockMenschClient(ctrl *gomock.Controller) *MockMenschClient {
	mock := &MockMenschClient{ctrl: ctrl}
	mock.recorder = &_MockMenschClientRecorder{mock}
	return mock
}

func (_m *MockMenschClient) EXPECT() *_MockMenschClientRecorder {
	return _m.recorder
}

func (_m *MockMenschClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Status", _s...)
	ret0, _ := ret[0].(*StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschClientRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Status", _s...)
}

func (_m *MockMenschClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Init", _s...)
	ret0, _ := ret[0].(*InitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschClientRecorder) Init(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", _s...)
}

func (_m *MockMenschClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Start", _s...)
	ret0, _ := ret[0].(*StartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschClientRecorder) Start(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", _s...)
}

func (_m *MockMenschClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Stop", _s...)
	ret0, _ := ret[0].(*StopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschClientRecorder) Stop(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop", _s...)
}

func (_m *MockMenschClient) Modify(ctx context.Context, in *ModifyRequest, opts ...grpc.CallOption) (*ModifyResponse, error) {
	_s := []interface{}{ctx, in}
	for _, _x := range opts {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Modify", _s...)
	ret0, _ := ret[0].(*ModifyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschClientRecorder) Modify(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Modify", _s...)
}

// Mock of MenschServer interface
type MockMenschServer struct {
	ctrl     *gomock.Controller
	recorder *_MockMenschServerRecorder
}

// Recorder for MockMenschServer (not exported)
type _MockMenschServerRecorder struct {
	mock *MockMenschServer
}

func NewMockMenschServer(ctrl *gomock.Controller) *MockMenschServer {
	mock := &MockMenschServer{ctrl: ctrl}
	mock.recorder = &_MockMenschServerRecorder{mock}
	return mock
}

func (_m *MockMenschServer) EXPECT() *_MockMenschServerRecorder {
	return _m.recorder
}

func (_m *MockMenschServer) Status(_param0 context.Context, _param1 *StatusRequest) (*StatusResponse, error) {
	ret := _m.ctrl.Call(_m, "Status", _param0, _param1)
	ret0, _ := ret[0].(*StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschServerRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Status", arg0, arg1)
}

func (_m *MockMenschServer) Init(_param0 context.Context, _param1 *InitRequest) (*InitResponse, error) {
	ret := _m.ctrl.Call(_m, "Init", _param0, _param1)
	ret0, _ := ret[0].(*InitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschServerRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0, arg1)
}

func (_m *MockMenschServer) Start(_param0 context.Context, _param1 *StartRequest) (*StartResponse, error) {
	ret := _m.ctrl.Call(_m, "Start", _param0, _param1)
	ret0, _ := ret[0].(*StartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschServerRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0, arg1)
}

func (_m *MockMenschServer) Stop(_param0 context.Context, _param1 *StopRequest) (*StopResponse, error) {
	ret := _m.ctrl.Call(_m, "Stop", _param0, _param1)
	ret0, _ := ret[0].(*StopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschServerRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop", arg0, arg1)
}

func (_m *MockMenschServer) Modify(_param0 context.Context, _param1 *ModifyRequest) (*ModifyResponse, error) {
	ret := _m.ctrl.Call(_m, "Modify", _param0, _param1)
	ret0, _ := ret[0].(*ModifyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockMenschServerRecorder) Modify(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Modify", arg0, arg1)
}