// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/ticket_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/akasaa101/ticketing/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockTicketService is a mock of TicketService interface.
type MockTicketService struct {
	ctrl     *gomock.Controller
	recorder *MockTicketServiceMockRecorder
}

// MockTicketServiceMockRecorder is the mock recorder for MockTicketService.
type MockTicketServiceMockRecorder struct {
	mock *MockTicketService
}

// NewMockTicketService creates a new mock instance.
func NewMockTicketService(ctrl *gomock.Controller) *MockTicketService {
	mock := &MockTicketService{ctrl: ctrl}
	mock.recorder = &MockTicketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketService) EXPECT() *MockTicketServiceMockRecorder {
	return m.recorder
}

// TicketGetById mocks base method.
func (m *MockTicketService) TicketGetById(id int16) (models.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TicketGetById", id)
	ret0, _ := ret[0].(models.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TicketGetById indicates an expected call of TicketGetById.
func (mr *MockTicketServiceMockRecorder) TicketGetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TicketGetById", reflect.TypeOf((*MockTicketService)(nil).TicketGetById), id)
}

// TicketInsert mocks base method.
func (m *MockTicketService) TicketInsert(ticket models.Ticket) (models.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TicketInsert", ticket)
	ret0, _ := ret[0].(models.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TicketInsert indicates an expected call of TicketInsert.
func (mr *MockTicketServiceMockRecorder) TicketInsert(ticket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TicketInsert", reflect.TypeOf((*MockTicketService)(nil).TicketInsert), ticket)
}
