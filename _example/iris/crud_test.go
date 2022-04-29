package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/wgbbiao/xadminent"
	"github.com/wgbbiao/xadminent/api"
	"github.com/wgbbiao/xadminent/database"
	"github.com/wgbbiao/xadminent/ent"
)

func TestCrud(t *testing.T) {
	xadminent.SetDatabase(&database.DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Database: "testent",
	})
	client := database.GetDb()
	u, _ := client.XadminUser.Get(context.Background(), 9)
	fmt.Println(u.ID)
	fmt.Println(api.CheckPermission(u, ent.XadminUser{}, "add"))
}
