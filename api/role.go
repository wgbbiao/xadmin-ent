package api

import (
	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent/xadminrole"
)

// 角色管理

func RoleList(ctx iris.Context) {
	var form struct {
		Name string `json:"name"`
	}
	if err := ctx.ReadJSON(&form); err != nil {
		ctx.JSON(iris.Map{
			"status": 1,
			"msg":    "表单读取错误",
		})
		return
	}
	result := AmisResult{}
	result.Status = 0
	result.Msg = "角色列表"
	q := database.GetDb().XadminRole.Query()
	if form.Name != "" {
		q = q.Where(xadminrole.NameContains(form.Name))
	}
	roles, err := q.All(ctx.Request().Context())
	cnt, _ := q.Count(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "角色列表获取失败"
	} else {
		result.Data = iris.Map{
			"items": roles,
			"total": cnt,
		}
	}
	ctx.JSON(result)
}
