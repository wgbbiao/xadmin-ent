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
		user.Put("/{id:int}", api.UserEdit)
		user.Get("/{id:int}", api.UserDetail)
		user.Post("/modify_password", api.UserPassword)
		user.Post("/delete", api.UserDelete)
	}
	{
		role := r.Party("/role", j.VerifyMiddleware())
		role.Get("/list", api.RoleList)
		role.Post("/", api.RoleAdd)
		role.Put("/{id:int}", api.RoleUpdate)
		role.Delete("/{id:int}", api.RoleDelete)
		role.Get("/{id:int}", api.RoleDetail)
	}
	{
		permission := r.Party("/permission", j.VerifyMiddleware())
		permission.Get("/list", api.PermissionList)
		permission.Post("/", api.PermissionAdd)
		permission.Get("/{id:int}", api.PermissionDetail)
		permission.Put("/{id:int}", api.PermissionEdit)
		permission.Delete("/{id:int}", api.PermissionDelete)

	}
	{
		contenttype := r.Party("/contenttype", j.VerifyMiddleware())
		contenttype.Get("/list", api.ContentTypeList)
		contenttype.Get("/{id:int}", api.ContentTypeDetail)
		contenttype.Put("/{id:int}", api.ContentTypeEdit)
		contenttype.Delete("/{id:int}", api.ContentTypeDelete)
	}
}
