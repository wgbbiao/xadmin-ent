package xadminent

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/wgbbiao/xadminent/api"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
	"github.com/wgbbiao/xadminent/ent/migrate"
	"github.com/wgbbiao/xadminent/ent/xadmincontenttype"
	"github.com/wgbbiao/xadminent/ent/xadminpermission"
)

func AutoMigrate() {
	client := database.GetDb()
	fmt.Println("开始迁移数据库...")
	beginAt := time.Now()
	err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithFixture(true),
	)
	endAt := time.Now()

	if err != nil {
		fmt.Println("迁移数据库失败:", err)
	} else {
		fmt.Println("数据库迁移完成,耗时:", endAt.Sub(beginAt))
	}
	// 创建权限
	CreatePermission("xadmin", []any{
		ent.XadminRole{},
		ent.XadminUser{},
	})
}

// 创建文档类型
func CreateDocType(appLabel, model string) *ent.XadminContenttype {
	client := database.GetDb()
	// 查询是否存在
	contentType, err := client.XadminContenttype.Query().
		Where(xadmincontenttype.AppLabel(appLabel), xadmincontenttype.Model(model)).
		First(context.Background())
	if err != nil {
		contentType, err = client.XadminContenttype.Create().
			SetModel(model).
			SetAppLabel(appLabel).
			Save(context.Background())
		if err != nil {
			panic(err)
		}
	}
	return contentType
}

// 根据models创建权限
func CreatePermission(appLabel string, models []any) {
	client := database.GetDb()
	// 权限列表
	codes := []string{
		api.CanView,
		api.CanAdd,
		api.CanEdit,
		api.CanDel,
	}
	for _, model := range models {
		rt := reflect.TypeOf(model)
		ct := CreateDocType(appLabel, rt.Name())
		fmt.Println("创建权限:", rt.Name())
		for _, code := range codes {
			// 查询是否存在
			name := fmt.Sprintf("Can %s %s", code, rt.Name())
			code := strings.ToLower(fmt.Sprintf("%s_%s", rt.Name(), code))
			cnt, _ := client.XadminPermission.Query().
				Where(xadminpermission.Code(code), xadminpermission.Name(name)).Count(context.Background())
			if cnt == 0 {
				client.XadminPermission.Create().
					SetName(name).
					SetCode(code).
					SetContentType(ct).
					Save(context.Background())
			}
		}
	}
}
