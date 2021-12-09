package database

// I get this Gorm wrapper from https://github.com/kacperjurak/igorm
// I only adjust some part to support latest gorm version

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

// Gormw is an interface which DB implements
type Gormw interface {
	Close() error
	DB() *sql.DB
	Where(query interface{}, args ...interface{}) Gormw
	Or(query interface{}, args ...interface{}) Gormw
	Not(query interface{}, args ...interface{}) Gormw
	Limit(value int) Gormw
	Offset(value int) Gormw
	Order(value interface{}) Gormw
	Select(query interface{}, args ...interface{}) Gormw
	Omit(columns ...string) Gormw
	Group(query string) Gormw
	Having(query string, values ...interface{}) Gormw
	Joins(query string, args ...interface{}) Gormw
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gormw
	Unscoped() Gormw
	Attrs(attrs ...interface{}) Gormw
	Assign(attrs ...interface{}) Gormw
	First(out interface{}, where ...interface{}) Gormw
	Take(out interface{}, where ...interface{}) Gormw
	Last(out interface{}, where ...interface{}) Gormw
	Find(out interface{}, where ...interface{}) Gormw
	Scan(dest interface{}) Gormw
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	ScanRows(rows *sql.Rows, result interface{}) error
	Pluck(column string, value interface{}) Gormw
	Count(count *int64) Gormw
	FirstOrInit(out interface{}, where ...interface{}) Gormw
	FirstOrCreate(out interface{}, where ...interface{}) Gormw
	Update(column string, value interface{}) Gormw
	Updates(values interface{}) Gormw
	UpdateColumn(column string, value interface{}) Gormw
	UpdateColumns(values interface{}) Gormw
	Save(value interface{}) Gormw
	Create(value interface{}) Gormw
	Delete(value interface{}, where ...interface{}) Gormw
	Raw(sql string, values ...interface{}) Gormw
	Exec(sql string, values ...interface{}) Gormw
	Model(value interface{}) Gormw
	Table(name string) Gormw
	Debug() Gormw
	Begin() Gormw
	Commit() Gormw
	Rollback() Gormw
	AutoMigrate(values ...interface{}) error
	Association(column string) *gorm.Association
	Preload(column string, conditions ...interface{}) Gormw
	Set(name string, value interface{}) Gormw
	Get(name string) (value interface{}, ok bool)
	AddError(err error) error

	// extra
	Error() error
	RowsAffected() int64
	WithContext(ctx context.Context) Gormw
}

type gormw struct {
	w *gorm.DB
}

// Openw is a drop-in replacement for Open()
func Openw(dialect gorm.Dialector, args ...gorm.Option) (db Gormw, err error) {
	gormdb, err := gorm.Open(dialect, args...)
	return Wrap(gormdb), err
}

// Wrap wraps gorm.DB in an interface
func Wrap(db *gorm.DB) Gormw {
	return &gormw{db}
}

func (it *gormw) Close() error {
	return it.DB().Close()
}

func (it *gormw) DB() *sql.DB {
	db, _ := it.w.DB()
	return db
}

func (it *gormw) Where(query interface{}, args ...interface{}) Gormw {
	return Wrap(it.w.Where(query, args...))
}

func (it *gormw) Or(query interface{}, args ...interface{}) Gormw {
	return Wrap(it.w.Or(query, args...))
}

func (it *gormw) Not(query interface{}, args ...interface{}) Gormw {
	return Wrap(it.w.Not(query, args...))
}

func (it *gormw) Limit(value int) Gormw {
	return Wrap(it.w.Limit(value))
}

func (it *gormw) Offset(value int) Gormw {
	return Wrap(it.w.Offset(value))
}

func (it *gormw) Order(value interface{}) Gormw {
	return Wrap(it.w.Order(value))
}

func (it *gormw) Select(query interface{}, args ...interface{}) Gormw {
	return Wrap(it.w.Select(query, args...))
}

func (it *gormw) Omit(columns ...string) Gormw {
	return Wrap(it.w.Omit(columns...))
}

func (it *gormw) Group(query string) Gormw {
	return Wrap(it.w.Group(query))
}

func (it *gormw) Having(query string, values ...interface{}) Gormw {
	return Wrap(it.w.Having(query, values...))
}

func (it *gormw) Joins(query string, args ...interface{}) Gormw {
	return Wrap(it.w.Joins(query, args...))
}

func (it *gormw) Scopes(funcs ...func(*gorm.DB) *gorm.DB) Gormw {
	return Wrap(it.w.Scopes(funcs...))
}

