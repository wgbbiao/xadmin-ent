package xadminent

import (
	"context"

	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent/migrate"
)

func AutoMigrate() {
	client := database.GetDb()
	client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithFixture(true),
	)
}
