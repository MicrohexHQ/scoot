// Automatically generated by MockGen. DO NOT EDIT!
// Source: queue.go

package queue

import (
	gomock "github.com/golang/mock/gomock"
	sched "github.com/scootdev/scoot/sched"
)

// Mock of Queue interface
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *_MockQueueRecorder
}

// Recorder for MockQueue (not exported)
type _MockQueueRecorder struct {
	mock *MockQueue
}

func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &_MockQueueRecorder{mock}
	return mock
}

func (_m *MockQueue) EXPECT() *_MockQueueRecorder {
	return _m.recorder
}

func (_m *MockQueue) Enqueue(job sched.JobDefinition) (string, error) {
	ret := _m.ctrl.Call(_m, "Enqueue", job)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockQueueRecorder) Enqueue(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Enqueue", arg0)
}

func (_m *MockQueue) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockQueueRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

// Mock of WorkItem interface
type MockWorkItem struct {
	ctrl     *gomock.Controller
	recorder *_MockWorkItemRecorder
}

// Recorder for MockWorkItem (not exported)
type _MockWorkItemRecorder struct {
	mock *MockWorkItem
}

func NewMockWorkItem(ctrl *gomock.Controller) *MockWorkItem {
	mock := &MockWorkItem{ctrl: ctrl}
	mock.recorder = &_MockWorkItemRecorder{mock}
	return mock
}

func (_m *MockWorkItem) EXPECT() *_MockWorkItemRecorder {
	return _m.recorder
}

func (_m *MockWorkItem) Job() sched.Job {
	ret := _m.ctrl.Call(_m, "Job")
	ret0, _ := ret[0].(sched.Job)
	return ret0
}

func (_mr *_MockWorkItemRecorder) Job() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Job")
}

func (_m *MockWorkItem) Dequeue() {
	_m.ctrl.Call(_m, "Dequeue")
}

func (_mr *_MockWorkItemRecorder) Dequeue() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Dequeue")
}