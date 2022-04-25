package xadminent

import (
	"github.com/wgbbiao/xadminent/database"
)

func SetDatabase(db *database.DatabaseConfig) {
	database.SetDatabase(db)
}
