// Code generated by MockGen. DO NOT EDIT.
// Source: domain/deckOfCards/service.go

// Package mock_deckOfCards is a generated GoMock package.
package mock_deckOfCards

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	dao "github.com/yoratyo/card-games/model/dao"
	deck "github.com/yoratyo/card-games/model/dto/deck"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateNewDeck mocks base method.
func (m *MockService) CreateNewDeck(ctx *gin.Context, req deck.CreateNewDeckRequestDTO) (dao.Deck, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewDeck", ctx, req)
	ret0, _ := ret[0].(dao.Deck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewDeck indicates an expected call of CreateNewDeck.
func (mr *MockServiceMockRecorder) CreateNewDeck(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewDeck", reflect.TypeOf((*MockService)(nil).CreateNewDeck), ctx, req)
}

// DrawCards mocks base method.
func (m *MockService) DrawCards(ctx *gin.Context, ID uuid.UUID, req deck.DrawCardsFromDeckRequestDTO) (deck.DrawCardsFromDeckResponseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DrawCards", ctx, ID, req)
	ret0, _ := ret[0].(deck.DrawCardsFromDeckResponseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DrawCards indicates an expected call of DrawCards.
func (mr *MockServiceMockRecorder) DrawCards(ctx, ID, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DrawCards", reflect.TypeOf((*MockService)(nil).DrawCards), ctx, ID, req)
}

// GetDetailByID mocks base method.
func (m *MockService) GetDetailByID(ctx *gin.Context, ID uuid.UUID) (deck.OpenDeckResponseDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailByID", ctx, ID)
	ret0, _ := ret[0].(deck.OpenDeckResponseDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailByID indicates an expected call of GetDetailByID.
func (mr *MockServiceMockRecorder) GetDetailByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailByID", reflect.TypeOf((*MockService)(nil).GetDetailByID), ctx, ID)
}