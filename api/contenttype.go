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
