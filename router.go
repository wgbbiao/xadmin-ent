package xadminent

import (
	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent/api"
	"github.com/wgbbiao/xadminent/j"
)

func AddRoute(r iris.Party) {
	r.Post("/login", api.Login)
	user := r.Party("/user", j.VerifyMiddleware())
	user.Get("info", api.UserInfo)
}