package api

import (
	"github.com/go-playground/validator/v10"
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
		Name          string `json:"name" validate:"required"`
		Code          string `json:"code" validate:"required"`
		ContentTypeID int    `json:"content_type_id" validate:"required"`
	}
	var result AmisResult = AmisResult{}
	if err := ctx.ReadJSON(&form); err != nil {
		result.Status = 1
		if errs, ok := err.(validator.ValidationErrors); ok {
			result.Msg = "表单验证错误"
			validationErrors := wrapValidationErrors(errs)
			result.Data = iris.NewProblem().
				Title("表单验证失败").
				Detail("One or more fields failed to be validated").
				Type("/Permission/validation-errors").
				Key("errors", validationErrors)
		} else {
			result.Msg = "表单读取错误"
		}
	} else {
		client := database.GetDb()
		per := client.XadminPermission.Create()
		p, err := per.SetName(form.Name).
			SetCode(form.Code).
			SetContentTypeID(form.ContentTypeID).
			Save(ctx.Request().Context())

		if err != nil {
			result.Status = 1
			result.Msg = "权限创建失败"
		} else {
			result.Status = 0
			result.Msg = "创建成功"
			result.Data = p
		}
	}
	ctx.JSON(result)
}

// 修改权限
func PermissionEdit(ctx iris.Context) {
	var form struct {
		ID            int    `json:"id"`
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
	per := client.XadminPermission.UpdateOneID(form.ID)
	p, err := per.
		SetName(form.Name).
		SetCode(form.Code).
		SetContentTypeID(form.ContentTypeID).
		Save(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "权限修改失败"
	} else {
		result.Status = 0
		result.Msg = "修改成功"
		result.Data = p
	}
	ctx.JSON(result)
}

// 删除权限
func PermissionDelete(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	var result AmisResult = AmisResult{}

	client := database.GetDb()
	err := client.XadminPermission.DeleteOneID(id).Exec(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "权限删除失败"
	}
	result.Status = 0
	result.Msg = "删除成功"
	ctx.JSON(result)
}

// 获取权限详情
func PermissionDetail(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	var result AmisResult = AmisResult{}
	p, err := database.GetDb().XadminPermission.Query().
		Where(xadminpermission.ID(id)).
		First(ctx.Request().Context())

	if err != nil {
		result.Status = 1
		result.Msg = "权限详情获取失败"
	}
	result.Status = 0
	result.Msg = "权限详情获取成功"
	result.Data = p
	ctx.JSON(result)
}
