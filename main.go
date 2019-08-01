package main

import (
	"dana-tech.com/orm"
	"dana-tech.com/web-api/options"
	_ "dana-tech.com/web-api/routers"
	"github.com/astaxie/beego"
)

func main() {
	cfg := options.Init()
	orm.Init(cfg)
	// rows, _ := orm.Engine.Query("select * from userinfo where username=?", username)
	// for _, row := range rows {
	// 	Uid := string(row["uid"])
	// 	Username := string(row["username"])
	// 	Password := string(row["password"])
	// 	beego.Info(Uid, Username, Password)
	// }

	beego.Run()
}
