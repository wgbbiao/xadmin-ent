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

func getSigner() *jwt.Signer {
	signer := jwt.NewSigner(jwt.HS256, JwtKey, time.Second*time.Duration(JwtTimeOut))
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
