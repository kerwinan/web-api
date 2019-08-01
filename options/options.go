package options

import (
	"github.com/astaxie/beego"
	"github.com/dlintw/goconf"
)

var (
	cfg     *goconf.ConfigFile
	cfgPath = "./conf/app.ini"
)

func Init() *goconf.ConfigFile {
	cfg, err := goconf.ReadConfigFile(cfgPath)
	if err != nil {
		beego.Error("ReadConfigFile err:", err.Error())
	}
	return cfg
}
