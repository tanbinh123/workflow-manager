// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/aws-sdk-go/service/sfn/sfniface (interfaces: SFNAPI)

package mock_sfniface

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	sfn "github.com/aws/aws-sdk-go/service/sfn"
	gomock "github.com/golang/mock/gomock"
)

// Mock of SFNAPI interface
type MockSFNAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockSFNAPIRecorder
}

// Recorder for MockSFNAPI (not exported)
type _MockSFNAPIRecorder struct {
	mock *MockSFNAPI
}

func NewMockSFNAPI(ctrl *gomock.Controller) *MockSFNAPI {
	mock := &MockSFNAPI{ctrl: ctrl}
	mock.recorder = &_MockSFNAPIRecorder{mock}
	return mock
}

func (_m *MockSFNAPI) EXPECT() *_MockSFNAPIRecorder {
	return _m.recorder
}

func (_m *MockSFNAPI) CreateActivity(_param0 *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateActivity", _param0)
	ret0, _ := ret[0].(*sfn.CreateActivityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) CreateActivity(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateActivity", arg0)
}

func (_m *MockSFNAPI) CreateActivityRequest(_param0 *sfn.CreateActivityInput) (*request.Request, *sfn.CreateActivityOutput) {
	ret := _m.ctrl.Call(_m, "CreateActivityRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.CreateActivityOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) CreateActivityRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateActivityRequest", arg0)
}

func (_m *MockSFNAPI) CreateActivityWithContext(_param0 aws.Context, _param1 *sfn.CreateActivityInput, _param2 ...request.Option) (*sfn.CreateActivityOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CreateActivityWithContext", _s...)
	ret0, _ := ret[0].(*sfn.CreateActivityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) CreateActivityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateActivityWithContext", _s...)
}

func (_m *MockSFNAPI) CreateStateMachine(_param0 *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateStateMachine", _param0)
	ret0, _ := ret[0].(*sfn.CreateStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) CreateStateMachine(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateStateMachine", arg0)
}

func (_m *MockSFNAPI) CreateStateMachineRequest(_param0 *sfn.CreateStateMachineInput) (*request.Request, *sfn.CreateStateMachineOutput) {
	ret := _m.ctrl.Call(_m, "CreateStateMachineRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.CreateStateMachineOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) CreateStateMachineRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateStateMachineRequest", arg0)
}

func (_m *MockSFNAPI) CreateStateMachineWithContext(_param0 aws.Context, _param1 *sfn.CreateStateMachineInput, _param2 ...request.Option) (*sfn.CreateStateMachineOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CreateStateMachineWithContext", _s...)
	ret0, _ := ret[0].(*sfn.CreateStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) CreateStateMachineWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateStateMachineWithContext", _s...)
}

func (_m *MockSFNAPI) DeleteActivity(_param0 *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteActivity", _param0)
	ret0, _ := ret[0].(*sfn.DeleteActivityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DeleteActivity(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteActivity", arg0)
}

func (_m *MockSFNAPI) DeleteActivityRequest(_param0 *sfn.DeleteActivityInput) (*request.Request, *sfn.DeleteActivityOutput) {
	ret := _m.ctrl.Call(_m, "DeleteActivityRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.DeleteActivityOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DeleteActivityRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteActivityRequest", arg0)
}

func (_m *MockSFNAPI) DeleteActivityWithContext(_param0 aws.Context, _param1 *sfn.DeleteActivityInput, _param2 ...request.Option) (*sfn.DeleteActivityOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeleteActivityWithContext", _s...)
	ret0, _ := ret[0].(*sfn.DeleteActivityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DeleteActivityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteActivityWithContext", _s...)
}

func (_m *MockSFNAPI) DeleteStateMachine(_param0 *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteStateMachine", _param0)
	ret0, _ := ret[0].(*sfn.DeleteStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DeleteStateMachine(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteStateMachine", arg0)
}

func (_m *MockSFNAPI) DeleteStateMachineRequest(_param0 *sfn.DeleteStateMachineInput) (*request.Request, *sfn.DeleteStateMachineOutput) {
	ret := _m.ctrl.Call(_m, "DeleteStateMachineRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.DeleteStateMachineOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DeleteStateMachineRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteStateMachineRequest", arg0)
}

func (_m *MockSFNAPI) DeleteStateMachineWithContext(_param0 aws.Context, _param1 *sfn.DeleteStateMachineInput, _param2 ...request.Option) (*sfn.DeleteStateMachineOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeleteStateMachineWithContext", _s...)
	ret0, _ := ret[0].(*sfn.DeleteStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DeleteStateMachineWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteStateMachineWithContext", _s...)
}

func (_m *MockSFNAPI) DescribeActivity(_param0 *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeActivity", _param0)
	ret0, _ := ret[0].(*sfn.DescribeActivityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeActivity(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeActivity", arg0)
}

func (_m *MockSFNAPI) DescribeActivityRequest(_param0 *sfn.DescribeActivityInput) (*request.Request, *sfn.DescribeActivityOutput) {
	ret := _m.ctrl.Call(_m, "DescribeActivityRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.DescribeActivityOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeActivityRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeActivityRequest", arg0)
}

func (_m *MockSFNAPI) DescribeActivityWithContext(_param0 aws.Context, _param1 *sfn.DescribeActivityInput, _param2 ...request.Option) (*sfn.DescribeActivityOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DescribeActivityWithContext", _s...)
	ret0, _ := ret[0].(*sfn.DescribeActivityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeActivityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeActivityWithContext", _s...)
}

func (_m *MockSFNAPI) DescribeExecution(_param0 *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeExecution", _param0)
	ret0, _ := ret[0].(*sfn.DescribeExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeExecution(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeExecution", arg0)
}

func (_m *MockSFNAPI) DescribeExecutionRequest(_param0 *sfn.DescribeExecutionInput) (*request.Request, *sfn.DescribeExecutionOutput) {
	ret := _m.ctrl.Call(_m, "DescribeExecutionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.DescribeExecutionOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeExecutionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeExecutionRequest", arg0)
}

func (_m *MockSFNAPI) DescribeExecutionWithContext(_param0 aws.Context, _param1 *sfn.DescribeExecutionInput, _param2 ...request.Option) (*sfn.DescribeExecutionOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DescribeExecutionWithContext", _s...)
	ret0, _ := ret[0].(*sfn.DescribeExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeExecutionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeExecutionWithContext", _s...)
}

func (_m *MockSFNAPI) DescribeStateMachine(_param0 *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeStateMachine", _param0)
	ret0, _ := ret[0].(*sfn.DescribeStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeStateMachine(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeStateMachine", arg0)
}

func (_m *MockSFNAPI) DescribeStateMachineRequest(_param0 *sfn.DescribeStateMachineInput) (*request.Request, *sfn.DescribeStateMachineOutput) {
	ret := _m.ctrl.Call(_m, "DescribeStateMachineRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.DescribeStateMachineOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeStateMachineRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeStateMachineRequest", arg0)
}

func (_m *MockSFNAPI) DescribeStateMachineWithContext(_param0 aws.Context, _param1 *sfn.DescribeStateMachineInput, _param2 ...request.Option) (*sfn.DescribeStateMachineOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DescribeStateMachineWithContext", _s...)
	ret0, _ := ret[0].(*sfn.DescribeStateMachineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) DescribeStateMachineWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeStateMachineWithContext", _s...)
}

func (_m *MockSFNAPI) GetActivityTask(_param0 *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "GetActivityTask", _param0)
	ret0, _ := ret[0].(*sfn.GetActivityTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) GetActivityTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetActivityTask", arg0)
}

func (_m *MockSFNAPI) GetActivityTaskRequest(_param0 *sfn.GetActivityTaskInput) (*request.Request, *sfn.GetActivityTaskOutput) {
	ret := _m.ctrl.Call(_m, "GetActivityTaskRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.GetActivityTaskOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) GetActivityTaskRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetActivityTaskRequest", arg0)
}

func (_m *MockSFNAPI) GetActivityTaskWithContext(_param0 aws.Context, _param1 *sfn.GetActivityTaskInput, _param2 ...request.Option) (*sfn.GetActivityTaskOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetActivityTaskWithContext", _s...)
	ret0, _ := ret[0].(*sfn.GetActivityTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) GetActivityTaskWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetActivityTaskWithContext", _s...)
}

func (_m *MockSFNAPI) GetExecutionHistory(_param0 *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
	ret := _m.ctrl.Call(_m, "GetExecutionHistory", _param0)
	ret0, _ := ret[0].(*sfn.GetExecutionHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) GetExecutionHistory(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetExecutionHistory", arg0)
}

func (_m *MockSFNAPI) GetExecutionHistoryPages(_param0 *sfn.GetExecutionHistoryInput, _param1 func(*sfn.GetExecutionHistoryOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "GetExecutionHistoryPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) GetExecutionHistoryPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetExecutionHistoryPages", arg0, arg1)
}

func (_m *MockSFNAPI) GetExecutionHistoryPagesWithContext(_param0 aws.Context, _param1 *sfn.GetExecutionHistoryInput, _param2 func(*sfn.GetExecutionHistoryOutput, bool) bool, _param3 ...request.Option) error {
	_s := []interface{}{_param0, _param1, _param2}
	for _, _x := range _param3 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetExecutionHistoryPagesWithContext", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) GetExecutionHistoryPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetExecutionHistoryPagesWithContext", _s...)
}

func (_m *MockSFNAPI) GetExecutionHistoryRequest(_param0 *sfn.GetExecutionHistoryInput) (*request.Request, *sfn.GetExecutionHistoryOutput) {
	ret := _m.ctrl.Call(_m, "GetExecutionHistoryRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.GetExecutionHistoryOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) GetExecutionHistoryRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetExecutionHistoryRequest", arg0)
}

func (_m *MockSFNAPI) GetExecutionHistoryWithContext(_param0 aws.Context, _param1 *sfn.GetExecutionHistoryInput, _param2 ...request.Option) (*sfn.GetExecutionHistoryOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetExecutionHistoryWithContext", _s...)
	ret0, _ := ret[0].(*sfn.GetExecutionHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) GetExecutionHistoryWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetExecutionHistoryWithContext", _s...)
}

func (_m *MockSFNAPI) ListActivities(_param0 *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListActivities", _param0)
	ret0, _ := ret[0].(*sfn.ListActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListActivities(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListActivities", arg0)
}

func (_m *MockSFNAPI) ListActivitiesPages(_param0 *sfn.ListActivitiesInput, _param1 func(*sfn.ListActivitiesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListActivitiesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) ListActivitiesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListActivitiesPages", arg0, arg1)
}

func (_m *MockSFNAPI) ListActivitiesPagesWithContext(_param0 aws.Context, _param1 *sfn.ListActivitiesInput, _param2 func(*sfn.ListActivitiesOutput, bool) bool, _param3 ...request.Option) error {
	_s := []interface{}{_param0, _param1, _param2}
	for _, _x := range _param3 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListActivitiesPagesWithContext", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) ListActivitiesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListActivitiesPagesWithContext", _s...)
}

func (_m *MockSFNAPI) ListActivitiesRequest(_param0 *sfn.ListActivitiesInput) (*request.Request, *sfn.ListActivitiesOutput) {
	ret := _m.ctrl.Call(_m, "ListActivitiesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.ListActivitiesOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListActivitiesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListActivitiesRequest", arg0)
}

func (_m *MockSFNAPI) ListActivitiesWithContext(_param0 aws.Context, _param1 *sfn.ListActivitiesInput, _param2 ...request.Option) (*sfn.ListActivitiesOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListActivitiesWithContext", _s...)
	ret0, _ := ret[0].(*sfn.ListActivitiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListActivitiesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListActivitiesWithContext", _s...)
}

func (_m *MockSFNAPI) ListExecutions(_param0 *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
	ret := _m.ctrl.Call(_m, "ListExecutions", _param0)
	ret0, _ := ret[0].(*sfn.ListExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListExecutions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListExecutions", arg0)
}

func (_m *MockSFNAPI) ListExecutionsPages(_param0 *sfn.ListExecutionsInput, _param1 func(*sfn.ListExecutionsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListExecutionsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) ListExecutionsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListExecutionsPages", arg0, arg1)
}

func (_m *MockSFNAPI) ListExecutionsPagesWithContext(_param0 aws.Context, _param1 *sfn.ListExecutionsInput, _param2 func(*sfn.ListExecutionsOutput, bool) bool, _param3 ...request.Option) error {
	_s := []interface{}{_param0, _param1, _param2}
	for _, _x := range _param3 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListExecutionsPagesWithContext", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) ListExecutionsPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListExecutionsPagesWithContext", _s...)
}

func (_m *MockSFNAPI) ListExecutionsRequest(_param0 *sfn.ListExecutionsInput) (*request.Request, *sfn.ListExecutionsOutput) {
	ret := _m.ctrl.Call(_m, "ListExecutionsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.ListExecutionsOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListExecutionsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListExecutionsRequest", arg0)
}

func (_m *MockSFNAPI) ListExecutionsWithContext(_param0 aws.Context, _param1 *sfn.ListExecutionsInput, _param2 ...request.Option) (*sfn.ListExecutionsOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListExecutionsWithContext", _s...)
	ret0, _ := ret[0].(*sfn.ListExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListExecutionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListExecutionsWithContext", _s...)
}

func (_m *MockSFNAPI) ListStateMachines(_param0 *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListStateMachines", _param0)
	ret0, _ := ret[0].(*sfn.ListStateMachinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListStateMachines(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListStateMachines", arg0)
}

func (_m *MockSFNAPI) ListStateMachinesPages(_param0 *sfn.ListStateMachinesInput, _param1 func(*sfn.ListStateMachinesOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListStateMachinesPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) ListStateMachinesPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListStateMachinesPages", arg0, arg1)
}

func (_m *MockSFNAPI) ListStateMachinesPagesWithContext(_param0 aws.Context, _param1 *sfn.ListStateMachinesInput, _param2 func(*sfn.ListStateMachinesOutput, bool) bool, _param3 ...request.Option) error {
	_s := []interface{}{_param0, _param1, _param2}
	for _, _x := range _param3 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListStateMachinesPagesWithContext", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockSFNAPIRecorder) ListStateMachinesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListStateMachinesPagesWithContext", _s...)
}

func (_m *MockSFNAPI) ListStateMachinesRequest(_param0 *sfn.ListStateMachinesInput) (*request.Request, *sfn.ListStateMachinesOutput) {
	ret := _m.ctrl.Call(_m, "ListStateMachinesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.ListStateMachinesOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListStateMachinesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListStateMachinesRequest", arg0)
}

func (_m *MockSFNAPI) ListStateMachinesWithContext(_param0 aws.Context, _param1 *sfn.ListStateMachinesInput, _param2 ...request.Option) (*sfn.ListStateMachinesOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListStateMachinesWithContext", _s...)
	ret0, _ := ret[0].(*sfn.ListStateMachinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) ListStateMachinesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListStateMachinesWithContext", _s...)
}

func (_m *MockSFNAPI) SendTaskFailure(_param0 *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
	ret := _m.ctrl.Call(_m, "SendTaskFailure", _param0)
	ret0, _ := ret[0].(*sfn.SendTaskFailureOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskFailure(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskFailure", arg0)
}

func (_m *MockSFNAPI) SendTaskFailureRequest(_param0 *sfn.SendTaskFailureInput) (*request.Request, *sfn.SendTaskFailureOutput) {
	ret := _m.ctrl.Call(_m, "SendTaskFailureRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.SendTaskFailureOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskFailureRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskFailureRequest", arg0)
}

func (_m *MockSFNAPI) SendTaskFailureWithContext(_param0 aws.Context, _param1 *sfn.SendTaskFailureInput, _param2 ...request.Option) (*sfn.SendTaskFailureOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SendTaskFailureWithContext", _s...)
	ret0, _ := ret[0].(*sfn.SendTaskFailureOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskFailureWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskFailureWithContext", _s...)
}

func (_m *MockSFNAPI) SendTaskHeartbeat(_param0 *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
	ret := _m.ctrl.Call(_m, "SendTaskHeartbeat", _param0)
	ret0, _ := ret[0].(*sfn.SendTaskHeartbeatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskHeartbeat(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskHeartbeat", arg0)
}

func (_m *MockSFNAPI) SendTaskHeartbeatRequest(_param0 *sfn.SendTaskHeartbeatInput) (*request.Request, *sfn.SendTaskHeartbeatOutput) {
	ret := _m.ctrl.Call(_m, "SendTaskHeartbeatRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.SendTaskHeartbeatOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskHeartbeatRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskHeartbeatRequest", arg0)
}

func (_m *MockSFNAPI) SendTaskHeartbeatWithContext(_param0 aws.Context, _param1 *sfn.SendTaskHeartbeatInput, _param2 ...request.Option) (*sfn.SendTaskHeartbeatOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SendTaskHeartbeatWithContext", _s...)
	ret0, _ := ret[0].(*sfn.SendTaskHeartbeatOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskHeartbeatWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskHeartbeatWithContext", _s...)
}

func (_m *MockSFNAPI) SendTaskSuccess(_param0 *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
	ret := _m.ctrl.Call(_m, "SendTaskSuccess", _param0)
	ret0, _ := ret[0].(*sfn.SendTaskSuccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskSuccess(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskSuccess", arg0)
}

func (_m *MockSFNAPI) SendTaskSuccessRequest(_param0 *sfn.SendTaskSuccessInput) (*request.Request, *sfn.SendTaskSuccessOutput) {
	ret := _m.ctrl.Call(_m, "SendTaskSuccessRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.SendTaskSuccessOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskSuccessRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskSuccessRequest", arg0)
}

func (_m *MockSFNAPI) SendTaskSuccessWithContext(_param0 aws.Context, _param1 *sfn.SendTaskSuccessInput, _param2 ...request.Option) (*sfn.SendTaskSuccessOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SendTaskSuccessWithContext", _s...)
	ret0, _ := ret[0].(*sfn.SendTaskSuccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) SendTaskSuccessWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendTaskSuccessWithContext", _s...)
}

func (_m *MockSFNAPI) StartExecution(_param0 *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
	ret := _m.ctrl.Call(_m, "StartExecution", _param0)
	ret0, _ := ret[0].(*sfn.StartExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) StartExecution(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartExecution", arg0)
}

func (_m *MockSFNAPI) StartExecutionRequest(_param0 *sfn.StartExecutionInput) (*request.Request, *sfn.StartExecutionOutput) {
	ret := _m.ctrl.Call(_m, "StartExecutionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.StartExecutionOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) StartExecutionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartExecutionRequest", arg0)
}

func (_m *MockSFNAPI) StartExecutionWithContext(_param0 aws.Context, _param1 *sfn.StartExecutionInput, _param2 ...request.Option) (*sfn.StartExecutionOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "StartExecutionWithContext", _s...)
	ret0, _ := ret[0].(*sfn.StartExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) StartExecutionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartExecutionWithContext", _s...)
}

func (_m *MockSFNAPI) StopExecution(_param0 *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
	ret := _m.ctrl.Call(_m, "StopExecution", _param0)
	ret0, _ := ret[0].(*sfn.StopExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) StopExecution(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopExecution", arg0)
}

func (_m *MockSFNAPI) StopExecutionRequest(_param0 *sfn.StopExecutionInput) (*request.Request, *sfn.StopExecutionOutput) {
	ret := _m.ctrl.Call(_m, "StopExecutionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sfn.StopExecutionOutput)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) StopExecutionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopExecutionRequest", arg0)
}

func (_m *MockSFNAPI) StopExecutionWithContext(_param0 aws.Context, _param1 *sfn.StopExecutionInput, _param2 ...request.Option) (*sfn.StopExecutionOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "StopExecutionWithContext", _s...)
	ret0, _ := ret[0].(*sfn.StopExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSFNAPIRecorder) StopExecutionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StopExecutionWithContext", _s...)
}