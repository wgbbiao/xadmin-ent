package api

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
	"github.com/wgbbiao/xadminent/ent/xadminpermission"
	"github.com/wgbbiao/xadminent/ent/xadminrole"
	"github.com/wgbbiao/xadminent/ent/xadminuser"
)

// 检查管理是否有权限进行操作

func CheckPermission(user *ent.XadminUser, model any, permission string) bool {
	if user.IsSuper {
		return true
	}
	rt := reflect.TypeOf(model)
	code := strings.ToLower(fmt.Sprintf("%s_%s", rt.Name(), permission))
	client := database.GetDb()
	// 查询是否存在权限管理
	{
		cnt, err := client.XadminPermission.Query().
			Where(xadminpermission.Code(code)).Count(context.Background())
		if err != nil {
			panic(err)
		}
		// 没有权限管理,任何人都可以操作
		if cnt == 0 {
			return true
		}
	}
	// 查询权限
	{
		cnt, err := client.XadminUser.Query().
			Where(
				xadminuser.ID(user.ID),
				xadminuser.HasPermissionsWith(xadminpermission.Code(code)),
			).
			Count(context.Background())
		if err != nil {
			fmt.Println(err)
		} else {
			if cnt > 0 {
				return true
			}
		}
	}
	// 查询角色
	{
		cnt, _ := client.XadminUser.Query().
			Where(
				xadminuser.ID(user.ID),
				xadminuser.HasRolesWith(
					xadminrole.HasPermissionsWith(
						xadminpermission.Code(code),
					),
				)).
			Count(context.Background())
		if cnt > 0 {
			return true
		}
	}
	return false
}
