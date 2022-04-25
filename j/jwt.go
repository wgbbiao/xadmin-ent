package j

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/wgbbiao/xadminent/ent"
)

//JwtKey JwtKey
var JwtKey string = "Auys7;fq272/csH6"

//JwtTimeOut jwt过期时间
var JwtTimeOut int64 = 86400

// JwtCheckFunc JwtCheckFunc
var JwtCheckFunc func(c iris.Context)

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

type FooClaims struct {
	Uid int `json:"uid"`
}

func GetToken(u *ent.XadminUser) (tokenString string) {
	signer := getSigner()

	foo := FooClaims{
		Uid: u.ID,
	}

	token, _ := signer.Sign(foo)
	tokenString = string(token)
	return
}

func VerifyMiddleware() context.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, JwtKey)
	// Enable server-side token block feature (even before its expiration time):
	verifier.WithDefaultBlocklist()
	// Enable payload decryption with:
	// verifier.WithDecryption(encKey, nil)
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(FooClaims)
	})
	return verifyMiddleware
}
