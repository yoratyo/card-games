// Code generated by MockGen. DO NOT EDIT.
// Source: config/database/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "github.com/yoratyo/card-games/config/database"
	gorm "gorm.io/gorm"
)

// MockGormw is a mock of Gormw interface.
type MockGormw struct {
	ctrl     *gomock.Controller
	recorder *MockGormwMockRecorder
}

// MockGormwMockRecorder is the mock recorder for MockGormw.
type MockGormwMockRecorder struct {
	mock *MockGormw
}

// NewMockGormw creates a new mock instance.
func NewMockGormw(ctrl *gomock.Controller) *MockGormw {
	mock := &MockGormw{ctrl: ctrl}
	mock.recorder = &MockGormwMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGormw) EXPECT() *MockGormwMockRecorder {
	return m.recorder
}

// AddError mocks base method.
func (m *MockGormw) AddError(err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddError", err)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddError indicates an expected call of AddError.
func (mr *MockGormwMockRecorder) AddError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddError", reflect.TypeOf((*MockGormw)(nil).AddError), err)
}

// Assign mocks base method.
func (m *MockGormw) Assign(attrs ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range attrs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Assign", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Assign indicates an expected call of Assign.
func (mr *MockGormwMockRecorder) Assign(attrs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*MockGormw)(nil).Assign), attrs...)
}

// Association mocks base method.
func (m *MockGormw) Association(column string) *gorm.Association {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Association", column)
	ret0, _ := ret[0].(*gorm.Association)
	return ret0
}

// Association indicates an expected call of Association.
func (mr *MockGormwMockRecorder) Association(column interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Association", reflect.TypeOf((*MockGormw)(nil).Association), column)
}

// Attrs mocks base method.
func (m *MockGormw) Attrs(attrs ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range attrs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Attrs", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Attrs indicates an expected call of Attrs.
func (mr *MockGormwMockRecorder) Attrs(attrs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attrs", reflect.TypeOf((*MockGormw)(nil).Attrs), attrs...)
}

// AutoMigrate mocks base method.
func (m *MockGormw) AutoMigrate(values ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AutoMigrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate.
func (mr *MockGormwMockRecorder) AutoMigrate(values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockGormw)(nil).AutoMigrate), values...)
}

// Begin mocks base method.
func (m *MockGormw) Begin() database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockGormwMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockGormw)(nil).Begin))
}

// Close mocks base method.
func (m *MockGormw) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockGormwMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGormw)(nil).Close))
}

// Commit mocks base method.
func (m *MockGormw) Commit() database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockGormwMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockGormw)(nil).Commit))
}

// Count mocks base method.
func (m *MockGormw) Count(count *int64) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", count)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockGormwMockRecorder) Count(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockGormw)(nil).Count), count)
}

// Create mocks base method.
func (m *MockGormw) Create(value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockGormwMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGormw)(nil).Create), value)
}

// DB mocks base method.
func (m *MockGormw) DB() *sql.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockGormwMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockGormw)(nil).DB))
}

// Debug mocks base method.
func (m *MockGormw) Debug() database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug")
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Debug indicates an expected call of Debug.
func (mr *MockGormwMockRecorder) Debug() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockGormw)(nil).Debug))
}

// Delete mocks base method.
func (m *MockGormw) Delete(value interface{}, where ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{value}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockGormwMockRecorder) Delete(value interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{value}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGormw)(nil).Delete), varargs...)
}

// Error mocks base method.
func (m *MockGormw) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockGormwMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockGormw)(nil).Error))
}

// Exec mocks base method.
func (m *MockGormw) Exec(sql string, values ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockGormwMockRecorder) Exec(sql interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockGormw)(nil).Exec), varargs...)
}

// Find mocks base method.
func (m *MockGormw) Find(out interface{}, where ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockGormwMockRecorder) Find(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockGormw)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockGormw) First(out interface{}, where ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockGormwMockRecorder) First(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockGormw)(nil).First), varargs...)
}

