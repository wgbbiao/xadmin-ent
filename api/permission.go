package api

import (
	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent/xadminpermission"
)

// 权限接口

// 获取权限列表
func PermissionList(ctx iris.Context) {
	var form struct {
		Name string `json:"name"`
	}
	if err := ctx.ReadQuery(&form); err != nil {
		ctx.JSON(iris.Map{
			"status": 1,
			"msg":    "表单读取错误",
		})
		return
	}
	result := AmisResult{}
	result.Status = 0
	result.Msg = "权限列表"
	q := database.GetDb().XadminPermission.Query()
	if form.Name != "" {
		q = q.Where(xadminpermission.NameContains(form.Name))
	}
	permissions, err := q.All(ctx.Request().Context())
	cnt, _ := q.Count(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "权限列表获取失败"
	} else {
		result.Data = iris.Map{
			"items": permissions,
			"total": cnt,
		}
	}
	ctx.JSON(result)
}

// 创建权限
func PermissionAdd(ctx iris.Context) {
	var form struct {
		Name          string `json:"name"`
		Code          string `json:"code"`
		ContentTypeID int    `json:"content_type_id"`
	}
	var result AmisResult = AmisResult{}
	if err := ctx.ReadJSON(&form); err != nil {
		result.Status = 1
		result.Msg = "表单读取错误"
	}
	client := database.GetDb()
	per := client.XadminPermission.Create()
	p, err := per.SetName(form.Name).
		SetCode(form.Code).
		SetContentTypeID(form.ContentTypeID).
		Save(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "权限创建失败"
	}
	result.Status = 0
	result.Msg = "创建成功"
	result.Data = p
	ctx.JSON(result)
}
