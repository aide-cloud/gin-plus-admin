package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        uint                  `gorm:"primary_key" json:"id"`
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
	DeletedAt soft_delete.DeletedAt `json:"-"`
}

type Scopemethod = func(db *gorm.DB) *gorm.DB

type Pagination interface {
	Page() int
	Size() int
	SetTotal(total int64)
}

// Paginate 分页
func Paginate(pgInfo Pagination) Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		if pgInfo == nil {
			return db
		}
		return db.Limit(pgInfo.Size()).Offset((pgInfo.Page() - 1) * pgInfo.Size())
	}
}

// WhereInColumn 通过字段名和值列表进行查询
func WhereInColumn[T any](column string, vals ...T) Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		idsLen := len(vals)
		switch idsLen {
		case 0:
			return db
		case 1:
			return db.Where("`"+column+"` = ?", vals[0])
		default:
			return db.Where("`"+column+"`in (?)", vals)
		}
	}
}

// WhereID 通过ID列表进行查询
func WhereID(ids ...int) Scopemethod {
	return WhereInColumn("id", ids...)
}

// WhereLikeKeyword 模糊查询
func WhereLikeKeyword(keyword string, columns ...string) Scopemethod {
	return func(db *gorm.DB) *gorm.DB {
		if keyword == "" || len(columns) == 0 {
			return db
		}
		likeKeyword := "%" + keyword + "%"
		for _, column := range columns {
			db = db.Or("`"+column+"` LIKE ?", likeKeyword)
		}
		return db
	}
}
