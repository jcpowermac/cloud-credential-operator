// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	iam "github.com/aws/aws-sdk-go/service/iam"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateAccessKey mocks base method
func (m *MockClient) CreateAccessKey(arg0 *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessKey", arg0)
	ret0, _ := ret[0].(*iam.CreateAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessKey indicates an expected call of CreateAccessKey
func (mr *MockClientMockRecorder) CreateAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessKey", reflect.TypeOf((*MockClient)(nil).CreateAccessKey), arg0)
}

// CreateUser mocks base method
func (m *MockClient) CreateUser(arg0 *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*iam.CreateUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockClientMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockClient)(nil).CreateUser), arg0)
}

// DeleteAccessKey mocks base method
func (m *MockClient) DeleteAccessKey(arg0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessKey", arg0)
	ret0, _ := ret[0].(*iam.DeleteAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccessKey indicates an expected call of DeleteAccessKey
func (mr *MockClientMockRecorder) DeleteAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessKey", reflect.TypeOf((*MockClient)(nil).DeleteAccessKey), arg0)
}

// DeleteUser mocks base method
func (m *MockClient) DeleteUser(arg0 *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*iam.DeleteUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockClientMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockClient)(nil).DeleteUser), arg0)
}

// DeleteUserPolicy mocks base method
func (m *MockClient) DeleteUserPolicy(arg0 *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.DeleteUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserPolicy indicates an expected call of DeleteUserPolicy
func (mr *MockClientMockRecorder) DeleteUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPolicy", reflect.TypeOf((*MockClient)(nil).DeleteUserPolicy), arg0)
}

// GetUser mocks base method
func (m *MockClient) GetUser(arg0 *iam.GetUserInput) (*iam.GetUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*iam.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockClientMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockClient)(nil).GetUser), arg0)
}

// ListAccessKeys mocks base method
func (m *MockClient) ListAccessKeys(arg0 *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessKeys", arg0)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessKeys indicates an expected call of ListAccessKeys
func (mr *MockClientMockRecorder) ListAccessKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessKeys", reflect.TypeOf((*MockClient)(nil).ListAccessKeys), arg0)
}

// ListUserPolicies mocks base method
func (m *MockClient) ListUserPolicies(arg0 *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserPolicies", arg0)
	ret0, _ := ret[0].(*iam.ListUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserPolicies indicates an expected call of ListUserPolicies
func (mr *MockClientMockRecorder) ListUserPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPolicies", reflect.TypeOf((*MockClient)(nil).ListUserPolicies), arg0)
}

// PutUserPolicy mocks base method
func (m *MockClient) PutUserPolicy(arg0 *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.PutUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutUserPolicy indicates an expected call of PutUserPolicy
func (mr *MockClientMockRecorder) PutUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUserPolicy", reflect.TypeOf((*MockClient)(nil).PutUserPolicy), arg0)
}

// GetUserPolicy mocks base method
func (m *MockClient) GetUserPolicy(arg0 *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.GetUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserPolicy indicates an expected call of GetUserPolicy
func (mr *MockClientMockRecorder) GetUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPolicy", reflect.TypeOf((*MockClient)(nil).GetUserPolicy), arg0)
}

// SimulatePrincipalPolicy mocks base method
func (m *MockClient) SimulatePrincipalPolicy(arg0 *iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulatePrincipalPolicy", arg0)
	ret0, _ := ret[0].(*iam.SimulatePolicyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SimulatePrincipalPolicy indicates an expected call of SimulatePrincipalPolicy
func (mr *MockClientMockRecorder) SimulatePrincipalPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulatePrincipalPolicy", reflect.TypeOf((*MockClient)(nil).SimulatePrincipalPolicy), arg0)
}

// TagUser mocks base method
func (m *MockClient) TagUser(arg0 *iam.TagUserInput) (*iam.TagUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagUser", arg0)
	ret0, _ := ret[0].(*iam.TagUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagUser indicates an expected call of TagUser
func (mr *MockClientMockRecorder) TagUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagUser", reflect.TypeOf((*MockClient)(nil).TagUser), arg0)
}