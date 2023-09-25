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

func InitMysqlDB(dsn string, debug bool) {
	once.Do(func() {
		sqlDB, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		db, err := gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}))
		if err != nil {
			panic(err)
		}

		if debug {
			db = db.Debug()
		}

		if err := db.Use(ginplus.NewOpentracingPlugin()); err != nil {
			panic(err)
		}
		_db = db
	})
}

func GetMysqlDB() *gorm.DB {
	if _db == nil {
		panic("mysql db not init")
	}
	return _db
}
