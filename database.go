package xadminent

import (
	"fmt"
	"sync"

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
	var err error
	clientSync.Do(func() {
		client, err = ent.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
				database.User,
				database.Password,
				database.Host,
				database.Port,
				database.Database))
		if err != nil {
			panic(err)
		}
	})
	return client
}
