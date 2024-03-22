// Code generated by MockGen. DO NOT EDIT.
// Source: transition.go

// Package cluster is a generated GoMock package.
package cluster

import (
	reflect "reflect"

	stateswitch "github.com/filanov/stateswitch"
	gomock "github.com/golang/mock/gomock"
)

// MockTransitionHandler is a mock of TransitionHandler interface.
type MockTransitionHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTransitionHandlerMockRecorder
}

// MockTransitionHandlerMockRecorder is the mock recorder for MockTransitionHandler.
type MockTransitionHandlerMockRecorder struct {
	mock *MockTransitionHandler
}

// NewMockTransitionHandler creates a new mock instance.
func NewMockTransitionHandler(ctrl *gomock.Controller) *MockTransitionHandler {
	mock := &MockTransitionHandler{ctrl: ctrl}
	mock.recorder = &MockTransitionHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransitionHandler) EXPECT() *MockTransitionHandlerMockRecorder {
	return m.recorder
}

// FinalizingStageTimeoutMinutes mocks base method.
func (m *MockTransitionHandler) FinalizingStageTimeoutMinutes(sCluster *stateCluster) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizingStageTimeoutMinutes", sCluster)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// FinalizingStageTimeoutMinutes indicates an expected call of FinalizingStageTimeoutMinutes.
func (mr *MockTransitionHandlerMockRecorder) FinalizingStageTimeoutMinutes(sCluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizingStageTimeoutMinutes", reflect.TypeOf((*MockTransitionHandler)(nil).FinalizingStageTimeoutMinutes), sCluster)
}

// InstallCluster mocks base method.
func (m *MockTransitionHandler) InstallCluster(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallCluster", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallCluster indicates an expected call of InstallCluster.
func (mr *MockTransitionHandlerMockRecorder) InstallCluster(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallCluster", reflect.TypeOf((*MockTransitionHandler)(nil).InstallCluster), sw, args)
}

// InstallationTimeoutMinutes mocks base method.
func (m *MockTransitionHandler) InstallationTimeoutMinutes(arg0 *stateCluster) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallationTimeoutMinutes", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// InstallationTimeoutMinutes indicates an expected call of InstallationTimeoutMinutes.
func (mr *MockTransitionHandlerMockRecorder) InstallationTimeoutMinutes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallationTimeoutMinutes", reflect.TypeOf((*MockTransitionHandler)(nil).InstallationTimeoutMinutes), arg0)
}

// IsFinalizing mocks base method.
func (m *MockTransitionHandler) IsFinalizing(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFinalizing", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFinalizing indicates an expected call of IsFinalizing.
func (mr *MockTransitionHandlerMockRecorder) IsFinalizing(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFinalizing", reflect.TypeOf((*MockTransitionHandler)(nil).IsFinalizing), sw, args)
}

// IsFinalizingStageTimedOut mocks base method.
func (m *MockTransitionHandler) IsFinalizingStageTimedOut(sw stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFinalizingStageTimedOut", sw, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFinalizingStageTimedOut indicates an expected call of IsFinalizingStageTimedOut.
func (mr *MockTransitionHandlerMockRecorder) IsFinalizingStageTimedOut(sw, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFinalizingStageTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).IsFinalizingStageTimedOut), sw, arg1)
}

// IsFinalizingTimedOut mocks base method.
func (m *MockTransitionHandler) IsFinalizingTimedOut(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFinalizingTimedOut", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFinalizingTimedOut indicates an expected call of IsFinalizingTimedOut.
func (mr *MockTransitionHandlerMockRecorder) IsFinalizingTimedOut(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFinalizingTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).IsFinalizingTimedOut), sw, args)
}

// IsInstallationTimedOut mocks base method.
func (m *MockTransitionHandler) IsInstallationTimedOut(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstallationTimedOut", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInstallationTimedOut indicates an expected call of IsInstallationTimedOut.
func (mr *MockTransitionHandlerMockRecorder) IsInstallationTimedOut(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstallationTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).IsInstallationTimedOut), sw, args)
}

