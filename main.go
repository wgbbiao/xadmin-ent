package xadminent

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/wgbbiao/xadminent/api"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
	"golang.org/x/term"
)

// 创建管理员
func CreateAdmin() (*ent.XadminUser, error) {
	client := database.GetDb()
	fmt.Print("请输入管理员账号:")
	var username string
	fmt.Scanln(&username)
	fmt.Print("\n请输入管理员密码:")
	password, err := term.ReadPassword(0)
	if err != nil {
		return nil, err
	}
	fmt.Print("\n请再次输入管理员密码:")
	password2, err := term.ReadPassword(0)
	if err != nil {
		return nil, err
	}
	if string(password) != string(password2) {
		return nil, fmt.Errorf("两次输入的密码不一致")
	}
	// 随机生成一个密码盐
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return nil, err
	}
	pw := api.MD5(string(password) + string(salt))
	return client.XadminUser.Create().
		SetUsername(username).
		SetSalt(string(salt)).
		SetIsSuper(true).
		SetPassword(pw).Save(context.Background())
}