// FirstOrCreate mocks base method.
func (m *MockGormw) FirstOrCreate(out interface{}, where ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FirstOrCreate", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// FirstOrCreate indicates an expected call of FirstOrCreate.
func (mr *MockGormwMockRecorder) FirstOrCreate(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrCreate", reflect.TypeOf((*MockGormw)(nil).FirstOrCreate), varargs...)
}

// FirstOrInit mocks base method.
func (m *MockGormw) FirstOrInit(out interface{}, where ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FirstOrInit", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// FirstOrInit indicates an expected call of FirstOrInit.
func (mr *MockGormwMockRecorder) FirstOrInit(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstOrInit", reflect.TypeOf((*MockGormw)(nil).FirstOrInit), varargs...)
}

// Get mocks base method.
func (m *MockGormw) Get(name string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGormwMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGormw)(nil).Get), name)
}

// Group mocks base method.
func (m *MockGormw) Group(query string) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", query)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Group indicates an expected call of Group.
func (mr *MockGormwMockRecorder) Group(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockGormw)(nil).Group), query)
}

// Having mocks base method.
func (m *MockGormw) Having(query string, values ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Having", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Having indicates an expected call of Having.
func (mr *MockGormwMockRecorder) Having(query interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Having", reflect.TypeOf((*MockGormw)(nil).Having), varargs...)
}

// Joins mocks base method.
func (m *MockGormw) Joins(query string, args ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Joins", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Joins indicates an expected call of Joins.
func (mr *MockGormwMockRecorder) Joins(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Joins", reflect.TypeOf((*MockGormw)(nil).Joins), varargs...)
}

// Last mocks base method.
func (m *MockGormw) Last(out interface{}, where ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Last", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockGormwMockRecorder) Last(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockGormw)(nil).Last), varargs...)
}

// Limit mocks base method.
func (m *MockGormw) Limit(value int) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit", value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Limit indicates an expected call of Limit.
func (mr *MockGormwMockRecorder) Limit(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockGormw)(nil).Limit), value)
}

// Model mocks base method.
func (m *MockGormw) Model(value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model", value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockGormwMockRecorder) Model(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockGormw)(nil).Model), value)
}

// Not mocks base method.
func (m *MockGormw) Not(query interface{}, args ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Not", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Not indicates an expected call of Not.
func (mr *MockGormwMockRecorder) Not(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Not", reflect.TypeOf((*MockGormw)(nil).Not), varargs...)
}

// Offset mocks base method.
func (m *MockGormw) Offset(value int) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset", value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Offset indicates an expected call of Offset.
func (mr *MockGormwMockRecorder) Offset(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockGormw)(nil).Offset), value)
}

// Omit mocks base method.
func (m *MockGormw) Omit(columns ...string) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range columns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Omit", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Omit indicates an expected call of Omit.
func (mr *MockGormwMockRecorder) Omit(columns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Omit", reflect.TypeOf((*MockGormw)(nil).Omit), columns...)
}

// Or mocks base method.
func (m *MockGormw) Or(query interface{}, args ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Or", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Or indicates an expected call of Or.
func (mr *MockGormwMockRecorder) Or(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Or", reflect.TypeOf((*MockGormw)(nil).Or), varargs...)
}

// Order mocks base method.
func (m *MockGormw) Order(value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Order", value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Order indicates an expected call of Order.
func (mr *MockGormwMockRecorder) Order(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Order", reflect.TypeOf((*MockGormw)(nil).Order), value)
}

// Pluck mocks base method.
func (m *MockGormw) Pluck(column string, value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pluck", column, value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Pluck indicates an expected call of Pluck.
func (mr *MockGormwMockRecorder) Pluck(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pluck", reflect.TypeOf((*MockGormw)(nil).Pluck), column, value)
}

// Preload mocks base method.
func (m *MockGormw) Preload(column string, conditions ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{column}
	for _, a := range conditions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Preload", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Preload indicates an expected call of Preload.
func (mr *MockGormwMockRecorder) Preload(column interface{}, conditions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{column}, conditions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Preload", reflect.TypeOf((*MockGormw)(nil).Preload), varargs...)
}

// Raw mocks base method.
func (m *MockGormw) Raw(sql string, values ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{sql}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Raw", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *MockGormwMockRecorder) Raw(sql interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sql}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockGormw)(nil).Raw), varargs...)
}

// Rollback mocks base method.
func (m *MockGormw) Rollback() database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockGormwMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockGormw)(nil).Rollback))
}

// Row mocks base method.
func (m *MockGormw) Row() *sql.Row {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Row")
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// Row indicates an expected call of Row.
func (mr *MockGormwMockRecorder) Row() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Row", reflect.TypeOf((*MockGormw)(nil).Row))
}

