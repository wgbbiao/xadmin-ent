package main

import (
	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent"
	"github.com/wgbbiao/xadminent/database"
)

func main() {
	xadminent.SetDatabase(&database.DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Database: "testent",
	})
	xadminent.AutoMigrate()

	app := iris.New()
	xadminent.AddRoute(app)
	app.Run(iris.Addr(":7070"))
}
