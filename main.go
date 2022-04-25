package xadminent

import (
	"context"
	"fmt"

	"github.com/unknwon/com"
	"github.com/wgbbiao/xadminent/api"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
	"github.com/wgbbiao/xadminent/ent/xadminuser"
	"golang.org/x/term"
)

// 创建管理员
func CreateAdmin() (*ent.XadminUser, error) {
	client := database.GetDb()
	var username string
	for {
		fmt.Print("请输入管理员账号:")
		fmt.Scanln(&username)
		cnt, _ := client.XadminUser.Query().Where(xadminuser.Username(username)).Count(context.Background())
		if cnt > 0 {
			fmt.Println("管理员账号已存在,请重新输入！")
			continue
		}
		break
	}
	fmt.Print("请输入管理员密码:")
	password, err := term.ReadPassword(0)
	if err != nil {
		return nil, err
	}
	for {
		fmt.Print("\n请再次输入管理员密码:")
		password2, err := term.ReadPassword(0)
		if err != nil {
			return nil, err
		}
		fmt.Println("")
		if string(password) != string(password2) {
			fmt.Println("两次输入的密码不一致,请重新输入!")
		} else {
			break
		}
	}
	// 随机生成长度为16的字符串
	salt := com.RandomCreateBytes(16)
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
