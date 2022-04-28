package api

import (
	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent/xadmincontenttype"
)

// 文档列表
func ContentTypeList(ctx iris.Context) {
	var form struct {
		AppLabel string `json:"app_label"`
		Model    string `json:"model"`
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
	result.Msg = "文档列表"
	q := database.GetDb().XadminContenttype.Query()
	if form.AppLabel != "" {
		q = q.Where(xadmincontenttype.AppLabel(form.AppLabel))
	}
	if form.Model != "" {
		q = q.Where(xadmincontenttype.ModelContains(form.Model))
	}
	contenttypes, err := q.All(ctx.Request().Context())
	cnt, _ := q.Count(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "文档列表获取失败"
	} else {
		result.Data = iris.Map{
			"items": contenttypes,
			"total": cnt,
		}
	}
	ctx.JSON(result)
}

// 文档详情
func ContentTypeDetail(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	result := AmisResult{}
	result.Status = 0
	result.Msg = "文档详情"
	contenttype, err := database.GetDb().XadminContenttype.Get(ctx.Request().Context(), id)
	if err != nil {
		result.Status = 1
		result.Msg = "文档详情获取失败"
	} else {
		result.Data = contenttype
	}
	ctx.JSON(result)
}

// 修改文档
func ContentTypeEdit(ctx iris.Context) {
	var form struct {
		ID       int    `json:"id"`
		AppLabel string `json:"app_label"`
		Model    string `json:"model"`
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
	result.Msg = "修改文档"

	if contenttype, err := database.GetDb().XadminContenttype.UpdateOneID(form.ID).
		SetAppLabel(form.AppLabel).
		SetModel(form.Model).Save(ctx.Request().Context()); err != nil {
		result.Status = 1
		result.Msg = "修改文档失败"
	} else {
		result.Msg = "修改文档成功"
		result.Data = contenttype
	}
	ctx.JSON(result)
}

// 删除文档
func ContentTypeDelete(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	result := AmisResult{}
	result.Status = 0
	result.Msg = "删除文档"
	if err := database.GetDb().XadminContenttype.DeleteOneID(id).Exec(ctx.Request().Context()); err != nil {
		result.Status = 1
		result.Msg = "删除文档失败"
	}
	ctx.JSON(result)
}
