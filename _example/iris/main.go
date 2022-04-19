package main

import (
	"github.com/wgbbiao/xadminent"
)

func main() {
	xadminent.SetDatabase(&xadminent.DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Database: "testent",
	})
	xadminent.AutoMigrate()
}
