// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/protobuf/comment/v1/comment_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/jxlwqq/blog-microservices/api/protobuf/comment/v1"
	grpc "google.golang.org/grpc"
)

// MockCommentServiceClient is a mock of CommentServiceClient interface.
type MockCommentServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCommentServiceClientMockRecorder
}

// MockCommentServiceClientMockRecorder is the mock recorder for MockCommentServiceClient.
type MockCommentServiceClientMockRecorder struct {
	mock *MockCommentServiceClient
}

// NewMockCommentServiceClient creates a new mock instance.
func NewMockCommentServiceClient(ctrl *gomock.Controller) *MockCommentServiceClient {
	mock := &MockCommentServiceClient{ctrl: ctrl}
	mock.recorder = &MockCommentServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentServiceClient) EXPECT() *MockCommentServiceClientMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockCommentServiceClient) CreateComment(ctx context.Context, in *v1.CreateCommentRequest, opts ...grpc.CallOption) (*v1.CreateCommentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateComment", varargs...)
	ret0, _ := ret[0].(*v1.CreateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentServiceClientMockRecorder) CreateComment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentServiceClient)(nil).CreateComment), varargs...)
}

// CreateCommentCompensate mocks base method.
func (m *MockCommentServiceClient) CreateCommentCompensate(ctx context.Context, in *v1.CreateCommentRequest, opts ...grpc.CallOption) (*v1.CreateCommentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCommentCompensate", varargs...)
	ret0, _ := ret[0].(*v1.CreateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommentCompensate indicates an expected call of CreateCommentCompensate.
func (mr *MockCommentServiceClientMockRecorder) CreateCommentCompensate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommentCompensate", reflect.TypeOf((*MockCommentServiceClient)(nil).CreateCommentCompensate), varargs...)
}

// DeleteComment mocks base method.
func (m *MockCommentServiceClient) DeleteComment(ctx context.Context, in *v1.DeleteCommentRequest, opts ...grpc.CallOption) (*v1.DeleteCommentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteComment", varargs...)
	ret0, _ := ret[0].(*v1.DeleteCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockCommentServiceClientMockRecorder) DeleteComment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCommentServiceClient)(nil).DeleteComment), varargs...)
}

// DeleteCommentCompensate mocks base method.
func (m *MockCommentServiceClient) DeleteCommentCompensate(ctx context.Context, in *v1.DeleteCommentRequest, opts ...grpc.CallOption) (*v1.DeleteCommentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCommentCompensate", varargs...)
	ret0, _ := ret[0].(*v1.DeleteCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCommentCompensate indicates an expected call of DeleteCommentCompensate.
func (mr *MockCommentServiceClientMockRecorder) DeleteCommentCompensate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentCompensate", reflect.TypeOf((*MockCommentServiceClient)(nil).DeleteCommentCompensate), varargs...)
}

// GetComment mocks base method.
func (m *MockCommentServiceClient) GetComment(ctx context.Context, in *v1.GetCommentRequest, opts ...grpc.CallOption) (*v1.GetCommentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetComment", varargs...)
	ret0, _ := ret[0].(*v1.GetCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockCommentServiceClientMockRecorder) GetComment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockCommentServiceClient)(nil).GetComment), varargs...)
}

// GetCommentByUUID mocks base method.
func (m *MockCommentServiceClient) GetCommentByUUID(ctx context.Context, in *v1.GetCommentByUUIDRequest, opts ...grpc.CallOption) (*v1.GetCommentByUUIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommentByUUID", varargs...)
	ret0, _ := ret[0].(*v1.GetCommentByUUIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentByUUID indicates an expected call of GetCommentByUUID.
func (mr *MockCommentServiceClientMockRecorder) GetCommentByUUID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByUUID", reflect.TypeOf((*MockCommentServiceClient)(nil).GetCommentByUUID), varargs...)
}

// ListCommentsByPostID mocks base method.
func (m *MockCommentServiceClient) ListCommentsByPostID(ctx context.Context, in *v1.ListCommentsByPostIDRequest, opts ...grpc.CallOption) (*v1.ListCommentsByPostIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCommentsByPostID", varargs...)
	ret0, _ := ret[0].(*v1.ListCommentsByPostIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommentsByPostID indicates an expected call of ListCommentsByPostID.
func (mr *MockCommentServiceClientMockRecorder) ListCommentsByPostID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommentsByPostID", reflect.TypeOf((*MockCommentServiceClient)(nil).ListCommentsByPostID), varargs...)
}

// UpdateComment mocks base method.
func (m *MockCommentServiceClient) UpdateComment(ctx context.Context, in *v1.UpdateCommentRequest, opts ...grpc.CallOption) (*v1.UpdateCommentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateComment", varargs...)
	ret0, _ := ret[0].(*v1.UpdateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockCommentServiceClientMockRecorder) UpdateComment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockCommentServiceClient)(nil).UpdateComment), varargs...)
}

// MockCommentServiceServer is a mock of CommentServiceServer interface.
type MockCommentServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockCommentServiceServerMockRecorder
}

// MockCommentServiceServerMockRecorder is the mock recorder for MockCommentServiceServer.
type MockCommentServiceServerMockRecorder struct {
	mock *MockCommentServiceServer
}

