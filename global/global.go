package global

import (
	"github.com/dlintw/goconf"
)

var (
	NodeID int64
	Cfg    *goconf.ConfigFile
)

func Init(conf *goconf.ConfigFile) {
	node, _ := conf.GetInt("global", "node_id")
	NodeID = int64(node)
}
