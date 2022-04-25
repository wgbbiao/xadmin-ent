package cmd

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
	"github.com/wgbbiao/xadminent"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "开启web服务",
	Long:  `开启web服务，默认端口为7070`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}
