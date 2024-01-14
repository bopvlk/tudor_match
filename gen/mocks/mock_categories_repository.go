// Code generated by MockGen. DO NOT EDIT.
// Source: study_marketplace/pkg/repositories (interfaces: CategoriesRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	queries "study_marketplace/database/queries"

	gomock "github.com/golang/mock/gomock"
)

// MockCategoriesRepository is a mock of CategoriesRepository interface.
type MockCategoriesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoriesRepositoryMockRecorder
}

// MockCategoriesRepositoryMockRecorder is the mock recorder for MockCategoriesRepository.
type MockCategoriesRepositoryMockRecorder struct {
	mock *MockCategoriesRepository
}

// NewMockCategoriesRepository creates a new mock instance.
func NewMockCategoriesRepository(ctrl *gomock.Controller) *MockCategoriesRepository {
	mock := &MockCategoriesRepository{ctrl: ctrl}
	mock.recorder = &MockCategoriesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoriesRepository) EXPECT() *MockCategoriesRepositoryMockRecorder {
	return m.recorder
}

// GetCategoriesWithChildren mocks base method.
func (m *MockCategoriesRepository) GetCategoriesWithChildren(arg0 context.Context) ([]queries.GetCategoriesWithChildrenRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesWithChildren", arg0)
	ret0, _ := ret[0].([]queries.GetCategoriesWithChildrenRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoriesWithChildren indicates an expected call of GetCategoriesWithChildren.
func (mr *MockCategoriesRepositoryMockRecorder) GetCategoriesWithChildren(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesWithChildren", reflect.TypeOf((*MockCategoriesRepository)(nil).GetCategoriesWithChildren), arg0)
}

// GetCategoryAndParent mocks base method.
func (m *MockCategoriesRepository) GetCategoryAndParent(arg0 context.Context, arg1 string) (queries.GetCategoryAndParentRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryAndParent", arg0, arg1)
	ret0, _ := ret[0].(queries.GetCategoryAndParentRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryAndParent indicates an expected call of GetCategoryAndParent.
func (mr *MockCategoriesRepositoryMockRecorder) GetCategoryAndParent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryAndParent", reflect.TypeOf((*MockCategoriesRepository)(nil).GetCategoryAndParent), arg0, arg1)
}

// GetCategoryByID mocks base method.
func (m *MockCategoriesRepository) GetCategoryByID(arg0 context.Context, arg1 int) (queries.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryByID", arg0, arg1)
	ret0, _ := ret[0].(queries.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryByID indicates an expected call of GetCategoryByID.
func (mr *MockCategoriesRepositoryMockRecorder) GetCategoryByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryByID", reflect.TypeOf((*MockCategoriesRepository)(nil).GetCategoryByID), arg0, arg1)
}

// GetCategoryByName mocks base method.
func (m *MockCategoriesRepository) GetCategoryByName(arg0 context.Context, arg1 string) (queries.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryByName", arg0, arg1)
	ret0, _ := ret[0].(queries.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryByName indicates an expected call of GetCategoryByName.
func (mr *MockCategoriesRepositoryMockRecorder) GetCategoryByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryByName", reflect.TypeOf((*MockCategoriesRepository)(nil).GetCategoryByName), arg0, arg1)
}

// GetCategoryParents mocks base method.
func (m *MockCategoriesRepository) GetCategoryParents(arg0 context.Context) ([]queries.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryParents", arg0)
	ret0, _ := ret[0].([]queries.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategoryParents indicates an expected call of GetCategoryParents.
func (mr *MockCategoriesRepositoryMockRecorder) GetCategoryParents(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryParents", reflect.TypeOf((*MockCategoriesRepository)(nil).GetCategoryParents), arg0)
}