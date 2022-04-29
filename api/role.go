package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
	"github.com/wgbbiao/xadminent/ent/xadminpermission"
	"github.com/wgbbiao/xadminent/ent/xadminrole"
)

// 角色管理

func RoleList(ctx iris.Context) {
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

// 创建角色
func RoleAdd(ctx iris.Context) {
	var form struct {
		Name          string `json:"name" validate:"required"`
		PermissionIDs []int  `json:"permission_ids" validate:"required"`
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
				Key("errors", validationErrors)
		}
		ctx.JSON(result)
		return
	}
	role := database.GetDb().XadminRole.Create()
	role.AddPermissionIDs(form.PermissionIDs...).SetName(form.Name)
	if r, err := role.Save(ctx.Request().Context()); err != nil {
		result.Status = 1
		result.Msg = "角色创建失败"
	} else {
		result.Msg = "角色创建成功"
		result.Data = r
	}
	ctx.JSON(result)
}

// 角色详情
func RoleDetail(ctx iris.Context) {
	var result AmisResult = AmisResult{}
	id, _ := ctx.Params().GetInt("id")
	role, err := database.GetDb().XadminRole.Query().
		Where(xadminrole.ID(id)).
		WithPermissions(func(xpq *ent.XadminPermissionQuery) {
			xpq.Select(xadminpermission.FieldID)
		}).
		First(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "角色获取失败"
	} else {
		result.Msg = "角色获取成功"
		result.Data = role
	}
	ctx.JSON(result)
}

// 更新角色
func RoleUpdate(ctx iris.Context) {
	var form struct {
		Name          string `json:"name" validate:"required"`
		PermissionIDs []int  `json:"permission_ids" validate:"required"`
	}
	id, _ := ctx.Params().GetInt("id")
	var result AmisResult = AmisResult{}
	if err := ctx.ReadJSON(&form); err != nil {
		result.Status = 1
		if errs, ok := err.(validator.ValidationErrors); ok {
			result.Msg = "表单验证错误"
			validationErrors := wrapValidationErrors(errs)
			result.Data = iris.NewProblem().
				Title("表单验证失败").
				Detail("One or more fields failed to be validated").
				Key("errors", validationErrors)
		}
		ctx.JSON(result)
		return
	}
	role := database.GetDb().XadminRole.UpdateOneID(id)
	role.ClearPermissions().AddPermissionIDs(form.PermissionIDs...).
		SetName(form.Name)
	if r, err := role.Save(ctx.Request().Context()); err != nil {
		result.Status = 1
		result.Msg = "角色更新失败"
	} else {
		result.Msg = "角色更新成功"
		result.Data = r
	}
	ctx.JSON(result)
}

// 删除角色
func RoleDelete(ctx iris.Context) {
	var result AmisResult = AmisResult{}
	id, _ := ctx.Params().GetInt("id")
	if err := database.GetDb().XadminRole.DeleteOneID(id).Exec(ctx.Request().Context()); err != nil {
		result.Status = 1
		result.Msg = "角色删除失败"
	} else {
		result.Msg = "角色删除成功"
	}
	ctx.JSON(result)
}
