package main

import (
	"github.com/iris-contrib/middleware/cors"
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
	app.UseRouter(cors.AllowAll())
	app.Options("{root:path}", func(context iris.Context) {
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Headers", "Origin,Authorization,Content-Type,Accept,X-Total,X-Limit,X-Offset")
		context.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,HEAD,PATCH")
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Expose-Headers", "Content-Length,Content-Encoding,Content-Type")
	})
	{
		xadmin := app.Party("/admin")
		xadminent.AddRoute(xadmin)
	}
	app.Run(iris.Addr(":7070"))
}
