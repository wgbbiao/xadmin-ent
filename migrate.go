package xadminent

import (
	"context"
	"fmt"

	"github.com/wgbbiao/xadminent/ent/migrate"
)

func AutoMigrate() {
	client := GetDb()
	err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithFixture(true),
	)
	fmt.Println(err)
}