// IsInstalling mocks base method.
func (m *MockTransitionHandler) IsInstalling(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstalling", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInstalling indicates an expected call of IsInstalling.
func (mr *MockTransitionHandlerMockRecorder) IsInstalling(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstalling", reflect.TypeOf((*MockTransitionHandler)(nil).IsInstalling), sw, args)
}

// IsInstallingPendingUserAction mocks base method.
func (m *MockTransitionHandler) IsInstallingPendingUserAction(sw stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstallingPendingUserAction", sw, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInstallingPendingUserAction indicates an expected call of IsInstallingPendingUserAction.
func (mr *MockTransitionHandlerMockRecorder) IsInstallingPendingUserAction(sw, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstallingPendingUserAction", reflect.TypeOf((*MockTransitionHandler)(nil).IsInstallingPendingUserAction), sw, arg1)
}

// IsLogCollectionTimedOut mocks base method.
func (m *MockTransitionHandler) IsLogCollectionTimedOut(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLogCollectionTimedOut", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLogCollectionTimedOut indicates an expected call of IsLogCollectionTimedOut.
func (mr *MockTransitionHandlerMockRecorder) IsLogCollectionTimedOut(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLogCollectionTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).IsLogCollectionTimedOut), sw, args)
}

// IsPreparingTimedOut mocks base method.
func (m *MockTransitionHandler) IsPreparingTimedOut(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPreparingTimedOut", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPreparingTimedOut indicates an expected call of IsPreparingTimedOut.
func (mr *MockTransitionHandlerMockRecorder) IsPreparingTimedOut(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPreparingTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).IsPreparingTimedOut), sw, args)
}

// PostCancelInstallation mocks base method.
func (m *MockTransitionHandler) PostCancelInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCancelInstallation", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostCancelInstallation indicates an expected call of PostCancelInstallation.
func (mr *MockTransitionHandlerMockRecorder) PostCancelInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCancelInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).PostCancelInstallation), sw, args)
}

// PostCompleteInstallation mocks base method.
func (m *MockTransitionHandler) PostCompleteInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostCompleteInstallation", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostCompleteInstallation indicates an expected call of PostCompleteInstallation.
func (mr *MockTransitionHandlerMockRecorder) PostCompleteInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCompleteInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).PostCompleteInstallation), sw, args)
}

// PostHandlePreInstallationError mocks base method.
func (m *MockTransitionHandler) PostHandlePreInstallationError(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostHandlePreInstallationError", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostHandlePreInstallationError indicates an expected call of PostHandlePreInstallationError.
func (mr *MockTransitionHandlerMockRecorder) PostHandlePreInstallationError(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostHandlePreInstallationError", reflect.TypeOf((*MockTransitionHandler)(nil).PostHandlePreInstallationError), sw, args)
}

// PostPrepareForInstallation mocks base method.
func (m *MockTransitionHandler) PostPrepareForInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostPrepareForInstallation", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostPrepareForInstallation indicates an expected call of PostPrepareForInstallation.
func (mr *MockTransitionHandlerMockRecorder) PostPrepareForInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostPrepareForInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).PostPrepareForInstallation), sw, args)
}

// PostPreparingTimedOut mocks base method.
func (m *MockTransitionHandler) PostPreparingTimedOut(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostPreparingTimedOut", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostPreparingTimedOut indicates an expected call of PostPreparingTimedOut.
func (mr *MockTransitionHandlerMockRecorder) PostPreparingTimedOut(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostPreparingTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).PostPreparingTimedOut), sw, args)
}

// PostRefreshCluster mocks base method.
func (m *MockTransitionHandler) PostRefreshCluster(reason string, formatFuncs ...formatFunction) stateswitch.PostTransition {
	m.ctrl.T.Helper()
	varargs := []interface{}{reason}
	for _, a := range formatFuncs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostRefreshCluster", varargs...)
	ret0, _ := ret[0].(stateswitch.PostTransition)
	return ret0
}

// PostRefreshCluster indicates an expected call of PostRefreshCluster.
func (mr *MockTransitionHandlerMockRecorder) PostRefreshCluster(reason interface{}, formatFuncs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{reason}, formatFuncs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRefreshCluster", reflect.TypeOf((*MockTransitionHandler)(nil).PostRefreshCluster), varargs...)
}