// NewMockCommentServiceServer creates a new mock instance.
func NewMockCommentServiceServer(ctrl *gomock.Controller) *MockCommentServiceServer {
	mock := &MockCommentServiceServer{ctrl: ctrl}
	mock.recorder = &MockCommentServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentServiceServer) EXPECT() *MockCommentServiceServerMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockCommentServiceServer) CreateComment(arg0 context.Context, arg1 *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(*v1.CreateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentServiceServerMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentServiceServer)(nil).CreateComment), arg0, arg1)
}

// CreateCommentCompensate mocks base method.
func (m *MockCommentServiceServer) CreateCommentCompensate(arg0 context.Context, arg1 *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommentCompensate", arg0, arg1)
	ret0, _ := ret[0].(*v1.CreateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCommentCompensate indicates an expected call of CreateCommentCompensate.
func (mr *MockCommentServiceServerMockRecorder) CreateCommentCompensate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommentCompensate", reflect.TypeOf((*MockCommentServiceServer)(nil).CreateCommentCompensate), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockCommentServiceServer) DeleteComment(arg0 context.Context, arg1 *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(*v1.DeleteCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockCommentServiceServerMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCommentServiceServer)(nil).DeleteComment), arg0, arg1)
}

// DeleteCommentCompensate mocks base method.
func (m *MockCommentServiceServer) DeleteCommentCompensate(arg0 context.Context, arg1 *v1.DeleteCommentRequest) (*v1.DeleteCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentCompensate", arg0, arg1)
	ret0, _ := ret[0].(*v1.DeleteCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCommentCompensate indicates an expected call of DeleteCommentCompensate.
func (mr *MockCommentServiceServerMockRecorder) DeleteCommentCompensate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentCompensate", reflect.TypeOf((*MockCommentServiceServer)(nil).DeleteCommentCompensate), arg0, arg1)
}

// GetComment mocks base method.
func (m *MockCommentServiceServer) GetComment(arg0 context.Context, arg1 *v1.GetCommentRequest) (*v1.GetCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockCommentServiceServerMockRecorder) GetComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockCommentServiceServer)(nil).GetComment), arg0, arg1)
}

// GetCommentByUUID mocks base method.
func (m *MockCommentServiceServer) GetCommentByUUID(arg0 context.Context, arg1 *v1.GetCommentByUUIDRequest) (*v1.GetCommentByUUIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentByUUID", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetCommentByUUIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentByUUID indicates an expected call of GetCommentByUUID.
func (mr *MockCommentServiceServerMockRecorder) GetCommentByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentByUUID", reflect.TypeOf((*MockCommentServiceServer)(nil).GetCommentByUUID), arg0, arg1)
}

// ListCommentsByPostID mocks base method.
func (m *MockCommentServiceServer) ListCommentsByPostID(arg0 context.Context, arg1 *v1.ListCommentsByPostIDRequest) (*v1.ListCommentsByPostIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommentsByPostID", arg0, arg1)
	ret0, _ := ret[0].(*v1.ListCommentsByPostIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommentsByPostID indicates an expected call of ListCommentsByPostID.
func (mr *MockCommentServiceServerMockRecorder) ListCommentsByPostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommentsByPostID", reflect.TypeOf((*MockCommentServiceServer)(nil).ListCommentsByPostID), arg0, arg1)
}

// UpdateComment mocks base method.
func (m *MockCommentServiceServer) UpdateComment(arg0 context.Context, arg1 *v1.UpdateCommentRequest) (*v1.UpdateCommentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(*v1.UpdateCommentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockCommentServiceServerMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockCommentServiceServer)(nil).UpdateComment), arg0, arg1)
}

// mustEmbedUnimplementedCommentServiceServer mocks base method.
func (m *MockCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCommentServiceServer")
}

// mustEmbedUnimplementedCommentServiceServer indicates an expected call of mustEmbedUnimplementedCommentServiceServer.
func (mr *MockCommentServiceServerMockRecorder) mustEmbedUnimplementedCommentServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCommentServiceServer", reflect.TypeOf((*MockCommentServiceServer)(nil).mustEmbedUnimplementedCommentServiceServer))
}

// MockUnsafeCommentServiceServer is a mock of UnsafeCommentServiceServer interface.
type MockUnsafeCommentServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeCommentServiceServerMockRecorder
}

// MockUnsafeCommentServiceServerMockRecorder is the mock recorder for MockUnsafeCommentServiceServer.
type MockUnsafeCommentServiceServerMockRecorder struct {
	mock *MockUnsafeCommentServiceServer
}

// NewMockUnsafeCommentServiceServer creates a new mock instance.
func NewMockUnsafeCommentServiceServer(ctrl *gomock.Controller) *MockUnsafeCommentServiceServer {
	mock := &MockUnsafeCommentServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeCommentServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeCommentServiceServer) EXPECT() *MockUnsafeCommentServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedCommentServiceServer mocks base method.
func (m *MockUnsafeCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedCommentServiceServer")
}

// mustEmbedUnimplementedCommentServiceServer indicates an expected call of mustEmbedUnimplementedCommentServiceServer.
func (mr *MockUnsafeCommentServiceServerMockRecorder) mustEmbedUnimplementedCommentServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedCommentServiceServer", reflect.TypeOf((*MockUnsafeCommentServiceServer)(nil).mustEmbedUnimplementedCommentServiceServer))
}
