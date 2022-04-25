package xadminent

import (
	"context"
	"fmt"
	"time"

	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent/migrate"
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
}
