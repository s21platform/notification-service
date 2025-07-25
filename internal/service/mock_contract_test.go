// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/s21platform/notification-service/internal/model"
)

// MockDbRepo is a mock of DbRepo interface.
type MockDbRepo struct {
	ctrl     *gomock.Controller
	recorder *MockDbRepoMockRecorder
}

// MockDbRepoMockRecorder is the mock recorder for MockDbRepo.
type MockDbRepoMockRecorder struct {
	mock *MockDbRepo
}

// NewMockDbRepo creates a new mock instance.
func NewMockDbRepo(ctrl *gomock.Controller) *MockDbRepo {
	mock := &MockDbRepo{ctrl: ctrl}
	mock.recorder = &MockDbRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDbRepo) EXPECT() *MockDbRepoMockRecorder {
	return m.recorder
}

// GetCountNotification mocks base method.
func (m *MockDbRepo) GetCountNotification(ctx context.Context, userUuid string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountNotification", ctx, userUuid)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountNotification indicates an expected call of GetCountNotification.
func (mr *MockDbRepoMockRecorder) GetCountNotification(ctx, userUuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountNotification", reflect.TypeOf((*MockDbRepo)(nil).GetCountNotification), ctx, userUuid)
}

// GetNotifications mocks base method.
func (m *MockDbRepo) GetNotifications(ctx context.Context, userUuid string, limit, offset int64) ([]model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifications", ctx, userUuid, limit, offset)
	ret0, _ := ret[0].([]model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockDbRepoMockRecorder) GetNotifications(ctx, userUuid, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockDbRepo)(nil).GetNotifications), ctx, userUuid, limit, offset)
}

// MarkNotificationsAsRead mocks base method.
func (m *MockDbRepo) MarkNotificationsAsRead(ctx context.Context, userUuid string, notificationId []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkNotificationsAsRead", ctx, userUuid, notificationId)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkNotificationsAsRead indicates an expected call of MarkNotificationsAsRead.
func (mr *MockDbRepoMockRecorder) MarkNotificationsAsRead(ctx, userUuid, notificationId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkNotificationsAsRead", reflect.TypeOf((*MockDbRepo)(nil).MarkNotificationsAsRead), ctx, userUuid, notificationId)
}

// MockEmailSender is a mock of EmailSender interface.
type MockEmailSender struct {
	ctrl     *gomock.Controller
	recorder *MockEmailSenderMockRecorder
}

// MockEmailSenderMockRecorder is the mock recorder for MockEmailSender.
type MockEmailSenderMockRecorder struct {
	mock *MockEmailSender
}

// NewMockEmailSender creates a new mock instance.
func NewMockEmailSender(ctrl *gomock.Controller) *MockEmailSender {
	mock := &MockEmailSender{ctrl: ctrl}
	mock.recorder = &MockEmailSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailSender) EXPECT() *MockEmailSenderMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockEmailSender) SendEmail(subject, to, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", subject, to, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockEmailSenderMockRecorder) SendEmail(subject, to, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockEmailSender)(nil).SendEmail), subject, to, content)
}

// MockVerificationCodeSender is a mock of VerificationCodeSender interface.
type MockVerificationCodeSender struct {
	ctrl     *gomock.Controller
	recorder *MockVerificationCodeSenderMockRecorder
}

// MockVerificationCodeSenderMockRecorder is the mock recorder for MockVerificationCodeSender.
type MockVerificationCodeSenderMockRecorder struct {
	mock *MockVerificationCodeSender
}

// NewMockVerificationCodeSender creates a new mock instance.
func NewMockVerificationCodeSender(ctrl *gomock.Controller) *MockVerificationCodeSender {
	mock := &MockVerificationCodeSender{ctrl: ctrl}
	mock.recorder = &MockVerificationCodeSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerificationCodeSender) EXPECT() *MockVerificationCodeSenderMockRecorder {
	return m.recorder
}

// SendVerificationCode mocks base method.
func (m *MockVerificationCodeSender) SendVerificationCode(email, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerificationCode", email, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerificationCode indicates an expected call of SendVerificationCode.
func (mr *MockVerificationCodeSenderMockRecorder) SendVerificationCode(email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationCode", reflect.TypeOf((*MockVerificationCodeSender)(nil).SendVerificationCode), email, code)
}

// MockVerificationEduCodeSender is a mock of VerificationEduCodeSender interface.
type MockVerificationEduCodeSender struct {
	ctrl     *gomock.Controller
	recorder *MockVerificationEduCodeSenderMockRecorder
}

// MockVerificationEduCodeSenderMockRecorder is the mock recorder for MockVerificationEduCodeSender.
type MockVerificationEduCodeSenderMockRecorder struct {
	mock *MockVerificationEduCodeSender
}

// NewMockVerificationEduCodeSender creates a new mock instance.
func NewMockVerificationEduCodeSender(ctrl *gomock.Controller) *MockVerificationEduCodeSender {
	mock := &MockVerificationEduCodeSender{ctrl: ctrl}
	mock.recorder = &MockVerificationEduCodeSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerificationEduCodeSender) EXPECT() *MockVerificationEduCodeSenderMockRecorder {
	return m.recorder
}

// SendVerificationCode mocks base method.
func (m *MockVerificationEduCodeSender) SendVerificationCode(email, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerificationCode", email, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerificationCode indicates an expected call of SendVerificationCode.
func (mr *MockVerificationEduCodeSenderMockRecorder) SendVerificationCode(email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationCode", reflect.TypeOf((*MockVerificationEduCodeSender)(nil).SendVerificationCode), email, code)
}