// Rows mocks base method.
func (m *MockGormw) Rows() (*sql.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rows")
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rows indicates an expected call of Rows.
func (mr *MockGormwMockRecorder) Rows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rows", reflect.TypeOf((*MockGormw)(nil).Rows))
}

// RowsAffected mocks base method.
func (m *MockGormw) RowsAffected() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowsAffected")
	ret0, _ := ret[0].(int64)
	return ret0
}

// RowsAffected indicates an expected call of RowsAffected.
func (mr *MockGormwMockRecorder) RowsAffected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowsAffected", reflect.TypeOf((*MockGormw)(nil).RowsAffected))
}

// Save mocks base method.
func (m *MockGormw) Save(value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockGormwMockRecorder) Save(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockGormw)(nil).Save), value)
}

// Scan mocks base method.
func (m *MockGormw) Scan(dest interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", dest)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockGormwMockRecorder) Scan(dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockGormw)(nil).Scan), dest)
}

// ScanRows mocks base method.
func (m *MockGormw) ScanRows(rows *sql.Rows, result interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanRows", rows, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanRows indicates an expected call of ScanRows.
func (mr *MockGormwMockRecorder) ScanRows(rows, result interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanRows", reflect.TypeOf((*MockGormw)(nil).ScanRows), rows, result)
}

// Scopes mocks base method.
func (m *MockGormw) Scopes(funcs ...func(*gorm.DB) *gorm.DB) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range funcs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scopes", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Scopes indicates an expected call of Scopes.
func (mr *MockGormwMockRecorder) Scopes(funcs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scopes", reflect.TypeOf((*MockGormw)(nil).Scopes), funcs...)
}

// Select mocks base method.
func (m *MockGormw) Select(query interface{}, args ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockGormwMockRecorder) Select(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockGormw)(nil).Select), varargs...)
}

// Set mocks base method.
func (m *MockGormw) Set(name string, value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", name, value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockGormwMockRecorder) Set(name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockGormw)(nil).Set), name, value)
}

// Table mocks base method.
func (m *MockGormw) Table(name string) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Table", name)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Table indicates an expected call of Table.
func (mr *MockGormwMockRecorder) Table(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Table", reflect.TypeOf((*MockGormw)(nil).Table), name)
}

// Take mocks base method.
func (m *MockGormw) Take(out interface{}, where ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Take", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Take indicates an expected call of Take.
func (mr *MockGormwMockRecorder) Take(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*MockGormw)(nil).Take), varargs...)
}

// Unscoped mocks base method.
func (m *MockGormw) Unscoped() database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unscoped")
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Unscoped indicates an expected call of Unscoped.
func (mr *MockGormwMockRecorder) Unscoped() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unscoped", reflect.TypeOf((*MockGormw)(nil).Unscoped))
}

// Update mocks base method.
func (m *MockGormw) Update(column string, value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", column, value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockGormwMockRecorder) Update(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGormw)(nil).Update), column, value)
}

// UpdateColumn mocks base method.
func (m *MockGormw) UpdateColumn(column string, value interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", column, value)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// UpdateColumn indicates an expected call of UpdateColumn.
func (mr *MockGormwMockRecorder) UpdateColumn(column, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*MockGormw)(nil).UpdateColumn), column, value)
}

// UpdateColumns mocks base method.
func (m *MockGormw) UpdateColumns(values interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumns", values)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// UpdateColumns indicates an expected call of UpdateColumns.
func (mr *MockGormwMockRecorder) UpdateColumns(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumns", reflect.TypeOf((*MockGormw)(nil).UpdateColumns), values)
}

// Updates mocks base method.
func (m *MockGormw) Updates(values interface{}) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", values)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockGormwMockRecorder) Updates(values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockGormw)(nil).Updates), values)
}

// Where mocks base method.
func (m *MockGormw) Where(query interface{}, args ...interface{}) database.Gormw {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Where", varargs...)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// Where indicates an expected call of Where.
func (mr *MockGormwMockRecorder) Where(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockGormw)(nil).Where), varargs...)
}

// WithContext mocks base method.
func (m *MockGormw) WithContext(ctx context.Context) database.Gormw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(database.Gormw)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockGormwMockRecorder) WithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockGormw)(nil).WithContext), ctx)
}