func (it *gormw) Unscoped() Gormw {
	return Wrap(it.w.Unscoped())
}

func (it *gormw) Attrs(attrs ...interface{}) Gormw {
	return Wrap(it.w.Attrs(attrs...))
}

func (it *gormw) Assign(attrs ...interface{}) Gormw {
	return Wrap(it.w.Assign(attrs...))
}

func (it *gormw) First(out interface{}, where ...interface{}) Gormw {
	return Wrap(it.w.First(out, where...))
}

func (it *gormw) Take(out interface{}, where ...interface{}) Gormw {
	return Wrap(it.w.Take(out, where...))
}

func (it *gormw) Last(out interface{}, where ...interface{}) Gormw {
	return Wrap(it.w.Last(out, where...))
}

func (it *gormw) Find(out interface{}, where ...interface{}) Gormw {
	return Wrap(it.w.Find(out, where...))
}

func (it *gormw) Scan(dest interface{}) Gormw {
	return Wrap(it.w.Scan(dest))
}

func (it *gormw) Row() *sql.Row {
	return it.w.Row()
}

func (it *gormw) Rows() (*sql.Rows, error) {
	return it.w.Rows()
}

func (it *gormw) ScanRows(rows *sql.Rows, result interface{}) error {
	return it.w.ScanRows(rows, result)
}

func (it *gormw) Pluck(column string, value interface{}) Gormw {
	return Wrap(it.w.Pluck(column, value))
}

func (it *gormw) Count(count *int64) Gormw {
	return Wrap(it.w.Count(count))
}

func (it *gormw) FirstOrInit(out interface{}, where ...interface{}) Gormw {
	return Wrap(it.w.FirstOrInit(out, where...))
}

func (it *gormw) FirstOrCreate(out interface{}, where ...interface{}) Gormw {
	return Wrap(it.w.FirstOrCreate(out, where...))
}

func (it *gormw) Update(column string, value interface{}) Gormw {
	return Wrap(it.w.Update(column, value))
}

func (it *gormw) Updates(values interface{}) Gormw {
	return Wrap(it.w.Updates(values))
}

func (it *gormw) UpdateColumn(column string, value interface{}) Gormw {
	return Wrap(it.w.UpdateColumn(column, value))
}

func (it *gormw) UpdateColumns(values interface{}) Gormw {
	return Wrap(it.w.UpdateColumns(values))
}

func (it *gormw) Save(value interface{}) Gormw {
	return Wrap(it.w.Save(value))
}

func (it *gormw) Create(value interface{}) Gormw {
	return Wrap(it.w.Create(value))
}

func (it *gormw) Delete(value interface{}, where ...interface{}) Gormw {
	return Wrap(it.w.Delete(value, where...))
}

func (it *gormw) Raw(sql string, values ...interface{}) Gormw {
	return Wrap(it.w.Raw(sql, values...))
}

func (it *gormw) Exec(sql string, values ...interface{}) Gormw {
	return Wrap(it.w.Exec(sql, values...))
}

func (it *gormw) Model(value interface{}) Gormw {
	return Wrap(it.w.Model(value))
}

func (it *gormw) Table(name string) Gormw {
	return Wrap(it.w.Table(name))
}

func (it *gormw) Debug() Gormw {
	return Wrap(it.w.Debug())
}

func (it *gormw) Begin() Gormw {
	return Wrap(it.w.Begin())
}

func (it *gormw) Commit() Gormw {
	return Wrap(it.w.Commit())
}

func (it *gormw) Rollback() Gormw {
	return Wrap(it.w.Rollback())
}

func (it *gormw) AutoMigrate(values ...interface{}) error {
	return it.w.AutoMigrate(values...)
}

func (it *gormw) Association(column string) *gorm.Association {
	return it.w.Association(column)
}

func (it *gormw) Preload(column string, conditions ...interface{}) Gormw {
	return Wrap(it.w.Preload(column, conditions...))
}

func (it *gormw) Set(name string, value interface{}) Gormw {
	return Wrap(it.w.Set(name, value))
}

func (it *gormw) Get(name string) (interface{}, bool) {
	return it.w.Get(name)
}

func (it *gormw) AddError(err error) error {
	return it.w.AddError(err)
}

func (it *gormw) RowsAffected() int64 {
	return it.w.RowsAffected
}

func (it *gormw) Error() error {
	return it.w.Error
}

func (it *gormw) WithContext(ctx context.Context) Gormw {
	return Wrap(it.w.WithContext(ctx))
}
