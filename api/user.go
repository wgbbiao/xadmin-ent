package api

import (
	"crypto/md5"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
	"github.com/wgbbiao/xadminent/ent/user"
	"github.com/wgbbiao/xadminent/j"
)

//RefreshJwt 刷新jwt
func RefreshJwt(c iris.Context) {
	u := c.Values().Get("u").(*ent.User)
	tokenString := j.GetToken(u)
	c.JSON(iris.Map{
		"status": 0,
		"data": iris.Map{
			"token": tokenString,
			// "username": u.Username,
		},
	})
}

// md5 加密
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 管理员登录
func Login(ctx iris.Context) {
	var form struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	result := AmisResult{}
	if err := ctx.ReadJSON(&form); err != nil {
		result.Status = 1
		if errs, ok := err.(validator.ValidationErrors); ok {
			result.Msg = "表单验证错误"
			validationErrors := wrapValidationErrors(errs)

			result.Data = iris.NewProblem().
				Title("Validation error").
				Detail("One or more fields failed to be validated").
				Type("/user/validation-errors").
				Key("errors", validationErrors)
		} else {
			result.Msg = "表单读取错误"
		}
		ctx.JSON(result)
		return
	}
	u, err := database.GetDb().User.
		Query().
		Where(user.Username(form.Username)).
		First(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "用户不存在"
		ctx.JSON(result)
		return
	}
	if u.Password != MD5(form.Password+u.Salt) {
		result.Status = 1
		result.Msg = "用户名或密码错误"
		ctx.JSON(result)
		return
	}
	result.Status = 0
	result.Msg = "登录成功"
	result.Data = iris.Map{
		"token": j.GetToken(u),
	}
	ctx.JSON(result)
}

// 用户信息
func UserInfo(ctx iris.Context) {
	u, err := getUser(ctx)

	if err != nil {
		ctx.JSON(iris.Map{
			"status": 1,
			"msg":    "用户不存在",
		})
		return
	}
	ctx.JSON(iris.Map{
		"status": 0,
		"data":   u,
	})
}

func getUser(ctx iris.Context) (*ent.User, error) {
	claims := jwt.Get(ctx).(*j.FooClaims)
	return database.GetDb().
		User.Query().
		Where(user.ID(claims.Uid)).
		First(ctx.Request().Context())
}
