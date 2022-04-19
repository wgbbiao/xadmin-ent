package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wgbbiao/xadminent"
)

func main() {
	xadmin := xadminent.NewXadmin(&xadminent.DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "123456",
		Database: "testent",
	})
	xadmin.Init()
	xadminent.AutoMigrate()

	client := xadminent.GetDb()

	ct, _ := client.ContentType.Create().SetAppLabel("dd").SetModel("mode").Save(context.Background())

	p, err := client.Permission.Create().
		SetCode("code").
		SetContentTypeID(ct.ID).
		SetName("name").Save(context.Background())

	log.Println("------")
	log.Println(err)
	ud := client.User
	u, err := ud.Create().SetUsername("123").
		SetPassword("qwe").
		SetSalt("123qwe").
		SetIsSuper(true).
		SetLastLoginAt(time.Now()).
		AddPermissionIDs(p.ID).
		Save(context.Background())
	fmt.Println(u, err)
	fmt.Println(ud)
	// v := reflect.ValueOf(*client)
	// fmt.Println(v.NumField())
	// v2 := v.FieldByName("User")

	// m := v2.MethodByName("Create")
	// v3 := m.Call([]reflect.Value{})[0]
	// m1 := v3.MethodByName("SetUsername")
	// m1.Call([]reflect.Value{reflect.ValueOf("123222222")})
	// {
	// 	m := v3.MethodByName("SetPassword")
	// 	m.Call([]reflect.Value{reflect.ValueOf("qwe")})
	// }
	// {
	// 	m := v3.MethodByName("SetSalt")
	// 	m.Call([]reflect.Value{reflect.ValueOf("123qwe")})
	// }
	// {
	// 	m := v3.MethodByName("SetIsSuper")
	// 	m.Call([]reflect.Value{reflect.ValueOf(true)})
	// }
	// {
	// 	m := v3.MethodByName("SetLastLoginAt")
	// 	m.Call([]reflect.Value{reflect.ValueOf(time.Now())})
	// }
	// m2 := v3.MethodByName("Save")
	// m2.Call([]reflect.Value{reflect.ValueOf(context.Background())})

}