// PostRefreshFinalizingStageSoftTimedOut mocks base method.
func (m *MockTransitionHandler) PostRefreshFinalizingStageSoftTimedOut(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRefreshFinalizingStageSoftTimedOut", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostRefreshFinalizingStageSoftTimedOut indicates an expected call of PostRefreshFinalizingStageSoftTimedOut.
func (mr *MockTransitionHandlerMockRecorder) PostRefreshFinalizingStageSoftTimedOut(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRefreshFinalizingStageSoftTimedOut", reflect.TypeOf((*MockTransitionHandler)(nil).PostRefreshFinalizingStageSoftTimedOut), sw, args)
}

// PostRefreshLogsProgress mocks base method.
func (m *MockTransitionHandler) PostRefreshLogsProgress(progress string) stateswitch.PostTransition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRefreshLogsProgress", progress)
	ret0, _ := ret[0].(stateswitch.PostTransition)
	return ret0
}

// PostRefreshLogsProgress indicates an expected call of PostRefreshLogsProgress.
func (mr *MockTransitionHandlerMockRecorder) PostRefreshLogsProgress(progress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRefreshLogsProgress", reflect.TypeOf((*MockTransitionHandler)(nil).PostRefreshLogsProgress), progress)
}

// PostResetCluster mocks base method.
func (m *MockTransitionHandler) PostResetCluster(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostResetCluster", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostResetCluster indicates an expected call of PostResetCluster.
func (mr *MockTransitionHandlerMockRecorder) PostResetCluster(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostResetCluster", reflect.TypeOf((*MockTransitionHandler)(nil).PostResetCluster), sw, args)
}

// PostUpdateFinalizingAMSConsoleUrl mocks base method.
func (m *MockTransitionHandler) PostUpdateFinalizingAMSConsoleUrl(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUpdateFinalizingAMSConsoleUrl", sw, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostUpdateFinalizingAMSConsoleUrl indicates an expected call of PostUpdateFinalizingAMSConsoleUrl.
func (mr *MockTransitionHandlerMockRecorder) PostUpdateFinalizingAMSConsoleUrl(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUpdateFinalizingAMSConsoleUrl", reflect.TypeOf((*MockTransitionHandler)(nil).PostUpdateFinalizingAMSConsoleUrl), sw, args)
}

// SoftTimeoutsEnabled mocks base method.
func (m *MockTransitionHandler) SoftTimeoutsEnabled(arg0 stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftTimeoutsEnabled", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SoftTimeoutsEnabled indicates an expected call of SoftTimeoutsEnabled.
func (mr *MockTransitionHandlerMockRecorder) SoftTimeoutsEnabled(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftTimeoutsEnabled", reflect.TypeOf((*MockTransitionHandler)(nil).SoftTimeoutsEnabled), arg0, arg1)
}

// WithAMSSubscriptions mocks base method.
func (m *MockTransitionHandler) WithAMSSubscriptions(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithAMSSubscriptions", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithAMSSubscriptions indicates an expected call of WithAMSSubscriptions.
func (mr *MockTransitionHandlerMockRecorder) WithAMSSubscriptions(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithAMSSubscriptions", reflect.TypeOf((*MockTransitionHandler)(nil).WithAMSSubscriptions), sw, args)
}

// areAllHostsDone mocks base method.
func (m *MockTransitionHandler) areAllHostsDone(sw stateswitch.StateSwitch, arg1 stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "areAllHostsDone", sw, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// areAllHostsDone indicates an expected call of areAllHostsDone.
func (mr *MockTransitionHandlerMockRecorder) areAllHostsDone(sw, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "areAllHostsDone", reflect.TypeOf((*MockTransitionHandler)(nil).areAllHostsDone), sw, arg1)
}

// hasClusterCompleteInstallation mocks base method.
func (m *MockTransitionHandler) hasClusterCompleteInstallation(sw stateswitch.StateSwitch, args stateswitch.TransitionArgs) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "hasClusterCompleteInstallation", sw, args)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// hasClusterCompleteInstallation indicates an expected call of hasClusterCompleteInstallation.
func (mr *MockTransitionHandlerMockRecorder) hasClusterCompleteInstallation(sw, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "hasClusterCompleteInstallation", reflect.TypeOf((*MockTransitionHandler)(nil).hasClusterCompleteInstallation), sw, args)
}
