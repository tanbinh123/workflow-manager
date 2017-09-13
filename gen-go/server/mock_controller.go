// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package server is a generated GoMock package.
package server

import (
	context "context"
	models "github.com/Clever/workflow-manager/gen-go/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method
func (m *MockController) HealthCheck(ctx context.Context) error {
	ret := m.ctrl.Call(m, "HealthCheck", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockControllerMockRecorder) HealthCheck(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockController)(nil).HealthCheck), ctx)
}

// PostStateResource mocks base method
func (m *MockController) PostStateResource(ctx context.Context, i *models.NewStateResource) (*models.StateResource, error) {
	ret := m.ctrl.Call(m, "PostStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostStateResource indicates an expected call of PostStateResource
func (mr *MockControllerMockRecorder) PostStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostStateResource", reflect.TypeOf((*MockController)(nil).PostStateResource), ctx, i)
}

// DeleteStateResource mocks base method
func (m *MockController) DeleteStateResource(ctx context.Context, i *models.DeleteStateResourceInput) error {
	ret := m.ctrl.Call(m, "DeleteStateResource", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStateResource indicates an expected call of DeleteStateResource
func (mr *MockControllerMockRecorder) DeleteStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStateResource", reflect.TypeOf((*MockController)(nil).DeleteStateResource), ctx, i)
}

// GetStateResource mocks base method
func (m *MockController) GetStateResource(ctx context.Context, i *models.GetStateResourceInput) (*models.StateResource, error) {
	ret := m.ctrl.Call(m, "GetStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateResource indicates an expected call of GetStateResource
func (mr *MockControllerMockRecorder) GetStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateResource", reflect.TypeOf((*MockController)(nil).GetStateResource), ctx, i)
}

// PutStateResource mocks base method
func (m *MockController) PutStateResource(ctx context.Context, i *models.PutStateResourceInput) (*models.StateResource, error) {
	ret := m.ctrl.Call(m, "PutStateResource", ctx, i)
	ret0, _ := ret[0].(*models.StateResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutStateResource indicates an expected call of PutStateResource
func (mr *MockControllerMockRecorder) PutStateResource(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutStateResource", reflect.TypeOf((*MockController)(nil).PutStateResource), ctx, i)
}

// GetWorkflowDefinitions mocks base method
func (m *MockController) GetWorkflowDefinitions(ctx context.Context) ([]models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "GetWorkflowDefinitions", ctx)
	ret0, _ := ret[0].([]models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowDefinitions indicates an expected call of GetWorkflowDefinitions
func (mr *MockControllerMockRecorder) GetWorkflowDefinitions(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowDefinitions", reflect.TypeOf((*MockController)(nil).GetWorkflowDefinitions), ctx)
}

// NewWorkflowDefinition mocks base method
func (m *MockController) NewWorkflowDefinition(ctx context.Context, i *models.NewWorkflowDefinitionRequest) (*models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "NewWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWorkflowDefinition indicates an expected call of NewWorkflowDefinition
func (mr *MockControllerMockRecorder) NewWorkflowDefinition(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWorkflowDefinition", reflect.TypeOf((*MockController)(nil).NewWorkflowDefinition), ctx, i)
}

// GetWorkflowDefinitionVersionsByName mocks base method
func (m *MockController) GetWorkflowDefinitionVersionsByName(ctx context.Context, i *models.GetWorkflowDefinitionVersionsByNameInput) ([]models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "GetWorkflowDefinitionVersionsByName", ctx, i)
	ret0, _ := ret[0].([]models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowDefinitionVersionsByName indicates an expected call of GetWorkflowDefinitionVersionsByName
func (mr *MockControllerMockRecorder) GetWorkflowDefinitionVersionsByName(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowDefinitionVersionsByName", reflect.TypeOf((*MockController)(nil).GetWorkflowDefinitionVersionsByName), ctx, i)
}

// UpdateWorkflowDefinition mocks base method
func (m *MockController) UpdateWorkflowDefinition(ctx context.Context, i *models.UpdateWorkflowDefinitionInput) (*models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "UpdateWorkflowDefinition", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkflowDefinition indicates an expected call of UpdateWorkflowDefinition
func (mr *MockControllerMockRecorder) UpdateWorkflowDefinition(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkflowDefinition", reflect.TypeOf((*MockController)(nil).UpdateWorkflowDefinition), ctx, i)
}

// GetWorkflowDefinitionByNameAndVersion mocks base method
func (m *MockController) GetWorkflowDefinitionByNameAndVersion(ctx context.Context, i *models.GetWorkflowDefinitionByNameAndVersionInput) (*models.WorkflowDefinition, error) {
	ret := m.ctrl.Call(m, "GetWorkflowDefinitionByNameAndVersion", ctx, i)
	ret0, _ := ret[0].(*models.WorkflowDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowDefinitionByNameAndVersion indicates an expected call of GetWorkflowDefinitionByNameAndVersion
func (mr *MockControllerMockRecorder) GetWorkflowDefinitionByNameAndVersion(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowDefinitionByNameAndVersion", reflect.TypeOf((*MockController)(nil).GetWorkflowDefinitionByNameAndVersion), ctx, i)
}

// GetWorkflows mocks base method
func (m *MockController) GetWorkflows(ctx context.Context, i *models.GetWorkflowsInput) ([]models.Workflow, error) {
	ret := m.ctrl.Call(m, "GetWorkflows", ctx, i)
	ret0, _ := ret[0].([]models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflows indicates an expected call of GetWorkflows
func (mr *MockControllerMockRecorder) GetWorkflows(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflows", reflect.TypeOf((*MockController)(nil).GetWorkflows), ctx, i)
}

// StartWorkflow mocks base method
func (m *MockController) StartWorkflow(ctx context.Context, i *models.WorkflowInput) (*models.Workflow, error) {
	ret := m.ctrl.Call(m, "StartWorkflow", ctx, i)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflow indicates an expected call of StartWorkflow
func (mr *MockControllerMockRecorder) StartWorkflow(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflow", reflect.TypeOf((*MockController)(nil).StartWorkflow), ctx, i)
}

// CancelWorkflow mocks base method
func (m *MockController) CancelWorkflow(ctx context.Context, i *models.CancelWorkflowInput) error {
	ret := m.ctrl.Call(m, "CancelWorkflow", ctx, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelWorkflow indicates an expected call of CancelWorkflow
func (mr *MockControllerMockRecorder) CancelWorkflow(ctx, i interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWorkflow", reflect.TypeOf((*MockController)(nil).CancelWorkflow), ctx, i)
}

// GetWorkflowByID mocks base method
func (m *MockController) GetWorkflowByID(ctx context.Context, workflowId string) (*models.Workflow, error) {
	ret := m.ctrl.Call(m, "GetWorkflowByID", ctx, workflowId)
	ret0, _ := ret[0].(*models.Workflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflowByID indicates an expected call of GetWorkflowByID
func (mr *MockControllerMockRecorder) GetWorkflowByID(ctx, workflowId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflowByID", reflect.TypeOf((*MockController)(nil).GetWorkflowByID), ctx, workflowId)
}
