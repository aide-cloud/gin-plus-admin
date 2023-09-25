package query

import (
	"context"

	"gin-plus-admin/pkg/model"

	"gorm.io/gorm"
)

var _ IAction[any] = (*Action[any])(nil)

type IAction[T any] interface {
	// First 查询单条数据
	First(ctx context.Context, wheres ...model.Scopemethod) (*T, error)
	// Last 查询单条数据
	Last(ctx context.Context, wheres ...model.Scopemethod) (*T, error)
	// List 查询多条数据
	List(ctx context.Context, pgInfo model.Pagination, wheres ...model.Scopemethod) ([]*T, error)
	// Create 创建数据
	Create(ctx context.Context, m *T) error
	// Update 更新数据
	Update(ctx context.Context, m *T, wheres ...model.Scopemethod) error
	// Delete 删除数据
	Delete(ctx context.Context, wheres ...model.Scopemethod) error
}

type Action[T any] struct {
	db *gorm.DB
}

type ActionOption[T any] func(a *Action[T])

func NewAction[T any](opts ...ActionOption[T]) *Action[T] {
	action := Action[T]{}

	for _, opt := range opts {
		opt(&action)
	}

	return &action
}

func WithDB[T any](db *gorm.DB) ActionOption[T] {
	return func(a *Action[T]) {
		a.db = db
	}
}

// First 查询单条数据
func (a *Action[T]) First(ctx context.Context, wheres ...model.Scopemethod) (*T, error) {
	var m T

	if err := a.db.WithContext(ctx).Scopes(wheres...).First(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

// Last 查询单条数据
func (a *Action[T]) Last(ctx context.Context, wheres ...model.Scopemethod) (*T, error) {
	var m T

	if err := a.db.Scopes(wheres...).Last(&m).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

// List 查询多条数据
func (a *Action[T]) List(ctx context.Context, pgInfo model.Pagination, wheres ...model.Scopemethod) ([]*T, error) {
	var ms []*T

	db := a.db.Scopes(wheres...)

	if pgInfo != nil {
		var total int64
		if err := db.Count(&total).Error; err != nil {
			return nil, err
		}
		pgInfo.SetTotal(total)
	}

	if err := db.Find(&ms).Error; err != nil {
		return nil, err
	}

	return ms, nil
}

// Create 创建数据
func (a *Action[T]) Create(ctx context.Context, newModel *T) error {
	return a.db.Create(newModel).Error
}

// Update 更新数据
func (a *Action[T]) Update(ctx context.Context, newModel *T, wheres ...model.Scopemethod) error {
	return a.db.Scopes(wheres...).Updates(newModel).Error
}

// Delete 删除数据
func (a *Action[T]) Delete(ctx context.Context, wheres ...model.Scopemethod) error {
	var m T
	return a.db.Scopes(wheres...).Delete(&m).Error
}
