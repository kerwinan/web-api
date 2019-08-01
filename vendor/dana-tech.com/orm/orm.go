package orm

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/dlintw/goconf"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

// type OrmOpt struct {
// 	Type    string `json:"type" yaml:"type"`
// 	User    string `json:"user" yaml:"user"`
// 	Pass    string `json:"pass" yaml:"pass"`
// 	Host    string `json:"host" yaml:"host"`
// 	Port    int    `json:"port" yaml:"port"`
// 	Name    string `json:"name" yaml:"name"`
// 	Charset string `json:"charset" yaml:"charset"`
// }

func Init(conf *goconf.ConfigFile) {
	user, _ := conf.GetString("mysql", "user")
	pwd, _ := conf.GetString("mysql", "pwd")
	host, _ := conf.GetString("mysql", "host")
	port, _ := conf.GetString("mysql", "port")
	dbname, _ := conf.GetString("mysql", "dbname")
	charset, _ := conf.GetString("mysql", "charset")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pwd, host, port, dbname, charset)
	beego.Info("url: ", url)
	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		beego.Error("NewEngine err: ", err.Error())
	}
	Engine = engine
}
