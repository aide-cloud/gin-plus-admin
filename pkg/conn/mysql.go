package conn

import (
	"database/sql"
	"sync"

	ginplus "github.com/aide-cloud/gin-plus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_db  *gorm.DB
	once sync.Once
)

func init() {
	once.Do(func() {
		sqlDB, err := sql.Open("mysql", "root:localhost@tcp(localhost:3306)/gin-plus-admin?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		db, err := gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}))
		if err != nil {
			panic(err)
		}

		db = db.Debug()

		if err := db.Use(ginplus.NewOpentracingPlugin()); err != nil {
			panic(err)
		}
		_db = db
	})
}

func GetDB() *gorm.DB {
	return _db
}
