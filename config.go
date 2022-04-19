package xadminent

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/wgbbiao/xadminent/ent"
)

//JwtKey JwtKey
var JwtKey string = "Auys7;fq272/csH6"

//JwtTimeOut jwt过期时间
var JwtTimeOut int64 = 86400

// JwtCheckFunc JwtCheckFunc
var JwtCheckFunc func(c iris.Context)

//Handle Handle
type Handle struct {
	Path   string
	Method []string
	Func   func(ctx iris.Context)
	Jwt    bool
	// Permissions []string //所需要的权限
}

// Config Config
type Config struct {
	ListField []string //列表页字段
	Model     interface{}
	// Layout     interface{} //表单排列
	PageSize        int //每页大小
	BeforeSave      func(obj interface{})
	BeforeListQuery func(query *map[string]string)
	//todo
	// Form interface{} //自定表单
	Sort          string
	DisableAction []string //禁止的操作 [create,update,detail,delete,list]
}

//HasPermission 检查是否有权限
func (c *Config) HasPermission(action string) bool {
	return true
}

//Title Title
func (c *Config) Title() string {
	return "Title"
}

//XadminConfig 配置
type XadminConfig struct {
	Models       map[string]Config
	Handles      []Handle
	IrisParty    iris.Party
	JwtCheckFunc func(c iris.Context)
}

//Xadmin 配置实例
var Xadmin *XadminConfig

//NewXadmin 新建配置
func NewXadmin() *XadminConfig {
	Xadmin = &XadminConfig{
		Models: make(map[string]Config),
	}
	return Xadmin
}

// //GetInfo 取得表结构（字段信息用于前台表单创建）
// func GetInfo(ctx iris.Context) {

// }

//Register Register
func (o *XadminConfig) Register(model interface{}, config Config) {
	resType := reflect.TypeOf(model)
	config.Model = model

	modelname := strings.Replace(resType.String(), "*", "", 1)
	modelname = strings.Replace(modelname, ".", "/", 1)
	fmt.Println(modelname)
	if config.Sort == "" {
		fmt.Println("a ...interface{}")
		// stmt := &gorm.Statement{DB: Db}
		// stmt.Parse(model)
		// config.Sort = fmt.Sprintf("-%s.id", stmt.Schema.Table)
	}
	o.Models[modelname] = config
}

//GetRegModels GetRegModels
func (o *XadminConfig) GetRegModels() (ms []interface{}) {
	for _, conf := range o.Models {
		ms = append(ms, conf.Model)
	}
	return
}

//RegisterView 注册新的api
func (o *XadminConfig) RegisterView(handle ...Handle) {
	o.Handles = append(o.Handles, handle...)
}

//SetIris 设置http
func (o *XadminConfig) SetIris(r iris.Party) {
	o.IrisParty = r
}

//SetJwtCheck 设置jwt检查函数
func (o *XadminConfig) SetJwtCheck(f func(c iris.Context)) {
	o.JwtCheckFunc = f
}

//Init Init
func (o *XadminConfig) Init() {
	// JwtCheckFunc = CheckJWTAndSetUser
	o.initUser()
	// initValidator()

	if o.JwtCheckFunc == nil {
		o.JwtCheckFunc = CheckJWTAndSetUser
	}
	for _, handel := range o.Handles {
		for _, method := range handel.Method {
			if handel.Jwt {
				o.IrisParty.Handle(method, handel.Path, o.JwtCheckFunc, handel.Func)
			} else {
				o.IrisParty.Handle(method, handel.Path, handel.Func)
			}
		}
	}

	o.IrisParty.Get("/refreshjwt", o.JwtCheckFunc, RefreshJwt)
	// o.IrisParty.Get("/{model:string  min(3)}/{table:string  min(3)}", o.JwtCheckFunc, ListHandel)
	// o.IrisParty.Get("/{model:string  min(3)}/{table:string  min(3)}/{id:int}", o.JwtCheckFunc, DetailHandel)
	// o.IrisParty.Put("/{model:string  min(3)}/{table:string  min(3)}/{id:int}", o.JwtCheckFunc, UpdateHandel)
	// o.IrisParty.Put("/{model:string  min(3)}/{table:string  min(3)}", o.JwtCheckFunc, BatchUpdateHandel)
	// o.IrisParty.Post("/{model:string  min(3)}/{table:string  min(3)}", o.JwtCheckFunc, PostHandel)
	// o.IrisParty.Delete("/{model:string  min(3)}/{table:string  min(3)}/{id:int}", o.JwtCheckFunc, DeleteHandel)
	// o.IrisParty.Delete("/{model:string  min(3)}/{table:string  min(3)}", o.JwtCheckFunc, BatchDeleteHandel)
}

//GetConfig 取得配置文件
func GetConfig(model, table string) Config {
	return Xadmin.Models[fmt.Sprintf("%s/%s", model, table)]
}

func (o *XadminConfig) initUser() {
	// o.RegisterView(
	// 	Handle{
	// 		Path:   "/login",
	// 		Method: []string{iris.MethodPost},
	// 		Func:   Login,
	// 		Jwt:    false,
	// 	},
	// 	Handle{
	// 		Path:   "/info",
	// 		Method: []string{iris.MethodGet},
	// 		Func:   GetInfo,
	// 		Jwt:    true,
	// 	},
	// 	Handle{
	// 		Path:   "/changepassword",
	// 		Method: []string{iris.MethodPost},
	// 		Func:   ChangePassword,
	// 		Jwt:    true,
	// 	},
	// )

	o.Register(&ent.User{}, Config{
		BeforeSave: func(obj interface{}) {
			pointer := reflect.ValueOf(obj)
			m := pointer.MethodByName("SetPassword")
			args := []reflect.Value{}
			m.Call(args)
		},
	})
	o.Register(&ent.Permission{}, Config{})
	o.Register(&ent.Role{}, Config{})
	o.Register(&ent.ContentType{}, Config{})
}
