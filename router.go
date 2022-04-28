package xadminent

import (
	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent/api"
	"github.com/wgbbiao/xadminent/j"
)

func AddRoute(r iris.Party) {
	r.Post("/login", api.Login)
	{
		user := r.Party("/user", j.VerifyMiddleware())
		user.Get("/info", api.UserInfo)
		user.Get("/list", api.UserList)
		user.Post("/add", api.UserAdd)
		user.Post("/edit", api.UserEdit)
		user.Get("/{id:int}", api.UserDetail)
		user.Post("/modify_password", api.UserPassword)
		user.Post("/delete", api.UserDelete)
	}
	{
		role := r.Party("/role", j.VerifyMiddleware())
		role.Get("/list", api.RoleList)
	}
	{
		permission := r.Party("/permission", j.VerifyMiddleware())
		permission.Get("/list", api.PermissionList)
		permission.Post("/add", api.PermissionAdd)
	}
	{
		contenttype := r.Party("/contenttype", j.VerifyMiddleware())
		contenttype.Get("/list", api.ContentTypeList)
		contenttype.Get("/{id:int}", api.ContentTypeDetail)
		contenttype.Put("/{id:int}", api.ContentTypeEdit)
		contenttype.Delete("/{id:int}", api.ContentTypeDelete)
	}
}
