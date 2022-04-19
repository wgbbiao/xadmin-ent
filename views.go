package xadminent

import (
	"reflect"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/wgbbiao/xadminent/ent"
)

//CheckJWTAndSetUser 检查jwt并把User放到Values
func CheckJWTAndSetUser(ctx iris.Context) {
	// if err := myJwtMiddleware.CheckJWT(ctx); err != nil {
	// 	myJwtMiddleware.Config.ErrorHandler(ctx, err)
	// 	return
	// }
	// // If everything ok then call next.
	// if ctx.GetStatusCode() != iris.StatusUnauthorized {
	// 	var u User
	// 	x, _ := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	// 	if rt := u.GetUserByID(int(x["uid"].(float64))); !(rt.Error == gorm.ErrRecordNotFound) && rt.Error == nil {
	// 		bl := true
	// 		if ctx.Params().Get("model") != "" {
	// 			config := GetConfig(ctx.Params().Get("model"), ctx.Params().GetString("table"))
	// 			bl = HasPermissionForModel(&u, config.Model, GetActionByMethod(ctx.Method()))
	// 		}
	// 		if bl {
	// 			ctx.Values().Set("u", u)
	// 			ctx.Next()
	// 		} else {
	// 			ctx.JSON(iris.Map{
	// 				"status": HTTPForbidden,
	// 				"msg":    UserNoPermission,
	// 			})
	// 		}
	// 	} else {
	// 		ctx.JSON(iris.Map{
	// 			"status": 1,
	// 			"msg":    UserDoesNotExist,
	// 		})
	// 	}
	// }
}

func getSigner() *jwt.Signer {
	signer := jwt.NewSigner(jwt.HS256, JwtKey, 10*time.Minute)
	return signer
}

type fooClaims struct {
	Uid int `json:"uid"`
}

func getToken(_ iris.Context, u ent.User) (tokenString string) {
	signer := getSigner()

	foo := fooClaims{
		Uid: u.ID,
	}

	token, _ := signer.Sign(foo)
	tokenString = string(token)
	return
}

//RefreshJwt 刷新jwt
func RefreshJwt(c iris.Context) {
	u := c.Values().Get("u").(ent.User)
	tokenString := getToken(c, u)
	c.JSON(iris.Map{
		"status": 0,
		"data": iris.Map{
			"token": tokenString,
			// "username": u.Username,
		},
	})
}

// 列表视图
func ListView(ctx iris.Context) {
	// config := GetConfig(ctx.Params().Get("model"), ctx.Params().Get("table"))

	client := GetDb()
	v := reflect.ValueOf(client)
	v2 := v.FieldByName("User")
	v2.MethodByName("Create").Call([]reflect.Value{})

}
