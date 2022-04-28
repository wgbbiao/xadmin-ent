package api

import (
	"crypto/md5"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/unknwon/com"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
	"github.com/wgbbiao/xadminent/ent/xadminuser"
	"github.com/wgbbiao/xadminent/j"
)

//RefreshJwt 刷新jwt
func RefreshJwt(c iris.Context) {
	u := c.Values().Get("u").(*ent.XadminUser)
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
	u, err := database.GetDb().XadminUser.
		Query().
		Where(xadminuser.Username(form.Username)).
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
	u, err := getUserFromJwt(ctx)

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

// 用户列表
func UserList(ctx iris.Context) {
	var query struct {
		Page     int    `url:"page"`
		Limit    int    `url:"limit"`
		Username string `url:"username"`
		IsSuper  *bool  `url:"is_super"`
	}
	result := AmisResult{}
	if err := ctx.ReadQuery(&query); err != nil {
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
	q := database.GetDb().XadminUser.
		Query().
		Offset((query.Page - 1) * query.Limit).
		Limit(query.Limit)

	if query.Username != "" {
		q = q.Where(xadminuser.UsernameContains(query.Username))
	}
	if query.IsSuper != nil {
		q = q.Where(xadminuser.IsSuper(true))
	}

	users, err := q.All(ctx.Request().Context())
	cnt, _ := q.Count(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "用户不存在"
		ctx.JSON(result)
		return
	}
	result.Status = 0
	result.Msg = "用户列表"
	result.Data = iris.Map{
		"items": users,
		"total": cnt,
	}
	ctx.JSON(result)
}

func getUserFromJwt(ctx iris.Context) (*ent.XadminUser, error) {
	claims := jwt.Get(ctx).(*j.FooClaims)
	return database.GetDb().
		XadminUser.Query().
		Where(xadminuser.ID(claims.Uid)).
		First(ctx.Request().Context())
}

// 添加用户
func UserAdd(ctx iris.Context) {
	var form struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		IsSuper  bool   `json:"is_super"`
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
	{
		u, err := database.GetDb().XadminUser.
			Query().
			Where(xadminuser.Username(form.Username)).
			First(ctx.Request().Context())
		if err == nil && u.ID != 0 {
			result.Status = 300
			result.Msg = "用户已存在"
			result.Data = u
			// ctx.StatusCode(iris.StatusBadGateway)
			ctx.JSON(result)
			return
		}
	}
	salt := com.RandomCreateBytes(16)
	u, err := database.GetDb().XadminUser.
		Create().
		SetUsername(form.Username).
		SetSalt(string(salt)).
		SetPassword(MD5(form.Password + string(salt))).
		SetIsSuper(form.IsSuper).
		Save(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "添加用户失败"
		ctx.JSON(result)
		return
	}
	result.Status = 0
	result.Msg = "添加用户成功"
	result.Data = u
	ctx.JSON(result)
}

// 用户详情
func UserDetail(ctx iris.Context) {
	id := ctx.Params().GetIntDefault("id", 0)
	result := AmisResult{}
	if id == 0 {
		result.Status = 1
		result.Msg = "用户不存在"
	} else {
		u, err := database.GetDb().XadminUser.
			Query().
			Where(xadminuser.ID(id)).
			First(ctx.Request().Context())
		if err != nil {
			result.Status = 1
			result.Msg = "用户不存在"
		} else {
			result.Status = 0
			result.Msg = "用户详情"
			result.Data = u
		}
	}
	ctx.JSON(result)
}

// 用户编辑
func UserEdit(ctx iris.Context) {
	var form struct {
		ID       int    `json:"id" validate:"required"`
		Username string `json:"username" validate:"required"`
		IsSuper  bool   `json:"is_super"`
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
	u, err := database.GetDb().XadminUser.
		Query().
		Where(xadminuser.ID(form.ID)).
		First(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "用户不存在"
		ctx.JSON(result)
		return
	}

	q := database.GetDb().XadminUser.Update()
	if form.Username != "" {
		q.SetUsername(form.Username)
	}
	if form.IsSuper != u.IsSuper {
		q.SetIsSuper(form.IsSuper)
	}
	_, err = q.Where(xadminuser.ID(form.ID)).Save(ctx.Request().Context())
	if err != nil {
		if sqlgraph.IsUniqueConstraintError(err) {
			result.Status = 300
			result.Msg = "用户已存在"
			result.Data = u
			ctx.JSON(result)
			return
		}
		result.Status = 1
		result.Msg = "编辑用户失败"
		ctx.JSON(result)
		return
	}
	result.Status = 0
	result.Msg = "用户编辑成功"
	ctx.JSON(result)
}

// 修改密码
func UserPassword(ctx iris.Context) {
	var form struct {
		ID                   int    `json:"id" validate:"required"`
		Password             string `json:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required"`
	}
	result := AmisResult{}
	if err := ctx.ReadJSON(&form); err != nil {
		fmt.Println(err)
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
	if form.Password != form.PasswordConfirmation {
		result.Status = 1
		result.Msg = "两次密码不一致"
		ctx.JSON(result)
		return
	}
	u, err := database.GetDb().XadminUser.
		Query().
		Where(xadminuser.ID(form.ID)).
		First(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "用户不存在"
		ctx.JSON(result)
		return
	}

	salt := com.RandomCreateBytes(16)
	_, err = database.GetDb().XadminUser.
		Update().
		SetPassword(MD5(form.Password + string(salt))).
		SetSalt(string(salt)).
		Where(xadminuser.ID(form.ID)).
		Save(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "修改密码失败"
		ctx.JSON(result)
		return
	}
	result.Data = u
	result.Status = 0
	result.Msg = "修改密码成功"
	ctx.JSON(result)
}

// 删除用户
func UserDelete(ctx iris.Context) {
	var form struct {
		ID int `json:"id" validate:"required"`
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
	_, err := database.GetDb().XadminUser.
		Query().
		Where(xadminuser.ID(form.ID)).
		First(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "用户不存在"
		ctx.JSON(result)
		return
	}

	err = database.GetDb().XadminUser.
		DeleteOneID(form.ID).Exec(ctx.Request().Context())
	if err != nil {
		result.Status = 1
		result.Msg = "删除用户失败"
		ctx.JSON(result)
		return
	}
	// result.Data = u
	result.Status = 0
	result.Msg = "删除用户成功"
	ctx.JSON(result)
}
