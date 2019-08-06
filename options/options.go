package options

import (
	"dana-tech.com/web-api/global"
	"github.com/astaxie/beego"
	"github.com/dlintw/goconf"
)

var (
	cfgPath = "./conf/app.ini"
)

func Init() {
	cfg, err := goconf.ReadConfigFile(cfgPath)
	if err != nil {
		beego.Error("ReadConfigFile err:", err.Error())
	}
	global.Cfg = cfg
}
