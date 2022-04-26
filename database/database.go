package database

import (
	"fmt"
	"sync"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wgbbiao/xadminent/ent"
)

type DatabaseConfig struct {
	// 地址
	Host string
	// 端口
	Port int
	// 用户名
	User string
	// 密码
	Password string
	// 数据库名
	Database string
}

var database *DatabaseConfig

func SetDatabase(db *DatabaseConfig) {
	database = db
}

var client *ent.Client
var clientSync sync.Once

func GetDb() *ent.Client {
	clientSync.Do(func() {
		drv, err := sql.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
				database.User,
				database.Password,
				database.Host,
				database.Port,
				database.Database))
		if err != nil {
			panic(err)
		}
		// 获取数据库驱动中的sql.DB对象。
		db := drv.DB()
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour)
		client = ent.NewClient(ent.Driver(drv))
	})
	return client
}
